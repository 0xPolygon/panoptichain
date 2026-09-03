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
	// Tag is the human-readable name configured for the requester.
	Tag string
	// Hour is the ISO-8601 hour boundary the usage falls in. It is
	// deliberately not a metric label: a new label value every hour would
	// grow the series count without bound.
	Hour string
	// NewHour reports that this bucket has not been reported before, and so
	// should advance the cumulative counters. The gauges are set either way:
	// they hold a level, and republishing the same hour is harmless. The
	// provider decides this, because it is the one that keeps state.
	NewHour bool
	// RatePerBillionGas prices the usage. Zero means no pricing is configured,
	// which suppresses the cost gauge rather than publishing a free hour.
	RatePerBillionGas float64
	// Currency denominates RatePerBillionGas.
	Currency string
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
// recent complete hour, split by how it was billed, and what that gas cost.
//
// It reports each quantity twice, because two different questions are being
// asked of it. The gauges answer "what is the burn rate right now": they hold a
// single hour's usage and are the right thing to alert a spike on. The counters
// answer "how much was consumed over some window": being cumulative, they can
// be range-summed in a query (increase(...[$__range])), which the gauges cannot
// be. A gauge holding an hourly bucket is republished on every poll, so summing
// it over time would count each hour once per poll rather than once.
//
// Which buckets are new is the provider's business, reported on the summary as
// NewHour. This observer holds no state of its own: EventBus.Publish delivers
// every message on its own goroutine, so state kept here would need its own
// synchronization, whereas a provider's cycle is single-threaded by contract.
type RequesterUsageObserver struct {
	reserved    *prometheus.GaugeVec
	on_demand   *prometheus.GaugeVec
	hourly      *prometheus.GaugeVec
	cost_hourly *prometheus.GaugeVec

	consumed   *prometheus.CounterVec
	cost_total *prometheus.CounterVec
}

func (o *RequesterUsageObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.RequesterUsage, o)

	// tag names the system that owns the requester, so a dashboard can read the
	// graph and group or exclude requesters without carrying its own address
	// table. It is drawn from config, so it adds no cardinality beyond the
	// requester label it accompanies.
	labels := []string{"requester", "tag"}

	o.reserved = metrics.NewGauge(
		metrics.SPN,
		"gas_reserved",
		"The reserved gas the requester used over the most recent complete hour",
		labels...,
	)
	o.on_demand = metrics.NewGauge(
		metrics.SPN,
		"gas_on_demand",
		"The on-demand gas the requester used over the most recent complete hour",
		labels...,
	)
	// Named hourly rather than total: a _total suffix reads as a counter to
	// anyone writing a query, and this is a gauge holding one hour.
	o.hourly = metrics.NewGauge(
		metrics.SPN,
		"gas_hourly",
		"The total gas (reserved plus on-demand) the requester used over the most recent complete hour",
		labels...,
	)
	o.cost_hourly = metrics.NewGauge(
		metrics.SPN,
		"cost_hourly",
		"The cost of the gas the requester used over the most recent complete hour, excluding any flat support fee",
		append(labels, "currency")...,
	)

	o.consumed = metrics.NewCounter(
		metrics.SPN,
		"gas_consumed_total",
		"Cumulative total gas the requester consumed, counting each hourly bucket once, for range totals",
		labels...,
	)
	o.cost_total = metrics.NewCounter(
		metrics.SPN,
		"cost_total",
		"Cumulative cost of the gas the requester consumed, excluding any flat support fee, for range totals",
		append(labels, "currency")...,
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
		usage.Tag,
	}

	// Gas values arrive as decimal strings because they can exceed uint64.
	// Skip a malformed value rather than resetting the gauge to zero, which
	// would read as "no usage" instead of "no data".
	reserved, reservedErr := strconv.ParseFloat(usage.ReservedGas, 64)
	if reservedErr == nil {
		o.reserved.WithLabelValues(labels...).Set(reserved)
	}

	onDemand, onDemandErr := strconv.ParseFloat(usage.OnDemandGas, 64)
	if onDemandErr == nil {
		o.on_demand.WithLabelValues(labels...).Set(onDemand)
	}

	total, ok := totalGas(usage, reserved, reservedErr, onDemand, onDemandErr)
	if !ok {
		return
	}

	o.hourly.WithLabelValues(labels...).Set(total)

	// A zero rate means no pricing is configured. Publishing a cost then would
	// claim the hour was free, so report gas alone and leave the cost series
	// absent for this network.
	priced := usage.RatePerBillionGas != 0
	cost := total / 1e9 * usage.RatePerBillionGas
	if priced {
		o.cost_hourly.WithLabelValues(append(labels, usage.Currency)...).Set(cost)
	}

	// Advance the counters only for a bucket the provider has not reported
	// before. Adding a repeat would overstate consumption permanently, since a
	// counter never recovers from it.
	if !usage.NewHour {
		return
	}

	o.consumed.WithLabelValues(labels...).Add(total)
	if priced {
		o.cost_total.WithLabelValues(append(labels, usage.Currency)...).Add(cost)
	}
}

// totalGas resolves the requester's total gas for the hour, preferring the
// total the network reports and falling back to the sum of the two components
// when it is absent or malformed. The fallback exists because the total is a
// convenience field: the components are what the network bills on, so a total
// that fails to parse must not cost us the cost gauge. It reports false when
// neither source is usable, which leaves both gauges at their last reading.
func totalGas(usage *UsageSummary, reserved float64, reservedErr error, onDemand float64, onDemandErr error) (float64, bool) {
	if total, err := strconv.ParseFloat(usage.TotalGas, 64); err == nil {
		return total, true
	}

	if reservedErr != nil || onDemandErr != nil {
		return 0, false
	}

	return reserved + onDemand, true
}

func (o *RequesterUsageObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{
		o.reserved,
		o.on_demand,
		o.hourly,
		o.cost_hourly,
		o.consumed,
		o.cost_total,
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
