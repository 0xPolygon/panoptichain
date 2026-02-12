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

// SPOLControllerFullNonceDetails is an auto generated low-level Go binding around an user-defined struct.
type SPOLControllerFullNonceDetails struct {
	ValidatorId    uint16
	Amount         *big.Int
	ValidatorNonce *big.Int
	Nonce          *big.Int
}

// SPOLControllerMetaData contains all meta data concerning the SPOLController contract.
var SPOLControllerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_polToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_maticToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_polygonMigration\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_sPOLToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stakeManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_DIVERGENCE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeValidators\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addValidator\",\"inputs\":[{\"name\":\"_validatorID\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authority\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"buySPOL\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_validator\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"buySPOL\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"buySPOLPermit\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_validator\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"buySPOLPermit\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"buySPOLWithDPOL\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_validatorOfDPOL\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"buySPOLWithDPOLPermit\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_validatorOfDPOL\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeFeeReceiver\",\"inputs\":[{\"name\":\"_newFeeReceiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeMaxDivergence\",\"inputs\":[{\"name\":\"_newDivergence\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeRewardFee\",\"inputs\":[{\"name\":\"_newFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cleanUpMaticPOL\",\"inputs\":[{\"name\":\"_validator\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"convertPOLtoSPOL\",\"inputs\":[{\"name\":\"_amountPOL\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"convertSPOLtoPOL\",\"inputs\":[{\"name\":\"_amountSPOL\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeReceiver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feedPOLBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMostOverfundedValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMostUnderfundedValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserOpenNonces\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structsPOLController.FullNonceDetails[]\",\"components\":[{\"name\":\"validatorId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"amount\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"validatorNonce\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalWithdrawNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_rewardFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_maxDivergence\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_authority\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isConsumingScheduledOp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maticToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxDivergence\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrateValidator\",\"inputs\":[{\"name\":\"_oldValidator\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_newValidator\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_restake\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseUserFunctions\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"polToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractERC20Permit\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"polygonMigration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPolygonMigration\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reloadAllActiveValidatorInfo\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reloadValidatorInfo\",\"inputs\":[{\"name\":\"_validator\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeValidator\",\"inputs\":[{\"name\":\"_removedValidator\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"restakeAllActiveValidators\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"restakeValidator\",\"inputs\":[{\"name\":\"_validator\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sPOLToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractsPOL\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sellSPOL\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_validator\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sellSPOL\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sellSPOLPermit\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_validator\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sellSPOLPermit\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAuthority\",\"inputs\":[{\"name\":\"newAuthority\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractStakeManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"takeFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totaldPOLBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalsPOLBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpauseUserFunctions\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateValidatorTargetShare\",\"inputs\":[{\"name\":\"_validatorID\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"_newTargetShare\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userNonces\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"_begin\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"_end\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorList\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validators\",\"inputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumsPOLController.ValidatorStatus\"},{\"name\":\"depositShare\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"index\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"validatorContract\",\"type\":\"address\",\"internalType\":\"contractValidatorShare\"},{\"name\":\"totalStaked\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawNonceDetails\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"validatorId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"amount\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"validatorNonce\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawPOL\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawPOL\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuthorityUpdated\",\"inputs\":[{\"name\":\"authority\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExchangeRateSnapshot\",\"inputs\":[{\"name\":\"totalsPOLSupply\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"totalbPOLBalance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeCollected\",\"inputs\":[{\"name\":\"feeReceiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feePOLAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feesPOLAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeReceiverChanged\",\"inputs\":[{\"name\":\"oldReceiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newReceiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaticTokensCleaned\",\"inputs\":[{\"name\":\"maticAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxDivergenceChanged\",\"inputs\":[{\"name\":\"oldDivergence\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"newDivergence\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"POLTokensCleaned\",\"inputs\":[{\"name\":\"validatorId\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"polAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"POLWithdrawn\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountPOL\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardFeeChanged\",\"inputs\":[{\"name\":\"oldFee\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newFee\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorAdded\",\"inputs\":[{\"name\":\"validatorId\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorMigrated\",\"inputs\":[{\"name\":\"oldValidator\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newValidator\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRemoved\",\"inputs\":[{\"name\":\"validatorId\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorTargetShareChanged\",\"inputs\":[{\"name\":\"validatorId\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newTargetShare\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"sPOLBurned\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountSPOL\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountPOL\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"sPOLMinted\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountPOL\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountSPOL\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessManagedInvalidAuthority\",\"inputs\":[{\"name\":\"authority\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessManagedRequiredDelay\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"AccessManagedUnauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AmountZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[{\"name\":\"validatorLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shareLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BuySharesMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"DPOLRestakeTransferFromFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositSharesTotalNotOneHundred\",\"inputs\":[{\"name\":\"totalPercent\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeTooLarge\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxAllowed\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"IncorrectValidatorShareExchangeRate\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxDivergenceTooLarge\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"NoNoncesReady\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NoOpenNonces\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NoUnlockedValidators\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughStake\",\"inputs\":[{\"name\":\"remaining\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorDepositShareNotZero\",\"inputs\":[{\"name\":\"validatorId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"depositShare\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"ValidatorNotActive\",\"inputs\":[{\"name\":\"validatorId\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"ValidatorNotDelegating\",\"inputs\":[{\"name\":\"validatorId\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"ValidatorNotInactive\",\"inputs\":[{\"name\":\"validatorId\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"ValidatorOverfunded\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ValidatorRewardsPending\",\"inputs\":[{\"name\":\"validatorId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rewards\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ValidatorSharesPending\",\"inputs\":[{\"name\":\"validatorId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ValidatorStillFunded\",\"inputs\":[{\"name\":\"validatorId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"totalStaked\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ValidatorUnderfunded\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// SPOLControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use SPOLControllerMetaData.ABI instead.
var SPOLControllerABI = SPOLControllerMetaData.ABI

// SPOLController is an auto generated Go binding around an Ethereum contract.
type SPOLController struct {
	SPOLControllerCaller     // Read-only binding to the contract
	SPOLControllerTransactor // Write-only binding to the contract
	SPOLControllerFilterer   // Log filterer for contract events
}

// SPOLControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SPOLControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SPOLControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SPOLControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SPOLControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SPOLControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SPOLControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SPOLControllerSession struct {
	Contract     *SPOLController   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SPOLControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SPOLControllerCallerSession struct {
	Contract *SPOLControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SPOLControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SPOLControllerTransactorSession struct {
	Contract     *SPOLControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SPOLControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SPOLControllerRaw struct {
	Contract *SPOLController // Generic contract binding to access the raw methods on
}

// SPOLControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SPOLControllerCallerRaw struct {
	Contract *SPOLControllerCaller // Generic read-only contract binding to access the raw methods on
}

// SPOLControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SPOLControllerTransactorRaw struct {
	Contract *SPOLControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSPOLController creates a new instance of SPOLController, bound to a specific deployed contract.
func NewSPOLController(address common.Address, backend bind.ContractBackend) (*SPOLController, error) {
	contract, err := bindSPOLController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SPOLController{SPOLControllerCaller: SPOLControllerCaller{contract: contract}, SPOLControllerTransactor: SPOLControllerTransactor{contract: contract}, SPOLControllerFilterer: SPOLControllerFilterer{contract: contract}}, nil
}

// NewSPOLControllerCaller creates a new read-only instance of SPOLController, bound to a specific deployed contract.
func NewSPOLControllerCaller(address common.Address, caller bind.ContractCaller) (*SPOLControllerCaller, error) {
	contract, err := bindSPOLController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SPOLControllerCaller{contract: contract}, nil
}

// NewSPOLControllerTransactor creates a new write-only instance of SPOLController, bound to a specific deployed contract.
func NewSPOLControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*SPOLControllerTransactor, error) {
	contract, err := bindSPOLController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SPOLControllerTransactor{contract: contract}, nil
}

// NewSPOLControllerFilterer creates a new log filterer instance of SPOLController, bound to a specific deployed contract.
func NewSPOLControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*SPOLControllerFilterer, error) {
	contract, err := bindSPOLController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SPOLControllerFilterer{contract: contract}, nil
}

// bindSPOLController binds a generic wrapper to an already deployed contract.
func bindSPOLController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SPOLControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SPOLController *SPOLControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SPOLController.Contract.SPOLControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SPOLController *SPOLControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SPOLController.Contract.SPOLControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SPOLController *SPOLControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SPOLController.Contract.SPOLControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SPOLController *SPOLControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SPOLController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SPOLController *SPOLControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SPOLController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SPOLController *SPOLControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SPOLController.Contract.contract.Transact(opts, method, params...)
}

// MAXDIVERGENCE is a free data retrieval call binding the contract method 0x8093cd64.
//
// Solidity: function MAX_DIVERGENCE() view returns(uint8)
func (_SPOLController *SPOLControllerCaller) MAXDIVERGENCE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "MAX_DIVERGENCE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXDIVERGENCE is a free data retrieval call binding the contract method 0x8093cd64.
//
// Solidity: function MAX_DIVERGENCE() view returns(uint8)
func (_SPOLController *SPOLControllerSession) MAXDIVERGENCE() (uint8, error) {
	return _SPOLController.Contract.MAXDIVERGENCE(&_SPOLController.CallOpts)
}

