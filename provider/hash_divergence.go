package provider

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog"

	"github.com/0xPolygon/panoptichain/network"
	"github.com/0xPolygon/panoptichain/observer"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

// HashDivergenceProvider is a special type of provider because it's a provider
// of providers, or a meta-provider. This providers store RPCProviders and
// queries their blockbuffers to determine hash divergence.
//
// See ../runner/runner.go to see how this provider is initialized.
type HashDivergenceProvider struct {
	bus      *observer.EventBus
	interval uint
	label    string
	logger   zerolog.Logger

	// networkProvidersMap maps the network name to the provider.
	networkProvidersMap map[string][]*RPCProvider
	// networkBlockNumbers keeps track of the latest block number that was queried
	// (exclusive).
	networkBlockNumbers map[string]uint64
	hashDivergences     []*observer.CoreMessage
	refreshStateTime    *time.Duration
}

func NewHashDivergenceProvider(rpcProviders []*RPCProvider, eb *observer.EventBus, interval uint) *HashDivergenceProvider {
	label := "hash-divergence"
	networkProvidersMap := make(map[string][]*RPCProvider)
	networkBlockNumbers := make(map[string]uint64)

	for _, provider := range rpcProviders {
		networkProvidersMap[provider.Network.GetName()] = append(networkProvidersMap[provider.Network.GetName()], provider)
		networkBlockNumbers[provider.Network.GetName()] = 0
	}

	return &HashDivergenceProvider{
		bus:                 eb,
		interval:            interval,
		label:               label,
		logger:              NewLogger(nil, label),
		networkProvidersMap: networkProvidersMap,
		networkBlockNumbers: networkBlockNumbers,
		refreshStateTime:    new(time.Duration),
	}
}

func (h *HashDivergenceProvider) RefreshState(context.Context) error {
	defer timer(h.refreshStateTime)()

	h.hashDivergences = nil

loop:
	for networkName, providers := range h.networkProvidersMap {
		prevBlockNumber := h.networkBlockNumbers[networkName]
		var blockNumber uint64

		for _, provider := range providers {
			// Find the minimum BlockNumber to ensure that all providers have at least
			// that block number.
			if blockNumber == 0 || blockNumber > provider.BlockNumber {
				blockNumber = provider.BlockNumber
			}
		}

		h.logger.Debug().
			Any("prev_block_number", prevBlockNumber).
			Any("block_number", blockNumber).
			Msg("Refreshing hash divergence state")

		h.networkBlockNumbers[networkName] = blockNumber

		if prevBlockNumber == 0 || blockNumber <= prevBlockNumber {
			continue
		}

		for i := prevBlockNumber; i < blockNumber; i++ {
			var blocks []*types.Block
			hashes := make(map[common.Hash]struct{})

			for _, provider := range providers {
				b, err := provider.blockBuffer.GetBlock(i)
				if err != nil {
					continue
				}

				block, ok := b.(*types.Block)
				if !ok {
					continue
				}

				blocks = append(blocks, block)
				hashes[block.Hash()] = struct{}{}
			}

			if len(hashes) <= 1 {
				continue
			}

			n, err := network.GetNetworkByName(networkName)
			if err != nil {
				h.logger.Error().Err(err).Msgf("Could not find network: %s", networkName)
				continue loop
			}

			// Create the messages here while we still have access to the network
			// variable.
			m := observer.NewMessage(n, h.label, &observer.HashDivergence{
				Blocks:      blocks,
				BlockNumber: i,
			})
			h.hashDivergences = append(h.hashDivergences, m)
		}
	}

	return nil
}

func (h *HashDivergenceProvider) PublishEvents(ctx context.Context) error {
	for _, m := range h.hashDivergences {
		h.bus.Publish(ctx, topics.HashDivergence, m)
	}

	h.bus.Publish(ctx, topics.RefreshStateTime, observer.NewMessage(nil, h.label, h.refreshStateTime))

	return nil
}

func (h *HashDivergenceProvider) SetEventBus(bus *observer.EventBus) {
	h.bus = bus
}

func (h *HashDivergenceProvider) PollingInterval() uint {
	return h.interval
}
