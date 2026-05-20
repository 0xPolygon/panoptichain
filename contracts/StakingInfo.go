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

// StakingInfoMetaData contains all meta data concerning the StakingInfo contract.
var StakingInfoMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"activationEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signerPubkey\",\"type\":\"bytes\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivationEpoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnstakeInit\",\"type\":\"event\"}]",
}

// StakingInfoABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingInfoMetaData.ABI instead.
var StakingInfoABI = StakingInfoMetaData.ABI

// StakingInfo is an auto generated Go binding around an Ethereum contract.
type StakingInfo struct {
	StakingInfoCaller     // Read-only binding to the contract
	StakingInfoTransactor // Write-only binding to the contract
	StakingInfoFilterer   // Log filterer for contract events
}

// StakingInfoCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingInfoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingInfoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingInfoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingInfoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingInfoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingInfoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingInfoSession struct {
	Contract     *StakingInfo      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingInfoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingInfoCallerSession struct {
	Contract *StakingInfoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// StakingInfoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingInfoTransactorSession struct {
	Contract     *StakingInfoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StakingInfoRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingInfoRaw struct {
	Contract *StakingInfo // Generic contract binding to access the raw methods on
}

// StakingInfoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingInfoCallerRaw struct {
	Contract *StakingInfoCaller // Generic read-only contract binding to access the raw methods on
}

// StakingInfoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingInfoTransactorRaw struct {
	Contract *StakingInfoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingInfo creates a new instance of StakingInfo, bound to a specific deployed contract.
func NewStakingInfo(address common.Address, backend bind.ContractBackend) (*StakingInfo, error) {
	contract, err := bindStakingInfo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingInfo{StakingInfoCaller: StakingInfoCaller{contract: contract}, StakingInfoTransactor: StakingInfoTransactor{contract: contract}, StakingInfoFilterer: StakingInfoFilterer{contract: contract}}, nil
}

// NewStakingInfoCaller creates a new read-only instance of StakingInfo, bound to a specific deployed contract.
func NewStakingInfoCaller(address common.Address, caller bind.ContractCaller) (*StakingInfoCaller, error) {
	contract, err := bindStakingInfo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingInfoCaller{contract: contract}, nil
}

// NewStakingInfoTransactor creates a new write-only instance of StakingInfo, bound to a specific deployed contract.
func NewStakingInfoTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingInfoTransactor, error) {
	contract, err := bindStakingInfo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingInfoTransactor{contract: contract}, nil
}

// NewStakingInfoFilterer creates a new log filterer instance of StakingInfo, bound to a specific deployed contract.
func NewStakingInfoFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingInfoFilterer, error) {
	contract, err := bindStakingInfo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingInfoFilterer{contract: contract}, nil
}

// bindStakingInfo binds a generic wrapper to an already deployed contract.
func bindStakingInfo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingInfoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingInfo *StakingInfoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingInfo.Contract.StakingInfoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingInfo *StakingInfoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingInfo.Contract.StakingInfoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingInfo *StakingInfoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingInfo.Contract.StakingInfoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingInfo *StakingInfoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingInfo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingInfo *StakingInfoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingInfo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingInfo *StakingInfoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingInfo.Contract.contract.Transact(opts, method, params...)
}

// StakingInfoStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the StakingInfo contract.
type StakingInfoStakedIterator struct {
	Event *StakingInfoStaked // Event containing the contract specifics and raw log

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
func (it *StakingInfoStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingInfoStaked)
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
		it.Event = new(StakingInfoStaked)
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
func (it *StakingInfoStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingInfoStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingInfoStaked represents a Staked event raised by the StakingInfo contract.
type StakingInfoStaked struct {
	Signer          common.Address
	ValidatorId     *big.Int
	Nonce           *big.Int
	ActivationEpoch *big.Int
	Amount          *big.Int
	Total           *big.Int
	SignerPubkey    []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x68c13e4125b983d7e2d6114246f443e567ec6c4ee5b4d4a7ef6100b1402bfd84.
//
// Solidity: event Staked(address indexed signer, uint256 indexed validatorId, uint256 nonce, uint256 indexed activationEpoch, uint256 amount, uint256 total, bytes signerPubkey)
func (_StakingInfo *StakingInfoFilterer) FilterStaked(opts *bind.FilterOpts, signer []common.Address, validatorId []*big.Int, activationEpoch []*big.Int) (*StakingInfoStakedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	var activationEpochRule []interface{}
	for _, activationEpochItem := range activationEpoch {
		activationEpochRule = append(activationEpochRule, activationEpochItem)
	}

	logs, sub, err := _StakingInfo.contract.FilterLogs(opts, "Staked", signerRule, validatorIdRule, activationEpochRule)
	if err != nil {
		return nil, err
	}
	return &StakingInfoStakedIterator{contract: _StakingInfo.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x68c13e4125b983d7e2d6114246f443e567ec6c4ee5b4d4a7ef6100b1402bfd84.
//
// Solidity: event Staked(address indexed signer, uint256 indexed validatorId, uint256 nonce, uint256 indexed activationEpoch, uint256 amount, uint256 total, bytes signerPubkey)
func (_StakingInfo *StakingInfoFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingInfoStaked, signer []common.Address, validatorId []*big.Int, activationEpoch []*big.Int) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	var activationEpochRule []interface{}
	for _, activationEpochItem := range activationEpoch {
		activationEpochRule = append(activationEpochRule, activationEpochItem)
	}

	logs, sub, err := _StakingInfo.contract.WatchLogs(opts, "Staked", signerRule, validatorIdRule, activationEpochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingInfoStaked)
				if err := _StakingInfo.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x68c13e4125b983d7e2d6114246f443e567ec6c4ee5b4d4a7ef6100b1402bfd84.
//
// Solidity: event Staked(address indexed signer, uint256 indexed validatorId, uint256 nonce, uint256 indexed activationEpoch, uint256 amount, uint256 total, bytes signerPubkey)
func (_StakingInfo *StakingInfoFilterer) ParseStaked(log types.Log) (*StakingInfoStaked, error) {
	event := new(StakingInfoStaked)
	if err := _StakingInfo.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingInfoUnstakeInitIterator is returned from FilterUnstakeInit and is used to iterate over the raw logs and unpacked data for UnstakeInit events raised by the StakingInfo contract.
type StakingInfoUnstakeInitIterator struct {
	Event *StakingInfoUnstakeInit // Event containing the contract specifics and raw log

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
func (it *StakingInfoUnstakeInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingInfoUnstakeInit)
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
		it.Event = new(StakingInfoUnstakeInit)
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
func (it *StakingInfoUnstakeInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingInfoUnstakeInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingInfoUnstakeInit represents a UnstakeInit event raised by the StakingInfo contract.
type StakingInfoUnstakeInit struct {
	User              common.Address
	ValidatorId       *big.Int
	Nonce             *big.Int
	DeactivationEpoch *big.Int
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUnstakeInit is a free log retrieval operation binding the contract event 0x69b288bb79cd5386c9fe0af060f650e823bcdfa96a44cdc07f862db060f57120.
//
// Solidity: event UnstakeInit(address indexed user, uint256 indexed validatorId, uint256 nonce, uint256 deactivationEpoch, uint256 indexed amount)
func (_StakingInfo *StakingInfoFilterer) FilterUnstakeInit(opts *bind.FilterOpts, user []common.Address, validatorId []*big.Int, amount []*big.Int) (*StakingInfoUnstakeInitIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _StakingInfo.contract.FilterLogs(opts, "UnstakeInit", userRule, validatorIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &StakingInfoUnstakeInitIterator{contract: _StakingInfo.contract, event: "UnstakeInit", logs: logs, sub: sub}, nil
}

// WatchUnstakeInit is a free log subscription operation binding the contract event 0x69b288bb79cd5386c9fe0af060f650e823bcdfa96a44cdc07f862db060f57120.
//
// Solidity: event UnstakeInit(address indexed user, uint256 indexed validatorId, uint256 nonce, uint256 deactivationEpoch, uint256 indexed amount)
func (_StakingInfo *StakingInfoFilterer) WatchUnstakeInit(opts *bind.WatchOpts, sink chan<- *StakingInfoUnstakeInit, user []common.Address, validatorId []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _StakingInfo.contract.WatchLogs(opts, "UnstakeInit", userRule, validatorIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingInfoUnstakeInit)
				if err := _StakingInfo.contract.UnpackLog(event, "UnstakeInit", log); err != nil {
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

// ParseUnstakeInit is a log parse operation binding the contract event 0x69b288bb79cd5386c9fe0af060f650e823bcdfa96a44cdc07f862db060f57120.
//
// Solidity: event UnstakeInit(address indexed user, uint256 indexed validatorId, uint256 nonce, uint256 deactivationEpoch, uint256 indexed amount)
func (_StakingInfo *StakingInfoFilterer) ParseUnstakeInit(log types.Log) (*StakingInfoUnstakeInit, error) {
	event := new(StakingInfoUnstakeInit)
	if err := _StakingInfo.contract.UnpackLog(event, "UnstakeInit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