// MAXDIVERGENCE is a free data retrieval call binding the contract method 0x8093cd64.
//
// Solidity: function MAX_DIVERGENCE() view returns(uint8)
func (_SPOLController *SPOLControllerCallerSession) MAXDIVERGENCE() (uint8, error) {
	return _SPOLController.Contract.MAXDIVERGENCE(&_SPOLController.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint16)
func (_SPOLController *SPOLControllerCaller) MAXFEE(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint16)
func (_SPOLController *SPOLControllerSession) MAXFEE() (uint16, error) {
	return _SPOLController.Contract.MAXFEE(&_SPOLController.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint16)
func (_SPOLController *SPOLControllerCallerSession) MAXFEE() (uint16, error) {
	return _SPOLController.Contract.MAXFEE(&_SPOLController.CallOpts)
}

// ActiveValidators is a free data retrieval call binding the contract method 0x14f64c78.
//
// Solidity: function activeValidators(uint256 ) view returns(uint16)
func (_SPOLController *SPOLControllerCaller) ActiveValidators(opts *bind.CallOpts, arg0 *big.Int) (uint16, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "activeValidators", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ActiveValidators is a free data retrieval call binding the contract method 0x14f64c78.
//
// Solidity: function activeValidators(uint256 ) view returns(uint16)
func (_SPOLController *SPOLControllerSession) ActiveValidators(arg0 *big.Int) (uint16, error) {
	return _SPOLController.Contract.ActiveValidators(&_SPOLController.CallOpts, arg0)
}

// ActiveValidators is a free data retrieval call binding the contract method 0x14f64c78.
//
// Solidity: function activeValidators(uint256 ) view returns(uint16)
func (_SPOLController *SPOLControllerCallerSession) ActiveValidators(arg0 *big.Int) (uint16, error) {
	return _SPOLController.Contract.ActiveValidators(&_SPOLController.CallOpts, arg0)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_SPOLController *SPOLControllerCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_SPOLController *SPOLControllerSession) Authority() (common.Address, error) {
	return _SPOLController.Contract.Authority(&_SPOLController.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_SPOLController *SPOLControllerCallerSession) Authority() (common.Address, error) {
	return _SPOLController.Contract.Authority(&_SPOLController.CallOpts)
}

// ConvertPOLtoSPOL is a free data retrieval call binding the contract method 0xc356a582.
//
// Solidity: function convertPOLtoSPOL(uint256 _amountPOL) view returns(uint256)
func (_SPOLController *SPOLControllerCaller) ConvertPOLtoSPOL(opts *bind.CallOpts, _amountPOL *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "convertPOLtoSPOL", _amountPOL)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertPOLtoSPOL is a free data retrieval call binding the contract method 0xc356a582.
//
// Solidity: function convertPOLtoSPOL(uint256 _amountPOL) view returns(uint256)
func (_SPOLController *SPOLControllerSession) ConvertPOLtoSPOL(_amountPOL *big.Int) (*big.Int, error) {
	return _SPOLController.Contract.ConvertPOLtoSPOL(&_SPOLController.CallOpts, _amountPOL)
}

// ConvertPOLtoSPOL is a free data retrieval call binding the contract method 0xc356a582.
//
// Solidity: function convertPOLtoSPOL(uint256 _amountPOL) view returns(uint256)
func (_SPOLController *SPOLControllerCallerSession) ConvertPOLtoSPOL(_amountPOL *big.Int) (*big.Int, error) {
	return _SPOLController.Contract.ConvertPOLtoSPOL(&_SPOLController.CallOpts, _amountPOL)
}

// ConvertSPOLtoPOL is a free data retrieval call binding the contract method 0xff8aaf7a.
//
// Solidity: function convertSPOLtoPOL(uint256 _amountSPOL) view returns(uint256)
func (_SPOLController *SPOLControllerCaller) ConvertSPOLtoPOL(opts *bind.CallOpts, _amountSPOL *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "convertSPOLtoPOL", _amountSPOL)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertSPOLtoPOL is a free data retrieval call binding the contract method 0xff8aaf7a.
//
// Solidity: function convertSPOLtoPOL(uint256 _amountSPOL) view returns(uint256)
func (_SPOLController *SPOLControllerSession) ConvertSPOLtoPOL(_amountSPOL *big.Int) (*big.Int, error) {
	return _SPOLController.Contract.ConvertSPOLtoPOL(&_SPOLController.CallOpts, _amountSPOL)
}

// ConvertSPOLtoPOL is a free data retrieval call binding the contract method 0xff8aaf7a.
//
// Solidity: function convertSPOLtoPOL(uint256 _amountSPOL) view returns(uint256)
func (_SPOLController *SPOLControllerCallerSession) ConvertSPOLtoPOL(_amountSPOL *big.Int) (*big.Int, error) {
	return _SPOLController.Contract.ConvertSPOLtoPOL(&_SPOLController.CallOpts, _amountSPOL)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_SPOLController *SPOLControllerCaller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "feeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_SPOLController *SPOLControllerSession) FeeReceiver() (common.Address, error) {
	return _SPOLController.Contract.FeeReceiver(&_SPOLController.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_SPOLController *SPOLControllerCallerSession) FeeReceiver() (common.Address, error) {
	return _SPOLController.Contract.FeeReceiver(&_SPOLController.CallOpts)
}

// FeedPOLBalance is a free data retrieval call binding the contract method 0x3656abde.
//
// Solidity: function feedPOLBalance() view returns(uint256)
func (_SPOLController *SPOLControllerCaller) FeedPOLBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "feedPOLBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeedPOLBalance is a free data retrieval call binding the contract method 0x3656abde.
//
// Solidity: function feedPOLBalance() view returns(uint256)
func (_SPOLController *SPOLControllerSession) FeedPOLBalance() (*big.Int, error) {
	return _SPOLController.Contract.FeedPOLBalance(&_SPOLController.CallOpts)
}

// FeedPOLBalance is a free data retrieval call binding the contract method 0x3656abde.
//
// Solidity: function feedPOLBalance() view returns(uint256)
func (_SPOLController *SPOLControllerCallerSession) FeedPOLBalance() (*big.Int, error) {
	return _SPOLController.Contract.FeedPOLBalance(&_SPOLController.CallOpts)
}

// GetMostOverfundedValidator is a free data retrieval call binding the contract method 0x374abe5e.
//
// Solidity: function getMostOverfundedValidator() view returns(uint16, uint256)
func (_SPOLController *SPOLControllerCaller) GetMostOverfundedValidator(opts *bind.CallOpts) (uint16, *big.Int, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "getMostOverfundedValidator")

	if err != nil {
		return *new(uint16), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetMostOverfundedValidator is a free data retrieval call binding the contract method 0x374abe5e.
//
// Solidity: function getMostOverfundedValidator() view returns(uint16, uint256)
func (_SPOLController *SPOLControllerSession) GetMostOverfundedValidator() (uint16, *big.Int, error) {
	return _SPOLController.Contract.GetMostOverfundedValidator(&_SPOLController.CallOpts)
}

// GetMostOverfundedValidator is a free data retrieval call binding the contract method 0x374abe5e.
//
// Solidity: function getMostOverfundedValidator() view returns(uint16, uint256)
func (_SPOLController *SPOLControllerCallerSession) GetMostOverfundedValidator() (uint16, *big.Int, error) {
	return _SPOLController.Contract.GetMostOverfundedValidator(&_SPOLController.CallOpts)
}

// GetMostUnderfundedValidator is a free data retrieval call binding the contract method 0xa5c2afc7.
//
// Solidity: function getMostUnderfundedValidator() view returns(uint16, uint256)
func (_SPOLController *SPOLControllerCaller) GetMostUnderfundedValidator(opts *bind.CallOpts) (uint16, *big.Int, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "getMostUnderfundedValidator")

	if err != nil {
		return *new(uint16), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetMostUnderfundedValidator is a free data retrieval call binding the contract method 0xa5c2afc7.
//
// Solidity: function getMostUnderfundedValidator() view returns(uint16, uint256)
func (_SPOLController *SPOLControllerSession) GetMostUnderfundedValidator() (uint16, *big.Int, error) {
	return _SPOLController.Contract.GetMostUnderfundedValidator(&_SPOLController.CallOpts)
}

// GetMostUnderfundedValidator is a free data retrieval call binding the contract method 0xa5c2afc7.
//
// Solidity: function getMostUnderfundedValidator() view returns(uint16, uint256)
func (_SPOLController *SPOLControllerCallerSession) GetMostUnderfundedValidator() (uint16, *big.Int, error) {
	return _SPOLController.Contract.GetMostUnderfundedValidator(&_SPOLController.CallOpts)
}

// GetUserOpenNonces is a free data retrieval call binding the contract method 0x58c51245.
//
// Solidity: function getUserOpenNonces(address _user) view returns((uint16,uint128,uint96,uint256)[])
func (_SPOLController *SPOLControllerCaller) GetUserOpenNonces(opts *bind.CallOpts, _user common.Address) ([]SPOLControllerFullNonceDetails, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "getUserOpenNonces", _user)

	if err != nil {
		return *new([]SPOLControllerFullNonceDetails), err
	}

	out0 := *abi.ConvertType(out[0], new([]SPOLControllerFullNonceDetails)).(*[]SPOLControllerFullNonceDetails)

	return out0, err

}

// GetUserOpenNonces is a free data retrieval call binding the contract method 0x58c51245.
//
// Solidity: function getUserOpenNonces(address _user) view returns((uint16,uint128,uint96,uint256)[])
func (_SPOLController *SPOLControllerSession) GetUserOpenNonces(_user common.Address) ([]SPOLControllerFullNonceDetails, error) {
	return _SPOLController.Contract.GetUserOpenNonces(&_SPOLController.CallOpts, _user)
}

// GetUserOpenNonces is a free data retrieval call binding the contract method 0x58c51245.
//
// Solidity: function getUserOpenNonces(address _user) view returns((uint16,uint128,uint96,uint256)[])
func (_SPOLController *SPOLControllerCallerSession) GetUserOpenNonces(_user common.Address) ([]SPOLControllerFullNonceDetails, error) {
	return _SPOLController.Contract.GetUserOpenNonces(&_SPOLController.CallOpts, _user)
}

// GlobalWithdrawNonce is a free data retrieval call binding the contract method 0xa4858b5a.
//
// Solidity: function globalWithdrawNonce() view returns(uint256)
func (_SPOLController *SPOLControllerCaller) GlobalWithdrawNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "globalWithdrawNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalWithdrawNonce is a free data retrieval call binding the contract method 0xa4858b5a.
//
// Solidity: function globalWithdrawNonce() view returns(uint256)
func (_SPOLController *SPOLControllerSession) GlobalWithdrawNonce() (*big.Int, error) {
	return _SPOLController.Contract.GlobalWithdrawNonce(&_SPOLController.CallOpts)
}

// GlobalWithdrawNonce is a free data retrieval call binding the contract method 0xa4858b5a.
//
// Solidity: function globalWithdrawNonce() view returns(uint256)
func (_SPOLController *SPOLControllerCallerSession) GlobalWithdrawNonce() (*big.Int, error) {
	return _SPOLController.Contract.GlobalWithdrawNonce(&_SPOLController.CallOpts)
}

// IsConsumingScheduledOp is a free data retrieval call binding the contract method 0x8fb36037.
//
// Solidity: function isConsumingScheduledOp() view returns(bytes4)
func (_SPOLController *SPOLControllerCaller) IsConsumingScheduledOp(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "isConsumingScheduledOp")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsConsumingScheduledOp is a free data retrieval call binding the contract method 0x8fb36037.
//
// Solidity: function isConsumingScheduledOp() view returns(bytes4)
func (_SPOLController *SPOLControllerSession) IsConsumingScheduledOp() ([4]byte, error) {
	return _SPOLController.Contract.IsConsumingScheduledOp(&_SPOLController.CallOpts)
}

// IsConsumingScheduledOp is a free data retrieval call binding the contract method 0x8fb36037.
//
// Solidity: function isConsumingScheduledOp() view returns(bytes4)
func (_SPOLController *SPOLControllerCallerSession) IsConsumingScheduledOp() ([4]byte, error) {
	return _SPOLController.Contract.IsConsumingScheduledOp(&_SPOLController.CallOpts)
}

// MaticToken is a free data retrieval call binding the contract method 0xdc354296.
//
// Solidity: function maticToken() view returns(address)
func (_SPOLController *SPOLControllerCaller) MaticToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "maticToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MaticToken is a free data retrieval call binding the contract method 0xdc354296.
//
// Solidity: function maticToken() view returns(address)
func (_SPOLController *SPOLControllerSession) MaticToken() (common.Address, error) {
	return _SPOLController.Contract.MaticToken(&_SPOLController.CallOpts)
}

// MaticToken is a free data retrieval call binding the contract method 0xdc354296.
//
// Solidity: function maticToken() view returns(address)
func (_SPOLController *SPOLControllerCallerSession) MaticToken() (common.Address, error) {
	return _SPOLController.Contract.MaticToken(&_SPOLController.CallOpts)
}

// MaxDivergence is a free data retrieval call binding the contract method 0x3de5d7d3.
//
// Solidity: function maxDivergence() view returns(uint8)
func (_SPOLController *SPOLControllerCaller) MaxDivergence(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "maxDivergence")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MaxDivergence is a free data retrieval call binding the contract method 0x3de5d7d3.
//
// Solidity: function maxDivergence() view returns(uint8)
func (_SPOLController *SPOLControllerSession) MaxDivergence() (uint8, error) {
	return _SPOLController.Contract.MaxDivergence(&_SPOLController.CallOpts)
}

// MaxDivergence is a free data retrieval call binding the contract method 0x3de5d7d3.
//
// Solidity: function maxDivergence() view returns(uint8)
func (_SPOLController *SPOLControllerCallerSession) MaxDivergence() (uint8, error) {
	return _SPOLController.Contract.MaxDivergence(&_SPOLController.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SPOLController *SPOLControllerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SPOLController *SPOLControllerSession) Paused() (bool, error) {
	return _SPOLController.Contract.Paused(&_SPOLController.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SPOLController *SPOLControllerCallerSession) Paused() (bool, error) {
	return _SPOLController.Contract.Paused(&_SPOLController.CallOpts)
}

// PolToken is a free data retrieval call binding the contract method 0xb6658d07.
//
// Solidity: function polToken() view returns(address)
func (_SPOLController *SPOLControllerCaller) PolToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "polToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolToken is a free data retrieval call binding the contract method 0xb6658d07.
//
// Solidity: function polToken() view returns(address)
func (_SPOLController *SPOLControllerSession) PolToken() (common.Address, error) {
	return _SPOLController.Contract.PolToken(&_SPOLController.CallOpts)
}

// PolToken is a free data retrieval call binding the contract method 0xb6658d07.
//
// Solidity: function polToken() view returns(address)
func (_SPOLController *SPOLControllerCallerSession) PolToken() (common.Address, error) {
	return _SPOLController.Contract.PolToken(&_SPOLController.CallOpts)
}

// PolygonMigration is a free data retrieval call binding the contract method 0x16ceca3c.
//
// Solidity: function polygonMigration() view returns(address)
func (_SPOLController *SPOLControllerCaller) PolygonMigration(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "polygonMigration")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolygonMigration is a free data retrieval call binding the contract method 0x16ceca3c.
//
// Solidity: function polygonMigration() view returns(address)
func (_SPOLController *SPOLControllerSession) PolygonMigration() (common.Address, error) {
	return _SPOLController.Contract.PolygonMigration(&_SPOLController.CallOpts)
}

// PolygonMigration is a free data retrieval call binding the contract method 0x16ceca3c.
//
// Solidity: function polygonMigration() view returns(address)
func (_SPOLController *SPOLControllerCallerSession) PolygonMigration() (common.Address, error) {
	return _SPOLController.Contract.PolygonMigration(&_SPOLController.CallOpts)
}

// RewardFee is a free data retrieval call binding the contract method 0x8b424267.
//
// Solidity: function rewardFee() view returns(uint16)
func (_SPOLController *SPOLControllerCaller) RewardFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "rewardFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RewardFee is a free data retrieval call binding the contract method 0x8b424267.
//
// Solidity: function rewardFee() view returns(uint16)
func (_SPOLController *SPOLControllerSession) RewardFee() (uint16, error) {
	return _SPOLController.Contract.RewardFee(&_SPOLController.CallOpts)
}

// RewardFee is a free data retrieval call binding the contract method 0x8b424267.
//
// Solidity: function rewardFee() view returns(uint16)
func (_SPOLController *SPOLControllerCallerSession) RewardFee() (uint16, error) {
	return _SPOLController.Contract.RewardFee(&_SPOLController.CallOpts)
}

// SPOLToken is a free data retrieval call binding the contract method 0x874a69e8.
//
// Solidity: function sPOLToken() view returns(address)
func (_SPOLController *SPOLControllerCaller) SPOLToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "sPOLToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SPOLToken is a free data retrieval call binding the contract method 0x874a69e8.
//
// Solidity: function sPOLToken() view returns(address)
func (_SPOLController *SPOLControllerSession) SPOLToken() (common.Address, error) {
	return _SPOLController.Contract.SPOLToken(&_SPOLController.CallOpts)
}

// SPOLToken is a free data retrieval call binding the contract method 0x874a69e8.
//
// Solidity: function sPOLToken() view returns(address)
func (_SPOLController *SPOLControllerCallerSession) SPOLToken() (common.Address, error) {
	return _SPOLController.Contract.SPOLToken(&_SPOLController.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_SPOLController *SPOLControllerCaller) StakeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "stakeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_SPOLController *SPOLControllerSession) StakeManager() (common.Address, error) {
	return _SPOLController.Contract.StakeManager(&_SPOLController.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_SPOLController *SPOLControllerCallerSession) StakeManager() (common.Address, error) {
	return _SPOLController.Contract.StakeManager(&_SPOLController.CallOpts)
}

// TotaldPOLBalance is a free data retrieval call binding the contract method 0x5d3a8f06.
//
// Solidity: function totaldPOLBalance() view returns(uint256)
func (_SPOLController *SPOLControllerCaller) TotaldPOLBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "totaldPOLBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotaldPOLBalance is a free data retrieval call binding the contract method 0x5d3a8f06.
//
// Solidity: function totaldPOLBalance() view returns(uint256)
func (_SPOLController *SPOLControllerSession) TotaldPOLBalance() (*big.Int, error) {
	return _SPOLController.Contract.TotaldPOLBalance(&_SPOLController.CallOpts)
}

// TotaldPOLBalance is a free data retrieval call binding the contract method 0x5d3a8f06.
//
// Solidity: function totaldPOLBalance() view returns(uint256)
func (_SPOLController *SPOLControllerCallerSession) TotaldPOLBalance() (*big.Int, error) {
	return _SPOLController.Contract.TotaldPOLBalance(&_SPOLController.CallOpts)
}

// TotalsPOLBalance is a free data retrieval call binding the contract method 0x94b3e542.
//
// Solidity: function totalsPOLBalance() view returns(uint256)
func (_SPOLController *SPOLControllerCaller) TotalsPOLBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "totalsPOLBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalsPOLBalance is a free data retrieval call binding the contract method 0x94b3e542.
//
// Solidity: function totalsPOLBalance() view returns(uint256)
func (_SPOLController *SPOLControllerSession) TotalsPOLBalance() (*big.Int, error) {
	return _SPOLController.Contract.TotalsPOLBalance(&_SPOLController.CallOpts)
}

// TotalsPOLBalance is a free data retrieval call binding the contract method 0x94b3e542.
//
// Solidity: function totalsPOLBalance() view returns(uint256)
func (_SPOLController *SPOLControllerCallerSession) TotalsPOLBalance() (*big.Int, error) {
	return _SPOLController.Contract.TotalsPOLBalance(&_SPOLController.CallOpts)
}

// UserNonces is a free data retrieval call binding the contract method 0x2f7801f4.
//
// Solidity: function userNonces(address ) view returns(uint128 _begin, uint128 _end)
func (_SPOLController *SPOLControllerCaller) UserNonces(opts *bind.CallOpts, arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "userNonces", arg0)

	outstruct := new(struct {
		Begin *big.Int
		End   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Begin = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserNonces is a free data retrieval call binding the contract method 0x2f7801f4.
//
// Solidity: function userNonces(address ) view returns(uint128 _begin, uint128 _end)
func (_SPOLController *SPOLControllerSession) UserNonces(arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _SPOLController.Contract.UserNonces(&_SPOLController.CallOpts, arg0)
}

// UserNonces is a free data retrieval call binding the contract method 0x2f7801f4.
//
// Solidity: function userNonces(address ) view returns(uint128 _begin, uint128 _end)
func (_SPOLController *SPOLControllerCallerSession) UserNonces(arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _SPOLController.Contract.UserNonces(&_SPOLController.CallOpts, arg0)
}

// ValidatorList is a free data retrieval call binding the contract method 0xb048e056.
//
// Solidity: function validatorList(uint256 ) view returns(uint16)
func (_SPOLController *SPOLControllerCaller) ValidatorList(opts *bind.CallOpts, arg0 *big.Int) (uint16, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "validatorList", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ValidatorList is a free data retrieval call binding the contract method 0xb048e056.
//
// Solidity: function validatorList(uint256 ) view returns(uint16)
func (_SPOLController *SPOLControllerSession) ValidatorList(arg0 *big.Int) (uint16, error) {
	return _SPOLController.Contract.ValidatorList(&_SPOLController.CallOpts, arg0)
}

// ValidatorList is a free data retrieval call binding the contract method 0xb048e056.
//
// Solidity: function validatorList(uint256 ) view returns(uint16)
func (_SPOLController *SPOLControllerCallerSession) ValidatorList(arg0 *big.Int) (uint16, error) {
	return _SPOLController.Contract.ValidatorList(&_SPOLController.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x69f7fa1f.
//
// Solidity: function validators(uint16 ) view returns(uint8 status, uint8 depositShare, uint16 index, address validatorContract, uint256 totalStaked)
func (_SPOLController *SPOLControllerCaller) Validators(opts *bind.CallOpts, arg0 uint16) (struct {
	Status            uint8
	DepositShare      uint8
	Index             uint16
	ValidatorContract common.Address
	TotalStaked       *big.Int
}, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		Status            uint8
		DepositShare      uint8
		Index             uint16
		ValidatorContract common.Address
		TotalStaked       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.DepositShare = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Index = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.ValidatorContract = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.TotalStaked = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0x69f7fa1f.
//
// Solidity: function validators(uint16 ) view returns(uint8 status, uint8 depositShare, uint16 index, address validatorContract, uint256 totalStaked)
func (_SPOLController *SPOLControllerSession) Validators(arg0 uint16) (struct {
	Status            uint8
	DepositShare      uint8
	Index             uint16
	ValidatorContract common.Address
	TotalStaked       *big.Int
}, error) {
	return _SPOLController.Contract.Validators(&_SPOLController.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x69f7fa1f.
//
// Solidity: function validators(uint16 ) view returns(uint8 status, uint8 depositShare, uint16 index, address validatorContract, uint256 totalStaked)
func (_SPOLController *SPOLControllerCallerSession) Validators(arg0 uint16) (struct {
	Status            uint8
	DepositShare      uint8
	Index             uint16
	ValidatorContract common.Address
	TotalStaked       *big.Int
}, error) {
	return _SPOLController.Contract.Validators(&_SPOLController.CallOpts, arg0)
}

// WithdrawNonceDetails is a free data retrieval call binding the contract method 0x795c95d5.
//
// Solidity: function withdrawNonceDetails(uint256 ) view returns(uint16 validatorId, uint128 amount, uint96 validatorNonce)
func (_SPOLController *SPOLControllerCaller) WithdrawNonceDetails(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ValidatorId    uint16
	Amount         *big.Int
	ValidatorNonce *big.Int
}, error) {
	var out []interface{}
	err := _SPOLController.contract.Call(opts, &out, "withdrawNonceDetails", arg0)

	outstruct := new(struct {
		ValidatorId    uint16
		Amount         *big.Int
		ValidatorNonce *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValidatorId = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ValidatorNonce = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// WithdrawNonceDetails is a free data retrieval call binding the contract method 0x795c95d5.
//
// Solidity: function withdrawNonceDetails(uint256 ) view returns(uint16 validatorId, uint128 amount, uint96 validatorNonce)
func (_SPOLController *SPOLControllerSession) WithdrawNonceDetails(arg0 *big.Int) (struct {
	ValidatorId    uint16
	Amount         *big.Int
	ValidatorNonce *big.Int
}, error) {
	return _SPOLController.Contract.WithdrawNonceDetails(&_SPOLController.CallOpts, arg0)
}

// WithdrawNonceDetails is a free data retrieval call binding the contract method 0x795c95d5.
//
// Solidity: function withdrawNonceDetails(uint256 ) view returns(uint16 validatorId, uint128 amount, uint96 validatorNonce)
func (_SPOLController *SPOLControllerCallerSession) WithdrawNonceDetails(arg0 *big.Int) (struct {
	ValidatorId    uint16
	Amount         *big.Int
	ValidatorNonce *big.Int
}, error) {
	return _SPOLController.Contract.WithdrawNonceDetails(&_SPOLController.CallOpts, arg0)
}

// AddValidator is a paid mutator transaction binding the contract method 0xc7e6d61c.
//
// Solidity: function addValidator(uint16 _validatorID) returns()
func (_SPOLController *SPOLControllerTransactor) AddValidator(opts *bind.TransactOpts, _validatorID uint16) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "addValidator", _validatorID)
}

// AddValidator is a paid mutator transaction binding the contract method 0xc7e6d61c.
//
// Solidity: function addValidator(uint16 _validatorID) returns()
func (_SPOLController *SPOLControllerSession) AddValidator(_validatorID uint16) (*types.Transaction, error) {
	return _SPOLController.Contract.AddValidator(&_SPOLController.TransactOpts, _validatorID)
}

// AddValidator is a paid mutator transaction binding the contract method 0xc7e6d61c.
//
// Solidity: function addValidator(uint16 _validatorID) returns()
func (_SPOLController *SPOLControllerTransactorSession) AddValidator(_validatorID uint16) (*types.Transaction, error) {
	return _SPOLController.Contract.AddValidator(&_SPOLController.TransactOpts, _validatorID)
}

// BuySPOL is a paid mutator transaction binding the contract method 0xb6722163.
//
// Solidity: function buySPOL(uint256 _amount, uint16 _validator) returns(uint256)
func (_SPOLController *SPOLControllerTransactor) BuySPOL(opts *bind.TransactOpts, _amount *big.Int, _validator uint16) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "buySPOL", _amount, _validator)
}

// BuySPOL is a paid mutator transaction binding the contract method 0xb6722163.
//
// Solidity: function buySPOL(uint256 _amount, uint16 _validator) returns(uint256)
func (_SPOLController *SPOLControllerSession) BuySPOL(_amount *big.Int, _validator uint16) (*types.Transaction, error) {
	return _SPOLController.Contract.BuySPOL(&_SPOLController.TransactOpts, _amount, _validator)
}

// BuySPOL is a paid mutator transaction binding the contract method 0xb6722163.
//
// Solidity: function buySPOL(uint256 _amount, uint16 _validator) returns(uint256)
func (_SPOLController *SPOLControllerTransactorSession) BuySPOL(_amount *big.Int, _validator uint16) (*types.Transaction, error) {
	return _SPOLController.Contract.BuySPOL(&_SPOLController.TransactOpts, _amount, _validator)
}

// BuySPOL0 is a paid mutator transaction binding the contract method 0xbb7914a3.
//
// Solidity: function buySPOL(uint256 _amount) returns(uint256)
func (_SPOLController *SPOLControllerTransactor) BuySPOL0(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "buySPOL0", _amount)
}

// BuySPOL0 is a paid mutator transaction binding the contract method 0xbb7914a3.
//
// Solidity: function buySPOL(uint256 _amount) returns(uint256)
func (_SPOLController *SPOLControllerSession) BuySPOL0(_amount *big.Int) (*types.Transaction, error) {
	return _SPOLController.Contract.BuySPOL0(&_SPOLController.TransactOpts, _amount)
}

// BuySPOL0 is a paid mutator transaction binding the contract method 0xbb7914a3.
//
// Solidity: function buySPOL(uint256 _amount) returns(uint256)
func (_SPOLController *SPOLControllerTransactorSession) BuySPOL0(_amount *big.Int) (*types.Transaction, error) {
	return _SPOLController.Contract.BuySPOL0(&_SPOLController.TransactOpts, _amount)
}

// BuySPOLPermit is a paid mutator transaction binding the contract method 0x27bbe03d.
//
// Solidity: function buySPOLPermit(uint256 _amount, uint16 _validator, address _user, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256)
func (_SPOLController *SPOLControllerTransactor) BuySPOLPermit(opts *bind.TransactOpts, _amount *big.Int, _validator uint16, _user common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "buySPOLPermit", _amount, _validator, _user, _deadline, _v, _r, _s)
}

// BuySPOLPermit is a paid mutator transaction binding the contract method 0x27bbe03d.
//
// Solidity: function buySPOLPermit(uint256 _amount, uint16 _validator, address _user, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256)
func (_SPOLController *SPOLControllerSession) BuySPOLPermit(_amount *big.Int, _validator uint16, _user common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SPOLController.Contract.BuySPOLPermit(&_SPOLController.TransactOpts, _amount, _validator, _user, _deadline, _v, _r, _s)
}

// BuySPOLPermit is a paid mutator transaction binding the contract method 0x27bbe03d.
//
// Solidity: function buySPOLPermit(uint256 _amount, uint16 _validator, address _user, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256)
func (_SPOLController *SPOLControllerTransactorSession) BuySPOLPermit(_amount *big.Int, _validator uint16, _user common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SPOLController.Contract.BuySPOLPermit(&_SPOLController.TransactOpts, _amount, _validator, _user, _deadline, _v, _r, _s)
}

// BuySPOLPermit0 is a paid mutator transaction binding the contract method 0x4d4778a1.
//
// Solidity: function buySPOLPermit(uint256 _amount, address _user, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256)
func (_SPOLController *SPOLControllerTransactor) BuySPOLPermit0(opts *bind.TransactOpts, _amount *big.Int, _user common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "buySPOLPermit0", _amount, _user, _deadline, _v, _r, _s)
}

// BuySPOLPermit0 is a paid mutator transaction binding the contract method 0x4d4778a1.
//
// Solidity: function buySPOLPermit(uint256 _amount, address _user, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256)
func (_SPOLController *SPOLControllerSession) BuySPOLPermit0(_amount *big.Int, _user common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SPOLController.Contract.BuySPOLPermit0(&_SPOLController.TransactOpts, _amount, _user, _deadline, _v, _r, _s)
}

// BuySPOLPermit0 is a paid mutator transaction binding the contract method 0x4d4778a1.
//
// Solidity: function buySPOLPermit(uint256 _amount, address _user, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256)
func (_SPOLController *SPOLControllerTransactorSession) BuySPOLPermit0(_amount *big.Int, _user common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SPOLController.Contract.BuySPOLPermit0(&_SPOLController.TransactOpts, _amount, _user, _deadline, _v, _r, _s)
}

// BuySPOLWithDPOL is a paid mutator transaction binding the contract method 0xf57ccae9.
//
// Solidity: function buySPOLWithDPOL(uint256 _amount, uint16 _validatorOfDPOL) returns(uint256)
func (_SPOLController *SPOLControllerTransactor) BuySPOLWithDPOL(opts *bind.TransactOpts, _amount *big.Int, _validatorOfDPOL uint16) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "buySPOLWithDPOL", _amount, _validatorOfDPOL)
}

// BuySPOLWithDPOL is a paid mutator transaction binding the contract method 0xf57ccae9.
//
// Solidity: function buySPOLWithDPOL(uint256 _amount, uint16 _validatorOfDPOL) returns(uint256)
func (_SPOLController *SPOLControllerSession) BuySPOLWithDPOL(_amount *big.Int, _validatorOfDPOL uint16) (*types.Transaction, error) {
	return _SPOLController.Contract.BuySPOLWithDPOL(&_SPOLController.TransactOpts, _amount, _validatorOfDPOL)
}

// BuySPOLWithDPOL is a paid mutator transaction binding the contract method 0xf57ccae9.
//
// Solidity: function buySPOLWithDPOL(uint256 _amount, uint16 _validatorOfDPOL) returns(uint256)
func (_SPOLController *SPOLControllerTransactorSession) BuySPOLWithDPOL(_amount *big.Int, _validatorOfDPOL uint16) (*types.Transaction, error) {
	return _SPOLController.Contract.BuySPOLWithDPOL(&_SPOLController.TransactOpts, _amount, _validatorOfDPOL)
}

// BuySPOLWithDPOLPermit is a paid mutator transaction binding the contract method 0x44354ade.
//
// Solidity: function buySPOLWithDPOLPermit(uint256 _amount, uint16 _validatorOfDPOL, address _user, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256)
func (_SPOLController *SPOLControllerTransactor) BuySPOLWithDPOLPermit(opts *bind.TransactOpts, _amount *big.Int, _validatorOfDPOL uint16, _user common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "buySPOLWithDPOLPermit", _amount, _validatorOfDPOL, _user, _deadline, _v, _r, _s)
}

// BuySPOLWithDPOLPermit is a paid mutator transaction binding the contract method 0x44354ade.
//
// Solidity: function buySPOLWithDPOLPermit(uint256 _amount, uint16 _validatorOfDPOL, address _user, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256)
func (_SPOLController *SPOLControllerSession) BuySPOLWithDPOLPermit(_amount *big.Int, _validatorOfDPOL uint16, _user common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SPOLController.Contract.BuySPOLWithDPOLPermit(&_SPOLController.TransactOpts, _amount, _validatorOfDPOL, _user, _deadline, _v, _r, _s)
}

// BuySPOLWithDPOLPermit is a paid mutator transaction binding the contract method 0x44354ade.
//
// Solidity: function buySPOLWithDPOLPermit(uint256 _amount, uint16 _validatorOfDPOL, address _user, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256)
func (_SPOLController *SPOLControllerTransactorSession) BuySPOLWithDPOLPermit(_amount *big.Int, _validatorOfDPOL uint16, _user common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SPOLController.Contract.BuySPOLWithDPOLPermit(&_SPOLController.TransactOpts, _amount, _validatorOfDPOL, _user, _deadline, _v, _r, _s)
}

// ChangeFeeReceiver is a paid mutator transaction binding the contract method 0x7c08b964.
//
// Solidity: function changeFeeReceiver(address _newFeeReceiver) returns()
func (_SPOLController *SPOLControllerTransactor) ChangeFeeReceiver(opts *bind.TransactOpts, _newFeeReceiver common.Address) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "changeFeeReceiver", _newFeeReceiver)
}

// ChangeFeeReceiver is a paid mutator transaction binding the contract method 0x7c08b964.
//
// Solidity: function changeFeeReceiver(address _newFeeReceiver) returns()
func (_SPOLController *SPOLControllerSession) ChangeFeeReceiver(_newFeeReceiver common.Address) (*types.Transaction, error) {
	return _SPOLController.Contract.ChangeFeeReceiver(&_SPOLController.TransactOpts, _newFeeReceiver)
}

// ChangeFeeReceiver is a paid mutator transaction binding the contract method 0x7c08b964.
//
// Solidity: function changeFeeReceiver(address _newFeeReceiver) returns()
func (_SPOLController *SPOLControllerTransactorSession) ChangeFeeReceiver(_newFeeReceiver common.Address) (*types.Transaction, error) {
	return _SPOLController.Contract.ChangeFeeReceiver(&_SPOLController.TransactOpts, _newFeeReceiver)
}

// ChangeMaxDivergence is a paid mutator transaction binding the contract method 0xf1935b51.
//
// Solidity: function changeMaxDivergence(uint8 _newDivergence) returns()
func (_SPOLController *SPOLControllerTransactor) ChangeMaxDivergence(opts *bind.TransactOpts, _newDivergence uint8) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "changeMaxDivergence", _newDivergence)
}

// ChangeMaxDivergence is a paid mutator transaction binding the contract method 0xf1935b51.
//
// Solidity: function changeMaxDivergence(uint8 _newDivergence) returns()
func (_SPOLController *SPOLControllerSession) ChangeMaxDivergence(_newDivergence uint8) (*types.Transaction, error) {
	return _SPOLController.Contract.ChangeMaxDivergence(&_SPOLController.TransactOpts, _newDivergence)
}

// ChangeMaxDivergence is a paid mutator transaction binding the contract method 0xf1935b51.
//
// Solidity: function changeMaxDivergence(uint8 _newDivergence) returns()
func (_SPOLController *SPOLControllerTransactorSession) ChangeMaxDivergence(_newDivergence uint8) (*types.Transaction, error) {
	return _SPOLController.Contract.ChangeMaxDivergence(&_SPOLController.TransactOpts, _newDivergence)
}

// ChangeRewardFee is a paid mutator transaction binding the contract method 0xed69a126.
//
// Solidity: function changeRewardFee(uint16 _newFee) returns()
func (_SPOLController *SPOLControllerTransactor) ChangeRewardFee(opts *bind.TransactOpts, _newFee uint16) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "changeRewardFee", _newFee)
}

// ChangeRewardFee is a paid mutator transaction binding the contract method 0xed69a126.
//
// Solidity: function changeRewardFee(uint16 _newFee) returns()
func (_SPOLController *SPOLControllerSession) ChangeRewardFee(_newFee uint16) (*types.Transaction, error) {
	return _SPOLController.Contract.ChangeRewardFee(&_SPOLController.TransactOpts, _newFee)
}

// ChangeRewardFee is a paid mutator transaction binding the contract method 0xed69a126.
//
// Solidity: function changeRewardFee(uint16 _newFee) returns()
func (_SPOLController *SPOLControllerTransactorSession) ChangeRewardFee(_newFee uint16) (*types.Transaction, error) {
	return _SPOLController.Contract.ChangeRewardFee(&_SPOLController.TransactOpts, _newFee)
}

// CleanUpMaticPOL is a paid mutator transaction binding the contract method 0xd07bc162.
//
// Solidity: function cleanUpMaticPOL(uint16 _validator) returns()
func (_SPOLController *SPOLControllerTransactor) CleanUpMaticPOL(opts *bind.TransactOpts, _validator uint16) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "cleanUpMaticPOL", _validator)
}

// CleanUpMaticPOL is a paid mutator transaction binding the contract method 0xd07bc162.
//
// Solidity: function cleanUpMaticPOL(uint16 _validator) returns()
func (_SPOLController *SPOLControllerSession) CleanUpMaticPOL(_validator uint16) (*types.Transaction, error) {
	return _SPOLController.Contract.CleanUpMaticPOL(&_SPOLController.TransactOpts, _validator)
}

// CleanUpMaticPOL is a paid mutator transaction binding the contract method 0xd07bc162.
//
// Solidity: function cleanUpMaticPOL(uint16 _validator) returns()
func (_SPOLController *SPOLControllerTransactorSession) CleanUpMaticPOL(_validator uint16) (*types.Transaction, error) {
	return _SPOLController.Contract.CleanUpMaticPOL(&_SPOLController.TransactOpts, _validator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf4dc5606.
//
// Solidity: function initialize(uint16 _rewardFee, address _feeReceiver, uint8 _maxDivergence, address _authority) returns()
func (_SPOLController *SPOLControllerTransactor) Initialize(opts *bind.TransactOpts, _rewardFee uint16, _feeReceiver common.Address, _maxDivergence uint8, _authority common.Address) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "initialize", _rewardFee, _feeReceiver, _maxDivergence, _authority)
}

// Initialize is a paid mutator transaction binding the contract method 0xf4dc5606.
//
// Solidity: function initialize(uint16 _rewardFee, address _feeReceiver, uint8 _maxDivergence, address _authority) returns()
func (_SPOLController *SPOLControllerSession) Initialize(_rewardFee uint16, _feeReceiver common.Address, _maxDivergence uint8, _authority common.Address) (*types.Transaction, error) {
	return _SPOLController.Contract.Initialize(&_SPOLController.TransactOpts, _rewardFee, _feeReceiver, _maxDivergence, _authority)
}

// Initialize is a paid mutator transaction binding the contract method 0xf4dc5606.
//
// Solidity: function initialize(uint16 _rewardFee, address _feeReceiver, uint8 _maxDivergence, address _authority) returns()
func (_SPOLController *SPOLControllerTransactorSession) Initialize(_rewardFee uint16, _feeReceiver common.Address, _maxDivergence uint8, _authority common.Address) (*types.Transaction, error) {
	return _SPOLController.Contract.Initialize(&_SPOLController.TransactOpts, _rewardFee, _feeReceiver, _maxDivergence, _authority)
}

// MigrateValidator is a paid mutator transaction binding the contract method 0x9210db2c.
//
// Solidity: function migrateValidator(uint16 _oldValidator, uint16 _newValidator, uint256 _amount, bool _restake) returns()
func (_SPOLController *SPOLControllerTransactor) MigrateValidator(opts *bind.TransactOpts, _oldValidator uint16, _newValidator uint16, _amount *big.Int, _restake bool) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "migrateValidator", _oldValidator, _newValidator, _amount, _restake)
}

