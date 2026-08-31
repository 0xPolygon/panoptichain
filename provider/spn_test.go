package provider

import (
	"context"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/0xPolygon/panoptichain/config"
	spnpb "github.com/0xPolygon/panoptichain/proto/network"
)

// newUsageServer starts a gRPC server that answers GetRequesterUsage with the
// given response, and returns a connection to it along with the last request it
// received. The service is registered by hand because the protos this package
// is generated from do not declare the method on ProverNetwork.
func newUsageServer(t *testing.T, res *spnpb.GetRequesterUsageResponse) (*grpc.ClientConn, *spnpb.GetRequesterUsageRequest) {
	t.Helper()

	got := new(spnpb.GetRequesterUsageRequest)

	desc := grpc.ServiceDesc{
		ServiceName: "network.ProverNetwork",
		HandlerType: (*any)(nil),
		Methods: []grpc.MethodDesc{{
			MethodName: "GetRequesterUsage",
			Handler: func(_ any, ctx context.Context, dec func(any) error, _ grpc.UnaryServerInterceptor) (any, error) {
				if err := dec(got); err != nil {
					return nil, err
				}
				return res, nil
			},
		}},
	}

	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}

	srv := grpc.NewServer()
	srv.RegisterService(&desc, new(struct{}))
	go func() { _ = srv.Serve(lis) }()
	t.Cleanup(srv.Stop)

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	t.Cleanup(func() { _ = conn.Close() })

	return conn, got
}

func newUsageProvider(usageRequesters ...string) *SuccinctProverNetworkProvider {
	requesters := make([]config.UsageRequester, 0, len(usageRequesters))
	for _, address := range usageRequesters {
		requesters = append(requesters, config.UsageRequester{Address: address})
	}

	return &SuccinctProverNetworkProvider{
		logger:          NewLogger(nil, "test"),
		usageRequesters: requesters,
	}
}

func TestRefreshRequesterUsage_PicksNewestHour(t *testing.T) {
	conn, got := newUsageServer(t, &spnpb.GetRequesterUsageResponse{
		UsageSummary: []*spnpb.RequesterUsageSummary{
			{
				Hour:         "2026-08-06T11:00:00Z",
				UsageSummary: &spnpb.UsageSummary{ReservedGas: "100", OnDemandGas: "200"},
			},
			{
				Hour:         "2026-08-06T12:00:00Z",
				UsageSummary: &spnpb.UsageSummary{ReservedGas: "300", OnDemandGas: "400"},
			},
		},
	})

	requester := "0x5428abf0e5aec1be48597a984a4f9570d9236f29"
	h := newUsageProvider(requester)
	h.refreshRequesterUsage(context.Background(), conn)

	if len(h.usage) != 1 {
		t.Fatalf("expected one usage summary, got %d", len(h.usage))
	}
	usage := h.usage[0]
	if usage.Hour != "2026-08-06T12:00:00Z" {
		t.Fatalf("expected newest hour, got %q", usage.Hour)
	}
	if usage.ReservedGas != "300" || usage.OnDemandGas != "400" {
		t.Fatalf("unexpected gas: reserved=%q on_demand=%q", usage.ReservedGas, usage.OnDemandGas)
	}

	// The request must round-trip over the wire with the field numbers taken
	// from sp1's bindings; a mismatch would decode to empty strings here.
	if got.Requester == "" || got.StartTime == "" || got.EndTime == "" {
		t.Fatalf("request did not decode server-side: %+v", got)
	}
	// Unlike the other requester filters on this service, this one is a hex
	// string rather than raw address bytes. It carries the 0x prefix the
	// server requires, and stays lowercased so the metric label does not
	// change with however the address is written in config.
	if got.Requester != requester {
		t.Fatalf("unexpected requester encoding: %q, want %q", got.Requester, requester)
	}
}

func TestRefreshRequesterUsage_EmptyHourLeavesUsageNil(t *testing.T) {
	conn, _ := newUsageServer(t, &spnpb.GetRequesterUsageResponse{})

	requester := "0x5428abf0e5aec1be48597a984a4f9570d9236f29"
	h := newUsageProvider(requester)
	h.refreshRequesterUsage(context.Background(), conn)

	if len(h.usage) != 0 {
		t.Fatalf("expected no usage for an empty hour, got %+v", h.usage)
	}
}

func TestRefreshRequesterUsage_NoRequesterSkipsCall(t *testing.T) {
	conn, got := newUsageServer(t, &spnpb.GetRequesterUsageResponse{
		UsageSummary: []*spnpb.RequesterUsageSummary{{
			Hour:         "2026-08-06T12:00:00Z",
			UsageSummary: &spnpb.UsageSummary{ReservedGas: "1", OnDemandGas: "2"},
		}},
	})

	h := newUsageProvider()
	h.refreshRequesterUsage(context.Background(), conn)

	if len(h.usage) != 0 {
		t.Fatalf("expected no usage without a configured requester, got %+v", h.usage)
	}
	if got.Requester != "" {
		t.Fatal("expected no RPC to be made without a configured requester")
	}
}

