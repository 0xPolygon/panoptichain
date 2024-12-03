package provider

import (
	"container/list"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/maticnetwork/polygon-cli/p2p/database"
	"github.com/rs/zerolog"
	"google.golang.org/api/iterator"

	"github.com/0xPolygon/panoptichain/api"
	"github.com/0xPolygon/panoptichain/network"
	"github.com/0xPolygon/panoptichain/observer"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

type SensorNetworkProvider struct {
	Network  network.Network
	Label    string
	bus      *observer.EventBus
	interval uint
	db       *datastore.Client
	logger   zerolog.Logger

	// Because blockbuffer stores blocks with the block number as the key, this
	// wouldn't be usable for the sensor network.
	blocks          *list.List
	BlockNumber     uint64
	prevBlockNumber uint64

	blockEvents     []*observer.SensorBlockEvents
	blockEventsLock sync.Mutex

	stolenBlocks []*types.Block

	// Store any new reorgs that are observed in this slice.
	reorgs          []*observer.DatastoreReorg
	latestReorgTime time.Time

	refreshStateTime *time.Duration
}

func NewSensorNetworkProvider(ctx context.Context, n network.Network, project, database, label string, eb *observer.EventBus, interval uint) *SensorNetworkProvider {
	logger := NewLogger(n, label)

	db, err := datastore.NewClientWithDatabase(ctx, project, database)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to connect to Datastore")
	}

	return &SensorNetworkProvider{
		Network:          n,
		Label:            label,
		bus:              eb,
		interval:         interval,
		db:               db,
		logger:           logger,
		blocks:           list.New(),
		latestReorgTime:  time.Now(),
		refreshStateTime: new(time.Duration),
	}
}

func (s *SensorNetworkProvider) RefreshState(ctx context.Context) error {
	defer timer(s.refreshStateTime)()

	if err := s.refreshBlockBuffer(ctx); err != nil {
		s.logger.Error().Err(err).Msg("Failed to refresh block buffer")
	}

	if err := s.refreshReorgs(ctx); err != nil {
		s.logger.Error().Err(err).Msg("Failed to refresh reorg")
	}

	if err := s.refreshStolenBlocks(); err != nil {
		s.logger.Error().Err(err).Msg("Failed to refresh stolen blocks")
	}

	return nil
}

func (s *SensorNetworkProvider) PublishEvents(ctx context.Context) error {
	blocks := make(types.Blocks, 0, s.blocks.Len())

	for e := s.blocks.Front(); e != nil; e = e.Next() {
		block := e.Value.(*types.Block)

		// Skip blocks that have already been published.
		if block.Number().Uint64() < s.prevBlockNumber {
			continue
		}

		blocks = append(blocks, block)

		m := observer.NewMessage(s.Network, s.Label, block)
		s.bus.Publish(ctx, topics.NewEVMBlock, m)
	}

	if len(blocks) > 0 {
		m := observer.NewMessage(s.Network, s.Label, &observer.SensorBlocks{
			Start:  s.prevBlockNumber,
			End:    s.BlockNumber,
			Blocks: blocks,
		})

		s.bus.Publish(ctx, topics.SensorBlocks, m)
	}

	for _, event := range s.blockEvents {
		m := observer.NewMessage(s.Network, s.Label, event)
		s.bus.Publish(ctx, topics.SensorBlockEvents, m)
	}

	for _, reorg := range s.reorgs {
		m := observer.NewMessage(s.Network, s.Label, reorg)
		s.bus.Publish(ctx, topics.Reorg, m)
	}

	for _, stolenBlock := range s.stolenBlocks {
		m := observer.NewMessage(s.Network, s.Label, stolenBlock)
		s.bus.Publish(ctx, topics.StolenBlock, m)
	}

	s.bus.Publish(ctx, topics.RefreshStateTime, observer.NewMessage(s.Network, s.Label, s.refreshStateTime))

	return nil
}

func (s *SensorNetworkProvider) SetEventBus(bus *observer.EventBus) {
	s.bus = bus
}

