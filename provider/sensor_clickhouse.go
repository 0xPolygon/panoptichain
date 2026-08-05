package provider

import (
	"container/list"
	"context"
	"errors"
	"math"
	"math/big"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/0xPolygon/polygon-cli/p2p/database"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/network"
	"github.com/0xPolygon/panoptichain/observer"
)

const (
	// headWindow bounds how far back the head-block query looks. It must comfortably
	// exceed the polling interval so a slow cycle never sees an empty window.
	headWindow = 5 * time.Minute

	// reorgClockSkewAllowance is how far below the reorg watermark each poll
	// re-reads. The watermark is the global max detected_at of the last batch,
	// while a deepening of an OLDER start height can carry a timestamp below it
	// (writer clock skew, out-of-order arrival) -- without the allowance such a
	// detection is never seen again. publishedReorgDepth keeps the re-read window
	// from re-publishing.
	reorgClockSkewAllowance = 5 * time.Minute

	// maxReorgsPerPoll caps the reorg backlog a single cycle will pull. The
	// watermark query is otherwise unbounded, so a long outage would return every
	// reorg since it began in one go.
	maxReorgsPerPoll = 1000

	// maxPublishedReorgDepths bounds the dedup map.
	maxPublishedReorgDepths = 4096
)

// ClickHouseSensorNetworkProvider reads the same sensor data as
// SensorNetworkProvider but from ClickHouse instead of GCP Datastore. It embeds
// SensorNetworkProvider to reuse the shared, backend-agnostic behavior
// (PublishEvents, stolen-block/bogon detection, block-buffer bounds) and only
// overrides the data-fetching path.
type ClickHouseSensorNetworkProvider struct {
	*SensorNetworkProvider
	conn driver.Conn

	// covered is the resume watermark: every height below it has been read in a
	// CLEAN pass, blocks and events both. It advances only after a whole cycle
	// commits, so a failed poll leaves it in place and the next poll re-covers the
	// same range. It may move DOWN, when a bogon leaves the head window.
	//
	// publishedThrough is the high-water mark of what has been handed to the
	// observers, and it is monotonic BY CONSTRUCTION -- advance() is the only writer
	// and it never lowers it. That is the property the Datastore path gets from
	// `if blockNumber > prevBlockNumber`, which the first separation of these roles
	// dropped: covered going down while the buffer still held higher blocks let the
	// publish filter un-skip them, republishing a block within one poll. Grouped by
	// height with one signer, that reads as a double sign, and DoubleSignObserver
	// raised sensor_double_sign against a real validator address.
	covered          uint64
	publishedThrough uint64

	// publishedReorgDepth remembers the deepest depth already published per start
	// height, so re-reading the clock-skew allowance window cannot re-publish a
	// detection while a genuine deepening (higher depth) still goes out.
	publishedReorgDepth map[uint64]uint32
}

func NewClickHouseSensorNetworkProvider(n network.Network, eb *observer.EventBus, cfg config.SensorNetworkClickHouse) *ClickHouseSensorNetworkProvider {
	logger := NewLogger(n, cfg.Label)

	base := &SensorNetworkProvider{
		network:          n,
		label:            cfg.Label,
		bus:              eb,
		interval:         GetInterval(cfg.Interval),
		logger:           logger,
		blocks:           list.New(),
		latestReorgTime:  time.Now(),
		refreshStateTime: new(time.Duration),
	}

	opts, err := clickhouse.ParseDSN(cfg.DSN)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to parse ClickHouse DSN")
		return &ClickHouseSensorNetworkProvider{SensorNetworkProvider: base, publishedReorgDepth: map[uint64]uint32{}}
	}

	conn, err := clickhouse.Open(opts)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to connect to ClickHouse")
		return &ClickHouseSensorNetworkProvider{SensorNetworkProvider: base, publishedReorgDepth: map[uint64]uint32{}}
	}

	return &ClickHouseSensorNetworkProvider{SensorNetworkProvider: base, conn: conn, publishedReorgDepth: map[uint64]uint32{}}
}

