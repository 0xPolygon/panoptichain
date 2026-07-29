package provider

import (
	"container/list"
	"context"
	"errors"
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

	// maxReorgsPerPoll caps the reorg backlog a single cycle will pull. The
	// watermark query is otherwise unbounded, so a long outage would return every
	// reorg since it began in one go.
	maxReorgsPerPoll = 1000
)

// ClickHouseSensorNetworkProvider reads the same sensor data as
// SensorNetworkProvider but from ClickHouse instead of GCP Datastore. It embeds
// SensorNetworkProvider to reuse the shared, backend-agnostic behavior
// (PublishEvents, stolen-block/bogon detection, block-buffer bounds) and only
// overrides the data-fetching path.
type ClickHouseSensorNetworkProvider struct {
	*SensorNetworkProvider
	conn driver.Conn
}

func NewClickHouseSensorNetworkProvider(ctx context.Context, n network.Network, eb *observer.EventBus, cfg config.SensorNetwork) *ClickHouseSensorNetworkProvider {
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

	opts, err := clickhouse.ParseDSN(cfg.ClickHouseDSN)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to parse ClickHouse DSN")
		return &ClickHouseSensorNetworkProvider{SensorNetworkProvider: base}
	}

	conn, err := clickhouse.Open(opts)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to connect to ClickHouse")
		return &ClickHouseSensorNetworkProvider{SensorNetworkProvider: base}
	}

	return &ClickHouseSensorNetworkProvider{SensorNetworkProvider: base, conn: conn}
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

func (s *ClickHouseSensorNetworkProvider) refreshBlockBuffer(ctx context.Context) error {
	s.blockEvents = nil

	if s.blockNumber > s.prevBlockNumber {
		s.prevBlockNumber = s.blockNumber
	}

	// The head is taken from recent events rather than from max(number) over
	// the whole blocks table. blocks is retained forever and ordered by number, so
	// a single bogon block announcing an absurd height would otherwise pin
	// blockNumber permanently and skew clampStart's backfill window for good. A
	// time-bounded window cannot prevent a bogon from moving the head, but it does
	// bound the damage to the window instead of forever.
	var bn uint64
	if err := s.conn.QueryRow(ctx, `
		SELECT max(block_number)
		FROM block_events
		WHERE seen_at > now() - INTERVAL ? SECOND`, int(headWindow.Seconds())).Scan(&bn); err != nil {
		return err
	}
	if bn == 0 {
		// No events in the window: the fleet is down or not yet writing. Leave
		// the previous head in place rather than resetting it to zero.
		return errors.New("no block events in the recent window")
	}
	s.blockNumber = bn

	s.logger.Trace().
		Uint64("block_number", s.blockNumber).
		Msg("Refreshing sensor network block state")

	// clampStart is reused from the embedded SensorNetworkProvider.
	if s.prevBlockNumber != 0 && s.prevBlockNumber != s.blockNumber {
		s.fillRange(ctx, s.clampStart(s.prevBlockNumber))
	}

	return nil
}

func (s *ClickHouseSensorNetworkProvider) fillRange(ctx context.Context, start uint64) {
	s.logger.Debug().
		Uint64("start", start).
		Uint64("end", s.blockNumber).
		Msg("Filling sensor network block range")

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
		LIMIT 1 BY number, hash`, start, s.blockNumber)
	if err != nil {
		s.logger.Warn().Err(err).Msg("Failed to query blocks")
		return
	}
	defer rows.Close()

	// blockTimes feeds the events pass below, which needs each block's header
	// timestamp to compute latency.
	blockTimes := make(map[string]time.Time)

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
		block := types.NewBlockWithHeader(header)
		blockTimes[hash] = blockTime

		if s.blocks.Len() >= blockBufferSize {
			s.blocks.Remove(s.blocks.Front())
		}
		s.blocks.PushBack(block)
	}

	if err := rows.Err(); err != nil {
		s.logger.Warn().Err(err).Msg("Failed to iterate blocks")
	}

	s.getBlockEvents(ctx, start, blockTimes)
}

// getBlockEvents loads every event for the block range in a single query.
//
// This used to issue one `WHERE block_hash = ?` query per block, fanned out over
// unbounded goroutines from inside the still-open blocks cursor -- up to 512
// concurrent point lookups against the largest table in the database after a
// stall. block_events now carries block_number as the leading sort key, so the
// whole range is one ordered scan instead.
func (s *ClickHouseSensorNetworkProvider) getBlockEvents(ctx context.Context, start uint64, blockTimes map[string]time.Time) {
	if len(blockTimes) == 0 {
		return
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
		ORDER BY block_number`, start, s.blockNumber)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to get block events")
		return
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
		s.logger.Warn().Err(err).Msg("Failed to iterate block events")
	}

	s.blockEventsLock.Lock()
	defer s.blockEventsLock.Unlock()

	for blockHash, events := range byBlock {
		blockTime, ok := blockTimes[blockHash]
		if !ok {
			// A event for a block whose header we have not stored (or that fell
			// outside the blocks query). Latency is measured against the header
			// timestamp, so there is nothing to compute without it.
			continue
		}
		// Only block.Time is consumed by BlockEventsObserver, so a minimal header
		// carrying the block timestamp is sufficient here.
		s.blockEvents = append(s.blockEvents, &observer.SensorBlockEvents{
			Block:  &database.DatastoreBlock{DatastoreHeader: &database.DatastoreHeader{Time: blockTime}},
			Events: events,
		})
	}
}

func (s *ClickHouseSensorNetworkProvider) refreshReorgs(ctx context.Context) error {
	// reorg_detections is append-only: the same reorg is re-reported as it deepens,
	// and the deepest detection is the one worth alerting on. Collapsing to one row
	// per start height here is what the old ReplacingMergeTree(depth) was trying to
	// express -- but it belongs in the read, since an insert cannot suppress itself.
	//
	// Note this table is written by the reorg-alerts job, which is still backed by
	// Datastore, so it stays empty (and the reorg and stolen-block metrics stay at
	// zero) until that job is ported.
	rows, err := s.conn.Query(ctx, `
		SELECT
			start_block,
			max(depth)                      AS depth,
			argMax(start_block_hash, depth) AS start_block_hash,
			argMax(end_block, depth)        AS end_block,
			argMax(end_block_hash, depth)   AS end_block_hash,
			max(detected_at)                AS detected_at
		FROM reorg_detections
		WHERE detected_at > ?
		GROUP BY start_block
		ORDER BY detected_at
		LIMIT ?`, s.latestReorgTime, maxReorgsPerPoll)
	if err != nil {
		return err
	}
	defer rows.Close()

	var reorgs []*observer.DatastoreReorg
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

	s.reorgs = reorgs
	if len(reorgs) > 0 {
		s.latestReorgTime = *reorgs[len(reorgs)-1].Time
	}

	return nil
}
