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
	"net/url"
	"strconv"
	"strings"
	"time"

	zkevmtypes "github.com/0xPolygonHermez/zkevm-node/jsonrpc/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	URL              string
	Network          network.Network
	Label            string
	parsedURL        *url.URL
	bus              *observer.EventBus
	interval         uint
	logger           zerolog.Logger
	BlockNumber      uint64
	prevBlockNumber  uint64
	blockBuffer      *blockbuffer.BlockBuffer
	txPool           *observer.TransactionPool
	refreshStateTime *time.Duration
	contracts        config.ContractAddresses
	timeToMine       *config.TimeToMine
	accounts         []string
	accountBalances  observer.AccountBalances
	timeToFinalized  *uint64

	// PoS
	stateSync               map[bool]*observer.StateSync
	checkpointSignatures    map[bool]*observer.CheckpointSignatures
	validatorWalletBalances *observer.ValidatorWalletBalances
	missedBlockProposal     *observer.MissedBlockProposal

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

type RPCProviderOpts struct {
	Network    network.Network
	URL        string
	Label      string
	EventBus   *observer.EventBus
	Interval   uint
	Contracts  config.ContractAddresses
	TimeToMine *config.TimeToMine
	Accounts   []string
}

// NewRPCProvider will create a new RPC provider and configure it's event bus.
func NewRPCProvider(opts RPCProviderOpts) *RPCProvider {
	logger := NewLogger(opts.Network, opts.Label)

	parsedURL, err := url.Parse(opts.URL)
	if err != nil {
		logger.Error().Err(err).Msg("Unable to parse RPC URL")
	}

	return &RPCProvider{
		URL:                  opts.URL,
		Label:                opts.Label,
		parsedURL:            parsedURL,
		blockBuffer:          blockbuffer.NewBlockBuffer(128),
		Network:              opts.Network,
		bus:                  opts.EventBus,
		interval:             opts.Interval,
		logger:               logger,
		refreshStateTime:     new(time.Duration),
		contracts:            opts.Contracts,
		timeToMine:           opts.TimeToMine,
		accounts:             opts.Accounts,
		accountBalances:      make(observer.AccountBalances),
		stateSync:            make(map[bool]*observer.StateSync),
		checkpointSignatures: make(map[bool]*observer.CheckpointSignatures),
		bridgeEventTimes:     make(observer.BridgeEventTimes),
		claimEventTimes:      make(observer.ClaimEventTimes),
		trustedSequencers:    make(map[uint32]*RPCProvider),
		trustedSequencerURL:  make(chan string),
		rollupContracts:      make(map[uint32]common.Address),
	}
}

func (r *RPCProvider) SetEventBus(bus *observer.EventBus) {
	r.bus = bus
}

// RefreshState is going to get the current head block and request all
// of the blocks between the current head and the last head. All of
// those blocks will be pushed into the buffer.
func (r *RPCProvider) RefreshState(ctx context.Context) error {
	defer timer(r.refreshStateTime)()

	c, err := ethclient.DialContext(ctx, r.URL)
	if err != nil {
		r.logger.Error().Err(err).Msg("Unable to create the client")
		return err
	}

	r.refreshBlockBuffer(ctx, c)
	r.refreshStateSync(ctx, c, true)
	r.refreshStateSync(ctx, c, false)
	r.refreshCheckpoint(ctx, c)
	r.refreshWalletBalance(ctx, c)
	r.refreshMissedBlockProposal(ctx, c)
	r.refreshTxPoolStatus(ctx, c)
	r.refreshTimeToMine(ctx, c)
	r.refreshAccountBalances(ctx, c)

	r.refreshBatches(ctx, c)
	r.refreshRollupManager(ctx, c)
	r.refreshExitRoots(ctx, c)
	r.refreshExitRootsL2(ctx, c)
	r.refreshBridge(ctx, c)

	return nil
}

