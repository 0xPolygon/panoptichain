package provider

import (
	"context"
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	orderedmap "github.com/wk8/go-ordered-map/v2"

	"github.com/0xPolygon/panoptichain/api"
	"github.com/0xPolygon/panoptichain/blockbuffer"
	"github.com/0xPolygon/panoptichain/network"
	"github.com/0xPolygon/panoptichain/observer"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

type HeimdallProvider struct {
	TendermintURL string
	HeimdallURL   string
	Network       network.Network
	Label         string
	bus           *observer.EventBus
	interval      uint
	logger        zerolog.Logger

	BlockNumber         uint64
	prevBlockNumber     uint64
	blockBuffer         *blockbuffer.BlockBuffer
	missedBlockProposal *observer.HeimdallMissedBlockProposal

	checkpoint               *observer.HeimdallCheckpoint
	checkpointProposers      *orderedmap.OrderedMap[string, struct{}]
	missedCheckpointProposal []string

	milestone               *observer.HeimdallMilestone
	prevMilestoneCount      uint64
	milestoneProposers      []*api.Validator
	prevMilestoneProposers  []*api.Validator
	missedMilestoneProposal []string

	span *observer.HeimdallSpan

	refreshStateTime *time.Duration
}

func NewHeimdallProvider(n network.Network, tendermintURL, heimdallURL, label string, eb *observer.EventBus, interval uint) *HeimdallProvider {
	return &HeimdallProvider{
		TendermintURL:       tendermintURL,
		HeimdallURL:         heimdallURL,
		Label:               label,
		blockBuffer:         blockbuffer.NewBlockBuffer(128),
		Network:             n,
		bus:                 eb,
		interval:            interval,
		logger:              NewLogger(n, label),
		checkpointProposers: orderedmap.New[string, struct{}](),
		refreshStateTime:    new(time.Duration),
	}
}

func (h *HeimdallProvider) SetEventBus(bus *observer.EventBus) {
	h.bus = bus
}

func (h *HeimdallProvider) RefreshState(ctx context.Context) error {
	defer timer(h.refreshStateTime)()

	h.logger.Debug().Msg("Refreshing Heimdall State")

	h.refreshBlockBuffer()
	h.refreshMilestone()
	h.refreshCheckpoint()
	h.refreshMissedCheckpointProposal()
	h.refreshMissedBlockProposal()
	h.refreshMissedMilestoneProposal()
	h.refreshSpan()

	return nil
}

func (h *HeimdallProvider) PublishEvents(ctx context.Context) error {
	for i := h.prevBlockNumber + 1; i <= h.BlockNumber && h.prevBlockNumber != 0; i++ {
		b, err := h.blockBuffer.GetBlock(i)
		if err != nil {
			continue
		}

		block, ok := b.(*observer.HeimdallBlock)
		if !ok {
			continue
		}

		m := observer.NewMessage(h.Network, h.Label, block)
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

		interval := observer.NewMessage(h.Network, h.Label, time-prevTime)
		h.bus.Publish(ctx, topics.HeimdallBlockInterval, interval)
	}

	if h.milestone != nil {
		m := observer.NewMessage(h.Network, h.Label, h.milestone)
		h.bus.Publish(ctx, topics.Milestone, m)
	}

	if h.checkpoint != nil {
		m := observer.NewMessage(h.Network, h.Label, h.checkpoint)
		h.bus.Publish(ctx, topics.Checkpoint, m)
	}

	if len(h.missedCheckpointProposal) > 0 {
		m := observer.NewMessage(h.Network, h.Label, h.missedCheckpointProposal)
		h.bus.Publish(ctx, topics.MissedCheckpointProposal, m)
	}

	if h.missedBlockProposal != nil {
		m := observer.NewMessage(h.Network, h.Label, h.missedBlockProposal)
		h.bus.Publish(ctx, topics.HeimdallMissedBlockProposal, m)
	}

	if len(h.missedMilestoneProposal) > 0 {
		m := observer.NewMessage(h.Network, h.Label, h.missedMilestoneProposal)
		h.bus.Publish(ctx, topics.MissedMilestoneProposal, m)
	}

	if h.span != nil {
		m := observer.NewMessage(h.Network, h.Label, h.span)
		h.bus.Publish(ctx, topics.Span, m)
	}

	h.bus.Publish(ctx, topics.RefreshStateTime, observer.NewMessage(h.Network, h.Label, h.refreshStateTime))

	return nil
}

func (h *HeimdallProvider) PollingInterval() uint {
	return h.interval
}

func (h *HeimdallProvider) refreshBlockBuffer() {
	h.prevBlockNumber = h.BlockNumber
	block := h.getBlock(0)
	if block == nil {
		return
	}

	bn := block.Number()
	if bn == nil {
		return
	}
	h.BlockNumber = bn.Uint64()

	h.logger.Debug().Uint64("block_number", h.BlockNumber).Msg("Refreshed Heimdall state")
	if h.prevBlockNumber != 0 && h.prevBlockNumber != h.BlockNumber {
		h.fillRange(h.prevBlockNumber)
	}
}

func (h *HeimdallProvider) getBlock(height uint64) *observer.HeimdallBlock {
	path, err := url.JoinPath(h.TendermintURL, "block")
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to join path when fetching Heimdall block")
		return nil
	}

	if height > 0 {
		path = fmt.Sprintf("%s?height=%d", path, height)
	}

	block := &observer.HeimdallBlock{}
	err = api.GetJSON(path, &block)
	if err != nil {
		h.logger.Warn().Err(err).Msg("Failed to get Heimdall block")
		return nil
	}

	return block
}