// MigrateValidator is a paid mutator transaction binding the contract method 0x9210db2c.
//
// Solidity: function migrateValidator(uint16 _oldValidator, uint16 _newValidator, uint256 _amount, bool _restake) returns()
func (_SPOLController *SPOLControllerSession) MigrateValidator(_oldValidator uint16, _newValidator uint16, _amount *big.Int, _restake bool) (*types.Transaction, error) {
	return _SPOLController.Contract.MigrateValidator(&_SPOLController.TransactOpts, _oldValidator, _newValidator, _amount, _restake)
}

// MigrateValidator is a paid mutator transaction binding the contract method 0x9210db2c.
//
// Solidity: function migrateValidator(uint16 _oldValidator, uint16 _newValidator, uint256 _amount, bool _restake) returns()
func (_SPOLController *SPOLControllerTransactorSession) MigrateValidator(_oldValidator uint16, _newValidator uint16, _amount *big.Int, _restake bool) (*types.Transaction, error) {
	return _SPOLController.Contract.MigrateValidator(&_SPOLController.TransactOpts, _oldValidator, _newValidator, _amount, _restake)
}

// PauseUserFunctions is a paid mutator transaction binding the contract method 0x57c380d9.
//
// Solidity: function pauseUserFunctions() returns()
func (_SPOLController *SPOLControllerTransactor) PauseUserFunctions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "pauseUserFunctions")
}

