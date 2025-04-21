// Package observer defines the event and message handing objects that
// are ultimately going to be used for metrics tracking. The observers should be fast and not connect to external data.
package observer

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/0xPolygon/panoptichain/api"
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
				TotalTxs        string `json:"total_txs"`
				ProposerAddress string `json:"proposer_address"`
			} `json:"header"`
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

func (b *HeimdallBlock) NumTxs() (*big.Int, error) {
	txs, ok := new(big.Int).SetString(b.Result.Block.Header.NumTxs, 10)
	if !ok {
		return nil, errors.New("failed to parse number of transactions")
	}

	return txs, nil
}

func (b *HeimdallBlock) TotalTxs() (*big.Int, error) {
	totalTxs, ok := new(big.Int).SetString(b.Result.Block.Header.TotalTxs, 10)
	if !ok {
		return nil, errors.New("failed to parse total number of transactions")
	}

	return totalTxs, nil
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

type HeimdallTransactionCountObserver struct {
	transactionCount *prometheus.HistogramVec
}

func (o *HeimdallTransactionCountObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewHeimdallBlock, o)

	o.transactionCount = metrics.NewHistogram(
		metrics.Heimdall,
		"transactions_per_block",
		"The number of transactions per Heimdall block",
		newExponentialBuckets(2, 11),
	)
}

func (o *HeimdallTransactionCountObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)

	block := m.Data().(*HeimdallBlock)

	txs, err := block.NumTxs()
	if err != nil {
		logger.Error().Msg("Failed to get number of transactions")
		return
	}

	o.transactionCount.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(float64(txs.Uint64()))
}

func (o *HeimdallTransactionCountObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.transactionCount}
}

type HeimdallTotalTransactionCountObserver struct {
	totalTransactionCount *prometheus.CounterVec
}

func (o *HeimdallTotalTransactionCountObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewHeimdallBlock, o)

	o.totalTransactionCount = metrics.NewCounter(
		metrics.Heimdall,
		"total_transaction_count",
		"The number of total transactions for Heimdall",
	)
}

func (o *HeimdallTotalTransactionCountObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)

	block := m.Data().(*HeimdallBlock)

	txs, err := block.TotalTxs()
	if err != nil {
		logger.Error().Msg("Failed to get total number of transactions")
		return
	}

	o.totalTransactionCount.WithLabelValues(m.Network().GetName(), m.Provider()).Add(float64(txs.Uint64()))
}

func (o *HeimdallTotalTransactionCountObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.totalTransactionCount}
}

type HeimdallHeightObserver struct {
	height *prometheus.GaugeVec
}

func (o *HeimdallHeightObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewHeimdallBlock, o)

	o.height = metrics.NewGauge(
		metrics.Heimdall,
		"height",
		"The block height for Heimdall",
	)
}

func (o *HeimdallHeightObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)

	block := m.Data().(*HeimdallBlock)

	height := block.Number()
	if height == nil {
		logger.Error().Msg("Failed to get Heimdall block number")
		return
	}

	o.height.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(height.Uint64()))
}

func (o *HeimdallHeightObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.height}
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
	Count uint64 `json:"count"`
}

type HeimdallMilestoneCountV1 HeimdallResult[HeimdallMilestoneCount]

type HeimdallMilestone struct {
	Proposer    string `json:"proposer"`
	StartBlock  uint64 `json:"start_block"`
	EndBlock    uint64 `json:"end_block"`
	Hash        string `json:"hash"`
	BorChainID  string `json:"bor_chain_id"`
	MilestoneID string `json:"milestone_id"`
	Timestamp   uint64 `json:"timestamp"`
	Count       uint64
	PrevCount   uint64
}

type HeimdallMilestoneV1 HeimdallResult[HeimdallMilestone]

type MilestoneObserver struct {
	time       *prometheus.GaugeVec
	height     *prometheus.GaugeVec
	count      *prometheus.GaugeVec
	startBlock *prometheus.GaugeVec
	endBlock   *prometheus.GaugeVec
	observed   *prometheus.CounterVec
	blockRange *prometheus.HistogramVec
}

