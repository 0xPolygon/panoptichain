package provider

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
	orderedmap "github.com/wk8/go-ordered-map/v2"

	"github.com/0xPolygon/panoptichain/api"
	"github.com/0xPolygon/panoptichain/blockbuffer"
	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/network"
	"github.com/0xPolygon/panoptichain/observer"
	"github.com/0xPolygon/panoptichain/observer/topics"
	"github.com/0xPolygon/panoptichain/proto/heimdall"
)

// ErrInvalidSpan is returned when a span has zero ID with zero blocks,
// indicating a malformed API response.
var ErrInvalidSpan = errors.New("invalid span: zero ID with zero blocks")

type HeimdallProvider struct {
	tendermintURL string
	heimdallURL   string
	network       network.Network
	label         string
	bus           *observer.EventBus
	interval      time.Duration
	logger        zerolog.Logger
	maxSpanLag    uint64

	blockNumber         uint64
	prevBlockNumber     uint64
	blockBuffer         *blockbuffer.BlockBuffer
	missedBlockProposal observer.HeimdallMissedBlockProposal

	checkpoint                *observer.HeimdallCheckpoint
	checkpointProposers       *orderedmap.OrderedMap[string, struct{}]
	missedCheckpointProposers []string

	milestones         []*observer.HeimdallMilestone
	prevMilestoneCount int64

	spans *observer.HeimdallSpans

	validatorSets *observer.HeimdallValidatorSets

	missedVotes    []*observer.HeimdallMissedVotes
	milestoneVotes []*observer.HeimdallMilestoneVotes
	validatorIDMap map[string]uint64 // normalized signer_address -> val_id

	refreshStateTime *time.Duration
}

func NewHeimdallProvider(n network.Network, eb *observer.EventBus, cfg config.HeimdallEndpoint) *HeimdallProvider {
	maxSpanLag := config.DefaultMaxSpanLag
	if cfg.MaxSpanLag != nil {
		maxSpanLag = *cfg.MaxSpanLag
	}

	return &HeimdallProvider{
		tendermintURL:       cfg.TendermintURL,
		heimdallURL:         cfg.HeimdallURL,
		network:             n,
		label:               cfg.Label,
		bus:                 eb,
		blockBuffer:         blockbuffer.NewBlockBuffer(128),
		interval:            GetInterval(cfg.Interval),
		logger:              NewLogger(n, cfg.Label),
		maxSpanLag:          maxSpanLag,
		checkpointProposers: orderedmap.New[string, struct{}](),
		refreshStateTime:    new(time.Duration),
		spans:               &observer.HeimdallSpans{},
		validatorSets:       &observer.HeimdallValidatorSets{},
	}
}

func (h *HeimdallProvider) RefreshState(ctx context.Context) error {
	defer timer(h.refreshStateTime)()

	h.logger.Debug().Msg("Refreshing Heimdall state")

	h.refreshBlockBuffer()
	h.refreshMilestone()
	h.refreshCheckpoint()
	h.refreshMissedCheckpointProposal()
	h.refreshMissedBlockProposal()
	h.refreshSpan()
	h.refreshValidatorSet()
	h.refreshMissedVotes()
	h.refreshMilestoneVotes()

	return nil
}