// RefreshState mirrors SensorNetworkProvider.RefreshState but fetches from
// ClickHouse. refreshStolenBlocks is reused unchanged from the embedded type.
func (s *ClickHouseSensorNetworkProvider) RefreshState(ctx context.Context) error {
	defer timer(s.refreshStateTime)()

	if s.conn == nil {
		return errors.New("ClickHouse client not initialized")
	}

	if err := s.refreshBlockBuffer(ctx); err != nil {
		s.logger.Error().Err(err).Msg("Failed to refresh block buffer")
	}

	if err := s.refreshReorgs(ctx); err != nil {
		s.logger.Error().Err(err).Msg("Failed to refresh reorg")
	}

	if err := s.refreshStolenBlocks(); err != nil {
		s.logger.Error().Err(err).Msg("Failed to refresh stolen blocks")
	}

	return nil
}

// addSaturating adds without wrapping. Every operand here derives from a
// peer-announced block_number (UInt64 in the schema), so an announcement anywhere in
// the top 512 of the range used to wrap the bound to near zero -- which stranded the
// watermark below the real head, made maxIf return 0 on every subsequent poll, and
// left the walk as the only recovery: ~175,000 polls, about ten days of silence.
func addSaturating(a, b uint64) uint64 {
	if a > math.MaxUint64-b {
		return math.MaxUint64
	}
	return a + b
}

// publish sets the range the observers will be shown, and is the ONLY writer of
// publishedThrough. It cannot lower it: PublishEvents both filters on it
// (block.Number() < prev is skipped) and reports it as SensorBlocks.Start, so
// lowering it re-publishes blocks still sitting in the buffer.
func (s *ClickHouseSensorNetworkProvider) publish(from uint64) {
	if from > s.publishedThrough {
		s.publishedThrough = from
	}
	s.prevBlockNumber = s.publishedThrough
}

func (s *ClickHouseSensorNetworkProvider) refreshBlockBuffer(ctx context.Context) error {
	s.blockEvents = nil

	head, err := s.queryHead(ctx)
	if err != nil {
		return err
	}
	s.blockNumber = head

	if s.covered == 0 {
		// First successful poll: nothing behind us to fill. Start covering from here.
		s.covered = head
		s.publish(head)
		return nil
	}

	// clampStart bounds the range to one buffer behind the head, logging any skip.
	start := s.clampStart(s.covered)
	if start >= head {
		// Nothing new, or the head dropped below the watermark because a bogon left
		// the window -- snapping covered down IS the recovery. publishedThrough does
		// not follow it down, so the blocks still buffered above the new head are not
		// re-published; they were published when they were read.
		s.covered = head
		s.publish(start)
		return nil
	}

	// The cycle is transactional: read both passes to the side, commit only if both
	// were clean. Committing blocks and then failing on events would advance
	// nothing visible yet still leave duplicates in the buffer when the range is
	// retried; committing neither makes the retry exact. On failure prevBlockNumber
	// is parked at the head so PublishEvents' filter skips everything this cycle --
	// the covered watermark has not moved, so the next poll re-covers the range,
	// which is what the previous version's comments claimed and its code did not do
	// (it advanced the watermark before the first query, so a failed poll lost its
	// range forever).
	blocks, blockTimes, err := s.readBlocks(ctx, start, head)
	if err != nil {
		s.publish(head)
		return err
	}
	events, err := s.readBlockEvents(ctx, start, head, blockTimes)
	if err != nil {
		s.publish(head)
		return err
	}

	for _, block := range blocks {
		if s.blocks.Len() >= blockBufferSize {
			s.blocks.Remove(s.blocks.Front())
		}
		s.blocks.PushBack(block)
	}
	s.blockEventsLock.Lock()
	s.blockEvents = events
	s.blockEventsLock.Unlock()

	// PublishEvents publishes [prevBlockNumber, blockNumber) -- the range committed
	// above, clamped by publishedThrough so nothing already shown goes out twice.
	s.covered = head
	s.publish(start)
	return nil
}

