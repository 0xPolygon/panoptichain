package provider

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"slices"
	"strconv"
	"strings"
	"time"

	zkevmtypes "github.com/0xPolygonHermez/zkevm-node/jsonrpc/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/rs/zerolog"

	"github.com/0xPolygon/panoptichain/api"
	"github.com/0xPolygon/panoptichain/blockbuffer"
	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/contracts"
	"github.com/0xPolygon/panoptichain/log"
	"github.com/0xPolygon/panoptichain/network"
	"github.com/0xPolygon/panoptichain/observer"
	"github.com/0xPolygon/panoptichain/observer/topics"
	"github.com/0xPolygon/panoptichain/util"
)

// RPCProvider is the generic struct for all EVM style JSON RPC services.
type RPCProvider struct {
	url              string
	network          network.Network
	label            string
	bus              *observer.EventBus
	interval         time.Duration
	logger           zerolog.Logger
	blockNumber      uint64
	prevBlockNumber  uint64
	finalizedHeight  uint64
	blockBuffer      *blockbuffer.BlockBuffer
	txPool           *observer.TransactionPool
	refreshStateTime *time.Duration
	contracts        config.Contracts
	timeToMine       *config.TimeToMine
	accounts         []string
	accountBalances  observer.AccountBalances
	timeToFinalized  *uint64
	blockLookBack    uint64
	hasTxPool        bool

	// PoS
	stateSync            map[bool]*observer.StateSync
	checkpointSignatures map[bool]*observer.CheckpointSignatures
	validatorBalances    observer.ValidatorWalletBalances
	missedBlockProposal  observer.MissedBlockProposal

	// zkEVM
	batches        observer.ZkEVMBatches
	trustedBatches []*zkevmtypes.Batch

	globalExitRoot   *observer.ExitRoot
	mainnetExitRoot  *observer.ExitRoot
	rollupExitRoot   *observer.ExitRoot
	rollupExitRootL2 *observer.ExitRoot

	bridgeEvents []*contracts.PolygonZkEVMBridgeV2BridgeEvent
	claimEvents  []*contracts.PolygonZkEVMBridgeV2ClaimEvent

	bridgeEventTimes observer.BridgeEventTimes
	claimEventTimes  observer.ClaimEventTimes

	depositCount            *big.Int
	lastUpdatedDepositCount *uint32

	rollupManager       *observer.RollupManager
	trustedSequencers   map[uint32]*RPCProvider
	trustedSequencerURL chan string

	// These contract addresses will be derived from the PolygonRollupManager
	// contract.
	rollupContracts       map[uint32]common.Address
	polTokenAddress       *common.Address
	globalExitRootAddress *common.Address
}

// NewRPCProvider creates a new RPC provider.
func NewRPCProvider(n network.Network, eb *observer.EventBus, cfg config.RPC) *RPCProvider {
	// Look back this number of blocks when filtering event logs.
	blb := config.DefaultBlockLookBack
	if cfg.BlockLookBack != nil {
		blb = *cfg.BlockLookBack
	}

	return &RPCProvider{
		url:                  cfg.URL,
		network:              n,
		label:                cfg.Label,
		bus:                  eb,
		blockBuffer:          blockbuffer.NewBlockBuffer(128),
		interval:             GetInterval(cfg.Interval),
		logger:               NewLogger(n, cfg.Label),
		refreshStateTime:     new(time.Duration),
		contracts:            cfg.Contracts,
		timeToMine:           cfg.TimeToMine,
		accounts:             cfg.Accounts,
		accountBalances:      make(observer.AccountBalances),
		stateSync:            make(map[bool]*observer.StateSync),
		checkpointSignatures: make(map[bool]*observer.CheckpointSignatures),
		validatorBalances:    make(observer.ValidatorWalletBalances),
		missedBlockProposal:  make(observer.MissedBlockProposal),
		bridgeEventTimes:     make(observer.BridgeEventTimes),
		claimEventTimes:      make(observer.ClaimEventTimes),
		trustedSequencers:    make(map[uint32]*RPCProvider),
		trustedSequencerURL:  make(chan string),
		rollupContracts:      make(map[uint32]common.Address),
		blockLookBack:        blb,
		hasTxPool:            cfg.TxPool,
	}
}

// RefreshState is going to get the current head block and request all
// of the blocks between the current head and the last head. All of
// those blocks will be pushed into the buffer.
func (r *RPCProvider) RefreshState(ctx context.Context) error {
	defer timer(r.refreshStateTime)()

	c, err := ethclient.DialContext(ctx, r.url)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to create the client")
		return err
	}

	r.refreshBlockBuffer(ctx, c)

	r.refreshStateSync(ctx, c, true)
	r.refreshStateSync(ctx, c, false)
	r.refreshCheckpoint(ctx, c)

	if r.network.IsPolygonPoS() {
		r.refreshValidatorBalances(ctx, c)
		r.refreshMissedBlockProposal(ctx, c)
	}

	if r.hasTxPool {
		r.refreshTxPoolStatus(ctx, c)
	}

	r.refreshTimeToMine(ctx, c)
	r.refreshAccountBalances(ctx, c)

	if r.network.IsPolygonZkEVM() {
		r.refreshBatches(ctx, c)
	}

	r.refreshRollupManager(ctx, c)
	r.refreshExitRoots(ctx, c)
	r.refreshExitRootsL2(ctx, c)
	r.refreshBridge(ctx, c)

	return nil
}

