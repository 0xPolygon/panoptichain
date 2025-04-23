// Package observer defines the event and message handing objects that
// are ultimately going to be used for metrics tracking. The observers should be fast and not connect to external data.
package observer

import (
	"context"
	"encoding/json"
	"math/big"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/0xPolygon/panoptichain/api"
	"github.com/0xPolygon/panoptichain/log"
	"github.com/0xPolygon/panoptichain/metrics"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

// HeimdallResult wraps responses payloads in Heimdall v1.
type HeimdallResult[T any] struct {
	Height string `json:"height"`
	Result T      `json:"result"`
}

type PreCommit struct {
	Type    uint64 `json:"type"`
	Height  string `json:"height"`
	Round   string `json:"round"`
	BlockId struct {
		Hash  string `json:"hash"`
		Parts struct {
			Total uint64 `json:"total"`
			Hash  string `json:"hash"`
		} `json:"parts"`
	} `json:"block_id"`
	Timestamp        string `json:"timestamp"`
	ValidatorAddress string `json:"validator_address"`
	ValidatorIndex   string `json:"validator_index"`
	Signature        string `json:"signature"`
	SideTxResults    []struct {
		TxHash string `json:"tx_hash"`
		Result uint64 `json:"result"`
		Sig    string `json:"sig"`
	} `json:"side_tx_results"`
}

type HeimdallBlock struct {
	Result struct {
		Block struct {
			Header struct {
				Time            string `json:"time"`
				Height          string `json:"height"`
				NumTxs          string `json:"num_txs"`
				ProposerAddress string `json:"proposer_address"`
			} `json:"header"`
			Data struct {
				Txs []string `json:"txs"`
			} `json:"data"`
			LastCommit struct {
				PreCommits []*PreCommit `json:"precommits"`
			} `json:"last_commit"`
		} `json:"block"`
	} `json:"result"`
}

type HeimdallValidator struct {
	Address          string `json:"address"`
	VotingPower      string `json:"voting_power"`
	ProposerPriority string `json:"proposer_priority"`
}

type HeimdallValidators struct {
	Result struct {
		BlockHeight string               `json:"block_height"`
		Validators  []*HeimdallValidator `json:"validators"`
		Count       string               `json:"count"`
		Total       string               `json:"total"`
	} `json:"result"`
}

func (b *HeimdallValidators) Validators() []*HeimdallValidator {
	return b.Result.Validators
}

// Number returns the Heimdall block number or nil if it doesn't exist.
func (b *HeimdallBlock) Number() *big.Int {
	height := b.Result.Block.Header.Height
	n, ok := new(big.Int).SetString(height, 10)
	if !ok {
		return nil
	}

	return n
}

func (b *HeimdallBlock) Time() (uint64, error) {
	parsedTime, err := time.Parse(time.RFC3339Nano, b.Result.Block.Header.Time)
	if err != nil {
		return 0, err
	}

	return uint64(parsedTime.Unix()), nil
}

func (b *HeimdallBlock) Txs() *big.Int {
	txs, ok := new(big.Int).SetString(b.Result.Block.Header.NumTxs, 10)
	if !ok {
		return big.NewInt(int64(len(b.Result.Block.Data.Txs)))
	}

	return txs
}

func (b *HeimdallBlock) PreCommits() []*PreCommit {
	return b.Result.Block.LastCommit.PreCommits
}

func (b *HeimdallBlock) ProposerAddress() string {
	return b.Result.Block.Header.ProposerAddress
}

type HeimdallBlockIntervalObserver struct {
	blockInterval *prometheus.HistogramVec
}

func (o *HeimdallBlockIntervalObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.HeimdallBlockInterval, o)

	o.blockInterval = metrics.NewHistogram(
		metrics.Heimdall,
		"block_interval",
		"The time interval (in seconds) between Heimdall blocks",
		newExponentialBuckets(2, 6),
	)
}

func (o *HeimdallBlockIntervalObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)

	interval := m.Data().(uint64)
	logger.Trace().Uint64("interval", interval).Msg("Heimdall block interval")

	o.blockInterval.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(float64(interval))
}

func (o *HeimdallBlockIntervalObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.blockInterval}
}

