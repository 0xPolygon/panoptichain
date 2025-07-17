package observer

import (
	"context"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/0xPolygon/panoptichain/metrics"
	"github.com/0xPolygon/panoptichain/observer/topics"
	"github.com/0xPolygon/panoptichain/proto"
)

type UsageSummary struct {
	*proto.UsageSummary
	Requester string
}

type SPNUsageSummaryObserver struct {
	total    *prometheus.GaugeVec
	reserved *prometheus.GaugeVec
	onDemand *prometheus.GaugeVec
}

func (o *SPNUsageSummaryObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.UsageSummary, o)

	o.total = metrics.NewGauge(
		metrics.SPN,
		"total_gas",
		"The total gas",
		"requester",
	)

	o.reserved = metrics.NewGauge(
		metrics.SPN,
		"reserved_gas",
		"The reserved gas",
		"requester",
	)

	o.onDemand = metrics.NewGauge(
		metrics.SPN,
		"on_demand_gas",
		"The on-demand gas",
		"requester",
	)
}

func (o *SPNUsageSummaryObserver) Notify(ctx context.Context, msg Message) {
	logger := NewLogger(o, msg)

	usageSummary := msg.Data().(*UsageSummary)

	total, err := strconv.ParseFloat(usageSummary.TotalGas, 64)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to parse total gas")
	} else {
		o.total.WithLabelValues(msg.Network().GetName(), msg.Provider(), usageSummary.Requester).Set(total)
	}

	reserved, err := strconv.ParseFloat(usageSummary.ReservedGas, 64)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to parse reserved gas")
	} else {
		o.reserved.WithLabelValues(msg.Network().GetName(), msg.Provider(), usageSummary.Requester).Set(reserved)
	}

	onDemand, err := strconv.ParseFloat(usageSummary.OnDemandGas, 64)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to parse on-demand gas")
	} else {
		o.onDemand.WithLabelValues(msg.Network().GetName(), msg.Provider(), usageSummary.Requester).Set(onDemand)
	}
}

func (o *SPNUsageSummaryObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.total, o.reserved, o.onDemand}
}
