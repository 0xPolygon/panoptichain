package observer

import (
	"context"
	"encoding/hex"
	"sort"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/maticnetwork/polygon-cli/p2p/database"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/0xPolygon/panoptichain/api"
	"github.com/0xPolygon/panoptichain/metrics"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

const ReorgsKind = "reorgs"

type DatastoreReorg struct {
	Depth      int
	Start      int
	End        int
	StartBlock *datastore.Key
	EndBlock   *datastore.Key
	Time       *time.Time
}

type ReorgObserver struct {
	depth *prometheus.HistogramVec
}

func (o *ReorgObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.Reorg, o)

	o.depth = metrics.NewHistogram(
		metrics.Sensor,
		"reorg_depth",
		"The number of blocks that were reorganized",
		newExponentialBuckets(2, 7),
	)
}

func (o *ReorgObserver) Notify(ctx context.Context, m Message) {
	reorg := m.Data().(*DatastoreReorg)
	o.depth.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(float64(reorg.Depth))
}

func (o *ReorgObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.depth}
}

type SensorBlocks struct {
	Start  uint64
	End    uint64
	Blocks types.Blocks
}

type SensorBlocksObserver struct {
	forksPerBlockNumber *prometheus.HistogramVec
	totalBlocks         *prometheus.CounterVec
}

func (o *SensorBlocksObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.SensorBlocks, o)

	o.forksPerBlockNumber = metrics.NewHistogram(
		metrics.Sensor,
		"forks_per_block_number",
		"The cardinality of forks for a block number",
		newLinearBuckets(0, 10, 1),
	)
	o.totalBlocks = metrics.NewCounter(
		metrics.Sensor,
		"total_blocks",
		"The total number of blocks observed by sensors including bogons",
	)
}

func (o *SensorBlocksObserver) Notify(ctx context.Context, msg Message) {
	data := msg.Data().(*SensorBlocks)

	n := float64(len(data.Blocks))
	o.totalBlocks.WithLabelValues(msg.Network().GetName(), msg.Provider()).Add(n)

	m := make(map[uint64]types.Blocks)

	for _, block := range data.Blocks {
		n := block.Number().Uint64()
		m[n] = append(m[n], block)
	}

	// Iterate over the entire collected range to ensure that there's no missing
	// blocks.
	for i := data.Start; i < data.End; i++ {
		var n float64 = 0
		if blocks, ok := m[i]; ok {
			n = float64(len(blocks))
		}

		o.forksPerBlockNumber.WithLabelValues(msg.Network().GetName(), msg.Provider()).Observe(n)
	}
}

func (o *SensorBlocksObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.forksPerBlockNumber, o.totalBlocks}
}

type DoubleSignObserver struct {
	doubleSign *prometheus.CounterVec
}

func (o *DoubleSignObserver) Notify(ctx context.Context, msg Message) {
	logger := NewLogger(o, msg)

	data := msg.Data().(*SensorBlocks)
	m := make(map[uint64]types.Blocks)

	for _, block := range data.Blocks {
		n := block.Number().Uint64()
		m[n] = append(m[n], block)
	}

	for blockNumber, blocks := range m {
		signerCount := make(map[string]int)

		for _, block := range blocks {
			bytes, err := api.Ecrecover(block.Header())
			if err != nil {
				logger.Warn().Err(err).Msg("Failed to get block signer")
				continue
			}

			signer := "0x" + hex.EncodeToString(bytes)
			signerCount[signer]++

			if signerCount[signer] > 1 {
				o.doubleSign.WithLabelValues(msg.Network().GetName(), msg.Provider(), signer).Inc()

				logger.Debug().
					Uint64("block_number", blockNumber).
					Str("signer", signer).
					Msg("Double sign detected")
			}
		}
	}
}

func (o *DoubleSignObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.SensorBlocks, o)

	o.doubleSign = metrics.NewCounter(
		metrics.Sensor,
		"double_sign",
		"Number of double sign events detected",
		"signer_address",
	)
}

func (o *DoubleSignObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.doubleSign}
}

type SensorBogonBlockObserver struct {
	bogonBlocks *prometheus.CounterVec
}