func (r *RPCProvider) PublishEvents(ctx context.Context) error {
	for i := r.prevBlockNumber + 1; i <= r.BlockNumber && r.prevBlockNumber != 0; i++ {
		b, err := r.blockBuffer.GetBlock(i)
		if err != nil {
			continue
		}
		block, ok := b.(*types.Block)
		if !ok {
			continue
		}

		m := observer.NewMessage(r.Network, r.Label, block)
		r.bus.Publish(ctx, topics.NewEVMBlock, m)

		pb, err := r.blockBuffer.GetBlock(b.Number().Uint64() - 1)
		if err != nil {
			continue
		}
		prev, ok := pb.(*types.Block)
		if !ok {
			continue
		}

		interval := observer.NewMessage(r.Network, r.Label, block.Time()-prev.Time())
		r.bus.Publish(ctx, topics.BlockInterval, interval)
	}

	if r.missedBlockProposal != nil {
		missedBlockProposal := observer.NewMessage(r.Network, r.Label, r.missedBlockProposal)
		r.bus.Publish(ctx, topics.BorMissedBlockProposal, missedBlockProposal)
	}

	for _, stateSync := range r.stateSync {
		r.bus.Publish(ctx, topics.BorStateSync, observer.NewMessage(r.Network, r.Label, stateSync))
	}

	for _, checkpointSignatures := range r.checkpointSignatures {
		m := observer.NewMessage(r.Network, r.Label, checkpointSignatures)
		r.bus.Publish(ctx, topics.CheckpointSignatures, m)
	}

	if r.validatorWalletBalances != nil {
		validatorWalletBalance := observer.NewMessage(r.Network, r.Label, r.validatorWalletBalances)
		r.bus.Publish(ctx, topics.ValidatorWallet, validatorWalletBalance)
	}

	if r.txPool != nil {
		transactionPoolStatus := observer.NewMessage(r.Network, r.Label, r.txPool)
		r.bus.Publish(ctx, topics.TransactionPool, transactionPoolStatus)
	}

	if r.batches.TrustedBatch.Number > 0 || r.batches.VirtualBatch.Number > 0 || r.batches.VerifiedBatch.Number > 0 {
		r.bus.Publish(ctx, topics.ZkEVMBatches, observer.NewMessage(r.Network, r.Label, r.batches))
	}

	if r.globalExitRoot != nil || r.mainnetExitRoot != nil || r.rollupExitRoot != nil {
		er := &observer.ExitRoots{
			GlobalExitRoot:  r.globalExitRoot,
			MainnetExitRoot: r.mainnetExitRoot,
			RollupExitRoot:  r.rollupExitRoot,
		}
		r.bus.Publish(ctx, topics.ExitRoots, observer.NewMessage(r.Network, r.Label, er))
	}

	if r.rollupExitRootL2 != nil {
		er := &observer.ExitRoots{
			RollupExitRoot: r.rollupExitRootL2,
		}
		r.bus.Publish(ctx, topics.ExitRoots, observer.NewMessage(r.Network, r.Label, er))
	}

	if r.depositCount != nil || r.lastUpdatedDepositCount != nil {
		m := observer.NewMessage(r.Network, r.Label, &observer.DepositCounts{
			DepositCount:            r.depositCount,
			LastUpdatedDepositCount: r.lastUpdatedDepositCount,
		})
		r.bus.Publish(ctx, topics.DepositCounts, m)
	}

	for _, bridgeEvent := range r.bridgeEvents {
		r.bus.Publish(ctx, topics.BridgeEvent, observer.NewMessage(r.Network, r.Label, bridgeEvent))
	}

	for _, claimEvent := range r.claimEvents {
		r.bus.Publish(ctx, topics.ClaimEvent, observer.NewMessage(r.Network, r.Label, claimEvent))
	}

	if len(r.bridgeEventTimes) > 0 {
		m := observer.NewMessage(r.Network, r.Label, r.bridgeEventTimes)
		r.bus.Publish(ctx, topics.BridgeEventTimes, m)
	}

	if len(r.claimEventTimes) > 0 {
		m := observer.NewMessage(r.Network, r.Label, r.claimEventTimes)
		r.bus.Publish(ctx, topics.ClaimEventTimes, m)
	}

	if r.rollupManager != nil {
		m := observer.NewMessage(r.Network, r.Label, r.rollupManager)
		r.bus.Publish(ctx, topics.RollupManager, m)
	}

	if len(r.accountBalances) > 0 {
		m := observer.NewMessage(r.Network, r.Label, r.accountBalances)
		r.bus.Publish(ctx, topics.AccountBalances, m)
	}

	for _, batch := range r.trustedBatches {
		m := observer.NewMessage(r.Network, r.Label, batch)
		r.bus.Publish(ctx, topics.TrustedBatch, m)
	}

	if r.timeToFinalized != nil {
		m := observer.NewMessage(r.Network, r.Label, r.timeToFinalized)
		r.bus.Publish(ctx, topics.TimeToFinalized, m)
	}

	r.bus.Publish(ctx, topics.RefreshStateTime, observer.NewMessage(r.Network, r.Label, r.refreshStateTime))

	return nil
}

