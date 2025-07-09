package util

import (
	"encoding/hex"
	"encoding/json"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// Base represents the hexadecimal base, which is 16
	Base = 16

	// BitSize64 64 bits
	BitSize64 = 64
)

// ArgUint64 helps to marshal uint64 values provided in the RPC requests
type ArgUint64 uint64

// MarshalText marshals into text
func (b ArgUint64) MarshalText() ([]byte, error) {
	buf := make([]byte, 2) //nolint:gomnd
	copy(buf, `0x`)
	buf = strconv.AppendUint(buf, uint64(b), Base)
	return buf, nil
}

// UnmarshalText unmarshals from text
func (b *ArgUint64) UnmarshalText(input []byte) error {
	str := strings.TrimPrefix(string(input), "0x")
	num, err := strconv.ParseUint(str, Base, BitSize64)
	if err != nil {
		return err
	}
	*b = ArgUint64(num)
	return nil
}

// ArgBig helps to marshal big number values provided in the RPC requests
type ArgBig big.Int

// UnmarshalText unmarshals an instance of ArgBig into an array of bytes
func (a *ArgBig) UnmarshalText(input []byte) error {
	buf, err := decodeToHex(input)
	if err != nil {
		return err
	}

	b := new(big.Int)
	b.SetBytes(buf)
	*a = ArgBig(*b)

	return nil
}

// ArgBytes helps to marshal byte array values provided in the RPC requests
type ArgBytes []byte

// MarshalText marshals into text
func (b ArgBytes) MarshalText() ([]byte, error) {
	return encodeToHex(b), nil
}

// UnmarshalText unmarshals from text
func (b *ArgBytes) UnmarshalText(input []byte) error {
	hh, err := decodeToHex(input)
	if err != nil {
		return nil
	}
	aux := make([]byte, len(hh))
	copy(aux, hh)
	*b = aux
	return nil
}

// Hex returns a hexadecimal representation
func (b ArgBytes) Hex() string {
	bb, _ := b.MarshalText()
	return string(bb)
}

// ArgBytesPtr helps to marshal byte array values provided in the RPC requests
func ArgBytesPtr(b []byte) *ArgBytes {
	bb := ArgBytes(b)

	return &bb
}

func decodeToHex(b []byte) ([]byte, error) {
	str := string(b)
	str = strings.TrimPrefix(str, "0x")
	if len(str)%2 != 0 {
		str = "0" + str
	}
	return hex.DecodeString(str)
}

func encodeToHex(b []byte) []byte {
	str := hex.EncodeToString(b)
	if len(str)%2 != 0 {
		str = "0" + str
	}
	return []byte("0x" + str)
}

// MarshalText marshals an array of bytes into an instance of ArgBig
func (a ArgBig) MarshalText() ([]byte, error) {
	b := (*big.Int)(&a)

	return []byte("0x" + b.Text(Base)), nil
}

type Batch struct {
	Number              ArgUint64           `json:"number"`
	Coinbase            common.Address      `json:"coinbase"`
	StateRoot           common.Hash         `json:"stateRoot"`
	GlobalExitRoot      common.Hash         `json:"globalExitRoot"`
	MainnetExitRoot     common.Hash         `json:"mainnetExitRoot"`
	RollupExitRoot      common.Hash         `json:"rollupExitRoot"`
	LocalExitRoot       common.Hash         `json:"localExitRoot"`
	AccInputHash        common.Hash         `json:"accInputHash"`
	Timestamp           ArgUint64           `json:"timestamp"`
	SendSequencesTxHash *common.Hash        `json:"sendSequencesTxHash"`
	VerifyBatchTxHash   *common.Hash        `json:"verifyBatchTxHash"`
	Transactions        []TransactionOrHash `json:"transactions"`
}

// TransactionOrHash for union type of transaction and types.Hash
type TransactionOrHash struct {
	Hash *common.Hash
	Tx   *Transaction
}

// MarshalJSON marshals into json
func (b TransactionOrHash) MarshalJSON() ([]byte, error) {
	if b.Hash != nil {
		return json.Marshal(b.Hash)
	}
	return json.Marshal(b.Tx)
}

// UnmarshalJSON unmarshals from json
func (b *TransactionOrHash) UnmarshalJSON(input []byte) error {
	v := string(input)
	if strings.HasPrefix(v, "0x") || strings.HasPrefix(v, "\"0x") {
		var h common.Hash
		err := json.Unmarshal(input, &h)
		if err != nil {
			return err
		}
		*b = TransactionOrHash{Hash: &h}
		return nil
	}

	var t Transaction
	err := json.Unmarshal(input, &t)
	if err != nil {
		return err
	}
	*b = TransactionOrHash{Tx: &t}
	return nil
}

// Transaction structure
type Transaction struct {
	Nonce       ArgUint64       `json:"nonce"`
	GasPrice    ArgBig          `json:"gasPrice"`
	Gas         ArgUint64       `json:"gas"`
	To          *common.Address `json:"to"`
	Value       ArgBig          `json:"value"`
	Input       ArgBytes        `json:"input"`
	V           ArgBig          `json:"v"`
	R           ArgBig          `json:"r"`
	S           ArgBig          `json:"s"`
	Hash        common.Hash     `json:"hash"`
	From        common.Address  `json:"from"`
	BlockHash   *common.Hash    `json:"blockHash"`
	BlockNumber *ArgUint64      `json:"blockNumber"`
	TxIndex     *ArgUint64      `json:"transactionIndex"`
	ChainID     ArgBig          `json:"chainId"`
	Type        ArgUint64       `json:"type"`
}