func (o *SensorBogonBlockObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)

	signers, err := api.Signers(m.Network())
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to get signers validator map")
		return
	}

	data := m.Data().(*SensorBlocks)

	for _, block := range data.Blocks {
		bytes, err := api.Ecrecover(block.Header())

		if err != nil {
			logger.Warn().Err(err).Msg("Failed to get block signer")
			return
		}
		signer := "0x" + hex.EncodeToString(bytes)

		if _, ok := signers[signer]; !ok {
			o.bogonBlocks.WithLabelValues(m.Network().GetName(), m.Provider(), signer).Inc()
		}
	}
}

func (o *SensorBogonBlockObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.SensorBlocks, o)

	o.bogonBlocks = metrics.NewCounter(
		metrics.Sensor,
		"bogon_block",
		"The total number of bogon blocks observed",
		"signer_address",
	)
}

func (o *SensorBogonBlockObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.bogonBlocks}
}

type SealedOutOfTurnObserver struct {
	sealedOutOfTurn *prometheus.CounterVec
}

func (o *SealedOutOfTurnObserver) Notify(ctx context.Context, msg Message) {
	logger := NewLogger(o, msg)

	signers, err := api.Signers(msg.Network())
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to get signers validator map")
		return
	}

	data := msg.Data().(*SensorBlocks)
	m := make(map[uint64][]*types.Block)

	for _, block := range data.Blocks {
		n := block.Number().Uint64()
		m[n] = append(m[n], block)
	}

	for _, blocks := range m {
		// Sort the blocks by descending difficulty.
		sort.Slice(blocks, func(i, j int) bool {
			return blocks[i].Difficulty().Uint64() > blocks[j].Difficulty().Uint64()
		})

		// This map is used to filter out double signers.
		seenSigners := make(map[string]struct{})

		for i, block := range blocks {
			// Skip the first iteration because this would be the highest difficulty
			// and not be out of turn.
			if i == 0 {
				continue
			}

			bytes, err := api.Ecrecover(block.Header())
			if err != nil {
				logger.Warn().Err(err).Msg("Failed to get block signer")
				return
			}

			signer := "0x" + hex.EncodeToString(bytes)

			// Skip blocks that have already been signed by the same signers for this
			// block number.
			if _, ok := seenSigners[signer]; ok {
				continue
			}
			seenSigners[signer] = struct{}{}

			if _, ok := signers[signer]; ok {
				o.sealedOutOfTurn.WithLabelValues(msg.Network().GetName(), msg.Provider(), signer).Inc()
			}
		}
	}
}

func (o *SealedOutOfTurnObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.SensorBlocks, o)

	o.sealedOutOfTurn = metrics.NewCounter(
		metrics.Sensor,
		"sealed_out_of_turn",
		"The number of blocks that were sealed out of turn",
		"signer_address",
	)
}

func (o *SealedOutOfTurnObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.sealedOutOfTurn}
}

type StolenBlockObserver struct {
	stolenBlock *prometheus.CounterVec
}

func (o *StolenBlockObserver) Notify(ctx context.Context, msg Message) {
	logger := NewLogger(o, msg)

	block := msg.Data().(*types.Block)

	bytes, err := api.Ecrecover(block.Header())
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to get block signer")
		return
	}
	signer := "0x" + hex.EncodeToString(bytes)

	o.stolenBlock.WithLabelValues(msg.Network().GetName(), msg.Provider(), signer).Inc()
}

func (o *StolenBlockObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.StolenBlock, o)

	o.stolenBlock = metrics.NewCounter(
		metrics.Sensor,
		"stolen_block",
		"The number blocks stolen from a validator",
		"signer_address",
	)
}

func (o *StolenBlockObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.stolenBlock}
}

type SensorBlockEvents struct {
	Block  *database.DatastoreBlock
	Events []database.DatastoreEvent
}

type BlockEventsObserver struct {
	latency     *prometheus.HistogramVec
	diff        *prometheus.HistogramVec
	first       *prometheus.HistogramVec
	last        *prometheus.HistogramVec
	peers       *prometheus.HistogramVec
	events      *prometheus.HistogramVec
	connections *prometheus.HistogramVec
	rank        *prometheus.HistogramVec
}

type blockLatencies struct {
	first time.Duration
	last  time.Duration
}