func (s *SensorNetworkProvider) PollingInterval() uint {
	return s.interval
}

func (s *SensorNetworkProvider) refreshBlockBuffer(ctx context.Context) error {
	s.blockEvents = nil

	if s.BlockNumber > s.prevBlockNumber {
		s.prevBlockNumber = s.BlockNumber
	}

	query := datastore.NewQuery(database.BlocksKind).Order("-TimeFirstSeen").Limit(1)
	var block database.DatastoreBlock

	if _, err := s.db.Run(ctx, query).Next(&block); err != nil {
		return err
	}

	bn, err := strconv.ParseUint(block.Number, 10, 0)
	if err != nil {
		return err
	}
	s.BlockNumber = bn

	if s.prevBlockNumber != 0 && s.prevBlockNumber != s.BlockNumber {
		s.fillRange(ctx, s.prevBlockNumber)
	}

	s.logger.Trace().
		Uint64("block_number", s.BlockNumber).
		Msg("Sensor network block state refreshed")

	return nil
}

func (s *SensorNetworkProvider) fillRange(ctx context.Context, start uint64) {
	s.logger.Debug().
		Uint64("start", start).
		Uint64("end", s.BlockNumber).
		Msg("Filling sensor network block range")

	// This query is slightly different from the other ones found in rpc.go and
	// heimdall.go. Here, the end block number is exclusive because the sensors
	// may still be writing multiple blocks to Datastore. This gives the sensors
	// around a two second buffer to write blocks before they are read by this
	// provider.
	query := datastore.NewQuery(database.BlocksKind).
		Order("Number").
		FilterField("Number", ">=", fmt.Sprint(start)).
		FilterField("Number", "<", fmt.Sprint(s.BlockNumber))
	iter := s.db.Run(ctx, query)

	var wg sync.WaitGroup

	for {
		var b database.DatastoreBlock

		key, err := iter.Next(&b)
		if err == iterator.Done {
			break
		}
		if err != nil {
			s.logger.Warn().Err(err).Msg("Failed to get next block")
			continue
		}

		go s.getBlockEvents(ctx, key, &b, &wg)

		block, err := NewBlockFromDatastoreBlock(&b)
		if err != nil {
			s.logger.Warn().Err(err).Msg("Failed to convert block")
			continue
		}

		// Only keep a certain amount of blocks in the buffer. Remove the oldest
		// block if it is full.
		if s.blocks.Len() >= 512 {
			s.blocks.Remove(s.blocks.Front())
		}

		s.blocks.PushBack(block)
	}

	wg.Wait()
}

func (s *SensorNetworkProvider) getBlockEvents(ctx context.Context, key *datastore.Key, block *database.DatastoreBlock, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	query := datastore.NewQuery(database.BlockEventsKind).FilterField("Hash", "=", key)
	var events []database.DatastoreEvent

	if _, err := s.db.GetAll(ctx, query, &events); err != nil {
		s.logger.Error().Err(err).Msg("Failed to get block events")
		return
	}

	s.blockEventsLock.Lock()
	defer s.blockEventsLock.Unlock()

	s.blockEvents = append(s.blockEvents, &observer.SensorBlockEvents{
		Block:  block,
		Events: events,
	})
}

