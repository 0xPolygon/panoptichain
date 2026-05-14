// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PolygonZkEVMGlobalExitRootV2MetaData contains all meta data concerning the PolygonZkEVMGlobalExitRootV2 contract.
var PolygonZkEVMGlobalExitRootV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"}],\"name\":\"UpdateL1InfoTree\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"calculateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastGlobalExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lastBlockHash\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"getLeafValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMainnetExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// PolygonZkEVMGlobalExitRootV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonZkEVMGlobalExitRootV2MetaData.ABI instead.
var PolygonZkEVMGlobalExitRootV2ABI = PolygonZkEVMGlobalExitRootV2MetaData.ABI

// PolygonZkEVMGlobalExitRootV2 is an auto generated Go binding around an Ethereum contract.
type PolygonZkEVMGlobalExitRootV2 struct {
	PolygonZkEVMGlobalExitRootV2Caller     // Read-only binding to the contract
	PolygonZkEVMGlobalExitRootV2Transactor // Write-only binding to the contract
	PolygonZkEVMGlobalExitRootV2Filterer   // Log filterer for contract events
}

// PolygonZkEVMGlobalExitRootV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonZkEVMGlobalExitRootV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonZkEVMGlobalExitRootV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonZkEVMGlobalExitRootV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonZkEVMGlobalExitRootV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonZkEVMGlobalExitRootV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonZkEVMGlobalExitRootV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonZkEVMGlobalExitRootV2Session struct {
	Contract     *PolygonZkEVMGlobalExitRootV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PolygonZkEVMGlobalExitRootV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonZkEVMGlobalExitRootV2CallerSession struct {
	Contract *PolygonZkEVMGlobalExitRootV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// PolygonZkEVMGlobalExitRootV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonZkEVMGlobalExitRootV2TransactorSession struct {
	Contract     *PolygonZkEVMGlobalExitRootV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// PolygonZkEVMGlobalExitRootV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonZkEVMGlobalExitRootV2Raw struct {
	Contract *PolygonZkEVMGlobalExitRootV2 // Generic contract binding to access the raw methods on
}

// PolygonZkEVMGlobalExitRootV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonZkEVMGlobalExitRootV2CallerRaw struct {
	Contract *PolygonZkEVMGlobalExitRootV2Caller // Generic read-only contract binding to access the raw methods on
}

// PolygonZkEVMGlobalExitRootV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonZkEVMGlobalExitRootV2TransactorRaw struct {
	Contract *PolygonZkEVMGlobalExitRootV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonZkEVMGlobalExitRootV2 creates a new instance of PolygonZkEVMGlobalExitRootV2, bound to a specific deployed contract.
func NewPolygonZkEVMGlobalExitRootV2(address common.Address, backend bind.ContractBackend) (*PolygonZkEVMGlobalExitRootV2, error) {
	contract, err := bindPolygonZkEVMGlobalExitRootV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMGlobalExitRootV2{PolygonZkEVMGlobalExitRootV2Caller: PolygonZkEVMGlobalExitRootV2Caller{contract: contract}, PolygonZkEVMGlobalExitRootV2Transactor: PolygonZkEVMGlobalExitRootV2Transactor{contract: contract}, PolygonZkEVMGlobalExitRootV2Filterer: PolygonZkEVMGlobalExitRootV2Filterer{contract: contract}}, nil
}

// NewPolygonZkEVMGlobalExitRootV2Caller creates a new read-only instance of PolygonZkEVMGlobalExitRootV2, bound to a specific deployed contract.
func NewPolygonZkEVMGlobalExitRootV2Caller(address common.Address, caller bind.ContractCaller) (*PolygonZkEVMGlobalExitRootV2Caller, error) {
	contract, err := bindPolygonZkEVMGlobalExitRootV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMGlobalExitRootV2Caller{contract: contract}, nil
}

