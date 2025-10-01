package provider

import (
	"context"
	"crypto/tls"
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
	start            time.Time
	apiKey           string
	requester        *string
	fulfiller        *string
	proofRequests    []*proto.ProofRequest
	seen             map[string]time.Time
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
		seen:             make(map[string]time.Time),
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

	r.refreshProofRequests(ctx, c)

	return nil
}

func (r *SuccinctProverNetworkProvider) refreshProofRequests(ctx context.Context, c proto.ProverNetworkClient) {
	r.proofRequests = nil

	limit := uint32(100)

	req := &proto.GetFilteredProofRequestsRequest{
		Limit:             &limit,
		ExecutionStatus:   proto.ExecutionStatus_EXECUTED.Enum(),
		FulfillmentStatus: proto.FulfillmentStatus_FULFILLED.Enum(),
	}

	if r.requester != nil {
		req.Requester = common.HexToAddress(*r.requester).Bytes()
	}

	if r.fulfiller != nil {
		req.Fulfiller = common.HexToAddress(*r.fulfiller).Bytes()
	}

	res, err := c.GetFilteredProofRequests(metadata.AppendToOutgoingContext(ctx, "api-key", r.apiKey), req)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to get requester usage")
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

	return nil
}

func (r *SuccinctProverNetworkProvider) PollingInterval() time.Duration {
	return r.interval
}