func NewBlockFromDatastoreBlock(b *database.DatastoreBlock) (*types.Block, error) {
	difficulty, ok := new(big.Int).SetString(b.Difficulty, 10)
	if !ok {
		return nil, errors.New("failed to parse difficulty")
	}

	number, ok := new(big.Int).SetString(b.Number, 10)
	if !ok {
		return nil, errors.New("failed to parse block number")
	}

	baseFee, ok := new(big.Int).SetString(b.BaseFee, 10)
	if !ok {
		return nil, errors.New("failed to parse base fee")
	}

	gasLimit, err := strconv.ParseUint(b.GasLimit, 10, 0)
	if err != nil {
		return nil, err
	}

	gasUsed, err := strconv.ParseUint(b.GasUsed, 10, 0)
	if err != nil {
		return nil, err
	}

	nonce, err := strconv.ParseUint(b.Nonce, 10, 0)
	if err != nil {
		return nil, err
	}

	header := &types.Header{
		ParentHash:  common.HexToHash(b.ParentHash.Name),
		UncleHash:   common.HexToHash(b.UncleHash),
		Coinbase:    common.HexToAddress(b.Coinbase),
		Root:        common.HexToHash(b.Root),
		TxHash:      common.HexToHash(b.TxHash),
		ReceiptHash: common.HexToHash(b.ReceiptHash),
		Bloom:       types.BytesToBloom(b.Bloom),
		Difficulty:  difficulty,
		Number:      number,
		GasLimit:    gasLimit,
		GasUsed:     gasUsed,
		Time:        uint64(b.Time.Unix()),
		Extra:       b.Extra,
		MixDigest:   common.HexToHash(b.MixDigest),
		Nonce:       types.EncodeNonce(nonce),
		BaseFee:     baseFee,
	}

	return types.NewBlockWithHeader(header), nil
}

func (s *SensorNetworkProvider) refreshReorgs(ctx context.Context) error {
	query := datastore.NewQuery(observer.ReorgsKind).
		Order("Time").
		FilterField("Time", ">", s.latestReorgTime)

	var reorgs []*observer.DatastoreReorg
	if _, err := s.db.GetAll(ctx, query, &reorgs); err != nil {
		return err
	}

	s.reorgs = reorgs
	if len(reorgs) > 0 {
		s.latestReorgTime = *reorgs[len(reorgs)-1].Time
	}

	return nil
}

func (s *SensorNetworkProvider) getNonBogonBlocks() []*types.Block {
	signers, err := api.Signers(s.Network)
	if err != nil {
		s.logger.Warn().Err(err).Msg("Failed to get signers validator map")
		return nil
	}

	var blocks []*types.Block
	for e := s.blocks.Front(); e != nil; e = e.Next() {
		block := e.Value.(*types.Block)

		bytes, err := api.Ecrecover(block.Header())
		if err != nil {
			s.logger.Warn().Err(err).Msg("Failed to get block signer")
			continue
		}
		signer := "0x" + hex.EncodeToString(bytes)

		// Filter out bogon blocks.
		if _, ok := signers[signer]; !ok {
			continue
		}

		blocks = append(blocks, block)
	}

	return blocks
}

func (s *SensorNetworkProvider) refreshStolenBlocks() error {
	s.stolenBlocks = nil

	blocks := s.getNonBogonBlocks()

	for _, reorg := range s.reorgs {
		reorgBlocks := make([]*types.Block, 0, reorg.Depth)

		// Find all the blocks in the reorg.
		for hash := reorg.EndBlock.Name; hash != reorg.StartBlock.Name; {
			block := findBlockWithHash(blocks, hash)
			if block == nil {
				break
			}

			reorgBlocks = append(reorgBlocks, block)
			hash = block.ParentHash().Hex()
		}

	loop:
		for _, reorgBlock := range reorgBlocks {
			blocksWithNumber := findBlocksWithNumber(blocks, reorgBlock.NumberU64())

			for _, block := range blocksWithNumber {
				if block.Difficulty().Cmp(reorgBlock.Difficulty()) == 1 {
					continue loop
				}
			}

			s.stolenBlocks = append(s.stolenBlocks, reorgBlock)
		}
	}

	return nil
}

func findBlockWithHash(blocks []*types.Block, hash string) *types.Block {
	for _, block := range blocks {
		if block.Hash().Hex() == hash {
			return block
		}
	}

	return nil
}

func findBlocksWithNumber(blocks []*types.Block, number uint64) []*types.Block {
	var blocksWithNumber []*types.Block

	for _, block := range blocks {
		if block.NumberU64() == number {
			blocksWithNumber = append(blocksWithNumber, block)
		}
	}

	return blocksWithNumber
}
