package provider

import (
	"context"
	"time"

	zkevmtypes "github.com/0xPolygonHermez/zkevm-node/jsonrpc/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/contracts"
	"github.com/0xPolygon/panoptichain/log"
	"github.com/0xPolygon/panoptichain/observer"
)

type SequenceSenderProvider struct {
	bus              *observer.EventBus
	interval         time.Duration
	refreshStateTime *time.Duration
	config           config.SequenceSender

	blockNumber     uint64
	prevBlockNumber uint64
}

func NewSequenceSenderProvider(eb *observer.EventBus, cfg config.SequenceSender) *SequenceSenderProvider {
	if cfg.Interval == nil {
		cfg.Interval = config.Config().Runner.Interval
	}

	return &SequenceSenderProvider{
		bus:              eb,
		refreshStateTime: new(time.Duration),
		config:           cfg,
		interval:         *cfg.Interval,
	}
}

func (s *SequenceSenderProvider) RefreshState(ctx context.Context) error {
	defer timer(s.refreshStateTime)()

	l1, err := ethclient.DialContext(ctx, s.config.L1URL)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create the L1 client")
		return err
	}

	s.prevBlockNumber = s.blockNumber
	s.blockNumber, err = l1.BlockNumber(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get block number")
		return err
	}

	if s.prevBlockNumber == 0 {
		s.prevBlockNumber = s.blockNumber - 1000
	}

	address := common.HexToAddress(s.config.RollupAddress)
	contract, err := contracts.NewPolygonZkEVMEtrog(address, l1)
	if err != nil {
		log.Error().Err(err).Msg("Failed to bind zkEVM Etrog contract")
		return err
	}

	opts := &bind.FilterOpts{
		Start: s.prevBlockNumber,
		End:   &s.blockNumber,
	}

	log.Info().Any("opts", opts).Send()

	iter, err := contract.FilterSequenceBatches(opts, nil)
	if iter == nil || err != nil {
		log.Error().Err(err).Msg("Failed to filter SequenceBatches events")
		return err
	}

	var event *contracts.PolygonZkEVMEtrogSequenceBatches
	for iter.Next() && iter.Event != nil {
		event = iter.Event
		// Sometimes the send_sequence_tx_hash is nil so just grab the first one. In
		// a real scenario, wait for the tx before continuing with logic.
		break
	}

	log.Info().Any("event", event).Send()

	if event == nil {
		return nil
	}

	l2, err := ethclient.DialContext(ctx, s.config.L2URL)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create the L2 client")
		return err
	}

	var batch zkevmtypes.Batch
	err = l2.Client().CallContext(ctx, &batch, "zkevm_getBatchByNumber", event.NumBatch)
	if err != nil {
		return err
	}

	hash := batch.Blocks[len(batch.Blocks)-1].Hash
	block, err := l2.BlockByHash(ctx, *hash)

	log.Info().Any("send_sequence_tx_hash", batch.SendSequencesTxHash).Send()

	tx, _, err := l1.TransactionByHash(ctx, *batch.SendSequencesTxHash)
	if err != nil {
		return err
	}

	abi, err := contracts.PolygonZkEVMEtrogMetaData.GetAbi()
	if err != nil {
		return err
	}

	method, err := abi.MethodById(tx.Data()[:4])
	if err != nil {
		return err
	}

	inputs := make(map[string]any)
	if err := method.Inputs.UnpackIntoMap(inputs, tx.Data()[4:]); err != nil {
		return err
	}

	log.Info().
		Any("virtual_batch", event.NumBatch).
		Any("max_sequence_timestamp", inputs["maxSequenceTimestamp"]).
		Any("batch_ts", uint64(batch.Timestamp)).
		Any("block_ts", block.Time()).
		Send()

	return nil
}

func (s *SequenceSenderProvider) PublishEvents(ctx context.Context) error {
	return nil
}

func (s *SequenceSenderProvider) SetEventBus(bus *observer.EventBus) {
	s.bus = bus
}

func (ss *SequenceSenderProvider) PollingInterval() time.Duration {
	return ss.interval
}