func (h *HeimdallProvider) PublishEvents(ctx context.Context) error {
	for i := h.prevBlockNumber + 1; i <= h.blockNumber && h.prevBlockNumber != 0; i++ {
		b, err := h.blockBuffer.GetBlock(i)
		if err != nil {
			continue
		}

		block, ok := b.(*observer.HeimdallBlock)
		if !ok {
			continue
		}

		m := observer.NewMessage(h.network, h.label, block)
		h.bus.Publish(ctx, topics.NewHeimdallBlock, m)

		bn := b.Number()
		if bn == nil {
			continue
		}

		pb, err := h.blockBuffer.GetBlock(bn.Uint64() - 1)
		if pb == nil {
			continue
		}

		prev, ok := pb.(*observer.HeimdallBlock)
		if !ok {
			continue
		}

		time, err := block.Time()
		if err != nil {
			h.logger.Warn().Err(err).Msg("Failed to get Heimdall block time")
			continue
		}

		prevTime, err := prev.Time()
		if err != nil {
			h.logger.Warn().Err(err).Msg("Failed to get previous Heimdall block time")
			continue
		}

		interval := observer.NewMessage(h.network, h.label, time-prevTime)
		h.bus.Publish(ctx, topics.HeimdallBlockInterval, interval)
	}

	if h.missedBlockProposal != nil {
		m := observer.NewMessage(h.network, h.label, h.missedBlockProposal)
		h.bus.Publish(ctx, topics.HeimdallMissedBlockProposal, m)
	}

	if h.checkpoint != nil {
		m := observer.NewMessage(h.network, h.label, h.checkpoint)
		h.bus.Publish(ctx, topics.Checkpoint, m)
	}

	if len(h.missedCheckpointProposers) > 0 {
		m := observer.NewMessage(h.network, h.label, h.missedCheckpointProposers)
		h.bus.Publish(ctx, topics.MissedCheckpointProposal, m)
	}

	for _, milestone := range h.milestones {
		m := observer.NewMessage(h.network, h.label, milestone)
		h.bus.Publish(ctx, topics.Milestone, m)
	}

	if h.spans != nil {
		m := observer.NewMessage(h.network, h.label, h.spans)
		h.bus.Publish(ctx, topics.Span, m)
	}

	if h.validatorSets != nil {
		m := observer.NewMessage(h.network, h.label, h.validatorSets)
		h.bus.Publish(ctx, topics.ValidatorSet, m)
	}

	for _, mv := range h.missedVotes {
		if mv.MissingCount > 0 {
			m := observer.NewMessage(h.network, h.label, mv)
			h.bus.Publish(ctx, topics.MissedVote, m)
		}
	}

	for _, mv := range h.milestoneVotes {
		m := observer.NewMessage(h.network, h.label, mv)
		h.bus.Publish(ctx, topics.MilestoneVote, m)
	}

	h.bus.Publish(ctx, topics.RefreshStateTime, observer.NewMessage(h.network, h.label, h.refreshStateTime))

	return nil
}

func (h *HeimdallProvider) PollingInterval() time.Duration {
	return h.interval
}

func (h *HeimdallProvider) refreshBlockBuffer() {
	h.prevBlockNumber = h.blockNumber
	block := h.getBlock(0)
	if block == nil {
		return
	}

	bn := block.Number()
	if bn == nil {
		return
	}
	h.blockNumber = bn.Uint64()

	h.logger.Debug().Uint64("block_number", h.blockNumber).Msg("Refreshing Heimdall state")
	if h.prevBlockNumber != 0 && h.prevBlockNumber != h.blockNumber {
		h.fillRange(h.prevBlockNumber)
	}
}

func (h *HeimdallProvider) getBlock(height uint64) *observer.HeimdallBlock {
	path, err := url.JoinPath(h.tendermintURL, "block")
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to join path when fetching Heimdall block")
		return nil
	}

	if height > 0 {
		path = fmt.Sprintf("%s?height=%d", path, height)
	}

	var block observer.HeimdallBlock
	err = api.GetJSON(path, &block)
	if err != nil {
		h.logger.Warn().Err(err).Msg("Failed to get Heimdall block")
		return nil
	}

	return &block
}

func (h *HeimdallProvider) getValidators(height uint64) *observer.HeimdallValidators {
	path, err := url.JoinPath(h.tendermintURL, "validators")
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to join path when fetching Heimdall validators")
		return nil
	}

	if height > 0 {
		path = fmt.Sprintf("%s?height=%d", path, height)
	}

	var validators observer.HeimdallValidators
	err = api.GetJSON(path, &validators)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall validators")
		return nil
	}

	return &validators
}

