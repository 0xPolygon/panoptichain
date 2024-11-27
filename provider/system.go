package provider

import (
	"context"
	"time"

	"github.com/0xPolygon/panoptichain/observer"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

type SystemProvider struct {
	bus      *observer.EventBus
	interval uint

	start time.Time
}

func NewSystemProvider(eb *observer.EventBus, interval uint) *SystemProvider {
	return &SystemProvider{
		bus:      eb,
		interval: interval,
		start:    time.Now(),
	}
}

func (s *SystemProvider) RefreshState(context.Context) error {
	return nil
}

func (s *SystemProvider) PublishEvents(ctx context.Context) error {
	m := observer.NewMessage(nil, "", &observer.System{
		StartTime:    s.start,
		EventBusJobs: s.bus.Jobs(),
	})
	s.bus.Publish(ctx, topics.System, m)

	return nil
}

func (s *SystemProvider) SetEventBus(bus *observer.EventBus) {
	s.bus = bus
}

func (s *SystemProvider) PollingInterval() uint {
	return s.interval
}
