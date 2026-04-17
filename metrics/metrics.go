// Package metrics exposes standardized functions for creating new
// counters, gauges and histograms. Ideally if we use this package
// everywhere, we can ensure consistently named metrics.
package metrics

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

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
	SPN
)

// registerOrExisting registers c and returns it, or returns the previously
// registered collector if one with the same name already exists.
func registerOrExisting[C prometheus.Collector](c C) C {
	if err := prometheus.Register(c); err != nil {
		if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
			return are.ExistingCollector.(C)
		}
		panic(err)
	}
	return c
}

// NewCounter returns a Prometheus counter object with labels for network
// and provider.
func NewCounter(subsystem Subsystem, name, help string, labels ...string) *prometheus.CounterVec {
	return registerOrExisting(prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: config.Config().Namespace,
		Subsystem: strings.ToLower(subsystem.String()),
		Name:      name,
		Help:      help,
	}, append([]string{"network", "provider"}, labels...)))
}

// NewGauge returns a Prometheus gauge with labels for network and provider.
func NewGauge(subsystem Subsystem, name, help string, labels ...string) *prometheus.GaugeVec {
	return registerOrExisting(prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: config.Config().Namespace,
		Subsystem: strings.ToLower(subsystem.String()),
		Name:      name,
		Help:      help,
	}, append([]string{"network", "provider"}, labels...)))
}

// NewGaugeWithoutLabels returns a Prometheus gauge without labels.
func NewGaugeWithoutLabels(subsystem Subsystem, name, help string) prometheus.Gauge {
	return registerOrExisting(prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: config.Config().Namespace,
		Subsystem: strings.ToLower(subsystem.String()),
		Name:      name,
		Help:      help,
	}))
}

// NewHistogram returns a configured histogram with labels for network and
// provider.
func NewHistogram(subsystem Subsystem, name, help string, buckets []float64, labels ...string) *prometheus.HistogramVec {
	return registerOrExisting(prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: config.Config().Namespace,
		Subsystem: strings.ToLower(subsystem.String()),
		Name:      name,
		Help:      help,
		Buckets:   buckets,
	}, append([]string{"network", "provider"}, labels...)))
}
