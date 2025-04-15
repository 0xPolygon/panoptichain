// Package observer defines the event and message handing objects that
// are ultimately going to be used for metrics tracking. The observers should be fast and not connect to external data.
package observer

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	zkevmtypes "github.com/0xPolygonHermez/zkevm-node/jsonrpc/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/0xPolygon/panoptichain/api"
	"github.com/0xPolygon/panoptichain/contracts"
	"github.com/0xPolygon/panoptichain/log"
	"github.com/0xPolygon/panoptichain/metrics"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

type EmptyBlockObserver struct {
	counter *prometheus.CounterVec
}

func (o *EmptyBlockObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)

	block := m.Data().(*types.Block)

	bytes, err := api.Ecrecover(block.Header())
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to get block signer")
		return
	}
	signer := "0x" + hex.EncodeToString(bytes)

	if len(block.Transactions()) > 0 {
		return
	}

	logger.Debug().
		Uint64("block_number", block.NumberU64()).
		Str("signer", signer).
		Msg("Empty block detected")

	o.counter.WithLabelValues(m.Network().GetName(), m.Provider()).Inc()
}

func (o *EmptyBlockObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewEVMBlock, o)

	o.counter = metrics.NewCounter(
		metrics.RPC,
		"empty_block",
		"The total number of empty blocks observed",
	)
}

func (o *EmptyBlockObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.counter}
}

type BlockObserver struct {
	blockCounter *prometheus.CounterVec
	height       *prometheus.GaugeVec
	difficulty   *prometheus.GaugeVec
	blockSize    *prometheus.HistogramVec
	extraSize    *prometheus.HistogramVec
}

func (o *BlockObserver) Notify(ctx context.Context, m Message) {
	block := m.Data().(*types.Block)

	o.blockCounter.WithLabelValues(m.Network().GetName(), m.Provider()).Inc()
	o.height.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(block.NumberU64()))
	o.difficulty.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(block.Header().Difficulty.Uint64()))
	o.blockSize.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(float64(block.Size()))
	o.extraSize.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(float64(len(block.Extra())))
}

func (o *BlockObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewEVMBlock, o)

	o.blockCounter = metrics.NewCounter(metrics.RPC, "block", "The total number of blocks observed")
	o.height = metrics.NewGauge(metrics.RPC, "height", "The latest known block height")
	o.difficulty = metrics.NewGauge(metrics.RPC, "difficulty", "The difficulty of the block")
	o.blockSize = metrics.NewHistogram(metrics.RPC, "block_size", "The block size per block (bytes)", newExponentialBuckets(2, 14))
	o.extraSize = metrics.NewHistogram(metrics.RPC, "extra_size", "The size of the extra data (bytes)", newExponentialBuckets(2, 14))
}

func (o *BlockObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{
		o.height,
		o.blockCounter,
		o.difficulty,
		o.blockSize,
		o.extraSize,
	}
}

type FinalizedHeightObserver struct {
	finalizedHeight *prometheus.GaugeVec
}

func (o *FinalizedHeightObserver) Notify(ctx context.Context, m Message) {
	finalizedHeight := m.Data().(uint64)
	o.finalizedHeight.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(finalizedHeight))
}

func (o *FinalizedHeightObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.FinalizedHeight, o)

	o.finalizedHeight = metrics.NewGauge(metrics.RPC, "finalized_height", "The latest known finalized block height")
}

func (o *FinalizedHeightObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.finalizedHeight}
}

type BogonBlockObserver struct {
	counter *prometheus.CounterVec
}

func (o *BogonBlockObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)

	signers, err := api.Signers(m.Network())
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to get signers validator map")
		return
	}

	block := m.Data().(*types.Block)

	bytes, err := api.Ecrecover(block.Header())
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to get block signer")
		return
	}
	signer := "0x" + hex.EncodeToString(bytes)

	if _, ok := signers[signer]; ok {
		return
	}

	logger.Debug().
		Uint64("block_number", block.NumberU64()).
		Str("signer", signer).
		Msg("Bogon block detected")

	o.counter.WithLabelValues(m.Network().GetName(), m.Provider()).Inc()
}

func (o *BogonBlockObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewEVMBlock, o)

	o.counter = metrics.NewCounter(
		metrics.RPC,
		"bogon_block",
		"The total number of bogon blocks observed",
	)
}

func (o *BogonBlockObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.counter}
}

type StateSync struct {
	ID        uint64
	Time      time.Time // The time the state sync was first observed
	Finalized bool
}

type StateSyncObserver struct {
	stateSyncID            *prometheus.GaugeVec
	timeSinceLastStateSync *prometheus.GaugeVec
}

func (o *StateSyncObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)

	stateSync := m.Data().(*StateSync)
	seconds := time.Now().Sub(stateSync.Time).Seconds()
	finalized := fmt.Sprint(stateSync.Finalized)

	logger.Debug().
		Uint64("state_sync_id", stateSync.ID).
		Msg("State sync detected")

	o.timeSinceLastStateSync.WithLabelValues(m.Network().GetName(), m.Provider(), finalized).Set(seconds)
	o.stateSyncID.WithLabelValues(m.Network().GetName(), m.Provider(), finalized).Set(float64(stateSync.ID))
}

func (o *StateSyncObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.BorStateSync, o)

	o.stateSyncID = metrics.NewGauge(
		metrics.RPC,
		"state_sync_id",
		"the latest observed state sync id",
		"finalized",
	)
	o.timeSinceLastStateSync = metrics.NewGauge(
		metrics.RPC,
		"time_since_last_state_sync",
		"The elapsed time since the last state sync",
		"finalized",
	)
}

func (o *StateSyncObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.stateSyncID, o.timeSinceLastStateSync}
}

type BlockIntervalObserver struct {
	blockInterval *prometheus.HistogramVec
}

func (o *BlockIntervalObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)

	interval := m.Data().(uint64)

	logger.Trace().
		Uint64("interval", interval).
		Msg("Block interval")

	o.blockInterval.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(float64(interval))
}

func (o *BlockIntervalObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.BlockInterval, o)

	o.blockInterval = metrics.NewHistogram(
		metrics.RPC,
		"block_interval",
		"the number of seconds between blocks",
		newExponentialBuckets(2, 6),
	)
}

