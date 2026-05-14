package provider

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"

	"github.com/0xPolygon/panoptichain/api"
	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/observer"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

type ExchangeRatesProvider struct {
	bus         *observer.EventBus
	interval    time.Duration
	logger      zerolog.Logger
	coinbaseURL string
	tokens      map[string][]string
	rates       []observer.ExchangeRate

	refreshStateTime *time.Duration
}

type CoinbaseExchangeRates struct {
	Data struct {
		Currency string            `json:"currency"`
		Rates    map[string]string `json:"rates"`
	} `json:"data"`
}

func NewExchangeRatesProvider(eb *observer.EventBus, cfg config.ExchangeRates) *ExchangeRatesProvider {
	return &ExchangeRatesProvider{
		bus:              eb,
		interval:         GetInterval(cfg.Interval),
		logger:           NewLogger(nil, "exchange-rates"),
		coinbaseURL:      cfg.CoinbaseURL,
		tokens:           cfg.Tokens,
		refreshStateTime: new(time.Duration),
	}
}

func (e *ExchangeRatesProvider) RefreshState(ctx context.Context) error {
	defer timer(e.refreshStateTime)()

	e.rates = nil
	for base, quotes := range e.tokens {
		e.fetchRates(base, quotes)
	}

	return nil
}

func (e *ExchangeRatesProvider) fetchRates(base string, quotes []string) {
	url := e.coinbaseURL + base

	var body CoinbaseExchangeRates
	if err := api.GetJSON(url, &body); err != nil {
		e.logger.Error().Err(err).Str("url", url).Msg("Failed to fetch exchange rates")
		return
	}

	for _, quote := range quotes {
		rate, ok := body.Data.Rates[strings.ToUpper(quote)]
		if !ok {
			e.logger.Error().Str("base", base).Str("quote", quote).Msg("Failed to get quote currency")
			continue
		}

		value, err := strconv.ParseFloat(rate, 64)
		if err != nil {
			e.logger.Error().Err(err).Str("rate", rate).Msg("Failed to parse exchange rate to float")
			continue
		}

		e.rates = append(e.rates, observer.ExchangeRate{
			Base:  strings.ToLower(base),
			Quote: strings.ToLower(quote),
			Rate:  value,
		})
	}
}

func (e *ExchangeRatesProvider) PublishEvents(ctx context.Context) error {
	e.bus.Publish(ctx, topics.RefreshStateTime, observer.NewMessage(nil, "", e.refreshStateTime))

	for _, rate := range e.rates {
		e.bus.Publish(ctx, topics.ExchangeRate, observer.NewMessage(nil, "", rate))
	}

	return nil
}

func (e *ExchangeRatesProvider) PollingInterval() time.Duration {
	return e.interval
}
