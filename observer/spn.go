package observer

import (
	"context"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/0xPolygon/panoptichain/metrics"
	"github.com/0xPolygon/panoptichain/observer/topics"
	spnpb "github.com/0xPolygon/panoptichain/proto/network"
)

// UsageSummary is the gas a requester consumed over a single hour.
type UsageSummary struct {
	*spnpb.UsageSummary
	Requester string
	// Hour is the ISO-8601 hour boundary the usage falls in. It is
	// deliberately not a metric label: a new label value every hour would
	// grow the series count without bound.
	Hour string
}

type ProofRequestObserver struct {
	gas_limit             *prometheus.HistogramVec
	gas_used              *prometheus.HistogramVec
	gas_used_by_fulfiller *prometheus.HistogramVec
	cycle_limit           *prometheus.HistogramVec
	cycles                *prometheus.HistogramVec
	time                  *prometheus.HistogramVec
}

func (o *ProofRequestObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.ProofRequest, o)

	buckets := []float64{
		1_000_000,
		10_000_000,
		100_000_000,
		1_000_000_000,
		2_000_000_000,
		4_000_000_000,
		8_000_000_000,
		16_000_000_000,
		32_000_000_000,
		64_000_000_000,
		128_000_000_000,
	}

	o.gas_limit = metrics.NewHistogram(
		metrics.SPN,
		"gas_limit",
		"The gas limit",
		buckets,
		"requester",
		"fulfiller",
		"program",
	)
	o.gas_used = metrics.NewHistogram(
		metrics.SPN,
		"gas_used",
		"The gas used",
		buckets,
		"requester",
		"fulfiller",
		"program",
	)
	o.gas_used_by_fulfiller = metrics.NewHistogram(
		metrics.SPN,
		"gas_used_by_fulfiller",
		"The gas used by fulfiller",
		buckets,
		"fulfiller",
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
	proof := msg.Data().(*spnpb.ProofRequest)
	labels := []string{
		msg.Network().GetName(),
		msg.Provider(),
		common.BytesToAddress(proof.Requester).Hex(),
		common.BytesToAddress(proof.Fulfiller).Hex(),
		common.BytesToHash(proof.VkHash).Hex(),
	}

	o.gas_limit.WithLabelValues(labels...).Observe(float64(proof.GasLimit))
	o.gas_used.WithLabelValues(labels...).Observe(float64(*proof.GasUsed))
	o.gas_used_by_fulfiller.WithLabelValues(
		msg.Network().GetName(),
		msg.Provider(),
		common.BytesToAddress(proof.Fulfiller).Hex(),
	).Observe(float64(*proof.GasUsed))
	o.cycle_limit.WithLabelValues(labels...).Observe(float64(proof.CycleLimit))
	o.cycles.WithLabelValues(labels...).Observe(float64(*proof.Cycles))

	created := time.Unix(int64(proof.CreatedAt), 0)
	fulfilled := time.Unix(int64(*proof.FulfilledAt), 0)
	dt := fulfilled.Sub(created).Seconds()
	o.time.WithLabelValues(labels...).Observe(float64(dt))
}

// RequesterUsageObserver tracks the gas a requester consumed over the most
// recent complete hour, split by how it was billed.
type RequesterUsageObserver struct {
	reserved  *prometheus.GaugeVec
	on_demand *prometheus.GaugeVec
}

func (o *RequesterUsageObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.RequesterUsage, o)

	o.reserved = metrics.NewGauge(
		metrics.SPN,
		"gas_reserved",
		"The reserved gas the requester used over the most recent complete hour",
		"requester",
	)
	o.on_demand = metrics.NewGauge(
		metrics.SPN,
		"gas_on_demand",
		"The on-demand gas the requester used over the most recent complete hour",
		"requester",
	)
}

func (o *RequesterUsageObserver) Notify(ctx context.Context, msg Message) {
	usage, ok := msg.Data().(*UsageSummary)
	if !ok || usage == nil || usage.UsageSummary == nil {
		return
	}

	labels := []string{
		msg.Network().GetName(),
		msg.Provider(),
		usage.Requester,
	}

	// Gas values arrive as decimal strings because they can exceed uint64.
	// Skip a malformed value rather than resetting the gauge to zero, which
	// would read as "no usage" instead of "no data".
	if reserved, err := strconv.ParseFloat(usage.ReservedGas, 64); err == nil {
		o.reserved.WithLabelValues(labels...).Set(reserved)
	}

	if onDemand, err := strconv.ParseFloat(usage.OnDemandGas, 64); err == nil {
		o.on_demand.WithLabelValues(labels...).Set(onDemand)
	}
}

func (o *RequesterUsageObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{
		o.reserved,
		o.on_demand,
	}
}

func (o *ProofRequestObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{
		o.gas_limit,
		o.gas_used,
		o.gas_used_by_fulfiller,
		o.cycle_limit,
		o.cycles,
		o.time,
	}
}