func (r *RPCProvider) PollingInterval() uint {
	return r.interval
}

func (r *RPCProvider) refreshBlockBuffer(ctx context.Context, c *ethclient.Client) (err error) {
	r.prevBlockNumber = r.BlockNumber
	r.BlockNumber, err = c.BlockNumber(ctx)
	if err != nil {
		r.logger.Error().Err(err).Any("provider", r).Msg("Unable to get block number")
		return err
	}

	r.logger.Info().Uint64("block_number", r.BlockNumber).Msg("Block state refreshed")

	if r.prevBlockNumber == 0 {
		r.prevBlockNumber = r.BlockNumber
	}

	if r.prevBlockNumber != r.BlockNumber {
		r.fillRange(ctx, r.prevBlockNumber, c)
	}

	finalized, err := c.HeaderByNumber(ctx, big.NewInt(int64(rpc.FinalizedBlockNumber)))
	if err != nil {
		r.logger.Warn().Err(err).Msg("Failed to get finalized block header")
		return err
	}

	latest, err := c.HeaderByNumber(ctx, big.NewInt(int64(r.BlockNumber)))
	if err != nil {
		r.logger.Warn().Err(err).Msg("Failed to get latest block header")
		return err
	}

	diff := latest.Time - finalized.Time
	r.timeToFinalized = &diff

	return nil
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
			r.logger.Error().Err(err).Msg("Unable to bind state sender contract")
			return err
		}

		counter, err = ss.Counter(&co)
		if err != nil {
			r.logger.Error().Err(err).Msg("Unable to get counter")
			return err
		}
	} else if r.contracts.StateSyncReceiverAddress != nil {
		address := common.HexToAddress(*r.contracts.StateSyncReceiverAddress)
		sr, err := contracts.NewStateReceiver(address, c)
		if err != nil {
			r.logger.Error().Err(err).Msg("Unable to bind state receiver contract")
			return err
		}

		counter, err = sr.LastStateId(&co)
		if err != nil {
			r.logger.Error().Err(err).Msg("Unable to get counter")
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
		r.logger.Warn().Err(err).Msg("Unable to bind root chain contract")
		return
	}

	iter, err := contract.FilterNewHeaderBlock(&bind.FilterOpts{Start: r.prevBlockNumber}, nil, nil, nil)
	if iter == nil || err != nil {
		r.logger.Warn().Err(err).Msg("No NewHeaderBlock events found")
		return
	}

	// Get the last NewHeaderBlock event.
	var event *contracts.RootChainNewHeaderBlock
	for iter.Next() && iter.Event != nil {
		event = iter.Event
	}

	if event == nil {
		r.logger.Error().Msg("NewHeaderBlock event is nil")
		return
	}

	// Grab that block so that we know when the transaction was mined.
	block, err := c.BlockByHash(ctx, event.Raw.BlockHash)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to get block by hash")
		return
	}

	r.logger.Trace().Any("event", event).Str("network", r.Network.GetName()).Msg("Latest NewHeaderBlock event")
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

	inputs := make(map[string]interface{})
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

