package provider

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	orderedmap "github.com/wk8/go-ordered-map/v2"

	"github.com/0xPolygon/panoptichain/api"
	"github.com/0xPolygon/panoptichain/blockbuffer"
	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/network"
	"github.com/0xPolygon/panoptichain/observer"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

type HeimdallProvider struct {
	tendermintURL string
	heimdallURL   string
	network       network.Network
	label         string
	bus           *observer.EventBus
	interval      time.Duration
	logger        zerolog.Logger

	blockNumber         uint64
	prevBlockNumber     uint64
	blockBuffer         *blockbuffer.BlockBuffer
	missedBlockProposal observer.HeimdallMissedBlockProposal

	checkpoint                *observer.HeimdallCheckpoint
	checkpointProposers       *orderedmap.OrderedMap[string, struct{}]
	missedCheckpointProposers []string

	milestone          *observer.HeimdallMilestone
	prevMilestoneCount int64

	span *observer.HeimdallSpan

	refreshStateTime *time.Duration
}

func NewHeimdallProvider(n network.Network, eb *observer.EventBus, cfg config.HeimdallEndpoint) *HeimdallProvider {
	return &HeimdallProvider{
		tendermintURL:       cfg.TendermintURL,
		heimdallURL:         cfg.HeimdallURL,
		network:             n,
		label:               cfg.Label,
		bus:                 eb,
		blockBuffer:         blockbuffer.NewBlockBuffer(128),
		interval:            GetInterval(cfg.Interval),
		logger:              NewLogger(n, cfg.Label),
		checkpointProposers: orderedmap.New[string, struct{}](),
		refreshStateTime:    new(time.Duration),
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

	if h.milestone != nil {
		m := observer.NewMessage(h.network, h.label, h.milestone)
		h.bus.Publish(ctx, topics.Milestone, m)
	}

	if h.span != nil {
		m := observer.NewMessage(h.network, h.label, h.span)
		h.bus.Publish(ctx, topics.Span, m)
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

func (h *HeimdallProvider) getHeimdallMilestoneCount() (*big.Int, error) {
	path, err := url.JoinPath(h.heimdallURL, "milestones", "count")
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall milestone count path")
		return nil, err
	}

	var count observer.HeimdallMilestoneCount
	if err := api.GetJSON(path, &count); err != nil {
		return nil, err
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

	path, err := url.JoinPath(h.heimdallURL, "milestones", count.String())
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall milestone path")
		return err
	}

	var milestone observer.HeimdallMilestone
	var v2 observer.HeimdallMilestoneV2
	if err = api.GetJSON(path, &v2); err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall milestone")
		return err
	}

	milestone = v2.Milestone
	h.milestone = &milestone
	h.milestone.PrevCount = h.prevMilestoneCount
	h.milestone.Count = count.Int64()

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

func (h *HeimdallProvider) getCurrentCheckpointProposer() (*api.Validator, error) {
	path, err := url.JoinPath(h.heimdallURL, "checkpoint", "proposers", "current")
	if err != nil {
		return nil, err
	}

	var proposer observer.HeimdallCurrentCheckpointProposer
	if err = api.GetJSON(path, &proposer); err != nil {
		return nil, err
	}

	return &proposer.Validator, nil
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

	signer := current.Signer
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
	url, err := url.JoinPath(h.heimdallURL, "bor", "span", "latest")
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall latest span path")
		return err
	}

	err = api.GetJSON(url, h.span)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall latest span")
		return err
	}

	return nil
}