func (r *RPCProvider) PublishEvents(ctx context.Context) error {
	for i := r.prevBlockNumber + 1; i <= r.blockNumber && r.prevBlockNumber != 0; i++ {
		b, err := r.blockBuffer.GetBlock(i)
		if err != nil {
			continue
		}
		block, ok := b.(*types.Block)
		if !ok {
			continue
		}

		m := observer.NewMessage(r.network, r.label, block)
		r.bus.Publish(ctx, topics.NewEVMBlock, m)

		pb, err := r.blockBuffer.GetBlock(b.Number().Uint64() - 1)
		if err != nil {
			continue
		}
		prev, ok := pb.(*types.Block)
		if !ok {
			continue
		}

		interval := observer.NewMessage(r.network, r.label, block.Time()-prev.Time())
		r.bus.Publish(ctx, topics.BlockInterval, interval)
	}

	if len(r.missedBlockProposal) > 0 {
		missedBlockProposal := observer.NewMessage(r.network, r.label, r.missedBlockProposal)
		r.bus.Publish(ctx, topics.BorMissedBlockProposal, missedBlockProposal)
	}

	for _, stateSync := range r.stateSync {
		r.bus.Publish(ctx, topics.BorStateSync, observer.NewMessage(r.network, r.label, stateSync))
	}

	for _, checkpointSignatures := range r.checkpointSignatures {
		m := observer.NewMessage(r.network, r.label, checkpointSignatures)
		r.bus.Publish(ctx, topics.CheckpointSignatures, m)
	}

	if len(r.validatorBalances) > 0 {
		validatorWalletBalance := observer.NewMessage(r.network, r.label, r.validatorBalances)
		r.bus.Publish(ctx, topics.ValidatorWallet, validatorWalletBalance)
	}

	if r.txPool != nil {
		txPool := observer.NewMessage(r.network, r.label, r.txPool)
		r.bus.Publish(ctx, topics.TransactionPool, txPool)
	}

	if r.batches.TrustedBatch.Number > 0 || r.batches.VirtualBatch.Number > 0 || r.batches.VerifiedBatch.Number > 0 {
		r.bus.Publish(ctx, topics.ZkEVMBatches, observer.NewMessage(r.network, r.label, r.batches))
	}

	if r.globalExitRoot != nil || r.mainnetExitRoot != nil || r.rollupExitRoot != nil {
		er := &observer.ExitRoots{
			GlobalExitRoot:  r.globalExitRoot,
			MainnetExitRoot: r.mainnetExitRoot,
			RollupExitRoot:  r.rollupExitRoot,
		}
		r.bus.Publish(ctx, topics.ExitRoots, observer.NewMessage(r.network, r.label, er))
	}

	if r.rollupExitRootL2 != nil {
		er := &observer.ExitRoots{
			RollupExitRoot: r.rollupExitRootL2,
		}
		r.bus.Publish(ctx, topics.ExitRoots, observer.NewMessage(r.network, r.label, er))
	}

	if r.depositCount != nil || r.lastUpdatedDepositCount != nil {
		m := observer.NewMessage(r.network, r.label, &observer.DepositCounts{
			DepositCount:            r.depositCount,
			LastUpdatedDepositCount: r.lastUpdatedDepositCount,
		})
		r.bus.Publish(ctx, topics.DepositCounts, m)
	}

	for _, bridgeEvent := range r.bridgeEvents {
		r.bus.Publish(ctx, topics.BridgeEvent, observer.NewMessage(r.network, r.label, bridgeEvent))
	}

	for _, claimEvent := range r.claimEvents {
		r.bus.Publish(ctx, topics.ClaimEvent, observer.NewMessage(r.network, r.label, claimEvent))
	}

	if len(r.bridgeEventTimes) > 0 {
		m := observer.NewMessage(r.network, r.label, r.bridgeEventTimes)
		r.bus.Publish(ctx, topics.BridgeEventTimes, m)
	}

	if len(r.claimEventTimes) > 0 {
		m := observer.NewMessage(r.network, r.label, r.claimEventTimes)
		r.bus.Publish(ctx, topics.ClaimEventTimes, m)
	}

	if r.rollupManager != nil {
		m := observer.NewMessage(r.network, r.label, r.rollupManager)
		r.bus.Publish(ctx, topics.RollupManager, m)
	}

	if len(r.accountBalances) > 0 {
		m := observer.NewMessage(r.network, r.label, r.accountBalances)
		r.bus.Publish(ctx, topics.AccountBalances, m)
	}

	for _, batch := range r.trustedBatches {
		m := observer.NewMessage(r.network, r.label, batch)
		r.bus.Publish(ctx, topics.TrustedBatch, m)
	}

	if r.timeToFinalized != nil {
		m := observer.NewMessage(r.network, r.label, r.timeToFinalized)
		r.bus.Publish(ctx, topics.TimeToFinalized, m)
	}

	if r.finalizedHeight > 0 {
		m := observer.NewMessage(r.network, r.label, r.finalizedHeight)
		r.bus.Publish(ctx, topics.FinalizedHeight, m)
	}

	r.bus.Publish(ctx, topics.RefreshStateTime, observer.NewMessage(r.network, r.label, r.refreshStateTime))

	return nil
}

func (r *RPCProvider) PollingInterval() time.Duration {
	return r.interval
}

func (r *RPCProvider) refreshBlockBuffer(ctx context.Context, c *ethclient.Client) (err error) {
	r.prevBlockNumber = r.blockNumber
	r.blockNumber, err = c.BlockNumber(ctx)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to get block number")
		return err
	}

	r.logger.Info().Uint64("block_number", r.blockNumber).Msg("Refreshing block state")

	if r.prevBlockNumber != 0 && r.prevBlockNumber != r.blockNumber {
		r.fillRange(ctx, r.prevBlockNumber, c)
	}

	finalized, err := c.HeaderByNumber(ctx, big.NewInt(int64(rpc.FinalizedBlockNumber)))
	if err != nil {
		r.logger.Warn().Err(err).Msg("Failed to get finalized block header")
		return err
	}
	r.finalizedHeight = finalized.Number.Uint64()

	latest, err := c.HeaderByNumber(ctx, big.NewInt(int64(r.blockNumber)))
	if err != nil {
		r.logger.Warn().Err(err).Msg("Failed to get latest block header")
		return err
	}

	diff := latest.Time - finalized.Time
	r.timeToFinalized = &diff

	return nil
}

func (r *RPCProvider) getFilterOpts() *bind.FilterOpts {
	opts := bind.FilterOpts{End: &r.blockNumber}

	if r.prevBlockNumber > 0 {
		opts.Start = r.prevBlockNumber
	} else if r.blockLookBack < r.blockNumber {
		opts.Start = r.blockNumber - r.blockLookBack
	}

	log.Trace().
		Any("opts", opts).
		Any("block_number", r.blockNumber).
		Any("prev_block_number", r.prevBlockNumber).
		Any("block_look_back", r.blockLookBack).
		Msg("Getting filter options")

	return &opts
}

// cast call --rpc-url https://eth.llamarpc.com 0x28e4F3a7f651294B9564800b2D01f35189A5bFbE 'function counter() view returns(uint256)'
// cast call --rpc-url https://polygon-rpc.com 0x0000000000000000000000000000000000001001 'function lastStateId() view returns(uint256)'
func (r *RPCProvider) refreshStateSync(ctx context.Context, c *ethclient.Client, finalized bool) error {
	var counter, blockNumber *big.Int
	if finalized {
		blockNumber = big.NewInt(rpc.FinalizedBlockNumber.Int64())
	}

	co := bind.CallOpts{
		Context:     ctx,
		BlockNumber: blockNumber,
	}

	if r.contracts.StateSyncSenderAddress != nil {
		address := common.HexToAddress(*r.contracts.StateSyncSenderAddress)
		ss, err := contracts.NewStateSender(address, c)
		if err != nil {
			r.logger.Error().Err(err).Msg("Failed to bind state sender contract")
			return err
		}

		counter, err = ss.Counter(&co)
		if err != nil {
			r.logger.Error().Err(err).Msg("Failed to get state sender counter")
			return err
		}
	} else if r.contracts.StateSyncReceiverAddress != nil {
		address := common.HexToAddress(*r.contracts.StateSyncReceiverAddress)
		sr, err := contracts.NewStateReceiver(address, c)
		if err != nil {
			r.logger.Error().Err(err).Msg("Failed to bind state receiver contract")
			return err
		}

		counter, err = sr.LastStateId(&co)
		if err != nil {
			r.logger.Error().Err(err).Msg("Failed to get state receiver counter")
			return err
		}
	} else {
		return nil
	}

	stateSync := &observer.StateSync{
		ID:        counter.Uint64(),
		Time:      time.Now(),
		Finalized: finalized,
	}

	if r.stateSync[finalized] == nil || r.stateSync[finalized].ID != stateSync.ID {
		r.stateSync[finalized] = stateSync
	}

	return nil
}

