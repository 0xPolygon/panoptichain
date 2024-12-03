// Package metrics exposes standardized functions for creating new
// counters, gauges and histograms. Ideally if we use this package
// everywhere, we can ensure consistently named metrics.
package metrics

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/0xPolygon/panoptichain/config"
)

// Subsystem defines the different types of providers that we use to
// get data.
//
//go:generate stringer -type=Subsystem
type Subsystem int

const (
	RPC Subsystem = iota
	Sensor
	Heimdall
	System
)

// NewCounter will return a prometheus counter object with labels for network
// and provider.
func NewCounter(subsystem Subsystem, name, help string, labels ...string) *prometheus.CounterVec {
	return promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: config.Config().Namespace,
		Subsystem: strings.ToLower(subsystem.String()),
		Name:      name,
		Help:      help,
	}, append([]string{"network", "provider"}, labels...))
}

// NewGauge will return a prometheus gauge with labels for network and provider.
func NewGauge(subsystem Subsystem, name, help string, labels ...string) *prometheus.GaugeVec {
	return promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: config.Config().Namespace,
		Subsystem: strings.ToLower(subsystem.String()),
		Name:      name,
		Help:      help,
	}, append([]string{"network", "provider"}, labels...))
}

// NewGaugeWithoutLabels will return a prometheus gauge without labels.
func NewGaugeWithoutLabels(subsystem Subsystem, name, help string) prometheus.Gauge {
	return promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: config.Config().Namespace,
		Subsystem: strings.ToLower(subsystem.String()),
		Name:      name,
		Help:      help,
	})
}

// NewHistogram will return a configured histogram with labels for network and
// provider.
func NewHistogram(subsystem Subsystem, name, help string, buckets []float64, labels ...string) *prometheus.HistogramVec {
	return promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: config.Config().Namespace,
		Subsystem: strings.ToLower(subsystem.String()),
		Name:      name,
		Help:      help,
		Buckets:   buckets,
	}, append([]string{"network", "provider"}, labels...))
}