func (o *BlockIntervalObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.blockInterval}
}

type TransactionCountObserver struct {
	histogram *prometheus.HistogramVec
}

func (o *TransactionCountObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewEVMBlock, o)

	o.histogram = metrics.NewHistogram(
		metrics.RPC,
		"transactions_per_block",
		"The number of transactions per block",
		newExponentialBuckets(2, 11),
	)
}

func (o *TransactionCountObserver) Notify(ctx context.Context, m Message) {
	block := m.Data().(*types.Block)
	txs := float64(len(block.Transactions()))
	o.histogram.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(txs)
}

func (o *TransactionCountObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.histogram}
}

type BaseFeePerGasObserver struct {
	gauge *prometheus.GaugeVec
}

func (o *BaseFeePerGasObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewEVMBlock, o)
	// TODO(praetoriasentry): is this worth having a histogram at some point?
	o.gauge = metrics.NewGauge(metrics.RPC, "base_fee_per_gas", "The base fee per gas (gwei)")
}

func (o *BaseFeePerGasObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)

	block := m.Data().(*types.Block)

	if block.BaseFee() == nil {
		logger.Warn().Msg("Base fee is nil")
		return
	}

	gwei, _ := weiToGwei(block.BaseFee()).Float64()
	o.gauge.WithLabelValues(m.Network().GetName(), m.Provider()).Set(gwei)
}

func (o *BaseFeePerGasObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.gauge}
}

type GasLimitObserver struct {
	gauge *prometheus.GaugeVec
}

func (o *GasLimitObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewEVMBlock, o)

	o.gauge = metrics.NewGauge(
		metrics.RPC,
		"gas_limit",
		"The gas limit of the block",
	)
}

func (o *GasLimitObserver) Notify(ctx context.Context, m Message) {
	block := m.Data().(*types.Block)
	gasLimit := float64(block.GasLimit())
	o.gauge.WithLabelValues(m.Network().GetName(), m.Provider()).Set(gasLimit)
}

func (o *GasLimitObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.gauge}
}

type GasUsedObserver struct {
	histogram *prometheus.HistogramVec
}

func (o *GasUsedObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewEVMBlock, o)

	o.histogram = metrics.NewHistogram(
		metrics.RPC,
		"gas_used",
		"The gas used in the block (million gas)",
		newExponentialBuckets(2, 6),
	)
}

func (o *GasUsedObserver) Notify(ctx context.Context, m Message) {
	block := m.Data().(*types.Block)
	gasUsed := float64(block.GasUsed()) / 1_000_000
	o.histogram.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(gasUsed)
}

func (o *GasUsedObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.histogram}
}

type TransactionCostObserver struct {
	histogram *prometheus.HistogramVec
}

func (o *TransactionCostObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewEVMBlock, o)

	o.histogram = metrics.NewHistogram(
		metrics.RPC,
		"transaction_cost",
		"The transaction cost (ether)",
		newExponentialBuckets(2, 11),
	)
}

func (o *TransactionCostObserver) Notify(ctx context.Context, m Message) {
	block := m.Data().(*types.Block)

	for _, tx := range block.Transactions() {
		ether, _ := weiToEther(tx.Cost()).Float64()
		o.histogram.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(ether)
	}
}

func (o *TransactionCostObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.histogram}
}

type TransactionGasLimitObserver struct {
	histogram *prometheus.HistogramVec
}

func (o *TransactionGasLimitObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewEVMBlock, o)
	o.histogram = metrics.NewHistogram(
		metrics.RPC,
		"transaction_gas_limit",
		"The transaction gas limit (gas)",
		newExponentialBuckets(2, 11),
	)
}

func (o *TransactionGasLimitObserver) Notify(ctx context.Context, m Message) {
	block := m.Data().(*types.Block)

	for _, tx := range block.Transactions() {
		gas := float64(tx.Gas()) / 100_000
		o.histogram.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(gas)
	}
}

func (o *TransactionGasLimitObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.histogram}
}

type TransactionGasPriceObserver struct {
	histogram *prometheus.HistogramVec
}

func (o *TransactionGasPriceObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewEVMBlock, o)

	o.histogram = metrics.NewHistogram(
		metrics.RPC,
		"transaction_gas_price",
		"The transaction gas price (gwei)",
		newExponentialBuckets(2, 12),
	)
}

func (o *TransactionGasPriceObserver) Notify(ctx context.Context, m Message) {
	block := m.Data().(*types.Block)

	for _, tx := range block.Transactions() {
		gwei, _ := weiToGwei(tx.GasPrice()).Float64()
		o.histogram.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(gwei)
	}
}

func (o *TransactionGasPriceObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.histogram}
}

type TransactionGasFeeCapObserver struct {
	histogram *prometheus.HistogramVec
}

func (o *TransactionGasFeeCapObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewEVMBlock, o)

	o.histogram = metrics.NewHistogram(
		metrics.RPC,
		"transaction_gas_fee_cap",
		"The transaction gas fee cap (gwei)",
		newExponentialBuckets(2, 12),
	)
}

func (o *TransactionGasFeeCapObserver) Notify(ctx context.Context, m Message) {
	block := m.Data().(*types.Block)

	for _, tx := range block.Transactions() {
		gwei, _ := weiToGwei(tx.GasFeeCap()).Float64()
		o.histogram.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(gwei)
	}
}

func (o *TransactionGasFeeCapObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.histogram}
}

type TransactionGasTipCapObserver struct {
	histogram *prometheus.HistogramVec
}

func (o *TransactionGasTipCapObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewEVMBlock, o)
	o.histogram = metrics.NewHistogram(
		metrics.RPC,
		"transaction_gas_tip_cap",
		"The transaction gas tip cap (gwei)",
		newExponentialBuckets(2, 12),
	)
}

func (o *TransactionGasTipCapObserver) Notify(ctx context.Context, m Message) {
	block := m.Data().(*types.Block)

	for _, tx := range block.Transactions() {
		gwei, _ := weiToGwei(tx.GasTipCap()).Float64()
		o.histogram.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(gwei)
	}
}

func (o *TransactionGasTipCapObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.histogram}
}

type TransactionValueObserver struct {
	histogram *prometheus.HistogramVec
}