func (r *RPCProvider) refreshCheckpoint(ctx context.Context, c *ethclient.Client) {
	if r.contracts.CheckpointAddress == nil {
		return
	}

	address := common.HexToAddress(*r.contracts.CheckpointAddress)
	contract, err := contracts.NewRootChain(address, c)
	if contract == nil || err != nil {
		r.logger.Warn().Err(err).Msg("Failed to bind root chain contract")
		return
	}

	iter, err := contract.FilterNewHeaderBlock(r.getFilterOpts(), nil, nil, nil)
	if iter == nil || err != nil {
		r.logger.Error().Err(err).Msg("Failed to filter NewHeaderBlock events")
		return
	}

	// Get the last NewHeaderBlock event.
	var event *contracts.RootChainNewHeaderBlock
	for iter.Next() && iter.Event != nil {
		event = iter.Event
	}

	if event == nil {
		r.logger.Debug().Msg("No NewHeaderBlock events found")
		return
	}

	// Grab that block so that we know when the transaction was mined.
	block, err := c.BlockByHash(ctx, event.Raw.BlockHash)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to get block by hash")
		return
	}

	r.logger.Trace().Any("event", event).Str("network", r.network.GetName()).Msg("Latest NewHeaderBlock event")
	tx, _, err := c.TransactionByHash(ctx, event.Raw.TxHash)
	if err != nil {
		r.logger.Error().Err(err).Msg("Could not find submitCheckpoint transaction")
		return
	}

	abi, err := contracts.RootChainMetaData.GetAbi()
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to get root chain ABI")
		return
	}

	method, err := abi.MethodById(tx.Data()[:4])
	if err != nil {
		r.logger.Error().Err(err).Msg("Contract method not found")
		return
	}

	inputs := make(map[string]any)
	if err := method.Inputs.UnpackIntoMap(inputs, tx.Data()[4:]); err != nil {
		r.logger.Error().Err(err).Msg("Failed to unpack input params")
		return
	}

	data := inputs["data"].([]byte)
	sigs := inputs["sigs"].([][3]*big.Int)
	vote := crypto.Keccak256(append([]byte{1}, data...))
	signers := make([]common.Address, 0, len(sigs))

	for _, sig := range sigs {
		R := padLeft(sig[0].Bytes(), 32)
		s := padLeft(sig[1].Bytes(), 32)
		v := padLeft(new(big.Int).Sub(sig[2], big.NewInt(27)).Bytes(), 1)

		signature := bytes.Join([][]byte{R, s, v}, nil)

		key, err := crypto.SigToPub(vote, signature)
		if err != nil {
			r.logger.Warn().Err(err).Msg("Failed to get public key from signature")
			continue
		}

		address := crypto.PubkeyToAddress(*key)
		signers = append(signers, address)
	}

	r.refreshFinalizedCheckpoint(ctx, c)

	finalized := false
	cs := r.checkpointSignatures[finalized]
	seen := cs != nil && event.HeaderBlockId.Cmp(cs.Event.HeaderBlockId) == 0
	r.checkpointSignatures[finalized] = &observer.CheckpointSignatures{
		Event:     event,
		Block:     block,
		Signers:   signers,
		Seen:      seen,
		Finalized: finalized,
	}
}

func (r *RPCProvider) refreshFinalizedCheckpoint(ctx context.Context, c *ethclient.Client) {
	finalized := true
	latest := r.checkpointSignatures[!finalized]
	if latest == nil {
		return
	}

	block, err := c.BlockByNumber(ctx, big.NewInt(rpc.FinalizedBlockNumber.Int64()))
	if err != nil {
		log.Error().Err(err).Msg("Failed to get block header by number")
		return
	}

	// The block with the checkpoint transaction hasn't been finalized yet.
	if block.Number().Cmp(latest.Block.Number()) > 0 {
		return
	}

	r.checkpointSignatures[finalized] = &observer.CheckpointSignatures{
		Event:     latest.Event,
		Block:     block,
		Signers:   latest.Signers,
		Seen:      latest.Seen,
		Finalized: finalized,
	}
}

type rpcTransaction struct {
	tx *types.Transaction
	txExtraInfo
}

func (tx *rpcTransaction) UnmarshalJSON(msg []byte) error {
	if err := json.Unmarshal(msg, &tx.tx); err != nil {
		return err
	}
	return json.Unmarshal(msg, &tx.txExtraInfo)
}

type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

type rpcBlock struct {
	Hash         common.Hash         `json:"hash"`
	Transactions []rpcTransaction    `json:"transactions"`
	UncleHashes  []common.Hash       `json:"uncles"`
	Withdrawals  []*types.Withdrawal `json:"withdrawals,omitempty"`
}