type HeimdallBlockObserver struct {
	height   *prometheus.GaugeVec
	txs      *prometheus.HistogramVec
	totalTxs *prometheus.CounterVec
}

func (o *HeimdallBlockObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewHeimdallBlock, o)

	o.height = metrics.NewGauge(
		metrics.Heimdall,
		"height",
		"The block height for Heimdall",
	)
	o.txs = metrics.NewHistogram(
		metrics.Heimdall,
		"transactions_per_block",
		"The number of transactions per Heimdall block",
		newExponentialBuckets(2, 11),
	)
	o.totalTxs = metrics.NewCounter(
		metrics.Heimdall,
		"total_transaction_count",
		"The number of total transactions observed for Heimdall",
	)
}

func (o *HeimdallBlockObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)

	block := m.Data().(*HeimdallBlock)

	height := block.Number()
	if height == nil {
		logger.Error().Msg("Failed to get Heimdall block number")
	} else {
		h, _ := height.Float64()
		o.height.WithLabelValues(m.Network().GetName(), m.Provider()).Set(h)
	}

	txs, _ := block.Txs().Float64()
	o.txs.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(txs)
	o.totalTxs.WithLabelValues(m.Network().GetName(), m.Provider()).Add(txs)
}

func (o *HeimdallBlockObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.height, o.txs, o.totalTxs}
}

type HeimdallSignatureCountObserver struct {
	signature *prometheus.GaugeVec
}

func (o *HeimdallSignatureCountObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewHeimdallBlock, o)

	o.signature = metrics.NewGauge(
		metrics.Heimdall,
		"signatures",
		"The number of signatures on block",
	)
}

func (o *HeimdallSignatureCountObserver) Notify(ctx context.Context, m Message) {
	block := m.Data().(*HeimdallBlock)
	o.signature.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(len(block.PreCommits())))
}

func (o *HeimdallSignatureCountObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.signature}
}

type HeimdallMilestoneCount struct {
	Count json.Number `json:"count"`
}

type HeimdallMilestoneCountV1 HeimdallResult[HeimdallMilestoneCount]

type HeimdallMilestone struct {
	Proposer    string      `json:"proposer"`
	StartBlock  json.Number `json:"start_block"`
	EndBlock    json.Number `json:"end_block"`
	Hash        string      `json:"hash"`
	BorChainID  json.Number `json:"bor_chain_id"`
	MilestoneID string      `json:"milestone_id"`
	Timestamp   json.Number `json:"timestamp"`
	Count       int64
	PrevCount   int64
}

type HeimdallMilestoneV1 HeimdallResult[HeimdallMilestone]

type HeimdallMilestoneV2 struct {
	Milestone HeimdallMilestone `json:"milestone"`
}

type MilestoneObserver struct {
	time       *prometheus.GaugeVec
	count      *prometheus.GaugeVec
	startBlock *prometheus.GaugeVec
	endBlock   *prometheus.GaugeVec
	observed   *prometheus.CounterVec
	blockRange *prometheus.HistogramVec
}

func (o *MilestoneObserver) Notify(ctx context.Context, m Message) {
	milestone := m.Data().(*HeimdallMilestone)

	timestamp, err := milestone.Timestamp.Int64()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get milestone timestamp")
	}
	seconds := time.Now().Sub(time.Unix(timestamp, 0)).Seconds()

	startBlock, err := milestone.StartBlock.Float64()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get milestone start block")
	}

	endBlock, err := milestone.EndBlock.Float64()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get milestone end block")
	}

	o.count.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(milestone.Count))
	o.time.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(seconds))
	o.startBlock.WithLabelValues(m.Network().GetName(), m.Provider()).Set(startBlock)
	o.endBlock.WithLabelValues(m.Network().GetName(), m.Provider()).Set(endBlock)

	if milestone.Count > milestone.PrevCount {
		o.observed.WithLabelValues(m.Network().GetName(), m.Provider()).Inc()
		o.blockRange.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(endBlock - startBlock)
	}
}