func (o *TransactionValueObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewEVMBlock, o)

	o.histogram = metrics.NewHistogram(
		metrics.RPC,
		"transaction_value",
		"The value of the transactions (ether)",
		newExponentialBuckets(2, 10),
	)
}

func (o *TransactionValueObserver) Notify(ctx context.Context, m Message) {
	block := m.Data().(*types.Block)

	for _, tx := range block.Transactions() {
		ether, _ := weiToEther(tx.Value()).Float64()
		o.histogram.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(ether)
	}
}

func (o *TransactionValueObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.histogram}
}

type UnclesObserver struct {
	counter *prometheus.CounterVec
}

func (o *UnclesObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.NewEVMBlock, o)
	o.counter = metrics.NewCounter(
		metrics.RPC,
		"uncles",
		"The number of uncles for the block",
	)
}

func (o *UnclesObserver) Notify(ctx context.Context, m Message) {
	block := m.Data().(*types.Block)
	uncles := block.Uncles()
	o.counter.WithLabelValues(m.Network().GetName(), m.Provider()).Add(float64(len(uncles)))
}

func (o *UnclesObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.counter}
}

func weiToGwei(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.GWei))
}

func weiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
}

type CheckpointSignatures struct {
	Event     *contracts.RootChainNewHeaderBlock
	Block     *types.Block
	Signers   []common.Address
	Seen      bool
	Finalized bool
}

type CheckpointObserver struct {
	checkpointID            *prometheus.GaugeVec
	checkpointSignatures    *prometheus.GaugeVec
	timeSinceLastCheckpoint *prometheus.GaugeVec
	signedCheckpoint        *prometheus.CounterVec
}

func (o *CheckpointObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.CheckpointSignatures, o)
	o.checkpointID = metrics.NewGauge(
		metrics.RPC,
		"checkpoint_id",
		"The last header block ID",
		"finalized",
	)
	o.checkpointSignatures = metrics.NewGauge(
		metrics.RPC,
		"checkpoint_signatures",
		"The number of validators that signed the latest checkpoint",
		"finalized",
	)
	o.timeSinceLastCheckpoint = metrics.NewGauge(
		metrics.RPC,
		"time_since_last_checkpoint",
		"The elapsed time since the last checkpoint",
		"finalized",
	)
	o.signedCheckpoint = metrics.NewCounter(
		metrics.RPC,
		"signed_checkpoint",
		"Counts the number of times a validator has signed a checkpoint",
		"signer",
	)
}

func (o *CheckpointObserver) Notify(ctx context.Context, m Message) {
	cs := m.Data().(*CheckpointSignatures)
	finalized := fmt.Sprint(cs.Finalized)

	id, _ := cs.Event.HeaderBlockId.Float64()
	checkpointTime := time.Unix(int64(cs.Block.Time()), 0)
	seconds := m.Time().Sub(checkpointTime).Seconds()

	o.checkpointID.WithLabelValues(m.Network().GetName(), m.Provider(), finalized).Set(id)
	o.checkpointSignatures.WithLabelValues(m.Network().GetName(), m.Provider(), finalized).Set(float64(len(cs.Signers)))
	o.timeSinceLastCheckpoint.WithLabelValues(m.Network().GetName(), m.Provider(), finalized).Set(seconds)

	// We only update the following metrics if there is a new checkpoint detected.
	if cs.Seen || cs.Finalized {
		return
	}

	for _, signer := range cs.Signers {
		o.signedCheckpoint.WithLabelValues(m.Network().GetName(), m.Provider(), signer.Hex()).Inc()
	}
}

func (o *CheckpointObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{
		o.checkpointID,
		o.checkpointSignatures,
		o.timeSinceLastCheckpoint,
		o.signedCheckpoint,
	}
}

type ValidatorWalletBalances map[string]*big.Int

type ValidatorWalletBalanceObserver struct {
	balances *prometheus.GaugeVec
}

func (o *ValidatorWalletBalanceObserver) Notify(ctx context.Context, m Message) {
	balances := m.Data().(ValidatorWalletBalances)

	for address, balance := range balances {
		wei := float64(balance.Uint64())
		o.balances.WithLabelValues(m.Network().GetName(), m.Provider(), address).Set(wei)
	}
}

func (o *ValidatorWalletBalanceObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.ValidatorWallet, o)
	o.balances = metrics.NewGauge(
		metrics.RPC,
		"validator_wallet_balance",
		"PoS validator wallet balance",
		"signer_address",
	)
}

func (o *ValidatorWalletBalanceObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.balances}
}

type MissedBlockProposal map[uint64][]string

type MissedBlockProposalObserver struct {
	counter *prometheus.CounterVec
}

func (o *MissedBlockProposalObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)

	data := m.Data().(MissedBlockProposal)
	for blockNumber, proposers := range data {
		logger.Debug().
			Uint64("block_number", blockNumber).
			Strs("proposers", proposers).
			Msg("Missed block proposer update")

		for _, proposer := range proposers {
			o.counter.WithLabelValues(m.Network().GetName(), m.Provider(), proposer).Inc()
		}
	}
}

func (o *MissedBlockProposalObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.BorMissedBlockProposal, o)

	o.counter = metrics.NewCounter(
		metrics.RPC,
		"missed_block_proposal",
		"Missed block proposals",
		"signer_address",
	)
}

func (o *MissedBlockProposalObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.counter}
}

type TransactionPool struct {
	Pending uint64
	Queued  uint64
}

type TransactionPoolObserver struct {
	pending *prometheus.GaugeVec
	queued  *prometheus.GaugeVec
}

func (o *TransactionPoolObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)

	txPool := m.Data().(*TransactionPool)

	logger.Debug().
		Uint64("pending", txPool.Pending).
		Uint64("queued", txPool.Queued).
		Msg("Transaction pool status")

	o.pending.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(txPool.Pending))
	o.queued.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(txPool.Queued))
}

func (o *TransactionPoolObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.TransactionPool, o)

	o.pending = metrics.NewGauge(metrics.RPC, "pending_tx_size", "Number of pending transactions")
	o.queued = metrics.NewGauge(metrics.RPC, "queued_tx_size", "Number of queued transactions")
}