// PauseUserFunctions is a paid mutator transaction binding the contract method 0x57c380d9.
//
// Solidity: function pauseUserFunctions() returns()
func (_SPOLController *SPOLControllerSession) PauseUserFunctions() (*types.Transaction, error) {
	return _SPOLController.Contract.PauseUserFunctions(&_SPOLController.TransactOpts)
}

// PauseUserFunctions is a paid mutator transaction binding the contract method 0x57c380d9.
//
// Solidity: function pauseUserFunctions() returns()
func (_SPOLController *SPOLControllerTransactorSession) PauseUserFunctions() (*types.Transaction, error) {
	return _SPOLController.Contract.PauseUserFunctions(&_SPOLController.TransactOpts)
}

// ReloadAllActiveValidatorInfo is a paid mutator transaction binding the contract method 0x6f62b8bc.
//
// Solidity: function reloadAllActiveValidatorInfo() returns()
func (_SPOLController *SPOLControllerTransactor) ReloadAllActiveValidatorInfo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "reloadAllActiveValidatorInfo")
}

// ReloadAllActiveValidatorInfo is a paid mutator transaction binding the contract method 0x6f62b8bc.
//
// Solidity: function reloadAllActiveValidatorInfo() returns()
func (_SPOLController *SPOLControllerSession) ReloadAllActiveValidatorInfo() (*types.Transaction, error) {
	return _SPOLController.Contract.ReloadAllActiveValidatorInfo(&_SPOLController.TransactOpts)
}

// ReloadAllActiveValidatorInfo is a paid mutator transaction binding the contract method 0x6f62b8bc.
//
// Solidity: function reloadAllActiveValidatorInfo() returns()
func (_SPOLController *SPOLControllerTransactorSession) ReloadAllActiveValidatorInfo() (*types.Transaction, error) {
	return _SPOLController.Contract.ReloadAllActiveValidatorInfo(&_SPOLController.TransactOpts)
}

// ReloadValidatorInfo is a paid mutator transaction binding the contract method 0xb4761418.
//
// Solidity: function reloadValidatorInfo(uint16 _validator) returns()
func (_SPOLController *SPOLControllerTransactor) ReloadValidatorInfo(opts *bind.TransactOpts, _validator uint16) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "reloadValidatorInfo", _validator)
}

// ReloadValidatorInfo is a paid mutator transaction binding the contract method 0xb4761418.
//
// Solidity: function reloadValidatorInfo(uint16 _validator) returns()
func (_SPOLController *SPOLControllerSession) ReloadValidatorInfo(_validator uint16) (*types.Transaction, error) {
	return _SPOLController.Contract.ReloadValidatorInfo(&_SPOLController.TransactOpts, _validator)
}

