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
	latestMilestone    *observer.HeimdallMilestone

	spans *observer.HeimdallSpans
	// spanUpdates holds the Prev/Curr snapshots to publish this cycle. refreshSpan
	// walks every span from the last known one up to the latest, so each
	// consecutive pair is published and overlap detection sees the full sequence.
	spanUpdates []*observer.HeimdallSpans

	validatorSets *observer.HeimdallValidatorSets

	missedVotes    []*observer.HeimdallMissedVotes
	validatorIDMap map[string]uint64 // normalized signer_address -> val_id

	refreshStateTime *time.Duration

	bufferedCheckpoint *observer.HeimdallCheckpoint
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

	// The runner bounds this cycle with a deadline and api.GetJSON honours it, so
	// a slow/degraded upstream cannot make RefreshState run unbounded. The
	// freshness milestone is fetched first inside refreshMilestone, so a
	// truncated cycle only drops backlog accounting, never the freshness gauge.
	h.logger.Debug().Msg("Refreshing Heimdall state")

	anchorTime := h.refreshBlockBuffer(ctx)
	h.refreshValidatorSet(ctx)
	h.refreshMilestone(ctx, anchorTime)
	h.refreshCheckpoint(ctx)
	h.refreshBufferedCheckpoint(ctx)
	h.refreshMissedCheckpointProposal(ctx)
	h.refreshMissedBlockProposal(ctx)
	h.refreshSpan(ctx)
	h.refreshMissedVotes(ctx)

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

	// Publish the tip milestone (drives the freshness gauges) independently of
	// the per-milestone backlog stream below, so freshness never lags a slow
	// catch-up.
	if h.latestMilestone != nil {
		latest := observer.NewMessage(h.network, h.label, h.latestMilestone)
		h.bus.Publish(ctx, topics.MilestoneLatest, latest)
	}

	for _, milestone := range h.milestones {
		m := observer.NewMessage(h.network, h.label, milestone)
		h.bus.Publish(ctx, topics.Milestone, m)
	}

	for _, s := range h.spanUpdates {
		m := observer.NewMessage(h.network, h.label, s)
		h.bus.Publish(ctx, topics.Span, m)
	}

	if h.validatorSets != nil {
		m := observer.NewMessage(h.network, h.label, h.validatorSets)
		h.bus.Publish(ctx, topics.ValidatorSet, m)
	}

	// Always publish buffered checkpoint (observer handles nil)
	h.bus.Publish(ctx, topics.BufferedCheckpoint, observer.NewMessage(h.network, h.label, h.bufferedCheckpoint))

	for _, mv := range h.missedVotes {
		if mv.MissingCount > 0 {
			m := observer.NewMessage(h.network, h.label, mv)
			h.bus.Publish(ctx, topics.MissedVote, m)
		}
	}

	h.bus.Publish(ctx, topics.RefreshStateTime, observer.NewMessage(h.network, h.label, h.refreshStateTime))

	return nil
}

func (h *HeimdallProvider) PollingInterval() time.Duration {
	return h.interval
}

// refreshBlockBuffer refreshes the block buffer and returns the latest block's
// timestamp (0 if unavailable) so callers can reuse it as a vote-height anchor
// without refetching the tip block.
func (h *HeimdallProvider) refreshBlockBuffer(ctx context.Context) uint64 {
	h.prevBlockNumber = h.blockNumber
	block := h.getBlock(ctx, 0)
	if block == nil {
		return 0
	}

	bn := block.Number()
	if bn == nil {
		return 0
	}
	h.blockNumber = bn.Uint64()

	h.logger.Debug().Uint64("block_number", h.blockNumber).Msg("Refreshing Heimdall state")
	if h.prevBlockNumber != 0 && h.prevBlockNumber != h.blockNumber {
		h.fillRange(ctx, h.prevBlockNumber)
	}

	anchorTime, err := block.Time()
	if err != nil {
		return 0
	}

	return anchorTime
}

