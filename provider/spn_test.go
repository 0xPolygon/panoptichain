package provider

import (
	"context"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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

func newUsageProvider(requester *string) *SuccinctProverNetworkProvider {
	return &SuccinctProverNetworkProvider{
		logger:    NewLogger(nil, "test"),
		requester: requester,
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
	h := newUsageProvider(&requester)
	h.refreshRequesterUsage(context.Background(), conn)

	if h.usage == nil {
		t.Fatal("expected usage to be set")
	}
	if h.usage.Hour != "2026-08-06T12:00:00Z" {
		t.Fatalf("expected newest hour, got %q", h.usage.Hour)
	}
	if h.usage.ReservedGas != "300" || h.usage.OnDemandGas != "400" {
		t.Fatalf("unexpected gas: reserved=%q on_demand=%q", h.usage.ReservedGas, h.usage.OnDemandGas)
	}

	// The request must round-trip over the wire with the field numbers taken
	// from sp1's bindings; a mismatch would decode to empty strings here.
	if got.Requester == "" || got.StartTime == "" || got.EndTime == "" {
		t.Fatalf("request did not decode server-side: %+v", got)
	}
	// Unlike the other requester filters on this service, this one is a hex
	// string rather than raw address bytes, and it must stay lowercased so a
	// case-sensitive comparison on the server still matches.
	if got.Requester != requester {
		t.Fatalf("unexpected requester encoding: %q, want %q", got.Requester, requester)
	}
}

func TestRefreshRequesterUsage_EmptyHourLeavesUsageNil(t *testing.T) {
	conn, _ := newUsageServer(t, &spnpb.GetRequesterUsageResponse{})

	requester := "0x5428abf0e5aec1be48597a984a4f9570d9236f29"
	h := newUsageProvider(&requester)
	h.refreshRequesterUsage(context.Background(), conn)

	if h.usage != nil {
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

	h := newUsageProvider(nil)
	h.refreshRequesterUsage(context.Background(), conn)

	if h.usage != nil {
		t.Fatalf("expected no usage without a configured requester, got %+v", h.usage)
	}
	if got.Requester != "" {
		t.Fatal("expected no RPC to be made without a configured requester")
	}
}
