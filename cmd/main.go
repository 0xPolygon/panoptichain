package main

import (
	"context"
	"flag"
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
	flag.Parse()

	if err := config.Init(flag.Args()); err != nil {
		log.Error().Err(err).Msg("Failed to initialize config")
		return
	}

	if err := log.Init(); err != nil {
		log.Error().Err(err).Msg("Failed to initialize logger")
		return
	}

	log.Info().Msg("Starting Panoptichain")
	cfg := config.Config().HTTP

	// There are two major components of this setup right now:
	// 1. The polling system to read state from various systems.
	// 2. The metrics / Prometheus system to expose those systems elsewhere.
	go func() {
		http.Handle(cfg.Path, promhttp.Handler())
		address := fmt.Sprintf("%s:%d", cfg.Address, cfg.PromPort)
		log.Info().Str("path", cfg.Path).Str("address", address).Msg("Starting Prometheus")

		if err := http.ListenAndServe(address, nil); err != nil {
			log.Error().Err(err).Msg("Failed to start Prometheus")
		}
	}()

	go func() {
		address := fmt.Sprintf("%s:%d", cfg.Address, cfg.PprofPort)
		log.Info().Str("address", address).Msg("Starting pprof")

		if err := http.ListenAndServe(address, nil); err != nil {
			log.Error().Err(err).Msg("Failed to start pprof")
		}
	}()

	if err := runner.Init(ctx); err != nil {
		log.Error().Err(err).Msg("Failed to initialize runner")
		return
	}

	runner.Start(ctx)
}
