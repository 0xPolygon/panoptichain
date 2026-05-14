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

// BorValidatorSetValidator is an auto generated low-level Go binding around an user-defined struct.
type BorValidatorSetValidator struct {
	Id     *big.Int
	Power  *big.Int
	Signer common.Address
}

// ValidatorSetMetaData contains all meta data concerning the ValidatorSet contract.
var ValidatorSetMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"SPRINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SYSTEM_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CHAIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FIRST_END_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"producers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUND_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BOR_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spanNumbers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VOTE_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spans\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"name\":\"NewSpan\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSprint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"span\",\"type\":\"uint256\"}],\"name\":\"getSpan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentSpan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNextSpan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"getSpanByBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSpanNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"span\",\"type\":\"uint256\"}],\"name\":\"getValidatorsTotalStakeBySpan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"span\",\"type\":\"uint256\"}],\"name\":\"getProducersTotalStakeBySpan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"span\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getValidatorBySigner\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"internalType\":\"structBorValidatorSet.Validator\",\"name\":\"result\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"span\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"span\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"isProducer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"isCurrentValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"isCurrentProducer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"getBorValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newSpan\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"validatorBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"producerBytes\",\"type\":\"bytes\"}],\"name\":\"commitSpan\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"span\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sigs\",\"type\":\"bytes\"}],\"name\":\"getStakePowerBySigs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"checkMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"d\",\"type\":\"bytes32\"}],\"name\":\"leafNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"right\",\"type\":\"bytes32\"}],\"name\":\"innerNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// ValidatorSetABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorSetMetaData.ABI instead.
var ValidatorSetABI = ValidatorSetMetaData.ABI

// ValidatorSet is an auto generated Go binding around an Ethereum contract.
type ValidatorSet struct {
	ValidatorSetCaller     // Read-only binding to the contract
	ValidatorSetTransactor // Write-only binding to the contract
	ValidatorSetFilterer   // Log filterer for contract events
}

// ValidatorSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSetSession struct {
	Contract     *ValidatorSet     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorSetCallerSession struct {
	Contract *ValidatorSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ValidatorSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorSetTransactorSession struct {
	Contract     *ValidatorSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ValidatorSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorSetRaw struct {
	Contract *ValidatorSet // Generic contract binding to access the raw methods on
}

// ValidatorSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorSetCallerRaw struct {
	Contract *ValidatorSetCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorSetTransactorRaw struct {
	Contract *ValidatorSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorSet creates a new instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSet(address common.Address, backend bind.ContractBackend) (*ValidatorSet, error) {
	contract, err := bindValidatorSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorSet{ValidatorSetCaller: ValidatorSetCaller{contract: contract}, ValidatorSetTransactor: ValidatorSetTransactor{contract: contract}, ValidatorSetFilterer: ValidatorSetFilterer{contract: contract}}, nil
}

// NewValidatorSetCaller creates a new read-only instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSetCaller(address common.Address, caller bind.ContractCaller) (*ValidatorSetCaller, error) {
	contract, err := bindValidatorSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetCaller{contract: contract}, nil
}

// NewValidatorSetTransactor creates a new write-only instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSetTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorSetTransactor, error) {
	contract, err := bindValidatorSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetTransactor{contract: contract}, nil
}

// NewValidatorSetFilterer creates a new log filterer instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSetFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorSetFilterer, error) {
	contract, err := bindValidatorSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetFilterer{contract: contract}, nil
}

// bindValidatorSet binds a generic wrapper to an already deployed contract.
func bindValidatorSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorSetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorSet *ValidatorSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorSet.Contract.ValidatorSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorSet *ValidatorSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSet.Contract.ValidatorSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorSet *ValidatorSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorSet.Contract.ValidatorSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorSet *ValidatorSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorSet *ValidatorSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorSet *ValidatorSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorSet.Contract.contract.Transact(opts, method, params...)
}

// BORID is a free data retrieval call binding the contract method 0xae756451.
//
// Solidity: function BOR_ID() view returns(bytes32)
func (_ValidatorSet *ValidatorSetCaller) BORID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "BOR_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BORID is a free data retrieval call binding the contract method 0xae756451.
//
// Solidity: function BOR_ID() view returns(bytes32)
func (_ValidatorSet *ValidatorSetSession) BORID() ([32]byte, error) {
	return _ValidatorSet.Contract.BORID(&_ValidatorSet.CallOpts)
}

// BORID is a free data retrieval call binding the contract method 0xae756451.
//
// Solidity: function BOR_ID() view returns(bytes32)
func (_ValidatorSet *ValidatorSetCallerSession) BORID() ([32]byte, error) {
	return _ValidatorSet.Contract.BORID(&_ValidatorSet.CallOpts)
}

// CHAIN is a free data retrieval call binding the contract method 0x43ee8213.
//
// Solidity: function CHAIN() view returns(bytes32)
func (_ValidatorSet *ValidatorSetCaller) CHAIN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "CHAIN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CHAIN is a free data retrieval call binding the contract method 0x43ee8213.
//
// Solidity: function CHAIN() view returns(bytes32)
func (_ValidatorSet *ValidatorSetSession) CHAIN() ([32]byte, error) {
	return _ValidatorSet.Contract.CHAIN(&_ValidatorSet.CallOpts)
}

// CHAIN is a free data retrieval call binding the contract method 0x43ee8213.
//
// Solidity: function CHAIN() view returns(bytes32)
func (_ValidatorSet *ValidatorSetCallerSession) CHAIN() ([32]byte, error) {
	return _ValidatorSet.Contract.CHAIN(&_ValidatorSet.CallOpts)
}

