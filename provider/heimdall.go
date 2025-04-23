package provider

import (
	"context"
	"errors"
	"fmt"
	"math/big"
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
	version       uint

	BlockNumber         uint64
	prevBlockNumber     uint64
	blockBuffer         *blockbuffer.BlockBuffer
	missedBlockProposal observer.HeimdallMissedBlockProposal

	checkpoint                *observer.HeimdallCheckpoint
	checkpointProposers       *orderedmap.OrderedMap[string, struct{}]
	missedCheckpointProposers []string

	milestone                *observer.HeimdallMilestone
	prevMilestoneCount       int64
	milestoneProposers       []api.ValidatorV1
	prevMilestoneProposers   []api.ValidatorV1
	missedMilestoneProposers []string

	span observer.HeimdallSpan

	refreshStateTime *time.Duration
}

func NewHeimdallProvider(n network.Network, tendermintURL, heimdallURL, label string, eb *observer.EventBus, interval uint, version uint) *HeimdallProvider {
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
		version:             version,
	}
}

func (h *HeimdallProvider) SetEventBus(bus *observer.EventBus) {
	h.bus = bus
}

func (h *HeimdallProvider) RefreshState(ctx context.Context) error {
	defer timer(h.refreshStateTime)()

	h.logger.Debug().Msg("Refreshing Heimdall state")

	h.refreshBlockBuffer()
	h.refreshMilestone()
	h.refreshMissedMilestoneProposal()
	h.refreshCheckpoint()
	h.refreshMissedCheckpointProposal()
	h.refreshMissedBlockProposal()
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

	if h.missedBlockProposal != nil {
		m := observer.NewMessage(h.Network, h.Label, h.missedBlockProposal)
		h.bus.Publish(ctx, topics.HeimdallMissedBlockProposal, m)
	}

	if h.checkpoint != nil {
		m := observer.NewMessage(h.Network, h.Label, h.checkpoint)
		h.bus.Publish(ctx, topics.Checkpoint, m)
	}

	if len(h.missedCheckpointProposers) > 0 {
		m := observer.NewMessage(h.Network, h.Label, h.missedCheckpointProposers)
		h.bus.Publish(ctx, topics.MissedCheckpointProposal, m)
	}

	if h.milestone != nil {
		m := observer.NewMessage(h.Network, h.Label, h.milestone)
		h.bus.Publish(ctx, topics.Milestone, m)
	}

	if len(h.missedMilestoneProposers) > 0 {
		m := observer.NewMessage(h.Network, h.Label, h.missedMilestoneProposers)
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

	var block observer.HeimdallBlock
	err = api.GetJSON(path, &block)
	if err != nil {
		h.logger.Warn().Err(err).Msg("Failed to get Heimdall block")
		return nil
	}

	return &block
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

	var validators observer.HeimdallValidators
	err = api.GetJSON(path, &validators)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall validators")
		return nil
	}

	return &validators
}

func (h *HeimdallProvider) fillRange(start uint64) {
	h.logger.Debug().
		Uint64("start_block", start).
		Uint64("end_block", h.BlockNumber).
		Msg("Filling block range")

	for i := start + 1; i <= h.BlockNumber; i++ {
		block := h.getBlock(i)
		if block == nil {
			h.logger.Warn().Uint64("block_number", i).Msg("Failed to get block")
			break
		}

		h.blockBuffer.PutBlock(block)
	}
}

func (h *HeimdallProvider) getHeimdallMilestoneCount() (*big.Int, error) {
	path, err := url.JoinPath(h.HeimdallURL, "milestone/count")
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall milestone count path")
		return nil, err
	}

	var count observer.HeimdallMilestoneCount
	switch h.version {
	case 1:
		var v1 observer.HeimdallMilestoneCountV1
		if err := api.GetJSON(path, &v1); err != nil {
			return nil, err
		}
		count = v1.Result
	case 2:
		if err := api.GetJSON(path, &count); err != nil {
			return nil, err
		}
	}

	c, ok := new(big.Int).SetString(count.Count.String(), 10)
	if !ok {
		return nil, errors.New("failed to parse milestone count")
	}

	return c, nil
}

