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
var observers observer.ObserverSet

// Start starts the main infinite loop of this program.
func Start(ctx context.Context) {
	log.Info().Msg("Starting main loop")

	var wg sync.WaitGroup
	wg.Add(len(providers))

	for _, p := range providers {
		go func(p provider.Provider) {
			defer wg.Done()

			for {
				if err := p.RefreshState(ctx); err != nil {
					log.Error().Err(err).Send()
				}

				if err := p.PublishEvents(ctx); err != nil {
					log.Error().Err(err).Send()
				}

				util.BlockFor(ctx, time.Second*time.Duration(p.PollingInterval()))
			}
		}(p)
	}

	wg.Wait()
}

// Init configures all the providers and observers of the system.
func Init(ctx context.Context) error {
	providers = make([]provider.Provider, 0)

	eb := observer.NewEventBus()

	var rpcProviders []*provider.RPCProvider
	for _, r := range config.Config().Providers.RPCs {
		n, err := network.GetNetworkByName(r.Name)
		if err != nil {
			return err
		}

		interval := config.Config().Runner.Interval
		if r.Interval > 0 {
			interval = r.Interval
		}

		// Look back this number of blocks when filtering event logs.
		var blockLookBack uint64 = 1000
		if r.BlockLookBack != nil {
			blockLookBack = *r.BlockLookBack
		}

		p := provider.NewRPCProvider(provider.RPCProviderOpts{
			Network:       n,
			URL:           r.URL,
			Label:         r.Label,
			EventBus:      eb,
			Interval:      interval,
			Contracts:     r.Contracts,
			TimeToMine:    r.TimeToMine,
			Accounts:      r.Accounts,
			BlockLookBack: blockLookBack,
		})

		providers = append(providers, p)
		rpcProviders = append(rpcProviders, p)
	}

	if hd := config.Config().Providers.HashDivergence; hd != nil {
		interval := config.Config().Runner.Interval
		if hd.Interval > 0 {
			interval = hd.Interval
		}

		p := provider.NewHashDivergenceProvider(rpcProviders, eb, interval)
		providers = append(providers, p)
	}

	for _, h := range config.Config().Providers.HeimdallEndpoints {
		n, err := network.GetNetworkByName(h.Name)
		if err != nil {
			return err
		}

		interval := config.Config().Runner.Interval
		if h.Interval > 0 {
			interval = h.Interval
		}

		p := provider.NewHeimdallProvider(n, h.TendermintURL, h.HeimdallURL, h.Label, eb, interval)
		providers = append(providers, p)
	}

	for _, s := range config.Config().Providers.SensorNetworks {
		n, err := network.GetNetworkByName(s.Name)
		if err != nil {
			return err
		}

		interval := config.Config().Runner.Interval
		if s.Interval > 0 {
			interval = s.Interval
		}

		p := provider.NewSensorNetworkProvider(ctx, n, s.Project, s.Database, s.Label, eb, interval)
		providers = append(providers, p)
	}

	if system := config.Config().Providers.System; system != nil {
		interval := config.Config().Runner.Interval
		if system.Interval > 0 {
			interval = system.Interval
		}

		p := provider.NewSystemProvider(eb, interval)
		providers = append(providers, p)
	}

	if er := config.Config().Providers.ExchangeRates; er != nil {
		interval := config.Config().Runner.Interval
		if er.Interval > 0 {
			interval = er.Interval
		}

		p := provider.NewExchangeRatesProvider(er.CoinbaseURL, er.Tokens, eb, interval)
		providers = append(providers, p)
	}

	observers = observer.GetEnabledObserverSet()
	observers.Register(eb)

	return nil
}
