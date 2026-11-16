package provider

import (
	"context"
	"net"
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

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
		countedHours:    make(map[string]string),
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
	mustRefreshUsage(t, h, conn)

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
	mustRefreshUsage(t, h, conn)

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
	mustRefreshUsage(t, h, conn)

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
	mustRefreshUsage(t, h, conn)

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

	mustRefreshUsage(t, h, conn)

	if len(h.usage) != 1 {
		t.Fatalf("expected one usage summary, got %d", len(h.usage))
	}
	if h.usage[0].Tag != "katana" {
		t.Fatalf("tag = %q, want katana", h.usage[0].Tag)
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

	mustRefreshUsage(t, h, conn)

	if len(h.usage) != 0 {
		t.Fatalf("expected no usage from the proof-request filter alone, got %+v", h.usage)
	}
	if got.Requester != "" {
		t.Fatal("expected no RPC to be made from the proof-request filter alone")
	}
}

// newUsageServerFor is newUsageServer with a per-call response, so a test can
// change what the network reports between cycles.
func newUsageServerFor(t *testing.T, next func() *spnpb.GetRequesterUsageResponse) *grpc.ClientConn {
	t.Helper()

	desc := grpc.ServiceDesc{
		ServiceName: "network.ProverNetwork",
		HandlerType: (*any)(nil),
		Methods: []grpc.MethodDesc{{
			MethodName: "GetRequesterUsage",
			Handler: func(_ any, ctx context.Context, dec func(any) error, _ grpc.UnaryServerInterceptor) (any, error) {
				if err := dec(new(spnpb.GetRequesterUsageRequest)); err != nil {
					return nil, err
				}
				return next(), nil
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

	return conn
}

func hourResponse(hour string) *spnpb.GetRequesterUsageResponse {
	return &spnpb.GetRequesterUsageResponse{
		UsageSummary: []*spnpb.RequesterUsageSummary{{
			Hour:         hour,
			UsageSummary: &spnpb.UsageSummary{ReservedGas: "1000000000", OnDemandGas: "0", TotalGas: "1000000000"},
		}},
	}
}

// The provider owns the once-only decision behind the cumulative counters. The
// network buckets usage hourly and every cycle re-reads the same bucket, so only
// the first cycle to see an hour may flag it.
func TestRefreshRequesterUsage_FlagsEachHourOnce(t *testing.T) {
	hour := "2026-08-31T10:00:00Z"
	conn := newUsageServerFor(t, func() *spnpb.GetRequesterUsageResponse { return hourResponse(hour) })

	h := newUsageProvider("0x5428abf0e5aec1be48597a984a4f9570d9236f29")

	mustRefreshUsage(t, h, conn)
	if len(h.usage) != 1 || !h.usage[0].NewHour {
		t.Fatalf("first cycle must flag the hour as new, got %+v", h.usage)
	}

	// Later cycles re-read the same bucket and must not flag it again.
	for cycle := 2; cycle <= 4; cycle++ {
		mustRefreshUsage(t, h, conn)
		if len(h.usage) != 1 {
			t.Fatalf("cycle %d: expected one summary, got %d", cycle, len(h.usage))
		}
		if h.usage[0].NewHour {
			t.Fatalf("cycle %d re-flagged an hour already reported", cycle)
		}
	}

	// The hour rolling over flags again.
	hour = "2026-08-31T11:00:00Z"
	mustRefreshUsage(t, h, conn)
	if len(h.usage) != 1 || !h.usage[0].NewHour {
		t.Fatalf("a new hour must be flagged, got %+v", h.usage)
	}
}

// An hour older than one already reported must not be flagged, so a late or
// reordered response cannot double count.
func TestRefreshRequesterUsage_DoesNotFlagOlderHour(t *testing.T) {
	hour := "2026-08-31T12:00:00Z"
	conn := newUsageServerFor(t, func() *spnpb.GetRequesterUsageResponse { return hourResponse(hour) })

	h := newUsageProvider("0x5428abf0e5aec1be48597a984a4f9570d9236f29")
	mustRefreshUsage(t, h, conn)

	hour = "2026-08-31T09:00:00Z"
	mustRefreshUsage(t, h, conn)

	if len(h.usage) != 1 || h.usage[0].NewHour {
		t.Fatalf("an older hour must not be flagged, got %+v", h.usage)
	}
}

// Requesters are tracked independently: one advancing must not suppress
// another's first sighting of the same hour.
func TestRefreshRequesterUsage_FlagsPerRequester(t *testing.T) {
	conn := newUsageServerFor(t, func() *spnpb.GetRequesterUsageResponse {
		return hourResponse("2026-08-31T10:00:00Z")
	})

	h := newUsageProvider(
		"0x5428abf0e5aec1be48597a984a4f9570d9236f29",
		"0xafb1d2c26654c85f51f550c97f16699da0293dee",
	)
	mustRefreshUsage(t, h, conn)

	if len(h.usage) != 2 {
		t.Fatalf("expected two summaries, got %d", len(h.usage))
	}
	for i, u := range h.usage {
		if !u.NewHour {
			t.Fatalf("usage[%d] (%s) was not flagged new", i, u.Requester)
		}
	}
}

// A response carrying no hour cannot be told apart from a repeat, so it must
// never be flagged: a flat counter is recoverable, an overstated one is not.
func TestRefreshRequesterUsage_NoHourIsNeverFlagged(t *testing.T) {
	conn := newUsageServerFor(t, func() *spnpb.GetRequesterUsageResponse {
		return &spnpb.GetRequesterUsageResponse{
			UsageSummary: []*spnpb.RequesterUsageSummary{{
				UsageSummary: &spnpb.UsageSummary{ReservedGas: "1", OnDemandGas: "2", TotalGas: "3"},
			}},
		}
	})

	h := newUsageProvider("0x5428abf0e5aec1be48597a984a4f9570d9236f29")
	mustRefreshUsage(t, h, conn)

	if len(h.usage) != 1 {
		t.Fatalf("expected one summary, got %d", len(h.usage))
	}
	if h.usage[0].NewHour {
		t.Fatal("a summary with no hour must not be flagged new")
	}
}

// mustRefreshUsage runs refreshRequesterUsage and fails the test on an error,
// which none of the usage tests expect.
func mustRefreshUsage(t *testing.T, h *SuccinctProverNetworkProvider, conn *grpc.ClientConn) {
	t.Helper()

	// RefreshState clears the previous cycle's results before calling this;
	// do the same so repeated calls model separate cycles.
	h.usage = nil
	if err := h.refreshRequesterUsage(context.Background(), conn); err != nil {
		t.Fatalf("refreshRequesterUsage: %v", err)
	}
}

// One requester the server answers with an error must not cost the others:
// they are still fetched, and the cycle fails with a count rather than
// silently.
func TestRefreshRequesterUsage_FailedRequesterIsCountedAndOthersFetched(t *testing.T) {
	bad := "0xacfe00ba538e753cd0a73ede2c5c27cc44a02fa8"

	desc := grpc.ServiceDesc{
		ServiceName: "network.ProverNetwork",
		HandlerType: (*any)(nil),
		Methods: []grpc.MethodDesc{{
			MethodName: "GetRequesterUsage",
			Handler: func(_ any, _ context.Context, dec func(any) error, _ grpc.UnaryServerInterceptor) (any, error) {
				req := new(spnpb.GetRequesterUsageRequest)
				if err := dec(req); err != nil {
					return nil, err
				}
				if req.Requester == bad {
					return nil, status.Error(codes.InvalidArgument, "unknown requester")
				}
				return hourResponse("2026-08-31T10:00:00Z"), nil
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

	h := newUsageProvider(bad, "0x5428abf0e5aec1be48597a984a4f9570d9236f29", "0x1111111111111111111111111111111111111111")
	err = h.refreshRequesterUsage(context.Background(), conn)
	if err == nil || !strings.Contains(err.Error(), "1 of 3 requesters") {
		t.Fatalf("expected a failure count for 1 of 3 requesters, got %v", err)
	}
	if len(h.usage) != 2 {
		t.Fatalf("expected the other two requesters' usage, got %d summaries", len(h.usage))
	}
}

// fulfilled_at is optional on the wire; a proof without it must be skipped,
// not dereferenced.
func TestPublishEvents_SkipsProofWithoutFulfilledAt(t *testing.T) {
	h := newUsageProvider()
	h.seen = make(map[string]time.Time)
	h.proofRequests = []*spnpb.ProofRequest{{RequestId: []byte("no-fulfilled-at")}}

	if err := h.PublishEvents(context.Background()); err != nil {
		t.Fatalf("PublishEvents: %v", err)
	}
	if len(h.seen) != 0 {
		t.Fatalf("a proof without fulfilled_at was published: %v", h.seen)
	}
}