// getBlockByNumber gets the block given a block number. This is lifted from
// https://github.com/ethereum/go-ethereum/blob/master/ethclient/ethclient.go
// with the change that unsupported transactions are treated as legacy
// transactions. This allows observation of chains that use transaction types
// that are not supported by Geth.
func (r *RPCProvider) getBlockByNumber(ctx context.Context, n *big.Int, c *ethclient.Client) (*types.Block, error) {
	var raw json.RawMessage
	err := c.Client().Call(&raw, "eth_getBlockByNumber", hexutil.EncodeBig(n), true)
	if err != nil {
		return nil, err
	}

	var head *types.Header
	if err := json.Unmarshal(raw, &head); err != nil {
		return nil, err
	}

	var body map[string]any
	if err := json.Unmarshal(raw, &body); err != nil {
		return nil, err
	}

	transactions, ok := body["transactions"].([]any)
	if !ok {
		return nil, errors.New("transactions type assertion failed")
	}

	for _, tx := range transactions {
		tx, ok := tx.(map[string]any)
		if !ok {
			continue
		}

		hex, ok := tx["type"].(string)
		if !ok {
			continue
		}

		decimal, err := hexutil.DecodeUint64(hex)
		if err != nil {
			log.Warn().Err(err).Send()
		}

		// Remove the transaction type field which would allow the transaction to be
		// treated as legacy.
		if decimal > 3 {
			delete(tx, "type")
		}
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var block rpcBlock
	if err := json.Unmarshal(bytes, &block); err != nil {
		return nil, err
	}

	var uncles []*types.Header
	if len(block.UncleHashes) > 0 {
		uncles = make([]*types.Header, len(block.UncleHashes))
		reqs := make([]rpc.BatchElem, len(block.UncleHashes))

		for i := range reqs {
			reqs[i] = rpc.BatchElem{
				Method: "eth_getUncleByBlockHashAndIndex",
				Args:   []any{block.Hash, hexutil.EncodeUint64(uint64(i))},
				Result: &uncles[i],
			}
		}

		if err := c.Client().BatchCallContext(ctx, reqs); err != nil {
			return nil, err
		}

		for i := range reqs {
			if reqs[i].Error != nil {
				return nil, reqs[i].Error
			}
			if uncles[i] == nil {
				return nil, fmt.Errorf("got null header for uncle %d of block %x", i, block.Hash[:])
			}
		}
	}

	txs := make([]*types.Transaction, len(block.Transactions))
	for i, tx := range block.Transactions {
		txs[i] = tx.tx
	}

	return types.NewBlockWithHeader(head).WithBody(txs, uncles).WithWithdrawals(block.Withdrawals), nil
}

// fillRange pulls all of the blocks between the start and the current head.
func (r *RPCProvider) fillRange(ctx context.Context, start uint64, c *ethclient.Client) {
	r.logger.Debug().
		Uint64("start_block", start).
		Uint64("end_block", r.blockNumber).
		Msg("Filling block range")

	for i := start + 1; i <= r.blockNumber; i++ {
		num := new(big.Int).SetUint64(i)
		block, err := c.BlockByNumber(ctx, num)

		// Retry fetching the block if it's an unrecognized transaction.
		if errors.Is(err, types.ErrTxTypeNotSupported) {
			block, err = r.getBlockByNumber(ctx, num, c)
		}

		if err != nil {
			r.logger.Warn().Err(err).Uint64("block_number", i).Msg("Failed to get block")
			break
		}

		r.blockBuffer.PutBlock(block)
	}
}

func padLeft(data []byte, size int) []byte {
	if len(data) < size {
		n := size - len(data)
		padded := make([]byte, n)

		return append(padded, data...)
	}

	return data[len(data)-size:]
}

func (r *RPCProvider) refreshValidatorBalances(ctx context.Context, c *ethclient.Client) (err error) {
	signers, err := api.Signers(r.network)
	if err != nil {
		r.logger.Warn().Err(err).Msg("Failed to get signers validator map")
		return
	}

	reqs := make([]rpc.BatchElem, 0, len(signers))
	addresses := make([]string, 0, len(signers))

	for address := range signers {
		addr := common.HexToAddress(address)
		addresses = append(addresses, address)
		reqs = append(reqs, rpc.BatchElem{
			Method: "eth_getBalance",
			Args:   []any{addr, "latest"},
			Result: new(json.RawMessage),
		})
	}

	err = c.Client().BatchCallContext(ctx, reqs)
	if err != nil {
		r.logger.Warn().Err(err).Msg("Failed to execute batch request for validator balances")
		return err
	}

	for i, req := range reqs {
		logger := r.logger.Warn().Int("index", i)

		if req.Error != nil {
			logger.Err(req.Error).Msg("Failed to get validator balance")
			continue
		}

		var b string
		if err := json.Unmarshal(*req.Result.(*json.RawMessage), &b); err != nil {
			logger.Err(err).Msg("Failed to unmarshal validator balance")
			continue
		}

		balance, err := hexutil.DecodeBig(b)
		if err != nil {
			logger.Err(err).Msg("Failed to decode validator balance")
			continue
		}

		address := addresses[i]
		r.validatorBalances[address] = balance
	}

	return nil
}

type SignerInfo struct {
	Difficulty int    `json:"Difficulty"`
	Signer     string `json:"Signer"`
}

type SnapshotProposerSequence struct {
	Author  string       `json:"Author"`
	Diff    int          `json:"Diff"`
	Signers []SignerInfo `json:"Signers"`
}

func (r *RPCProvider) refreshMissedBlockProposal(ctx context.Context, c *ethclient.Client) error {
	for i := r.prevBlockNumber + 1; i <= r.blockNumber && r.prevBlockNumber != 0; i++ {
		var response SnapshotProposerSequence
		err := c.Client().CallContext(ctx, &response, "bor_getSnapshotProposerSequence", hexutil.EncodeUint64(i))
		if err != nil {
			r.logger.Warn().Err(err).Msg("Failed to execute request for snapshot proposer sequence")
			return err
		}

		b, err := r.blockBuffer.GetBlock(i)
		if err != nil {
			continue
		}
		block := b.(*types.Block)

		bytes, err := api.Ecrecover(block.Header())
		if err != nil {
			r.logger.Warn().Err(err).Msg("Failed to get block signer")
			continue
		}

		signer := "0x" + hex.EncodeToString(bytes)
		if signer == response.Author {
			continue
		}

		for _, info := range response.Signers {
			if signer == info.Signer {
				break
			}

			r.missedBlockProposal[i] = append(r.missedBlockProposal[i], info.Signer)
		}
	}

	return nil
}

type TxPoolStatus struct {
	Pending string `json:"pending"`
	Queued  string `json:"queued"`
}

func (r *RPCProvider) refreshTxPoolStatus(ctx context.Context, c *ethclient.Client) error {
	var response TxPoolStatus
	err := c.Client().CallContext(ctx, &response, "txpool_status")
	if err != nil {
		r.logger.Warn().Err(err).Msg("Failed to execute request to get transaction pool status")
		return err
	}

	pending, err := strconv.ParseUint(strings.TrimPrefix(response.Pending, "0x"), 16, 0)
	if err != nil {
		return nil
	}
	queued, err := strconv.ParseUint(strings.TrimPrefix(response.Queued, "0x"), 16, 0)
	if err != nil {
		return nil
	}

	r.txPool = &observer.TransactionPool{
		Pending: pending,
		Queued:  queued,
	}

	return nil
}

// refreshTimeToMine sends a transaction to the network and records the time it
// took to be included in a block.
func (r *RPCProvider) refreshTimeToMine(ctx context.Context, c *ethclient.Client) error {
	if r.timeToMine == nil {
		return nil
	}

	gasPriceFactor := r.timeToMine.GasPriceFactor
	if gasPriceFactor == 0 {
		gasPriceFactor = 1
	}

	sender := common.HexToAddress(r.timeToMine.Sender)
	receiver := common.HexToAddress(r.timeToMine.Receiver)
	privateKey, err := crypto.HexToECDSA(r.timeToMine.SenderPrivateKey)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to parse SenderPrivateKey")
		return err
	}

	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("PublicKey type assertion failed")
		r.logger.Error().Err(err).Send()
		return err
	}

	address := crypto.PubkeyToAddress(*publicKey)
	if address != sender {
		err = fmt.Errorf("sender address mismatch %v != %v", sender, address)
		r.logger.Error().Err(err).Send()
		return err
	}

	gasPrice, err := c.SuggestGasPrice(ctx)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to get suggested gas price")
		return err
	}
	gasPrice.Mul(gasPrice, big.NewInt(r.timeToMine.GasPriceFactor))

	nonce, err := c.PendingNonceAt(ctx, sender)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to get pending nonce")
		return err
	}

	value := big.NewInt(r.timeToMine.Value)
	gasLimit := r.timeToMine.GasLimit
	data := []byte(r.timeToMine.Data)

	tx := types.NewTransaction(nonce, receiver, value, gasLimit, gasPrice, data)

	chainID, err := c.ChainID(ctx)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to get network ID")
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to sign transaction")
		return err
	}

	err = c.SendTransaction(ctx, signedTx)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to send transaction")
		return err
	}

	start := time.Now()

	// Generally, all messages sent to topics should be done in the PublishEvents
	// method. This is the exception because of its asynchronous nature. This
	// implementation reduces complexity by not needing to manage shared variables.
	go func() {
		ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
		defer cancel()

		_, err := bind.WaitMined(ctx, c, signedTx)
		if err != nil {
			r.logger.Error().Err(err).Msg("Failed to wait for transaction")
		}

		ttm := &observer.TimeToMine{
			Seconds:        time.Since(start).Seconds(),
			GasPrice:       gasPrice,
			GasPriceFactor: gasPriceFactor,
		}

		m := observer.NewMessage(r.network, r.label, ttm)
		r.bus.Publish(ctx, topics.TimeToMine, m)
	}()

	return nil
}