// Usage is tracked per requester, so several configured requesters each get
// their own summary rather than collapsing into one.
func TestRefreshRequesterUsage_TracksEachRequester(t *testing.T) {
	conn, _ := newUsageServer(t, &spnpb.GetRequesterUsageResponse{
		UsageSummary: []*spnpb.RequesterUsageSummary{{
			Hour:         "2026-08-06T12:00:00Z",
			UsageSummary: &spnpb.UsageSummary{ReservedGas: "1", OnDemandGas: "2"},
		}},
	})

	requesters := []string{
		"0x5428abf0e5aec1be48597a984a4f9570d9236f29",
		"0xafb1d2c26654c85f51f550c97f16699da0293dee",
	}

	h := newUsageProvider(requesters...)
	h.refreshRequesterUsage(context.Background(), conn)

	if len(h.usage) != len(requesters) {
		t.Fatalf("expected %d usage summaries, got %d", len(requesters), len(h.usage))
	}
	for i, want := range requesters {
		if h.usage[i].Requester != want {
			t.Fatalf("usage[%d]: got requester %q, want %q", i, h.usage[i].Requester, want)
		}
	}
}

// The tag is config, not an observation, so it must ride along to the observer
// that turns it into a metric label.
func TestRefreshRequesterUsage_CarriesRequesterTag(t *testing.T) {
	conn, _ := newUsageServer(t, &spnpb.GetRequesterUsageResponse{
		UsageSummary: []*spnpb.RequesterUsageSummary{{
			Hour:         "2026-08-06T12:00:00Z",
			UsageSummary: &spnpb.UsageSummary{ReservedGas: "1", OnDemandGas: "2"},
		}},
	})

	h := newUsageProvider()
	h.usageRequesters = []config.UsageRequester{{
		Address: "0x5428abf0e5aec1be48597a984a4f9570d9236f29",
		Tag:     "katana",
	}}

	h.refreshRequesterUsage(context.Background(), conn)

	if len(h.usage) != 1 {
		t.Fatalf("expected one usage summary, got %d", len(h.usage))
	}
	if h.usage[0].Tag != "katana" {
		t.Fatalf("tag = %q, want katana", h.usage[0].Tag)
	}
}

// Pricing is optional. Without it the summary carries a zero rate, which the
// observer reads as "gas only".
func TestRefreshRequesterUsage_PricingIsOptional(t *testing.T) {
	res := &spnpb.GetRequesterUsageResponse{
		UsageSummary: []*spnpb.RequesterUsageSummary{{
			Hour:         "2026-08-06T12:00:00Z",
			UsageSummary: &spnpb.UsageSummary{ReservedGas: "1", OnDemandGas: "2"},
		}},
	}
	requester := "0x5428abf0e5aec1be48597a984a4f9570d9236f29"

	conn, _ := newUsageServer(t, res)
	h := newUsageProvider(requester)
	h.refreshRequesterUsage(context.Background(), conn)

	if len(h.usage) != 1 || h.usage[0].RatePerBillionGas != 0 {
		t.Fatalf("expected a zero rate without pricing, got %+v", h.usage)
	}

	h = newUsageProvider(requester)
	h.pricing = &config.SuccinctPricing{RatePerBillionGas: 0.5, Currency: "USD"}
	h.refreshRequesterUsage(context.Background(), conn)

	if len(h.usage) != 1 {
		t.Fatalf("expected one usage summary, got %d", len(h.usage))
	}
	if h.usage[0].RatePerBillionGas != 0.5 || h.usage[0].Currency != "USD" {
		t.Fatalf("pricing not carried through: %+v", h.usage[0])
	}
}

// The proof-request filter and the usage requesters are independent: setting
// only Requester must not cause a usage call, since narrowing proof requests
// says nothing about whose gas budget is being tracked.
func TestRefreshRequesterUsage_RequesterFilterDoesNotDriveUsage(t *testing.T) {
	conn, got := newUsageServer(t, &spnpb.GetRequesterUsageResponse{
		UsageSummary: []*spnpb.RequesterUsageSummary{{
			Hour:         "2026-08-06T12:00:00Z",
			UsageSummary: &spnpb.UsageSummary{ReservedGas: "1", OnDemandGas: "2"},
		}},
	})

	requester := "0x5428abf0e5aec1be48597a984a4f9570d9236f29"
	h := newUsageProvider()
	h.requester = &requester

	h.refreshRequesterUsage(context.Background(), conn)

	if len(h.usage) != 0 {
		t.Fatalf("expected no usage from the proof-request filter alone, got %+v", h.usage)
	}
	if got.Requester != "" {
		t.Fatal("expected no RPC to be made from the proof-request filter alone")
	}
}
