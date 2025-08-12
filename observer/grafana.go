package observer

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/0xPolygon/panoptichain/log"
	"github.com/0xPolygon/panoptichain/metrics"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

type GrafanaResponse struct {
	Results struct {
		A struct {
			Frames []struct {
				Schema struct {
					Fields []struct {
						Labels struct {
							Fulfiller      *string `json:"fulfiller"`
							FulfillerLabel *string `json:"fulfiller_label"`
						} `json:"labels"`
					} `json:"fields"`
				} `json:"schema"`
				Data struct {
					Values [][]float64 `json:"values"`
				} `json:"data"`
			} `json:"frames"`
		} `json:"A"`
	} `json:"results"`
}

type GrafanaObserver struct {
	gauge *prometheus.GaugeVec
}

func (o *GrafanaObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.Grafana, o)

	o.gauge = metrics.NewGauge(
		metrics.SPN,
		"mpgu_per_second",
		"The gas used (million prover gas units per second)",
		"fulfiller",
	)
}

func (o *GrafanaObserver) Notify(ctx context.Context, m Message) {
	gr := m.Data().(*GrafanaResponse)

	for _, frame := range gr.Results.A.Frames {
		var fulfiller *string
		for _, field := range frame.Schema.Fields {
			if field.Labels.Fulfiller != nil {
				fulfiller = field.Labels.Fulfiller
			}
		}

		if fulfiller == nil {
			log.Warn().Msg("No fulfiller found")
			continue
		}

		ms := int64(frame.Data.Values[0][len(frame.Data.Values[0])-1])
		timestamp := time.UnixMilli(ms)

		if time.Since(timestamp) > 2*time.Minute {
			log.Warn().Time("timestamp", timestamp).Msg("Skipping stale data")
			continue
		}

		value := frame.Data.Values[1][len(frame.Data.Values[1])-1]

		log.Trace().
			Any("timestamp", timestamp).
			Any("fulfiller", fulfiller).
			Any("ts", ms).
			Any("value", value).
			Send()

		o.gauge.WithLabelValues(m.Network().GetName(), m.Provider(), *fulfiller).Set(value)
	}
}

func (o *GrafanaObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.gauge}
}