func (r *RPCProvider) refreshAccountBalances(ctx context.Context, c *ethclient.Client) {
	co := &bind.CallOpts{Context: ctx}

	for _, account := range r.accounts {
		address := common.HexToAddress(account)
		balances, ok := r.accountBalances[address]
		if !ok {
			balances = &observer.TokenBalances{}
			r.accountBalances[address] = balances
		}

		eth, err := c.BalanceAt(ctx, address, nil)
		if err != nil || eth == nil {
			r.logger.Error().Err(err).
				Any("address", address).
				Str("token", observer.ETH).
				Msg("Failed to get balance")
		} else {
			balances.ETH = eth
		}

		balances.POL = r.getPOL(c, address, co, balances.POL)
	}
}

func (r *RPCProvider) refreshBatches(ctx context.Context, c *ethclient.Client) {
	r.trustedBatches = nil
	prev := r.batches.TrustedBatch.Number

	r.refreshBatch(ctx, c, "zkevm_batchNumber", &r.batches.TrustedBatch)
	for i := prev + 1; i <= r.batches.TrustedBatch.Number && prev != 0; i++ {
		var batch zkevmtypes.Batch

		err := c.Client().CallContext(ctx, &batch, "zkevm_getBatchByNumber", i)
		if err != nil {
			r.logger.Warn().Err(err).Msg("Failed to get trusted batch by number")
			continue
		}

		r.trustedBatches = append(r.trustedBatches, &batch)
	}

	r.refreshBatch(ctx, c, "zkevm_virtualBatchNumber", &r.batches.VirtualBatch)
	r.refreshBatch(ctx, c, "zkevm_verifiedBatchNumber", &r.batches.VerifiedBatch)
}

func (r *RPCProvider) refreshBatch(ctx context.Context, c *ethclient.Client, endpoint string, batch *observer.ZkEVMBatch) {
	var response string
	err := c.Client().CallContext(ctx, &response, endpoint)
	if err != nil {
		r.logger.Warn().Err(err).Msgf("Failed to get %s", endpoint)
		return
	}

	result, err := strconv.ParseUint(response, 0, 0)
	if err != nil {
		r.logger.Warn().Err(err).Msgf("Failed to parse %s", endpoint)
		return
	}

	if result > batch.Number {
		batch.Number = result
		batch.Time = time.Now()
	}
}

// refreshExitRoot updates the exit root. If it has not been seen, set the exit
// root `Time` to `t`. If it has already been observed, the time will remain the
// same and the `Seen` value will be set to true.
func refreshExitRoot(er *observer.ExitRoot, bytes [32]byte, t time.Time) *observer.ExitRoot {
	hash := common.BytesToHash(bytes[:])
	if er == nil || er.Hash.Cmp(hash) != 0 {
		return &observer.ExitRoot{
			Hash: hash,
			Time: t,
		}
	}

	// Don't modify the passed in exit root, so return a new one.
	return &observer.ExitRoot{
		Hash: er.Hash,
		Time: er.Time,
		Seen: true,
	}
}

func (r *RPCProvider) refreshExitRoots(ctx context.Context, c *ethclient.Client) error {
	if r.globalExitRootAddress == nil {
		return nil
	}

	co := &bind.CallOpts{Context: ctx}
	contract, err := contracts.NewPolygonZkEVMGlobalExitRootV2(*r.globalExitRootAddress, c)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to bind global exit root contract")
		return nil
	}

	r.refreshGlobalExitRoot(ctx, c, contract, co)

	if mainnetExitRoot, err := contract.LastMainnetExitRoot(co); err != nil {
		r.logger.Error().Err(err).Msg("Failed to get last mainnet exit root")
	} else {
		r.mainnetExitRoot = refreshExitRoot(r.mainnetExitRoot, mainnetExitRoot, time.Now())
	}

	if rollupExitRoot, err := contract.LastRollupExitRoot(co); err != nil {
		r.logger.Error().Err(err).Msg("Failed to get last rollup exit root")
	} else {
		r.rollupExitRoot = refreshExitRoot(r.rollupExitRoot, rollupExitRoot, time.Now())
	}

	return nil
}

func (r *RPCProvider) refreshGlobalExitRoot(ctx context.Context, c *ethclient.Client, contract *contracts.PolygonZkEVMGlobalExitRootV2, co *bind.CallOpts) {
	globalExitRoot, err := contract.GetLastGlobalExitRoot(co)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to get last global exit root")
		return
	}

	t := time.Now()
	hash, err := contract.GlobalExitRootMap(co, globalExitRoot)
	if err != nil || hash == nil {
		r.logger.Error().Err(err).Msg("Failed to get block hash from global exit root map")
	} else {
		header, err := c.HeaderByHash(ctx, common.BigToHash(hash))
		if err != nil || header == nil {
			r.logger.Error().Err(err).Msg("Failed to get block header from global exit root block hash")
		} else {
			t = time.Unix(int64(header.Time), 0)
		}
	}

	r.globalExitRoot = refreshExitRoot(r.globalExitRoot, globalExitRoot, t)
}

func (r *RPCProvider) refreshExitRootsL2(ctx context.Context, c *ethclient.Client) error {
	if r.contracts.GlobalExitRootL2Address == nil {
		return nil
	}

	address := common.HexToAddress(*r.contracts.GlobalExitRootL2Address)
	contract, err := contracts.NewPolygonZkEVMGlobalExitRootL2(address, c)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to bind global exit root l2 contract")
		return err
	}

	co := bind.CallOpts{Context: ctx}
	rollupExitRoot, err := contract.LastRollupExitRoot(&co)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to get last rollup exit root")
		return err
	}

	r.rollupExitRootL2 = refreshExitRoot(r.rollupExitRootL2, rollupExitRoot, time.Now())

	return nil
}

func (r *RPCProvider) refreshBridge(ctx context.Context, c *ethclient.Client) error {
	if r.contracts.ZkEVMBridgeAddress == nil {
		return nil
	}

	r.bridgeEvents = nil
	r.claimEvents = nil

	co := bind.CallOpts{Context: ctx}
	address := common.HexToAddress(*r.contracts.ZkEVMBridgeAddress)
	contract, err := contracts.NewPolygonZkEVMBridgeV2(address, c)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to bind zkEVM bridge contract")
		return nil
	}

	dc, err := contract.DepositCount(&co)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to get deposit count")
	} else {
		r.depositCount = dc
	}

	ludc, err := contract.LastUpdatedDepositCount(&co)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to get last updated deposit count")
	} else {
		r.lastUpdatedDepositCount = &ludc
	}

	opts := r.getFilterOpts()
	r.refreshBridgeEvents(ctx, c, contract, opts)
	r.refreshClaimEvents(ctx, c, contract, opts)

	return nil
}

