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

// ValidatorShareMetaData contains all meta data concerning the ValidatorShare contract.
var ValidatorShareMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLiquidRewards\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"exchangeRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"}]",
}

// ValidatorShareABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorShareMetaData.ABI instead.
var ValidatorShareABI = ValidatorShareMetaData.ABI

// ValidatorShare is an auto generated Go binding around an Ethereum contract.
type ValidatorShare struct {
	ValidatorShareCaller     // Read-only binding to the contract
	ValidatorShareTransactor // Write-only binding to the contract
	ValidatorShareFilterer   // Log filterer for contract events
}

// ValidatorShareCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorShareCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorShareTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorShareTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorShareFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorShareFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorShareSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorShareSession struct {
	Contract     *ValidatorShare   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorShareCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorShareCallerSession struct {
	Contract *ValidatorShareCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ValidatorShareTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorShareTransactorSession struct {
	Contract     *ValidatorShareTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ValidatorShareRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorShareRaw struct {
	Contract *ValidatorShare // Generic contract binding to access the raw methods on
}

// ValidatorShareCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorShareCallerRaw struct {
	Contract *ValidatorShareCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorShareTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorShareTransactorRaw struct {
	Contract *ValidatorShareTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorShare creates a new instance of ValidatorShare, bound to a specific deployed contract.
func NewValidatorShare(address common.Address, backend bind.ContractBackend) (*ValidatorShare, error) {
	contract, err := bindValidatorShare(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorShare{ValidatorShareCaller: ValidatorShareCaller{contract: contract}, ValidatorShareTransactor: ValidatorShareTransactor{contract: contract}, ValidatorShareFilterer: ValidatorShareFilterer{contract: contract}}, nil
}

// NewValidatorShareCaller creates a new read-only instance of ValidatorShare, bound to a specific deployed contract.
func NewValidatorShareCaller(address common.Address, caller bind.ContractCaller) (*ValidatorShareCaller, error) {
	contract, err := bindValidatorShare(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorShareCaller{contract: contract}, nil
}

// NewValidatorShareTransactor creates a new write-only instance of ValidatorShare, bound to a specific deployed contract.
func NewValidatorShareTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorShareTransactor, error) {
	contract, err := bindValidatorShare(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorShareTransactor{contract: contract}, nil
}

// NewValidatorShareFilterer creates a new log filterer instance of ValidatorShare, bound to a specific deployed contract.
func NewValidatorShareFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorShareFilterer, error) {
	contract, err := bindValidatorShare(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorShareFilterer{contract: contract}, nil
}

// bindValidatorShare binds a generic wrapper to an already deployed contract.
func bindValidatorShare(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorShareMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorShare *ValidatorShareRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorShare.Contract.ValidatorShareCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorShare *ValidatorShareRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorShare.Contract.ValidatorShareTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorShare *ValidatorShareRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorShare.Contract.ValidatorShareTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorShare *ValidatorShareCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorShare.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorShare *ValidatorShareTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorShare.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorShare *ValidatorShareTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorShare.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ValidatorShare *ValidatorShareCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorShare.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ValidatorShare *ValidatorShareSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ValidatorShare.Contract.BalanceOf(&_ValidatorShare.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ValidatorShare *ValidatorShareCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ValidatorShare.Contract.BalanceOf(&_ValidatorShare.CallOpts, account)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(bool)
func (_ValidatorShare *ValidatorShareCaller) Delegation(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ValidatorShare.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(bool)
func (_ValidatorShare *ValidatorShareSession) Delegation() (bool, error) {
	return _ValidatorShare.Contract.Delegation(&_ValidatorShare.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(bool)
func (_ValidatorShare *ValidatorShareCallerSession) Delegation() (bool, error) {
	return _ValidatorShare.Contract.Delegation(&_ValidatorShare.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_ValidatorShare *ValidatorShareCaller) ExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorShare.contract.Call(opts, &out, "exchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_ValidatorShare *ValidatorShareSession) ExchangeRate() (*big.Int, error) {
	return _ValidatorShare.Contract.ExchangeRate(&_ValidatorShare.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_ValidatorShare *ValidatorShareCallerSession) ExchangeRate() (*big.Int, error) {
	return _ValidatorShare.Contract.ExchangeRate(&_ValidatorShare.CallOpts)
}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address user) view returns(uint256)
func (_ValidatorShare *ValidatorShareCaller) GetLiquidRewards(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorShare.contract.Call(opts, &out, "getLiquidRewards", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address user) view returns(uint256)
func (_ValidatorShare *ValidatorShareSession) GetLiquidRewards(user common.Address) (*big.Int, error) {
	return _ValidatorShare.Contract.GetLiquidRewards(&_ValidatorShare.CallOpts, user)
}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address user) view returns(uint256)
func (_ValidatorShare *ValidatorShareCallerSession) GetLiquidRewards(user common.Address) (*big.Int, error) {
	return _ValidatorShare.Contract.GetLiquidRewards(&_ValidatorShare.CallOpts, user)
}
