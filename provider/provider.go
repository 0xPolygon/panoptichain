// Package provider defines the basic provider subsystems for
// observing our networks of interest
package provider

import (
	"context"
	"time"

	"github.com/rs/zerolog"

	"github.com/0xPolygon/panoptichain/log"
	"github.com/0xPolygon/panoptichain/network"
	"github.com/0xPolygon/panoptichain/observer"
)

// Provider must be implemented by any system that's monitoring the
// state of a network.
type Provider interface {
	// Refresh state is responsible for updating the provider state. All state
	// updates should happen in this method of the provider. There should not be
	// event publishing done is this method. The Start function in runner.go will
	// call RefreshState of every provider before PublishEvents.
	RefreshState(context.Context) error

	// PublishEvents should given the current state of the provider, publish those
	// messages to the corresponding event bus. PublishEvents should not modify
	// state at all. The Start function in runner.go will call PublishEvents of
	// every provider after RefreshState.
	PublishEvents(context.Context) error

	// SetEventBus is used to configure the providers currently used message bus.
	SetEventBus(*observer.EventBus)

	// PollingInterval returns how often the provider should refresh it state and
	// publish events in seconds.
	PollingInterval() uint
}

func timer(duration *time.Duration) func() {
	start := time.Now()
	return func() {
		*duration = time.Since(start)
	}
}

func NewLogger(n network.Network, provider string) zerolog.Logger {
	network := ""
	if n != nil {
		network = n.GetName()
	}

	return log.With().
		Str("network", network).
		Str("provider", provider).
		Logger()
}
