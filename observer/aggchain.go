package observer

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/0xPolygon/panoptichain/contracts"
	"github.com/0xPolygon/panoptichain/metrics"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

type AggchainEvent struct {
	OutputProposed *contracts.AggchainFEPOutputProposed
	L2Block        *types.Block
}

type AggchainObserver struct {
	latency *prometheus.HistogramVec
	bn      *prometheus.GaugeVec
	index   *prometheus.GaugeVec
}

func (o *AggchainObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.AggchainEvent, o)

	buckets := newExponentialBuckets(2, 5)
	for i := range buckets {
		buckets[i] *= time.Hour.Seconds()
	}

	o.latency = metrics.NewHistogram(
		metrics.RPC,
		"aggchain_latency",
		"The difference between the L1 timestamp and the L2 block timestamp (in seconds)",
		buckets,
	)
	o.bn = metrics.NewGauge(
		metrics.RPC,
		"aggchain_output_proposed_block_number",
		"The block number in the OutputProposed event",
	)
	o.index = metrics.NewGauge(
		metrics.RPC,
		"aggchain_output_index",
		"The output index in the OutputProposed event",
	)
}

func (o *AggchainObserver) Notify(ctx context.Context, msg Message) {
	event := msg.Data().(*AggchainEvent)

	bn, _ := event.OutputProposed.L2BlockNumber.Float64()
	o.bn.WithLabelValues(msg.Network().GetName(), msg.Provider()).Set(bn)

	index, _ := event.OutputProposed.L2OutputIndex.Float64()
	o.index.WithLabelValues(msg.Network().GetName(), msg.Provider()).Set(index)

	if event.L2Block != nil {
		latency := float64(event.OutputProposed.L1Timestamp.Uint64() - event.L2Block.Time())
		o.latency.WithLabelValues(msg.Network().GetName(), msg.Provider()).Observe(latency)
	}
}

func (o *AggchainObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.latency, o.bn, o.index}
}
