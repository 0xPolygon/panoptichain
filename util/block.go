package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/0xPolygon/panoptichain/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

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

// BlockByNumber gets the block given a block number.
func BlockByNumber(ctx context.Context, n *big.Int, c *ethclient.Client) (*types.Block, error) {
	var raw json.RawMessage
	err := c.Client().Call(&raw, "eth_getBlockByNumber", hexutil.EncodeBig(n), true)
	if err != nil {
		return nil, err
	}

	return getBlock(ctx, raw, c)
}

// BlockByHash gets the block given a block hash.
func BlockByHash(ctx context.Context, hash common.Hash, c *ethclient.Client) (*types.Block, error) {
	var raw json.RawMessage
	err := c.Client().Call(&raw, "eth_getBlockByHash", hash, true)
	if err != nil {
		return nil, err
	}

	return getBlock(ctx, raw, c)
}

// getBlock marshals a block given the json.RawMessage. This is lifted from
// https://github.com/ethereum/go-ethereum/blob/master/ethclient/ethclient.go
// with the change that unsupported transactions are treated as legacy
// transactions. This allows observation of chains that use transaction types
// that are not supported by Geth.
func getBlock(ctx context.Context, raw json.RawMessage, c *ethclient.Client) (*types.Block, error) {
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

	return types.NewBlockWithHeader(head).WithBody(
		types.Body{
			Transactions: txs,
			Uncles:       uncles,
			Withdrawals:  block.Withdrawals,
		},
	), nil
}