// getValidatorsAtHeight fetches all validators at a specific height with pagination.
func (h *HeimdallProvider) getValidatorsAtHeight(height uint64) ([]*observer.HeimdallValidator, error) {
	const perPage = 100
	const maxPages = 10

	var v []*observer.HeimdallValidator

	for page := 1; page <= maxPages; page++ {
		path, err := url.JoinPath(h.tendermintURL, "validators")
		if err != nil {
			return nil, fmt.Errorf("failed to join validators path: %w", err)
		}

		path = fmt.Sprintf("%s?height=%d&per_page=%d&page=%d", path, height, perPage, page)

		var validators observer.HeimdallValidators
		if err := api.GetJSON(path, &validators); err != nil {
			return nil, fmt.Errorf("failed to get validators at height %d page %d: %w", height, page, err)
		}

		v = append(v, validators.Validators()...)

		total, _ := strconv.Atoi(validators.Result.Total)
		if len(v) >= total {
			break
		}
	}

	return v, nil
}

func (h *HeimdallProvider) fillRange(start uint64) {
	h.logger.Debug().
		Uint64("start_block", start).
		Uint64("end_block", h.blockNumber).
		Msg("Filling block range")

	for i := start + 1; i <= h.blockNumber; i++ {
		block := h.getBlock(i)
		if block == nil {
			h.logger.Warn().Uint64("block_number", i).Msg("Failed to get block")
			break
		}

		h.blockBuffer.PutBlock(block)
	}
}

func (h *HeimdallProvider) refreshMilestone() error {
	path, err := url.JoinPath(h.heimdallURL, "milestones", "count")
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall milestone count path")
		return err
	}

	var count observer.HeimdallMilestoneCount
	if err := api.GetJSON(path, &count); err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall milestone count")
		return err
	}

	currentCount := int64(count.Count)
	h.milestones = nil

	// On first poll, fetch only the latest milestone to establish baseline.
	// On subsequent polls, fetch all new milestones in range.
	start := h.prevMilestoneCount + 1
	if h.prevMilestoneCount == 0 {
		start = currentCount
	}

	for i := start; i <= currentCount; i++ {
		path, err := url.JoinPath(h.heimdallURL, "milestones", strconv.FormatInt(i, 10))
		if err != nil {
			h.logger.Error().Err(err).Msg("Failed to get Heimdall milestone path")
			continue
		}

		var v2 observer.HeimdallMilestoneV2
		if err = api.GetJSON(path, &v2); err != nil {
			h.logger.Error().Err(err).Int64("milestone", i).Msg("Failed to get Heimdall milestone")
			continue
		}

		milestone := &v2.Milestone
		milestone.Count = i
		h.milestones = append(h.milestones, milestone)
	}

	h.prevMilestoneCount = currentCount
	return nil
}

func (h *HeimdallProvider) refreshCheckpoint() error {
	path, err := url.JoinPath(h.heimdallURL, "checkpoints", "latest")
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall latest checkpoint path")
		return err
	}

	var v2 observer.HeimdallCheckpointV2
	if err = api.GetJSON(path, &v2); err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall latest checkpoint")
		return err
	}

	h.checkpoint = &v2.Checkpoint

	return nil
}

func (h *HeimdallProvider) getCurrentCheckpointProposer() (string, error) {
	path, err := url.JoinPath(h.heimdallURL, "checkpoints", "prepare-next")
	if err != nil {
		return "", err
	}

	var resp observer.HeimdallPrepareNextCheckpoint
	if err = api.GetJSON(path, &resp); err != nil {
		return "", err
	}

	return resp.Checkpoint.Proposer, nil
}