func (h *HeimdallProvider) getValidators(height uint64) *observer.HeimdallValidators {
	path, err := url.JoinPath(h.TendermintURL, "validators")
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to join path when fetching Heimdall validators")
		return nil
	}

	if height > 0 {
		path = fmt.Sprintf("%s?height=%d", path, height)
	}

	validators := &observer.HeimdallValidators{}
	err = api.GetJSON(path, &validators)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall validators")
		return nil
	}

	return validators
}

func (h *HeimdallProvider) fillRange(start uint64) {
	h.logger.Debug().
		Uint64("start_block", start).
		Uint64("end_block", h.BlockNumber).
		Str("url", h.TendermintURL).
		Msg("Filling block range")

	for i := start + 1; i <= h.BlockNumber; i++ {
		block := h.getBlock(i)
		if block == nil {
			h.logger.Warn().Uint64("block_number", i).Msg("Unable to get block")
			break
		}

		h.blockBuffer.PutBlock(block)
	}
}

func (h *HeimdallProvider) refreshMilestone() error {
	if h.milestone != nil {
		h.prevMilestoneCount = h.milestone.Count
	}

	path, err := url.JoinPath(h.HeimdallURL, "milestone/count")
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to join path when fetching Heimdall milestone count")
		return err
	}

	count := &observer.HeimdallMilestoneCount{}

	err = api.GetJSON(path, &count)
	if err != nil {
		h.logger.Error().Err(err).Msg("Unable to get latest Heimdall milestone count")
		return err
	}

	path, err = url.JoinPath(h.HeimdallURL, "milestone", fmt.Sprint(count.Result.Count))
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to join path when fetching Heimdall milestone")
		return err
	}

	err = api.GetJSON(path, &h.milestone)
	if err != nil {
		h.logger.Error().Err(err).Msg("Unable to get latest heimdall milestone")
		return err
	}

	h.milestone.PrevCount = max(h.milestone.PrevCount, h.milestone.Count)
	h.milestone.Count = count.Result.Count
	h.logger.Info().Any("milestone", h.milestone).Msg("Received Heimdall milestone")

	return nil
}

func (h *HeimdallProvider) refreshCheckpoint() error {
	path, err := url.JoinPath(h.HeimdallURL, "checkpoints/latest")
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to join path when fetching Heimdall checkpoint")
		return err
	}

	err = api.GetJSON(path, &h.checkpoint)
	if err != nil {
		h.logger.Error().Err(err).Msg("Unable to get latest Heimdall checkpoint")
		return err
	}

	h.logger.Info().Any("checkpoint", h.checkpoint).Msg("Received Heimdall checkpoint")

	return nil
}