func (o *MilestoneObserver) Notify(ctx context.Context, m Message) {
	milestone := m.Data().(*HeimdallMilestone)
	seconds := time.Now().Sub(time.Unix(int64(milestone.Timestamp), 0)).Seconds()

	start := float64(milestone.StartBlock)
	end := float64(milestone.EndBlock)

	o.count.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(milestone.Count))
	o.time.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(seconds))
	o.startBlock.WithLabelValues(m.Network().GetName(), m.Provider()).Set(start)
	o.endBlock.WithLabelValues(m.Network().GetName(), m.Provider()).Set(end)

	if milestone.Count > milestone.PrevCount {
		o.observed.WithLabelValues(m.Network().GetName(), m.Provider()).Inc()
		o.blockRange.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(end - start)
	}
}

func (o *MilestoneObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.Milestone, o)

	o.time = metrics.NewGauge(metrics.Heimdall, "time_since_last_milestone", "The time since last milestone")
	o.height = metrics.NewGauge(metrics.Heimdall, "milestone_block_height", "The milestone block height")
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
		o.height,
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

type HeimdallCheckpointBuffer struct {
	Proposer   string `json:"proposer"`
	StartBlock uint64 `json:"start_block"`
	EndBlock   uint64 `json:"end_block"`
	RootHash   string `json:"root_hash"`
	BorChainID string `json:"bor_chain_id"`
	Timestamp  uint64 `json:"timestamp"`
}

type HeimdallCheckpoint struct {
	HeimdallCheckpointBuffer
	ID uint64 `json:"id"`
}

type HeimdallCheckpointV1 HeimdallResult[HeimdallCheckpoint]

type HeimdallCheckpointObserver struct {
	startBlock *prometheus.GaugeVec
	endBlock   *prometheus.GaugeVec
	id         *prometheus.GaugeVec
	time       *prometheus.GaugeVec
}

func (o *HeimdallCheckpointObserver) Notify(ctx context.Context, m Message) {
	checkpoint := m.Data().(*HeimdallCheckpoint)
	checkpointTime := time.Unix(int64(checkpoint.Timestamp), 0)
	seconds := m.Time().Sub(checkpointTime).Seconds()

	o.startBlock.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(checkpoint.StartBlock))
	o.endBlock.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(checkpoint.EndBlock))
	o.id.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(checkpoint.ID))
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

type ValidatorV1 HeimdallResult[api.Validator]

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

type ValidatorsV1 HeimdallResult[[]api.Validator]

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

type HeimdallSpan struct {
	SpanID     int64 `json:"span_id"`
	StartBlock int64 `json:"start_block"`
	EndBlock   int64 `json:"end_block"`
}

type HeimdallSpanV1 HeimdallResult[HeimdallSpan]

type HeimdallSpanObserver struct {
	height     *prometheus.GaugeVec
	spanID     *prometheus.GaugeVec
	startBlock *prometheus.GaugeVec
	endBlock   *prometheus.GaugeVec
}

func (o *HeimdallSpanObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.Span, o)

	o.height = metrics.NewGauge(metrics.Heimdall, "span_height", "The span height")
	o.spanID = metrics.NewGauge(metrics.Heimdall, "span_id", "The span id")
	o.startBlock = metrics.NewGauge(metrics.Heimdall, "span_start_block", "The span start block")
	o.endBlock = metrics.NewGauge(metrics.Heimdall, "span_end_block", "The span end block")
}

func (o *HeimdallSpanObserver) Notify(ctx context.Context, m Message) {
	span := m.Data().(*HeimdallSpan)

	o.spanID.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(span.SpanID))
	o.startBlock.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(span.StartBlock))
	o.endBlock.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(span.EndBlock))
}

func (o *HeimdallSpanObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.height, o.spanID, o.startBlock, o.endBlock}
}
