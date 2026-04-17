// Package runner is the main function for running our program
package runner

import (
	"context"
	"sync"

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
				if err := p.RefreshState(ctx); err != nil {
					log.Error().Err(err).Send()
				}

				if err := p.PublishEvents(ctx); err != nil {
					log.Error().Err(err).Send()
				}

				util.BlockFor(ctx, p.PollingInterval())
			}
		})
	}

	wg.Wait()
}

// Init configures all the providers and observers of the system.
func Init(ctx context.Context) error {
	providers = []provider.Provider{}
	rpcProviders := []*provider.RPCProvider{}

	// systemEB is used by non-RPC providers (exchange_rates, system,
	// hash_divergence, heimdall, sensor, etc.).
	systemEB := observer.NewEventBus()
	systemObs := observer.GetObserverSetFrom(config.Config().Observers.System)
	systemObs.Register(systemEB)

	// rpcFallbackEB is used by RPC providers that don't specify a custom
	// observers group.
	rpcFallbackEB := observer.NewEventBus()
	rpcFallbackObs := observer.GetObserverSetFrom(config.Config().Observers.RPC)
	rpcFallbackObs.Register(rpcFallbackEB)

	for _, r := range config.Config().Providers.RPCs {
		n, err := network.GetNetworkByName(r.Name)
		if err != nil {
			return err
		}

		eb := rpcFallbackEB
		if r.Observers != nil {
			eb = observer.NewEventBus()
			obs := observer.GetObserverSetFrom(*r.Observers)
			obs.Register(eb)
		}

		p := provider.NewRPCProvider(n, eb, r)
		providers = append(providers, p)
		rpcProviders = append(rpcProviders, p)
	}

	if hd := config.Config().Providers.HashDivergence; hd != nil {
		p := provider.NewHashDivergenceProvider(
			rpcProviders,
			systemEB,
			provider.GetInterval(hd.Interval),
		)
		providers = append(providers, p)
	}

	for _, h := range config.Config().Providers.HeimdallEndpoints {
		n, err := network.GetNetworkByName(h.Name)
		if err != nil {
			return err
		}

		p := provider.NewHeimdallProvider(n, systemEB, h)
		providers = append(providers, p)
	}

	for _, s := range config.Config().Providers.SensorNetworks {
		n, err := network.GetNetworkByName(s.Name)
		if err != nil {
			return err
		}

		p := provider.NewSensorNetworkProvider(ctx, n, systemEB, s)
		providers = append(providers, p)
	}

	for _, p := range config.Config().Providers.SuccinctProverNetworks {
		n, err := network.GetNetworkByName(p.Name)
		if err != nil {
			return err
		}

		p := provider.NewProverNetworkProvider(n, systemEB, p)
		providers = append(providers, p)
	}

	for _, p := range config.Config().Providers.Aggchains {
		n, err := network.GetNetworkByName(p.Name)
		if err != nil {
			return err
		}

		p := provider.NewAggchainProvider(n, systemEB, p)
		providers = append(providers, p)
	}

	for _, p := range config.Config().Providers.Grafana {
		n, err := network.GetNetworkByName(p.Name)
		if err != nil {
			return err
		}

		p := provider.NewGrafanaProvider(n, systemEB, p)
		providers = append(providers, p)
	}

	if system := config.Config().Providers.System; system != nil {
		p := provider.NewSystemProvider(systemEB, provider.GetInterval(system.Interval))
		providers = append(providers, p)
	}

	if er := config.Config().Providers.ExchangeRates; er != nil {
		p := provider.NewExchangeRatesProvider(systemEB, *er)
		providers = append(providers, p)
	}

	return nil
}