func (h *HeimdallProvider) refreshMilestone() error {
	if h.milestone != nil {
		h.prevMilestoneCount = h.milestone.Count
	}

	count, err := h.getHeimdallMilestoneCount()
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall milestone count")
		return err
	}

	path, err := url.JoinPath(h.HeimdallURL, "milestone", count.String())
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall milestone path")
		return err
	}

	var milestone observer.HeimdallMilestone
	switch h.version {
	case 1:
		var v1 observer.HeimdallMilestoneV1
		if err = api.GetJSON(path, &v1); err == nil {
			milestone = v1.Result
		}
	case 2:
		var v2 observer.HeimdallMilestoneV2
		if err = api.GetJSON(path, &v2); err == nil {
			milestone = v2.Milestone
		}
	}

	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall milestone")
		return err
	}

	h.milestone = &milestone
	h.milestone.PrevCount = h.prevMilestoneCount
	h.milestone.Count = count.Int64()

	return nil
}

func (h *HeimdallProvider) refreshCheckpoint() error {
	path, err := url.JoinPath(h.HeimdallURL, "checkpoints/latest")
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall latest checkpoint path")
		return err
	}

	switch h.version {
	case 1:
		var v1 observer.HeimdallCheckpointV1
		if err = api.GetJSON(path, &v1); err == nil {
			h.checkpoint = &v1.Result
		}
	case 2:
		var v2 observer.HeimdallCheckpointV2
		if err = api.GetJSON(path, &v2); err == nil {
			h.checkpoint = &v2.Checkpoint
		}
	}

	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall latest checkpoint")
		return err
	}

	return nil
}

func (h *HeimdallProvider) getCurrentCheckpointProposer() (api.Validator, error) {
	var proposer api.Validator

	switch h.version {
	case 1:
		path, err := url.JoinPath(h.HeimdallURL, "staking/current-proposer")
		if err != nil {
			return nil, err
		}

		var v1 observer.HeimdallCurrentCheckpointProposerV1
		err = api.GetJSON(path, &v1)
		if err != nil {
			return nil, err
		}

		proposer = v1.Result
	case 2:
		path, err := url.JoinPath(h.HeimdallURL, "checkpoint/proposers/current")
		if err != nil {
			return nil, err
		}

		var v2 observer.HeimdallCurrentCheckpointProposerV2
		err = api.GetJSON(path, &v2)
		if err != nil {
			return nil, err
		}

		proposer = v2.Validator
	}

	return proposer, nil
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

	current, err := h.getCurrentCheckpointProposer()
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall current checkpoint proposer")
		return err
	}

	signer := current.GetSigner()
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
	for i := h.prevBlockNumber + 1; i <= h.BlockNumber && h.prevBlockNumber != 0; i++ {
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

func (h *HeimdallProvider) refreshMissedMilestoneProposal() error {
	if h.version != 1 {
		return nil
	}

	h.missedMilestoneProposers = nil

	path, err := url.JoinPath(h.HeimdallURL, "staking/milestoneProposer", fmt.Sprint(500))
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall milestone proposers path")
		return err
	}

	var proposers *observer.ValidatorsV1
	err = api.GetJSON(path, &proposers)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall milestone proposers")
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
		if validator.Signer == h.milestone.Proposer {
			break
		}

		h.missedMilestoneProposers = append(h.missedMilestoneProposers, validator.Signer)
	}

	if len(h.missedMilestoneProposers) > 0 {
		h.logger.Info().
			Any("missed_milestone_proposers", h.missedMilestoneProposers).
			Msg("Validators missed milestone proposal")
	}

	return nil
}

func (h *HeimdallProvider) refreshSpan() error {
	switch h.version {
	case 1:
		url, err := url.JoinPath(h.HeimdallURL, "bor/latest-span")
		if err != nil {
			h.logger.Error().Err(err).Msg("Failed to get Heimdall v1 latest span path")
			return err
		}

		var v1 observer.HeimdallSpanV1
		err = api.GetJSON(url, &v1)
		if err != nil {
			h.logger.Error().Err(err).Msg("Failed to get Heimdall v1 latest span")
			return err
		}

		h.span = v1
	case 2:
		url, err := url.JoinPath(h.HeimdallURL, "bor/span/latest")
		if err != nil {
			h.logger.Error().Err(err).Msg("Failed to get Heimdall v2 latest span path")
			return err
		}

		var v2 observer.HeimdallSpanV2
		err = api.GetJSON(url, &v2)
		if err != nil {
			h.logger.Error().Err(err).Msg("Failed to get Heimdall v2 latest span")
			return err
		}

		h.span = v2
	}

	return nil
}