func (o *MilestoneObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.Milestone, o)

	o.time = metrics.NewGauge(metrics.Heimdall, "time_since_last_milestone", "The time since last milestone")
	o.count = metrics.NewGauge(metrics.Heimdall, "milestone_count", "The milestone count")
	o.startBlock = metrics.NewGauge(metrics.Heimdall, "milestone_start_block", "The milestone start block")
	o.endBlock = metrics.NewGauge(metrics.Heimdall, "milestone_end_block", "The milestone end block")
	o.observed = metrics.NewCounter(metrics.Heimdall, "milestone_observed", "The number of milestones observed")
	o.blockRange = metrics.NewHistogram(
		metrics.Heimdall,
		"milestone_block_range",
		"The number of blocks in the milestone",
		newExponentialBuckets(2, 10),
	)
}

func (o *MilestoneObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{
		o.time,
		o.count,
		o.startBlock,
		o.endBlock,
		o.observed,
		o.blockRange,
	}
}

// HeimdallMissedBlockProposal maps the block number to the list of proposers
// that missed proposing the block.
type HeimdallMissedBlockProposal map[uint64][]string

type HeimdallMissedBlockProposalObserver struct {
	missedBlockProposal *prometheus.CounterVec
}

func (o *HeimdallMissedBlockProposalObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)

	missedBlockProposal := m.Data().(HeimdallMissedBlockProposal)
	for blockNumber, proposers := range missedBlockProposal {
		if len(proposers) > 0 {
			logger.Debug().
				Uint64("block_number", blockNumber).
				Strs("proposers", proposers).
				Msg("Updating Heimdall missed block proposal")
		}

		for _, proposer := range proposers {
			o.missedBlockProposal.WithLabelValues(m.Network().GetName(), m.Provider(), proposer).Inc()
		}
	}
}

func (o *HeimdallMissedBlockProposalObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.HeimdallMissedBlockProposal, o)

	o.missedBlockProposal = metrics.NewCounter(
		metrics.Heimdall,
		"missed_block_proposal",
		"Missed block proposals",
		"signer_address",
	)
}

func (o *HeimdallMissedBlockProposalObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.missedBlockProposal}
}

type HeimdallCheckpoint struct {
	ID         json.Number `json:"id"`
	StartBlock json.Number `json:"start_block"`
	EndBlock   json.Number `json:"end_block"`
	RootHash   string      `json:"root_hash"`
	BorChainID json.Number `json:"bor_chain_id"`
	Timestamp  json.Number `json:"timestamp"`
	Proposer   string      `json:"proposer"`
}

type HeimdallCheckpointV1 HeimdallResult[HeimdallCheckpoint]

type HeimdallCheckpointV2 struct {
	Checkpoint HeimdallCheckpoint `json:"checkpoint"`
}

type HeimdallCheckpointObserver struct {
	startBlock *prometheus.GaugeVec
	endBlock   *prometheus.GaugeVec
	id         *prometheus.GaugeVec
	time       *prometheus.GaugeVec
}

func (o *HeimdallCheckpointObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)

	checkpoint := m.Data().(*HeimdallCheckpoint)

	timestamp, err := checkpoint.Timestamp.Int64()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get checkpoint timestamp")
	}
	seconds := m.Time().Sub(time.Unix(timestamp, 0)).Seconds()

	startBlock, err := checkpoint.StartBlock.Float64()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get checkpoint start block")
	}

	endBlock, err := checkpoint.EndBlock.Float64()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get checkpoint end block")
	}

	id, err := checkpoint.ID.Float64()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get checkpoint id")
	}

	o.startBlock.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(startBlock))
	o.endBlock.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(endBlock))
	o.id.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(id))
	o.time.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(seconds))
}

func (o *HeimdallCheckpointObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.Checkpoint, o)

	o.startBlock = metrics.NewGauge(metrics.Heimdall, "checkpoint_start_block", "The checkpoint start block")
	o.endBlock = metrics.NewGauge(metrics.Heimdall, "checkpoint_end_block", "The checkpoint end block")
	o.id = metrics.NewGauge(metrics.Heimdall, "checkpoint_id", "The checkpoint id")
	o.time = metrics.NewGauge(metrics.Heimdall, "time_since_last_checkpoint", "The time since last checkpoint")
}