func (o *TransactionPoolObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.pending, o.queued}
}

type HashDivergence struct {
	Blocks      []*types.Block
	BlockNumber uint64
}

type HashDivergenceObserver struct {
	counter *prometheus.CounterVec
}

func (o *HashDivergenceObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)

	hashDivergence := m.Data().(*HashDivergence)

	var hashes []common.Hash
	for _, block := range hashDivergence.Blocks {
		hashes = append(hashes, block.Hash())
	}

	logger.Debug().
		Uint64("block_number", hashDivergence.BlockNumber).
		Any("hashes", hashes).
		Msg("Hash divergence detected")

	o.counter.WithLabelValues(m.Network().GetName(), m.Provider()).Inc()
}

func (o *HashDivergenceObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.HashDivergence, o)

	o.counter = metrics.NewCounter(
		metrics.RPC,
		"hash_divergence",
		"The number of blocks that have different hashes across different RPC providers",
	)
}

func (o *HashDivergenceObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.counter}
}

type ZkEVMBatches struct {
	TrustedBatch  ZkEVMBatch
	VirtualBatch  ZkEVMBatch
	VerifiedBatch ZkEVMBatch
}

type ZkEVMBatch struct {
	Number uint64
	Time   time.Time
}

type ZkEVMBatchObserver struct {
	trustedBatch  *prometheus.GaugeVec
	virtualBatch  *prometheus.GaugeVec
	verifiedBatch *prometheus.GaugeVec

	timeSinceLastTrustedBatch  *prometheus.GaugeVec
	timeSinceLastVirtualBatch  *prometheus.GaugeVec
	timeSinceLastVerifiedBatch *prometheus.GaugeVec
}

func (o *ZkEVMBatchObserver) Notify(ctx context.Context, m Message) {
	batches := m.Data().(ZkEVMBatches)

	o.trustedBatch.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(batches.TrustedBatch.Number))
	o.virtualBatch.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(batches.VirtualBatch.Number))
	o.verifiedBatch.WithLabelValues(m.Network().GetName(), m.Provider()).Set(float64(batches.VerifiedBatch.Number))

	tbt := float64(time.Since(batches.TrustedBatch.Time).Seconds())
	o.timeSinceLastTrustedBatch.WithLabelValues(m.Network().GetName(), m.Provider()).Set(tbt)

	vibt := float64(time.Since(batches.VirtualBatch.Time).Seconds())
	o.timeSinceLastVirtualBatch.WithLabelValues(m.Network().GetName(), m.Provider()).Set(vibt)

	vebt := float64(time.Since(batches.VerifiedBatch.Time).Seconds())
	o.timeSinceLastVerifiedBatch.WithLabelValues(m.Network().GetName(), m.Provider()).Set(vebt)
}

func (o *ZkEVMBatchObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.ZkEVMBatches, o)

	o.trustedBatch = metrics.NewGauge(metrics.RPC, "trusted_batch", "zkEVM trusted batch number")
	o.virtualBatch = metrics.NewGauge(metrics.RPC, "virtual_batch", "zkEVM virtual batch number")
	o.verifiedBatch = metrics.NewGauge(metrics.RPC, "verified_batch", "zkEVM verified batch number")

	o.timeSinceLastTrustedBatch = metrics.NewGauge(metrics.RPC, "time_since_last_trusted_batch", "time since last zkEVM trusted batch (in seconds)")
	o.timeSinceLastVirtualBatch = metrics.NewGauge(metrics.RPC, "time_since_last_virtual_batch", "time since last zkEVM virtual batch (in seconds)")
	o.timeSinceLastVerifiedBatch = metrics.NewGauge(metrics.RPC, "time_since_last_verified_batch", "time since last zkEVM verified batch (in seconds)")
}

func (o *ZkEVMBatchObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{
		o.trustedBatch,
		o.virtualBatch,
		o.verifiedBatch,
		o.timeSinceLastTrustedBatch,
		o.timeSinceLastVirtualBatch,
		o.timeSinceLastVerifiedBatch,
	}
}

type ExitRoot struct {
	Hash common.Hash
	Time time.Time
	Seen bool
}

type ExitRoots struct {
	GlobalExitRoot  *ExitRoot
	MainnetExitRoot *ExitRoot
	RollupExitRoot  *ExitRoot
}

type ExitRootsObserver struct {
	timeSinceLastGlobalExitRoot  *prometheus.GaugeVec
	timeSinceLastMainnetExitRoot *prometheus.GaugeVec
	timeSinceLastRollupExitRoot  *prometheus.GaugeVec

	globalExitRoots  *prometheus.CounterVec
	mainnetExitRoots *prometheus.CounterVec
	rollupExitRoots  *prometheus.CounterVec
}

func (o *ExitRootsObserver) Notify(ctx context.Context, m Message) {
	er := m.Data().(*ExitRoots)

	if er.GlobalExitRoot != nil {
		seconds := float64(time.Since(er.GlobalExitRoot.Time).Seconds())
		o.timeSinceLastGlobalExitRoot.WithLabelValues(m.Network().GetName(), m.Provider()).Set(seconds)

		if !er.GlobalExitRoot.Seen {
			o.globalExitRoots.WithLabelValues(m.Network().GetName(), m.Provider()).Inc()
		}
	}

	if er.MainnetExitRoot != nil {
		seconds := float64(time.Since(er.MainnetExitRoot.Time).Seconds())
		o.timeSinceLastMainnetExitRoot.WithLabelValues(m.Network().GetName(), m.Provider()).Set(seconds)

		if !er.MainnetExitRoot.Seen {
			o.mainnetExitRoots.WithLabelValues(m.Network().GetName(), m.Provider()).Inc()
		}
	}

	if er.RollupExitRoot != nil {
		seconds := float64(time.Since(er.RollupExitRoot.Time).Seconds())
		o.timeSinceLastRollupExitRoot.WithLabelValues(m.Network().GetName(), m.Provider()).Set(seconds)

		if !er.RollupExitRoot.Seen {
			o.rollupExitRoots.WithLabelValues(m.Network().GetName(), m.Provider()).Inc()
		}
	}
}