// NewPolygonZkEVMGlobalExitRootV2Transactor creates a new write-only instance of PolygonZkEVMGlobalExitRootV2, bound to a specific deployed contract.
func NewPolygonZkEVMGlobalExitRootV2Transactor(address common.Address, transactor bind.ContractTransactor) (*PolygonZkEVMGlobalExitRootV2Transactor, error) {
	contract, err := bindPolygonZkEVMGlobalExitRootV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMGlobalExitRootV2Transactor{contract: contract}, nil
}

// NewPolygonZkEVMGlobalExitRootV2Filterer creates a new log filterer instance of PolygonZkEVMGlobalExitRootV2, bound to a specific deployed contract.
func NewPolygonZkEVMGlobalExitRootV2Filterer(address common.Address, filterer bind.ContractFilterer) (*PolygonZkEVMGlobalExitRootV2Filterer, error) {
	contract, err := bindPolygonZkEVMGlobalExitRootV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMGlobalExitRootV2Filterer{contract: contract}, nil
}

// bindPolygonZkEVMGlobalExitRootV2 binds a generic wrapper to an already deployed contract.
func bindPolygonZkEVMGlobalExitRootV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonZkEVMGlobalExitRootV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolygonZkEVMGlobalExitRootV2.Contract.PolygonZkEVMGlobalExitRootV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.PolygonZkEVMGlobalExitRootV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.PolygonZkEVMGlobalExitRootV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolygonZkEVMGlobalExitRootV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Caller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonZkEVMGlobalExitRootV2.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Session) BridgeAddress() (common.Address, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.BridgeAddress(&_PolygonZkEVMGlobalExitRootV2.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2CallerSession) BridgeAddress() (common.Address, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.BridgeAddress(&_PolygonZkEVMGlobalExitRootV2.CallOpts)
}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Caller) CalculateRoot(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	var out []interface{}
	err := _PolygonZkEVMGlobalExitRootV2.contract.Call(opts, &out, "calculateRoot", leafHash, smtProof, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Session) CalculateRoot(leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.CalculateRoot(&_PolygonZkEVMGlobalExitRootV2.CallOpts, leafHash, smtProof, index)
}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2CallerSession) CalculateRoot(leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.CalculateRoot(&_PolygonZkEVMGlobalExitRootV2.CallOpts, leafHash, smtProof, index)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Caller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolygonZkEVMGlobalExitRootV2.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Session) DepositCount() (*big.Int, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.DepositCount(&_PolygonZkEVMGlobalExitRootV2.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2CallerSession) DepositCount() (*big.Int, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.DepositCount(&_PolygonZkEVMGlobalExitRootV2.CallOpts)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Caller) GetLastGlobalExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolygonZkEVMGlobalExitRootV2.contract.Call(opts, &out, "getLastGlobalExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Session) GetLastGlobalExitRoot() ([32]byte, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.GetLastGlobalExitRoot(&_PolygonZkEVMGlobalExitRootV2.CallOpts)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2CallerSession) GetLastGlobalExitRoot() ([32]byte, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.GetLastGlobalExitRoot(&_PolygonZkEVMGlobalExitRootV2.CallOpts)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x5d810501.
//
// Solidity: function getLeafValue(bytes32 newGlobalExitRoot, uint256 lastBlockHash, uint64 timestamp) pure returns(bytes32)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Caller) GetLeafValue(opts *bind.CallOpts, newGlobalExitRoot [32]byte, lastBlockHash *big.Int, timestamp uint64) ([32]byte, error) {
	var out []interface{}
	err := _PolygonZkEVMGlobalExitRootV2.contract.Call(opts, &out, "getLeafValue", newGlobalExitRoot, lastBlockHash, timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLeafValue is a free data retrieval call binding the contract method 0x5d810501.
//
// Solidity: function getLeafValue(bytes32 newGlobalExitRoot, uint256 lastBlockHash, uint64 timestamp) pure returns(bytes32)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Session) GetLeafValue(newGlobalExitRoot [32]byte, lastBlockHash *big.Int, timestamp uint64) ([32]byte, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.GetLeafValue(&_PolygonZkEVMGlobalExitRootV2.CallOpts, newGlobalExitRoot, lastBlockHash, timestamp)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x5d810501.
//
// Solidity: function getLeafValue(bytes32 newGlobalExitRoot, uint256 lastBlockHash, uint64 timestamp) pure returns(bytes32)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2CallerSession) GetLeafValue(newGlobalExitRoot [32]byte, lastBlockHash *big.Int, timestamp uint64) ([32]byte, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.GetLeafValue(&_PolygonZkEVMGlobalExitRootV2.CallOpts, newGlobalExitRoot, lastBlockHash, timestamp)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Caller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolygonZkEVMGlobalExitRootV2.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Session) GetRoot() ([32]byte, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.GetRoot(&_PolygonZkEVMGlobalExitRootV2.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2CallerSession) GetRoot() ([32]byte, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.GetRoot(&_PolygonZkEVMGlobalExitRootV2.CallOpts)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Caller) GlobalExitRootMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PolygonZkEVMGlobalExitRootV2.contract.Call(opts, &out, "globalExitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Session) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.GlobalExitRootMap(&_PolygonZkEVMGlobalExitRootV2.CallOpts, arg0)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2CallerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.GlobalExitRootMap(&_PolygonZkEVMGlobalExitRootV2.CallOpts, arg0)
}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Caller) LastMainnetExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolygonZkEVMGlobalExitRootV2.contract.Call(opts, &out, "lastMainnetExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Session) LastMainnetExitRoot() ([32]byte, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.LastMainnetExitRoot(&_PolygonZkEVMGlobalExitRootV2.CallOpts)
}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2CallerSession) LastMainnetExitRoot() ([32]byte, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.LastMainnetExitRoot(&_PolygonZkEVMGlobalExitRootV2.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Caller) LastRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolygonZkEVMGlobalExitRootV2.contract.Call(opts, &out, "lastRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Session) LastRollupExitRoot() ([32]byte, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.LastRollupExitRoot(&_PolygonZkEVMGlobalExitRootV2.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2CallerSession) LastRollupExitRoot() ([32]byte, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.LastRollupExitRoot(&_PolygonZkEVMGlobalExitRootV2.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Caller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonZkEVMGlobalExitRootV2.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Session) RollupManager() (common.Address, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.RollupManager(&_PolygonZkEVMGlobalExitRootV2.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2CallerSession) RollupManager() (common.Address, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.RollupManager(&_PolygonZkEVMGlobalExitRootV2.CallOpts)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Caller) VerifyMerkleProof(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	var out []interface{}
	err := _PolygonZkEVMGlobalExitRootV2.contract.Call(opts, &out, "verifyMerkleProof", leafHash, smtProof, index, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Session) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.VerifyMerkleProof(&_PolygonZkEVMGlobalExitRootV2.CallOpts, leafHash, smtProof, index, root)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2CallerSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.VerifyMerkleProof(&_PolygonZkEVMGlobalExitRootV2.CallOpts, leafHash, smtProof, index, root)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Transactor) UpdateExitRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _PolygonZkEVMGlobalExitRootV2.contract.Transact(opts, "updateExitRoot", newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Session) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.UpdateExitRoot(&_PolygonZkEVMGlobalExitRootV2.TransactOpts, newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2TransactorSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _PolygonZkEVMGlobalExitRootV2.Contract.UpdateExitRoot(&_PolygonZkEVMGlobalExitRootV2.TransactOpts, newRoot)
}

// PolygonZkEVMGlobalExitRootV2UpdateL1InfoTreeIterator is returned from FilterUpdateL1InfoTree and is used to iterate over the raw logs and unpacked data for UpdateL1InfoTree events raised by the PolygonZkEVMGlobalExitRootV2 contract.
type PolygonZkEVMGlobalExitRootV2UpdateL1InfoTreeIterator struct {
	Event *PolygonZkEVMGlobalExitRootV2UpdateL1InfoTree // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonZkEVMGlobalExitRootV2UpdateL1InfoTreeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonZkEVMGlobalExitRootV2UpdateL1InfoTree)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonZkEVMGlobalExitRootV2UpdateL1InfoTree)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonZkEVMGlobalExitRootV2UpdateL1InfoTreeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonZkEVMGlobalExitRootV2UpdateL1InfoTreeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonZkEVMGlobalExitRootV2UpdateL1InfoTree represents a UpdateL1InfoTree event raised by the PolygonZkEVMGlobalExitRootV2 contract.
type PolygonZkEVMGlobalExitRootV2UpdateL1InfoTree struct {
	MainnetExitRoot [32]byte
	RollupExitRoot  [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateL1InfoTree is a free log retrieval operation binding the contract event 0xda61aa7823fcd807e37b95aabcbe17f03a6f3efd514176444dae191d27fd66b3.
//
// Solidity: event UpdateL1InfoTree(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Filterer) FilterUpdateL1InfoTree(opts *bind.FilterOpts, mainnetExitRoot [][32]byte, rollupExitRoot [][32]byte) (*PolygonZkEVMGlobalExitRootV2UpdateL1InfoTreeIterator, error) {

	var mainnetExitRootRule []interface{}
	for _, mainnetExitRootItem := range mainnetExitRoot {
		mainnetExitRootRule = append(mainnetExitRootRule, mainnetExitRootItem)
	}
	var rollupExitRootRule []interface{}
	for _, rollupExitRootItem := range rollupExitRoot {
		rollupExitRootRule = append(rollupExitRootRule, rollupExitRootItem)
	}

	logs, sub, err := _PolygonZkEVMGlobalExitRootV2.contract.FilterLogs(opts, "UpdateL1InfoTree", mainnetExitRootRule, rollupExitRootRule)
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMGlobalExitRootV2UpdateL1InfoTreeIterator{contract: _PolygonZkEVMGlobalExitRootV2.contract, event: "UpdateL1InfoTree", logs: logs, sub: sub}, nil
}

// WatchUpdateL1InfoTree is a free log subscription operation binding the contract event 0xda61aa7823fcd807e37b95aabcbe17f03a6f3efd514176444dae191d27fd66b3.
//
// Solidity: event UpdateL1InfoTree(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Filterer) WatchUpdateL1InfoTree(opts *bind.WatchOpts, sink chan<- *PolygonZkEVMGlobalExitRootV2UpdateL1InfoTree, mainnetExitRoot [][32]byte, rollupExitRoot [][32]byte) (event.Subscription, error) {

	var mainnetExitRootRule []interface{}
	for _, mainnetExitRootItem := range mainnetExitRoot {
		mainnetExitRootRule = append(mainnetExitRootRule, mainnetExitRootItem)
	}
	var rollupExitRootRule []interface{}
	for _, rollupExitRootItem := range rollupExitRoot {
		rollupExitRootRule = append(rollupExitRootRule, rollupExitRootItem)
	}

	logs, sub, err := _PolygonZkEVMGlobalExitRootV2.contract.WatchLogs(opts, "UpdateL1InfoTree", mainnetExitRootRule, rollupExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonZkEVMGlobalExitRootV2UpdateL1InfoTree)
				if err := _PolygonZkEVMGlobalExitRootV2.contract.UnpackLog(event, "UpdateL1InfoTree", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateL1InfoTree is a log parse operation binding the contract event 0xda61aa7823fcd807e37b95aabcbe17f03a6f3efd514176444dae191d27fd66b3.
//
// Solidity: event UpdateL1InfoTree(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot)
func (_PolygonZkEVMGlobalExitRootV2 *PolygonZkEVMGlobalExitRootV2Filterer) ParseUpdateL1InfoTree(log types.Log) (*PolygonZkEVMGlobalExitRootV2UpdateL1InfoTree, error) {
	event := new(PolygonZkEVMGlobalExitRootV2UpdateL1InfoTree)
	if err := _PolygonZkEVMGlobalExitRootV2.contract.UnpackLog(event, "UpdateL1InfoTree", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
