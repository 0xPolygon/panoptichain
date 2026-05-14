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

// PolygonZkEVMGlobalExitRootL2MetaData contains all meta data concerning the PolygonZkEVMGlobalExitRootL2 contract.
var PolygonZkEVMGlobalExitRootL2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PolygonZkEVMGlobalExitRootL2ABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonZkEVMGlobalExitRootL2MetaData.ABI instead.
var PolygonZkEVMGlobalExitRootL2ABI = PolygonZkEVMGlobalExitRootL2MetaData.ABI

// PolygonZkEVMGlobalExitRootL2 is an auto generated Go binding around an Ethereum contract.
type PolygonZkEVMGlobalExitRootL2 struct {
	PolygonZkEVMGlobalExitRootL2Caller     // Read-only binding to the contract
	PolygonZkEVMGlobalExitRootL2Transactor // Write-only binding to the contract
	PolygonZkEVMGlobalExitRootL2Filterer   // Log filterer for contract events
}

// PolygonZkEVMGlobalExitRootL2Caller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonZkEVMGlobalExitRootL2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonZkEVMGlobalExitRootL2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonZkEVMGlobalExitRootL2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonZkEVMGlobalExitRootL2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonZkEVMGlobalExitRootL2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonZkEVMGlobalExitRootL2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonZkEVMGlobalExitRootL2Session struct {
	Contract     *PolygonZkEVMGlobalExitRootL2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PolygonZkEVMGlobalExitRootL2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonZkEVMGlobalExitRootL2CallerSession struct {
	Contract *PolygonZkEVMGlobalExitRootL2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// PolygonZkEVMGlobalExitRootL2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonZkEVMGlobalExitRootL2TransactorSession struct {
	Contract     *PolygonZkEVMGlobalExitRootL2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// PolygonZkEVMGlobalExitRootL2Raw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonZkEVMGlobalExitRootL2Raw struct {
	Contract *PolygonZkEVMGlobalExitRootL2 // Generic contract binding to access the raw methods on
}

// PolygonZkEVMGlobalExitRootL2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonZkEVMGlobalExitRootL2CallerRaw struct {
	Contract *PolygonZkEVMGlobalExitRootL2Caller // Generic read-only contract binding to access the raw methods on
}

// PolygonZkEVMGlobalExitRootL2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonZkEVMGlobalExitRootL2TransactorRaw struct {
	Contract *PolygonZkEVMGlobalExitRootL2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonZkEVMGlobalExitRootL2 creates a new instance of PolygonZkEVMGlobalExitRootL2, bound to a specific deployed contract.
func NewPolygonZkEVMGlobalExitRootL2(address common.Address, backend bind.ContractBackend) (*PolygonZkEVMGlobalExitRootL2, error) {
	contract, err := bindPolygonZkEVMGlobalExitRootL2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMGlobalExitRootL2{PolygonZkEVMGlobalExitRootL2Caller: PolygonZkEVMGlobalExitRootL2Caller{contract: contract}, PolygonZkEVMGlobalExitRootL2Transactor: PolygonZkEVMGlobalExitRootL2Transactor{contract: contract}, PolygonZkEVMGlobalExitRootL2Filterer: PolygonZkEVMGlobalExitRootL2Filterer{contract: contract}}, nil
}

// NewPolygonZkEVMGlobalExitRootL2Caller creates a new read-only instance of PolygonZkEVMGlobalExitRootL2, bound to a specific deployed contract.
func NewPolygonZkEVMGlobalExitRootL2Caller(address common.Address, caller bind.ContractCaller) (*PolygonZkEVMGlobalExitRootL2Caller, error) {
	contract, err := bindPolygonZkEVMGlobalExitRootL2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMGlobalExitRootL2Caller{contract: contract}, nil
}

// NewPolygonZkEVMGlobalExitRootL2Transactor creates a new write-only instance of PolygonZkEVMGlobalExitRootL2, bound to a specific deployed contract.
func NewPolygonZkEVMGlobalExitRootL2Transactor(address common.Address, transactor bind.ContractTransactor) (*PolygonZkEVMGlobalExitRootL2Transactor, error) {
	contract, err := bindPolygonZkEVMGlobalExitRootL2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMGlobalExitRootL2Transactor{contract: contract}, nil
}

// NewPolygonZkEVMGlobalExitRootL2Filterer creates a new log filterer instance of PolygonZkEVMGlobalExitRootL2, bound to a specific deployed contract.
func NewPolygonZkEVMGlobalExitRootL2Filterer(address common.Address, filterer bind.ContractFilterer) (*PolygonZkEVMGlobalExitRootL2Filterer, error) {
	contract, err := bindPolygonZkEVMGlobalExitRootL2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMGlobalExitRootL2Filterer{contract: contract}, nil
}