func (o *ExitRootsObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.ExitRoots, o)

	o.timeSinceLastGlobalExitRoot = metrics.NewGauge(
		metrics.RPC,
		"time_since_last_global_exit_root",
		"The elapsed time since the last global exit root (in seconds)",
	)
	o.timeSinceLastMainnetExitRoot = metrics.NewGauge(
		metrics.RPC,
		"time_since_last_mainnet_exit_root",
		"The elapsed time since the last mainnet exit root (in seconds)",
	)
	o.timeSinceLastRollupExitRoot = metrics.NewGauge(
		metrics.RPC,
		"time_since_last_rollup_exit_root",
		"The elapsed time since the last rollup exit root (in seconds)",
	)
	o.globalExitRoots = metrics.NewCounter(
		metrics.RPC,
		"global_exit_roots",
		"The number of unique global exit roots that have been observed",
	)
	o.mainnetExitRoots = metrics.NewCounter(
		metrics.RPC,
		"mainnet_exit_roots",
		"The number of unique mainnet exit roots that have been observed",
	)
	o.rollupExitRoots = metrics.NewCounter(
		metrics.RPC,
		"rollup_exit_roots",
		"The number of unique rollup exit roots that have been observed",
	)
}

func (o *ExitRootsObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{
		o.timeSinceLastGlobalExitRoot,
		o.timeSinceLastMainnetExitRoot,
		o.timeSinceLastRollupExitRoot,
	}
}

type DepositCounts struct {
	DepositCount            *big.Int
	LastUpdatedDepositCount *uint32
}

type DepositCountObserver struct {
	depositCount            *prometheus.GaugeVec
	lastUpdatedDepositCount *prometheus.GaugeVec
}

func (o *DepositCountObserver) Notify(ctx context.Context, m Message) {
	data := m.Data().(*DepositCounts)

	if data.DepositCount != nil {
		dc, _ := data.DepositCount.Float64()
		o.depositCount.WithLabelValues(m.Network().GetName(), m.Provider()).Set(dc)
	}

	if data.LastUpdatedDepositCount != nil {
		ludc := float64(*data.LastUpdatedDepositCount)
		o.lastUpdatedDepositCount.WithLabelValues(m.Network().GetName(), m.Provider()).Set(ludc)
	}
}

func (o *DepositCountObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.DepositCounts, o)

	o.depositCount = metrics.NewGauge(
		metrics.RPC,
		"bridge_deposit_count",
		"zkEVM bridge deposit count",
	)
	o.lastUpdatedDepositCount = metrics.NewGauge(
		metrics.RPC,
		"bridge_last_updated_deposit_count",
		"zkEVM bridge last updated deposit count",
	)
}

func (o *DepositCountObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.depositCount, o.lastUpdatedDepositCount}
}

type BridgeEventNetworks struct {
	OriginNetwork      uint32
	DestinationNetwork uint32
}

type BridgeEventTimes map[BridgeEventNetworks]time.Time

type BridgeEventObserver struct {
	timeSinceLastBridgeEvent *prometheus.GaugeVec
	depositCount             *prometheus.GaugeVec
	amount                   *prometheus.HistogramVec
}

func (o *BridgeEventObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)

	switch v := m.Data().(type) {
	case *contracts.PolygonZkEVMBridgeV2BridgeEvent:
		origin := fmt.Sprint(v.OriginNetwork)
		destination := fmt.Sprint(v.DestinationNetwork)

		dc := float64(v.DepositCount)
		o.depositCount.WithLabelValues(m.Network().GetName(), m.Provider(), origin, destination).Set(dc)

		gwei, _ := weiToGwei(v.Amount).Float64()
		o.amount.WithLabelValues(m.Network().GetName(), m.Provider(), origin, destination).Observe(gwei)

	case BridgeEventTimes:
		for k, t := range v {
			origin := fmt.Sprint(k.OriginNetwork)
			destination := fmt.Sprint(k.DestinationNetwork)
			seconds := time.Since(t).Seconds()
			o.timeSinceLastBridgeEvent.WithLabelValues(m.Network().GetName(), m.Provider(), origin, destination).Set(seconds)
		}

	default:
		logger.Error().Msg("Failed to match any types")
	}
}

func (o *BridgeEventObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.BridgeEvent, o)
	eb.Subscribe(topics.BridgeEventTimes, o)

	o.timeSinceLastBridgeEvent = metrics.NewGauge(
		metrics.RPC,
		"time_since_last_bridge_event",
		"The time since the last zkEVM bridge event (in seconds)",
		"origin_network",
		"destination_network",
	)
	o.depositCount = metrics.NewGauge(
		metrics.RPC,
		"bridge_event_deposit_count",
		"The deposit count of the latest bridge event",
		"origin_network",
		"destination_network",
	)
	o.amount = metrics.NewHistogram(
		metrics.RPC,
		"bridge_event_amount",
		"The amount in bridged (gwei)",
		newExponentialBuckets(10, 9),
		"origin_network",
		"destination_network",
	)
}

func (o *BridgeEventObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.timeSinceLastBridgeEvent, o.depositCount, o.amount}
}

type ClaimEventTimes map[uint32]time.Time

type ClaimEventObserver struct {
	timeSinceLastClaimEvent *prometheus.GaugeVec
	amount                  *prometheus.HistogramVec
	observedClaimEvents     *prometheus.CounterVec
}

func (o *ClaimEventObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)

	switch v := m.Data().(type) {
	case *contracts.PolygonZkEVMBridgeV2ClaimEvent:
		origin := fmt.Sprint(v.OriginNetwork)
		gwei, _ := weiToGwei(v.Amount).Float64()
		o.amount.WithLabelValues(m.Network().GetName(), m.Provider(), origin).Observe(gwei)
		o.observedClaimEvents.WithLabelValues(m.Network().GetName(), m.Provider(), origin).Inc()

	case ClaimEventTimes:
		for k, t := range v {
			origin := fmt.Sprint(k)
			seconds := time.Since(t).Seconds()
			o.timeSinceLastClaimEvent.WithLabelValues(m.Network().GetName(), m.Provider(), origin).Set(seconds)
		}

	default:
		logger.Error().Msg("Failed to match any types")
	}
}