func (h *HeimdallProvider) refreshMissedCheckpointProposal() error {
	h.logger.Info().
		Any("checkpoint_proposers", h.checkpointProposers).
		Any("missed_checkpoint_proposers", h.missedCheckpointProposal).
		Msg("Refreshing missed checkpoint proposal")

	h.missedCheckpointProposal = nil
	var currentProposer *observer.HeimdallCurrentProposer

	path, err := url.JoinPath(h.HeimdallURL, "staking/current-proposer")
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to join path when fetching Heimdall current checkpoint proposer")
		return err
	}

	err = api.GetJSON(path, &currentProposer)
	if err != nil {
		h.logger.Error().Err(err).Msg("Unable to get Heimdall current checkpoint proposer")
		return err
	}

	signer := currentProposer.Result.Signer
	if _, ok := h.checkpointProposers.Get(signer); !ok {
		h.checkpointProposers.Set(signer, struct{}{})
	}

	latestProposer := h.checkpoint.Result.Proposer
	if _, ok := h.checkpointProposers.Get(latestProposer); !ok {
		return nil
	}

	for pair := h.checkpointProposers.Oldest(); pair != nil; pair = pair.Next() {
		proposer := pair.Key

		h.checkpointProposers.Delete(proposer)
		if proposer == latestProposer {
			break
		}

		h.missedCheckpointProposal = append(h.missedCheckpointProposal, proposer)
	}

	return nil
}

func (h *HeimdallProvider) refreshMissedBlockProposal() error {
	missedBlockProposal := make(observer.HeimdallMissedBlockProposal)
	for i := h.prevBlockNumber + 1; i <= h.BlockNumber && h.prevBlockNumber != 0; i++ {
		currentBlock := h.getBlock(i)
		if currentBlock == nil {
			h.logger.Debug().Msg("Failed to get current block. Skipping.")
			continue
		}
		currentBlockProposer := currentBlock.ProposerAddress()

		validatorsData := h.getValidators(i - 1)
		if validatorsData == nil {
			h.logger.Debug().Msg("Failed to get current validator data. Skipping.")
			continue
		}

		validators := validatorsData.Validators()
		// in descending order
		sort.Slice(validators, func(i, j int) bool {
			pi, _ := strconv.Atoi(validators[i].ProposerPriority)
			pj, _ := strconv.Atoi(validators[j].ProposerPriority)
			return pi > pj
		})

		failedProposers := observer.FailedProposerInfo{}

		for _, validator := range validators {
			if validator.Address == currentBlockProposer {
				break
			}
			failedProposers.FailedProposers = append(failedProposers.FailedProposers, validator.Address)
		}

		missedBlockProposal[i] = failedProposers
	}

	h.missedBlockProposal = &missedBlockProposal

	return nil
}

func (h *HeimdallProvider) refreshMissedMilestoneProposal() error {
	h.missedMilestoneProposal = nil
	var proposers *observer.HeimdallMilestoneProposers

	path, err := url.JoinPath(h.HeimdallURL, "staking/milestoneProposer", fmt.Sprint(500))
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to join path when fetching Heimdall milestone proposers")
		return err
	}

	err = api.GetJSON(path, &proposers)
	if err != nil {
		h.logger.Error().Err(err).Msg("Unable to get Heimdall milestone proposers")
		return err
	}

	// While h.prevMilestoneCount and h.milestone keep track of the actual
	// previous count, h.prevMilestoneProposers will only update the milestone
	// proposers state if it has changed.
	//
	// The milestone proposers will change before the milestone, so at this point,
	// there is confidence that h.prevMilestoneProposers is the proposer set for
	// h.milestone (the latest milestone) and h.milestoneProposers is the
	// proposer set for the next milestone.
	if !reflect.DeepEqual(proposers.Result, h.milestoneProposers) {
		h.prevMilestoneProposers = h.milestoneProposers
		h.milestoneProposers = proposers.Result
	}

	// This checks if the there is a new milestone.
	if h.milestone == nil || h.prevMilestoneCount == h.milestone.Count {
		return nil
	}

	for _, validator := range h.prevMilestoneProposers {
		// Stop when we see the latest milestone.
		if validator.Signer == h.milestone.Result.Proposer {
			break
		}

		h.missedMilestoneProposal = append(h.missedMilestoneProposal, validator.Signer)
	}

	if len(h.missedMilestoneProposal) > 0 {
		h.logger.Info().
			Any("missed_milestone_proposal", h.missedMilestoneProposal).
			Msg("Validators missed milestone proposal")
	}

	return nil
}

func (h *HeimdallProvider) refreshSpan() error {
	url, err := url.JoinPath(h.HeimdallURL, "bor/latest-span")
	if err != nil {
		h.logger.Error().Err(err).Msg("Unable to get Heimdall URL")
		return err
	}

	err = api.GetJSON(url, &h.span)
	if err != nil {
		h.logger.Error().Err(err).Msg("Unable to get latest Heimdall span")
		return err
	}

	h.logger.Info().Any("span", h.span).Msg("Received Heimdall span")

	return nil
}
