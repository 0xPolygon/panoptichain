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

// PolygonZkEVMBridgeMetaData contains all meta data concerning the PolygonZkEVMBridge contract.
var PolygonZkEVMBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountDoesNotMatchMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DestinationNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EtherTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSmtProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPolygonZkEVM\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"depositCount\",\"type\":\"uint32\"}],\"name\":\"BridgeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"NewWrappedToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"bridgeAsset\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimedBitMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDepositRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"getLeafValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenWrappedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_networkID\",\"type\":\"uint32\"},{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonZkEVMaddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedDepositCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polygonZkEVMaddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"precalculatedWrapperAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenInfoToWrappedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedTokenToTokenInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PolygonZkEVMBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonZkEVMBridgeMetaData.ABI instead.
var PolygonZkEVMBridgeABI = PolygonZkEVMBridgeMetaData.ABI

// PolygonZkEVMBridge is an auto generated Go binding around an Ethereum contract.
type PolygonZkEVMBridge struct {
	PolygonZkEVMBridgeCaller     // Read-only binding to the contract
	PolygonZkEVMBridgeTransactor // Write-only binding to the contract
	PolygonZkEVMBridgeFilterer   // Log filterer for contract events
}

// PolygonZkEVMBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonZkEVMBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonZkEVMBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonZkEVMBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonZkEVMBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonZkEVMBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonZkEVMBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonZkEVMBridgeSession struct {
	Contract     *PolygonZkEVMBridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PolygonZkEVMBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonZkEVMBridgeCallerSession struct {
	Contract *PolygonZkEVMBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// PolygonZkEVMBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonZkEVMBridgeTransactorSession struct {
	Contract     *PolygonZkEVMBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PolygonZkEVMBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonZkEVMBridgeRaw struct {
	Contract *PolygonZkEVMBridge // Generic contract binding to access the raw methods on
}

// PolygonZkEVMBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonZkEVMBridgeCallerRaw struct {
	Contract *PolygonZkEVMBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonZkEVMBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonZkEVMBridgeTransactorRaw struct {
	Contract *PolygonZkEVMBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonZkEVMBridge creates a new instance of PolygonZkEVMBridge, bound to a specific deployed contract.
func NewPolygonZkEVMBridge(address common.Address, backend bind.ContractBackend) (*PolygonZkEVMBridge, error) {
	contract, err := bindPolygonZkEVMBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMBridge{PolygonZkEVMBridgeCaller: PolygonZkEVMBridgeCaller{contract: contract}, PolygonZkEVMBridgeTransactor: PolygonZkEVMBridgeTransactor{contract: contract}, PolygonZkEVMBridgeFilterer: PolygonZkEVMBridgeFilterer{contract: contract}}, nil
}

// NewPolygonZkEVMBridgeCaller creates a new read-only instance of PolygonZkEVMBridge, bound to a specific deployed contract.
func NewPolygonZkEVMBridgeCaller(address common.Address, caller bind.ContractCaller) (*PolygonZkEVMBridgeCaller, error) {
	contract, err := bindPolygonZkEVMBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMBridgeCaller{contract: contract}, nil
}

// NewPolygonZkEVMBridgeTransactor creates a new write-only instance of PolygonZkEVMBridge, bound to a specific deployed contract.
func NewPolygonZkEVMBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonZkEVMBridgeTransactor, error) {
	contract, err := bindPolygonZkEVMBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMBridgeTransactor{contract: contract}, nil
}

// NewPolygonZkEVMBridgeFilterer creates a new log filterer instance of PolygonZkEVMBridge, bound to a specific deployed contract.
func NewPolygonZkEVMBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonZkEVMBridgeFilterer, error) {
	contract, err := bindPolygonZkEVMBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMBridgeFilterer{contract: contract}, nil
}