func (o *ClaimEventObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.ClaimEvent, o)
	eb.Subscribe(topics.ClaimEventTimes, o)

	o.timeSinceLastClaimEvent = metrics.NewGauge(
		metrics.RPC,
		"time_since_last_claim_event",
		"The time since the last zkEVM claim event (in seconds)",
		"origin_network",
	)
	o.amount = metrics.NewHistogram(
		metrics.RPC,
		"claim_event_amount",
		"The amount in claimed (gwei)",
		newExponentialBuckets(10, 9),
		"origin_network",
	)
	o.observedClaimEvents = metrics.NewCounter(
		metrics.RPC,
		"observed_claim_events",
		"The number of claim events observed",
		"origin_network",
	)
}

func (o *ClaimEventObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.timeSinceLastClaimEvent, o.amount, o.observedClaimEvents}
}

const (
	ETH = "eth"
	POL = "pol"
)

type TokenBalances struct {
	ETH *big.Int
	POL *big.Int
}

type RollupManager struct {
	TotalSequencedBatches *uint64
	TotalVerifiedBatches  *uint64

	BatchFee       *big.Int
	ForcedBatchFee *big.Int
	RewardPerBatch *big.Int

	AggregatorBalances       map[common.Address]TokenBalances
	LastAggregationTimestamp *uint64

	Rollups         map[uint32]*RollupData
	RollupCount     *uint32
	RollupTypeCount *uint32
}

type RollupTx struct {
	Fee     *big.Int
	Address common.Address
}

type RollupData struct {
	LastBatchSequenced          *uint64
	LastSequencedTimestamp      *uint64
	TimeBetweenSequencedBatches []uint64
	SequencedBatchesTxFees      []RollupTx

	LastForceBatch          *uint64
	LastForceBatchSequenced *uint64

	LastVerifiedBatch          *uint64
	LastVerifiedTimestamp      *uint64
	TimeBetweenVerifiedBatches []uint64
	VerifiedBatchesTxFees      []RollupTx

	TrustedSequencerBalances TokenBalances

	ChainID     *uint64
	Pessimistic bool
}

type RollupManagerObserver struct {
	lastBatchSequenced          *prometheus.GaugeVec
	timeSinceLastSequenced      *prometheus.GaugeVec
	totalSequencedBatches       *prometheus.GaugeVec
	timeBetweenSequencedBatches *prometheus.HistogramVec
	sequencedBatchesTxFee       *prometheus.HistogramVec
	observedSequencedBatches    *prometheus.CounterVec

	lastVerifiedBatch          *prometheus.GaugeVec
	timeSinceLastVerified      *prometheus.GaugeVec
	totalVerifiedBatches       *prometheus.GaugeVec
	timeBetweenVerifiedBatches *prometheus.HistogramVec
	verifiedBatchesTxFee       *prometheus.HistogramVec
	observedVerifiedBatches    *prometheus.CounterVec

	lastForceBatch          *prometheus.GaugeVec
	lastForceBatchSequenced *prometheus.GaugeVec

	chainID *prometheus.GaugeVec

	batchFee       *prometheus.GaugeVec
	rewardPerBatch *prometheus.GaugeVec

	trustedSequencerBalance  *prometheus.GaugeVec
	aggregatorBalance        *prometheus.GaugeVec
	timeSinceLastAggregation *prometheus.GaugeVec

	rollupCount     *prometheus.GaugeVec
	rollupTypeCount *prometheus.GaugeVec
}

func (o *RollupManagerObserver) Notify(ctx context.Context, m Message) {
	data := m.Data().(*RollupManager)

	if data.BatchFee != nil {
		bf, _ := weiToGwei(data.BatchFee).Float64()
		log.Info().Float64("batch_fee", bf).Send()
		o.batchFee.WithLabelValues(m.Network().GetName(), m.Provider()).Set(bf)
	}

	if data.RewardPerBatch != nil {
		rpb, _ := weiToGwei(data.RewardPerBatch).Float64()
		log.Info().Float64("reward_per_batch", rpb).Send()
		o.rewardPerBatch.WithLabelValues(m.Network().GetName(), m.Provider()).Set(rpb)
	}

	if data.RollupCount != nil {
		rc := float64(*data.RollupCount)
		o.rollupCount.WithLabelValues(m.Network().GetName(), m.Provider()).Set(rc)
	}

	if data.RollupTypeCount != nil {
		rtc := float64(*data.RollupTypeCount)
		o.rollupTypeCount.WithLabelValues(m.Network().GetName(), m.Provider()).Set(rtc)
	}

	if data.LastAggregationTimestamp != nil {
		seconds := time.Since(time.Unix(int64(*data.LastAggregationTimestamp), 0)).Seconds()
		o.timeSinceLastAggregation.WithLabelValues(m.Network().GetName(), m.Provider()).Set(seconds)
	}

	for id, rollup := range data.Rollups {
		o.notifyRollup(m, rollup, fmt.Sprint(id))
	}

	if data.TotalSequencedBatches != nil {
		tsb := float64(*data.TotalSequencedBatches)
		o.totalSequencedBatches.WithLabelValues(m.Network().GetName(), m.Provider()).Set(tsb)
	}

	if data.TotalVerifiedBatches != nil {
		tvb := float64(*data.TotalVerifiedBatches)
		o.totalVerifiedBatches.WithLabelValues(m.Network().GetName(), m.Provider()).Set(tvb)
	}

	if len(data.AggregatorBalances) > 0 {
		o.notifyAggregatorBalances(m, data.AggregatorBalances)
	}
}

func (o *RollupManagerObserver) notifyAggregatorBalances(m Message, balances map[common.Address]TokenBalances) {
	for address, balances := range balances {
		if balances.ETH != nil {
			eth, _ := balances.ETH.Float64()
			o.aggregatorBalance.WithLabelValues(m.Network().GetName(), m.Provider(), address.Hex(), ETH).Set(eth)
		}

		if balances.POL != nil {
			pol, _ := balances.POL.Float64()
			o.aggregatorBalance.WithLabelValues(m.Network().GetName(), m.Provider(), address.Hex(), POL).Set(pol)
		}
	}
}