// FIRSTENDBLOCK is a free data retrieval call binding the contract method 0x66332354.
//
// Solidity: function FIRST_END_BLOCK() view returns(uint256)
func (_ValidatorSet *ValidatorSetCaller) FIRSTENDBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "FIRST_END_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FIRSTENDBLOCK is a free data retrieval call binding the contract method 0x66332354.
//
// Solidity: function FIRST_END_BLOCK() view returns(uint256)
func (_ValidatorSet *ValidatorSetSession) FIRSTENDBLOCK() (*big.Int, error) {
	return _ValidatorSet.Contract.FIRSTENDBLOCK(&_ValidatorSet.CallOpts)
}

// FIRSTENDBLOCK is a free data retrieval call binding the contract method 0x66332354.
//
// Solidity: function FIRST_END_BLOCK() view returns(uint256)
func (_ValidatorSet *ValidatorSetCallerSession) FIRSTENDBLOCK() (*big.Int, error) {
	return _ValidatorSet.Contract.FIRSTENDBLOCK(&_ValidatorSet.CallOpts)
}

// ROUNDTYPE is a free data retrieval call binding the contract method 0x98ab2b62.
//
// Solidity: function ROUND_TYPE() view returns(bytes32)
func (_ValidatorSet *ValidatorSetCaller) ROUNDTYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "ROUND_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROUNDTYPE is a free data retrieval call binding the contract method 0x98ab2b62.
//
// Solidity: function ROUND_TYPE() view returns(bytes32)
func (_ValidatorSet *ValidatorSetSession) ROUNDTYPE() ([32]byte, error) {
	return _ValidatorSet.Contract.ROUNDTYPE(&_ValidatorSet.CallOpts)
}

// ROUNDTYPE is a free data retrieval call binding the contract method 0x98ab2b62.
//
// Solidity: function ROUND_TYPE() view returns(bytes32)
func (_ValidatorSet *ValidatorSetCallerSession) ROUNDTYPE() ([32]byte, error) {
	return _ValidatorSet.Contract.ROUNDTYPE(&_ValidatorSet.CallOpts)
}

// SPRINT is a free data retrieval call binding the contract method 0x2bc06564.
//
// Solidity: function SPRINT() view returns(uint256)
func (_ValidatorSet *ValidatorSetCaller) SPRINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "SPRINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SPRINT is a free data retrieval call binding the contract method 0x2bc06564.
//
// Solidity: function SPRINT() view returns(uint256)
func (_ValidatorSet *ValidatorSetSession) SPRINT() (*big.Int, error) {
	return _ValidatorSet.Contract.SPRINT(&_ValidatorSet.CallOpts)
}

// SPRINT is a free data retrieval call binding the contract method 0x2bc06564.
//
// Solidity: function SPRINT() view returns(uint256)
func (_ValidatorSet *ValidatorSetCallerSession) SPRINT() (*big.Int, error) {
	return _ValidatorSet.Contract.SPRINT(&_ValidatorSet.CallOpts)
}

// SYSTEMADDRESS is a free data retrieval call binding the contract method 0x3434735f.
//
// Solidity: function SYSTEM_ADDRESS() view returns(address)
func (_ValidatorSet *ValidatorSetCaller) SYSTEMADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "SYSTEM_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEMADDRESS is a free data retrieval call binding the contract method 0x3434735f.
//
// Solidity: function SYSTEM_ADDRESS() view returns(address)
func (_ValidatorSet *ValidatorSetSession) SYSTEMADDRESS() (common.Address, error) {
	return _ValidatorSet.Contract.SYSTEMADDRESS(&_ValidatorSet.CallOpts)
}

// SYSTEMADDRESS is a free data retrieval call binding the contract method 0x3434735f.
//
// Solidity: function SYSTEM_ADDRESS() view returns(address)
func (_ValidatorSet *ValidatorSetCallerSession) SYSTEMADDRESS() (common.Address, error) {
	return _ValidatorSet.Contract.SYSTEMADDRESS(&_ValidatorSet.CallOpts)
}

// VOTETYPE is a free data retrieval call binding the contract method 0xd5b844eb.
//
// Solidity: function VOTE_TYPE() view returns(uint8)
func (_ValidatorSet *ValidatorSetCaller) VOTETYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "VOTE_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTETYPE is a free data retrieval call binding the contract method 0xd5b844eb.
//
// Solidity: function VOTE_TYPE() view returns(uint8)
func (_ValidatorSet *ValidatorSetSession) VOTETYPE() (uint8, error) {
	return _ValidatorSet.Contract.VOTETYPE(&_ValidatorSet.CallOpts)
}

// VOTETYPE is a free data retrieval call binding the contract method 0xd5b844eb.
//
// Solidity: function VOTE_TYPE() view returns(uint8)
func (_ValidatorSet *ValidatorSetCallerSession) VOTETYPE() (uint8, error) {
	return _ValidatorSet.Contract.VOTETYPE(&_ValidatorSet.CallOpts)
}