// bindPolygonZkEVMBridge binds a generic wrapper to an already deployed contract.
func bindPolygonZkEVMBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonZkEVMBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolygonZkEVMBridge.Contract.PolygonZkEVMBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.PolygonZkEVMBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.PolygonZkEVMBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolygonZkEVMBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.contract.Transact(opts, method, params...)
}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCaller) ClaimedBitMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PolygonZkEVMBridge.contract.Call(opts, &out, "claimedBitMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) ClaimedBitMap(arg0 *big.Int) (*big.Int, error) {
	return _PolygonZkEVMBridge.Contract.ClaimedBitMap(&_PolygonZkEVMBridge.CallOpts, arg0)
}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCallerSession) ClaimedBitMap(arg0 *big.Int) (*big.Int, error) {
	return _PolygonZkEVMBridge.Contract.ClaimedBitMap(&_PolygonZkEVMBridge.CallOpts, arg0)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCaller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolygonZkEVMBridge.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) DepositCount() (*big.Int, error) {
	return _PolygonZkEVMBridge.Contract.DepositCount(&_PolygonZkEVMBridge.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCallerSession) DepositCount() (*big.Int, error) {
	return _PolygonZkEVMBridge.Contract.DepositCount(&_PolygonZkEVMBridge.CallOpts)
}

// GetDepositRoot is a free data retrieval call binding the contract method 0x3ae05047.
//
// Solidity: function getDepositRoot() view returns(bytes32)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCaller) GetDepositRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolygonZkEVMBridge.contract.Call(opts, &out, "getDepositRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDepositRoot is a free data retrieval call binding the contract method 0x3ae05047.
//
// Solidity: function getDepositRoot() view returns(bytes32)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) GetDepositRoot() ([32]byte, error) {
	return _PolygonZkEVMBridge.Contract.GetDepositRoot(&_PolygonZkEVMBridge.CallOpts)
}

// GetDepositRoot is a free data retrieval call binding the contract method 0x3ae05047.
//
// Solidity: function getDepositRoot() view returns(bytes32)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCallerSession) GetDepositRoot() ([32]byte, error) {
	return _PolygonZkEVMBridge.Contract.GetDepositRoot(&_PolygonZkEVMBridge.CallOpts)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCaller) GetLeafValue(opts *bind.CallOpts, leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PolygonZkEVMBridge.contract.Call(opts, &out, "getLeafValue", leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) GetLeafValue(leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	return _PolygonZkEVMBridge.Contract.GetLeafValue(&_PolygonZkEVMBridge.CallOpts, leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCallerSession) GetLeafValue(leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	return _PolygonZkEVMBridge.Contract.GetLeafValue(&_PolygonZkEVMBridge.CallOpts, leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)
}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCaller) GetTokenWrappedAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _PolygonZkEVMBridge.contract.Call(opts, &out, "getTokenWrappedAddress", originNetwork, originTokenAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) GetTokenWrappedAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _PolygonZkEVMBridge.Contract.GetTokenWrappedAddress(&_PolygonZkEVMBridge.CallOpts, originNetwork, originTokenAddress)
}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCallerSession) GetTokenWrappedAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _PolygonZkEVMBridge.Contract.GetTokenWrappedAddress(&_PolygonZkEVMBridge.CallOpts, originNetwork, originTokenAddress)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonZkEVMBridge.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) GlobalExitRootManager() (common.Address, error) {
	return _PolygonZkEVMBridge.Contract.GlobalExitRootManager(&_PolygonZkEVMBridge.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _PolygonZkEVMBridge.Contract.GlobalExitRootManager(&_PolygonZkEVMBridge.CallOpts)
}

// IsClaimed is a free data retrieval call binding the contract method 0x9e34070f.
//
// Solidity: function isClaimed(uint256 index) view returns(bool)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCaller) IsClaimed(opts *bind.CallOpts, index *big.Int) (bool, error) {
	var out []interface{}
	err := _PolygonZkEVMBridge.contract.Call(opts, &out, "isClaimed", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0x9e34070f.
//
// Solidity: function isClaimed(uint256 index) view returns(bool)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) IsClaimed(index *big.Int) (bool, error) {
	return _PolygonZkEVMBridge.Contract.IsClaimed(&_PolygonZkEVMBridge.CallOpts, index)
}

// IsClaimed is a free data retrieval call binding the contract method 0x9e34070f.
//
// Solidity: function isClaimed(uint256 index) view returns(bool)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCallerSession) IsClaimed(index *big.Int) (bool, error) {
	return _PolygonZkEVMBridge.Contract.IsClaimed(&_PolygonZkEVMBridge.CallOpts, index)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PolygonZkEVMBridge.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) IsEmergencyState() (bool, error) {
	return _PolygonZkEVMBridge.Contract.IsEmergencyState(&_PolygonZkEVMBridge.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCallerSession) IsEmergencyState() (bool, error) {
	return _PolygonZkEVMBridge.Contract.IsEmergencyState(&_PolygonZkEVMBridge.CallOpts)
}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCaller) LastUpdatedDepositCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PolygonZkEVMBridge.contract.Call(opts, &out, "lastUpdatedDepositCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) LastUpdatedDepositCount() (uint32, error) {
	return _PolygonZkEVMBridge.Contract.LastUpdatedDepositCount(&_PolygonZkEVMBridge.CallOpts)
}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCallerSession) LastUpdatedDepositCount() (uint32, error) {
	return _PolygonZkEVMBridge.Contract.LastUpdatedDepositCount(&_PolygonZkEVMBridge.CallOpts)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCaller) NetworkID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PolygonZkEVMBridge.contract.Call(opts, &out, "networkID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) NetworkID() (uint32, error) {
	return _PolygonZkEVMBridge.Contract.NetworkID(&_PolygonZkEVMBridge.CallOpts)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCallerSession) NetworkID() (uint32, error) {
	return _PolygonZkEVMBridge.Contract.NetworkID(&_PolygonZkEVMBridge.CallOpts)
}