// fillRange will pull all of the blocks between the start and the current head.
func (r *RPCProvider) fillRange(ctx context.Context, start uint64, c *ethclient.Client) {
	r.logger.Debug().
		Uint64("start_block", start).
		Uint64("end_block", r.BlockNumber).
		Str("url", r.parsedURL.Host).
		Msg("Filling block range")

	for i := start + 1; i <= r.BlockNumber; i++ {
		b, err := c.BlockByNumber(ctx, new(big.Int).SetUint64(i))
		if err != nil {
			r.logger.Warn().Err(err).Uint64("block_number", i).Msg("Unable to get block")
			break
		}
		r.blockBuffer.PutBlock(b)
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

func (r *RPCProvider) refreshWalletBalance(ctx context.Context, c *ethclient.Client) (err error) {
	signers, err := api.Signers(r.Network)
	if err != nil {
		r.logger.Warn().Err(err).Msg("Failed to get signers validator map")
		return
	}

	reqs := make([]rpc.BatchElem, len(signers))
	signerAddresses := make([]string, 0, len(signers))

	index := 0
	for signerAddress := range signers {
		addr := common.HexToAddress(signerAddress)
		signerAddresses = append(signerAddresses, signerAddress)
		r := new(json.RawMessage)
		reqs[index] = rpc.BatchElem{
			Method: "eth_getBalance",
			Args:   []interface{}{addr, "latest"},
			Result: r,
			Error:  nil,
		}
		index++
	}

	err = c.Client().BatchCallContext(ctx, reqs)
	if err != nil {
		r.logger.Warn().Err(err).Msg("Failed to execute batch request for balances")
		return err
	}

	balances := make(observer.ValidatorWalletBalances)
	for i, req := range reqs {
		logger := r.logger.Warn().Int("index", i)

		if req.Error != nil {
			logger.Err(req.Error).Msg("Failed to get balance for validator")
			continue
		}

		var balanceStr string
		if err := json.Unmarshal(*req.Result.(*json.RawMessage), &balanceStr); err != nil {
			logger.Err(err).Msg("Failed to unmarshal balance for validator")
			continue
		}

		balance, ok := new(big.Int).SetString(balanceStr[2:], 16)
		if !ok {
			logger.Msg("Invalid value for validator")
			continue
		}

		signerAddress := signerAddresses[i]
		balances[signerAddress] = balance
	}

	r.validatorWalletBalances = &balances

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
	missedBlockProposal := make(observer.MissedBlockProposal)
	for i := r.prevBlockNumber + 1; i <= r.BlockNumber && r.prevBlockNumber != 0; i++ {
		var response SnapshotProposerSequence
		blockNumberHex := fmt.Sprintf("0x%x", i)
		err := c.Client().CallContext(ctx, &response, "bor_getSnapshotProposerSequence", blockNumberHex)
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
		actualSigner := "0x" + hex.EncodeToString(bytes)

		if actualSigner != response.Author {
			for _, signerInfo := range response.Signers {
				if actualSigner == signerInfo.Signer {
					break
				}
				missedBlockProposal[i] = append(missedBlockProposal[i], signerInfo.Signer)
			}
		}
	}
	r.missedBlockProposal = &missedBlockProposal

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

		m := observer.NewMessage(r.Network, r.Label, ttm)
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

// refreshExitRoot will update the exit root. If it has not seen, the exit root
// it will set the Time to t. If it has already been observed, the time will
// remain the same and the Seen value will be set to true.
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
		r.logger.Error().Err(err).Msg("Unable to bind global exit root contract")
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
		r.logger.Error().Err(err).Msg("Failed to block hash from global exit root map")
	} else {
		header, err := c.HeaderByHash(ctx, common.BigToHash(hash))
		if err != nil || header == nil {
			r.logger.Error().Err(err).Msg("Failed to block header from global exit root block hash")
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
		r.logger.Error().Err(err).Msg("Unable to bind global exit root l2 contract")
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
		r.logger.Error().Err(err).Msg("Unable to bind zkEVM bridge contract")
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

	if r.prevBlockNumber == 0 {
		return nil
	}

	opts := &bind.FilterOpts{Start: r.prevBlockNumber}
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
		r.logger.Error().Err(err).Msg("Unable to bind ERC20 contract")
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

	if address.Cmp(common.Address{}) == 0 {
		r.logger.Warn().Msg("Invalid trusted sequencer address")
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

func (r *RPCProvider) newRollupNetwork(rollupID uint32) network.Network {
	var name string

	switch r.Network.GetName() {
	case network.EthereumName:
		name = "Mainnet"
	case network.SepoliaName:
		if strings.HasPrefix(r.Label, "cardona.") {
			name = "Cardona"
		} else if strings.HasPrefix(r.Label, "bali.") {
			name = "Bali"
		}
	default:
		r.logger.Error().
			Str("network", r.Network.GetName()).
			Msg("Failed to create new rollup network")

		return nil
	}

	if len(name) == 0 {
		r.logger.Error().Msg("New rollup network has no name")
		return nil
	}

	return &network.EVMNetwork{
		Name: fmt.Sprintf("%s Rollup %d", name, rollupID),
	}
}

func (r *RPCProvider) refreshTrustedSequencerURL(ctx context.Context, contract *contracts.PolygonZkEVMEtrog, co *bind.CallOpts, rollupID uint32) error {
	url, err := contract.TrustedSequencerURL(co)
	if err != nil {
		return err
	}

	if url == "https://decomm.invalid" {
		return nil
	}

	rollupNetwork := r.newRollupNetwork(rollupID)
	if rollupNetwork == nil {
		return nil
	}

	provider, ok := r.trustedSequencers[rollupID]
	if !ok {
		r.trustedSequencers[rollupID] = NewRPCProvider(RPCProviderOpts{
			Network:  rollupNetwork,
			URL:      url,
			Label:    url,
			EventBus: r.bus,
			Interval: r.interval,
		})
		go runProvider(ctx, r.trustedSequencers[rollupID])
		return nil
	}

	if provider.URL != url {
		provider.trustedSequencerURL <- url
	}

	return nil
}

func runProvider(ctx context.Context, p *RPCProvider) {
	for {
		select {
		case url := <-p.trustedSequencerURL:
			p.URL = url
		default:
			if err := p.RefreshState(ctx); err != nil {
				p.logger.Error().Err(err).Send()
			}

			if err := p.PublishEvents(ctx); err != nil {
				p.logger.Error().Err(err).Send()
			}

			util.BlockFor(ctx, time.Second*time.Duration(p.PollingInterval()))
		}
	}
}

func (r *RPCProvider) refreshZkEVMEtrog(ctx context.Context, c *ethclient.Client, address common.Address, co *bind.CallOpts, rollupID uint32, chainID uint64) error {
	contract, err := contracts.NewPolygonZkEVMEtrog(address, c)
	if err != nil {
		r.logger.Error().Err(err).Msg("Unable to bind zkEVM Etrog contract")
		return nil
	}

	if _, ok := r.rollupManager.Rollups[rollupID]; !ok {
		r.rollupManager.Rollups[rollupID] = &observer.RollupData{}
	}

	r.rollupManager.Rollups[rollupID].ChainID = &chainID

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
		r.logger.Error().Err(err).Msg("Unable to bind rollup manager contract")
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

	r.refreshBatchFees(contract, co)
	r.refreshBatchTotals(contract, co)
	r.refreshRollupCounts(contract, co)
	r.refreshRollups(ctx, c, contract, co)

	if r.prevBlockNumber == 0 {
		return nil
	}

	opts := &bind.FilterOpts{Start: r.prevBlockNumber}
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

		r.refreshZkEVMEtrog(ctx, c, rollup.RollupContract, co, id, rollup.ChainID)
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

		if rollup.LastVerifiedBatch != nil {
			if *rollup.LastVerifiedBatch >= event.NumBatch {
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