// CheckMembership is a free data retrieval call binding the contract method 0x35ddfeea.
//
// Solidity: function checkMembership(bytes32 rootHash, bytes32 leaf, bytes proof) pure returns(bool)
func (_ValidatorSet *ValidatorSetCaller) CheckMembership(opts *bind.CallOpts, rootHash [32]byte, leaf [32]byte, proof []byte) (bool, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "checkMembership", rootHash, leaf, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMembership is a free data retrieval call binding the contract method 0x35ddfeea.
//
// Solidity: function checkMembership(bytes32 rootHash, bytes32 leaf, bytes proof) pure returns(bool)
func (_ValidatorSet *ValidatorSetSession) CheckMembership(rootHash [32]byte, leaf [32]byte, proof []byte) (bool, error) {
	return _ValidatorSet.Contract.CheckMembership(&_ValidatorSet.CallOpts, rootHash, leaf, proof)
}

// CheckMembership is a free data retrieval call binding the contract method 0x35ddfeea.
//
// Solidity: function checkMembership(bytes32 rootHash, bytes32 leaf, bytes proof) pure returns(bool)
func (_ValidatorSet *ValidatorSetCallerSession) CheckMembership(rootHash [32]byte, leaf [32]byte, proof []byte) (bool, error) {
	return _ValidatorSet.Contract.CheckMembership(&_ValidatorSet.CallOpts, rootHash, leaf, proof)
}

// CurrentSpanNumber is a free data retrieval call binding the contract method 0x4dbc959f.
//
// Solidity: function currentSpanNumber() view returns(uint256)
func (_ValidatorSet *ValidatorSetCaller) CurrentSpanNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "currentSpanNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentSpanNumber is a free data retrieval call binding the contract method 0x4dbc959f.
//
// Solidity: function currentSpanNumber() view returns(uint256)
func (_ValidatorSet *ValidatorSetSession) CurrentSpanNumber() (*big.Int, error) {
	return _ValidatorSet.Contract.CurrentSpanNumber(&_ValidatorSet.CallOpts)
}

// CurrentSpanNumber is a free data retrieval call binding the contract method 0x4dbc959f.
//
// Solidity: function currentSpanNumber() view returns(uint256)
func (_ValidatorSet *ValidatorSetCallerSession) CurrentSpanNumber() (*big.Int, error) {
	return _ValidatorSet.Contract.CurrentSpanNumber(&_ValidatorSet.CallOpts)
}

// CurrentSprint is a free data retrieval call binding the contract method 0xe3b7c924.
//
// Solidity: function currentSprint() view returns(uint256)
func (_ValidatorSet *ValidatorSetCaller) CurrentSprint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "currentSprint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentSprint is a free data retrieval call binding the contract method 0xe3b7c924.
//
// Solidity: function currentSprint() view returns(uint256)
func (_ValidatorSet *ValidatorSetSession) CurrentSprint() (*big.Int, error) {
	return _ValidatorSet.Contract.CurrentSprint(&_ValidatorSet.CallOpts)
}

// CurrentSprint is a free data retrieval call binding the contract method 0xe3b7c924.
//
// Solidity: function currentSprint() view returns(uint256)
func (_ValidatorSet *ValidatorSetCallerSession) CurrentSprint() (*big.Int, error) {
	return _ValidatorSet.Contract.CurrentSprint(&_ValidatorSet.CallOpts)
}

// GetBorValidators is a free data retrieval call binding the contract method 0x0c35b1cb.
//
// Solidity: function getBorValidators(uint256 number) view returns(address[], uint256[])
func (_ValidatorSet *ValidatorSetCaller) GetBorValidators(opts *bind.CallOpts, number *big.Int) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "getBorValidators", number)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetBorValidators is a free data retrieval call binding the contract method 0x0c35b1cb.
//
// Solidity: function getBorValidators(uint256 number) view returns(address[], uint256[])
func (_ValidatorSet *ValidatorSetSession) GetBorValidators(number *big.Int) ([]common.Address, []*big.Int, error) {
	return _ValidatorSet.Contract.GetBorValidators(&_ValidatorSet.CallOpts, number)
}

// GetBorValidators is a free data retrieval call binding the contract method 0x0c35b1cb.
//
// Solidity: function getBorValidators(uint256 number) view returns(address[], uint256[])
func (_ValidatorSet *ValidatorSetCallerSession) GetBorValidators(number *big.Int) ([]common.Address, []*big.Int, error) {
	return _ValidatorSet.Contract.GetBorValidators(&_ValidatorSet.CallOpts, number)
}