// ReloadValidatorInfo is a paid mutator transaction binding the contract method 0xb4761418.
//
// Solidity: function reloadValidatorInfo(uint16 _validator) returns()
func (_SPOLController *SPOLControllerTransactorSession) ReloadValidatorInfo(_validator uint16) (*types.Transaction, error) {
	return _SPOLController.Contract.ReloadValidatorInfo(&_SPOLController.TransactOpts, _validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x280b0bae.
//
// Solidity: function removeValidator(uint16 _removedValidator) returns()
func (_SPOLController *SPOLControllerTransactor) RemoveValidator(opts *bind.TransactOpts, _removedValidator uint16) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "removeValidator", _removedValidator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x280b0bae.
//
// Solidity: function removeValidator(uint16 _removedValidator) returns()
func (_SPOLController *SPOLControllerSession) RemoveValidator(_removedValidator uint16) (*types.Transaction, error) {
	return _SPOLController.Contract.RemoveValidator(&_SPOLController.TransactOpts, _removedValidator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x280b0bae.
//
// Solidity: function removeValidator(uint16 _removedValidator) returns()
func (_SPOLController *SPOLControllerTransactorSession) RemoveValidator(_removedValidator uint16) (*types.Transaction, error) {
	return _SPOLController.Contract.RemoveValidator(&_SPOLController.TransactOpts, _removedValidator)
}

// RestakeAllActiveValidators is a paid mutator transaction binding the contract method 0x01cc750b.
//
// Solidity: function restakeAllActiveValidators() returns()
func (_SPOLController *SPOLControllerTransactor) RestakeAllActiveValidators(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "restakeAllActiveValidators")
}

// RestakeAllActiveValidators is a paid mutator transaction binding the contract method 0x01cc750b.
//
// Solidity: function restakeAllActiveValidators() returns()
func (_SPOLController *SPOLControllerSession) RestakeAllActiveValidators() (*types.Transaction, error) {
	return _SPOLController.Contract.RestakeAllActiveValidators(&_SPOLController.TransactOpts)
}

// RestakeAllActiveValidators is a paid mutator transaction binding the contract method 0x01cc750b.
//
// Solidity: function restakeAllActiveValidators() returns()
func (_SPOLController *SPOLControllerTransactorSession) RestakeAllActiveValidators() (*types.Transaction, error) {
	return _SPOLController.Contract.RestakeAllActiveValidators(&_SPOLController.TransactOpts)
}

// RestakeValidator is a paid mutator transaction binding the contract method 0x25f17fe0.
//
// Solidity: function restakeValidator(uint16 _validator) returns()
func (_SPOLController *SPOLControllerTransactor) RestakeValidator(opts *bind.TransactOpts, _validator uint16) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "restakeValidator", _validator)
}

// RestakeValidator is a paid mutator transaction binding the contract method 0x25f17fe0.
//
// Solidity: function restakeValidator(uint16 _validator) returns()
func (_SPOLController *SPOLControllerSession) RestakeValidator(_validator uint16) (*types.Transaction, error) {
	return _SPOLController.Contract.RestakeValidator(&_SPOLController.TransactOpts, _validator)
}

// RestakeValidator is a paid mutator transaction binding the contract method 0x25f17fe0.
//
// Solidity: function restakeValidator(uint16 _validator) returns()
func (_SPOLController *SPOLControllerTransactorSession) RestakeValidator(_validator uint16) (*types.Transaction, error) {
	return _SPOLController.Contract.RestakeValidator(&_SPOLController.TransactOpts, _validator)
}

// SellSPOL is a paid mutator transaction binding the contract method 0x32f42f13.
//
// Solidity: function sellSPOL(uint256 _amount, uint16 _validator) returns(uint256)
func (_SPOLController *SPOLControllerTransactor) SellSPOL(opts *bind.TransactOpts, _amount *big.Int, _validator uint16) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "sellSPOL", _amount, _validator)
}

// SellSPOL is a paid mutator transaction binding the contract method 0x32f42f13.
//
// Solidity: function sellSPOL(uint256 _amount, uint16 _validator) returns(uint256)
func (_SPOLController *SPOLControllerSession) SellSPOL(_amount *big.Int, _validator uint16) (*types.Transaction, error) {
	return _SPOLController.Contract.SellSPOL(&_SPOLController.TransactOpts, _amount, _validator)
}

// SellSPOL is a paid mutator transaction binding the contract method 0x32f42f13.
//
// Solidity: function sellSPOL(uint256 _amount, uint16 _validator) returns(uint256)
func (_SPOLController *SPOLControllerTransactorSession) SellSPOL(_amount *big.Int, _validator uint16) (*types.Transaction, error) {
	return _SPOLController.Contract.SellSPOL(&_SPOLController.TransactOpts, _amount, _validator)
}

// SellSPOL0 is a paid mutator transaction binding the contract method 0x5d43011f.
//
// Solidity: function sellSPOL(uint256 _amount) returns(uint256[])
func (_SPOLController *SPOLControllerTransactor) SellSPOL0(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "sellSPOL0", _amount)
}

// SellSPOL0 is a paid mutator transaction binding the contract method 0x5d43011f.
//
// Solidity: function sellSPOL(uint256 _amount) returns(uint256[])
func (_SPOLController *SPOLControllerSession) SellSPOL0(_amount *big.Int) (*types.Transaction, error) {
	return _SPOLController.Contract.SellSPOL0(&_SPOLController.TransactOpts, _amount)
}

// SellSPOL0 is a paid mutator transaction binding the contract method 0x5d43011f.
//
// Solidity: function sellSPOL(uint256 _amount) returns(uint256[])
func (_SPOLController *SPOLControllerTransactorSession) SellSPOL0(_amount *big.Int) (*types.Transaction, error) {
	return _SPOLController.Contract.SellSPOL0(&_SPOLController.TransactOpts, _amount)
}

// SellSPOLPermit is a paid mutator transaction binding the contract method 0x1341248d.
//
// Solidity: function sellSPOLPermit(uint256 _amount, uint16 _validator, address _user, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256)
func (_SPOLController *SPOLControllerTransactor) SellSPOLPermit(opts *bind.TransactOpts, _amount *big.Int, _validator uint16, _user common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "sellSPOLPermit", _amount, _validator, _user, _deadline, _v, _r, _s)
}

// SellSPOLPermit is a paid mutator transaction binding the contract method 0x1341248d.
//
// Solidity: function sellSPOLPermit(uint256 _amount, uint16 _validator, address _user, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256)
func (_SPOLController *SPOLControllerSession) SellSPOLPermit(_amount *big.Int, _validator uint16, _user common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SPOLController.Contract.SellSPOLPermit(&_SPOLController.TransactOpts, _amount, _validator, _user, _deadline, _v, _r, _s)
}

// SellSPOLPermit is a paid mutator transaction binding the contract method 0x1341248d.
//
// Solidity: function sellSPOLPermit(uint256 _amount, uint16 _validator, address _user, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256)
func (_SPOLController *SPOLControllerTransactorSession) SellSPOLPermit(_amount *big.Int, _validator uint16, _user common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SPOLController.Contract.SellSPOLPermit(&_SPOLController.TransactOpts, _amount, _validator, _user, _deadline, _v, _r, _s)
}

// SellSPOLPermit0 is a paid mutator transaction binding the contract method 0xacc150d0.
//
// Solidity: function sellSPOLPermit(uint256 _amount, address _user, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256[])
func (_SPOLController *SPOLControllerTransactor) SellSPOLPermit0(opts *bind.TransactOpts, _amount *big.Int, _user common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "sellSPOLPermit0", _amount, _user, _deadline, _v, _r, _s)
}

// SellSPOLPermit0 is a paid mutator transaction binding the contract method 0xacc150d0.
//
// Solidity: function sellSPOLPermit(uint256 _amount, address _user, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256[])
func (_SPOLController *SPOLControllerSession) SellSPOLPermit0(_amount *big.Int, _user common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SPOLController.Contract.SellSPOLPermit0(&_SPOLController.TransactOpts, _amount, _user, _deadline, _v, _r, _s)
}

// SellSPOLPermit0 is a paid mutator transaction binding the contract method 0xacc150d0.
//
// Solidity: function sellSPOLPermit(uint256 _amount, address _user, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256[])
func (_SPOLController *SPOLControllerTransactorSession) SellSPOLPermit0(_amount *big.Int, _user common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SPOLController.Contract.SellSPOLPermit0(&_SPOLController.TransactOpts, _amount, _user, _deadline, _v, _r, _s)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_SPOLController *SPOLControllerTransactor) SetAuthority(opts *bind.TransactOpts, newAuthority common.Address) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "setAuthority", newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_SPOLController *SPOLControllerSession) SetAuthority(newAuthority common.Address) (*types.Transaction, error) {
	return _SPOLController.Contract.SetAuthority(&_SPOLController.TransactOpts, newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_SPOLController *SPOLControllerTransactorSession) SetAuthority(newAuthority common.Address) (*types.Transaction, error) {
	return _SPOLController.Contract.SetAuthority(&_SPOLController.TransactOpts, newAuthority)
}

// TakeFee is a paid mutator transaction binding the contract method 0x181aa1fd.
//
// Solidity: function takeFee() returns()
func (_SPOLController *SPOLControllerTransactor) TakeFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "takeFee")
}

// TakeFee is a paid mutator transaction binding the contract method 0x181aa1fd.
//
// Solidity: function takeFee() returns()
func (_SPOLController *SPOLControllerSession) TakeFee() (*types.Transaction, error) {
	return _SPOLController.Contract.TakeFee(&_SPOLController.TransactOpts)
}

// TakeFee is a paid mutator transaction binding the contract method 0x181aa1fd.
//
// Solidity: function takeFee() returns()
func (_SPOLController *SPOLControllerTransactorSession) TakeFee() (*types.Transaction, error) {
	return _SPOLController.Contract.TakeFee(&_SPOLController.TransactOpts)
}

// UnpauseUserFunctions is a paid mutator transaction binding the contract method 0x7f06b47a.
//
// Solidity: function unpauseUserFunctions() returns()
func (_SPOLController *SPOLControllerTransactor) UnpauseUserFunctions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "unpauseUserFunctions")
}

// UnpauseUserFunctions is a paid mutator transaction binding the contract method 0x7f06b47a.
//
// Solidity: function unpauseUserFunctions() returns()
func (_SPOLController *SPOLControllerSession) UnpauseUserFunctions() (*types.Transaction, error) {
	return _SPOLController.Contract.UnpauseUserFunctions(&_SPOLController.TransactOpts)
}

// UnpauseUserFunctions is a paid mutator transaction binding the contract method 0x7f06b47a.
//
// Solidity: function unpauseUserFunctions() returns()
func (_SPOLController *SPOLControllerTransactorSession) UnpauseUserFunctions() (*types.Transaction, error) {
	return _SPOLController.Contract.UnpauseUserFunctions(&_SPOLController.TransactOpts)
}

// UpdateValidatorTargetShare is a paid mutator transaction binding the contract method 0x2dd11a0a.
//
// Solidity: function updateValidatorTargetShare(uint16[] _validatorID, uint8[] _newTargetShare) returns()
func (_SPOLController *SPOLControllerTransactor) UpdateValidatorTargetShare(opts *bind.TransactOpts, _validatorID []uint16, _newTargetShare []uint8) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "updateValidatorTargetShare", _validatorID, _newTargetShare)
}

// UpdateValidatorTargetShare is a paid mutator transaction binding the contract method 0x2dd11a0a.
//
// Solidity: function updateValidatorTargetShare(uint16[] _validatorID, uint8[] _newTargetShare) returns()
func (_SPOLController *SPOLControllerSession) UpdateValidatorTargetShare(_validatorID []uint16, _newTargetShare []uint8) (*types.Transaction, error) {
	return _SPOLController.Contract.UpdateValidatorTargetShare(&_SPOLController.TransactOpts, _validatorID, _newTargetShare)
}

// UpdateValidatorTargetShare is a paid mutator transaction binding the contract method 0x2dd11a0a.
//
// Solidity: function updateValidatorTargetShare(uint16[] _validatorID, uint8[] _newTargetShare) returns()
func (_SPOLController *SPOLControllerTransactorSession) UpdateValidatorTargetShare(_validatorID []uint16, _newTargetShare []uint8) (*types.Transaction, error) {
	return _SPOLController.Contract.UpdateValidatorTargetShare(&_SPOLController.TransactOpts, _validatorID, _newTargetShare)
}

// WithdrawPOL is a paid mutator transaction binding the contract method 0x61ad860b.
//
// Solidity: function withdrawPOL() returns()
func (_SPOLController *SPOLControllerTransactor) WithdrawPOL(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "withdrawPOL")
}

// WithdrawPOL is a paid mutator transaction binding the contract method 0x61ad860b.
//
// Solidity: function withdrawPOL() returns()
func (_SPOLController *SPOLControllerSession) WithdrawPOL() (*types.Transaction, error) {
	return _SPOLController.Contract.WithdrawPOL(&_SPOLController.TransactOpts)
}

// WithdrawPOL is a paid mutator transaction binding the contract method 0x61ad860b.
//
// Solidity: function withdrawPOL() returns()
func (_SPOLController *SPOLControllerTransactorSession) WithdrawPOL() (*types.Transaction, error) {
	return _SPOLController.Contract.WithdrawPOL(&_SPOLController.TransactOpts)
}

// WithdrawPOL0 is a paid mutator transaction binding the contract method 0x8ffcca07.
//
// Solidity: function withdrawPOL(address _user) returns()
func (_SPOLController *SPOLControllerTransactor) WithdrawPOL0(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _SPOLController.contract.Transact(opts, "withdrawPOL0", _user)
}

// WithdrawPOL0 is a paid mutator transaction binding the contract method 0x8ffcca07.
//
// Solidity: function withdrawPOL(address _user) returns()
func (_SPOLController *SPOLControllerSession) WithdrawPOL0(_user common.Address) (*types.Transaction, error) {
	return _SPOLController.Contract.WithdrawPOL0(&_SPOLController.TransactOpts, _user)
}

// WithdrawPOL0 is a paid mutator transaction binding the contract method 0x8ffcca07.
//
// Solidity: function withdrawPOL(address _user) returns()
func (_SPOLController *SPOLControllerTransactorSession) WithdrawPOL0(_user common.Address) (*types.Transaction, error) {
	return _SPOLController.Contract.WithdrawPOL0(&_SPOLController.TransactOpts, _user)
}