func (h *HeimdallProvider) getBlock(ctx context.Context, height uint64) *observer.HeimdallBlock {
	path, err := url.JoinPath(h.tendermintURL, "block")
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to join path when fetching Heimdall block")
		return nil
	}

	if height > 0 {
		path = fmt.Sprintf("%s?height=%d", path, height)
	}

	var block observer.HeimdallBlock
	err = api.GetJSON(ctx, path, &block)
	if err != nil {
		h.logger.Warn().Err(err).Msg("Failed to get Heimdall block")
		return nil
	}

	return &block
}

func (h *HeimdallProvider) getValidators(ctx context.Context, height uint64) *observer.HeimdallValidators {
	path, err := url.JoinPath(h.tendermintURL, "validators")
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to join path when fetching Heimdall validators")
		return nil
	}

	if height > 0 {
		path = fmt.Sprintf("%s?height=%d", path, height)
	}

	var validators observer.HeimdallValidators
	err = api.GetJSON(ctx, path, &validators)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall validators")
		return nil
	}

	return &validators
}

// getValidatorsAtHeight fetches all validators at a specific height with pagination.
func (h *HeimdallProvider) getValidatorsAtHeight(ctx context.Context, height uint64) ([]*observer.HeimdallValidator, error) {
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
		if err := api.GetJSON(ctx, path, &validators); err != nil {
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

func (h *HeimdallProvider) fillRange(ctx context.Context, start uint64) {
	h.logger.Debug().
		Uint64("start_block", start).
		Uint64("end_block", h.blockNumber).
		Msg("Filling block range")

	for i := start; i <= h.blockNumber; i++ {
		block := h.getBlock(ctx, i)
		if block == nil {
			h.logger.Warn().Uint64("block_number", i).Msg("Failed to get block")
			break
		}

		h.blockBuffer.PutBlock(block)
	}
}

func (h *HeimdallProvider) refreshMilestone(ctx context.Context, anchorTime uint64) error {
	h.milestones = nil

	// Always fetch the tip milestone first; it drives the freshness gauges
	// regardless of how far behind the per-milestone backfill below is.
	latest, currentCount, err := h.getLatestMilestone(ctx)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get latest Heimdall milestone")
		return err
	}

	// Keep the cursor monotonic: a load-balanced node briefly reporting a lower
	// count must not rewind it, or already-counted milestones get re-published
	// and double-counted when the count recovers.
	if currentCount <= h.prevMilestoneCount {
		return nil
	}

	// Advance the freshness tip only on forward progress, so a stale lower tip
	// can't rewind the freshness gauges either. When there is no new milestone
	// the previous tip is still republished each cycle, so
	// time_since_last_milestone keeps climbing.
	if latest != nil {
		h.latestMilestone = latest
	}

	// On first poll, baseline to the tip; otherwise backfill new milestones.
	start := h.prevMilestoneCount + 1
	if h.prevMilestoneCount == 0 {
		start = currentCount
	}

	voteCache := make(map[uint64]*observer.HeimdallMilestoneVotes)

	// Walk milestones one at a time, advancing the cursor only over those we
	// actually process. The per-cycle deadline bounds the work; a catch-up
	// truncated by the deadline resumes from where it stopped next cycle. A
	// milestone that fails to fetch on its own (pruned/404/malformed) is skipped
	// so it can't stall every later milestone behind it.
	processed := start - 1
	for i := start; i <= currentCount; i++ {
		if ctx.Err() != nil {
			h.logger.Warn().
				Err(ctx.Err()).
				Int64("from", i).
				Int64("to", currentCount).
				Msg("Milestone refresh deadline reached; resuming next cycle")
			break
		}

		milestone, err := h.getMilestone(ctx, i, latest, currentCount)
		if err != nil {
			// Deadline hit mid-fetch: resume from here next cycle. Any other
			// error is specific to this milestone, so skip past it.
			if ctx.Err() != nil {
				h.logger.Warn().
					Err(err).
					Int64("milestone", i).
					Msg("Milestone refresh deadline reached; resuming next cycle")
				break
			}

			h.logger.Warn().
				Err(err).
				Int64("milestone", i).
				Msg("Failed to get Heimdall milestone; skipping")
			processed = i
			continue
		}

		milestone.Votes = h.findMilestoneVotes(ctx, milestone, anchorTime, voteCache)

		h.milestones = append(h.milestones, milestone)
		processed = i
	}

	h.prevMilestoneCount = processed
	return nil
}

// getMilestone returns milestone i. The tip (i == currentCount) is reused from
// latest when available so the backfill loop doesn't refetch a milestone
// getLatestMilestone already fetched this cycle.
func (h *HeimdallProvider) getMilestone(ctx context.Context, i int64, latest *observer.HeimdallMilestone, currentCount int64) (*observer.HeimdallMilestone, error) {
	if i == currentCount && latest != nil {
		return latest, nil
	}

	path, err := url.JoinPath(h.heimdallURL, "milestones", strconv.FormatInt(i, 10))
	if err != nil {
		return nil, err
	}

	var v2 observer.HeimdallMilestoneV2
	if err := api.GetJSON(ctx, path, &v2); err != nil {
		return nil, err
	}

	milestone := &v2.Milestone
	milestone.Count = i

	return milestone, nil
}

// getLatestMilestone fetches the current tip milestone along with the total
// milestone count. The milestone's Count is set to the total so downstream
// freshness gauges report the tip index.
func (h *HeimdallProvider) getLatestMilestone(ctx context.Context) (*observer.HeimdallMilestone, int64, error) {
	countPath, err := url.JoinPath(h.heimdallURL, "milestones", "count")
	if err != nil {
		return nil, 0, err
	}

	var count observer.HeimdallMilestoneCount
	if err := api.GetJSON(ctx, countPath, &count); err != nil {
		return nil, 0, err
	}
	currentCount := int64(count.Count)
	if currentCount <= 0 {
		// Fresh chain with no milestones yet; nothing to report.
		return nil, 0, nil
	}

	// Prefer the dedicated latest endpoint, but fall back to fetching the tip by
	// count when it errors or returns an empty/zero-value body (e.g. a 200 with
	// an unexpected shape), so a provider that doesn't serve a usable
	// /milestones/latest still keeps the freshness gauges alive.
	var v2 observer.HeimdallMilestoneV2
	latestPath, err := url.JoinPath(h.heimdallURL, "milestones", "latest")
	if err != nil {
		return nil, currentCount, err
	}
	if err := api.GetJSON(ctx, latestPath, &v2); err != nil || v2.Milestone.Timestamp == 0 {
		if err != nil {
			h.logger.Warn().
				Err(err).
				Msg("milestones/latest unavailable; falling back to fetch by count")
		} else {
			h.logger.Warn().
				Msg("milestones/latest returned an empty body; falling back to fetch by count")
		}

		byCount, err := url.JoinPath(h.heimdallURL, "milestones", strconv.FormatInt(currentCount, 10))
		if err != nil {
			return nil, currentCount, err
		}
		if err := api.GetJSON(ctx, byCount, &v2); err != nil {
			return nil, currentCount, err
		}
	}

	milestone := &v2.Milestone
	if milestone.Timestamp == 0 {
		// Still empty after the fallback; don't drive the freshness gauge
		// (time.Since of a zero timestamp) from it. Backfill by count still proceeds.
		return nil, currentCount, nil
	}
	milestone.Count = currentCount

	return milestone, currentCount, nil
}

func (h *HeimdallProvider) refreshCheckpoint(ctx context.Context) error {
	path, err := url.JoinPath(h.heimdallURL, "checkpoints", "latest")
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall latest checkpoint path")
		return err
	}

	var v2 observer.HeimdallCheckpointV2
	if err = api.GetJSON(ctx, path, &v2); err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall latest checkpoint")
		return err
	}

	h.checkpoint = &v2.Checkpoint

	return nil
}