func (r *RPCProvider) refreshBridgeEvents(ctx context.Context, c *ethclient.Client, contract *contracts.PolygonZkEVMBridgeV2, opts *bind.FilterOpts) {
	iter, err := contract.FilterBridgeEvent(opts)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to filter bridge events")
		return
	}

	for iter.Next() && iter.Event != nil {
		event := iter.Event
		r.bridgeEvents = append(r.bridgeEvents, event)

		block, err := c.BlockByHash(ctx, event.Raw.BlockHash)
		if err != nil {
			r.logger.Error().Err(err).Msg("Failed to get block by hash")
			continue
		}

		networks := observer.BridgeEventNetworks{
			OriginNetwork:      event.OriginNetwork,
			DestinationNetwork: event.DestinationNetwork,
		}

		r.bridgeEventTimes[networks] = time.Unix(int64(block.Time()), 0)
	}
}

func (r *RPCProvider) refreshClaimEvents(ctx context.Context, c *ethclient.Client, contract *contracts.PolygonZkEVMBridgeV2, opts *bind.FilterOpts) {
	iter, err := contract.FilterClaimEvent(opts)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to filter claim events")
		return
	}

	for iter.Next() && iter.Event != nil {
		event := iter.Event
		r.claimEvents = append(r.claimEvents, event)

		block, err := c.BlockByHash(ctx, event.Raw.BlockHash)
		if err != nil {
			r.logger.Error().Err(err).Msg("Failed to get block by hash")
			continue
		}

		r.claimEventTimes[event.OriginNetwork] = time.Unix(int64(block.Time()), 0)
	}
}

func (r *RPCProvider) getPOL(c *ethclient.Client, address common.Address, co *bind.CallOpts, prev *big.Int) *big.Int {
	if r.polTokenAddress == nil {
		return prev
	}

	erc20, err := contracts.NewERC20(*r.polTokenAddress, c)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to bind ERC20 contract")
		return prev
	}

	pol, err := erc20.BalanceOf(co, address)
	if err != nil || pol == nil {
		r.logger.Error().Err(err).
			Any("address", address).
			Str("token", observer.POL).
			Msg("Failed to get balance")

		return prev
	}

	return pol
}

func (r *RPCProvider) refreshTrustedSequencerBalance(ctx context.Context, c *ethclient.Client, contract *contracts.PolygonZkEVMEtrog, co *bind.CallOpts, rollupID uint32) {
	address, err := contract.TrustedSequencer(co)
	if err != nil {
		r.logger.Error().Err(err).Msg("Could not get trusted sequencer address")
		return
	}

	balances := &r.rollupManager.Rollups[rollupID].TrustedSequencerBalances

	eth, err := c.BalanceAt(ctx, address, nil)
	if err != nil || eth == nil {
		r.logger.Error().Err(err).
			Any("address", address).
			Str("token", observer.ETH).
			Msg("Failed to get balance")
	} else {
		balances.ETH = eth
	}

	balances.POL = r.getPOL(c, address, co, balances.POL)
}

// rollupManagers maps L1 network names to their corresponding rollup manager
// addresses and labels. It is used to identify the rollup manager contracts
// deployed on different networks. This should not be modified during runtime.
var rollupManagers = map[string]map[common.Address]string{
	network.EthereumName: {
		common.HexToAddress("0x5132A183E9F3CB7C848b0AAC5Ae0c4f0491B7aB2"): "Mainnet",
	},
	network.SepoliaName: {
		common.HexToAddress("0x32d33D5137a7cFFb54c5Bf8371172bcEc5f310ff"): "Cardona",
		common.HexToAddress("0xE2EF6215aDc132Df6913C8DD16487aBF118d1764"): "Bali",
	},
}

// getRollupNetwork determines the appropriate network configuration for a given
// rollup ID. It prioritizes a network name override if available, then checks
// known rollup manager addresses, and finally constructs a default name if
// necessary.
func (r *RPCProvider) getRollupNetwork(contract *contracts.PolygonZkEVMEtrog, co *bind.CallOpts, rollupID uint32) network.Network {
	// Use the network name override if available.
	if rollup, ok := r.contracts.RollupManager.Rollups[rollupID]; ok && rollup.Name != nil {
		if n, err := network.GetNetworkByName(*rollup.Name); err == nil {
			return n
		}
	}

	address := common.HexToAddress(*r.contracts.RollupManagerAddress)

	name := fmt.Sprintf("%s %s Rollup %d",
		r.network.GetName(),
		address.Hex(),
		rollupID,
	)

	if addresses, ok := rollupManagers[r.network.GetName()]; ok {
		if rollupManagerName, ok := addresses[address]; ok {
			name = fmt.Sprintf("%s Rollup %d", rollupManagerName, rollupID)
		}
	}

	if networkName, err := contract.NetworkName(co); err != nil {
		log.Warn().Err(err).Msg("Failed to get rollup network name")
	} else if networkName != "" {
		name = fmt.Sprintf("%s %s", name, networkName)
	}

	return &config.Network{
		Name: name,
	}
}

// isRollupEnabled determines if a rollup with the given rollupID is enabled. By
// default, all rollups are enabled.
func (r *RPCProvider) isRollupEnabled(rollupID uint32) bool {
	if slices.Contains(r.contracts.RollupManager.Disabled, rollupID) {
		return false
	}

	if slices.Contains(r.contracts.RollupManager.Enabled, rollupID) {
		return true
	}

	return len(r.contracts.RollupManager.Enabled) == 0
}

func (r *RPCProvider) refreshTrustedSequencerURL(ctx context.Context, contract *contracts.PolygonZkEVMEtrog, co *bind.CallOpts, rollupID uint32) (err error) {
	if !r.isRollupEnabled(rollupID) {
		return nil
	}

	network := r.getRollupNetwork(contract, co, rollupID)
	if network == nil {
		return nil
	}

	url, err := contract.TrustedSequencerURL(co)
	if err != nil {
		return err
	}

	rollup, ok := r.contracts.RollupManager.Rollups[rollupID]
	if ok && rollup.URL != nil {
		rollup.RPC.URL = *rollup.URL
	} else {
		rollup.RPC.URL = url
	}

	rollup.RPC.Label = r.label
	if ok && rollup.Label != nil {
		rollup.RPC.Label = *rollup.Label
	}

	provider, ok := r.trustedSequencers[rollupID]
	if !ok {
		r.trustedSequencers[rollupID] = NewRPCProvider(network, r.bus, rollup.RPC)
		go runProvider(ctx, r.trustedSequencers[rollupID])
		return nil
	}

	if provider.url != url {
		provider.trustedSequencerURL <- url
	}

	return nil
}

func runProvider(ctx context.Context, p *RPCProvider) {
	for {
		select {
		case url := <-p.trustedSequencerURL:
			p.url = url
		default:
			if err := p.RefreshState(ctx); err != nil {
				p.logger.Error().Err(err).Send()
			}

			if err := p.PublishEvents(ctx); err != nil {
				p.logger.Error().Err(err).Send()
			}

			util.BlockFor(ctx, p.PollingInterval())
		}
	}
}

