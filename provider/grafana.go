package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/network"
	"github.com/0xPolygon/panoptichain/observer"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

type GrafanaProvider struct {
	network          network.Network
	label            string
	bus              *observer.EventBus
	interval         time.Duration
	url              string
	response         *observer.GrafanaResponse
	refreshStateTime *time.Duration
}

func NewGrafanaProvider(n network.Network, eb *observer.EventBus, cfg config.Grafana) *GrafanaProvider {
	return &GrafanaProvider{
		network:          n,
		label:            cfg.Label,
		bus:              eb,
		interval:         GetInterval(cfg.Interval),
		url:              cfg.URL,
		refreshStateTime: new(time.Duration),
	}
}

func (g *GrafanaProvider) RefreshState(context.Context) error {
	defer timer(g.refreshStateTime)()

	payload := []byte(`{"intervalMs":10000}`)
	req, err := http.NewRequest("POST", g.url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 256))
		return fmt.Errorf("HTTP %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}

	var gr observer.GrafanaResponse
	if err := json.NewDecoder(resp.Body).Decode(&gr); err != nil {
		return err
	}
	g.response = &gr

	return nil
}

func (g *GrafanaProvider) PublishEvents(ctx context.Context) error {
	if g.response != nil {
		m := observer.NewMessage(g.network, g.label, g.response)
		g.bus.Publish(ctx, topics.Grafana, m)
	}

	return nil
}

func (g *GrafanaProvider) PollingInterval() time.Duration {
	return g.interval
}
