package provider

import (
	"container/list"
	"context"
	"errors"
	"math/big"
	"sync"
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

	var bn uint64
	if err := s.conn.QueryRow(ctx, "SELECT number FROM blocks ORDER BY number DESC LIMIT 1").Scan(&bn); err != nil {
		return err
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
	rows, err := s.conn.Query(ctx, `
		SELECT hash, number, parent_hash, block_time, coinbase, difficulty,
		       gas_used, gas_limit, base_fee, uncle_hash, state_root, tx_root,
		       receipt_root, logs_bloom, extra_data, mix_digest, nonce
		FROM blocks
		WHERE number >= ? AND number < ?
		ORDER BY number`, start, s.blockNumber)
	if err != nil {
		s.logger.Warn().Err(err).Msg("Failed to query blocks")
		return
	}
	defer rows.Close()

	var wg sync.WaitGroup

	for rows.Next() {
		var (
			hash, parentHash, coinbase                string
			uncleHash, stateRoot, txRoot, receiptRoot string
			mixDigest, logsBloom, extraData           string
			number, difficulty, gasUsed, gasLimit     uint64
			baseFee, nonce                            uint64
			blockTime                                 time.Time
		)

		if err := rows.Scan(&hash, &number, &parentHash, &blockTime, &coinbase, &difficulty,
			&gasUsed, &gasLimit, &baseFee, &uncleHash, &stateRoot, &txRoot,
			&receiptRoot, &logsBloom, &extraData, &mixDigest, &nonce); err != nil {
			s.logger.Warn().Err(err).Msg("Failed to scan block")
			continue
		}

		// Full header so downstream ecrecover (signer/bogon/stolen) works.
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
			BaseFee:     new(big.Int).SetUint64(baseFee),
		}
		block := types.NewBlockWithHeader(header)

		wg.Go(func() {
			s.getBlockEvents(ctx, hash, blockTime)
		})

		if s.blocks.Len() >= blockBufferSize {
			s.blocks.Remove(s.blocks.Front())
		}
		s.blocks.PushBack(block)
	}

	if err := rows.Err(); err != nil {
		s.logger.Warn().Err(err).Msg("Failed to iterate blocks")
	}

	wg.Wait()
}

func (s *ClickHouseSensorNetworkProvider) getBlockEvents(ctx context.Context, blockHash string, blockTime time.Time) {
	rows, err := s.conn.Query(ctx,
		"SELECT sensor_id, peer_id, seen_at FROM block_events WHERE block_hash = ?", blockHash)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to get block events")
		return
	}
	defer rows.Close()

	var events []database.DatastoreEvent
	for rows.Next() {
		var sensorID, peerID string
		var seenAt time.Time
		if err := rows.Scan(&sensorID, &peerID, &seenAt); err != nil {
			s.logger.Warn().Err(err).Msg("Failed to scan block event")
			continue
		}
		events = append(events, database.DatastoreEvent{
			SensorId: sensorID,
			PeerId:   peerID,
			Time:     seenAt,
		})
	}

	s.blockEventsLock.Lock()
	defer s.blockEventsLock.Unlock()

	// Only block.Time is consumed by BlockEventsObserver, so a minimal header
	// carrying the block timestamp is sufficient here.
	s.blockEvents = append(s.blockEvents, &observer.SensorBlockEvents{
		Block:  &database.DatastoreBlock{DatastoreHeader: &database.DatastoreHeader{Time: blockTime}},
		Events: events,
	})
}

func (s *ClickHouseSensorNetworkProvider) refreshReorgs(ctx context.Context) error {
	rows, err := s.conn.Query(ctx, `
		SELECT start_block, depth, start_block_hash, end_block, end_block_hash, detected_at
		FROM reorgs
		WHERE detected_at > ?
		ORDER BY detected_at`, s.latestReorgTime)
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