// queryHead derives the fleet head from recent announcements, refusing implausible
// heights in the query itself rather than clamping after the fact.
//
// block_number is peer-announced and written as received, so max() alone hands a
// single bogon announcement control of the head -- and with it the size of the
// range the fork observer iterates per height (~32ns each; a bogon at 2^63 parks
// that goroutine ~9,000 years from returning). The bound excludes anything more
// than a buffer above the watermark, so a mid-life bogon never enters the
// arithmetic at all and publishing continues underneath it. Two cases remain:
//
//   - First poll: no watermark to bound against, so a bogon can take the head. But
//     the head is assigned fresh from this query every poll, never maintained as a
//     running max, so when the bogon ages out of the window the head snaps back and
//     the watermark follows. Damage is bounded by the window, not permanent.
//   - Every event in the window above the bound: either the fleet is recovering
//     from an outage and the real head genuinely jumped, or the window holds only
//     bogons. The two are indistinguishable, so advance one buffer and walk: each
//     walked band is genuinely filled, so a real gap converges at a buffer per
//     poll, and a pure-bogon window stops mattering the moment the bogon expires.
func (s *ClickHouseSensorNetworkProvider) queryHead(ctx context.Context) (uint64, error) {
	bound := uint64(math.MaxUint64)
	if s.covered != 0 {
		bound = addSaturating(s.covered, blockBufferSize)
	}

	var rawMax, head uint64
	if err := s.conn.QueryRow(ctx, `
		SELECT max(block_number), maxIf(block_number, block_number < ?)
		FROM block_events
		WHERE seen_at > now() - INTERVAL ? SECOND`,
		bound, int(headWindow.Seconds())).Scan(&rawMax, &head); err != nil {
		return 0, err
	}
	if rawMax == 0 {
		// No events in the window: the fleet is down or not yet writing. Leave the
		// previous head in place rather than resetting it to zero.
		return 0, errors.New("no block events in the recent window")
	}
	if rawMax >= bound {
		s.logger.Warn().
			Uint64("max_announced", rawMax).
			Uint64("bound", bound).
			Uint64("watermark", s.covered).
			Msg("Ignoring implausibly high announced heights")
	}
	if head == 0 {
		head = addSaturating(s.covered, blockBufferSize)
	}
	return head, nil
}

// readBlocks reads the header facts for [start, end) without touching provider
// state, so a failed cursor commits nothing and the range can be retried exactly.
func (s *ClickHouseSensorNetworkProvider) readBlocks(ctx context.Context, start, end uint64) (types.Blocks, map[string]time.Time, error) {
	s.logger.Debug().
		Uint64("start", start).
		Uint64("end", end).
		Msg("Reading sensor network block range")

	// End is exclusive so the sensors have a moment to finish writing the head.
	//
	// LIMIT 1 BY collapses rows that have not yet been merged away. blocks is a
	// ReplacingMergeTree whose rows for a hash are byte-identical, so without this
	// the same block can be returned more than once and downstream fork counting
	// (forks_per_block_number) reads the copies as competing blocks.
	rows, err := s.conn.Query(ctx, `
		SELECT hash, number, parent_hash, block_time, coinbase, difficulty,
		       gas_used, gas_limit, base_fee, uncle_hash, state_root, tx_root,
		       receipt_root, logs_bloom, extra_data, mix_digest, nonce
		FROM blocks
		WHERE number >= ? AND number < ?
		ORDER BY number
		LIMIT 1 BY number, hash`, start, end)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	// blockTimes feeds the events pass, which needs each block's header timestamp
	// to compute latency.
	blockTimes := make(map[string]time.Time)
	var blocks types.Blocks

	for rows.Next() {
		var (
			hash, parentHash, coinbase                string
			uncleHash, stateRoot, txRoot, receiptRoot string
			mixDigest, logsBloom, extraData           string
			number, difficulty, gasUsed, gasLimit     uint64
			nonce                                     uint64
			baseFee                                   big.Int
			blockTime                                 time.Time
		)

		if err := rows.Scan(&hash, &number, &parentHash, &blockTime, &coinbase, &difficulty,
			&gasUsed, &gasLimit, &baseFee, &uncleHash, &stateRoot, &txRoot,
			&receiptRoot, &logsBloom, &extraData, &mixDigest, &nonce); err != nil {
			s.logger.Warn().Err(err).Msg("Failed to scan block")
			continue
		}

		// Full header so downstream ecrecover (signer/bogon/stolen) works.
		//
		// logs_bloom and extra_data are stored as RAW BYTES in their String
		// columns, not as hex text (see the conventions in clickhouse_schema.sql),
		// so they are converted directly. Hex-decoding them here would corrupt
		// Extra and silently break every signer-derived metric.
		header := &types.Header{
			ParentHash:  common.HexToHash(parentHash),
			UncleHash:   common.HexToHash(uncleHash),
			Coinbase:    common.HexToAddress(coinbase),
			Root:        common.HexToHash(stateRoot),
			TxHash:      common.HexToHash(txRoot),
			ReceiptHash: common.HexToHash(receiptRoot),
			Bloom:       types.BytesToBloom([]byte(logsBloom)),
			Difficulty:  new(big.Int).SetUint64(difficulty),
			Number:      new(big.Int).SetUint64(number),
			GasLimit:    gasLimit,
			GasUsed:     gasUsed,
			Time:        uint64(blockTime.Unix()),
			Extra:       []byte(extraData),
			MixDigest:   common.HexToHash(mixDigest),
			Nonce:       types.EncodeNonce(nonce),
			BaseFee:     new(big.Int).Set(&baseFee),
		}
		blocks = append(blocks, types.NewBlockWithHeader(header))
		blockTimes[hash] = blockTime
	}

	if err := rows.Err(); err != nil {
		// A truncated cursor would leave blockTimes incomplete, and the events pass
		// drops every event whose block is missing from it -- a partial read here
		// silently discards events too. Nothing has been committed, so the caller
		// retries the whole range on the next poll.
		return nil, nil, err
	}

	return blocks, blockTimes, nil
}