func (o *RollupManagerObserver) notifyRollup(m Message, rollup *RollupData, id string) {
	if rollup.TrustedSequencerBalances.ETH != nil {
		eth, _ := rollup.TrustedSequencerBalances.ETH.Float64()
		o.trustedSequencerBalance.WithLabelValues(m.Network().GetName(), m.Provider(), id, ETH).Set(eth)
	}

	if rollup.TrustedSequencerBalances.POL != nil {
		pol, _ := rollup.TrustedSequencerBalances.POL.Float64()
		o.trustedSequencerBalance.WithLabelValues(m.Network().GetName(), m.Provider(), id, POL).Set(pol)
	}

	if rollup.LastBatchSequenced != nil {
		lbs := float64(*rollup.LastBatchSequenced)
		o.lastBatchSequenced.WithLabelValues(m.Network().GetName(), m.Provider(), id).Set(lbs)
	}

	if rollup.LastSequencedTimestamp != nil {
		seconds := time.Since(time.Unix(int64(*rollup.LastSequencedTimestamp), 0)).Seconds()
		o.timeSinceLastSequenced.WithLabelValues(m.Network().GetName(), m.Provider(), id).Set(seconds)
	}

	for _, seconds := range rollup.TimeBetweenSequencedBatches {
		o.timeBetweenSequencedBatches.WithLabelValues(m.Network().GetName(), m.Provider(), id).Observe(float64(seconds))
	}

	for _, tx := range rollup.SequencedBatchesTxFees {
		gwei, _ := weiToGwei(tx.Fee).Float64()
		o.sequencedBatchesTxFee.WithLabelValues(m.Network().GetName(), m.Provider(), id, tx.Address.Hex()).Observe(gwei)
		o.observedSequencedBatches.WithLabelValues(m.Network().GetName(), m.Provider(), id, tx.Address.Hex()).Inc()
	}

	if rollup.LastForceBatch != nil {
		lfb := float64(*rollup.LastForceBatch)
		o.lastForceBatch.WithLabelValues(m.Network().GetName(), m.Provider(), id).Set(lfb)
	}

	if rollup.LastForceBatchSequenced != nil {
		lfbs := float64(*rollup.LastForceBatchSequenced)
		o.lastForceBatch.WithLabelValues(m.Network().GetName(), m.Provider(), id).Set(lfbs)
	}

	if rollup.LastVerifiedBatch != nil {
		lvb := float64(*rollup.LastVerifiedBatch)
		o.lastVerifiedBatch.WithLabelValues(m.Network().GetName(), m.Provider(), id).Set(lvb)
	}

	if rollup.LastVerifiedTimestamp != nil {
		seconds := time.Since(time.Unix(int64(*rollup.LastVerifiedTimestamp), 0)).Seconds()
		pessimistic := fmt.Sprint(rollup.Pessimistic)
		o.timeSinceLastVerified.WithLabelValues(m.Network().GetName(), m.Provider(), id, pessimistic).Set(seconds)
	}

	for _, seconds := range rollup.TimeBetweenVerifiedBatches {
		o.timeBetweenVerifiedBatches.WithLabelValues(m.Network().GetName(), m.Provider(), id).Observe(float64(seconds))
	}

	for _, tx := range rollup.VerifiedBatchesTxFees {
		gwei, _ := weiToGwei(tx.Fee).Float64()
		o.verifiedBatchesTxFee.WithLabelValues(m.Network().GetName(), m.Provider(), id, tx.Address.Hex()).Observe(gwei)
		o.observedVerifiedBatches.WithLabelValues(m.Network().GetName(), m.Provider(), id, tx.Address.Hex()).Inc()
	}

	if rollup.ChainID != nil {
		o.chainID.WithLabelValues(m.Network().GetName(), m.Provider(), id).Set(float64(*rollup.ChainID))
	}
}

func (o *RollupManagerObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.RollupManager, o)

	o.lastBatchSequenced = metrics.NewGauge(
		metrics.RPC,
		"zkevm_last_batch_sequenced",
		"The last batch sequenced number",
		"rollup",
	)
	o.timeSinceLastSequenced = metrics.NewGauge(
		metrics.RPC,
		"zkevm_time_since_last_sequenced",
		"The time since the last sequenced batch (in seconds)",
		"rollup",
	)
	o.totalSequencedBatches = metrics.NewGauge(
		metrics.RPC,
		"zkevm_total_sequenced_batches",
		"The total number of sequenced batches",
	)
	o.timeBetweenSequencedBatches = metrics.NewHistogram(
		metrics.RPC,
		"zkevm_time_between_sequenced_batches",
		"The time between sequenced batches (in seconds)",
		newExponentialBuckets(2, 14),
		"rollup",
	)
	o.sequencedBatchesTxFee = metrics.NewHistogram(
		metrics.RPC,
		"zkevm_sequenced_batches_tx_fee",
		"The transaction fee of OnSequencedBatches events (in gwei)",
		newExponentialBuckets(10, 9),
		"rollup",
		"address",
	)
	o.observedSequencedBatches = metrics.NewCounter(
		metrics.RPC,
		"zkevm_observed_sequenced_batches",
		"The number of sequenced batches observed",
		"rollup",
		"address",
	)

	o.lastVerifiedBatch = metrics.NewGauge(
		metrics.RPC,
		"zkevm_last_verified_batch",
		"The last verified batch number",
		"rollup",
	)
	o.timeSinceLastVerified = metrics.NewGauge(
		metrics.RPC,
		"zkevm_time_since_last_verified",
		"The time since the last verified batch (in seconds)",
		"rollup",
		"pessimistic",
	)
	o.totalVerifiedBatches = metrics.NewGauge(
		metrics.RPC,
		"zkevm_total_verified_batches",
		"The total number of verified batches",
	)
	o.timeBetweenVerifiedBatches = metrics.NewHistogram(
		metrics.RPC,
		"zkevm_time_between_verified_batches",
		"The time between verified batches (in seconds)",
		newExponentialBuckets(2, 14),
		"rollup",
	)
	o.verifiedBatchesTxFee = metrics.NewHistogram(
		metrics.RPC,
		"zkevm_verified_batches_tx_fee",
		"The transaction fee of verifyBatches and verifyBatchesTrustedAggregator events (in gwei)",
		newExponentialBuckets(10, 9),
		"rollup",
		"address",
	)
	o.observedVerifiedBatches = metrics.NewCounter(
		metrics.RPC,
		"zkevm_observed_verified_batches",
		"The number of verified batches observed",
		"rollup",
		"address",
	)

	o.lastForceBatch = metrics.NewGauge(
		metrics.RPC,
		"zkevm_last_force_batch",
		"The last force batch number",
		"rollup",
	)
	o.lastForceBatchSequenced = metrics.NewGauge(
		metrics.RPC,
		"zkevm_last_force_batch_sequenced",
		"The last force batch sequenced number",
		"rollup",
	)

	o.chainID = metrics.NewGauge(
		metrics.RPC,
		"zkevm_rollup_chain_id",
		"The rollup chain ID",
		"rollup",
	)

	o.batchFee = metrics.NewGauge(
		metrics.RPC,
		"zkevm_batch_fee",
		"The batch fee (gwei)",
	)
	o.rewardPerBatch = metrics.NewGauge(
		metrics.RPC,
		"zkevm_reward_per_batch",
		"The reward per batch (gwei)",
	)

	o.trustedSequencerBalance = metrics.NewGauge(
		metrics.RPC,
		"zkevm_trusted_sequencer_balance",
		"The trusted sequencer balance (wei)",
		"rollup",
		"token",
	)
	o.aggregatorBalance = metrics.NewGauge(
		metrics.RPC,
		"zkevm_aggregator_balance",
		"The aggregator balance (wei)",
		"address",
		"token",
	)
	o.timeSinceLastAggregation = metrics.NewGauge(
		metrics.RPC,
		"zkevm_time_since_last_aggregation",
		"The time since the last aggregation",
	)

	o.rollupCount = metrics.NewGauge(
		metrics.RPC,
		"zkevm_rollup_count",
		"The number of rollups",
	)
	o.rollupTypeCount = metrics.NewGauge(
		metrics.RPC,
		"zkevm_rollup_type_count",
		"The number of rollup types",
	)
}