// GetCurrentSpan is a free data retrieval call binding the contract method 0xaf26aa96.
//
// Solidity: function getCurrentSpan() view returns(uint256 number, uint256 startBlock, uint256 endBlock)
func (_ValidatorSet *ValidatorSetCaller) GetCurrentSpan(opts *bind.CallOpts) (struct {
	Number     *big.Int
	StartBlock *big.Int
	EndBlock   *big.Int
}, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "getCurrentSpan")

	outstruct := new(struct {
		Number     *big.Int
		StartBlock *big.Int
		EndBlock   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Number = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCurrentSpan is a free data retrieval call binding the contract method 0xaf26aa96.
//
// Solidity: function getCurrentSpan() view returns(uint256 number, uint256 startBlock, uint256 endBlock)
func (_ValidatorSet *ValidatorSetSession) GetCurrentSpan() (struct {
	Number     *big.Int
	StartBlock *big.Int
	EndBlock   *big.Int
}, error) {
	return _ValidatorSet.Contract.GetCurrentSpan(&_ValidatorSet.CallOpts)
}

// GetCurrentSpan is a free data retrieval call binding the contract method 0xaf26aa96.
//
// Solidity: function getCurrentSpan() view returns(uint256 number, uint256 startBlock, uint256 endBlock)
func (_ValidatorSet *ValidatorSetCallerSession) GetCurrentSpan() (struct {
	Number     *big.Int
	StartBlock *big.Int
	EndBlock   *big.Int
}, error) {
	return _ValidatorSet.Contract.GetCurrentSpan(&_ValidatorSet.CallOpts)
}

// GetInitialValidators is a free data retrieval call binding the contract method 0x65b3a1e2.
//
// Solidity: function getInitialValidators() view returns(address[], uint256[])
func (_ValidatorSet *ValidatorSetCaller) GetInitialValidators(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "getInitialValidators")

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetInitialValidators is a free data retrieval call binding the contract method 0x65b3a1e2.
//
// Solidity: function getInitialValidators() view returns(address[], uint256[])
func (_ValidatorSet *ValidatorSetSession) GetInitialValidators() ([]common.Address, []*big.Int, error) {
	return _ValidatorSet.Contract.GetInitialValidators(&_ValidatorSet.CallOpts)
}

// GetInitialValidators is a free data retrieval call binding the contract method 0x65b3a1e2.
//
// Solidity: function getInitialValidators() view returns(address[], uint256[])
func (_ValidatorSet *ValidatorSetCallerSession) GetInitialValidators() ([]common.Address, []*big.Int, error) {
	return _ValidatorSet.Contract.GetInitialValidators(&_ValidatorSet.CallOpts)
}

// GetNextSpan is a free data retrieval call binding the contract method 0x60c8614d.
//
// Solidity: function getNextSpan() view returns(uint256 number, uint256 startBlock, uint256 endBlock)
func (_ValidatorSet *ValidatorSetCaller) GetNextSpan(opts *bind.CallOpts) (struct {
	Number     *big.Int
	StartBlock *big.Int
	EndBlock   *big.Int
}, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "getNextSpan")

	outstruct := new(struct {
		Number     *big.Int
		StartBlock *big.Int
		EndBlock   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Number = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNextSpan is a free data retrieval call binding the contract method 0x60c8614d.
//
// Solidity: function getNextSpan() view returns(uint256 number, uint256 startBlock, uint256 endBlock)
func (_ValidatorSet *ValidatorSetSession) GetNextSpan() (struct {
	Number     *big.Int
	StartBlock *big.Int
	EndBlock   *big.Int
}, error) {
	return _ValidatorSet.Contract.GetNextSpan(&_ValidatorSet.CallOpts)
}

// GetNextSpan is a free data retrieval call binding the contract method 0x60c8614d.
//
// Solidity: function getNextSpan() view returns(uint256 number, uint256 startBlock, uint256 endBlock)
func (_ValidatorSet *ValidatorSetCallerSession) GetNextSpan() (struct {
	Number     *big.Int
	StartBlock *big.Int
	EndBlock   *big.Int
}, error) {
	return _ValidatorSet.Contract.GetNextSpan(&_ValidatorSet.CallOpts)
}

// GetProducersTotalStakeBySpan is a free data retrieval call binding the contract method 0x9d11b807.
//
// Solidity: function getProducersTotalStakeBySpan(uint256 span) view returns(uint256)
func (_ValidatorSet *ValidatorSetCaller) GetProducersTotalStakeBySpan(opts *bind.CallOpts, span *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "getProducersTotalStakeBySpan", span)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProducersTotalStakeBySpan is a free data retrieval call binding the contract method 0x9d11b807.
//
// Solidity: function getProducersTotalStakeBySpan(uint256 span) view returns(uint256)
func (_ValidatorSet *ValidatorSetSession) GetProducersTotalStakeBySpan(span *big.Int) (*big.Int, error) {
	return _ValidatorSet.Contract.GetProducersTotalStakeBySpan(&_ValidatorSet.CallOpts, span)
}

// GetProducersTotalStakeBySpan is a free data retrieval call binding the contract method 0x9d11b807.
//
// Solidity: function getProducersTotalStakeBySpan(uint256 span) view returns(uint256)
func (_ValidatorSet *ValidatorSetCallerSession) GetProducersTotalStakeBySpan(span *big.Int) (*big.Int, error) {
	return _ValidatorSet.Contract.GetProducersTotalStakeBySpan(&_ValidatorSet.CallOpts, span)
}

// GetSpan is a free data retrieval call binding the contract method 0x047a6c5b.
//
// Solidity: function getSpan(uint256 span) view returns(uint256 number, uint256 startBlock, uint256 endBlock)
func (_ValidatorSet *ValidatorSetCaller) GetSpan(opts *bind.CallOpts, span *big.Int) (struct {
	Number     *big.Int
	StartBlock *big.Int
	EndBlock   *big.Int
}, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "getSpan", span)

	outstruct := new(struct {
		Number     *big.Int
		StartBlock *big.Int
		EndBlock   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Number = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSpan is a free data retrieval call binding the contract method 0x047a6c5b.
//
// Solidity: function getSpan(uint256 span) view returns(uint256 number, uint256 startBlock, uint256 endBlock)
func (_ValidatorSet *ValidatorSetSession) GetSpan(span *big.Int) (struct {
	Number     *big.Int
	StartBlock *big.Int
	EndBlock   *big.Int
}, error) {
	return _ValidatorSet.Contract.GetSpan(&_ValidatorSet.CallOpts, span)
}

// GetSpan is a free data retrieval call binding the contract method 0x047a6c5b.
//
// Solidity: function getSpan(uint256 span) view returns(uint256 number, uint256 startBlock, uint256 endBlock)
func (_ValidatorSet *ValidatorSetCallerSession) GetSpan(span *big.Int) (struct {
	Number     *big.Int
	StartBlock *big.Int
	EndBlock   *big.Int
}, error) {
	return _ValidatorSet.Contract.GetSpan(&_ValidatorSet.CallOpts, span)
}

// GetSpanByBlock is a free data retrieval call binding the contract method 0xb71d7a69.
//
// Solidity: function getSpanByBlock(uint256 number) view returns(uint256)
func (_ValidatorSet *ValidatorSetCaller) GetSpanByBlock(opts *bind.CallOpts, number *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "getSpanByBlock", number)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSpanByBlock is a free data retrieval call binding the contract method 0xb71d7a69.
//
// Solidity: function getSpanByBlock(uint256 number) view returns(uint256)
func (_ValidatorSet *ValidatorSetSession) GetSpanByBlock(number *big.Int) (*big.Int, error) {
	return _ValidatorSet.Contract.GetSpanByBlock(&_ValidatorSet.CallOpts, number)
}

// GetSpanByBlock is a free data retrieval call binding the contract method 0xb71d7a69.
//
// Solidity: function getSpanByBlock(uint256 number) view returns(uint256)
func (_ValidatorSet *ValidatorSetCallerSession) GetSpanByBlock(number *big.Int) (*big.Int, error) {
	return _ValidatorSet.Contract.GetSpanByBlock(&_ValidatorSet.CallOpts, number)
}

// GetStakePowerBySigs is a free data retrieval call binding the contract method 0x44c15cb1.
//
// Solidity: function getStakePowerBySigs(uint256 span, bytes32 dataHash, bytes sigs) view returns(uint256)
func (_ValidatorSet *ValidatorSetCaller) GetStakePowerBySigs(opts *bind.CallOpts, span *big.Int, dataHash [32]byte, sigs []byte) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "getStakePowerBySigs", span, dataHash, sigs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakePowerBySigs is a free data retrieval call binding the contract method 0x44c15cb1.
//
// Solidity: function getStakePowerBySigs(uint256 span, bytes32 dataHash, bytes sigs) view returns(uint256)
func (_ValidatorSet *ValidatorSetSession) GetStakePowerBySigs(span *big.Int, dataHash [32]byte, sigs []byte) (*big.Int, error) {
	return _ValidatorSet.Contract.GetStakePowerBySigs(&_ValidatorSet.CallOpts, span, dataHash, sigs)
}

// GetStakePowerBySigs is a free data retrieval call binding the contract method 0x44c15cb1.
//
// Solidity: function getStakePowerBySigs(uint256 span, bytes32 dataHash, bytes sigs) view returns(uint256)
func (_ValidatorSet *ValidatorSetCallerSession) GetStakePowerBySigs(span *big.Int, dataHash [32]byte, sigs []byte) (*big.Int, error) {
	return _ValidatorSet.Contract.GetStakePowerBySigs(&_ValidatorSet.CallOpts, span, dataHash, sigs)
}

// GetValidatorBySigner is a free data retrieval call binding the contract method 0x44d6528f.
//
// Solidity: function getValidatorBySigner(uint256 span, address signer) view returns((uint256,uint256,address) result)
func (_ValidatorSet *ValidatorSetCaller) GetValidatorBySigner(opts *bind.CallOpts, span *big.Int, signer common.Address) (BorValidatorSetValidator, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "getValidatorBySigner", span, signer)

	if err != nil {
		return *new(BorValidatorSetValidator), err
	}

	out0 := *abi.ConvertType(out[0], new(BorValidatorSetValidator)).(*BorValidatorSetValidator)

	return out0, err

}

// GetValidatorBySigner is a free data retrieval call binding the contract method 0x44d6528f.
//
// Solidity: function getValidatorBySigner(uint256 span, address signer) view returns((uint256,uint256,address) result)
func (_ValidatorSet *ValidatorSetSession) GetValidatorBySigner(span *big.Int, signer common.Address) (BorValidatorSetValidator, error) {
	return _ValidatorSet.Contract.GetValidatorBySigner(&_ValidatorSet.CallOpts, span, signer)
}

// GetValidatorBySigner is a free data retrieval call binding the contract method 0x44d6528f.
//
// Solidity: function getValidatorBySigner(uint256 span, address signer) view returns((uint256,uint256,address) result)
func (_ValidatorSet *ValidatorSetCallerSession) GetValidatorBySigner(span *big.Int, signer common.Address) (BorValidatorSetValidator, error) {
	return _ValidatorSet.Contract.GetValidatorBySigner(&_ValidatorSet.CallOpts, span, signer)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[], uint256[])
func (_ValidatorSet *ValidatorSetCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[], uint256[])
func (_ValidatorSet *ValidatorSetSession) GetValidators() ([]common.Address, []*big.Int, error) {
	return _ValidatorSet.Contract.GetValidators(&_ValidatorSet.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[], uint256[])
func (_ValidatorSet *ValidatorSetCallerSession) GetValidators() ([]common.Address, []*big.Int, error) {
	return _ValidatorSet.Contract.GetValidators(&_ValidatorSet.CallOpts)
}

// GetValidatorsTotalStakeBySpan is a free data retrieval call binding the contract method 0x2eddf352.
//
// Solidity: function getValidatorsTotalStakeBySpan(uint256 span) view returns(uint256)
func (_ValidatorSet *ValidatorSetCaller) GetValidatorsTotalStakeBySpan(opts *bind.CallOpts, span *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "getValidatorsTotalStakeBySpan", span)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorsTotalStakeBySpan is a free data retrieval call binding the contract method 0x2eddf352.
//
// Solidity: function getValidatorsTotalStakeBySpan(uint256 span) view returns(uint256)
func (_ValidatorSet *ValidatorSetSession) GetValidatorsTotalStakeBySpan(span *big.Int) (*big.Int, error) {
	return _ValidatorSet.Contract.GetValidatorsTotalStakeBySpan(&_ValidatorSet.CallOpts, span)
}

// GetValidatorsTotalStakeBySpan is a free data retrieval call binding the contract method 0x2eddf352.
//
// Solidity: function getValidatorsTotalStakeBySpan(uint256 span) view returns(uint256)
func (_ValidatorSet *ValidatorSetCallerSession) GetValidatorsTotalStakeBySpan(span *big.Int) (*big.Int, error) {
	return _ValidatorSet.Contract.GetValidatorsTotalStakeBySpan(&_ValidatorSet.CallOpts, span)
}

// InnerNode is a free data retrieval call binding the contract method 0x2de3a180.
//
// Solidity: function innerNode(bytes32 left, bytes32 right) pure returns(bytes32)
func (_ValidatorSet *ValidatorSetCaller) InnerNode(opts *bind.CallOpts, left [32]byte, right [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "innerNode", left, right)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InnerNode is a free data retrieval call binding the contract method 0x2de3a180.
//
// Solidity: function innerNode(bytes32 left, bytes32 right) pure returns(bytes32)
func (_ValidatorSet *ValidatorSetSession) InnerNode(left [32]byte, right [32]byte) ([32]byte, error) {
	return _ValidatorSet.Contract.InnerNode(&_ValidatorSet.CallOpts, left, right)
}

// InnerNode is a free data retrieval call binding the contract method 0x2de3a180.
//
// Solidity: function innerNode(bytes32 left, bytes32 right) pure returns(bytes32)
func (_ValidatorSet *ValidatorSetCallerSession) InnerNode(left [32]byte, right [32]byte) ([32]byte, error) {
	return _ValidatorSet.Contract.InnerNode(&_ValidatorSet.CallOpts, left, right)
}

// IsCurrentProducer is a free data retrieval call binding the contract method 0x70ba5707.
//
// Solidity: function isCurrentProducer(address signer) view returns(bool)
func (_ValidatorSet *ValidatorSetCaller) IsCurrentProducer(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "isCurrentProducer", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrentProducer is a free data retrieval call binding the contract method 0x70ba5707.
//
// Solidity: function isCurrentProducer(address signer) view returns(bool)
func (_ValidatorSet *ValidatorSetSession) IsCurrentProducer(signer common.Address) (bool, error) {
	return _ValidatorSet.Contract.IsCurrentProducer(&_ValidatorSet.CallOpts, signer)
}

// IsCurrentProducer is a free data retrieval call binding the contract method 0x70ba5707.
//
// Solidity: function isCurrentProducer(address signer) view returns(bool)
func (_ValidatorSet *ValidatorSetCallerSession) IsCurrentProducer(signer common.Address) (bool, error) {
	return _ValidatorSet.Contract.IsCurrentProducer(&_ValidatorSet.CallOpts, signer)
}

// IsCurrentValidator is a free data retrieval call binding the contract method 0x55614fcc.
//
// Solidity: function isCurrentValidator(address signer) view returns(bool)
func (_ValidatorSet *ValidatorSetCaller) IsCurrentValidator(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "isCurrentValidator", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrentValidator is a free data retrieval call binding the contract method 0x55614fcc.
//
// Solidity: function isCurrentValidator(address signer) view returns(bool)
func (_ValidatorSet *ValidatorSetSession) IsCurrentValidator(signer common.Address) (bool, error) {
	return _ValidatorSet.Contract.IsCurrentValidator(&_ValidatorSet.CallOpts, signer)
}

// IsCurrentValidator is a free data retrieval call binding the contract method 0x55614fcc.
//
// Solidity: function isCurrentValidator(address signer) view returns(bool)
func (_ValidatorSet *ValidatorSetCallerSession) IsCurrentValidator(signer common.Address) (bool, error) {
	return _ValidatorSet.Contract.IsCurrentValidator(&_ValidatorSet.CallOpts, signer)
}

// IsProducer is a free data retrieval call binding the contract method 0x1270b574.
//
// Solidity: function isProducer(uint256 span, address signer) view returns(bool)
func (_ValidatorSet *ValidatorSetCaller) IsProducer(opts *bind.CallOpts, span *big.Int, signer common.Address) (bool, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "isProducer", span, signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProducer is a free data retrieval call binding the contract method 0x1270b574.
//
// Solidity: function isProducer(uint256 span, address signer) view returns(bool)
func (_ValidatorSet *ValidatorSetSession) IsProducer(span *big.Int, signer common.Address) (bool, error) {
	return _ValidatorSet.Contract.IsProducer(&_ValidatorSet.CallOpts, span, signer)
}

// IsProducer is a free data retrieval call binding the contract method 0x1270b574.
//
// Solidity: function isProducer(uint256 span, address signer) view returns(bool)
func (_ValidatorSet *ValidatorSetCallerSession) IsProducer(span *big.Int, signer common.Address) (bool, error) {
	return _ValidatorSet.Contract.IsProducer(&_ValidatorSet.CallOpts, span, signer)
}

// IsValidator is a free data retrieval call binding the contract method 0x23f2a73f.
//
// Solidity: function isValidator(uint256 span, address signer) view returns(bool)
func (_ValidatorSet *ValidatorSetCaller) IsValidator(opts *bind.CallOpts, span *big.Int, signer common.Address) (bool, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "isValidator", span, signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0x23f2a73f.
//
// Solidity: function isValidator(uint256 span, address signer) view returns(bool)
func (_ValidatorSet *ValidatorSetSession) IsValidator(span *big.Int, signer common.Address) (bool, error) {
	return _ValidatorSet.Contract.IsValidator(&_ValidatorSet.CallOpts, span, signer)
}

// IsValidator is a free data retrieval call binding the contract method 0x23f2a73f.
//
// Solidity: function isValidator(uint256 span, address signer) view returns(bool)
func (_ValidatorSet *ValidatorSetCallerSession) IsValidator(span *big.Int, signer common.Address) (bool, error) {
	return _ValidatorSet.Contract.IsValidator(&_ValidatorSet.CallOpts, span, signer)
}

// LeafNode is a free data retrieval call binding the contract method 0x582a8d08.
//
// Solidity: function leafNode(bytes32 d) pure returns(bytes32)
func (_ValidatorSet *ValidatorSetCaller) LeafNode(opts *bind.CallOpts, d [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "leafNode", d)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LeafNode is a free data retrieval call binding the contract method 0x582a8d08.
//
// Solidity: function leafNode(bytes32 d) pure returns(bytes32)
func (_ValidatorSet *ValidatorSetSession) LeafNode(d [32]byte) ([32]byte, error) {
	return _ValidatorSet.Contract.LeafNode(&_ValidatorSet.CallOpts, d)
}

// LeafNode is a free data retrieval call binding the contract method 0x582a8d08.
//
// Solidity: function leafNode(bytes32 d) pure returns(bytes32)
func (_ValidatorSet *ValidatorSetCallerSession) LeafNode(d [32]byte) ([32]byte, error) {
	return _ValidatorSet.Contract.LeafNode(&_ValidatorSet.CallOpts, d)
}

// Producers is a free data retrieval call binding the contract method 0x687a9bd6.
//
// Solidity: function producers(uint256 , uint256 ) view returns(uint256 id, uint256 power, address signer)
func (_ValidatorSet *ValidatorSetCaller) Producers(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Id     *big.Int
	Power  *big.Int
	Signer common.Address
}, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "producers", arg0, arg1)

	outstruct := new(struct {
		Id     *big.Int
		Power  *big.Int
		Signer common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Power = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Signer = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Producers is a free data retrieval call binding the contract method 0x687a9bd6.
//
// Solidity: function producers(uint256 , uint256 ) view returns(uint256 id, uint256 power, address signer)
func (_ValidatorSet *ValidatorSetSession) Producers(arg0 *big.Int, arg1 *big.Int) (struct {
	Id     *big.Int
	Power  *big.Int
	Signer common.Address
}, error) {
	return _ValidatorSet.Contract.Producers(&_ValidatorSet.CallOpts, arg0, arg1)
}

// Producers is a free data retrieval call binding the contract method 0x687a9bd6.
//
// Solidity: function producers(uint256 , uint256 ) view returns(uint256 id, uint256 power, address signer)
func (_ValidatorSet *ValidatorSetCallerSession) Producers(arg0 *big.Int, arg1 *big.Int) (struct {
	Id     *big.Int
	Power  *big.Int
	Signer common.Address
}, error) {
	return _ValidatorSet.Contract.Producers(&_ValidatorSet.CallOpts, arg0, arg1)
}

// SpanNumbers is a free data retrieval call binding the contract method 0xc1b3c919.
//
// Solidity: function spanNumbers(uint256 ) view returns(uint256)
func (_ValidatorSet *ValidatorSetCaller) SpanNumbers(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "spanNumbers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpanNumbers is a free data retrieval call binding the contract method 0xc1b3c919.
//
// Solidity: function spanNumbers(uint256 ) view returns(uint256)
func (_ValidatorSet *ValidatorSetSession) SpanNumbers(arg0 *big.Int) (*big.Int, error) {
	return _ValidatorSet.Contract.SpanNumbers(&_ValidatorSet.CallOpts, arg0)
}

// SpanNumbers is a free data retrieval call binding the contract method 0xc1b3c919.
//
// Solidity: function spanNumbers(uint256 ) view returns(uint256)
func (_ValidatorSet *ValidatorSetCallerSession) SpanNumbers(arg0 *big.Int) (*big.Int, error) {
	return _ValidatorSet.Contract.SpanNumbers(&_ValidatorSet.CallOpts, arg0)
}

// Spans is a free data retrieval call binding the contract method 0xf59cf565.
//
// Solidity: function spans(uint256 ) view returns(uint256 number, uint256 startBlock, uint256 endBlock)
func (_ValidatorSet *ValidatorSetCaller) Spans(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Number     *big.Int
	StartBlock *big.Int
	EndBlock   *big.Int
}, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "spans", arg0)

	outstruct := new(struct {
		Number     *big.Int
		StartBlock *big.Int
		EndBlock   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Number = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Spans is a free data retrieval call binding the contract method 0xf59cf565.
//
// Solidity: function spans(uint256 ) view returns(uint256 number, uint256 startBlock, uint256 endBlock)
func (_ValidatorSet *ValidatorSetSession) Spans(arg0 *big.Int) (struct {
	Number     *big.Int
	StartBlock *big.Int
	EndBlock   *big.Int
}, error) {
	return _ValidatorSet.Contract.Spans(&_ValidatorSet.CallOpts, arg0)
}

// Spans is a free data retrieval call binding the contract method 0xf59cf565.
//
// Solidity: function spans(uint256 ) view returns(uint256 number, uint256 startBlock, uint256 endBlock)
func (_ValidatorSet *ValidatorSetCallerSession) Spans(arg0 *big.Int) (struct {
	Number     *big.Int
	StartBlock *big.Int
	EndBlock   *big.Int
}, error) {
	return _ValidatorSet.Contract.Spans(&_ValidatorSet.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xdcf2793a.
//
// Solidity: function validators(uint256 , uint256 ) view returns(uint256 id, uint256 power, address signer)
func (_ValidatorSet *ValidatorSetCaller) Validators(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Id     *big.Int
	Power  *big.Int
	Signer common.Address
}, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "validators", arg0, arg1)

	outstruct := new(struct {
		Id     *big.Int
		Power  *big.Int
		Signer common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Power = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Signer = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0xdcf2793a.
//
// Solidity: function validators(uint256 , uint256 ) view returns(uint256 id, uint256 power, address signer)
func (_ValidatorSet *ValidatorSetSession) Validators(arg0 *big.Int, arg1 *big.Int) (struct {
	Id     *big.Int
	Power  *big.Int
	Signer common.Address
}, error) {
	return _ValidatorSet.Contract.Validators(&_ValidatorSet.CallOpts, arg0, arg1)
}

// Validators is a free data retrieval call binding the contract method 0xdcf2793a.
//
// Solidity: function validators(uint256 , uint256 ) view returns(uint256 id, uint256 power, address signer)
func (_ValidatorSet *ValidatorSetCallerSession) Validators(arg0 *big.Int, arg1 *big.Int) (struct {
	Id     *big.Int
	Power  *big.Int
	Signer common.Address
}, error) {
	return _ValidatorSet.Contract.Validators(&_ValidatorSet.CallOpts, arg0, arg1)
}

// CommitSpan is a paid mutator transaction binding the contract method 0x23c2a2b4.
//
// Solidity: function commitSpan(uint256 newSpan, uint256 startBlock, uint256 endBlock, bytes validatorBytes, bytes producerBytes) returns()
func (_ValidatorSet *ValidatorSetTransactor) CommitSpan(opts *bind.TransactOpts, newSpan *big.Int, startBlock *big.Int, endBlock *big.Int, validatorBytes []byte, producerBytes []byte) (*types.Transaction, error) {
	return _ValidatorSet.contract.Transact(opts, "commitSpan", newSpan, startBlock, endBlock, validatorBytes, producerBytes)
}

// CommitSpan is a paid mutator transaction binding the contract method 0x23c2a2b4.
//
// Solidity: function commitSpan(uint256 newSpan, uint256 startBlock, uint256 endBlock, bytes validatorBytes, bytes producerBytes) returns()
func (_ValidatorSet *ValidatorSetSession) CommitSpan(newSpan *big.Int, startBlock *big.Int, endBlock *big.Int, validatorBytes []byte, producerBytes []byte) (*types.Transaction, error) {
	return _ValidatorSet.Contract.CommitSpan(&_ValidatorSet.TransactOpts, newSpan, startBlock, endBlock, validatorBytes, producerBytes)
}

// CommitSpan is a paid mutator transaction binding the contract method 0x23c2a2b4.
//
// Solidity: function commitSpan(uint256 newSpan, uint256 startBlock, uint256 endBlock, bytes validatorBytes, bytes producerBytes) returns()
func (_ValidatorSet *ValidatorSetTransactorSession) CommitSpan(newSpan *big.Int, startBlock *big.Int, endBlock *big.Int, validatorBytes []byte, producerBytes []byte) (*types.Transaction, error) {
	return _ValidatorSet.Contract.CommitSpan(&_ValidatorSet.TransactOpts, newSpan, startBlock, endBlock, validatorBytes, producerBytes)
}

// ValidatorSetNewSpanIterator is returned from FilterNewSpan and is used to iterate over the raw logs and unpacked data for NewSpan events raised by the ValidatorSet contract.
type ValidatorSetNewSpanIterator struct {
	Event *ValidatorSetNewSpan // Event containing the contract specifics and raw log

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
func (it *ValidatorSetNewSpanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorSetNewSpan)
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
		it.Event = new(ValidatorSetNewSpan)
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
func (it *ValidatorSetNewSpanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorSetNewSpanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorSetNewSpan represents a NewSpan event raised by the ValidatorSet contract.
type ValidatorSetNewSpan struct {
	Id         *big.Int
	StartBlock *big.Int
	EndBlock   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewSpan is a free log retrieval operation binding the contract event 0xac9e8537ff98ccc9f53f34174714e4b56eee15cb11eafa9512d2914c60790c62.
//
// Solidity: event NewSpan(uint256 indexed id, uint256 indexed startBlock, uint256 indexed endBlock)
func (_ValidatorSet *ValidatorSetFilterer) FilterNewSpan(opts *bind.FilterOpts, id []*big.Int, startBlock []*big.Int, endBlock []*big.Int) (*ValidatorSetNewSpanIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var startBlockRule []interface{}
	for _, startBlockItem := range startBlock {
		startBlockRule = append(startBlockRule, startBlockItem)
	}
	var endBlockRule []interface{}
	for _, endBlockItem := range endBlock {
		endBlockRule = append(endBlockRule, endBlockItem)
	}

	logs, sub, err := _ValidatorSet.contract.FilterLogs(opts, "NewSpan", idRule, startBlockRule, endBlockRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetNewSpanIterator{contract: _ValidatorSet.contract, event: "NewSpan", logs: logs, sub: sub}, nil
}

// WatchNewSpan is a free log subscription operation binding the contract event 0xac9e8537ff98ccc9f53f34174714e4b56eee15cb11eafa9512d2914c60790c62.
//
// Solidity: event NewSpan(uint256 indexed id, uint256 indexed startBlock, uint256 indexed endBlock)
func (_ValidatorSet *ValidatorSetFilterer) WatchNewSpan(opts *bind.WatchOpts, sink chan<- *ValidatorSetNewSpan, id []*big.Int, startBlock []*big.Int, endBlock []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var startBlockRule []interface{}
	for _, startBlockItem := range startBlock {
		startBlockRule = append(startBlockRule, startBlockItem)
	}
	var endBlockRule []interface{}
	for _, endBlockItem := range endBlock {
		endBlockRule = append(endBlockRule, endBlockItem)
	}

	logs, sub, err := _ValidatorSet.contract.WatchLogs(opts, "NewSpan", idRule, startBlockRule, endBlockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorSetNewSpan)
				if err := _ValidatorSet.contract.UnpackLog(event, "NewSpan", log); err != nil {
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

// ParseNewSpan is a log parse operation binding the contract event 0xac9e8537ff98ccc9f53f34174714e4b56eee15cb11eafa9512d2914c60790c62.
//
// Solidity: event NewSpan(uint256 indexed id, uint256 indexed startBlock, uint256 indexed endBlock)
func (_ValidatorSet *ValidatorSetFilterer) ParseNewSpan(log types.Log) (*ValidatorSetNewSpan, error) {
	event := new(ValidatorSetNewSpan)
	if err := _ValidatorSet.contract.UnpackLog(event, "NewSpan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