func (h *HeimdallProvider) refreshBufferedCheckpoint(ctx context.Context) error {
	path, err := url.JoinPath(h.heimdallURL, "checkpoints", "buffer")
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall buffered checkpoint path")
		return err
	}

	var v2 observer.HeimdallCheckpointV2
	if err = api.GetJSON(ctx, path, &v2); err != nil {
		h.logger.Warn().Err(err).Msg("Failed to get Heimdall buffered checkpoint")
		h.bufferedCheckpoint = nil
		return err
	}

	// API returns zero ID when no buffered checkpoint exists
	if v2.Checkpoint.ID == 0 {
		h.bufferedCheckpoint = nil
	} else {
		h.bufferedCheckpoint = &v2.Checkpoint
	}

	return nil
}

func (h *HeimdallProvider) getCurrentCheckpointProposer(ctx context.Context) (string, error) {
	path, err := url.JoinPath(h.heimdallURL, "checkpoints", "prepare-next")
	if err != nil {
		return "", err
	}

	var resp observer.HeimdallPrepareNextCheckpoint
	if err = api.GetJSON(ctx, path, &resp); err != nil {
		return "", err
	}

	return resp.Checkpoint.Proposer, nil
}

func (h *HeimdallProvider) refreshMissedCheckpointProposal(ctx context.Context) error {
	var proposers []string
	for pair := h.checkpointProposers.Oldest(); pair != nil; pair = pair.Next() {
		proposers = append(proposers, pair.Key)
	}

	h.logger.Debug().
		Any("checkpoint_proposers", proposers).
		Any("missed_checkpoint_proposers", h.missedCheckpointProposers).
		Msg("Refreshing missed checkpoint proposal")

	h.missedCheckpointProposers = nil

	signer, err := h.getCurrentCheckpointProposer(ctx)
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

func (h *HeimdallProvider) refreshMissedBlockProposal(ctx context.Context) error {
	missedBlockProposal := make(observer.HeimdallMissedBlockProposal)
	for i := h.prevBlockNumber + 1; i <= h.blockNumber && h.prevBlockNumber != 0; i++ {
		if ctx.Err() != nil {
			// Deadline reached mid-scan; stop rather than hammering failing
			// requests. The remaining blocks are a hole this cycle.
			h.logger.Warn().
				Err(ctx.Err()).
				Uint64("from", i).
				Uint64("to", h.blockNumber).
				Msg("Missed-block-proposal scan deadline reached; skipping remaining")
			break
		}

		block := h.getBlock(ctx, i)
		if block == nil {
			h.logger.Debug().
				Uint64("height", i).
				Msg("Failed to get current block")
			continue
		}
		proposer := block.ProposerAddress()

		v := h.getValidators(ctx, i-1)
		if v == nil {
			h.logger.Debug().
				Uint64("height", i).
				Msg("Failed to get validators")
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

func (h *HeimdallProvider) refreshSpan(ctx context.Context) error {
	h.spanUpdates = nil

	// Always fetch the latest span
	latest, err := h.getLatestSpan(ctx)
	if err != nil {
		return err
	}

	// Set current span on startup
	if h.spans.Curr == nil {
		h.spans.Curr = latest
		h.recordSpanUpdate()
		return nil
	}

	// No new span: nothing to publish. The span gauges retain their last value,
	// so republishing would only make the observer re-count the same overlap
	// every idle cycle.
	if latest.ID <= h.spans.Curr.ID {
		return nil
	}

	// If we've fallen too far behind, jump straight to the latest span rather
	// than walking every intermediate one.
	if lag := latest.ID - h.spans.Curr.ID; lag > h.maxSpanLag {
		h.logger.Warn().
			Uint64("current_span_id", h.spans.Curr.ID).
			Uint64("latest_span_id", latest.ID).
			Uint64("lag", lag).
			Uint64("max_span_lag", h.maxSpanLag).
			Msg("Span lag exceeds maximum, jumping to latest")

		h.spans.Prev = h.spans.Curr
		h.spans.Curr = latest
		h.recordSpanUpdate()
		return nil
	}

	// Walk every new span up to the latest so overlap detection sees each
	// consecutive pair (mirrors refreshMilestone). Each step is published as its
	// own Prev/Curr snapshot; the final one carries the latest span forward.
	for id := h.spans.Curr.ID + 1; id <= latest.ID; id++ {
		span := latest
		if id != latest.ID {
			span, err = h.getSpan(ctx, id)
			if err != nil {
				if ctx.Err() != nil {
					// Deadline reached mid-walk; resume from here next cycle.
					h.logger.Warn().
						Uint64("span_id", id).
						Err(err).
						Msg("Span walk deadline reached; resuming next cycle")
					return nil
				}

				// Skip a span that fails on its own (pruned/404/transient) so one
				// bad span can't freeze span progress until the lag-jump kicks in.
				h.logger.Warn().
					Uint64("span_id", id).
					Err(err).
					Msg("Failed to fetch span; skipping")
				continue
			}
		}

		h.spans.Prev = h.spans.Curr
		h.spans.Curr = span
		h.recordSpanUpdate()
	}

	return nil
}

// recordSpanUpdate snapshots the current Prev/Curr span pair for publishing this
// cycle. Each call captures the pointers as they stand, so a walk over several
// spans records every consecutive pair in order.
func (h *HeimdallProvider) recordSpanUpdate() {
	h.spanUpdates = append(h.spanUpdates, &observer.HeimdallSpans{Prev: h.spans.Prev, Curr: h.spans.Curr})
}

func (h *HeimdallProvider) getLatestSpan(ctx context.Context) (*observer.HeimdallSpan, error) {
	return h.fetchSpan(ctx, "latest")
}

func (h *HeimdallProvider) getSpan(ctx context.Context, id uint64) (*observer.HeimdallSpan, error) {
	return h.fetchSpan(ctx, strconv.FormatUint(id, 10))
}

func (h *HeimdallProvider) fetchSpan(ctx context.Context, spanID string) (*observer.HeimdallSpan, error) {
	path, err := url.JoinPath(h.heimdallURL, "bor", "spans", spanID)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get Heimdall span path")
		return nil, err
	}

	var v2 observer.HeimdallSpanV2
	if err = api.GetJSON(ctx, path, &v2); err != nil {
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

func (h *HeimdallProvider) refreshValidatorSet(ctx context.Context) error {
	validators, err := api.GetValidators(ctx, h.heimdallURL)
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

func (h *HeimdallProvider) getCommit(ctx context.Context, height uint64) (*observer.HeimdallCommit, error) {
	path, err := url.JoinPath(h.tendermintURL, "commit")
	if err != nil {
		return nil, fmt.Errorf("failed to join commit path: %w", err)
	}

	var commit observer.HeimdallCommit
	if err := api.GetJSON(ctx, fmt.Sprintf("%s?height=%d", path, height), &commit); err != nil {
		return nil, fmt.Errorf("failed to get commit at height %d: %w", height, err)
	}

	return &commit, nil
}

func normalizeAddress(addr string) string {
	return strings.ToLower(strings.TrimPrefix(addr, "0x"))
}

func (h *HeimdallProvider) getMissedVotes(ctx context.Context, height uint64) (*observer.HeimdallMissedVotes, error) {
	validators, err := h.getValidatorsAtHeight(ctx, height)
	if err != nil {
		return nil, fmt.Errorf("failed to get validators at height %d: %w", height, err)
	}

	commit, err := h.getCommit(ctx, height)
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

func (h *HeimdallProvider) refreshMissedVotes(ctx context.Context) {
	if h.validatorIDMap == nil {
		return
	}

	h.missedVotes = nil

	h.logger.Debug().Msg("Refreshing missed consensus votes")

	for height := h.prevBlockNumber + 1; height <= h.blockNumber && h.prevBlockNumber != 0; height++ {
		if ctx.Err() != nil {
			// Deadline reached mid-scan; stop rather than hammering failing
			// requests. The remaining blocks are a hole this cycle.
			h.logger.Warn().
				Err(ctx.Err()).
				Uint64("from", height).
				Uint64("to", h.blockNumber).
				Msg("Missed-votes scan deadline reached; skipping remaining")
			break
		}

		mv, err := h.getMissedVotes(ctx, height)
		if err != nil {
			h.logger.Warn().
				Err(err).
				Uint64("height", height).
				Msg("Failed to detect missed votes")
			continue
		}
		if mv != nil {
			h.missedVotes = append(h.missedVotes, mv)
		}
	}
}

// getExtendedCommitInfo fetches and decodes the ExtendedCommitInfo from txs[0]
// of a Heimdall block. Vote extensions from block H-1 are stored in block H's txs[0].
func (h *HeimdallProvider) getExtendedCommitInfo(ctx context.Context, height uint64) (*heimdall.ExtendedCommitInfo, error) {
	block := h.getBlock(ctx, height)
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
func (h *HeimdallProvider) getMilestoneVotes(ctx context.Context, height uint64) (*observer.HeimdallMilestoneVotes, error) {
	extCommit, err := h.getExtendedCommitInfo(ctx, height)
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

// getBlockHeight estimates the Heimdall block height for a given timestamp,
// anchored on the latest block's number and time (anchorTime). Assumes ~2
// second block time.
func (h *HeimdallProvider) getBlockHeight(target int64, anchorTime uint64) uint64 {
	if h.blockNumber == 0 || anchorTime == 0 {
		return 0
	}

	// Estimate based on ~2 second block time
	diff := (int64(anchorTime) - target) / 2
	if diff < 0 {
		return 0
	}

	if uint64(diff) > h.blockNumber {
		return 0
	}

	return h.blockNumber - uint64(diff)
}

// findMilestoneVotes searches for votes matching this milestone's range.
// Uses the milestone timestamp to estimate the finalization block.
// Returns the first matching vote block, or nil if not found. The cache
// memoizes successful per-height lookups across milestones within the same
// cycle; failed fetches are not cached so a later milestone can retry them.
func (h *HeimdallProvider) findMilestoneVotes(ctx context.Context, milestone *observer.HeimdallMilestone, anchorTime uint64, cache map[uint64]*observer.HeimdallMilestoneVotes) *observer.HeimdallMilestoneVotes {
	if h.validatorIDMap == nil {
		return nil
	}

	height := h.getBlockHeight(milestone.Timestamp, anchorTime)
	if height == 0 {
		return nil
	}

	// Search a small window around the estimated height
	const window = 5
	start := max(1, height-window)
	end := min(h.blockNumber, height+window)

	for i := start; i <= end; i++ {
		mv, ok := cache[i]
		if !ok {
			votes, err := h.getMilestoneVotes(ctx, i)
			if err != nil {
				// Don't cache the error; a later milestone can retry this height.
				continue
			}
			mv = votes
			cache[i] = mv
		}

		if mv == nil || mv.MilestoneVoters == 0 {
			continue
		}

		if h.votesMatchMilestone(mv, milestone) {
			return mv
		}
	}

	return nil
}

// votesMatchMilestone checks if any vote in the votes matches the milestone's block range.
func (h *HeimdallProvider) votesMatchMilestone(mv *observer.HeimdallMilestoneVotes, milestone *observer.HeimdallMilestone) bool {
	for _, vote := range mv.Votes {
		if !vote.HasMilestone {
			continue
		}

		if vote.MilestoneStart == milestone.StartBlock && vote.MilestoneEnd == milestone.EndBlock {
			return true
		}
	}

	return false
}
