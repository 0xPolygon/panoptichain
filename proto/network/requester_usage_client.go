package network

import (
	"context"

	"google.golang.org/grpc"
)

// ProverNetwork_GetRequesterUsage_FullMethodName is the gRPC method that
// returns per-hour gas usage for a single requester.
//
// This method is hand-written rather than generated. The upstream
// `succinctlabs/network` protos this package is generated from do not declare
// GetRequesterUsage on the ProverNetwork service, on any branch or tag, even
// though the production endpoint serves it. The request and response messages
// in requester_usage.pb.go were generated from the field numbers published in
// `succinctlabs/sp1` (crates/sdk/src/network/proto/base/types.rs), which does
// carry the full definition.
//
// Adding the method to the generated ProverNetworkClient interface would break
// every other implementer of it, so this is exposed as a plain function taking
// the connection instead. Drop it and regenerate once the service definition
// lands in the protos we compile from.
const ProverNetwork_GetRequesterUsage_FullMethodName = "/network.ProverNetwork/GetRequesterUsage"

// GetRequesterUsage returns gas usage for the given requester, bucketed by
// hour over the request's time range.
func GetRequesterUsage(ctx context.Context, cc grpc.ClientConnInterface, in *GetRequesterUsageRequest, opts ...grpc.CallOption) (*GetRequesterUsageResponse, error) {
	out := new(GetRequesterUsageResponse)
	if err := cc.Invoke(ctx, ProverNetwork_GetRequesterUsage_FullMethodName, in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}
