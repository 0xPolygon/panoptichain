package provider

import (
	"context"
	"crypto/tls"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/network"
	"github.com/0xPolygon/panoptichain/observer"
	"github.com/0xPolygon/panoptichain/observer/topics"
	spnpb "github.com/0xPolygon/panoptichain/proto/network"
)

type SuccinctProverNetworkProvider struct {
	url              string
	network          network.Network
	label            string
	bus              *observer.EventBus
	interval         time.Duration
	logger           zerolog.Logger
	refreshStateTime *time.Duration
	start            time.Time
	apiKey           string
	requester        *string
	fulfiller        *string
	usageRequesters  []config.UsageRequester
	pricing          *config.SuccinctPricing
	proofRequests    []*spnpb.ProofRequest
	seen             map[string]time.Time
	// countedHours is the newest usage hour already reported for each requester,
	// keyed by address. The network buckets usage hourly and every cycle re-reads
	// the same bucket, so this is what lets the observer add each hour to its
	// cumulative counters exactly once instead of once per cycle. It lives here
	// rather than in the observer because a provider's cycle is single-threaded
	// by contract, so it needs no synchronization.
	//
	// Only hours this process observed are recorded: there is no backfill, so gas
	// consumed while panoptichain was down is never counted. Unlike a gauge,
	// which self-heals on the next cycle, a counter's gap is permanent and every
	// later range total inherits it.
	countedHours map[string]string
	usage        []*observer.UsageSummary
}

func NewProverNetworkProvider(n network.Network, eb *observer.EventBus, cfg config.SuccinctProverNetwork) *SuccinctProverNetworkProvider {
	return &SuccinctProverNetworkProvider{
		url:              cfg.URL,
		network:          n,
		label:            cfg.Label,
		bus:              eb,
		interval:         GetInterval(cfg.Interval),
		logger:           NewLogger(n, cfg.Label),
		refreshStateTime: new(time.Duration),
		start:            time.Now(),
		apiKey:           cfg.APIKey,
		requester:        cfg.Requester,
		fulfiller:        cfg.Fulfiller,
		usageRequesters:  cfg.UsageRequesters,
		pricing:          cfg.Pricing,
		seen:             make(map[string]time.Time),
		countedHours:     make(map[string]string),
	}
}

func (r *SuccinctProverNetworkProvider) RefreshState(ctx context.Context) error {
	defer timer(r.refreshStateTime)()

	creds := credentials.NewTLS(&tls.Config{})
	conn, err := grpc.NewClient(r.url, grpc.WithTransportCredentials(creds))
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to create gRPC client")
	}
	defer conn.Close()

	c := spnpb.NewProverNetworkClient(conn)

	r.refreshProofRequests(ctx, c)
	r.refreshRequesterUsage(ctx, conn)

	return nil
}

// refreshRequesterUsage fetches each configured requester's gas usage for the
// most recent complete hour. The network buckets usage hourly, so an in-progress
// hour would report a partial total that only climbs until the hour closes;
// reading one hour behind keeps the gauge on settled numbers.
func (r *SuccinctProverNetworkProvider) refreshRequesterUsage(ctx context.Context, conn *grpc.ClientConn) {
	r.usage = nil

	end := time.Now().UTC().Truncate(time.Hour)
	start := end.Add(-time.Hour)

	// The RPC is requester-scoped and has no "all requesters" mode, so each
	// configured requester costs a call. With none configured, none are made.
	for _, requester := range r.usageRequesters {
		if usage := r.requesterUsage(ctx, conn, requester, start, end); usage != nil {
			r.usage = append(r.usage, usage)
		}
	}
}