// PolygonZkEVMaddress is a free data retrieval call binding the contract method 0x34ac9cf2.
//
// Solidity: function polygonZkEVMaddress() view returns(address)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCaller) PolygonZkEVMaddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonZkEVMBridge.contract.Call(opts, &out, "polygonZkEVMaddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolygonZkEVMaddress is a free data retrieval call binding the contract method 0x34ac9cf2.
//
// Solidity: function polygonZkEVMaddress() view returns(address)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) PolygonZkEVMaddress() (common.Address, error) {
	return _PolygonZkEVMBridge.Contract.PolygonZkEVMaddress(&_PolygonZkEVMBridge.CallOpts)
}

// PolygonZkEVMaddress is a free data retrieval call binding the contract method 0x34ac9cf2.
//
// Solidity: function polygonZkEVMaddress() view returns(address)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCallerSession) PolygonZkEVMaddress() (common.Address, error) {
	return _PolygonZkEVMBridge.Contract.PolygonZkEVMaddress(&_PolygonZkEVMBridge.CallOpts)
}

// PrecalculatedWrapperAddress is a free data retrieval call binding the contract method 0xaaa13cc2.
//
// Solidity: function precalculatedWrapperAddress(uint32 originNetwork, address originTokenAddress, string name, string symbol, uint8 decimals) view returns(address)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCaller) PrecalculatedWrapperAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address, name string, symbol string, decimals uint8) (common.Address, error) {
	var out []interface{}
	err := _PolygonZkEVMBridge.contract.Call(opts, &out, "precalculatedWrapperAddress", originNetwork, originTokenAddress, name, symbol, decimals)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrecalculatedWrapperAddress is a free data retrieval call binding the contract method 0xaaa13cc2.
//
// Solidity: function precalculatedWrapperAddress(uint32 originNetwork, address originTokenAddress, string name, string symbol, uint8 decimals) view returns(address)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) PrecalculatedWrapperAddress(originNetwork uint32, originTokenAddress common.Address, name string, symbol string, decimals uint8) (common.Address, error) {
	return _PolygonZkEVMBridge.Contract.PrecalculatedWrapperAddress(&_PolygonZkEVMBridge.CallOpts, originNetwork, originTokenAddress, name, symbol, decimals)
}

