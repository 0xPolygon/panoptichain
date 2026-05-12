// Package observer defines the event and message handing objects that
// are ultimately going to be used for metrics tracking. The observers should be fast and not connect to external data.
package observer

import (
	"context"
	"math/big"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"

	"github.com/0xPolygon/panoptichain/api"
	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/metrics"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

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

	for _, h := range config.Config().Providers.HeimdallEndpoints {
		o.totalTxs.WithLabelValues(h.Name, h.Label).Add(0)
	}
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
	Count uint64 `json:"count,string"`
}

type HeimdallMilestone struct {
	Proposer    string `json:"proposer"`
	StartBlock  uint64 `json:"start_block,string"`
	EndBlock    uint64 `json:"end_block,string"`
	Hash        string `json:"hash"`
	BorChainID  uint64 `json:"bor_chain_id,string"`
	MilestoneID string `json:"milestone_id"`
	Timestamp   int64  `json:"timestamp,string"`
	Count       int64
	PrevCount   int64
}

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

	seconds := time.Since(time.Unix(milestone.Timestamp, 0)).Seconds()
	startBlock := milestone.StartBlock
	endBlock := milestone.EndBlock

	o.count.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(milestone.Count))
	o.time.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(seconds))
	o.startBlock.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(milestone.StartBlock))
	o.endBlock.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(milestone.EndBlock))

	if milestone.Count > milestone.PrevCount {
		o.observed.WithLabelValues(m.Network().GetName(), m.Provider()).Inc()
		o.blockRange.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(float64(endBlock - startBlock))
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

	for _, h := range config.Config().Providers.HeimdallEndpoints {
		o.observed.WithLabelValues(h.Name, h.Label).Add(0)
	}
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
	ID         uint64 `json:"id,string"`
	StartBlock uint64 `json:"start_block,string"`
	EndBlock   uint64 `json:"end_block,string"`
	RootHash   string `json:"root_hash"`
	BorChainID uint64 `json:"bor_chain_id,string"`
	Timestamp  uint64 `json:"timestamp,string"`
	Proposer   string `json:"proposer"`
}

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
	checkpoint := m.Data().(*HeimdallCheckpoint)

	var seconds float64
	if checkpoint.Timestamp != 0 {
		seconds = m.Time().Sub(time.Unix(int64(checkpoint.Timestamp), 0)).Seconds()
	}

	o.startBlock.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(checkpoint.StartBlock))
	o.endBlock.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(checkpoint.EndBlock))
	o.id.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(checkpoint.ID))
	o.time.WithLabelValues(m.Network().GetName(), m.Provider()).Set(seconds)
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

type HeimdallCurrentCheckpointProposer struct {
	Validator api.Validator `json:"validator"`
}

type HeimdallPrepareNextCheckpoint struct {
	Checkpoint struct {
		Proposer string `json:"proposer"`
	} `json:"checkpoint"`
}

type HeimdallMissedCheckpointProposalObserver struct {
	counter *prometheus.CounterVec
}

func (o *HeimdallMissedCheckpointProposalObserver) Notify(ctx context.Context, m Message) {
	proposers := m.Data().([]string)
	for _, proposer := range proposers {
		o.counter.WithLabelValues(m.Network().GetName(), m.Provider(), proposer).Inc()
	}
}

func (o *HeimdallMissedCheckpointProposalObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.MissedCheckpointProposal, o)
	o.counter = metrics.NewCounter(
		metrics.Heimdall,
		"missed_checkpoint_proposal",
		"Missed checkpoint proposals",
		"signer_address",
	)
}

func (o *HeimdallMissedCheckpointProposalObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.counter}
}

type HeimdallSpan struct {
	ID                uint64          `json:"id,string"`
	StartBlock        uint64          `json:"start_block,string"`
	EndBlock          uint64          `json:"end_block,string"`
	SelectedProducers []api.Validator `json:"selected_producers"`
}

type HeimdallSpanV2 struct {
	Span HeimdallSpan `json:"span"`
}

type HeimdallSpans struct {
	Curr *HeimdallSpan
	Prev *HeimdallSpan
}

// ValidatorMap maps validator IDs to validators.
type ValidatorMap = map[uint64]api.Validator

// HeimdallValidatorSets contains the current and previous validator sets.
type HeimdallValidatorSets struct {
	Curr ValidatorMap
	Prev ValidatorMap
}

type HeimdallCommitSignature struct {
	BlockIDFlag      int    `json:"block_id_flag"`
	ValidatorAddress string `json:"validator_address"`
	Timestamp        string `json:"timestamp"`
	Signature        string `json:"signature"`
}

