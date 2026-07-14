// Package runner is the main function for running our program
package runner

import (
	"context"
	"sync"
	"time"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/log"
	"github.com/0xPolygon/panoptichain/network"
	"github.com/0xPolygon/panoptichain/observer"
	"github.com/0xPolygon/panoptichain/provider"
	"github.com/0xPolygon/panoptichain/util"
)

var providers []provider.Provider

// Start starts the main infinite loop of this program.
func Start(ctx context.Context) {
	log.Info().Msg("Starting main loop")

	var wg sync.WaitGroup

	for _, p := range providers {
		wg.Go(func() {
			for {
				runCycle(ctx, p)
			}
		})
	}

	wg.Wait()
}

// runCycle runs a single refresh/publish cycle for a provider and blocks until
// the next scheduled tick.
func runCycle(ctx context.Context, p provider.Provider) {
	interval := p.PollingInterval()
	start := time.Now()

	// Bound each cycle so a hung upstream request can't pin a provider forever.
	// This is a recovery net, not a scheduling knob: cancelling mid-cycle can
	// drop in-progress work, so keep it generous (interval*4, floored at 30s)
	// and let it fire only on real stalls. The overrun warning below is the
	// non-destructive signal for "slower than the interval". Every request path
	// honours this context (ethclient and api.GetJSON both take it).
	cycleCtx, cancel := context.WithTimeout(ctx, refreshTimeout(interval))
	if err := p.RefreshState(cycleCtx); err != nil {
		log.Error().Err(err).Send()
	}
	cancel()

	if err := p.PublishEvents(ctx); err != nil {
		log.Error().Err(err).Send()
	}

	// If a cycle runs longer than its polling interval, BlockFor will not pause,
	// so cycles run back-to-back and the provider falls behind schedule. Surface
	// it rather than letting it hide.
	if elapsed := time.Since(start); elapsed >= interval {
		log.Warn().
			Dur("elapsed", elapsed).
			Dur("interval", interval).
			Msg("Provider refresh cycle overran its polling interval")
	}

	util.BlockFor(ctx, interval)
}

// refreshTimeout is the hard ceiling on a single refresh cycle — a recovery net
// for a hung upstream, not a scheduling bound. It is deliberately generous
// (interval*4, floored at 30s) so it only trips on a genuine stall; a cycle
// merely running slower than its interval is surfaced by the overrun warning
// instead, since cancelling mid-cycle can drop in-progress work.
func refreshTimeout(interval time.Duration) time.Duration {
	const minTimeout = 30 * time.Second
	if t := interval * 4; t > minTimeout {
		return t
	}
	return minTimeout
}

// Init configures all the providers and observers of the system.
func Init(ctx context.Context) error {
	// Global EventBus for all providers without custom observers.
	eb := observer.NewEventBus()
	observers := observer.GetEnabledObserverSet()
	observers.Register(eb)

	providers = []provider.Provider{}
	rpcProviders := []*provider.RPCProvider{}

	for _, r := range config.Config().Providers.RPCs {
		n, err := network.GetNetworkByName(r.Name)
		if err != nil {
			return err
		}

		providerEB := eb // default to global
		if r.Observers != nil {
			// Provider has custom observers - create dedicated EventBus.
			providerEB = observer.NewEventBus()
			obs := observer.GetObserverSetFrom(*r.Observers)
			obs.Register(providerEB)
		}

		p := provider.NewRPCProvider(n, providerEB, r)
		providers = append(providers, p)
		rpcProviders = append(rpcProviders, p)
	}

	if hd := config.Config().Providers.HashDivergence; hd != nil {
		p := provider.NewHashDivergenceProvider(
			rpcProviders,
			eb,
			provider.GetInterval(hd.Interval),
		)
		providers = append(providers, p)
	}

	for _, h := range config.Config().Providers.HeimdallEndpoints {
		n, err := network.GetNetworkByName(h.Name)
		if err != nil {
			return err
		}

		p := provider.NewHeimdallProvider(n, eb, h)
		providers = append(providers, p)
	}

	for _, s := range config.Config().Providers.SensorNetworks {
		n, err := network.GetNetworkByName(s.Name)
		if err != nil {
			return err
		}

		p := provider.NewSensorNetworkProvider(ctx, n, eb, s)
		providers = append(providers, p)
	}

	for _, p := range config.Config().Providers.SuccinctProverNetworks {
		n, err := network.GetNetworkByName(p.Name)
		if err != nil {
			return err
		}

		p := provider.NewProverNetworkProvider(n, eb, p)
		providers = append(providers, p)
	}

	for _, p := range config.Config().Providers.Aggchains {
		n, err := network.GetNetworkByName(p.Name)
		if err != nil {
			return err
		}

		p := provider.NewAggchainProvider(n, eb, p)
		providers = append(providers, p)
	}

	for _, p := range config.Config().Providers.Grafana {
		n, err := network.GetNetworkByName(p.Name)
		if err != nil {
			return err
		}

		p := provider.NewGrafanaProvider(n, eb, p)
		providers = append(providers, p)
	}

	if system := config.Config().Providers.System; system != nil {
		p := provider.NewSystemProvider(eb, provider.GetInterval(system.Interval))
		providers = append(providers, p)
	}

	if er := config.Config().Providers.ExchangeRates; er != nil {
		p := provider.NewExchangeRatesProvider(eb, *er)
		providers = append(providers, p)
	}

	return nil
}
