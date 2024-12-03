package observer

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/0xPolygon/panoptichain/metrics"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

type System struct {
	StartTime    time.Time
	EventBusJobs int
}

type SystemObserver struct {
	uptime prometheus.Gauge
	jobs   prometheus.Gauge
}

func (o *SystemObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.System, o)

	o.uptime = metrics.NewGaugeWithoutLabels(
		metrics.System,
		"uptime",
		"How long panoptichain has been running in seconds",
	)

	o.jobs = metrics.NewGaugeWithoutLabels(
		metrics.System,
		"event_bus_jobs",
		"The number of goroutines being run in the event bus",
	)
}

func (o *SystemObserver) Notify(ctx context.Context, m Message) {
	system := m.Data().(*System)

	o.uptime.Set(float64(time.Since(system.StartTime).Seconds()))
	o.jobs.Set(float64(system.EventBusJobs))
}

func (o *SystemObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.uptime}
}

type RefreshStateTimeObserver struct {
	histogram *prometheus.HistogramVec
}

func (o *RefreshStateTimeObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.RefreshStateTime, o)

	o.histogram = metrics.NewHistogram(
		metrics.System,
		"refresh_state_time",
		"The amount of time it took to refresh the state in milliseconds",
		newExponentialBuckets(2, 20),
	)
}

func (o *RefreshStateTimeObserver) Notify(ctx context.Context, m Message) {
	duration := m.Data().(*time.Duration)

	network := ""
	if m.Network() != nil {
		network = m.Network().GetName()
	}

	o.histogram.WithLabelValues(network, m.Provider()).Observe(float64(duration.Milliseconds()))
}

func (o *RefreshStateTimeObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.histogram}
}
