package observer

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

type ExchangeRate struct {
	Base  string
	Quote string
	Rate  float64
}

type ExchangeRatesObserver struct {
	gauge *prometheus.GaugeVec
}

func (o *ExchangeRatesObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.ExchangeRate, o)

	o.gauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: config.Config().Namespace,
		Name:      "exchange_rates",
		Help:      "The exchange rate between the base and quote currencies",
	}, []string{"base", "quote"})
}

func (o *ExchangeRatesObserver) Notify(ctx context.Context, m Message) {
	rate := m.Data().(ExchangeRate)

	o.gauge.WithLabelValues(rate.Base, rate.Quote).Set(rate.Rate)
}

func (o *ExchangeRatesObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.gauge}
}