func (o *RollupManagerObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{
		o.lastBatchSequenced,
		o.timeSinceLastSequenced,
		o.totalSequencedBatches,
		o.timeBetweenSequencedBatches,
		o.sequencedBatchesTxFee,
		o.observedSequencedBatches,

		o.lastVerifiedBatch,
		o.timeSinceLastVerified,
		o.totalVerifiedBatches,
		o.timeBetweenVerifiedBatches,
		o.verifiedBatchesTxFee,
		o.observedVerifiedBatches,

		o.lastForceBatch,
		o.lastForceBatchSequenced,

		o.chainID,

		o.batchFee,
		o.rewardPerBatch,

		o.trustedSequencerBalance,
		o.aggregatorBalance,

		o.rollupCount,
		o.rollupTypeCount,
	}
}

type TimeToMine struct {
	Seconds        float64
	GasPrice       *big.Int
	GasPriceFactor int64
}

type TimeToMineObserver struct {
	timeToMine *prometheus.HistogramVec
	gasPrice   *prometheus.HistogramVec
}

func (o *TimeToMineObserver) Notify(ctx context.Context, m Message) {
	data := m.Data().(*TimeToMine)

	o.timeToMine.WithLabelValues(m.Network().GetName(), m.Provider(), fmt.Sprint(data.GasPriceFactor)).Observe(data.Seconds)

	gasPrice, _ := weiToGwei(data.GasPrice).Float64()
	o.gasPrice.WithLabelValues(m.Network().GetName(), m.Provider(), fmt.Sprint(data.GasPriceFactor)).Observe(gasPrice)
}

func (o *TimeToMineObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.TimeToMine, o)

	o.timeToMine = metrics.NewHistogram(
		metrics.RPC,
		"time_to_mine",
		"Time it takes for sent transaction to be included in a block (in seconds)",
		newExponentialBuckets(2, 8),
		"gas_price_factor",
	)
	o.gasPrice = metrics.NewHistogram(
		metrics.RPC,
		"time_to_mine_gas_price",
		"The gas price for the time to mine transactions (in gwei)",
		newExponentialBuckets(2, 10),
		"gas_price_factor",
	)
}

func (o *TimeToMineObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.timeToMine, o.gasPrice}
}

type AccountBalances map[common.Address]*TokenBalances

type AccountBalancesObserver struct {
	balance *prometheus.GaugeVec
}

func (o *AccountBalancesObserver) Notify(ctx context.Context, m Message) {
	data := m.Data().(AccountBalances)

	for account, balances := range data {
		address := account.Hex()

		if balances.ETH != nil {
			eth, _ := balances.ETH.Float64()
			o.balance.WithLabelValues(m.Network().GetName(), m.Provider(), address, ETH).Set(eth)
		}

		if balances.POL != nil {
			pol, _ := balances.POL.Float64()
			o.balance.WithLabelValues(m.Network().GetName(), m.Provider(), address, POL).Set(pol)
		}
	}

}

func (o *AccountBalancesObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.AccountBalances, o)

	o.balance = metrics.NewGauge(
		metrics.RPC,
		"account_balance",
		"The account balance (wei)",
		"address",
		"token",
	)
}

func (o *AccountBalancesObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.balance}
}

type TrustedBatchObserver struct {
	length *prometheus.HistogramVec
}

func (o *TrustedBatchObserver) Notify(ctx context.Context, m Message) {
	batch := m.Data().(*zkevmtypes.Batch)

	length := float64(len(batch.Transactions))
	o.length.WithLabelValues(m.Network().GetName(), m.Provider()).Observe(length)
}

func (o *TrustedBatchObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.TrustedBatch, o)

	o.length = metrics.NewHistogram(
		metrics.RPC,
		"transactions_per_batch",
		"The number of transactions per trusted batch",
		newExponentialBuckets(2, 8),
	)
}

func (o *TrustedBatchObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.length}
}

type TimeToFinalizedObserver struct {
	gauge *prometheus.GaugeVec
}

func (o *TimeToFinalizedObserver) Notify(ctx context.Context, m Message) {
	data := m.Data().(*uint64)

	seconds := float64(*data)
	o.gauge.WithLabelValues(m.Network().GetName(), m.Provider()).Set(seconds)
}

func (o *TimeToFinalizedObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.TimeToFinalized, o)

	o.gauge = metrics.NewGauge(
		metrics.RPC,
		"time_to_finalized",
		"The time difference between the latest block and the last finalized block (in seconds)",
	)
}

func (o *TimeToFinalizedObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.gauge}
}