func (o *BlockEventsObserver) Notify(ctx context.Context, m Message) {
	data := m.Data().(*SensorBlockEvents)

	block := data.Block
	events := data.Events

	sort.Slice(events, func(i, j int) bool {
		return events[i].Time.Compare(events[j].Time) < 0
	})

	latencies := make(map[string]*blockLatencies)

	peers := make(map[string]struct{})
	peersBySensor := make(map[string]map[string]struct{})
	peersConnected := make(map[string]map[string]struct{})

	eventsBySensor := make(map[string]int)

	rank := 0
	ranks := make(map[string]int)

	for _, event := range events {
		latency := event.Time.Sub(block.Time)

		peers[event.PeerId] = struct{}{}

		if _, ok := peersBySensor[event.SensorId]; !ok {
			peersBySensor[event.SensorId] = make(map[string]struct{})
		}
		peersBySensor[event.SensorId][event.PeerId] = struct{}{}

		if _, ok := peersConnected[event.PeerId]; !ok {
			peersConnected[event.PeerId] = make(map[string]struct{})
		}
		peersConnected[event.PeerId][event.SensorId] = struct{}{}

		eventsBySensor[event.SensorId]++

		if _, ok := ranks[event.SensorId]; !ok {
			rank++
			ranks[event.SensorId] = rank
		}

		if _, ok := latencies[event.SensorId]; !ok {
			latencies[event.SensorId] = &blockLatencies{latency, latency}
		}

		latencies[event.SensorId].last = latency
		o.latency.WithLabelValues(m.Network().GetName(), m.Provider(), event.SensorId).Observe(float64(latency.Milliseconds()))
	}

	o.peers.WithLabelValues(m.Network().GetName(), m.Provider(), "all").Observe(float64(len(peers)))
	for sensor, p := range peersBySensor {
		o.peers.WithLabelValues(m.Network().GetName(), m.Provider(), sensor).Observe(float64(len(p)))
	}

	for _, sensors := range peersConnected {
		o.connections.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(float64(len(sensors)))
	}

	for sensor, e := range eventsBySensor {
		o.events.WithLabelValues(m.Network().GetName(), m.Provider(), sensor).Observe(float64(e))
	}

	for sensor, r := range ranks {
		o.rank.WithLabelValues(m.Network().GetName(), m.Provider(), sensor).Observe(float64(r))
	}

	for sensor, latency := range latencies {
		dt := latency.last - latency.first
		o.diff.WithLabelValues(m.Network().GetName(), m.Provider(), sensor).Observe(float64(dt.Milliseconds()))
		o.first.WithLabelValues(m.Network().GetName(), m.Provider(), sensor).Observe(float64(latency.first.Milliseconds()))
		o.last.WithLabelValues(m.Network().GetName(), m.Provider(), sensor).Observe(float64(latency.last.Milliseconds()))
	}
}

func (o *BlockEventsObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.SensorBlockEvents, o)

	o.latency = metrics.NewHistogram(
		metrics.Sensor,
		"block_latency",
		"The difference between the block time and the time seen for all blocks by a sensor (in milliseconds)",
		newExponentialBuckets(2, 16),
		"sensor",
	)
	o.diff = metrics.NewHistogram(
		metrics.Sensor,
		"block_latency_diff",
		"The difference between the first and last time a block was received by a sensor (in milliseconds)",
		newExponentialBuckets(2, 16),
		"sensor",
	)
	o.first = metrics.NewHistogram(
		metrics.Sensor,
		"first_block_latency",
		"The difference between the block time and the time first seen by a sensor (in milliseconds)",
		newExponentialBuckets(2, 16),
		"sensor",
	)
	o.last = metrics.NewHistogram(
		metrics.Sensor,
		"last_block_latency",
		"The difference between the block time and the time last seen by a sensor (in milliseconds)",
		newExponentialBuckets(2, 16),
		"sensor",
	)
	o.peers = metrics.NewHistogram(
		metrics.Sensor,
		"peers",
		"The number of unique block propagators observed by a sensor",
		newExponentialBuckets(2, 10),
		"sensor",
	)
	o.events = metrics.NewHistogram(
		metrics.Sensor,
		"block_events",
		"The number of block events that occurred for a block by a sensor",
		newExponentialBuckets(2, 14),
		"sensor",
	)
	o.connections = metrics.NewHistogram(
		metrics.Sensor,
		"connections",
		"The number of sensors a peer propagated a block to",
		newLinearBuckets(0, 10, 1),
	)
	o.rank = metrics.NewHistogram(
		metrics.Sensor,
		"latency_rank",
		"The rank order in which a sensor received a block relative to other sensors",
		newLinearBuckets(0, 10, 1),
		"sensor",
	)
}

func (o *BlockEventsObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.latency, o.diff, o.first, o.last, o.peers, o.events, o.connections, o.rank}
}