func (r *RPCProvider) refreshZkEVMEtrog(ctx context.Context, c *ethclient.Client, co *bind.CallOpts, rollupID uint32, rollup RollupData) error {
	contract, err := contracts.NewPolygonZkEVMEtrog(rollup.RollupContract, c)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to bind zkEVM Etrog contract")
		return nil
	}

	if _, ok := r.rollupManager.Rollups[rollupID]; !ok {
		r.rollupManager.Rollups[rollupID] = &observer.RollupData{}
	}

	r.rollupManager.Rollups[rollupID].ChainID = &rollup.ChainID

	lfb, err := contract.LastForceBatch(co)
	if err != nil {
		r.logger.Error().Err(err).Msg("Could not get last force batch")
	} else {
		r.rollupManager.Rollups[rollupID].LastForceBatch = &lfb
	}

	lfbs, err := contract.LastForceBatchSequenced(co)
	if err != nil {
		r.logger.Error().Err(err).Msg("Could not get last force batch sequenced")
	} else {
		r.rollupManager.Rollups[rollupID].LastForceBatchSequenced = &lfbs
	}

	r.refreshTrustedSequencerBalance(ctx, c, contract, co, rollupID)
	r.refreshTrustedSequencerURL(ctx, contract, co, rollupID)

	return nil
}

func (r *RPCProvider) refreshAggregatorBalances(ctx context.Context, c *ethclient.Client, aggregator common.Address) {
	eth, err := c.BalanceAt(ctx, aggregator, nil)
	if err != nil || eth == nil {
		r.logger.Error().Err(err).Any("address", aggregator).Msg("Failed to get aggregator balance")
		return
	}

	co := &bind.CallOpts{Context: ctx}
	balances := r.rollupManager.AggregatorBalances[aggregator]
	balances.ETH = eth
	balances.POL = r.getPOL(c, aggregator, co, balances.POL)
	r.rollupManager.AggregatorBalances[aggregator] = balances
}

func (r *RPCProvider) refreshRollupManager(ctx context.Context, c *ethclient.Client) error {
	if r.contracts.RollupManagerAddress == nil {
		return nil
	}

	if r.rollupManager == nil {
		r.rollupManager = &observer.RollupManager{
			Rollups:            make(map[uint32]*observer.RollupData),
			AggregatorBalances: make(map[common.Address]observer.TokenBalances),
		}
	}

	for _, rollup := range r.rollupManager.Rollups {
		rollup.TimeBetweenSequencedBatches = nil
		rollup.TimeBetweenVerifiedBatches = nil
		rollup.SequencedBatchesTxFees = nil
		rollup.VerifiedBatchesTxFees = nil
	}

	co := &bind.CallOpts{Context: ctx}
	address := common.HexToAddress(*r.contracts.RollupManagerAddress)
	contract, err := contracts.NewPolygonRollupManager(address, c)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to bind rollup manager contract")
		return nil
	}

	// Fetch the other contract addresses from the rollup manager contract.
	if err := r.refreshZkEVMContracts(contract, co); err != nil {
		return err
	}

	rpb, err := contract.CalculateRewardPerBatch(co)
	if err != nil {
		r.logger.Error().Err(err).Msg("Could not calculate reward per batch")
	} else {
		r.rollupManager.RewardPerBatch = rpb
	}

	lat, err := contract.LastAggregationTimestamp(co)
	if err != nil {
		r.logger.Error().Err(err).Msg("Could not get last aggregation timestamp")
	} else {
		r.rollupManager.LastAggregationTimestamp = &lat
	}

	opts := r.getFilterOpts()

	r.refreshBatchFees(contract, co)
	r.refreshBatchTotals(contract, co)
	r.refreshRollupCounts(contract, co)
	r.refreshRollups(ctx, c, contract, co)
	r.refreshOnSequenceBatches(ctx, c, contract, opts)
	r.refreshRollupVerifyBatches(ctx, c, contract, opts)
	r.refreshRollupVerifyBatchesTrustedAggregator(ctx, c, contract, opts)

	return nil
}

func (r *RPCProvider) refreshZkEVMContracts(contract *contracts.PolygonRollupManager, co *bind.CallOpts) error {
	if r.contracts.ZkEVMBridgeAddress == nil {
		bridgeAddress, err := contract.BridgeAddress(co)
		if err != nil {
			return err
		}

		r.contracts.ZkEVMBridgeAddress = new(string)
		*r.contracts.ZkEVMBridgeAddress = bridgeAddress.Hex()
	}

	if r.globalExitRootAddress == nil {
		germ, err := contract.GlobalExitRootManager(co)
		if err != nil {
			return nil
		}

		r.globalExitRootAddress = &germ
	}

	if r.polTokenAddress == nil {
		pol, err := contract.Pol(co)
		if err != nil {
			return err
		}

		r.polTokenAddress = &pol
	}

	return nil
}

// RollupData is the struct returned by the RollupIDToRollupData method. This is
// here because the abigen tool doesn't generate a named struct for this data.
// Update this value if the response ever changes.
type RollupData struct {
	RollupContract                 common.Address
	ChainID                        uint64
	Verifier                       common.Address
	ForkID                         uint64
	LastLocalExitRoot              [32]byte
	LastBatchSequenced             uint64
	LastVerifiedBatch              uint64
	LastPendingState               uint64
	LastPendingStateConsolidated   uint64
	LastVerifiedBatchBeforeUpgrade uint64
	RollupTypeID                   uint64
	RollupCompatibilityID          uint8
}

func (r *RPCProvider) refreshRollups(ctx context.Context, c *ethclient.Client, contract *contracts.PolygonRollupManager, co *bind.CallOpts) {
	if r.rollupManager.RollupCount == nil {
		return
	}

	for id := uint32(1); id <= *r.rollupManager.RollupCount; id++ {
		rollup, err := contract.RollupIDToRollupData(co, id)
		if err != nil {
			r.logger.Error().Err(err).Uint32("rollup", id).Msg("Failed to get rollup data")
			continue
		}

		r.refreshZkEVMEtrog(ctx, c, co, id, rollup)
	}
}

func (r *RPCProvider) refreshBatchFees(contract *contracts.PolygonRollupManager, co *bind.CallOpts) {
	bf, err := contract.GetBatchFee(co)
	if err != nil || bf == nil {
		r.logger.Error().Err(err).Msg("Could not get batch fee")
	} else {
		r.rollupManager.BatchFee = bf
	}

	fbf, err := contract.GetForcedBatchFee(co)
	if err != nil {
		r.logger.Error().Err(err).Msg("Could not get forced batch fee")
	} else {
		r.rollupManager.ForcedBatchFee = fbf
	}
}

func (r *RPCProvider) refreshRollupCounts(contract *contracts.PolygonRollupManager, co *bind.CallOpts) {
	rc, err := contract.RollupCount(co)
	if err != nil {
		r.logger.Error().Err(err).Msg("Could not get rollup count")
	} else {
		r.rollupManager.RollupCount = &rc
	}

	rtc, err := contract.RollupTypeCount(co)
	if err != nil {
		r.logger.Error().Err(err).Msg("Could not get rollup type count")
	} else {
		r.rollupManager.RollupTypeCount = &rtc
	}
}