// SPOLControllerAuthorityUpdatedIterator is returned from FilterAuthorityUpdated and is used to iterate over the raw logs and unpacked data for AuthorityUpdated events raised by the SPOLController contract.
type SPOLControllerAuthorityUpdatedIterator struct {
	Event *SPOLControllerAuthorityUpdated // Event containing the contract specifics and raw log

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
func (it *SPOLControllerAuthorityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPOLControllerAuthorityUpdated)
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
		it.Event = new(SPOLControllerAuthorityUpdated)
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
func (it *SPOLControllerAuthorityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPOLControllerAuthorityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPOLControllerAuthorityUpdated represents a AuthorityUpdated event raised by the SPOLController contract.
type SPOLControllerAuthorityUpdated struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuthorityUpdated is a free log retrieval operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address authority)
func (_SPOLController *SPOLControllerFilterer) FilterAuthorityUpdated(opts *bind.FilterOpts) (*SPOLControllerAuthorityUpdatedIterator, error) {

	logs, sub, err := _SPOLController.contract.FilterLogs(opts, "AuthorityUpdated")
	if err != nil {
		return nil, err
	}
	return &SPOLControllerAuthorityUpdatedIterator{contract: _SPOLController.contract, event: "AuthorityUpdated", logs: logs, sub: sub}, nil
}

// WatchAuthorityUpdated is a free log subscription operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address authority)
func (_SPOLController *SPOLControllerFilterer) WatchAuthorityUpdated(opts *bind.WatchOpts, sink chan<- *SPOLControllerAuthorityUpdated) (event.Subscription, error) {

	logs, sub, err := _SPOLController.contract.WatchLogs(opts, "AuthorityUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPOLControllerAuthorityUpdated)
				if err := _SPOLController.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
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

// ParseAuthorityUpdated is a log parse operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address authority)
func (_SPOLController *SPOLControllerFilterer) ParseAuthorityUpdated(log types.Log) (*SPOLControllerAuthorityUpdated, error) {
	event := new(SPOLControllerAuthorityUpdated)
	if err := _SPOLController.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPOLControllerExchangeRateSnapshotIterator is returned from FilterExchangeRateSnapshot and is used to iterate over the raw logs and unpacked data for ExchangeRateSnapshot events raised by the SPOLController contract.
type SPOLControllerExchangeRateSnapshotIterator struct {
	Event *SPOLControllerExchangeRateSnapshot // Event containing the contract specifics and raw log

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
func (it *SPOLControllerExchangeRateSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPOLControllerExchangeRateSnapshot)
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
		it.Event = new(SPOLControllerExchangeRateSnapshot)
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
func (it *SPOLControllerExchangeRateSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPOLControllerExchangeRateSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPOLControllerExchangeRateSnapshot represents a ExchangeRateSnapshot event raised by the SPOLController contract.
type SPOLControllerExchangeRateSnapshot struct {
	TotalsPOLSupply  *big.Int
	TotalbPOLBalance *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterExchangeRateSnapshot is a free log retrieval operation binding the contract event 0x224ca655cc667083faf1f4801d73e7fcf269f4f638d87c62a3a400b5fb179b8a.
//
// Solidity: event ExchangeRateSnapshot(uint256 totalsPOLSupply, uint256 totalbPOLBalance)
func (_SPOLController *SPOLControllerFilterer) FilterExchangeRateSnapshot(opts *bind.FilterOpts) (*SPOLControllerExchangeRateSnapshotIterator, error) {

	logs, sub, err := _SPOLController.contract.FilterLogs(opts, "ExchangeRateSnapshot")
	if err != nil {
		return nil, err
	}
	return &SPOLControllerExchangeRateSnapshotIterator{contract: _SPOLController.contract, event: "ExchangeRateSnapshot", logs: logs, sub: sub}, nil
}

// WatchExchangeRateSnapshot is a free log subscription operation binding the contract event 0x224ca655cc667083faf1f4801d73e7fcf269f4f638d87c62a3a400b5fb179b8a.
//
// Solidity: event ExchangeRateSnapshot(uint256 totalsPOLSupply, uint256 totalbPOLBalance)
func (_SPOLController *SPOLControllerFilterer) WatchExchangeRateSnapshot(opts *bind.WatchOpts, sink chan<- *SPOLControllerExchangeRateSnapshot) (event.Subscription, error) {

	logs, sub, err := _SPOLController.contract.WatchLogs(opts, "ExchangeRateSnapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPOLControllerExchangeRateSnapshot)
				if err := _SPOLController.contract.UnpackLog(event, "ExchangeRateSnapshot", log); err != nil {
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

// ParseExchangeRateSnapshot is a log parse operation binding the contract event 0x224ca655cc667083faf1f4801d73e7fcf269f4f638d87c62a3a400b5fb179b8a.
//
// Solidity: event ExchangeRateSnapshot(uint256 totalsPOLSupply, uint256 totalbPOLBalance)
func (_SPOLController *SPOLControllerFilterer) ParseExchangeRateSnapshot(log types.Log) (*SPOLControllerExchangeRateSnapshot, error) {
	event := new(SPOLControllerExchangeRateSnapshot)
	if err := _SPOLController.contract.UnpackLog(event, "ExchangeRateSnapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPOLControllerFeeCollectedIterator is returned from FilterFeeCollected and is used to iterate over the raw logs and unpacked data for FeeCollected events raised by the SPOLController contract.
type SPOLControllerFeeCollectedIterator struct {
	Event *SPOLControllerFeeCollected // Event containing the contract specifics and raw log

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
func (it *SPOLControllerFeeCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPOLControllerFeeCollected)
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
		it.Event = new(SPOLControllerFeeCollected)
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
func (it *SPOLControllerFeeCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPOLControllerFeeCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPOLControllerFeeCollected represents a FeeCollected event raised by the SPOLController contract.
type SPOLControllerFeeCollected struct {
	FeeReceiver   common.Address
	FeePOLAmount  *big.Int
	FeesPOLAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeCollected is a free log retrieval operation binding the contract event 0x108516ddcf5ba43cea6bb2cd5ff6d59ac196c1c86ccb9178332b9dd72d1ca561.
//
// Solidity: event FeeCollected(address feeReceiver, uint256 feePOLAmount, uint256 feesPOLAmount)
func (_SPOLController *SPOLControllerFilterer) FilterFeeCollected(opts *bind.FilterOpts) (*SPOLControllerFeeCollectedIterator, error) {

	logs, sub, err := _SPOLController.contract.FilterLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return &SPOLControllerFeeCollectedIterator{contract: _SPOLController.contract, event: "FeeCollected", logs: logs, sub: sub}, nil
}

// WatchFeeCollected is a free log subscription operation binding the contract event 0x108516ddcf5ba43cea6bb2cd5ff6d59ac196c1c86ccb9178332b9dd72d1ca561.
//
// Solidity: event FeeCollected(address feeReceiver, uint256 feePOLAmount, uint256 feesPOLAmount)
func (_SPOLController *SPOLControllerFilterer) WatchFeeCollected(opts *bind.WatchOpts, sink chan<- *SPOLControllerFeeCollected) (event.Subscription, error) {

	logs, sub, err := _SPOLController.contract.WatchLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPOLControllerFeeCollected)
				if err := _SPOLController.contract.UnpackLog(event, "FeeCollected", log); err != nil {
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

// ParseFeeCollected is a log parse operation binding the contract event 0x108516ddcf5ba43cea6bb2cd5ff6d59ac196c1c86ccb9178332b9dd72d1ca561.
//
// Solidity: event FeeCollected(address feeReceiver, uint256 feePOLAmount, uint256 feesPOLAmount)
func (_SPOLController *SPOLControllerFilterer) ParseFeeCollected(log types.Log) (*SPOLControllerFeeCollected, error) {
	event := new(SPOLControllerFeeCollected)
	if err := _SPOLController.contract.UnpackLog(event, "FeeCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPOLControllerFeeReceiverChangedIterator is returned from FilterFeeReceiverChanged and is used to iterate over the raw logs and unpacked data for FeeReceiverChanged events raised by the SPOLController contract.
type SPOLControllerFeeReceiverChangedIterator struct {
	Event *SPOLControllerFeeReceiverChanged // Event containing the contract specifics and raw log

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
func (it *SPOLControllerFeeReceiverChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPOLControllerFeeReceiverChanged)
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
		it.Event = new(SPOLControllerFeeReceiverChanged)
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
func (it *SPOLControllerFeeReceiverChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPOLControllerFeeReceiverChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPOLControllerFeeReceiverChanged represents a FeeReceiverChanged event raised by the SPOLController contract.
type SPOLControllerFeeReceiverChanged struct {
	OldReceiver common.Address
	NewReceiver common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeReceiverChanged is a free log retrieval operation binding the contract event 0xa4b009cc442411b602eaf94bc0579b6abdb8fd90b4ef5b9426e270038906bd03.
//
// Solidity: event FeeReceiverChanged(address oldReceiver, address newReceiver)
func (_SPOLController *SPOLControllerFilterer) FilterFeeReceiverChanged(opts *bind.FilterOpts) (*SPOLControllerFeeReceiverChangedIterator, error) {

	logs, sub, err := _SPOLController.contract.FilterLogs(opts, "FeeReceiverChanged")
	if err != nil {
		return nil, err
	}
	return &SPOLControllerFeeReceiverChangedIterator{contract: _SPOLController.contract, event: "FeeReceiverChanged", logs: logs, sub: sub}, nil
}

// WatchFeeReceiverChanged is a free log subscription operation binding the contract event 0xa4b009cc442411b602eaf94bc0579b6abdb8fd90b4ef5b9426e270038906bd03.
//
// Solidity: event FeeReceiverChanged(address oldReceiver, address newReceiver)
func (_SPOLController *SPOLControllerFilterer) WatchFeeReceiverChanged(opts *bind.WatchOpts, sink chan<- *SPOLControllerFeeReceiverChanged) (event.Subscription, error) {

	logs, sub, err := _SPOLController.contract.WatchLogs(opts, "FeeReceiverChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPOLControllerFeeReceiverChanged)
				if err := _SPOLController.contract.UnpackLog(event, "FeeReceiverChanged", log); err != nil {
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

// ParseFeeReceiverChanged is a log parse operation binding the contract event 0xa4b009cc442411b602eaf94bc0579b6abdb8fd90b4ef5b9426e270038906bd03.
//
// Solidity: event FeeReceiverChanged(address oldReceiver, address newReceiver)
func (_SPOLController *SPOLControllerFilterer) ParseFeeReceiverChanged(log types.Log) (*SPOLControllerFeeReceiverChanged, error) {
	event := new(SPOLControllerFeeReceiverChanged)
	if err := _SPOLController.contract.UnpackLog(event, "FeeReceiverChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPOLControllerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SPOLController contract.
type SPOLControllerInitializedIterator struct {
	Event *SPOLControllerInitialized // Event containing the contract specifics and raw log

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
func (it *SPOLControllerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPOLControllerInitialized)
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
		it.Event = new(SPOLControllerInitialized)
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
func (it *SPOLControllerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPOLControllerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPOLControllerInitialized represents a Initialized event raised by the SPOLController contract.
type SPOLControllerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SPOLController *SPOLControllerFilterer) FilterInitialized(opts *bind.FilterOpts) (*SPOLControllerInitializedIterator, error) {

	logs, sub, err := _SPOLController.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SPOLControllerInitializedIterator{contract: _SPOLController.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SPOLController *SPOLControllerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SPOLControllerInitialized) (event.Subscription, error) {

	logs, sub, err := _SPOLController.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPOLControllerInitialized)
				if err := _SPOLController.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SPOLController *SPOLControllerFilterer) ParseInitialized(log types.Log) (*SPOLControllerInitialized, error) {
	event := new(SPOLControllerInitialized)
	if err := _SPOLController.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPOLControllerMaticTokensCleanedIterator is returned from FilterMaticTokensCleaned and is used to iterate over the raw logs and unpacked data for MaticTokensCleaned events raised by the SPOLController contract.
type SPOLControllerMaticTokensCleanedIterator struct {
	Event *SPOLControllerMaticTokensCleaned // Event containing the contract specifics and raw log

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
func (it *SPOLControllerMaticTokensCleanedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPOLControllerMaticTokensCleaned)
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
		it.Event = new(SPOLControllerMaticTokensCleaned)
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
func (it *SPOLControllerMaticTokensCleanedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPOLControllerMaticTokensCleanedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPOLControllerMaticTokensCleaned represents a MaticTokensCleaned event raised by the SPOLController contract.
type SPOLControllerMaticTokensCleaned struct {
	MaticAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMaticTokensCleaned is a free log retrieval operation binding the contract event 0xe903f99508af7d1bb92f7a6e13edc84c4b29e027758d6861d0bc81d37710cf43.
//
// Solidity: event MaticTokensCleaned(uint256 maticAmount)
func (_SPOLController *SPOLControllerFilterer) FilterMaticTokensCleaned(opts *bind.FilterOpts) (*SPOLControllerMaticTokensCleanedIterator, error) {

	logs, sub, err := _SPOLController.contract.FilterLogs(opts, "MaticTokensCleaned")
	if err != nil {
		return nil, err
	}
	return &SPOLControllerMaticTokensCleanedIterator{contract: _SPOLController.contract, event: "MaticTokensCleaned", logs: logs, sub: sub}, nil
}

// WatchMaticTokensCleaned is a free log subscription operation binding the contract event 0xe903f99508af7d1bb92f7a6e13edc84c4b29e027758d6861d0bc81d37710cf43.
//
// Solidity: event MaticTokensCleaned(uint256 maticAmount)
func (_SPOLController *SPOLControllerFilterer) WatchMaticTokensCleaned(opts *bind.WatchOpts, sink chan<- *SPOLControllerMaticTokensCleaned) (event.Subscription, error) {

	logs, sub, err := _SPOLController.contract.WatchLogs(opts, "MaticTokensCleaned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPOLControllerMaticTokensCleaned)
				if err := _SPOLController.contract.UnpackLog(event, "MaticTokensCleaned", log); err != nil {
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

// ParseMaticTokensCleaned is a log parse operation binding the contract event 0xe903f99508af7d1bb92f7a6e13edc84c4b29e027758d6861d0bc81d37710cf43.
//
// Solidity: event MaticTokensCleaned(uint256 maticAmount)
func (_SPOLController *SPOLControllerFilterer) ParseMaticTokensCleaned(log types.Log) (*SPOLControllerMaticTokensCleaned, error) {
	event := new(SPOLControllerMaticTokensCleaned)
	if err := _SPOLController.contract.UnpackLog(event, "MaticTokensCleaned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPOLControllerMaxDivergenceChangedIterator is returned from FilterMaxDivergenceChanged and is used to iterate over the raw logs and unpacked data for MaxDivergenceChanged events raised by the SPOLController contract.
type SPOLControllerMaxDivergenceChangedIterator struct {
	Event *SPOLControllerMaxDivergenceChanged // Event containing the contract specifics and raw log

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
func (it *SPOLControllerMaxDivergenceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPOLControllerMaxDivergenceChanged)
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
		it.Event = new(SPOLControllerMaxDivergenceChanged)
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
func (it *SPOLControllerMaxDivergenceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPOLControllerMaxDivergenceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPOLControllerMaxDivergenceChanged represents a MaxDivergenceChanged event raised by the SPOLController contract.
type SPOLControllerMaxDivergenceChanged struct {
	OldDivergence uint8
	NewDivergence uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxDivergenceChanged is a free log retrieval operation binding the contract event 0xd2760d9b555325e879e19c942bd4506f46f5a8f19f52eaa21f620102d335edd8.
//
// Solidity: event MaxDivergenceChanged(uint8 oldDivergence, uint8 newDivergence)
func (_SPOLController *SPOLControllerFilterer) FilterMaxDivergenceChanged(opts *bind.FilterOpts) (*SPOLControllerMaxDivergenceChangedIterator, error) {

	logs, sub, err := _SPOLController.contract.FilterLogs(opts, "MaxDivergenceChanged")
	if err != nil {
		return nil, err
	}
	return &SPOLControllerMaxDivergenceChangedIterator{contract: _SPOLController.contract, event: "MaxDivergenceChanged", logs: logs, sub: sub}, nil
}

// WatchMaxDivergenceChanged is a free log subscription operation binding the contract event 0xd2760d9b555325e879e19c942bd4506f46f5a8f19f52eaa21f620102d335edd8.
//
// Solidity: event MaxDivergenceChanged(uint8 oldDivergence, uint8 newDivergence)
func (_SPOLController *SPOLControllerFilterer) WatchMaxDivergenceChanged(opts *bind.WatchOpts, sink chan<- *SPOLControllerMaxDivergenceChanged) (event.Subscription, error) {

	logs, sub, err := _SPOLController.contract.WatchLogs(opts, "MaxDivergenceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPOLControllerMaxDivergenceChanged)
				if err := _SPOLController.contract.UnpackLog(event, "MaxDivergenceChanged", log); err != nil {
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

// ParseMaxDivergenceChanged is a log parse operation binding the contract event 0xd2760d9b555325e879e19c942bd4506f46f5a8f19f52eaa21f620102d335edd8.
//
// Solidity: event MaxDivergenceChanged(uint8 oldDivergence, uint8 newDivergence)
func (_SPOLController *SPOLControllerFilterer) ParseMaxDivergenceChanged(log types.Log) (*SPOLControllerMaxDivergenceChanged, error) {
	event := new(SPOLControllerMaxDivergenceChanged)
	if err := _SPOLController.contract.UnpackLog(event, "MaxDivergenceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPOLControllerPOLTokensCleanedIterator is returned from FilterPOLTokensCleaned and is used to iterate over the raw logs and unpacked data for POLTokensCleaned events raised by the SPOLController contract.
type SPOLControllerPOLTokensCleanedIterator struct {
	Event *SPOLControllerPOLTokensCleaned // Event containing the contract specifics and raw log

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
func (it *SPOLControllerPOLTokensCleanedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPOLControllerPOLTokensCleaned)
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
		it.Event = new(SPOLControllerPOLTokensCleaned)
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
func (it *SPOLControllerPOLTokensCleanedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPOLControllerPOLTokensCleanedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPOLControllerPOLTokensCleaned represents a POLTokensCleaned event raised by the SPOLController contract.
type SPOLControllerPOLTokensCleaned struct {
	ValidatorId uint16
	PolAmount   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPOLTokensCleaned is a free log retrieval operation binding the contract event 0x5e3bbed4d7d4bcd63a120d1d3e1d66e2ec9c608bcfe9633800cfd7929cb6ef1e.
//
// Solidity: event POLTokensCleaned(uint16 validatorId, uint256 polAmount)
func (_SPOLController *SPOLControllerFilterer) FilterPOLTokensCleaned(opts *bind.FilterOpts) (*SPOLControllerPOLTokensCleanedIterator, error) {

	logs, sub, err := _SPOLController.contract.FilterLogs(opts, "POLTokensCleaned")
	if err != nil {
		return nil, err
	}
	return &SPOLControllerPOLTokensCleanedIterator{contract: _SPOLController.contract, event: "POLTokensCleaned", logs: logs, sub: sub}, nil
}

// WatchPOLTokensCleaned is a free log subscription operation binding the contract event 0x5e3bbed4d7d4bcd63a120d1d3e1d66e2ec9c608bcfe9633800cfd7929cb6ef1e.
//
// Solidity: event POLTokensCleaned(uint16 validatorId, uint256 polAmount)
func (_SPOLController *SPOLControllerFilterer) WatchPOLTokensCleaned(opts *bind.WatchOpts, sink chan<- *SPOLControllerPOLTokensCleaned) (event.Subscription, error) {

	logs, sub, err := _SPOLController.contract.WatchLogs(opts, "POLTokensCleaned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPOLControllerPOLTokensCleaned)
				if err := _SPOLController.contract.UnpackLog(event, "POLTokensCleaned", log); err != nil {
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

// ParsePOLTokensCleaned is a log parse operation binding the contract event 0x5e3bbed4d7d4bcd63a120d1d3e1d66e2ec9c608bcfe9633800cfd7929cb6ef1e.
//
// Solidity: event POLTokensCleaned(uint16 validatorId, uint256 polAmount)
func (_SPOLController *SPOLControllerFilterer) ParsePOLTokensCleaned(log types.Log) (*SPOLControllerPOLTokensCleaned, error) {
	event := new(SPOLControllerPOLTokensCleaned)
	if err := _SPOLController.contract.UnpackLog(event, "POLTokensCleaned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPOLControllerPOLWithdrawnIterator is returned from FilterPOLWithdrawn and is used to iterate over the raw logs and unpacked data for POLWithdrawn events raised by the SPOLController contract.
type SPOLControllerPOLWithdrawnIterator struct {
	Event *SPOLControllerPOLWithdrawn // Event containing the contract specifics and raw log

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
func (it *SPOLControllerPOLWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPOLControllerPOLWithdrawn)
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
		it.Event = new(SPOLControllerPOLWithdrawn)
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
func (it *SPOLControllerPOLWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPOLControllerPOLWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPOLControllerPOLWithdrawn represents a POLWithdrawn event raised by the SPOLController contract.
type SPOLControllerPOLWithdrawn struct {
	User      common.Address
	AmountPOL *big.Int
	Nonce     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPOLWithdrawn is a free log retrieval operation binding the contract event 0x78f3c96c1dbda8d5d7386a5d5b6185c4c7d2537dc5d4c14976754cf0ef121990.
//
// Solidity: event POLWithdrawn(address indexed user, uint256 amountPOL, uint256 nonce)
func (_SPOLController *SPOLControllerFilterer) FilterPOLWithdrawn(opts *bind.FilterOpts, user []common.Address) (*SPOLControllerPOLWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SPOLController.contract.FilterLogs(opts, "POLWithdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &SPOLControllerPOLWithdrawnIterator{contract: _SPOLController.contract, event: "POLWithdrawn", logs: logs, sub: sub}, nil
}

// WatchPOLWithdrawn is a free log subscription operation binding the contract event 0x78f3c96c1dbda8d5d7386a5d5b6185c4c7d2537dc5d4c14976754cf0ef121990.
//
// Solidity: event POLWithdrawn(address indexed user, uint256 amountPOL, uint256 nonce)
func (_SPOLController *SPOLControllerFilterer) WatchPOLWithdrawn(opts *bind.WatchOpts, sink chan<- *SPOLControllerPOLWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SPOLController.contract.WatchLogs(opts, "POLWithdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPOLControllerPOLWithdrawn)
				if err := _SPOLController.contract.UnpackLog(event, "POLWithdrawn", log); err != nil {
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

// ParsePOLWithdrawn is a log parse operation binding the contract event 0x78f3c96c1dbda8d5d7386a5d5b6185c4c7d2537dc5d4c14976754cf0ef121990.
//
// Solidity: event POLWithdrawn(address indexed user, uint256 amountPOL, uint256 nonce)
func (_SPOLController *SPOLControllerFilterer) ParsePOLWithdrawn(log types.Log) (*SPOLControllerPOLWithdrawn, error) {
	event := new(SPOLControllerPOLWithdrawn)
	if err := _SPOLController.contract.UnpackLog(event, "POLWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPOLControllerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SPOLController contract.
type SPOLControllerPausedIterator struct {
	Event *SPOLControllerPaused // Event containing the contract specifics and raw log

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
func (it *SPOLControllerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPOLControllerPaused)
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
		it.Event = new(SPOLControllerPaused)
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
func (it *SPOLControllerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPOLControllerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPOLControllerPaused represents a Paused event raised by the SPOLController contract.
type SPOLControllerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SPOLController *SPOLControllerFilterer) FilterPaused(opts *bind.FilterOpts) (*SPOLControllerPausedIterator, error) {

	logs, sub, err := _SPOLController.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SPOLControllerPausedIterator{contract: _SPOLController.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SPOLController *SPOLControllerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SPOLControllerPaused) (event.Subscription, error) {

	logs, sub, err := _SPOLController.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPOLControllerPaused)
				if err := _SPOLController.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SPOLController *SPOLControllerFilterer) ParsePaused(log types.Log) (*SPOLControllerPaused, error) {
	event := new(SPOLControllerPaused)
	if err := _SPOLController.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPOLControllerRewardFeeChangedIterator is returned from FilterRewardFeeChanged and is used to iterate over the raw logs and unpacked data for RewardFeeChanged events raised by the SPOLController contract.
type SPOLControllerRewardFeeChangedIterator struct {
	Event *SPOLControllerRewardFeeChanged // Event containing the contract specifics and raw log

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
func (it *SPOLControllerRewardFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPOLControllerRewardFeeChanged)
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
		it.Event = new(SPOLControllerRewardFeeChanged)
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
func (it *SPOLControllerRewardFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPOLControllerRewardFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPOLControllerRewardFeeChanged represents a RewardFeeChanged event raised by the SPOLController contract.
type SPOLControllerRewardFeeChanged struct {
	OldFee uint16
	NewFee uint16
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardFeeChanged is a free log retrieval operation binding the contract event 0x93b9b28eb59d392e802bcd43f6f8b19a9fdc46bb46600d9927baa35e0f765acf.
//
// Solidity: event RewardFeeChanged(uint16 oldFee, uint16 newFee)
func (_SPOLController *SPOLControllerFilterer) FilterRewardFeeChanged(opts *bind.FilterOpts) (*SPOLControllerRewardFeeChangedIterator, error) {

	logs, sub, err := _SPOLController.contract.FilterLogs(opts, "RewardFeeChanged")
	if err != nil {
		return nil, err
	}
	return &SPOLControllerRewardFeeChangedIterator{contract: _SPOLController.contract, event: "RewardFeeChanged", logs: logs, sub: sub}, nil
}

// WatchRewardFeeChanged is a free log subscription operation binding the contract event 0x93b9b28eb59d392e802bcd43f6f8b19a9fdc46bb46600d9927baa35e0f765acf.
//
// Solidity: event RewardFeeChanged(uint16 oldFee, uint16 newFee)
func (_SPOLController *SPOLControllerFilterer) WatchRewardFeeChanged(opts *bind.WatchOpts, sink chan<- *SPOLControllerRewardFeeChanged) (event.Subscription, error) {

	logs, sub, err := _SPOLController.contract.WatchLogs(opts, "RewardFeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPOLControllerRewardFeeChanged)
				if err := _SPOLController.contract.UnpackLog(event, "RewardFeeChanged", log); err != nil {
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

// ParseRewardFeeChanged is a log parse operation binding the contract event 0x93b9b28eb59d392e802bcd43f6f8b19a9fdc46bb46600d9927baa35e0f765acf.
//
// Solidity: event RewardFeeChanged(uint16 oldFee, uint16 newFee)
func (_SPOLController *SPOLControllerFilterer) ParseRewardFeeChanged(log types.Log) (*SPOLControllerRewardFeeChanged, error) {
	event := new(SPOLControllerRewardFeeChanged)
	if err := _SPOLController.contract.UnpackLog(event, "RewardFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPOLControllerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SPOLController contract.
type SPOLControllerUnpausedIterator struct {
	Event *SPOLControllerUnpaused // Event containing the contract specifics and raw log

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
func (it *SPOLControllerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPOLControllerUnpaused)
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
		it.Event = new(SPOLControllerUnpaused)
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
func (it *SPOLControllerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPOLControllerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPOLControllerUnpaused represents a Unpaused event raised by the SPOLController contract.
type SPOLControllerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SPOLController *SPOLControllerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SPOLControllerUnpausedIterator, error) {

	logs, sub, err := _SPOLController.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SPOLControllerUnpausedIterator{contract: _SPOLController.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SPOLController *SPOLControllerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SPOLControllerUnpaused) (event.Subscription, error) {

	logs, sub, err := _SPOLController.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPOLControllerUnpaused)
				if err := _SPOLController.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SPOLController *SPOLControllerFilterer) ParseUnpaused(log types.Log) (*SPOLControllerUnpaused, error) {
	event := new(SPOLControllerUnpaused)
	if err := _SPOLController.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPOLControllerValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the SPOLController contract.
type SPOLControllerValidatorAddedIterator struct {
	Event *SPOLControllerValidatorAdded // Event containing the contract specifics and raw log

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
func (it *SPOLControllerValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPOLControllerValidatorAdded)
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
		it.Event = new(SPOLControllerValidatorAdded)
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
func (it *SPOLControllerValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPOLControllerValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPOLControllerValidatorAdded represents a ValidatorAdded event raised by the SPOLController contract.
type SPOLControllerValidatorAdded struct {
	ValidatorId uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0x9015c2eaa6de7af24c629710f26f78b217803b3ffb5f0fb05aca18ad1ea1e69a.
//
// Solidity: event ValidatorAdded(uint16 validatorId)
func (_SPOLController *SPOLControllerFilterer) FilterValidatorAdded(opts *bind.FilterOpts) (*SPOLControllerValidatorAddedIterator, error) {

	logs, sub, err := _SPOLController.contract.FilterLogs(opts, "ValidatorAdded")
	if err != nil {
		return nil, err
	}
	return &SPOLControllerValidatorAddedIterator{contract: _SPOLController.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0x9015c2eaa6de7af24c629710f26f78b217803b3ffb5f0fb05aca18ad1ea1e69a.
//
// Solidity: event ValidatorAdded(uint16 validatorId)
func (_SPOLController *SPOLControllerFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *SPOLControllerValidatorAdded) (event.Subscription, error) {

	logs, sub, err := _SPOLController.contract.WatchLogs(opts, "ValidatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPOLControllerValidatorAdded)
				if err := _SPOLController.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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

// ParseValidatorAdded is a log parse operation binding the contract event 0x9015c2eaa6de7af24c629710f26f78b217803b3ffb5f0fb05aca18ad1ea1e69a.
//
// Solidity: event ValidatorAdded(uint16 validatorId)
func (_SPOLController *SPOLControllerFilterer) ParseValidatorAdded(log types.Log) (*SPOLControllerValidatorAdded, error) {
	event := new(SPOLControllerValidatorAdded)
	if err := _SPOLController.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPOLControllerValidatorMigratedIterator is returned from FilterValidatorMigrated and is used to iterate over the raw logs and unpacked data for ValidatorMigrated events raised by the SPOLController contract.
type SPOLControllerValidatorMigratedIterator struct {
	Event *SPOLControllerValidatorMigrated // Event containing the contract specifics and raw log

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
func (it *SPOLControllerValidatorMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPOLControllerValidatorMigrated)
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
		it.Event = new(SPOLControllerValidatorMigrated)
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
func (it *SPOLControllerValidatorMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPOLControllerValidatorMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPOLControllerValidatorMigrated represents a ValidatorMigrated event raised by the SPOLController contract.
type SPOLControllerValidatorMigrated struct {
	OldValidator uint16
	NewValidator uint16
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorMigrated is a free log retrieval operation binding the contract event 0x35b06bcdf7673eeb5a63ec39e7f9057cc7790921340d2b053cee552e0f1a8760.
//
// Solidity: event ValidatorMigrated(uint16 oldValidator, uint16 newValidator, uint256 amount)
func (_SPOLController *SPOLControllerFilterer) FilterValidatorMigrated(opts *bind.FilterOpts) (*SPOLControllerValidatorMigratedIterator, error) {

	logs, sub, err := _SPOLController.contract.FilterLogs(opts, "ValidatorMigrated")
	if err != nil {
		return nil, err
	}
	return &SPOLControllerValidatorMigratedIterator{contract: _SPOLController.contract, event: "ValidatorMigrated", logs: logs, sub: sub}, nil
}

// WatchValidatorMigrated is a free log subscription operation binding the contract event 0x35b06bcdf7673eeb5a63ec39e7f9057cc7790921340d2b053cee552e0f1a8760.
//
// Solidity: event ValidatorMigrated(uint16 oldValidator, uint16 newValidator, uint256 amount)
func (_SPOLController *SPOLControllerFilterer) WatchValidatorMigrated(opts *bind.WatchOpts, sink chan<- *SPOLControllerValidatorMigrated) (event.Subscription, error) {

	logs, sub, err := _SPOLController.contract.WatchLogs(opts, "ValidatorMigrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPOLControllerValidatorMigrated)
				if err := _SPOLController.contract.UnpackLog(event, "ValidatorMigrated", log); err != nil {
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

// ParseValidatorMigrated is a log parse operation binding the contract event 0x35b06bcdf7673eeb5a63ec39e7f9057cc7790921340d2b053cee552e0f1a8760.
//
// Solidity: event ValidatorMigrated(uint16 oldValidator, uint16 newValidator, uint256 amount)
func (_SPOLController *SPOLControllerFilterer) ParseValidatorMigrated(log types.Log) (*SPOLControllerValidatorMigrated, error) {
	event := new(SPOLControllerValidatorMigrated)
	if err := _SPOLController.contract.UnpackLog(event, "ValidatorMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPOLControllerValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the SPOLController contract.
type SPOLControllerValidatorRemovedIterator struct {
	Event *SPOLControllerValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *SPOLControllerValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPOLControllerValidatorRemoved)
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
		it.Event = new(SPOLControllerValidatorRemoved)
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
func (it *SPOLControllerValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPOLControllerValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPOLControllerValidatorRemoved represents a ValidatorRemoved event raised by the SPOLController contract.
type SPOLControllerValidatorRemoved struct {
	ValidatorId uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0x2c35389480b5cba2a073de5b8ba9ab59ded63507bb6dfb2466da2aa6717954c0.
//
// Solidity: event ValidatorRemoved(uint16 validatorId)
func (_SPOLController *SPOLControllerFilterer) FilterValidatorRemoved(opts *bind.FilterOpts) (*SPOLControllerValidatorRemovedIterator, error) {

	logs, sub, err := _SPOLController.contract.FilterLogs(opts, "ValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return &SPOLControllerValidatorRemovedIterator{contract: _SPOLController.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0x2c35389480b5cba2a073de5b8ba9ab59ded63507bb6dfb2466da2aa6717954c0.
//
// Solidity: event ValidatorRemoved(uint16 validatorId)
func (_SPOLController *SPOLControllerFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *SPOLControllerValidatorRemoved) (event.Subscription, error) {

	logs, sub, err := _SPOLController.contract.WatchLogs(opts, "ValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPOLControllerValidatorRemoved)
				if err := _SPOLController.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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

// ParseValidatorRemoved is a log parse operation binding the contract event 0x2c35389480b5cba2a073de5b8ba9ab59ded63507bb6dfb2466da2aa6717954c0.
//
// Solidity: event ValidatorRemoved(uint16 validatorId)
func (_SPOLController *SPOLControllerFilterer) ParseValidatorRemoved(log types.Log) (*SPOLControllerValidatorRemoved, error) {
	event := new(SPOLControllerValidatorRemoved)
	if err := _SPOLController.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPOLControllerValidatorTargetShareChangedIterator is returned from FilterValidatorTargetShareChanged and is used to iterate over the raw logs and unpacked data for ValidatorTargetShareChanged events raised by the SPOLController contract.
type SPOLControllerValidatorTargetShareChangedIterator struct {
	Event *SPOLControllerValidatorTargetShareChanged // Event containing the contract specifics and raw log

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
func (it *SPOLControllerValidatorTargetShareChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPOLControllerValidatorTargetShareChanged)
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
		it.Event = new(SPOLControllerValidatorTargetShareChanged)
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
func (it *SPOLControllerValidatorTargetShareChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPOLControllerValidatorTargetShareChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPOLControllerValidatorTargetShareChanged represents a ValidatorTargetShareChanged event raised by the SPOLController contract.
type SPOLControllerValidatorTargetShareChanged struct {
	ValidatorId    uint16
	NewTargetShare uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorTargetShareChanged is a free log retrieval operation binding the contract event 0x278fc29ecf3ec7b2b3e9bbcd3122772d466ec0e657e0e42f8a421367e5c5691f.
//
// Solidity: event ValidatorTargetShareChanged(uint16 validatorId, uint8 newTargetShare)
func (_SPOLController *SPOLControllerFilterer) FilterValidatorTargetShareChanged(opts *bind.FilterOpts) (*SPOLControllerValidatorTargetShareChangedIterator, error) {

	logs, sub, err := _SPOLController.contract.FilterLogs(opts, "ValidatorTargetShareChanged")
	if err != nil {
		return nil, err
	}
	return &SPOLControllerValidatorTargetShareChangedIterator{contract: _SPOLController.contract, event: "ValidatorTargetShareChanged", logs: logs, sub: sub}, nil
}

// WatchValidatorTargetShareChanged is a free log subscription operation binding the contract event 0x278fc29ecf3ec7b2b3e9bbcd3122772d466ec0e657e0e42f8a421367e5c5691f.
//
// Solidity: event ValidatorTargetShareChanged(uint16 validatorId, uint8 newTargetShare)
func (_SPOLController *SPOLControllerFilterer) WatchValidatorTargetShareChanged(opts *bind.WatchOpts, sink chan<- *SPOLControllerValidatorTargetShareChanged) (event.Subscription, error) {

	logs, sub, err := _SPOLController.contract.WatchLogs(opts, "ValidatorTargetShareChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPOLControllerValidatorTargetShareChanged)
				if err := _SPOLController.contract.UnpackLog(event, "ValidatorTargetShareChanged", log); err != nil {
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

// ParseValidatorTargetShareChanged is a log parse operation binding the contract event 0x278fc29ecf3ec7b2b3e9bbcd3122772d466ec0e657e0e42f8a421367e5c5691f.
//
// Solidity: event ValidatorTargetShareChanged(uint16 validatorId, uint8 newTargetShare)
func (_SPOLController *SPOLControllerFilterer) ParseValidatorTargetShareChanged(log types.Log) (*SPOLControllerValidatorTargetShareChanged, error) {
	event := new(SPOLControllerValidatorTargetShareChanged)
	if err := _SPOLController.contract.UnpackLog(event, "ValidatorTargetShareChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPOLControllerSPOLBurnedIterator is returned from FilterSPOLBurned and is used to iterate over the raw logs and unpacked data for SPOLBurned events raised by the SPOLController contract.
type SPOLControllerSPOLBurnedIterator struct {
	Event *SPOLControllerSPOLBurned // Event containing the contract specifics and raw log

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
func (it *SPOLControllerSPOLBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPOLControllerSPOLBurned)
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
		it.Event = new(SPOLControllerSPOLBurned)
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
func (it *SPOLControllerSPOLBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPOLControllerSPOLBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPOLControllerSPOLBurned represents a SPOLBurned event raised by the SPOLController contract.
type SPOLControllerSPOLBurned struct {
	User       common.Address
	AmountSPOL *big.Int
	AmountPOL  *big.Int
	Nonce      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSPOLBurned is a free log retrieval operation binding the contract event 0x995a8622255b30b649ea1635297eb4636d016d7d9ee35bd846c2d6c646f63ea9.
//
// Solidity: event sPOLBurned(address indexed user, uint256 amountSPOL, uint256 amountPOL, uint256 nonce)
func (_SPOLController *SPOLControllerFilterer) FilterSPOLBurned(opts *bind.FilterOpts, user []common.Address) (*SPOLControllerSPOLBurnedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SPOLController.contract.FilterLogs(opts, "sPOLBurned", userRule)
	if err != nil {
		return nil, err
	}
	return &SPOLControllerSPOLBurnedIterator{contract: _SPOLController.contract, event: "sPOLBurned", logs: logs, sub: sub}, nil
}

// WatchSPOLBurned is a free log subscription operation binding the contract event 0x995a8622255b30b649ea1635297eb4636d016d7d9ee35bd846c2d6c646f63ea9.
//
// Solidity: event sPOLBurned(address indexed user, uint256 amountSPOL, uint256 amountPOL, uint256 nonce)
func (_SPOLController *SPOLControllerFilterer) WatchSPOLBurned(opts *bind.WatchOpts, sink chan<- *SPOLControllerSPOLBurned, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SPOLController.contract.WatchLogs(opts, "sPOLBurned", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPOLControllerSPOLBurned)
				if err := _SPOLController.contract.UnpackLog(event, "sPOLBurned", log); err != nil {
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

// ParseSPOLBurned is a log parse operation binding the contract event 0x995a8622255b30b649ea1635297eb4636d016d7d9ee35bd846c2d6c646f63ea9.
//
// Solidity: event sPOLBurned(address indexed user, uint256 amountSPOL, uint256 amountPOL, uint256 nonce)
func (_SPOLController *SPOLControllerFilterer) ParseSPOLBurned(log types.Log) (*SPOLControllerSPOLBurned, error) {
	event := new(SPOLControllerSPOLBurned)
	if err := _SPOLController.contract.UnpackLog(event, "sPOLBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPOLControllerSPOLMintedIterator is returned from FilterSPOLMinted and is used to iterate over the raw logs and unpacked data for SPOLMinted events raised by the SPOLController contract.
type SPOLControllerSPOLMintedIterator struct {
	Event *SPOLControllerSPOLMinted // Event containing the contract specifics and raw log

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
func (it *SPOLControllerSPOLMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPOLControllerSPOLMinted)
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
		it.Event = new(SPOLControllerSPOLMinted)
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
func (it *SPOLControllerSPOLMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPOLControllerSPOLMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPOLControllerSPOLMinted represents a SPOLMinted event raised by the SPOLController contract.
type SPOLControllerSPOLMinted struct {
	User       common.Address
	AmountPOL  *big.Int
	AmountSPOL *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSPOLMinted is a free log retrieval operation binding the contract event 0xbee1c00aa78821ec45fe3686f893ebd117bf6d97a7d5c926c4041c93334cf03d.
//
// Solidity: event sPOLMinted(address indexed user, uint256 amountPOL, uint256 amountSPOL)
func (_SPOLController *SPOLControllerFilterer) FilterSPOLMinted(opts *bind.FilterOpts, user []common.Address) (*SPOLControllerSPOLMintedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SPOLController.contract.FilterLogs(opts, "sPOLMinted", userRule)
	if err != nil {
		return nil, err
	}
	return &SPOLControllerSPOLMintedIterator{contract: _SPOLController.contract, event: "sPOLMinted", logs: logs, sub: sub}, nil
}

// WatchSPOLMinted is a free log subscription operation binding the contract event 0xbee1c00aa78821ec45fe3686f893ebd117bf6d97a7d5c926c4041c93334cf03d.
//
// Solidity: event sPOLMinted(address indexed user, uint256 amountPOL, uint256 amountSPOL)
func (_SPOLController *SPOLControllerFilterer) WatchSPOLMinted(opts *bind.WatchOpts, sink chan<- *SPOLControllerSPOLMinted, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SPOLController.contract.WatchLogs(opts, "sPOLMinted", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPOLControllerSPOLMinted)
				if err := _SPOLController.contract.UnpackLog(event, "sPOLMinted", log); err != nil {
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

// ParseSPOLMinted is a log parse operation binding the contract event 0xbee1c00aa78821ec45fe3686f893ebd117bf6d97a7d5c926c4041c93334cf03d.
//
// Solidity: event sPOLMinted(address indexed user, uint256 amountPOL, uint256 amountSPOL)
func (_SPOLController *SPOLControllerFilterer) ParseSPOLMinted(log types.Log) (*SPOLControllerSPOLMinted, error) {
	event := new(SPOLControllerSPOLMinted)
	if err := _SPOLController.contract.UnpackLog(event, "sPOLMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