func (h *HeimdallProvider) refreshMissedCheckpointProposal() error {
	var proposers []string
	for pair := h.checkpointProposers.Oldest(); pair != nil; pair = pair.Next() {
		proposers = append(proposers, pair.Key)
	}

	h.logger.Debug().
		Any("checkpoint_proposers", proposers).
		Any("missed_checkpoint_proposers", h.missedCheckpointProposers).
		Msg("Refreshing missed checkpoint proposal")

	h.missedCheckpointProposers = nil

	signer, err := h.getCurrentCheckpointProposer()
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall current checkpoint proposer")
		return err
	}

	if _, ok := h.checkpointProposers.Get(signer); !ok {
		h.checkpointProposers.Set(signer, struct{}{})
	}

	latest := h.checkpoint.Proposer
	if _, ok := h.checkpointProposers.Get(latest); !ok {
		return nil
	}

	for pair := h.checkpointProposers.Oldest(); pair != nil; pair = pair.Next() {
		proposer := pair.Key

		h.checkpointProposers.Delete(proposer)
		if proposer == latest {
			break
		}

		h.missedCheckpointProposers = append(h.missedCheckpointProposers, proposer)
	}

	return nil
}

func (h *HeimdallProvider) refreshMissedBlockProposal() error {
	missedBlockProposal := make(observer.HeimdallMissedBlockProposal)
	for i := h.prevBlockNumber + 1; i <= h.blockNumber && h.prevBlockNumber != 0; i++ {
		block := h.getBlock(i)
		if block == nil {
			h.logger.Debug().Msg("Failed to get current block")
			continue
		}
		proposer := block.ProposerAddress()

		v := h.getValidators(i - 1)
		if v == nil {
			h.logger.Debug().Msg("Failed to get validators")
			continue
		}
		validators := v.Validators()

		// Sort validators in descending order.
		sort.Slice(validators, func(i, j int) bool {
			pi, _ := strconv.Atoi(validators[i].ProposerPriority)
			pj, _ := strconv.Atoi(validators[j].ProposerPriority)
			return pi > pj
		})

		var proposers []string
		for _, validator := range validators {
			if validator.Address == proposer {
				break
			}
			proposers = append(proposers, validator.Address)
		}

		missedBlockProposal[i] = proposers
	}

	h.missedBlockProposal = missedBlockProposal

	return nil
}

func (h *HeimdallProvider) refreshSpan() error {
	// Always fetch the latest span
	latest, err := h.getLatestSpan()
	if err != nil {
		return err
	}

	// Set current span on startup
	if h.spans.Curr == nil {
		h.spans.Curr = latest
		return nil
	}

	// No new span available
	if latest.ID == h.spans.Curr.ID {
		return nil
	}

	// Check if lag exceeds maximum threshold
	lag := latest.ID - h.spans.Curr.ID
	if lag > h.maxSpanLag {
		h.logger.Warn().
			Uint64("current_span_id", h.spans.Curr.ID).
			Uint64("latest_span_id", latest.ID).
			Uint64("lag", lag).
			Uint64("max_span_lag", h.maxSpanLag).
			Msg("Span lag exceeds maximum, jumping to latest")

		h.spans.Prev = h.spans.Curr
		h.spans.Curr = latest
		return nil
	}

	// Fetch next span sequentially to ensure overlap detection works
	next := h.spans.Curr.ID + 1
	span, err := h.getSpan(next)
	if err != nil {
		h.logger.Warn().Uint64("span_id", next).Err(err).Msg("Failed to fetch span")
		return err
	}

	h.spans.Prev = h.spans.Curr
	h.spans.Curr = span
	return nil
}

func (h *HeimdallProvider) getLatestSpan() (*observer.HeimdallSpan, error) {
	return h.fetchSpan("latest")
}

func (h *HeimdallProvider) getSpan(id uint64) (*observer.HeimdallSpan, error) {
	return h.fetchSpan(strconv.FormatUint(id, 10))
}