type HeimdallCommitData struct {
	Height     string                    `json:"height"`
	Signatures []HeimdallCommitSignature `json:"signatures"`
}

type HeimdallSignedHeader struct {
	Commit HeimdallCommitData `json:"commit"`
}

type HeimdallCommit struct {
	Result struct {
		SignedHeader HeimdallSignedHeader `json:"signed_header"`
	} `json:"result"`
}

type HeimdallMissedVote struct {
	ValidatorID   uint64
	SignerAddress string
	FlagLabel     string
}

type HeimdallMissedVotes struct {
	Height       uint64
	MissingCount int
	MissedVotes  []HeimdallMissedVote
}

// HeimdallMilestoneVote represents a validator's milestone vote from vote extensions.
type HeimdallMilestoneVote struct {
	ValidatorAddress string
	ValidatorID      uint64
	VotingPower      int64
	BlockIDFlag      int // see heimdall.BlockIDFlag* constants
	HasMilestone     bool
	MilestoneStart   uint64
	MilestoneEnd     uint64
}

// HeimdallMilestoneVotes represents all milestone votes for a block.
type HeimdallMilestoneVotes struct {
	Height               uint64
	TotalValidators      int
	TotalVotingPower     int64
	MilestoneVoters      int   // validators who proposed milestone
	MilestoneVotingPower int64 // VP of milestone voters
	Votes                []HeimdallMilestoneVote
}

type HeimdallSpanObserver struct {
	spanID     *prometheus.GaugeVec
	startBlock *prometheus.GaugeVec
	endBlock   *prometheus.GaugeVec
	producer   *prometheus.GaugeVec
	overlaps   *prometheus.CounterVec
	overlapped *prometheus.CounterVec
}

func (o *HeimdallSpanObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.Span, o)

	o.spanID = metrics.NewGauge(metrics.Heimdall, "span_id", "The span id")
	o.startBlock = metrics.NewGauge(metrics.Heimdall, "span_start_block", "The span start block")
	o.endBlock = metrics.NewGauge(metrics.Heimdall, "span_end_block", "The span end block")
	o.producer = metrics.NewGauge(metrics.Heimdall, "span_producer", "The span selected producer")
	o.overlaps = metrics.NewCounter(metrics.Heimdall, "span_overlaps", "The number of overlapping spans")
	o.overlapped = metrics.NewCounter(metrics.Heimdall, "span_overlapped_blocks", "The number of overlapped blocks between spans")

	for _, h := range config.Config().Providers.HeimdallEndpoints {
		o.overlaps.WithLabelValues(h.Name, h.Label).Add(0)
		o.overlapped.WithLabelValues(h.Name, h.Label).Add(0)
	}
}

func (o *HeimdallSpanObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)
	data := m.Data().(*HeimdallSpans)

	network := m.Network().GetName()
	provider := m.Provider()

	curr := data.Curr
	if curr == nil {
		return
	}

	o.spanID.WithLabelValues(network, provider).Set(float64(curr.ID))
	o.startBlock.WithLabelValues(network, provider).Set(float64(curr.StartBlock))
	o.endBlock.WithLabelValues(network, provider).Set(float64(curr.EndBlock))

	if len(curr.SelectedProducers) != 1 {
		logger.Warn().
			Int("selected_producers", len(curr.SelectedProducers)).
			Msg("Unexpected number of selected producers")
	} else {
		producer := float64(curr.SelectedProducers[0].ID)
		o.producer.WithLabelValues(network, provider).Set(producer)
	}

	prev := data.Prev
	if prev == nil || curr.ID <= prev.ID {
		return
	}

	if curr.StartBlock <= prev.EndBlock || curr.StartBlock == prev.StartBlock {
		o.overlaps.WithLabelValues(network, provider).Add(1)
		blocks := prev.EndBlock - curr.StartBlock + 1
		o.overlapped.WithLabelValues(network, provider).Add(float64(blocks))
	}
}

func (o *HeimdallSpanObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{
		o.spanID,
		o.startBlock,
		o.endBlock,
		o.producer,
		o.overlaps,
		o.overlapped,
	}
}

type HeimdallValidatorSetChangeObserver struct {
	counter *prometheus.CounterVec
}

func (o *HeimdallValidatorSetChangeObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.ValidatorSet, o)
	o.counter = metrics.NewCounter(
		metrics.Heimdall,
		"validator_set_change",
		"The number of validator set changes (onboarded or unbonded)",
		"change_type",
	)

	for _, h := range config.Config().Providers.HeimdallEndpoints {
		o.counter.WithLabelValues(h.Name, h.Label, "onboarded").Add(0)
		o.counter.WithLabelValues(h.Name, h.Label, "unbonded").Add(0)
	}
}

func (o *HeimdallValidatorSetChangeObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)
	data := m.Data().(*HeimdallValidatorSets)

	if data.Prev == nil {
		return
	}

	o.detectChanges(logger, m, data.Curr, data.Prev, "onboarded")
	o.detectChanges(logger, m, data.Prev, data.Curr, "unbonded")
}

// detectChanges finds validators in set A that are not in set B and records
// them as the specified change type.
func (o *HeimdallValidatorSetChangeObserver) detectChanges(logger zerolog.Logger, m Message, a, b ValidatorMap, changeType string) {
	for id, v := range a {
		if _, exists := b[id]; exists {
			continue
		}

		logger.Info().
			Uint64("validator_id", v.ID).
			Str("signer", v.Signer).
			Str("change_type", changeType).
			Msg("Validator set change detected")

		o.counter.WithLabelValues(m.Network().GetName(), m.Provider(), changeType).Inc()
	}
}

func (o *HeimdallValidatorSetChangeObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.counter}
}

// HeimdallMissedVoteObserver tracks missed consensus votes.
type HeimdallMissedVoteObserver struct {
	consensus *prometheus.CounterVec
}

func (o *HeimdallMissedVoteObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.MissedVote, o)

	o.consensus = metrics.NewCounter(
		metrics.Heimdall,
		"missed_consensus_vote",
		"Missed Heimdall consensus votes",
		"validator_id", "signer_address", "flag",
	)
}

func (o *HeimdallMissedVoteObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)
	missed := m.Data().(*HeimdallMissedVotes)
	network := m.Network().GetName()
	provider := m.Provider()

	for _, vote := range missed.MissedVotes {
		id := strconv.FormatUint(vote.ValidatorID, 10)
		o.consensus.WithLabelValues(network, provider, id, vote.SignerAddress, vote.FlagLabel).Inc()
	}

	if missed.MissingCount > 0 {
		logger.Trace().
			Uint64("height", missed.Height).
			Int("missing_count", missed.MissingCount).
			Msg("Detected missed votes")
	}
}

func (o *HeimdallMissedVoteObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.consensus}
}

// HeimdallMilestoneVoteObserver tracks milestone votes from vote extensions.
type HeimdallMilestoneVoteObserver struct {
	proposed    *prometheus.CounterVec
	missed      *prometheus.CounterVec
	votingPower *prometheus.GaugeVec
}

func (o *HeimdallMilestoneVoteObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.MilestoneVote, o)

	o.proposed = metrics.NewCounter(
		metrics.Heimdall,
		"milestone_vote_proposed",
		"Validators who proposed milestone in vote extension",
		"validator_id", "signer_address",
	)

	o.missed = metrics.NewCounter(
		metrics.Heimdall,
		"milestone_vote_missed",
		"Validators who signed but didn't propose milestone",
		"validator_id", "signer_address",
	)

	o.votingPower = metrics.NewGauge(
		metrics.Heimdall,
		"milestone_voting_power",
		"Percentage of voting power that proposed milestone",
	)
}

func (o *HeimdallMilestoneVoteObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)
	votes := m.Data().(*HeimdallMilestoneVotes)
	network := m.Network().GetName()
	provider := m.Provider()

	var vpPct float64
	if votes.TotalVotingPower > 0 {
		vpPct = float64(votes.MilestoneVotingPower) / float64(votes.TotalVotingPower) * 100
	}
	o.votingPower.WithLabelValues(network, provider).Set(vpPct)

	for _, vote := range votes.Votes {
		id := strconv.FormatUint(vote.ValidatorID, 10)
		switch {
		case vote.HasMilestone:
			o.proposed.WithLabelValues(network, provider, id, vote.ValidatorAddress).Inc()
		default:
			// Did not propose milestone (regardless of whether they signed)
			o.missed.WithLabelValues(network, provider, id, vote.ValidatorAddress).Inc()
		}
	}

	missed := votes.TotalValidators - votes.MilestoneVoters
	if missed > 0 {
		logger.Trace().
			Uint64("height", votes.Height).
			Int("total_validators", votes.TotalValidators).
			Int("milestone_voters", votes.MilestoneVoters).
			Int("missed", missed).
			Msg("Detected validators without milestone vote")
	}
}

func (o *HeimdallMilestoneVoteObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.proposed, o.missed, o.votingPower}
}