func (o *HeimdallCheckpointObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.startBlock, o.endBlock, o.id, o.time}
}

type HeimdallCurrentCheckpointProposerV1 HeimdallResult[api.ValidatorV1]

type HeimdallCurrentCheckpointProposerV2 struct {
	Validator api.ValidatorV2 `json:"validator"`
}

type HeimdallMissedCheckpointProposalObserver struct {
	missedCheckpointProposal *prometheus.CounterVec
}

func (o *HeimdallMissedCheckpointProposalObserver) Notify(ctx context.Context, m Message) {
	proposers := m.Data().([]string)
	for _, proposer := range proposers {
		o.missedCheckpointProposal.WithLabelValues(m.Network().GetName(), m.Provider(), proposer).Inc()
	}
}

func (o *HeimdallMissedCheckpointProposalObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.MissedCheckpointProposal, o)
	o.missedCheckpointProposal = metrics.NewCounter(
		metrics.Heimdall,
		"missed_checkpoint_proposal",
		"Missed checkpoint proposals",
		"signer_address",
	)
}

func (o *HeimdallMissedCheckpointProposalObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.missedCheckpointProposal}
}

type ValidatorsV1 HeimdallResult[[]api.ValidatorV1]

type HeimdallMissedMilestoneProposal struct {
	missedMilestoneProposal *prometheus.CounterVec
}

func (o *HeimdallMissedMilestoneProposal) Notify(ctx context.Context, m Message) {
	proposers := m.Data().([]string)
	for _, proposer := range proposers {
		o.missedMilestoneProposal.WithLabelValues(m.Network().GetName(), m.Provider(), proposer).Inc()
	}
}

func (o *HeimdallMissedMilestoneProposal) Register(eb *EventBus) {
	eb.Subscribe(topics.MissedMilestoneProposal, o)
	o.missedMilestoneProposal = metrics.NewCounter(
		metrics.Heimdall,
		"missed_milestone_proposal",
		"Missed milestone proposals",
		"signer_address",
	)
}

func (o *HeimdallMissedMilestoneProposal) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.missedMilestoneProposal}
}

type HeimdallSpan interface {
	GetID() uint64
	GetStartBlock() uint64
	GetEndBlock() uint64
}

type HeimdallSpanV1 HeimdallResult[struct {
	SpanID     uint64 `json:"span_id"`
	StartBlock uint64 `json:"start_block"`
	EndBlock   uint64 `json:"end_block"`
}]

func (h HeimdallSpanV1) GetID() uint64         { return h.Result.SpanID }
func (h HeimdallSpanV1) GetStartBlock() uint64 { return h.Result.StartBlock }
func (h HeimdallSpanV1) GetEndBlock() uint64   { return h.Result.EndBlock }

type HeimdallSpanV2 struct {
	Span struct {
		ID         uint64 `json:"id,string"`
		StartBlock uint64 `json:"start_block,string"`
		EndBlock   uint64 `json:"end_block,string"`
	} `json:"span"`
}

func (h HeimdallSpanV2) GetID() uint64         { return h.Span.ID }
func (h HeimdallSpanV2) GetStartBlock() uint64 { return h.Span.StartBlock }
func (h HeimdallSpanV2) GetEndBlock() uint64   { return h.Span.EndBlock }

type HeimdallSpanObserver struct {
	spanID     *prometheus.GaugeVec
	startBlock *prometheus.GaugeVec
	endBlock   *prometheus.GaugeVec
}

func (o *HeimdallSpanObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.Span, o)

	o.spanID = metrics.NewGauge(metrics.Heimdall, "span_id", "The span id")
	o.startBlock = metrics.NewGauge(metrics.Heimdall, "span_start_block", "The span start block")
	o.endBlock = metrics.NewGauge(metrics.Heimdall, "span_end_block", "The span end block")
}

func (o *HeimdallSpanObserver) Notify(ctx context.Context, m Message) {
	span := m.Data().(HeimdallSpan)

	o.spanID.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(span.GetID()))
	o.startBlock.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(span.GetStartBlock()))
	o.endBlock.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(span.GetEndBlock()))
}

func (o *HeimdallSpanObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.spanID, o.startBlock, o.endBlock}
}