func (h *HeimdallProvider) fetchSpan(spanID string) (*observer.HeimdallSpan, error) {
	path, err := url.JoinPath(h.heimdallURL, "bor", "spans", spanID)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall span path")
		return nil, err
	}

	var v2 observer.HeimdallSpanV2
	if err = api.GetJSON(path, &v2); err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall span")
		return nil, err
	}

	span := &v2.Span
	if span.ID == 0 && span.StartBlock == 0 && span.EndBlock == 0 {
		h.logger.Error().
			Str("requested_span", spanID).
			Msg("Received invalid zero-value span from API")
		return nil, ErrInvalidSpan
	}

	return span, nil
}

func (h *HeimdallProvider) refreshValidatorSet() error {
	validators, err := api.GetValidators(h.heimdallURL)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get validator set")
		return err
	}

	curr := make(observer.ValidatorMap, len(validators))
	ids := make(map[string]uint64, len(validators))
	for _, v := range validators {
		curr[v.ID] = v
		ids[normalizeAddress(v.Signer)] = v.ID
	}

	if h.validatorSets.Curr != nil {
		h.validatorSets.Prev = h.validatorSets.Curr
	}
	h.validatorSets.Curr = curr
	h.validatorIDMap = ids

	return nil
}

func (h *HeimdallProvider) getCommit(height uint64) (*observer.HeimdallCommit, error) {
	path, err := url.JoinPath(h.tendermintURL, "commit")
	if err != nil {
		return nil, fmt.Errorf("failed to join commit path: %w", err)
	}

	var commit observer.HeimdallCommit
	if err := api.GetJSON(fmt.Sprintf("%s?height=%d", path, height), &commit); err != nil {
		return nil, fmt.Errorf("failed to get commit at height %d: %w", height, err)
	}

	return &commit, nil
}

func normalizeAddress(addr string) string {
	return strings.ToLower(strings.TrimPrefix(addr, "0x"))
}

func (h *HeimdallProvider) getMissedVotes(height uint64) (*observer.HeimdallMissedVotes, error) {
	validators, err := h.getValidatorsAtHeight(height)
	if err != nil {
		return nil, fmt.Errorf("failed to get validators at height %d: %w", height, err)
	}

	commit, err := h.getCommit(height)
	if err != nil {
		return nil, err
	}

	signatures := commit.Result.SignedHeader.Commit.Signatures
	if len(signatures) != len(validators) {
		h.logger.Warn().
			Int("validators", len(validators)).
			Int("signatures", len(signatures)).
			Uint64("height", height).
			Msg("Validator and signature array length mismatch")
		return nil, nil
	}

	var missedVotes []observer.HeimdallMissedVote
	for i, sig := range signatures {
		if sig.BlockIDFlag == 2 {
			continue
		}

		validator := validators[i]
		valID := h.validatorIDMap[normalizeAddress(validator.Address)]

		flagLabel := "absent"
		if sig.BlockIDFlag == 3 {
			flagLabel = "nil"
		}

		missedVotes = append(missedVotes, observer.HeimdallMissedVote{
			ValidatorID:   valID,
			SignerAddress: validator.Address,
			FlagLabel:     flagLabel,
		})
	}

	return &observer.HeimdallMissedVotes{
		Height:       height,
		MissingCount: len(missedVotes),
		MissedVotes:  missedVotes,
	}, nil
}

func (h *HeimdallProvider) refreshMissedVotes() {
	if h.validatorIDMap == nil {
		return
	}

	h.missedVotes = nil

	h.logger.Debug().Msg("Refreshing missed consensus votes")

	for height := h.prevBlockNumber + 1; height <= h.blockNumber && h.prevBlockNumber != 0; height++ {
		mv, err := h.getMissedVotes(height)
		if err != nil {
			h.logger.Warn().Err(err).Uint64("height", height).Msg("Failed to detect missed votes")
			continue
		}
		if mv != nil {
			h.missedVotes = append(h.missedVotes, mv)
		}
	}
}

