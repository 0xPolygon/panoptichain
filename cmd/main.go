package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/log"
	"github.com/0xPolygon/panoptichain/runner"
)

func main() {
	ctx := context.Background()

	log.Info().Msg("Starting panoptichain")

	// There are two major components of this setup right now. We
	// have polling system to read state from various systems and
	// then we have the metrics / prometheus system to expose
	// those systems elsewhere.
	go func() {
		http.Handle(config.Config().HTTP.Path, promhttp.Handler())
		address := fmt.Sprintf("%s:%d", config.Config().HTTP.Address, config.Config().HTTP.Port)

		log.Info().
			Str("path", config.Config().HTTP.Path).
			Str("address", address).
			Msg("Starting Prometheus handler")

		if err := http.ListenAndServe(address, nil); err != nil {
			log.Error().Err(err).Msg("Failed to start server")
		}
	}()

	go func() {
		address := fmt.Sprintf("%s:%d", config.Config().HTTP.Address, 6060)
		log.Info().Str("address", address).Msg("Starting pprof")

		if err := http.ListenAndServe(address, nil); err != nil {
			log.Error().Err(err).Msg("Failed to start pprof")
		}
	}()

	if err := runner.Init(ctx); err != nil {
		log.Error().Err(err).Msg("Failed to initialize panoptichain")
		return
	}

	runner.Start(ctx)
}
