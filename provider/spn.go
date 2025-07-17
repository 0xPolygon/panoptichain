package provider

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/network"
	"github.com/0xPolygon/panoptichain/observer"
	"github.com/0xPolygon/panoptichain/observer/topics"
	"github.com/0xPolygon/panoptichain/proto"
)

type SuccinctProverNetworkProvider struct {
	url              string
	network          network.Network
	label            string
	bus              *observer.EventBus
	interval         time.Duration
	logger           zerolog.Logger
	refreshStateTime *time.Duration
	apiKey           string
	requesters       []string
	usageSummaries   []*observer.UsageSummary
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
		apiKey:           cfg.APIKey,
		requesters:       cfg.Requesters,
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

	c := proto.NewProverNetworkClient(conn)

	r.refreshUsageSummaries(ctx, c)

	return nil
}

func (r *SuccinctProverNetworkProvider) refreshUsageSummaries(ctx context.Context, c proto.ProverNetworkClient) {
	r.usageSummaries = nil
	now := time.Now()

	for _, requester := range r.requesters {
		req := &proto.GetRequesterUsageRequest{
			StartTime: now.Truncate(time.Hour).Format(time.RFC3339),
			EndTime:   now.Format(time.RFC3339),
			Requester: requester,
		}

		ctx := metadata.AppendToOutgoingContext(ctx, "api-key", r.apiKey)
		res, err := c.GetRequesterUsage(ctx, req)
		if err != nil {
			r.logger.Error().Err(err).Msg("Failed to get requester usage")
		}

		if len(res.UsageSummary) == 0 {
			r.logger.Warn().Msg("Failed to get UsageSummary")
			continue
		}

		if len(res.UsageSummary) > 1 {
			r.logger.Warn().Msg("Expected only one UsageSummary")
		}

		usageSummary := &observer.UsageSummary{
			// The first UsageSummary will be the latest hour bucket.
			UsageSummary: res.UsageSummary[0].UsageSummary,
			Requester:    requester,
		}

		r.usageSummaries = append(r.usageSummaries, usageSummary)

		r.logger.Info().Any("request", req).Any("response", res).Send()
	}

	r.logger.Info().Any("usage_summaries", r.usageSummaries).Send()
}

func (r *SuccinctProverNetworkProvider) PublishEvents(ctx context.Context) error {
	for _, usageSummary := range r.usageSummaries {
		msg := observer.NewMessage(r.network, r.label, usageSummary)
		r.bus.Publish(ctx, topics.UsageSummary, msg)
	}

	return nil
}

func (r *SuccinctProverNetworkProvider) PollingInterval() time.Duration {
	return r.interval
}