// getBlockEvents loads every event for the block range in a single query.
//
// This used to issue one `WHERE block_hash = ?` query per block, fanned out over
// unbounded goroutines from inside the still-open blocks cursor -- up to 512
// concurrent point lookups against the largest table in the database after a
// stall. block_events now carries block_number as the leading sort key, so the
// whole range is one ordered scan instead.
func (s *ClickHouseSensorNetworkProvider) readBlockEvents(ctx context.Context, start, end uint64, blockTimes map[string]time.Time) ([]*observer.SensorBlockEvents, error) {
	if len(blockTimes) == 0 {
		return nil, nil
	}

	// Only propagation events: header, header_backfill and body are things the sensor
	// requested, so they carry no peer, and counting them would add an empty peer to
	// every unique-propagator metric. block_events_first is keyed by source and keeps
	// all of them, so the restriction has to be stated here -- the same one
	// v_block_latency applies.
	rows, err := s.conn.Query(ctx, `
		SELECT block_hash, sensor_id, node_id, seen_at
		FROM block_events
		WHERE block_number >= ? AND block_number < ?
		  AND source IN ('hash_announce', 'new_block')
		ORDER BY block_number`, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	byBlock := make(map[string][]database.DatastoreEvent, len(blockTimes))
	for rows.Next() {
		var blockHash, sensorID, nodeID string
		var seenAt time.Time
		if err := rows.Scan(&blockHash, &sensorID, &nodeID, &seenAt); err != nil {
			s.logger.Warn().Err(err).Msg("Failed to scan block event")
			continue
		}
		byBlock[blockHash] = append(byBlock[blockHash], database.DatastoreEvent{
			SensorId: sensorID,
			PeerId:   nodeID,
			Time:     seenAt,
		})
	}

	if err := rows.Err(); err != nil {
		// Nothing committed; the caller retries the whole range next poll.
		return nil, err
	}

	events := make([]*observer.SensorBlockEvents, 0, len(byBlock))
	for blockHash, blockEvents := range byBlock {
		blockTime, ok := blockTimes[blockHash]
		if !ok {
			// An event for a block whose header we have not stored (or that fell
			// outside the blocks query). Latency is measured against the header
			// timestamp, so there is nothing to compute without it.
			continue
		}
		// Only block.Time is consumed by BlockEventsObserver, so a minimal header
		// carrying the block timestamp is sufficient here.
		events = append(events, &observer.SensorBlockEvents{
			Block:  &database.DatastoreBlock{DatastoreHeader: &database.DatastoreHeader{Time: blockTime}},
			Events: blockEvents,
		})
	}
	return events, nil
}

func (s *ClickHouseSensorNetworkProvider) refreshReorgs(ctx context.Context) error {
	// reorg_detections is append-only: the same reorg is re-reported as it deepens,
	// and the deepest detection is the one worth alerting on. Collapsing to one row
	// per start height here is what the old ReplacingMergeTree(depth) was trying to
	// express -- but it belongs in the read, since an insert cannot suppress itself.
	//
	// That collapse is v_reorgs in the schema, so read the view rather than
	// restating it. A hand-rolled copy here aliased each aggregate to its own
	// source column name, which makes ClickHouse resolve the WHERE and the argMax
	// argument to the alias instead of the column -- two ILLEGAL_AGGREGATION
	// errors, so the query failed on every poll regardless of the data. The view
	// aggregates in a subquery under distinct names precisely to avoid that.
	//
	// The view also takes detected_at as argMax(detected_at, depth) rather than
	// max(detected_at), so every field describes the same detection; the copy
	// mixed the newest timestamp with the deepest detection's other fields.
	//
	// reorg_detections is written by the reorg-alerts job. On the local stack (and
	// any freshly-reset schema) it is empty until a reorg is detected, so the reorg
	// and stolen-block metrics stay at zero until then.
	rows, err := s.conn.Query(ctx, `
		SELECT
			start_block,
			depth,
			start_block_hash,
			end_block,
			end_block_hash,
			detected_at
		FROM v_reorgs
		WHERE detected_at > ?
		ORDER BY detected_at
		LIMIT ?`, s.latestReorgTime.Add(-reorgClockSkewAllowance), maxReorgsPerPoll)
	if err != nil {
		return err
	}
	defer rows.Close()

	var reorgs []*observer.DatastoreReorg
	readThrough := s.latestReorgTime
	rowsRead := 0
	for rows.Next() {
		var (
			startBlock, endBlock uint64
			depth                uint32
			startHash, endHash   string
			detectedAt           time.Time
		)
		if err := rows.Scan(&startBlock, &depth, &startHash, &endBlock, &endHash, &detectedAt); err != nil {
			s.logger.Warn().Err(err).Msg("Failed to scan reorg")
			continue
		}

		// The cursor advances on every row READ, not every row published. Advancing
		// it only on publication deadlocked the read: once maxReorgsPerPoll
		// already-published rows sat in one skew window, each poll pulled the same
		// page, deduped all of it, published nothing, and so never moved the cursor
		// past them -- every later reorg lost permanently, and stolen-block detection
		// with it, since refreshStolenBlocks walks s.reorgs.
		if detectedAt.After(readThrough) {
			readThrough = detectedAt
		}
		rowsRead++

		// The skew allowance re-reads a window below the cursor, so rows already
		// published come back; only a genuine deepening (greater depth for the same
		// start height) goes out again.
		if published, ok := s.publishedReorgDepth[startBlock]; ok && depth <= published {
			continue
		}
		s.publishedReorgDepth[startBlock] = depth

		t := detectedAt
		reorgs = append(reorgs, &observer.DatastoreReorg{
			Depth:      int(depth),
			Start:      int(startBlock),
			End:        int(endBlock),
			StartBlock: datastore.NameKey(database.BlocksKind, startHash, nil),
			EndBlock:   datastore.NameKey(database.BlocksKind, endHash, nil),
			Time:       &t,
		})
	}

	if err := rows.Err(); err != nil {
		return err
	}

	// Bound the dedup map: reorgs near the head are the live ones; anything far
	// below it can no longer deepen meaningfully and is dropped so the map cannot
	// grow without bound.
	if len(s.publishedReorgDepth) > maxPublishedReorgDepths {
		if s.blockNumber == 0 {
			// No head to measure distance from -- queryHead has failed on every poll
			// so far -- and the distance predicate would be vacuously false, so the
			// map grew past its own guard. Reset instead: the worst case is
			// re-publishing a reorg that is still inside the skew window.
			s.publishedReorgDepth = make(map[uint64]uint32)
		} else {
			for start := range s.publishedReorgDepth {
				if addSaturating(start, 10*blockBufferSize) < s.blockNumber {
					delete(s.publishedReorgDepth, start)
				}
			}
		}
	}

	s.reorgs = reorgs
	// A full page means there is more behind it. The cursor has moved to the newest
	// row read, so the next poll continues from there rather than re-reading the
	// same page; a backlog drains a page per poll instead of wedging.
	if rowsRead >= maxReorgsPerPoll {
		s.logger.Warn().
			Int("rows_read", rowsRead).
			Time("read_through", readThrough).
			Msg("Reorg page was full; more remain and will be read next poll")
	}
	if readThrough.After(s.latestReorgTime) {
		s.latestReorgTime = readThrough
	}

	return nil
}