// requesterUsage fetches a single requester's usage for the given window,
// returning nil when the hour has no settled usage to report.
func (r *SuccinctProverNetworkProvider) requesterUsage(
	ctx context.Context,
	conn *grpc.ClientConn,
	requester config.UsageRequester,
	start, end time.Time,
) *observer.UsageSummary {
	req := &spnpb.GetRequesterUsageRequest{
		StartTime: start.Format(time.RFC3339),
		EndTime:   end.Format(time.RFC3339),
		// Unlike the other requester filters on this service, this field is a
		// hex string rather than raw address bytes. HexToAddress normalizes
		// whatever the config holds and guarantees the 0x prefix, which the
		// server rejects the request without. Casing is not significant to the
		// server, but this value becomes a metric label, so lowercase it to
		// keep the label stable however the address is written in config.
		Requester: strings.ToLower(common.HexToAddress(requester.Address).Hex()),
	}

	res, err := spnpb.GetRequesterUsage(metadata.AppendToOutgoingContext(ctx, "api-key", r.apiKey), conn, req)
	if err != nil {
		r.logger.Error().Err(err).Str("requester", req.Requester).Msg("Failed to get requester usage")
		return nil
	}

	// An hour with no usage comes back empty rather than zeroed. Leave the
	// gauges at their last value instead of publishing a zero that would be
	// indistinguishable from real idleness.
	if len(res.UsageSummary) == 0 {
		r.logger.Debug().
			Str("requester", req.Requester).
			Str("start_time", req.StartTime).
			Str("end_time", req.EndTime).
			Msg("No requester usage reported for the hour")
		return nil
	}

	// The window spans a single hour, but take the newest bucket rather than
	// assuming the response holds exactly one.
	latest := res.UsageSummary[0]
	for _, summary := range res.UsageSummary[1:] {
		if summary.Hour > latest.Hour {
			latest = summary
		}
	}

	if latest.UsageSummary == nil {
		return nil
	}

	usage := &observer.UsageSummary{
		UsageSummary: latest.UsageSummary,
		Requester:    req.Requester,
		Tag:          requester.Tag,
		Hour:         latest.Hour,
	}

	// Flag a bucket the observer has not been given before, so it can advance
	// its cumulative counters. Hours are RFC 3339 UTC at a fixed width, so they
	// order lexicographically and compare as strings. A response with no hour
	// cannot be told apart from a repeat, so it is never flagged: a flat counter
	// is recoverable, an overstated one is not.
	if latest.Hour != "" && latest.Hour > r.countedHours[req.Requester] {
		r.countedHours[req.Requester] = latest.Hour
		usage.NewHour = true
	}

	// Pricing is optional, so leave the rate at zero when none is configured;
	// the observer reads that as "gas only" and emits no cost.
	if r.pricing != nil {
		usage.RatePerBillionGas = r.pricing.RatePerBillionGas
		usage.Currency = r.pricing.Currency
	}

	return usage
}

func (r *SuccinctProverNetworkProvider) refreshProofRequests(ctx context.Context, c spnpb.ProverNetworkClient) {
	r.proofRequests = nil

	limit := uint32(100)

	req := &spnpb.GetFilteredProofRequestsRequest{
		Limit:             &limit,
		ExecutionStatus:   spnpb.ExecutionStatus_EXECUTED.Enum(),
		FulfillmentStatus: spnpb.FulfillmentStatus_FULFILLED.Enum(),
	}

	if r.requester != nil {
		req.Requester = common.HexToAddress(*r.requester).Bytes()
	}

	if r.fulfiller != nil {
		req.Fulfiller = common.HexToAddress(*r.fulfiller).Bytes()
	}

	res, err := c.GetFilteredProofRequests(metadata.AppendToOutgoingContext(ctx, "api-key", r.apiKey), req)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to get filtered proof requests")
	}

	if res != nil {
		r.proofRequests = res.Requests
	}
}

func (r *SuccinctProverNetworkProvider) PublishEvents(ctx context.Context) error {
	for _, proof := range r.proofRequests {
		if time.Unix(int64(*proof.FulfilledAt), 0).Compare(r.start) < 0 {
			continue
		}

		id := string(proof.RequestId)
		if seen, ok := r.seen[id]; ok {
			if time.Since(seen) > time.Hour {
				delete(r.seen, id)
			}
			continue
		}

		msg := observer.NewMessage(r.network, r.label, proof)
		r.bus.Publish(ctx, topics.ProofRequest, msg)
		r.seen[id] = time.Now()
	}

	for _, usage := range r.usage {
		msg := observer.NewMessage(r.network, r.label, usage)
		r.bus.Publish(ctx, topics.RequesterUsage, msg)
	}

	return nil
}

func (r *SuccinctProverNetworkProvider) Logger() zerolog.Logger {
	return r.logger
}

func (r *SuccinctProverNetworkProvider) PollingInterval() time.Duration {
	return r.interval
}
