package observer

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/0xPolygon/panoptichain/metrics"
	"github.com/0xPolygon/panoptichain/observer/topics"
	"github.com/0xPolygon/panoptichain/proto"
)

type UsageSummary struct {
	*proto.UsageSummary
	Requester string
}

type ProofRequestObserver struct {
	gas_limit   *prometheus.HistogramVec
	gas_used    *prometheus.HistogramVec
	cycle_limit *prometheus.HistogramVec
	cycles      *prometheus.HistogramVec
}

func (o *ProofRequestObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.ProofRequest, o)

	o.gas_limit = metrics.NewHistogram(
		metrics.SPN,
		"gas_limit",
		"The gas limit",
		newExponentialBuckets(10, 9),
		"requester",
		"fulfiller",
	)
	o.gas_used = metrics.NewHistogram(
		metrics.SPN,
		"gas_used",
		"The gas used",
		newExponentialBuckets(10, 9),
		"requester",
		"fulfiller",
	)
}

func (o *ProofRequestObserver) Notify(ctx context.Context, msg Message) {
	// logger := NewLogger(o, msg)

	proofRequest := msg.Data().(*proto.ProofRequest)
	labels := []string{
		msg.Network().GetName(),
		msg.Provider(),
		common.BytesToAddress(proofRequest.Requester).Hex(),
		common.BytesToAddress(proofRequest.Fulfiller).Hex(),
	}

	gl := float64(proofRequest.GasLimit)
	o.gas_limit.WithLabelValues(labels...).Observe(gl)
}

func (o *ProofRequestObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.gas_limit}
}