// bindPolygonZkEVMGlobalExitRootL2 binds a generic wrapper to an already deployed contract.
func bindPolygonZkEVMGlobalExitRootL2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonZkEVMGlobalExitRootL2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolygonZkEVMGlobalExitRootL2 *PolygonZkEVMGlobalExitRootL2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolygonZkEVMGlobalExitRootL2.Contract.PolygonZkEVMGlobalExitRootL2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolygonZkEVMGlobalExitRootL2 *PolygonZkEVMGlobalExitRootL2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonZkEVMGlobalExitRootL2.Contract.PolygonZkEVMGlobalExitRootL2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolygonZkEVMGlobalExitRootL2 *PolygonZkEVMGlobalExitRootL2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolygonZkEVMGlobalExitRootL2.Contract.PolygonZkEVMGlobalExitRootL2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolygonZkEVMGlobalExitRootL2 *PolygonZkEVMGlobalExitRootL2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolygonZkEVMGlobalExitRootL2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolygonZkEVMGlobalExitRootL2 *PolygonZkEVMGlobalExitRootL2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonZkEVMGlobalExitRootL2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolygonZkEVMGlobalExitRootL2 *PolygonZkEVMGlobalExitRootL2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolygonZkEVMGlobalExitRootL2.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_PolygonZkEVMGlobalExitRootL2 *PolygonZkEVMGlobalExitRootL2Caller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonZkEVMGlobalExitRootL2.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_PolygonZkEVMGlobalExitRootL2 *PolygonZkEVMGlobalExitRootL2Session) BridgeAddress() (common.Address, error) {
	return _PolygonZkEVMGlobalExitRootL2.Contract.BridgeAddress(&_PolygonZkEVMGlobalExitRootL2.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_PolygonZkEVMGlobalExitRootL2 *PolygonZkEVMGlobalExitRootL2CallerSession) BridgeAddress() (common.Address, error) {
	return _PolygonZkEVMGlobalExitRootL2.Contract.BridgeAddress(&_PolygonZkEVMGlobalExitRootL2.CallOpts)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_PolygonZkEVMGlobalExitRootL2 *PolygonZkEVMGlobalExitRootL2Caller) GlobalExitRootMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PolygonZkEVMGlobalExitRootL2.contract.Call(opts, &out, "globalExitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_PolygonZkEVMGlobalExitRootL2 *PolygonZkEVMGlobalExitRootL2Session) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _PolygonZkEVMGlobalExitRootL2.Contract.GlobalExitRootMap(&_PolygonZkEVMGlobalExitRootL2.CallOpts, arg0)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_PolygonZkEVMGlobalExitRootL2 *PolygonZkEVMGlobalExitRootL2CallerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _PolygonZkEVMGlobalExitRootL2.Contract.GlobalExitRootMap(&_PolygonZkEVMGlobalExitRootL2.CallOpts, arg0)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_PolygonZkEVMGlobalExitRootL2 *PolygonZkEVMGlobalExitRootL2Caller) LastRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolygonZkEVMGlobalExitRootL2.contract.Call(opts, &out, "lastRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_PolygonZkEVMGlobalExitRootL2 *PolygonZkEVMGlobalExitRootL2Session) LastRollupExitRoot() ([32]byte, error) {
	return _PolygonZkEVMGlobalExitRootL2.Contract.LastRollupExitRoot(&_PolygonZkEVMGlobalExitRootL2.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_PolygonZkEVMGlobalExitRootL2 *PolygonZkEVMGlobalExitRootL2CallerSession) LastRollupExitRoot() ([32]byte, error) {
	return _PolygonZkEVMGlobalExitRootL2.Contract.LastRollupExitRoot(&_PolygonZkEVMGlobalExitRootL2.CallOpts)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_PolygonZkEVMGlobalExitRootL2 *PolygonZkEVMGlobalExitRootL2Transactor) UpdateExitRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _PolygonZkEVMGlobalExitRootL2.contract.Transact(opts, "updateExitRoot", newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_PolygonZkEVMGlobalExitRootL2 *PolygonZkEVMGlobalExitRootL2Session) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _PolygonZkEVMGlobalExitRootL2.Contract.UpdateExitRoot(&_PolygonZkEVMGlobalExitRootL2.TransactOpts, newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_PolygonZkEVMGlobalExitRootL2 *PolygonZkEVMGlobalExitRootL2TransactorSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _PolygonZkEVMGlobalExitRootL2.Contract.UpdateExitRoot(&_PolygonZkEVMGlobalExitRootL2.TransactOpts, newRoot)
}
