package observer

import (
	"context"
	"time"

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
	time        *prometheus.HistogramVec
}

func (o *ProofRequestObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.ProofRequest, o)

	o.gas_limit = metrics.NewHistogram(
		metrics.SPN,
		"gas_limit",
		"The gas limit",
		newExponentialBuckets(10, 12),
		"requester",
		"fulfiller",
		"program",
	)
	o.gas_used = metrics.NewHistogram(
		metrics.SPN,
		"gas_used",
		"The gas used",
		newExponentialBuckets(10, 12),
		"requester",
		"fulfiller",
		"program",
	)
	o.cycle_limit = metrics.NewHistogram(
		metrics.SPN,
		"cycle_limit",
		"The cycle limit",
		newExponentialBuckets(10, 12),
		"requester",
		"fulfiller",
		"program",
	)
	o.cycles = metrics.NewHistogram(
		metrics.SPN,
		"cycles",
		"The number of cycles",
		newExponentialBuckets(10, 12),
		"requester",
		"fulfiller",
		"program",
	)
	o.time = metrics.NewHistogram(
		metrics.SPN,
		"time_to_fulfilled",
		"The time the proof took to be fulfilled",
		newExponentialBuckets(2, 12),
		"requester",
		"fulfiller",
		"program",
	)
}

func (o *ProofRequestObserver) Notify(ctx context.Context, msg Message) {
	proof := msg.Data().(*proto.ProofRequest)
	labels := []string{
		msg.Network().GetName(),
		msg.Provider(),
		common.BytesToAddress(proof.Requester).Hex(),
		common.BytesToAddress(proof.Fulfiller).Hex(),
		common.BytesToHash(proof.VkHash).Hex(),
	}

	o.gas_limit.WithLabelValues(labels...).Observe(float64(proof.GasLimit))
	o.gas_used.WithLabelValues(labels...).Observe(float64(*proof.GasUsed))
	o.cycle_limit.WithLabelValues(labels...).Observe(float64(proof.CycleLimit))
	o.cycles.WithLabelValues(labels...).Observe(float64(*proof.Cycles))

	created := time.Unix(int64(proof.CreatedAt), 0)
	fulfilled := time.Unix(int64(*proof.FulfilledAt), 0)
	dt := fulfilled.Sub(created).Seconds()
	o.time.WithLabelValues(labels...).Observe(float64(dt))
}

func (o *ProofRequestObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{
		o.gas_limit,
		o.gas_used,
		o.cycle_limit,
		o.cycles,
		o.time,
	}
}
