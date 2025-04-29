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

				util.BlockFor(ctx, p.PollingInterval())
			}
		}(p)
	}

	wg.Wait()
}

// Init configures all the providers and observers of the system.
func Init(ctx context.Context) error {
	providers = make([]provider.Provider, 0)

	eb := observer.NewEventBus()
	interval := config.Config().Runner.Interval

	var rpcProviders []*provider.RPCProvider
	for _, r := range config.Config().Providers.RPCs {
		n, err := network.GetNetworkByName(r.Name)
		if err != nil {
			return err
		}

		if r.Interval == nil {
			r.Interval = interval
		}

		// Look back this number of blocks when filtering event logs.
		blockLookBack := config.DefaultBlockLookBack
		if r.BlockLookBack != nil {
			blockLookBack = *r.BlockLookBack
		}

		p := provider.NewRPCProvider(provider.RPCProviderOpts{
			Network:       n,
			URL:           r.URL,
			Label:         r.Label,
			EventBus:      eb,
			Interval:      *r.Interval,
			Contracts:     r.Contracts,
			TimeToMine:    r.TimeToMine,
			Accounts:      r.Accounts,
			BlockLookBack: blockLookBack,
			TxPool:        r.TxPool,
		})

		providers = append(providers, p)
		rpcProviders = append(rpcProviders, p)
	}

	if hd := config.Config().Providers.HashDivergence; hd != nil {
		if hd.Interval == nil {
			hd.Interval = interval
		}

		p := provider.NewHashDivergenceProvider(rpcProviders, eb, *hd.Interval)
		providers = append(providers, p)
	}

	for _, h := range config.Config().Providers.HeimdallEndpoints {
		n, err := network.GetNetworkByName(h.Name)
		if err != nil {
			return err
		}

		if h.Interval == nil {
			h.Interval = interval
		}

		if h.Version == 0 {
			h.Version = 1
		}

		p := provider.NewHeimdallProvider(
			n,
			h.TendermintURL,
			h.HeimdallURL,
			h.Label,
			eb,
			*h.Interval,
			h.Version,
		)
		providers = append(providers, p)
	}

	for _, s := range config.Config().Providers.SensorNetworks {
		n, err := network.GetNetworkByName(s.Name)
		if err != nil {
			return err
		}

		if s.Interval == nil {
			s.Interval = interval
		}

		p := provider.NewSensorNetworkProvider(
			ctx,
			n,
			s.Project,
			s.Database,
			s.Label,
			eb,
			*interval,
		)
		providers = append(providers, p)
	}

	if system := config.Config().Providers.System; system != nil {
		if system.Interval != nil {
			system.Interval = interval
		}

		p := provider.NewSystemProvider(eb, *system.Interval)
		providers = append(providers, p)
	}

	if er := config.Config().Providers.ExchangeRates; er != nil {
		if er.Interval == nil {
			er.Interval = interval
		}

		p := provider.NewExchangeRatesProvider(
			er.CoinbaseURL,
			er.Tokens,
			eb,
			*er.Interval,
		)
		providers = append(providers, p)
	}

	observers = observer.GetEnabledObserverSet()
	observers.Register(eb)

	return nil
}
