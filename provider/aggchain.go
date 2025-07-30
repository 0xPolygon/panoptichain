package provider

import (
	"context"
	"time"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/contracts"
	"github.com/0xPolygon/panoptichain/network"
	"github.com/0xPolygon/panoptichain/observer"
	"github.com/0xPolygon/panoptichain/observer/topics"
	"github.com/0xPolygon/panoptichain/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
)

type AggchainProvider struct {
	l1URL            string
	l2URL            string
	network          network.Network
	bus              *observer.EventBus
	interval         time.Duration
	label            string
	rollupAddress    common.Address
	refreshStateTime *time.Duration
	logger           zerolog.Logger
	blockNumber      uint64
	prevBlockNumber  uint64
	events           []*observer.AggchainEvent
}

func NewAggchainProvider(n network.Network, eb *observer.EventBus, cfg config.Aggchain) *AggchainProvider {
	return &AggchainProvider{
		l1URL:            cfg.L1URL,
		l2URL:            cfg.L2URL,
		network:          n,
		bus:              eb,
		interval:         GetInterval(cfg.Interval),
		label:            cfg.Label,
		rollupAddress:    common.HexToAddress(cfg.RollupAddress),
		refreshStateTime: new(time.Duration),
		logger:           NewLogger(n, cfg.Label),
	}
}

func (a *AggchainProvider) RefreshState(ctx context.Context) error {
	defer timer(a.refreshStateTime)()

	l1, err := ethclient.DialContext(ctx, a.l1URL)
	if err != nil {
		a.logger.Error().Err(err).Msg("Failed to create the L1 client")
		return err
	}

	a.prevBlockNumber = a.blockNumber
	a.blockNumber, err = l1.BlockNumber(ctx)
	if err != nil {
		a.logger.Error().Err(err).Msg("Failed to get block number")
		return err
	}

	if a.prevBlockNumber == 0 {
		return nil
	}

	contract, err := contracts.NewAggchainFEP(a.rollupAddress, l1)
	if err != nil {
		a.logger.Error().Err(err).Msg("Failed to bind AggchainFEP contract")
		return err
	}

	opts := &bind.FilterOpts{
		Start: a.prevBlockNumber,
		End:   &a.blockNumber,
	}

	iter, err := contract.FilterOutputProposed(opts, nil, nil, nil)
	if iter == nil || err != nil {
		a.logger.Error().Err(err).Msg("Failed to filter OutputProposed events")
		return err
	}

	l2, err := ethclient.DialContext(ctx, a.l2URL)
	if err != nil {
		a.logger.Error().Err(err).Msg("Failed to create the L2 client")
		return err
	}

	a.events = nil

	for iter.Next() && iter.Event != nil {
		event := &observer.AggchainEvent{
			OutputProposed: iter.Event,
		}
		a.events = append(a.events, event)

		event.L1Block, err = util.BlockByHash(ctx, iter.Event.Raw.BlockHash, l1)
		if err != nil {
			a.logger.Error().Err(err).Msg("Failed to get block by hash")
		}

		event.L2Block, err = util.BlockByNumber(ctx, iter.Event.L2BlockNumber, l2)
		if err != nil {
			a.logger.Error().Err(err).Msg("Failed to get block by number")
		}
	}

	return nil
}

func (a *AggchainProvider) PublishEvents(ctx context.Context) error {
	for _, event := range a.events {
		msg := observer.NewMessage(a.network, a.label, event)
		a.bus.Publish(ctx, topics.AggchainEvent, msg)
	}

	return nil
}

func (a *AggchainProvider) PollingInterval() time.Duration {
	return a.interval
}