// getExtendedCommitInfo fetches and decodes the ExtendedCommitInfo from txs[0]
// of a Heimdall block. Vote extensions from block H-1 are stored in block H's txs[0].
func (h *HeimdallProvider) getExtendedCommitInfo(height uint64) (*heimdall.ExtendedCommitInfo, error) {
	block := h.getBlock(height)
	if block == nil {
		return nil, fmt.Errorf("failed to get block at height %d", height)
	}

	txs := block.Result.Block.Data.Txs
	if len(txs) == 0 {
		return nil, fmt.Errorf("no transactions in block at height %d", height)
	}

	// txs[0] contains the base64-encoded ExtendedCommitInfo
	veBytes, err := base64.StdEncoding.DecodeString(txs[0])
	if err != nil {
		return nil, fmt.Errorf("failed to base64 decode txs[0]: %w", err)
	}

	extCommit, err := heimdall.UnmarshalExtendedCommitInfo(veBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal ExtendedCommitInfo: %w", err)
	}

	return extCommit, nil
}

// getMilestoneVotes processes vote extensions from a block and returns milestone vote data.
func (h *HeimdallProvider) getMilestoneVotes(height uint64) (*observer.HeimdallMilestoneVotes, error) {
	extCommit, err := h.getExtendedCommitInfo(height)
	if err != nil {
		return nil, err
	}

	var (
		votes           []observer.HeimdallMilestoneVote
		totalVP         int64
		milestoneVP     int64
		milestoneVoters int
	)

	for _, v := range extCommit.Votes {
		if v.Validator == nil {
			continue
		}
		validatorAddr := hex.EncodeToString(v.Validator.Address)
		votingPower := v.Validator.Power
		totalVP += votingPower

		vote := observer.HeimdallMilestoneVote{
			ValidatorAddress: validatorAddr,
			ValidatorID:      h.validatorIDMap[normalizeAddress(validatorAddr)],
			VotingPower:      votingPower,
			BlockIDFlag:      int(v.BlockIdFlag),
		}

		if len(v.VoteExtension) > 0 {
			h.getMilestoneFromVoteExtension(&vote, v.VoteExtension)
			if vote.HasMilestone {
				milestoneVP += votingPower
				milestoneVoters++
			}
		}

		votes = append(votes, vote)
	}

	return &observer.HeimdallMilestoneVotes{
		Height:               height,
		TotalValidators:      len(votes),
		TotalVotingPower:     totalVP,
		MilestoneVoters:      milestoneVoters,
		MilestoneVotingPower: milestoneVP,
		Votes:                votes,
	}, nil
}

// getMilestoneFromVoteExtension decodes the vote extension and populates milestone data if present.
func (h *HeimdallProvider) getMilestoneFromVoteExtension(vote *observer.HeimdallMilestoneVote, data []byte) {
	ve, err := heimdall.UnmarshalVoteExtension(data)
	if err != nil {
		h.logger.Warn().
			Err(err).
			Str("validator", vote.ValidatorAddress).
			Msg("Failed to decode vote extension")
		return
	}

	if ve.MilestoneProposition == nil {
		return
	}

	mp := ve.MilestoneProposition
	if len(mp.BlockHashes) == 0 {
		return
	}
	vote.HasMilestone = true
	vote.MilestoneStart = mp.StartBlockNumber
	vote.MilestoneEnd = mp.StartBlockNumber + uint64(len(mp.BlockHashes)) - 1
}

func (h *HeimdallProvider) refreshMilestoneVotes() {
	if h.validatorIDMap == nil {
		return
	}

	h.milestoneVotes = nil

	h.logger.Debug().Msg("Refreshing milestone votes")

	for height := h.prevBlockNumber + 1; height <= h.blockNumber && h.prevBlockNumber != 0; height++ {
		mv, err := h.getMilestoneVotes(height)
		if err != nil {
			h.logger.Warn().Err(err).Uint64("height", height).Msg("Failed to detect milestone votes")
			continue
		}
		h.milestoneVotes = append(h.milestoneVotes, mv)
	}
}
