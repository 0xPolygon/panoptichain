package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"

	"github.com/0xPolygon/panoptichain/observer"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

type ExchangeRatesProvider struct {
	bus         *observer.EventBus
	interval    uint
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

func NewExchangeRatesProvider(coinbaseURL string, tokens map[string][]string, eb *observer.EventBus, interval uint) *ExchangeRatesProvider {
	return &ExchangeRatesProvider{
		bus:              eb,
		interval:         interval,
		logger:           NewLogger(nil, "exchange-rates"),
		coinbaseURL:      coinbaseURL,
		tokens:           tokens,
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
	r, err := http.Get(url)
	if err != nil {
		e.logger.Error().Err(err).Str("url", url).Send()
		return
	}
	defer r.Body.Close()

	var body CoinbaseExchangeRates
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get exchange rates")
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

func (e *ExchangeRatesProvider) SetEventBus(bus *observer.EventBus) {
	e.bus = bus
}

func (e *ExchangeRatesProvider) PollingInterval() uint {
	return e.interval
}