// PrecalculatedWrapperAddress is a free data retrieval call binding the contract method 0xaaa13cc2.
//
// Solidity: function precalculatedWrapperAddress(uint32 originNetwork, address originTokenAddress, string name, string symbol, uint8 decimals) view returns(address)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCallerSession) PrecalculatedWrapperAddress(originNetwork uint32, originTokenAddress common.Address, name string, symbol string, decimals uint8) (common.Address, error) {
	return _PolygonZkEVMBridge.Contract.PrecalculatedWrapperAddress(&_PolygonZkEVMBridge.CallOpts, originNetwork, originTokenAddress, name, symbol, decimals)
}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCaller) TokenInfoToWrappedToken(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _PolygonZkEVMBridge.contract.Call(opts, &out, "tokenInfoToWrappedToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) TokenInfoToWrappedToken(arg0 [32]byte) (common.Address, error) {
	return _PolygonZkEVMBridge.Contract.TokenInfoToWrappedToken(&_PolygonZkEVMBridge.CallOpts, arg0)
}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCallerSession) TokenInfoToWrappedToken(arg0 [32]byte) (common.Address, error) {
	return _PolygonZkEVMBridge.Contract.TokenInfoToWrappedToken(&_PolygonZkEVMBridge.CallOpts, arg0)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCaller) VerifyMerkleProof(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	var out []interface{}
	err := _PolygonZkEVMBridge.contract.Call(opts, &out, "verifyMerkleProof", leafHash, smtProof, index, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _PolygonZkEVMBridge.Contract.VerifyMerkleProof(&_PolygonZkEVMBridge.CallOpts, leafHash, smtProof, index, root)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCallerSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _PolygonZkEVMBridge.Contract.VerifyMerkleProof(&_PolygonZkEVMBridge.CallOpts, leafHash, smtProof, index, root)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCaller) WrappedTokenToTokenInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	var out []interface{}
	err := _PolygonZkEVMBridge.contract.Call(opts, &out, "wrappedTokenToTokenInfo", arg0)

	outstruct := new(struct {
		OriginNetwork      uint32
		OriginTokenAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OriginNetwork = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.OriginTokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) WrappedTokenToTokenInfo(arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	return _PolygonZkEVMBridge.Contract.WrappedTokenToTokenInfo(&_PolygonZkEVMBridge.CallOpts, arg0)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeCallerSession) WrappedTokenToTokenInfo(arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	return _PolygonZkEVMBridge.Contract.WrappedTokenToTokenInfo(&_PolygonZkEVMBridge.CallOpts, arg0)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.ActivateEmergencyState(&_PolygonZkEVMBridge.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.ActivateEmergencyState(&_PolygonZkEVMBridge.TransactOpts)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeTransactor) BridgeAsset(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.contract.Transact(opts, "bridgeAsset", destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.BridgeAsset(&_PolygonZkEVMBridge.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeTransactorSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.BridgeAsset(&_PolygonZkEVMBridge.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeTransactor) BridgeMessage(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.contract.Transact(opts, "bridgeMessage", destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.BridgeMessage(&_PolygonZkEVMBridge.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeTransactorSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.BridgeMessage(&_PolygonZkEVMBridge.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0x2cffd02e.
//
// Solidity: function claimAsset(bytes32[32] smtProof, uint32 index, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeTransactor) ClaimAsset(opts *bind.TransactOpts, smtProof [32][32]byte, index uint32, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.contract.Transact(opts, "claimAsset", smtProof, index, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0x2cffd02e.
//
// Solidity: function claimAsset(bytes32[32] smtProof, uint32 index, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) ClaimAsset(smtProof [32][32]byte, index uint32, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.ClaimAsset(&_PolygonZkEVMBridge.TransactOpts, smtProof, index, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0x2cffd02e.
//
// Solidity: function claimAsset(bytes32[32] smtProof, uint32 index, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeTransactorSession) ClaimAsset(smtProof [32][32]byte, index uint32, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.ClaimAsset(&_PolygonZkEVMBridge.TransactOpts, smtProof, index, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x2d2c9d94.
//
// Solidity: function claimMessage(bytes32[32] smtProof, uint32 index, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeTransactor) ClaimMessage(opts *bind.TransactOpts, smtProof [32][32]byte, index uint32, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.contract.Transact(opts, "claimMessage", smtProof, index, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x2d2c9d94.
//
// Solidity: function claimMessage(bytes32[32] smtProof, uint32 index, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) ClaimMessage(smtProof [32][32]byte, index uint32, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.ClaimMessage(&_PolygonZkEVMBridge.TransactOpts, smtProof, index, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x2d2c9d94.
//
// Solidity: function claimMessage(bytes32[32] smtProof, uint32 index, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeTransactorSession) ClaimMessage(smtProof [32][32]byte, index uint32, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.ClaimMessage(&_PolygonZkEVMBridge.TransactOpts, smtProof, index, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.DeactivateEmergencyState(&_PolygonZkEVMBridge.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.DeactivateEmergencyState(&_PolygonZkEVMBridge.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x647c576c.
//
// Solidity: function initialize(uint32 _networkID, address _globalExitRootManager, address _polygonZkEVMaddress) returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeTransactor) Initialize(opts *bind.TransactOpts, _networkID uint32, _globalExitRootManager common.Address, _polygonZkEVMaddress common.Address) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.contract.Transact(opts, "initialize", _networkID, _globalExitRootManager, _polygonZkEVMaddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x647c576c.
//
// Solidity: function initialize(uint32 _networkID, address _globalExitRootManager, address _polygonZkEVMaddress) returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) Initialize(_networkID uint32, _globalExitRootManager common.Address, _polygonZkEVMaddress common.Address) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.Initialize(&_PolygonZkEVMBridge.TransactOpts, _networkID, _globalExitRootManager, _polygonZkEVMaddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x647c576c.
//
// Solidity: function initialize(uint32 _networkID, address _globalExitRootManager, address _polygonZkEVMaddress) returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeTransactorSession) Initialize(_networkID uint32, _globalExitRootManager common.Address, _polygonZkEVMaddress common.Address) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.Initialize(&_PolygonZkEVMBridge.TransactOpts, _networkID, _globalExitRootManager, _polygonZkEVMaddress)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeTransactor) UpdateGlobalExitRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonZkEVMBridge.contract.Transact(opts, "updateGlobalExitRoot")
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.UpdateGlobalExitRoot(&_PolygonZkEVMBridge.TransactOpts)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeTransactorSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _PolygonZkEVMBridge.Contract.UpdateGlobalExitRoot(&_PolygonZkEVMBridge.TransactOpts)
}

// PolygonZkEVMBridgeBridgeEventIterator is returned from FilterBridgeEvent and is used to iterate over the raw logs and unpacked data for BridgeEvent events raised by the PolygonZkEVMBridge contract.
type PolygonZkEVMBridgeBridgeEventIterator struct {
	Event *PolygonZkEVMBridgeBridgeEvent // Event containing the contract specifics and raw log

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
func (it *PolygonZkEVMBridgeBridgeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonZkEVMBridgeBridgeEvent)
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
		it.Event = new(PolygonZkEVMBridgeBridgeEvent)
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
func (it *PolygonZkEVMBridgeBridgeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonZkEVMBridgeBridgeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonZkEVMBridgeBridgeEvent represents a BridgeEvent event raised by the PolygonZkEVMBridge contract.
type PolygonZkEVMBridgeBridgeEvent struct {
	LeafType           uint8
	OriginNetwork      uint32
	OriginAddress      common.Address
	DestinationNetwork uint32
	DestinationAddress common.Address
	Amount             *big.Int
	Metadata           []byte
	DepositCount       uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBridgeEvent is a free log retrieval operation binding the contract event 0x501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b.
//
// Solidity: event BridgeEvent(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata, uint32 depositCount)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeFilterer) FilterBridgeEvent(opts *bind.FilterOpts) (*PolygonZkEVMBridgeBridgeEventIterator, error) {

	logs, sub, err := _PolygonZkEVMBridge.contract.FilterLogs(opts, "BridgeEvent")
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMBridgeBridgeEventIterator{contract: _PolygonZkEVMBridge.contract, event: "BridgeEvent", logs: logs, sub: sub}, nil
}

// WatchBridgeEvent is a free log subscription operation binding the contract event 0x501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b.
//
// Solidity: event BridgeEvent(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata, uint32 depositCount)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeFilterer) WatchBridgeEvent(opts *bind.WatchOpts, sink chan<- *PolygonZkEVMBridgeBridgeEvent) (event.Subscription, error) {

	logs, sub, err := _PolygonZkEVMBridge.contract.WatchLogs(opts, "BridgeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonZkEVMBridgeBridgeEvent)
				if err := _PolygonZkEVMBridge.contract.UnpackLog(event, "BridgeEvent", log); err != nil {
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

// ParseBridgeEvent is a log parse operation binding the contract event 0x501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b.
//
// Solidity: event BridgeEvent(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata, uint32 depositCount)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeFilterer) ParseBridgeEvent(log types.Log) (*PolygonZkEVMBridgeBridgeEvent, error) {
	event := new(PolygonZkEVMBridgeBridgeEvent)
	if err := _PolygonZkEVMBridge.contract.UnpackLog(event, "BridgeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonZkEVMBridgeClaimEventIterator is returned from FilterClaimEvent and is used to iterate over the raw logs and unpacked data for ClaimEvent events raised by the PolygonZkEVMBridge contract.
type PolygonZkEVMBridgeClaimEventIterator struct {
	Event *PolygonZkEVMBridgeClaimEvent // Event containing the contract specifics and raw log

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
func (it *PolygonZkEVMBridgeClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonZkEVMBridgeClaimEvent)
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
		it.Event = new(PolygonZkEVMBridgeClaimEvent)
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
func (it *PolygonZkEVMBridgeClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonZkEVMBridgeClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonZkEVMBridgeClaimEvent represents a ClaimEvent event raised by the PolygonZkEVMBridge contract.
type PolygonZkEVMBridgeClaimEvent struct {
	Index              uint32
	OriginNetwork      uint32
	OriginAddress      common.Address
	DestinationAddress common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterClaimEvent is a free log retrieval operation binding the contract event 0x25308c93ceeed162da955b3f7ce3e3f93606579e40fb92029faa9efe27545983.
//
// Solidity: event ClaimEvent(uint32 index, uint32 originNetwork, address originAddress, address destinationAddress, uint256 amount)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeFilterer) FilterClaimEvent(opts *bind.FilterOpts) (*PolygonZkEVMBridgeClaimEventIterator, error) {

	logs, sub, err := _PolygonZkEVMBridge.contract.FilterLogs(opts, "ClaimEvent")
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMBridgeClaimEventIterator{contract: _PolygonZkEVMBridge.contract, event: "ClaimEvent", logs: logs, sub: sub}, nil
}

// WatchClaimEvent is a free log subscription operation binding the contract event 0x25308c93ceeed162da955b3f7ce3e3f93606579e40fb92029faa9efe27545983.
//
// Solidity: event ClaimEvent(uint32 index, uint32 originNetwork, address originAddress, address destinationAddress, uint256 amount)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeFilterer) WatchClaimEvent(opts *bind.WatchOpts, sink chan<- *PolygonZkEVMBridgeClaimEvent) (event.Subscription, error) {

	logs, sub, err := _PolygonZkEVMBridge.contract.WatchLogs(opts, "ClaimEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonZkEVMBridgeClaimEvent)
				if err := _PolygonZkEVMBridge.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
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

// ParseClaimEvent is a log parse operation binding the contract event 0x25308c93ceeed162da955b3f7ce3e3f93606579e40fb92029faa9efe27545983.
//
// Solidity: event ClaimEvent(uint32 index, uint32 originNetwork, address originAddress, address destinationAddress, uint256 amount)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeFilterer) ParseClaimEvent(log types.Log) (*PolygonZkEVMBridgeClaimEvent, error) {
	event := new(PolygonZkEVMBridgeClaimEvent)
	if err := _PolygonZkEVMBridge.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonZkEVMBridgeEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the PolygonZkEVMBridge contract.
type PolygonZkEVMBridgeEmergencyStateActivatedIterator struct {
	Event *PolygonZkEVMBridgeEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *PolygonZkEVMBridgeEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonZkEVMBridgeEmergencyStateActivated)
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
		it.Event = new(PolygonZkEVMBridgeEmergencyStateActivated)
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
func (it *PolygonZkEVMBridgeEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonZkEVMBridgeEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonZkEVMBridgeEmergencyStateActivated represents a EmergencyStateActivated event raised by the PolygonZkEVMBridge contract.
type PolygonZkEVMBridgeEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*PolygonZkEVMBridgeEmergencyStateActivatedIterator, error) {

	logs, sub, err := _PolygonZkEVMBridge.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMBridgeEmergencyStateActivatedIterator{contract: _PolygonZkEVMBridge.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *PolygonZkEVMBridgeEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _PolygonZkEVMBridge.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonZkEVMBridgeEmergencyStateActivated)
				if err := _PolygonZkEVMBridge.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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

// ParseEmergencyStateActivated is a log parse operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeFilterer) ParseEmergencyStateActivated(log types.Log) (*PolygonZkEVMBridgeEmergencyStateActivated, error) {
	event := new(PolygonZkEVMBridgeEmergencyStateActivated)
	if err := _PolygonZkEVMBridge.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonZkEVMBridgeEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the PolygonZkEVMBridge contract.
type PolygonZkEVMBridgeEmergencyStateDeactivatedIterator struct {
	Event *PolygonZkEVMBridgeEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *PolygonZkEVMBridgeEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonZkEVMBridgeEmergencyStateDeactivated)
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
		it.Event = new(PolygonZkEVMBridgeEmergencyStateDeactivated)
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
func (it *PolygonZkEVMBridgeEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonZkEVMBridgeEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonZkEVMBridgeEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the PolygonZkEVMBridge contract.
type PolygonZkEVMBridgeEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*PolygonZkEVMBridgeEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _PolygonZkEVMBridge.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMBridgeEmergencyStateDeactivatedIterator{contract: _PolygonZkEVMBridge.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *PolygonZkEVMBridgeEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _PolygonZkEVMBridge.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonZkEVMBridgeEmergencyStateDeactivated)
				if err := _PolygonZkEVMBridge.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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

// ParseEmergencyStateDeactivated is a log parse operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeFilterer) ParseEmergencyStateDeactivated(log types.Log) (*PolygonZkEVMBridgeEmergencyStateDeactivated, error) {
	event := new(PolygonZkEVMBridgeEmergencyStateDeactivated)
	if err := _PolygonZkEVMBridge.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonZkEVMBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PolygonZkEVMBridge contract.
type PolygonZkEVMBridgeInitializedIterator struct {
	Event *PolygonZkEVMBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonZkEVMBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonZkEVMBridgeInitialized)
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
		it.Event = new(PolygonZkEVMBridgeInitialized)
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
func (it *PolygonZkEVMBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonZkEVMBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonZkEVMBridgeInitialized represents a Initialized event raised by the PolygonZkEVMBridge contract.
type PolygonZkEVMBridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonZkEVMBridgeInitializedIterator, error) {

	logs, sub, err := _PolygonZkEVMBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMBridgeInitializedIterator{contract: _PolygonZkEVMBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonZkEVMBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _PolygonZkEVMBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonZkEVMBridgeInitialized)
				if err := _PolygonZkEVMBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeFilterer) ParseInitialized(log types.Log) (*PolygonZkEVMBridgeInitialized, error) {
	event := new(PolygonZkEVMBridgeInitialized)
	if err := _PolygonZkEVMBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonZkEVMBridgeNewWrappedTokenIterator is returned from FilterNewWrappedToken and is used to iterate over the raw logs and unpacked data for NewWrappedToken events raised by the PolygonZkEVMBridge contract.
type PolygonZkEVMBridgeNewWrappedTokenIterator struct {
	Event *PolygonZkEVMBridgeNewWrappedToken // Event containing the contract specifics and raw log

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
func (it *PolygonZkEVMBridgeNewWrappedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonZkEVMBridgeNewWrappedToken)
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
		it.Event = new(PolygonZkEVMBridgeNewWrappedToken)
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
func (it *PolygonZkEVMBridgeNewWrappedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonZkEVMBridgeNewWrappedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonZkEVMBridgeNewWrappedToken represents a NewWrappedToken event raised by the PolygonZkEVMBridge contract.
type PolygonZkEVMBridgeNewWrappedToken struct {
	OriginNetwork       uint32
	OriginTokenAddress  common.Address
	WrappedTokenAddress common.Address
	Metadata            []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewWrappedToken is a free log retrieval operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeFilterer) FilterNewWrappedToken(opts *bind.FilterOpts) (*PolygonZkEVMBridgeNewWrappedTokenIterator, error) {

	logs, sub, err := _PolygonZkEVMBridge.contract.FilterLogs(opts, "NewWrappedToken")
	if err != nil {
		return nil, err
	}
	return &PolygonZkEVMBridgeNewWrappedTokenIterator{contract: _PolygonZkEVMBridge.contract, event: "NewWrappedToken", logs: logs, sub: sub}, nil
}

// WatchNewWrappedToken is a free log subscription operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeFilterer) WatchNewWrappedToken(opts *bind.WatchOpts, sink chan<- *PolygonZkEVMBridgeNewWrappedToken) (event.Subscription, error) {

	logs, sub, err := _PolygonZkEVMBridge.contract.WatchLogs(opts, "NewWrappedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonZkEVMBridgeNewWrappedToken)
				if err := _PolygonZkEVMBridge.contract.UnpackLog(event, "NewWrappedToken", log); err != nil {
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

// ParseNewWrappedToken is a log parse operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_PolygonZkEVMBridge *PolygonZkEVMBridgeFilterer) ParseNewWrappedToken(log types.Log) (*PolygonZkEVMBridgeNewWrappedToken, error) {
	event := new(PolygonZkEVMBridgeNewWrappedToken)
	if err := _PolygonZkEVMBridge.contract.UnpackLog(event, "NewWrappedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