func (r *RPCProvider) refreshBatchTotals(contract *contracts.PolygonRollupManager, co *bind.CallOpts) {
	tsb, err := contract.TotalSequencedBatches(co)
	if err != nil {
		r.logger.Error().Err(err).Msg("Could not get total sequenced batches")
	} else {
		r.rollupManager.TotalSequencedBatches = &tsb
	}

	tvb, err := contract.TotalVerifiedBatches(co)
	if err != nil {
		r.logger.Error().Err(err).Msg("Could not get total verified batches")
	} else {
		r.rollupManager.TotalVerifiedBatches = &tvb
	}
}

func (r *RPCProvider) refreshOnSequenceBatches(ctx context.Context, c *ethclient.Client, contract *contracts.PolygonRollupManager, opts *bind.FilterOpts) {
	iter, err := contract.FilterOnSequenceBatches(opts, nil)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to filter on sequence batches events")
		return
	}

	var event *contracts.PolygonRollupManagerOnSequenceBatches
	for iter.Next() && iter.Event != nil {
		event = iter.Event

		block, err := c.BlockByHash(ctx, event.Raw.BlockHash)
		if err != nil {
			r.logger.Error().Err(err).Msg("Failed to get block by hash")
			continue
		}

		time := block.Time()

		rollup, ok := r.rollupManager.Rollups[event.RollupID]
		if !ok {
			continue
		}

		if rollup.LastBatchSequenced != nil {
			if *rollup.LastBatchSequenced >= event.LastBatchSequenced {
				continue
			}

			rollup.TimeBetweenSequencedBatches = append(
				rollup.TimeBetweenSequencedBatches,
				time-*rollup.LastSequencedTimestamp,
			)
		}

		// Any logic after this point means that there is a new event that hasn't
		// been published to an observer yet.

		rollup.LastSequencedTimestamp = &time
		rollup.LastBatchSequenced = &event.LastBatchSequenced

		receipt, err := c.TransactionReceipt(ctx, event.Raw.TxHash)
		if err != nil {
			r.logger.Error().Err(err).Msg("Failed to get transaction receipt")
			continue
		}

		fee := new(big.Int).Mul(big.NewInt(int64(receipt.GasUsed)), receipt.EffectiveGasPrice)
		rollup.SequencedBatchesTxFees = append(rollup.SequencedBatchesTxFees, observer.RollupTx{
			Fee:     fee,
			Address: event.Raw.Address,
		})
	}
}

func (r *RPCProvider) refreshRollupVerifyBatches(ctx context.Context, c *ethclient.Client, contract *contracts.PolygonRollupManager, opts *bind.FilterOpts) {
	iter, err := contract.FilterVerifyBatches(opts, nil, nil)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to filter rollup verify batches events")
		return
	}

	var event *contracts.PolygonRollupManagerVerifyBatches
	for iter.Next() && iter.Event != nil {
		event = iter.Event

		r.refreshAggregatorBalances(ctx, c, event.Aggregator)

		block, err := c.BlockByHash(ctx, event.Raw.BlockHash)
		if err != nil {
			r.logger.Error().Err(err).Msg("Failed to get block by hash")
			continue
		}

		time := block.Time()

		rollup, ok := r.rollupManager.Rollups[event.RollupID]
		if !ok {
			continue
		}

		if rollup.LastVerifiedBatch != nil {
			if *rollup.LastVerifiedBatch >= event.NumBatch {
				continue
			}

			// There should not be an instance where a rollup has events for both the
			// VerifyBatches and VerifyBatchesTrustedAggregator, so the
			// TimeBetweenVerifiedBatches and VerifiedBatchesTxFees slices should not
			// be not be affected.
			rollup.TimeBetweenVerifiedBatches = append(
				rollup.TimeBetweenVerifiedBatches,
				time-*rollup.LastVerifiedTimestamp,
			)
		}

		// Any logic after this point means that there is a new event that hasn't
		// been published to an observer yet.

		rollup.LastVerifiedTimestamp = &time
		rollup.LastVerifiedBatch = &event.NumBatch

		receipt, err := c.TransactionReceipt(ctx, event.Raw.TxHash)
		if err != nil {
			r.logger.Error().Err(err).Msg("Failed to get transaction receipt")
			continue
		}

		fee := new(big.Int).Mul(big.NewInt(int64(receipt.GasUsed)), receipt.EffectiveGasPrice)
		rollup.VerifiedBatchesTxFees = append(rollup.VerifiedBatchesTxFees, observer.RollupTx{
			Fee:     fee,
			Address: event.Aggregator,
		})
	}
}

func (r *RPCProvider) refreshRollupVerifyBatchesTrustedAggregator(ctx context.Context, c *ethclient.Client, contract *contracts.PolygonRollupManager, opts *bind.FilterOpts) {
	iter, err := contract.FilterVerifyBatchesTrustedAggregator(opts, nil, nil)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to filter rollup verify batches trusted aggregator batches events")
		return
	}

	var event *contracts.PolygonRollupManagerVerifyBatchesTrustedAggregator
	for iter.Next() && iter.Event != nil {
		event = iter.Event

		r.refreshAggregatorBalances(ctx, c, event.Aggregator)

		block, err := c.BlockByHash(ctx, event.Raw.BlockHash)
		if err != nil {
			r.logger.Error().Err(err).Msg("Failed to get block by hash")
			continue
		}

		time := block.Time()

		rollup, ok := r.rollupManager.Rollups[event.RollupID]
		if !ok {
			continue
		}

		pessimistic := event.NumBatch == 0 && event.StateRoot == [32]byte{}

		if rollup.LastVerifiedBatch != nil {
			// Here, pessimistic chains are handled differently because the NumBatch
			// will always be 0. The last verified timestamp is used to determine if
			// the event has already been seen.
			//
			// There is an edge case here where if both events are included in the
			// same block, there will be a missing time between verified batches
			// event. This will be rare and won't significantly impact the metric.
			if pessimistic && *rollup.LastVerifiedTimestamp >= time {
				continue
			}

			if !pessimistic && *rollup.LastVerifiedBatch >= event.NumBatch {
				continue
			}

			// At this point rollup.LastVerifiedBatch and rollup.LastVerifiedTimestamp
			// contain the data of the previous verified batch event, not the current
			// one (which is stored in event). This is used to calculate the time
			// between the verified batches.
			//
			// This iterator traverses the logs in order, so even if there's more than
			// one verified batches event, the logic will hold.
			rollup.TimeBetweenVerifiedBatches = append(
				rollup.TimeBetweenVerifiedBatches,
				time-*rollup.LastVerifiedTimestamp,
			)
		}

		// Any logic after this point means that there is a new event that hasn't
		// been published to an observer yet.

		rollup.LastVerifiedTimestamp = &time
		rollup.LastVerifiedBatch = &event.NumBatch
		rollup.Pessimistic = pessimistic

		receipt, err := c.TransactionReceipt(ctx, event.Raw.TxHash)
		if err != nil {
			r.logger.Error().Err(err).Msg("Failed to get transaction receipt")
			continue
		}

		fee := new(big.Int).Mul(big.NewInt(int64(receipt.GasUsed)), receipt.EffectiveGasPrice)
		rollup.VerifiedBatchesTxFees = append(rollup.VerifiedBatchesTxFees, observer.RollupTx{
			Fee:     fee,
			Address: event.Aggregator,
		})
	}
}
