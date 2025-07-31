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

// AggchainFEPOutputProposal is an auto generated low-level Go binding around an user-defined struct.
type AggchainFEPOutputProposal struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2BlockNumber *big.Int
}

// AggchainFEPMetaData contains all meta data concerning the AggchainFEP contract.
var AggchainFEPMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"_aggLayerGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AggchainManagerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggregationVkeyMustBeDifferentThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotProposeFutureL2Output\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggLayerGatewayAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2BlockNumberLessThanNextBlockNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2BlockTimeMustBeGreaterThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2OutputRootCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOptimisticModeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingOptimisticModeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimisticModeEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimisticModeNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RangeVkeyCommitmentMustBeDifferentThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupConfigHashMustBeDifferentThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StartL2TimestampMustBeLessThanCurrentTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubmissionIntervalMustBeGreaterThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValueAggchainVKey\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"AcceptAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOptimisticModeManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOptimisticModeManager\",\"type\":\"address\"}],\"name\":\"AcceptOptimisticModeManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"AcceptVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"AddAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldAggregationVkey\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAggregationVkey\",\"type\":\"bytes32\"}],\"name\":\"AggregationVkeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableOptimisticMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableUseDefaultGatewayFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableOptimisticMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableUseDefaultGatewayFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l1Timestamp\",\"type\":\"uint256\"}],\"name\":\"OutputProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldRangeVkeyCommitment\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newRangeVkeyCommitment\",\"type\":\"bytes32\"}],\"name\":\"RangeVkeyCommitmentUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldRollupConfigHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newRollupConfigHash\",\"type\":\"bytes32\"}],\"name\":\"RollupConfigHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldSubmissionInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSubmissionInterval\",\"type\":\"uint256\"}],\"name\":\"SubmissionIntervalUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAggchainManager\",\"type\":\"address\"}],\"name\":\"TransferAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentOptimisticModeManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingOptimisticModeManager\",\"type\":\"address\"}],\"name\":\"TransferOptimisticModeManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingVKeyManager\",\"type\":\"address\"}],\"name\":\"TransferVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousAggchainVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateAggchainVKey\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGCHAIN_TYPE\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONSENSUS_TYPE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_BLOCK_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUBMISSION_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAggchainManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOptimisticModeManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptVKeyManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"addOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggLayerGateway\",\"outputs\":[{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggchainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregationVkey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"computeL2Timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableOptimisticMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableOptimisticMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getAggchainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainTypeFromSelector\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"aggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes2\",\"name\":\"aggchainVKeyVersion\",\"type\":\"bytes2\"},{\"internalType\":\"bytes2\",\"name\":\"aggchainType\",\"type\":\"bytes2\"}],\"name\":\"getAggchainVKeySelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKeyVersionFromSelector\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2OutputIndex\",\"type\":\"uint256\"}],\"name\":\"getL2Output\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structAggchainFEP.OutputProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"initAggchainManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializeBytesAggchain\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2BlockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestOutputIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOutputIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"onVerifyPessimistic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimisticMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimisticModeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"ownedAggchainVKeys\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"ownedAggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAggchainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOptimisticModeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingVKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rangeVkeyCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupConfigHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submissionInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"transferAggchainManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOptimisticModeManager\",\"type\":\"address\"}],\"name\":\"transferOptimisticModeManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"transferVKeyManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_aggregationVkey\",\"type\":\"bytes32\"}],\"name\":\"updateAggregationVkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"updatedAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"updateOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_rangeVkeyCommitment\",\"type\":\"bytes32\"}],\"name\":\"updateRangeVkeyCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_rollupConfigHash\",\"type\":\"bytes32\"}],\"name\":\"updateRollupConfigHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_submissionInterval\",\"type\":\"uint256\"}],\"name\":\"updateSubmissionInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useDefaultGateway\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AggchainFEPABI is the input ABI used to generate the binding from.
// Deprecated: Use AggchainFEPMetaData.ABI instead.
var AggchainFEPABI = AggchainFEPMetaData.ABI

// AggchainFEP is an auto generated Go binding around an Ethereum contract.
type AggchainFEP struct {
	AggchainFEPCaller     // Read-only binding to the contract
	AggchainFEPTransactor // Write-only binding to the contract
	AggchainFEPFilterer   // Log filterer for contract events
}

// AggchainFEPCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggchainFEPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggchainFEPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggchainFEPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggchainFEPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggchainFEPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggchainFEPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggchainFEPSession struct {
	Contract     *AggchainFEP      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AggchainFEPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggchainFEPCallerSession struct {
	Contract *AggchainFEPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AggchainFEPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggchainFEPTransactorSession struct {
	Contract     *AggchainFEPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AggchainFEPRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggchainFEPRaw struct {
	Contract *AggchainFEP // Generic contract binding to access the raw methods on
}

// AggchainFEPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggchainFEPCallerRaw struct {
	Contract *AggchainFEPCaller // Generic read-only contract binding to access the raw methods on
}

// AggchainFEPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggchainFEPTransactorRaw struct {
	Contract *AggchainFEPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggchainFEP creates a new instance of AggchainFEP, bound to a specific deployed contract.
func NewAggchainFEP(address common.Address, backend bind.ContractBackend) (*AggchainFEP, error) {
	contract, err := bindAggchainFEP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AggchainFEP{AggchainFEPCaller: AggchainFEPCaller{contract: contract}, AggchainFEPTransactor: AggchainFEPTransactor{contract: contract}, AggchainFEPFilterer: AggchainFEPFilterer{contract: contract}}, nil
}

// NewAggchainFEPCaller creates a new read-only instance of AggchainFEP, bound to a specific deployed contract.
func NewAggchainFEPCaller(address common.Address, caller bind.ContractCaller) (*AggchainFEPCaller, error) {
	contract, err := bindAggchainFEP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggchainFEPCaller{contract: contract}, nil
}

// NewAggchainFEPTransactor creates a new write-only instance of AggchainFEP, bound to a specific deployed contract.
func NewAggchainFEPTransactor(address common.Address, transactor bind.ContractTransactor) (*AggchainFEPTransactor, error) {
	contract, err := bindAggchainFEP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggchainFEPTransactor{contract: contract}, nil
}

// NewAggchainFEPFilterer creates a new log filterer instance of AggchainFEP, bound to a specific deployed contract.
func NewAggchainFEPFilterer(address common.Address, filterer bind.ContractFilterer) (*AggchainFEPFilterer, error) {
	contract, err := bindAggchainFEP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggchainFEPFilterer{contract: contract}, nil
}

// bindAggchainFEP binds a generic wrapper to an already deployed contract.
func bindAggchainFEP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AggchainFEPMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggchainFEP *AggchainFEPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggchainFEP.Contract.AggchainFEPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggchainFEP *AggchainFEPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggchainFEP.Contract.AggchainFEPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggchainFEP *AggchainFEPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggchainFEP.Contract.AggchainFEPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggchainFEP *AggchainFEPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggchainFEP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggchainFEP *AggchainFEPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggchainFEP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggchainFEP *AggchainFEPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggchainFEP.Contract.contract.Transact(opts, method, params...)
}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(bytes2)
func (_AggchainFEP *AggchainFEPCaller) AGGCHAINTYPE(opts *bind.CallOpts) ([2]byte, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "AGGCHAIN_TYPE")

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(bytes2)
func (_AggchainFEP *AggchainFEPSession) AGGCHAINTYPE() ([2]byte, error) {
	return _AggchainFEP.Contract.AGGCHAINTYPE(&_AggchainFEP.CallOpts)
}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(bytes2)
func (_AggchainFEP *AggchainFEPCallerSession) AGGCHAINTYPE() ([2]byte, error) {
	return _AggchainFEP.Contract.AGGCHAINTYPE(&_AggchainFEP.CallOpts)
}

// CONSENSUSTYPE is a free data retrieval call binding the contract method 0xcea5a4c0.
//
// Solidity: function CONSENSUS_TYPE() view returns(uint32)
func (_AggchainFEP *AggchainFEPCaller) CONSENSUSTYPE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "CONSENSUS_TYPE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CONSENSUSTYPE is a free data retrieval call binding the contract method 0xcea5a4c0.
//
// Solidity: function CONSENSUS_TYPE() view returns(uint32)
func (_AggchainFEP *AggchainFEPSession) CONSENSUSTYPE() (uint32, error) {
	return _AggchainFEP.Contract.CONSENSUSTYPE(&_AggchainFEP.CallOpts)
}

// CONSENSUSTYPE is a free data retrieval call binding the contract method 0xcea5a4c0.
//
// Solidity: function CONSENSUS_TYPE() view returns(uint32)
func (_AggchainFEP *AggchainFEPCallerSession) CONSENSUSTYPE() (uint32, error) {
	return _AggchainFEP.Contract.CONSENSUSTYPE(&_AggchainFEP.CallOpts)
}

// L2BLOCKTIME is a free data retrieval call binding the contract method 0x002134cc.
//
// Solidity: function L2_BLOCK_TIME() view returns(uint256)
func (_AggchainFEP *AggchainFEPCaller) L2BLOCKTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "L2_BLOCK_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BLOCKTIME is a free data retrieval call binding the contract method 0x002134cc.
//
// Solidity: function L2_BLOCK_TIME() view returns(uint256)
func (_AggchainFEP *AggchainFEPSession) L2BLOCKTIME() (*big.Int, error) {
	return _AggchainFEP.Contract.L2BLOCKTIME(&_AggchainFEP.CallOpts)
}

// L2BLOCKTIME is a free data retrieval call binding the contract method 0x002134cc.
//
// Solidity: function L2_BLOCK_TIME() view returns(uint256)
func (_AggchainFEP *AggchainFEPCallerSession) L2BLOCKTIME() (*big.Int, error) {
	return _AggchainFEP.Contract.L2BLOCKTIME(&_AggchainFEP.CallOpts)
}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_AggchainFEP *AggchainFEPCaller) SUBMISSIONINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "SUBMISSION_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_AggchainFEP *AggchainFEPSession) SUBMISSIONINTERVAL() (*big.Int, error) {
	return _AggchainFEP.Contract.SUBMISSIONINTERVAL(&_AggchainFEP.CallOpts)
}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_AggchainFEP *AggchainFEPCallerSession) SUBMISSIONINTERVAL() (*big.Int, error) {
	return _AggchainFEP.Contract.SUBMISSIONINTERVAL(&_AggchainFEP.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AggchainFEP *AggchainFEPCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AggchainFEP *AggchainFEPSession) Admin() (common.Address, error) {
	return _AggchainFEP.Contract.Admin(&_AggchainFEP.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AggchainFEP *AggchainFEPCallerSession) Admin() (common.Address, error) {
	return _AggchainFEP.Contract.Admin(&_AggchainFEP.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_AggchainFEP *AggchainFEPCaller) AggLayerGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "aggLayerGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_AggchainFEP *AggchainFEPSession) AggLayerGateway() (common.Address, error) {
	return _AggchainFEP.Contract.AggLayerGateway(&_AggchainFEP.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_AggchainFEP *AggchainFEPCallerSession) AggLayerGateway() (common.Address, error) {
	return _AggchainFEP.Contract.AggLayerGateway(&_AggchainFEP.CallOpts)
}

// AggchainManager is a free data retrieval call binding the contract method 0x7388c436.
//
// Solidity: function aggchainManager() view returns(address)
func (_AggchainFEP *AggchainFEPCaller) AggchainManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "aggchainManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggchainManager is a free data retrieval call binding the contract method 0x7388c436.
//
// Solidity: function aggchainManager() view returns(address)
func (_AggchainFEP *AggchainFEPSession) AggchainManager() (common.Address, error) {
	return _AggchainFEP.Contract.AggchainManager(&_AggchainFEP.CallOpts)
}

// AggchainManager is a free data retrieval call binding the contract method 0x7388c436.
//
// Solidity: function aggchainManager() view returns(address)
func (_AggchainFEP *AggchainFEPCallerSession) AggchainManager() (common.Address, error) {
	return _AggchainFEP.Contract.AggchainManager(&_AggchainFEP.CallOpts)
}

// AggregationVkey is a free data retrieval call binding the contract method 0xc32e4e3e.
//
// Solidity: function aggregationVkey() view returns(bytes32)
func (_AggchainFEP *AggchainFEPCaller) AggregationVkey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "aggregationVkey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AggregationVkey is a free data retrieval call binding the contract method 0xc32e4e3e.
//
// Solidity: function aggregationVkey() view returns(bytes32)
func (_AggchainFEP *AggchainFEPSession) AggregationVkey() ([32]byte, error) {
	return _AggchainFEP.Contract.AggregationVkey(&_AggchainFEP.CallOpts)
}

// AggregationVkey is a free data retrieval call binding the contract method 0xc32e4e3e.
//
// Solidity: function aggregationVkey() view returns(bytes32)
func (_AggchainFEP *AggchainFEPCallerSession) AggregationVkey() ([32]byte, error) {
	return _AggchainFEP.Contract.AggregationVkey(&_AggchainFEP.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_AggchainFEP *AggchainFEPCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_AggchainFEP *AggchainFEPSession) BridgeAddress() (common.Address, error) {
	return _AggchainFEP.Contract.BridgeAddress(&_AggchainFEP.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_AggchainFEP *AggchainFEPCallerSession) BridgeAddress() (common.Address, error) {
	return _AggchainFEP.Contract.BridgeAddress(&_AggchainFEP.CallOpts)
}

// ComputeL2Timestamp is a free data retrieval call binding the contract method 0xd1de856c.
//
// Solidity: function computeL2Timestamp(uint256 _l2BlockNumber) view returns(uint256)
func (_AggchainFEP *AggchainFEPCaller) ComputeL2Timestamp(opts *bind.CallOpts, _l2BlockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "computeL2Timestamp", _l2BlockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeL2Timestamp is a free data retrieval call binding the contract method 0xd1de856c.
//
// Solidity: function computeL2Timestamp(uint256 _l2BlockNumber) view returns(uint256)
func (_AggchainFEP *AggchainFEPSession) ComputeL2Timestamp(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _AggchainFEP.Contract.ComputeL2Timestamp(&_AggchainFEP.CallOpts, _l2BlockNumber)
}

// ComputeL2Timestamp is a free data retrieval call binding the contract method 0xd1de856c.
//
// Solidity: function computeL2Timestamp(uint256 _l2BlockNumber) view returns(uint256)
func (_AggchainFEP *AggchainFEPCallerSession) ComputeL2Timestamp(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _AggchainFEP.Contract.ComputeL2Timestamp(&_AggchainFEP.CallOpts, _l2BlockNumber)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_AggchainFEP *AggchainFEPCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_AggchainFEP *AggchainFEPSession) ForceBatchAddress() (common.Address, error) {
	return _AggchainFEP.Contract.ForceBatchAddress(&_AggchainFEP.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_AggchainFEP *AggchainFEPCallerSession) ForceBatchAddress() (common.Address, error) {
	return _AggchainFEP.Contract.ForceBatchAddress(&_AggchainFEP.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_AggchainFEP *AggchainFEPCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_AggchainFEP *AggchainFEPSession) ForceBatchTimeout() (uint64, error) {
	return _AggchainFEP.Contract.ForceBatchTimeout(&_AggchainFEP.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_AggchainFEP *AggchainFEPCallerSession) ForceBatchTimeout() (uint64, error) {
	return _AggchainFEP.Contract.ForceBatchTimeout(&_AggchainFEP.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_AggchainFEP *AggchainFEPCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_AggchainFEP *AggchainFEPSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _AggchainFEP.Contract.ForcedBatches(&_AggchainFEP.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_AggchainFEP *AggchainFEPCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _AggchainFEP.Contract.ForcedBatches(&_AggchainFEP.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_AggchainFEP *AggchainFEPCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_AggchainFEP *AggchainFEPSession) GasTokenAddress() (common.Address, error) {
	return _AggchainFEP.Contract.GasTokenAddress(&_AggchainFEP.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_AggchainFEP *AggchainFEPCallerSession) GasTokenAddress() (common.Address, error) {
	return _AggchainFEP.Contract.GasTokenAddress(&_AggchainFEP.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_AggchainFEP *AggchainFEPCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_AggchainFEP *AggchainFEPSession) GasTokenNetwork() (uint32, error) {
	return _AggchainFEP.Contract.GasTokenNetwork(&_AggchainFEP.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_AggchainFEP *AggchainFEPCallerSession) GasTokenNetwork() (uint32, error) {
	return _AggchainFEP.Contract.GasTokenNetwork(&_AggchainFEP.CallOpts)
}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes aggchainData) view returns(bytes32)
func (_AggchainFEP *AggchainFEPCaller) GetAggchainHash(opts *bind.CallOpts, aggchainData []byte) ([32]byte, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "getAggchainHash", aggchainData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes aggchainData) view returns(bytes32)
func (_AggchainFEP *AggchainFEPSession) GetAggchainHash(aggchainData []byte) ([32]byte, error) {
	return _AggchainFEP.Contract.GetAggchainHash(&_AggchainFEP.CallOpts, aggchainData)
}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes aggchainData) view returns(bytes32)
func (_AggchainFEP *AggchainFEPCallerSession) GetAggchainHash(aggchainData []byte) ([32]byte, error) {
	return _AggchainFEP.Contract.GetAggchainHash(&_AggchainFEP.CallOpts, aggchainData)
}

// GetAggchainTypeFromSelector is a free data retrieval call binding the contract method 0x26f9b76d.
//
// Solidity: function getAggchainTypeFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_AggchainFEP *AggchainFEPCaller) GetAggchainTypeFromSelector(opts *bind.CallOpts, aggchainVKeySelector [4]byte) ([2]byte, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "getAggchainTypeFromSelector", aggchainVKeySelector)

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// GetAggchainTypeFromSelector is a free data retrieval call binding the contract method 0x26f9b76d.
//
// Solidity: function getAggchainTypeFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_AggchainFEP *AggchainFEPSession) GetAggchainTypeFromSelector(aggchainVKeySelector [4]byte) ([2]byte, error) {
	return _AggchainFEP.Contract.GetAggchainTypeFromSelector(&_AggchainFEP.CallOpts, aggchainVKeySelector)
}

// GetAggchainTypeFromSelector is a free data retrieval call binding the contract method 0x26f9b76d.
//
// Solidity: function getAggchainTypeFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_AggchainFEP *AggchainFEPCallerSession) GetAggchainTypeFromSelector(aggchainVKeySelector [4]byte) ([2]byte, error) {
	return _AggchainFEP.Contract.GetAggchainTypeFromSelector(&_AggchainFEP.CallOpts, aggchainVKeySelector)
}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey)
func (_AggchainFEP *AggchainFEPCaller) GetAggchainVKey(opts *bind.CallOpts, aggchainVKeySelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "getAggchainVKey", aggchainVKeySelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey)
func (_AggchainFEP *AggchainFEPSession) GetAggchainVKey(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _AggchainFEP.Contract.GetAggchainVKey(&_AggchainFEP.CallOpts, aggchainVKeySelector)
}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey)
func (_AggchainFEP *AggchainFEPCallerSession) GetAggchainVKey(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _AggchainFEP.Contract.GetAggchainVKey(&_AggchainFEP.CallOpts, aggchainVKeySelector)
}

// GetAggchainVKeySelector is a free data retrieval call binding the contract method 0x1d0b435e.
//
// Solidity: function getAggchainVKeySelector(bytes2 aggchainVKeyVersion, bytes2 aggchainType) pure returns(bytes4)
func (_AggchainFEP *AggchainFEPCaller) GetAggchainVKeySelector(opts *bind.CallOpts, aggchainVKeyVersion [2]byte, aggchainType [2]byte) ([4]byte, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "getAggchainVKeySelector", aggchainVKeyVersion, aggchainType)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// GetAggchainVKeySelector is a free data retrieval call binding the contract method 0x1d0b435e.
//
// Solidity: function getAggchainVKeySelector(bytes2 aggchainVKeyVersion, bytes2 aggchainType) pure returns(bytes4)
func (_AggchainFEP *AggchainFEPSession) GetAggchainVKeySelector(aggchainVKeyVersion [2]byte, aggchainType [2]byte) ([4]byte, error) {
	return _AggchainFEP.Contract.GetAggchainVKeySelector(&_AggchainFEP.CallOpts, aggchainVKeyVersion, aggchainType)
}

// GetAggchainVKeySelector is a free data retrieval call binding the contract method 0x1d0b435e.
//
// Solidity: function getAggchainVKeySelector(bytes2 aggchainVKeyVersion, bytes2 aggchainType) pure returns(bytes4)
func (_AggchainFEP *AggchainFEPCallerSession) GetAggchainVKeySelector(aggchainVKeyVersion [2]byte, aggchainType [2]byte) ([4]byte, error) {
	return _AggchainFEP.Contract.GetAggchainVKeySelector(&_AggchainFEP.CallOpts, aggchainVKeyVersion, aggchainType)
}

// GetAggchainVKeyVersionFromSelector is a free data retrieval call binding the contract method 0xe90a3409.
//
// Solidity: function getAggchainVKeyVersionFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_AggchainFEP *AggchainFEPCaller) GetAggchainVKeyVersionFromSelector(opts *bind.CallOpts, aggchainVKeySelector [4]byte) ([2]byte, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "getAggchainVKeyVersionFromSelector", aggchainVKeySelector)

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// GetAggchainVKeyVersionFromSelector is a free data retrieval call binding the contract method 0xe90a3409.
//
// Solidity: function getAggchainVKeyVersionFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_AggchainFEP *AggchainFEPSession) GetAggchainVKeyVersionFromSelector(aggchainVKeySelector [4]byte) ([2]byte, error) {
	return _AggchainFEP.Contract.GetAggchainVKeyVersionFromSelector(&_AggchainFEP.CallOpts, aggchainVKeySelector)
}

// GetAggchainVKeyVersionFromSelector is a free data retrieval call binding the contract method 0xe90a3409.
//
// Solidity: function getAggchainVKeyVersionFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_AggchainFEP *AggchainFEPCallerSession) GetAggchainVKeyVersionFromSelector(aggchainVKeySelector [4]byte) ([2]byte, error) {
	return _AggchainFEP.Contract.GetAggchainVKeyVersionFromSelector(&_AggchainFEP.CallOpts, aggchainVKeySelector)
}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((bytes32,uint128,uint128))
func (_AggchainFEP *AggchainFEPCaller) GetL2Output(opts *bind.CallOpts, _l2OutputIndex *big.Int) (AggchainFEPOutputProposal, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "getL2Output", _l2OutputIndex)

	if err != nil {
		return *new(AggchainFEPOutputProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(AggchainFEPOutputProposal)).(*AggchainFEPOutputProposal)

	return out0, err

}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((bytes32,uint128,uint128))
func (_AggchainFEP *AggchainFEPSession) GetL2Output(_l2OutputIndex *big.Int) (AggchainFEPOutputProposal, error) {
	return _AggchainFEP.Contract.GetL2Output(&_AggchainFEP.CallOpts, _l2OutputIndex)
}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((bytes32,uint128,uint128))
func (_AggchainFEP *AggchainFEPCallerSession) GetL2Output(_l2OutputIndex *big.Int) (AggchainFEPOutputProposal, error) {
	return _AggchainFEP.Contract.GetL2Output(&_AggchainFEP.CallOpts, _l2OutputIndex)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_AggchainFEP *AggchainFEPCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_AggchainFEP *AggchainFEPSession) GlobalExitRootManager() (common.Address, error) {
	return _AggchainFEP.Contract.GlobalExitRootManager(&_AggchainFEP.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_AggchainFEP *AggchainFEPCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _AggchainFEP.Contract.GlobalExitRootManager(&_AggchainFEP.CallOpts)
}

// Initialize0 is a free data retrieval call binding the contract method 0x71257022.
//
// Solidity: function initialize(address , address , uint32 , address , string , string ) pure returns()
func (_AggchainFEP *AggchainFEPCaller) Initialize0(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 string, arg5 string) error {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "initialize0", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return err
	}

	return err

}

// Initialize0 is a free data retrieval call binding the contract method 0x71257022.
//
// Solidity: function initialize(address , address , uint32 , address , string , string ) pure returns()
func (_AggchainFEP *AggchainFEPSession) Initialize0(arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 string, arg5 string) error {
	return _AggchainFEP.Contract.Initialize0(&_AggchainFEP.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize0 is a free data retrieval call binding the contract method 0x71257022.
//
// Solidity: function initialize(address , address , uint32 , address , string , string ) pure returns()
func (_AggchainFEP *AggchainFEPCallerSession) Initialize0(arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 string, arg5 string) error {
	return _AggchainFEP.Contract.Initialize0(&_AggchainFEP.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// L2BlockTime is a free data retrieval call binding the contract method 0x93991af3.
//
// Solidity: function l2BlockTime() view returns(uint256)
func (_AggchainFEP *AggchainFEPCaller) L2BlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "l2BlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BlockTime is a free data retrieval call binding the contract method 0x93991af3.
//
// Solidity: function l2BlockTime() view returns(uint256)
func (_AggchainFEP *AggchainFEPSession) L2BlockTime() (*big.Int, error) {
	return _AggchainFEP.Contract.L2BlockTime(&_AggchainFEP.CallOpts)
}

// L2BlockTime is a free data retrieval call binding the contract method 0x93991af3.
//
// Solidity: function l2BlockTime() view returns(uint256)
func (_AggchainFEP *AggchainFEPCallerSession) L2BlockTime() (*big.Int, error) {
	return _AggchainFEP.Contract.L2BlockTime(&_AggchainFEP.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_AggchainFEP *AggchainFEPCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_AggchainFEP *AggchainFEPSession) LastAccInputHash() ([32]byte, error) {
	return _AggchainFEP.Contract.LastAccInputHash(&_AggchainFEP.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_AggchainFEP *AggchainFEPCallerSession) LastAccInputHash() ([32]byte, error) {
	return _AggchainFEP.Contract.LastAccInputHash(&_AggchainFEP.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_AggchainFEP *AggchainFEPCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_AggchainFEP *AggchainFEPSession) LastForceBatch() (uint64, error) {
	return _AggchainFEP.Contract.LastForceBatch(&_AggchainFEP.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_AggchainFEP *AggchainFEPCallerSession) LastForceBatch() (uint64, error) {
	return _AggchainFEP.Contract.LastForceBatch(&_AggchainFEP.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_AggchainFEP *AggchainFEPCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_AggchainFEP *AggchainFEPSession) LastForceBatchSequenced() (uint64, error) {
	return _AggchainFEP.Contract.LastForceBatchSequenced(&_AggchainFEP.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_AggchainFEP *AggchainFEPCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _AggchainFEP.Contract.LastForceBatchSequenced(&_AggchainFEP.CallOpts)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_AggchainFEP *AggchainFEPCaller) LatestBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "latestBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_AggchainFEP *AggchainFEPSession) LatestBlockNumber() (*big.Int, error) {
	return _AggchainFEP.Contract.LatestBlockNumber(&_AggchainFEP.CallOpts)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_AggchainFEP *AggchainFEPCallerSession) LatestBlockNumber() (*big.Int, error) {
	return _AggchainFEP.Contract.LatestBlockNumber(&_AggchainFEP.CallOpts)
}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_AggchainFEP *AggchainFEPCaller) LatestOutputIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "latestOutputIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_AggchainFEP *AggchainFEPSession) LatestOutputIndex() (*big.Int, error) {
	return _AggchainFEP.Contract.LatestOutputIndex(&_AggchainFEP.CallOpts)
}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_AggchainFEP *AggchainFEPCallerSession) LatestOutputIndex() (*big.Int, error) {
	return _AggchainFEP.Contract.LatestOutputIndex(&_AggchainFEP.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_AggchainFEP *AggchainFEPCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_AggchainFEP *AggchainFEPSession) NetworkName() (string, error) {
	return _AggchainFEP.Contract.NetworkName(&_AggchainFEP.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_AggchainFEP *AggchainFEPCallerSession) NetworkName() (string, error) {
	return _AggchainFEP.Contract.NetworkName(&_AggchainFEP.CallOpts)
}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_AggchainFEP *AggchainFEPCaller) NextBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "nextBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_AggchainFEP *AggchainFEPSession) NextBlockNumber() (*big.Int, error) {
	return _AggchainFEP.Contract.NextBlockNumber(&_AggchainFEP.CallOpts)
}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_AggchainFEP *AggchainFEPCallerSession) NextBlockNumber() (*big.Int, error) {
	return _AggchainFEP.Contract.NextBlockNumber(&_AggchainFEP.CallOpts)
}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_AggchainFEP *AggchainFEPCaller) NextOutputIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "nextOutputIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_AggchainFEP *AggchainFEPSession) NextOutputIndex() (*big.Int, error) {
	return _AggchainFEP.Contract.NextOutputIndex(&_AggchainFEP.CallOpts)
}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_AggchainFEP *AggchainFEPCallerSession) NextOutputIndex() (*big.Int, error) {
	return _AggchainFEP.Contract.NextOutputIndex(&_AggchainFEP.CallOpts)
}

// OptimisticMode is a free data retrieval call binding the contract method 0x60caf7a0.
//
// Solidity: function optimisticMode() view returns(bool)
func (_AggchainFEP *AggchainFEPCaller) OptimisticMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "optimisticMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OptimisticMode is a free data retrieval call binding the contract method 0x60caf7a0.
//
// Solidity: function optimisticMode() view returns(bool)
func (_AggchainFEP *AggchainFEPSession) OptimisticMode() (bool, error) {
	return _AggchainFEP.Contract.OptimisticMode(&_AggchainFEP.CallOpts)
}

// OptimisticMode is a free data retrieval call binding the contract method 0x60caf7a0.
//
// Solidity: function optimisticMode() view returns(bool)
func (_AggchainFEP *AggchainFEPCallerSession) OptimisticMode() (bool, error) {
	return _AggchainFEP.Contract.OptimisticMode(&_AggchainFEP.CallOpts)
}

// OptimisticModeManager is a free data retrieval call binding the contract method 0x1cf810ee.
//
// Solidity: function optimisticModeManager() view returns(address)
func (_AggchainFEP *AggchainFEPCaller) OptimisticModeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "optimisticModeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptimisticModeManager is a free data retrieval call binding the contract method 0x1cf810ee.
//
// Solidity: function optimisticModeManager() view returns(address)
func (_AggchainFEP *AggchainFEPSession) OptimisticModeManager() (common.Address, error) {
	return _AggchainFEP.Contract.OptimisticModeManager(&_AggchainFEP.CallOpts)
}

// OptimisticModeManager is a free data retrieval call binding the contract method 0x1cf810ee.
//
// Solidity: function optimisticModeManager() view returns(address)
func (_AggchainFEP *AggchainFEPCallerSession) OptimisticModeManager() (common.Address, error) {
	return _AggchainFEP.Contract.OptimisticModeManager(&_AggchainFEP.CallOpts)
}

// OwnedAggchainVKeys is a free data retrieval call binding the contract method 0xeffb8479.
//
// Solidity: function ownedAggchainVKeys(bytes4 aggchainVKeySelector) view returns(bytes32 ownedAggchainVKey)
func (_AggchainFEP *AggchainFEPCaller) OwnedAggchainVKeys(opts *bind.CallOpts, aggchainVKeySelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "ownedAggchainVKeys", aggchainVKeySelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OwnedAggchainVKeys is a free data retrieval call binding the contract method 0xeffb8479.
//
// Solidity: function ownedAggchainVKeys(bytes4 aggchainVKeySelector) view returns(bytes32 ownedAggchainVKey)
func (_AggchainFEP *AggchainFEPSession) OwnedAggchainVKeys(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _AggchainFEP.Contract.OwnedAggchainVKeys(&_AggchainFEP.CallOpts, aggchainVKeySelector)
}

// OwnedAggchainVKeys is a free data retrieval call binding the contract method 0xeffb8479.
//
// Solidity: function ownedAggchainVKeys(bytes4 aggchainVKeySelector) view returns(bytes32 ownedAggchainVKey)
func (_AggchainFEP *AggchainFEPCallerSession) OwnedAggchainVKeys(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _AggchainFEP.Contract.OwnedAggchainVKeys(&_AggchainFEP.CallOpts, aggchainVKeySelector)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_AggchainFEP *AggchainFEPCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_AggchainFEP *AggchainFEPSession) PendingAdmin() (common.Address, error) {
	return _AggchainFEP.Contract.PendingAdmin(&_AggchainFEP.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_AggchainFEP *AggchainFEPCallerSession) PendingAdmin() (common.Address, error) {
	return _AggchainFEP.Contract.PendingAdmin(&_AggchainFEP.CallOpts)
}

// PendingAggchainManager is a free data retrieval call binding the contract method 0x527570f1.
//
// Solidity: function pendingAggchainManager() view returns(address)
func (_AggchainFEP *AggchainFEPCaller) PendingAggchainManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "pendingAggchainManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAggchainManager is a free data retrieval call binding the contract method 0x527570f1.
//
// Solidity: function pendingAggchainManager() view returns(address)
func (_AggchainFEP *AggchainFEPSession) PendingAggchainManager() (common.Address, error) {
	return _AggchainFEP.Contract.PendingAggchainManager(&_AggchainFEP.CallOpts)
}

// PendingAggchainManager is a free data retrieval call binding the contract method 0x527570f1.
//
// Solidity: function pendingAggchainManager() view returns(address)
func (_AggchainFEP *AggchainFEPCallerSession) PendingAggchainManager() (common.Address, error) {
	return _AggchainFEP.Contract.PendingAggchainManager(&_AggchainFEP.CallOpts)
}

// PendingOptimisticModeManager is a free data retrieval call binding the contract method 0xadb8696c.
//
// Solidity: function pendingOptimisticModeManager() view returns(address)
func (_AggchainFEP *AggchainFEPCaller) PendingOptimisticModeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "pendingOptimisticModeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOptimisticModeManager is a free data retrieval call binding the contract method 0xadb8696c.
//
// Solidity: function pendingOptimisticModeManager() view returns(address)
func (_AggchainFEP *AggchainFEPSession) PendingOptimisticModeManager() (common.Address, error) {
	return _AggchainFEP.Contract.PendingOptimisticModeManager(&_AggchainFEP.CallOpts)
}

// PendingOptimisticModeManager is a free data retrieval call binding the contract method 0xadb8696c.
//
// Solidity: function pendingOptimisticModeManager() view returns(address)
func (_AggchainFEP *AggchainFEPCallerSession) PendingOptimisticModeManager() (common.Address, error) {
	return _AggchainFEP.Contract.PendingOptimisticModeManager(&_AggchainFEP.CallOpts)
}

// PendingVKeyManager is a free data retrieval call binding the contract method 0xbfb193b6.
//
// Solidity: function pendingVKeyManager() view returns(address)
func (_AggchainFEP *AggchainFEPCaller) PendingVKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "pendingVKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingVKeyManager is a free data retrieval call binding the contract method 0xbfb193b6.
//
// Solidity: function pendingVKeyManager() view returns(address)
func (_AggchainFEP *AggchainFEPSession) PendingVKeyManager() (common.Address, error) {
	return _AggchainFEP.Contract.PendingVKeyManager(&_AggchainFEP.CallOpts)
}

// PendingVKeyManager is a free data retrieval call binding the contract method 0xbfb193b6.
//
// Solidity: function pendingVKeyManager() view returns(address)
func (_AggchainFEP *AggchainFEPCallerSession) PendingVKeyManager() (common.Address, error) {
	return _AggchainFEP.Contract.PendingVKeyManager(&_AggchainFEP.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_AggchainFEP *AggchainFEPCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_AggchainFEP *AggchainFEPSession) Pol() (common.Address, error) {
	return _AggchainFEP.Contract.Pol(&_AggchainFEP.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_AggchainFEP *AggchainFEPCallerSession) Pol() (common.Address, error) {
	return _AggchainFEP.Contract.Pol(&_AggchainFEP.CallOpts)
}

// RangeVkeyCommitment is a free data retrieval call binding the contract method 0x2b31841e.
//
// Solidity: function rangeVkeyCommitment() view returns(bytes32)
func (_AggchainFEP *AggchainFEPCaller) RangeVkeyCommitment(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "rangeVkeyCommitment")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RangeVkeyCommitment is a free data retrieval call binding the contract method 0x2b31841e.
//
// Solidity: function rangeVkeyCommitment() view returns(bytes32)
func (_AggchainFEP *AggchainFEPSession) RangeVkeyCommitment() ([32]byte, error) {
	return _AggchainFEP.Contract.RangeVkeyCommitment(&_AggchainFEP.CallOpts)
}

// RangeVkeyCommitment is a free data retrieval call binding the contract method 0x2b31841e.
//
// Solidity: function rangeVkeyCommitment() view returns(bytes32)
func (_AggchainFEP *AggchainFEPCallerSession) RangeVkeyCommitment() ([32]byte, error) {
	return _AggchainFEP.Contract.RangeVkeyCommitment(&_AggchainFEP.CallOpts)
}

// RollupConfigHash is a free data retrieval call binding the contract method 0x6d9a1c8b.
//
// Solidity: function rollupConfigHash() view returns(bytes32)
func (_AggchainFEP *AggchainFEPCaller) RollupConfigHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "rollupConfigHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RollupConfigHash is a free data retrieval call binding the contract method 0x6d9a1c8b.
//
// Solidity: function rollupConfigHash() view returns(bytes32)
func (_AggchainFEP *AggchainFEPSession) RollupConfigHash() ([32]byte, error) {
	return _AggchainFEP.Contract.RollupConfigHash(&_AggchainFEP.CallOpts)
}

// RollupConfigHash is a free data retrieval call binding the contract method 0x6d9a1c8b.
//
// Solidity: function rollupConfigHash() view returns(bytes32)
func (_AggchainFEP *AggchainFEPCallerSession) RollupConfigHash() ([32]byte, error) {
	return _AggchainFEP.Contract.RollupConfigHash(&_AggchainFEP.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_AggchainFEP *AggchainFEPCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_AggchainFEP *AggchainFEPSession) RollupManager() (common.Address, error) {
	return _AggchainFEP.Contract.RollupManager(&_AggchainFEP.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_AggchainFEP *AggchainFEPCallerSession) RollupManager() (common.Address, error) {
	return _AggchainFEP.Contract.RollupManager(&_AggchainFEP.CallOpts)
}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_AggchainFEP *AggchainFEPCaller) StartingBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "startingBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_AggchainFEP *AggchainFEPSession) StartingBlockNumber() (*big.Int, error) {
	return _AggchainFEP.Contract.StartingBlockNumber(&_AggchainFEP.CallOpts)
}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_AggchainFEP *AggchainFEPCallerSession) StartingBlockNumber() (*big.Int, error) {
	return _AggchainFEP.Contract.StartingBlockNumber(&_AggchainFEP.CallOpts)
}

// StartingTimestamp is a free data retrieval call binding the contract method 0x88786272.
//
// Solidity: function startingTimestamp() view returns(uint256)
func (_AggchainFEP *AggchainFEPCaller) StartingTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "startingTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingTimestamp is a free data retrieval call binding the contract method 0x88786272.
//
// Solidity: function startingTimestamp() view returns(uint256)
func (_AggchainFEP *AggchainFEPSession) StartingTimestamp() (*big.Int, error) {
	return _AggchainFEP.Contract.StartingTimestamp(&_AggchainFEP.CallOpts)
}

// StartingTimestamp is a free data retrieval call binding the contract method 0x88786272.
//
// Solidity: function startingTimestamp() view returns(uint256)
func (_AggchainFEP *AggchainFEPCallerSession) StartingTimestamp() (*big.Int, error) {
	return _AggchainFEP.Contract.StartingTimestamp(&_AggchainFEP.CallOpts)
}

// SubmissionInterval is a free data retrieval call binding the contract method 0xe1a41bcf.
//
// Solidity: function submissionInterval() view returns(uint256)
func (_AggchainFEP *AggchainFEPCaller) SubmissionInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "submissionInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmissionInterval is a free data retrieval call binding the contract method 0xe1a41bcf.
//
// Solidity: function submissionInterval() view returns(uint256)
func (_AggchainFEP *AggchainFEPSession) SubmissionInterval() (*big.Int, error) {
	return _AggchainFEP.Contract.SubmissionInterval(&_AggchainFEP.CallOpts)
}

// SubmissionInterval is a free data retrieval call binding the contract method 0xe1a41bcf.
//
// Solidity: function submissionInterval() view returns(uint256)
func (_AggchainFEP *AggchainFEPCallerSession) SubmissionInterval() (*big.Int, error) {
	return _AggchainFEP.Contract.SubmissionInterval(&_AggchainFEP.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_AggchainFEP *AggchainFEPCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_AggchainFEP *AggchainFEPSession) TrustedSequencer() (common.Address, error) {
	return _AggchainFEP.Contract.TrustedSequencer(&_AggchainFEP.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_AggchainFEP *AggchainFEPCallerSession) TrustedSequencer() (common.Address, error) {
	return _AggchainFEP.Contract.TrustedSequencer(&_AggchainFEP.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_AggchainFEP *AggchainFEPCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_AggchainFEP *AggchainFEPSession) TrustedSequencerURL() (string, error) {
	return _AggchainFEP.Contract.TrustedSequencerURL(&_AggchainFEP.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_AggchainFEP *AggchainFEPCallerSession) TrustedSequencerURL() (string, error) {
	return _AggchainFEP.Contract.TrustedSequencerURL(&_AggchainFEP.CallOpts)
}

// UseDefaultGateway is a free data retrieval call binding the contract method 0xff904079.
//
// Solidity: function useDefaultGateway() view returns(bool)
func (_AggchainFEP *AggchainFEPCaller) UseDefaultGateway(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "useDefaultGateway")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseDefaultGateway is a free data retrieval call binding the contract method 0xff904079.
//
// Solidity: function useDefaultGateway() view returns(bool)
func (_AggchainFEP *AggchainFEPSession) UseDefaultGateway() (bool, error) {
	return _AggchainFEP.Contract.UseDefaultGateway(&_AggchainFEP.CallOpts)
}

// UseDefaultGateway is a free data retrieval call binding the contract method 0xff904079.
//
// Solidity: function useDefaultGateway() view returns(bool)
func (_AggchainFEP *AggchainFEPCallerSession) UseDefaultGateway() (bool, error) {
	return _AggchainFEP.Contract.UseDefaultGateway(&_AggchainFEP.CallOpts)
}

// VKeyManager is a free data retrieval call binding the contract method 0xe279984e.
//
// Solidity: function vKeyManager() view returns(address)
func (_AggchainFEP *AggchainFEPCaller) VKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "vKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VKeyManager is a free data retrieval call binding the contract method 0xe279984e.
//
// Solidity: function vKeyManager() view returns(address)
func (_AggchainFEP *AggchainFEPSession) VKeyManager() (common.Address, error) {
	return _AggchainFEP.Contract.VKeyManager(&_AggchainFEP.CallOpts)
}

// VKeyManager is a free data retrieval call binding the contract method 0xe279984e.
//
// Solidity: function vKeyManager() view returns(address)
func (_AggchainFEP *AggchainFEPCallerSession) VKeyManager() (common.Address, error) {
	return _AggchainFEP.Contract.VKeyManager(&_AggchainFEP.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_AggchainFEP *AggchainFEPCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AggchainFEP.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_AggchainFEP *AggchainFEPSession) Version() (string, error) {
	return _AggchainFEP.Contract.Version(&_AggchainFEP.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_AggchainFEP *AggchainFEPCallerSession) Version() (string, error) {
	return _AggchainFEP.Contract.Version(&_AggchainFEP.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_AggchainFEP *AggchainFEPTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_AggchainFEP *AggchainFEPSession) AcceptAdminRole() (*types.Transaction, error) {
	return _AggchainFEP.Contract.AcceptAdminRole(&_AggchainFEP.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_AggchainFEP *AggchainFEPTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _AggchainFEP.Contract.AcceptAdminRole(&_AggchainFEP.TransactOpts)
}

// AcceptAggchainManagerRole is a paid mutator transaction binding the contract method 0x15981b29.
//
// Solidity: function acceptAggchainManagerRole() returns()
func (_AggchainFEP *AggchainFEPTransactor) AcceptAggchainManagerRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "acceptAggchainManagerRole")
}

// AcceptAggchainManagerRole is a paid mutator transaction binding the contract method 0x15981b29.
//
// Solidity: function acceptAggchainManagerRole() returns()
func (_AggchainFEP *AggchainFEPSession) AcceptAggchainManagerRole() (*types.Transaction, error) {
	return _AggchainFEP.Contract.AcceptAggchainManagerRole(&_AggchainFEP.TransactOpts)
}

// AcceptAggchainManagerRole is a paid mutator transaction binding the contract method 0x15981b29.
//
// Solidity: function acceptAggchainManagerRole() returns()
func (_AggchainFEP *AggchainFEPTransactorSession) AcceptAggchainManagerRole() (*types.Transaction, error) {
	return _AggchainFEP.Contract.AcceptAggchainManagerRole(&_AggchainFEP.TransactOpts)
}

// AcceptOptimisticModeManagerRole is a paid mutator transaction binding the contract method 0x12634900.
//
// Solidity: function acceptOptimisticModeManagerRole() returns()
func (_AggchainFEP *AggchainFEPTransactor) AcceptOptimisticModeManagerRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "acceptOptimisticModeManagerRole")
}

// AcceptOptimisticModeManagerRole is a paid mutator transaction binding the contract method 0x12634900.
//
// Solidity: function acceptOptimisticModeManagerRole() returns()
func (_AggchainFEP *AggchainFEPSession) AcceptOptimisticModeManagerRole() (*types.Transaction, error) {
	return _AggchainFEP.Contract.AcceptOptimisticModeManagerRole(&_AggchainFEP.TransactOpts)
}

// AcceptOptimisticModeManagerRole is a paid mutator transaction binding the contract method 0x12634900.
//
// Solidity: function acceptOptimisticModeManagerRole() returns()
func (_AggchainFEP *AggchainFEPTransactorSession) AcceptOptimisticModeManagerRole() (*types.Transaction, error) {
	return _AggchainFEP.Contract.AcceptOptimisticModeManagerRole(&_AggchainFEP.TransactOpts)
}

// AcceptVKeyManagerRole is a paid mutator transaction binding the contract method 0x368c822c.
//
// Solidity: function acceptVKeyManagerRole() returns()
func (_AggchainFEP *AggchainFEPTransactor) AcceptVKeyManagerRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "acceptVKeyManagerRole")
}

// AcceptVKeyManagerRole is a paid mutator transaction binding the contract method 0x368c822c.
//
// Solidity: function acceptVKeyManagerRole() returns()
func (_AggchainFEP *AggchainFEPSession) AcceptVKeyManagerRole() (*types.Transaction, error) {
	return _AggchainFEP.Contract.AcceptVKeyManagerRole(&_AggchainFEP.TransactOpts)
}

// AcceptVKeyManagerRole is a paid mutator transaction binding the contract method 0x368c822c.
//
// Solidity: function acceptVKeyManagerRole() returns()
func (_AggchainFEP *AggchainFEPTransactorSession) AcceptVKeyManagerRole() (*types.Transaction, error) {
	return _AggchainFEP.Contract.AcceptVKeyManagerRole(&_AggchainFEP.TransactOpts)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 newAggchainVKey) returns()
func (_AggchainFEP *AggchainFEPTransactor) AddOwnedAggchainVKey(opts *bind.TransactOpts, aggchainVKeySelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "addOwnedAggchainVKey", aggchainVKeySelector, newAggchainVKey)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 newAggchainVKey) returns()
func (_AggchainFEP *AggchainFEPSession) AddOwnedAggchainVKey(aggchainVKeySelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _AggchainFEP.Contract.AddOwnedAggchainVKey(&_AggchainFEP.TransactOpts, aggchainVKeySelector, newAggchainVKey)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 newAggchainVKey) returns()
func (_AggchainFEP *AggchainFEPTransactorSession) AddOwnedAggchainVKey(aggchainVKeySelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _AggchainFEP.Contract.AddOwnedAggchainVKey(&_AggchainFEP.TransactOpts, aggchainVKeySelector, newAggchainVKey)
}

// DisableOptimisticMode is a paid mutator transaction binding the contract method 0x0822dc61.
//
// Solidity: function disableOptimisticMode() returns()
func (_AggchainFEP *AggchainFEPTransactor) DisableOptimisticMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "disableOptimisticMode")
}

// DisableOptimisticMode is a paid mutator transaction binding the contract method 0x0822dc61.
//
// Solidity: function disableOptimisticMode() returns()
func (_AggchainFEP *AggchainFEPSession) DisableOptimisticMode() (*types.Transaction, error) {
	return _AggchainFEP.Contract.DisableOptimisticMode(&_AggchainFEP.TransactOpts)
}

// DisableOptimisticMode is a paid mutator transaction binding the contract method 0x0822dc61.
//
// Solidity: function disableOptimisticMode() returns()
func (_AggchainFEP *AggchainFEPTransactorSession) DisableOptimisticMode() (*types.Transaction, error) {
	return _AggchainFEP.Contract.DisableOptimisticMode(&_AggchainFEP.TransactOpts)
}

// DisableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xdc8c4249.
//
// Solidity: function disableUseDefaultGatewayFlag() returns()
func (_AggchainFEP *AggchainFEPTransactor) DisableUseDefaultGatewayFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "disableUseDefaultGatewayFlag")
}

// DisableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xdc8c4249.
//
// Solidity: function disableUseDefaultGatewayFlag() returns()
func (_AggchainFEP *AggchainFEPSession) DisableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _AggchainFEP.Contract.DisableUseDefaultGatewayFlag(&_AggchainFEP.TransactOpts)
}

// DisableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xdc8c4249.
//
// Solidity: function disableUseDefaultGatewayFlag() returns()
func (_AggchainFEP *AggchainFEPTransactorSession) DisableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _AggchainFEP.Contract.DisableUseDefaultGatewayFlag(&_AggchainFEP.TransactOpts)
}

// EnableOptimisticMode is a paid mutator transaction binding the contract method 0x81eb0baf.
//
// Solidity: function enableOptimisticMode() returns()
func (_AggchainFEP *AggchainFEPTransactor) EnableOptimisticMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "enableOptimisticMode")
}

// EnableOptimisticMode is a paid mutator transaction binding the contract method 0x81eb0baf.
//
// Solidity: function enableOptimisticMode() returns()
func (_AggchainFEP *AggchainFEPSession) EnableOptimisticMode() (*types.Transaction, error) {
	return _AggchainFEP.Contract.EnableOptimisticMode(&_AggchainFEP.TransactOpts)
}

// EnableOptimisticMode is a paid mutator transaction binding the contract method 0x81eb0baf.
//
// Solidity: function enableOptimisticMode() returns()
func (_AggchainFEP *AggchainFEPTransactorSession) EnableOptimisticMode() (*types.Transaction, error) {
	return _AggchainFEP.Contract.EnableOptimisticMode(&_AggchainFEP.TransactOpts)
}

// EnableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xe631476c.
//
// Solidity: function enableUseDefaultGatewayFlag() returns()
func (_AggchainFEP *AggchainFEPTransactor) EnableUseDefaultGatewayFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "enableUseDefaultGatewayFlag")
}

// EnableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xe631476c.
//
// Solidity: function enableUseDefaultGatewayFlag() returns()
func (_AggchainFEP *AggchainFEPSession) EnableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _AggchainFEP.Contract.EnableUseDefaultGatewayFlag(&_AggchainFEP.TransactOpts)
}

// EnableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xe631476c.
//
// Solidity: function enableUseDefaultGatewayFlag() returns()
func (_AggchainFEP *AggchainFEPTransactorSession) EnableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _AggchainFEP.Contract.EnableUseDefaultGatewayFlag(&_AggchainFEP.TransactOpts)
}

// InitAggchainManager is a paid mutator transaction binding the contract method 0xb3a326f7.
//
// Solidity: function initAggchainManager(address newAggchainManager) returns()
func (_AggchainFEP *AggchainFEPTransactor) InitAggchainManager(opts *bind.TransactOpts, newAggchainManager common.Address) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "initAggchainManager", newAggchainManager)
}

// InitAggchainManager is a paid mutator transaction binding the contract method 0xb3a326f7.
//
// Solidity: function initAggchainManager(address newAggchainManager) returns()
func (_AggchainFEP *AggchainFEPSession) InitAggchainManager(newAggchainManager common.Address) (*types.Transaction, error) {
	return _AggchainFEP.Contract.InitAggchainManager(&_AggchainFEP.TransactOpts, newAggchainManager)
}

// InitAggchainManager is a paid mutator transaction binding the contract method 0xb3a326f7.
//
// Solidity: function initAggchainManager(address newAggchainManager) returns()
func (_AggchainFEP *AggchainFEPTransactorSession) InitAggchainManager(newAggchainManager common.Address) (*types.Transaction, error) {
	return _AggchainFEP.Contract.InitAggchainManager(&_AggchainFEP.TransactOpts, newAggchainManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesAggchain) returns()
func (_AggchainFEP *AggchainFEPTransactor) Initialize(opts *bind.TransactOpts, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "initialize", initializeBytesAggchain)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesAggchain) returns()
func (_AggchainFEP *AggchainFEPSession) Initialize(initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _AggchainFEP.Contract.Initialize(&_AggchainFEP.TransactOpts, initializeBytesAggchain)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesAggchain) returns()
func (_AggchainFEP *AggchainFEPTransactorSession) Initialize(initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _AggchainFEP.Contract.Initialize(&_AggchainFEP.TransactOpts, initializeBytesAggchain)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes aggchainData) returns()
func (_AggchainFEP *AggchainFEPTransactor) OnVerifyPessimistic(opts *bind.TransactOpts, aggchainData []byte) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "onVerifyPessimistic", aggchainData)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes aggchainData) returns()
func (_AggchainFEP *AggchainFEPSession) OnVerifyPessimistic(aggchainData []byte) (*types.Transaction, error) {
	return _AggchainFEP.Contract.OnVerifyPessimistic(&_AggchainFEP.TransactOpts, aggchainData)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes aggchainData) returns()
func (_AggchainFEP *AggchainFEPTransactorSession) OnVerifyPessimistic(aggchainData []byte) (*types.Transaction, error) {
	return _AggchainFEP.Contract.OnVerifyPessimistic(&_AggchainFEP.TransactOpts, aggchainData)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_AggchainFEP *AggchainFEPTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_AggchainFEP *AggchainFEPSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _AggchainFEP.Contract.SetTrustedSequencer(&_AggchainFEP.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_AggchainFEP *AggchainFEPTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _AggchainFEP.Contract.SetTrustedSequencer(&_AggchainFEP.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_AggchainFEP *AggchainFEPTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_AggchainFEP *AggchainFEPSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _AggchainFEP.Contract.SetTrustedSequencerURL(&_AggchainFEP.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_AggchainFEP *AggchainFEPTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _AggchainFEP.Contract.SetTrustedSequencerURL(&_AggchainFEP.TransactOpts, newTrustedSequencerURL)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_AggchainFEP *AggchainFEPTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_AggchainFEP *AggchainFEPSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _AggchainFEP.Contract.TransferAdminRole(&_AggchainFEP.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_AggchainFEP *AggchainFEPTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _AggchainFEP.Contract.TransferAdminRole(&_AggchainFEP.TransactOpts, newPendingAdmin)
}

// TransferAggchainManagerRole is a paid mutator transaction binding the contract method 0xbdfbed7e.
//
// Solidity: function transferAggchainManagerRole(address newAggchainManager) returns()
func (_AggchainFEP *AggchainFEPTransactor) TransferAggchainManagerRole(opts *bind.TransactOpts, newAggchainManager common.Address) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "transferAggchainManagerRole", newAggchainManager)
}

// TransferAggchainManagerRole is a paid mutator transaction binding the contract method 0xbdfbed7e.
//
// Solidity: function transferAggchainManagerRole(address newAggchainManager) returns()
func (_AggchainFEP *AggchainFEPSession) TransferAggchainManagerRole(newAggchainManager common.Address) (*types.Transaction, error) {
	return _AggchainFEP.Contract.TransferAggchainManagerRole(&_AggchainFEP.TransactOpts, newAggchainManager)
}

// TransferAggchainManagerRole is a paid mutator transaction binding the contract method 0xbdfbed7e.
//
// Solidity: function transferAggchainManagerRole(address newAggchainManager) returns()
func (_AggchainFEP *AggchainFEPTransactorSession) TransferAggchainManagerRole(newAggchainManager common.Address) (*types.Transaction, error) {
	return _AggchainFEP.Contract.TransferAggchainManagerRole(&_AggchainFEP.TransactOpts, newAggchainManager)
}

// TransferOptimisticModeManagerRole is a paid mutator transaction binding the contract method 0xfdbbc19b.
//
// Solidity: function transferOptimisticModeManagerRole(address newOptimisticModeManager) returns()
func (_AggchainFEP *AggchainFEPTransactor) TransferOptimisticModeManagerRole(opts *bind.TransactOpts, newOptimisticModeManager common.Address) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "transferOptimisticModeManagerRole", newOptimisticModeManager)
}

// TransferOptimisticModeManagerRole is a paid mutator transaction binding the contract method 0xfdbbc19b.
//
// Solidity: function transferOptimisticModeManagerRole(address newOptimisticModeManager) returns()
func (_AggchainFEP *AggchainFEPSession) TransferOptimisticModeManagerRole(newOptimisticModeManager common.Address) (*types.Transaction, error) {
	return _AggchainFEP.Contract.TransferOptimisticModeManagerRole(&_AggchainFEP.TransactOpts, newOptimisticModeManager)
}

// TransferOptimisticModeManagerRole is a paid mutator transaction binding the contract method 0xfdbbc19b.
//
// Solidity: function transferOptimisticModeManagerRole(address newOptimisticModeManager) returns()
func (_AggchainFEP *AggchainFEPTransactorSession) TransferOptimisticModeManagerRole(newOptimisticModeManager common.Address) (*types.Transaction, error) {
	return _AggchainFEP.Contract.TransferOptimisticModeManagerRole(&_AggchainFEP.TransactOpts, newOptimisticModeManager)
}

// TransferVKeyManagerRole is a paid mutator transaction binding the contract method 0x85018182.
//
// Solidity: function transferVKeyManagerRole(address newVKeyManager) returns()
func (_AggchainFEP *AggchainFEPTransactor) TransferVKeyManagerRole(opts *bind.TransactOpts, newVKeyManager common.Address) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "transferVKeyManagerRole", newVKeyManager)
}

// TransferVKeyManagerRole is a paid mutator transaction binding the contract method 0x85018182.
//
// Solidity: function transferVKeyManagerRole(address newVKeyManager) returns()
func (_AggchainFEP *AggchainFEPSession) TransferVKeyManagerRole(newVKeyManager common.Address) (*types.Transaction, error) {
	return _AggchainFEP.Contract.TransferVKeyManagerRole(&_AggchainFEP.TransactOpts, newVKeyManager)
}

// TransferVKeyManagerRole is a paid mutator transaction binding the contract method 0x85018182.
//
// Solidity: function transferVKeyManagerRole(address newVKeyManager) returns()
func (_AggchainFEP *AggchainFEPTransactorSession) TransferVKeyManagerRole(newVKeyManager common.Address) (*types.Transaction, error) {
	return _AggchainFEP.Contract.TransferVKeyManagerRole(&_AggchainFEP.TransactOpts, newVKeyManager)
}

// UpdateAggregationVkey is a paid mutator transaction binding the contract method 0xc4cb03ec.
//
// Solidity: function updateAggregationVkey(bytes32 _aggregationVkey) returns()
func (_AggchainFEP *AggchainFEPTransactor) UpdateAggregationVkey(opts *bind.TransactOpts, _aggregationVkey [32]byte) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "updateAggregationVkey", _aggregationVkey)
}

// UpdateAggregationVkey is a paid mutator transaction binding the contract method 0xc4cb03ec.
//
// Solidity: function updateAggregationVkey(bytes32 _aggregationVkey) returns()
func (_AggchainFEP *AggchainFEPSession) UpdateAggregationVkey(_aggregationVkey [32]byte) (*types.Transaction, error) {
	return _AggchainFEP.Contract.UpdateAggregationVkey(&_AggchainFEP.TransactOpts, _aggregationVkey)
}

// UpdateAggregationVkey is a paid mutator transaction binding the contract method 0xc4cb03ec.
//
// Solidity: function updateAggregationVkey(bytes32 _aggregationVkey) returns()
func (_AggchainFEP *AggchainFEPTransactorSession) UpdateAggregationVkey(_aggregationVkey [32]byte) (*types.Transaction, error) {
	return _AggchainFEP.Contract.UpdateAggregationVkey(&_AggchainFEP.TransactOpts, _aggregationVkey)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 updatedAggchainVKey) returns()
func (_AggchainFEP *AggchainFEPTransactor) UpdateOwnedAggchainVKey(opts *bind.TransactOpts, aggchainVKeySelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "updateOwnedAggchainVKey", aggchainVKeySelector, updatedAggchainVKey)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 updatedAggchainVKey) returns()
func (_AggchainFEP *AggchainFEPSession) UpdateOwnedAggchainVKey(aggchainVKeySelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _AggchainFEP.Contract.UpdateOwnedAggchainVKey(&_AggchainFEP.TransactOpts, aggchainVKeySelector, updatedAggchainVKey)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 updatedAggchainVKey) returns()
func (_AggchainFEP *AggchainFEPTransactorSession) UpdateOwnedAggchainVKey(aggchainVKeySelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _AggchainFEP.Contract.UpdateOwnedAggchainVKey(&_AggchainFEP.TransactOpts, aggchainVKeySelector, updatedAggchainVKey)
}

// UpdateRangeVkeyCommitment is a paid mutator transaction binding the contract method 0xbc91ce33.
//
// Solidity: function updateRangeVkeyCommitment(bytes32 _rangeVkeyCommitment) returns()
func (_AggchainFEP *AggchainFEPTransactor) UpdateRangeVkeyCommitment(opts *bind.TransactOpts, _rangeVkeyCommitment [32]byte) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "updateRangeVkeyCommitment", _rangeVkeyCommitment)
}

// UpdateRangeVkeyCommitment is a paid mutator transaction binding the contract method 0xbc91ce33.
//
// Solidity: function updateRangeVkeyCommitment(bytes32 _rangeVkeyCommitment) returns()
func (_AggchainFEP *AggchainFEPSession) UpdateRangeVkeyCommitment(_rangeVkeyCommitment [32]byte) (*types.Transaction, error) {
	return _AggchainFEP.Contract.UpdateRangeVkeyCommitment(&_AggchainFEP.TransactOpts, _rangeVkeyCommitment)
}

// UpdateRangeVkeyCommitment is a paid mutator transaction binding the contract method 0xbc91ce33.
//
// Solidity: function updateRangeVkeyCommitment(bytes32 _rangeVkeyCommitment) returns()
func (_AggchainFEP *AggchainFEPTransactorSession) UpdateRangeVkeyCommitment(_rangeVkeyCommitment [32]byte) (*types.Transaction, error) {
	return _AggchainFEP.Contract.UpdateRangeVkeyCommitment(&_AggchainFEP.TransactOpts, _rangeVkeyCommitment)
}

// UpdateRollupConfigHash is a paid mutator transaction binding the contract method 0x1bdd450c.
//
// Solidity: function updateRollupConfigHash(bytes32 _rollupConfigHash) returns()
func (_AggchainFEP *AggchainFEPTransactor) UpdateRollupConfigHash(opts *bind.TransactOpts, _rollupConfigHash [32]byte) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "updateRollupConfigHash", _rollupConfigHash)
}

// UpdateRollupConfigHash is a paid mutator transaction binding the contract method 0x1bdd450c.
//
// Solidity: function updateRollupConfigHash(bytes32 _rollupConfigHash) returns()
func (_AggchainFEP *AggchainFEPSession) UpdateRollupConfigHash(_rollupConfigHash [32]byte) (*types.Transaction, error) {
	return _AggchainFEP.Contract.UpdateRollupConfigHash(&_AggchainFEP.TransactOpts, _rollupConfigHash)
}

// UpdateRollupConfigHash is a paid mutator transaction binding the contract method 0x1bdd450c.
//
// Solidity: function updateRollupConfigHash(bytes32 _rollupConfigHash) returns()
func (_AggchainFEP *AggchainFEPTransactorSession) UpdateRollupConfigHash(_rollupConfigHash [32]byte) (*types.Transaction, error) {
	return _AggchainFEP.Contract.UpdateRollupConfigHash(&_AggchainFEP.TransactOpts, _rollupConfigHash)
}

// UpdateSubmissionInterval is a paid mutator transaction binding the contract method 0x336c9e81.
//
// Solidity: function updateSubmissionInterval(uint256 _submissionInterval) returns()
func (_AggchainFEP *AggchainFEPTransactor) UpdateSubmissionInterval(opts *bind.TransactOpts, _submissionInterval *big.Int) (*types.Transaction, error) {
	return _AggchainFEP.contract.Transact(opts, "updateSubmissionInterval", _submissionInterval)
}

// UpdateSubmissionInterval is a paid mutator transaction binding the contract method 0x336c9e81.
//
// Solidity: function updateSubmissionInterval(uint256 _submissionInterval) returns()
func (_AggchainFEP *AggchainFEPSession) UpdateSubmissionInterval(_submissionInterval *big.Int) (*types.Transaction, error) {
	return _AggchainFEP.Contract.UpdateSubmissionInterval(&_AggchainFEP.TransactOpts, _submissionInterval)
}

// UpdateSubmissionInterval is a paid mutator transaction binding the contract method 0x336c9e81.
//
// Solidity: function updateSubmissionInterval(uint256 _submissionInterval) returns()
func (_AggchainFEP *AggchainFEPTransactorSession) UpdateSubmissionInterval(_submissionInterval *big.Int) (*types.Transaction, error) {
	return _AggchainFEP.Contract.UpdateSubmissionInterval(&_AggchainFEP.TransactOpts, _submissionInterval)
}

// AggchainFEPAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the AggchainFEP contract.
type AggchainFEPAcceptAdminRoleIterator struct {
	Event *AggchainFEPAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *AggchainFEPAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPAcceptAdminRole)
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
		it.Event = new(AggchainFEPAcceptAdminRole)
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
func (it *AggchainFEPAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPAcceptAdminRole represents a AcceptAdminRole event raised by the AggchainFEP contract.
type AggchainFEPAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_AggchainFEP *AggchainFEPFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*AggchainFEPAcceptAdminRoleIterator, error) {

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &AggchainFEPAcceptAdminRoleIterator{contract: _AggchainFEP.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_AggchainFEP *AggchainFEPFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *AggchainFEPAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPAcceptAdminRole)
				if err := _AggchainFEP.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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

// ParseAcceptAdminRole is a log parse operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_AggchainFEP *AggchainFEPFilterer) ParseAcceptAdminRole(log types.Log) (*AggchainFEPAcceptAdminRole, error) {
	event := new(AggchainFEPAcceptAdminRole)
	if err := _AggchainFEP.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPAcceptAggchainManagerRoleIterator is returned from FilterAcceptAggchainManagerRole and is used to iterate over the raw logs and unpacked data for AcceptAggchainManagerRole events raised by the AggchainFEP contract.
type AggchainFEPAcceptAggchainManagerRoleIterator struct {
	Event *AggchainFEPAcceptAggchainManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainFEPAcceptAggchainManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPAcceptAggchainManagerRole)
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
		it.Event = new(AggchainFEPAcceptAggchainManagerRole)
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
func (it *AggchainFEPAcceptAggchainManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPAcceptAggchainManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPAcceptAggchainManagerRole represents a AcceptAggchainManagerRole event raised by the AggchainFEP contract.
type AggchainFEPAcceptAggchainManagerRole struct {
	OldAggchainManager common.Address
	NewAggchainManager common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAcceptAggchainManagerRole is a free log retrieval operation binding the contract event 0x67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f.
//
// Solidity: event AcceptAggchainManagerRole(address oldAggchainManager, address newAggchainManager)
func (_AggchainFEP *AggchainFEPFilterer) FilterAcceptAggchainManagerRole(opts *bind.FilterOpts) (*AggchainFEPAcceptAggchainManagerRoleIterator, error) {

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "AcceptAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainFEPAcceptAggchainManagerRoleIterator{contract: _AggchainFEP.contract, event: "AcceptAggchainManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAggchainManagerRole is a free log subscription operation binding the contract event 0x67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f.
//
// Solidity: event AcceptAggchainManagerRole(address oldAggchainManager, address newAggchainManager)
func (_AggchainFEP *AggchainFEPFilterer) WatchAcceptAggchainManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainFEPAcceptAggchainManagerRole) (event.Subscription, error) {

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "AcceptAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPAcceptAggchainManagerRole)
				if err := _AggchainFEP.contract.UnpackLog(event, "AcceptAggchainManagerRole", log); err != nil {
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

// ParseAcceptAggchainManagerRole is a log parse operation binding the contract event 0x67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f.
//
// Solidity: event AcceptAggchainManagerRole(address oldAggchainManager, address newAggchainManager)
func (_AggchainFEP *AggchainFEPFilterer) ParseAcceptAggchainManagerRole(log types.Log) (*AggchainFEPAcceptAggchainManagerRole, error) {
	event := new(AggchainFEPAcceptAggchainManagerRole)
	if err := _AggchainFEP.contract.UnpackLog(event, "AcceptAggchainManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPAcceptOptimisticModeManagerRoleIterator is returned from FilterAcceptOptimisticModeManagerRole and is used to iterate over the raw logs and unpacked data for AcceptOptimisticModeManagerRole events raised by the AggchainFEP contract.
type AggchainFEPAcceptOptimisticModeManagerRoleIterator struct {
	Event *AggchainFEPAcceptOptimisticModeManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainFEPAcceptOptimisticModeManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPAcceptOptimisticModeManagerRole)
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
		it.Event = new(AggchainFEPAcceptOptimisticModeManagerRole)
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
func (it *AggchainFEPAcceptOptimisticModeManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPAcceptOptimisticModeManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPAcceptOptimisticModeManagerRole represents a AcceptOptimisticModeManagerRole event raised by the AggchainFEP contract.
type AggchainFEPAcceptOptimisticModeManagerRole struct {
	OldOptimisticModeManager common.Address
	NewOptimisticModeManager common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterAcceptOptimisticModeManagerRole is a free log retrieval operation binding the contract event 0x9a58f1fef974b760afdc36e96f8d4af9162ba9fec7cd8ce7ca397aa3399f3319.
//
// Solidity: event AcceptOptimisticModeManagerRole(address oldOptimisticModeManager, address newOptimisticModeManager)
func (_AggchainFEP *AggchainFEPFilterer) FilterAcceptOptimisticModeManagerRole(opts *bind.FilterOpts) (*AggchainFEPAcceptOptimisticModeManagerRoleIterator, error) {

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "AcceptOptimisticModeManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainFEPAcceptOptimisticModeManagerRoleIterator{contract: _AggchainFEP.contract, event: "AcceptOptimisticModeManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptOptimisticModeManagerRole is a free log subscription operation binding the contract event 0x9a58f1fef974b760afdc36e96f8d4af9162ba9fec7cd8ce7ca397aa3399f3319.
//
// Solidity: event AcceptOptimisticModeManagerRole(address oldOptimisticModeManager, address newOptimisticModeManager)
func (_AggchainFEP *AggchainFEPFilterer) WatchAcceptOptimisticModeManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainFEPAcceptOptimisticModeManagerRole) (event.Subscription, error) {

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "AcceptOptimisticModeManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPAcceptOptimisticModeManagerRole)
				if err := _AggchainFEP.contract.UnpackLog(event, "AcceptOptimisticModeManagerRole", log); err != nil {
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

// ParseAcceptOptimisticModeManagerRole is a log parse operation binding the contract event 0x9a58f1fef974b760afdc36e96f8d4af9162ba9fec7cd8ce7ca397aa3399f3319.
//
// Solidity: event AcceptOptimisticModeManagerRole(address oldOptimisticModeManager, address newOptimisticModeManager)
func (_AggchainFEP *AggchainFEPFilterer) ParseAcceptOptimisticModeManagerRole(log types.Log) (*AggchainFEPAcceptOptimisticModeManagerRole, error) {
	event := new(AggchainFEPAcceptOptimisticModeManagerRole)
	if err := _AggchainFEP.contract.UnpackLog(event, "AcceptOptimisticModeManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPAcceptVKeyManagerRoleIterator is returned from FilterAcceptVKeyManagerRole and is used to iterate over the raw logs and unpacked data for AcceptVKeyManagerRole events raised by the AggchainFEP contract.
type AggchainFEPAcceptVKeyManagerRoleIterator struct {
	Event *AggchainFEPAcceptVKeyManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainFEPAcceptVKeyManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPAcceptVKeyManagerRole)
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
		it.Event = new(AggchainFEPAcceptVKeyManagerRole)
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
func (it *AggchainFEPAcceptVKeyManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPAcceptVKeyManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPAcceptVKeyManagerRole represents a AcceptVKeyManagerRole event raised by the AggchainFEP contract.
type AggchainFEPAcceptVKeyManagerRole struct {
	OldVKeyManager common.Address
	NewVKeyManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAcceptVKeyManagerRole is a free log retrieval operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_AggchainFEP *AggchainFEPFilterer) FilterAcceptVKeyManagerRole(opts *bind.FilterOpts) (*AggchainFEPAcceptVKeyManagerRoleIterator, error) {

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainFEPAcceptVKeyManagerRoleIterator{contract: _AggchainFEP.contract, event: "AcceptVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptVKeyManagerRole is a free log subscription operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_AggchainFEP *AggchainFEPFilterer) WatchAcceptVKeyManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainFEPAcceptVKeyManagerRole) (event.Subscription, error) {

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPAcceptVKeyManagerRole)
				if err := _AggchainFEP.contract.UnpackLog(event, "AcceptVKeyManagerRole", log); err != nil {
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

// ParseAcceptVKeyManagerRole is a log parse operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_AggchainFEP *AggchainFEPFilterer) ParseAcceptVKeyManagerRole(log types.Log) (*AggchainFEPAcceptVKeyManagerRole, error) {
	event := new(AggchainFEPAcceptVKeyManagerRole)
	if err := _AggchainFEP.contract.UnpackLog(event, "AcceptVKeyManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPAddAggchainVKeyIterator is returned from FilterAddAggchainVKey and is used to iterate over the raw logs and unpacked data for AddAggchainVKey events raised by the AggchainFEP contract.
type AggchainFEPAddAggchainVKeyIterator struct {
	Event *AggchainFEPAddAggchainVKey // Event containing the contract specifics and raw log

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
func (it *AggchainFEPAddAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPAddAggchainVKey)
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
		it.Event = new(AggchainFEPAddAggchainVKey)
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
func (it *AggchainFEPAddAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPAddAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPAddAggchainVKey represents a AddAggchainVKey event raised by the AggchainFEP contract.
type AggchainFEPAddAggchainVKey struct {
	Selector        [4]byte
	NewAggchainVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddAggchainVKey is a free log retrieval operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_AggchainFEP *AggchainFEPFilterer) FilterAddAggchainVKey(opts *bind.FilterOpts) (*AggchainFEPAddAggchainVKeyIterator, error) {

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "AddAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AggchainFEPAddAggchainVKeyIterator{contract: _AggchainFEP.contract, event: "AddAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchAddAggchainVKey is a free log subscription operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_AggchainFEP *AggchainFEPFilterer) WatchAddAggchainVKey(opts *bind.WatchOpts, sink chan<- *AggchainFEPAddAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "AddAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPAddAggchainVKey)
				if err := _AggchainFEP.contract.UnpackLog(event, "AddAggchainVKey", log); err != nil {
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

// ParseAddAggchainVKey is a log parse operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_AggchainFEP *AggchainFEPFilterer) ParseAddAggchainVKey(log types.Log) (*AggchainFEPAddAggchainVKey, error) {
	event := new(AggchainFEPAddAggchainVKey)
	if err := _AggchainFEP.contract.UnpackLog(event, "AddAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPAggregationVkeyUpdatedIterator is returned from FilterAggregationVkeyUpdated and is used to iterate over the raw logs and unpacked data for AggregationVkeyUpdated events raised by the AggchainFEP contract.
type AggchainFEPAggregationVkeyUpdatedIterator struct {
	Event *AggchainFEPAggregationVkeyUpdated // Event containing the contract specifics and raw log

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
func (it *AggchainFEPAggregationVkeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPAggregationVkeyUpdated)
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
		it.Event = new(AggchainFEPAggregationVkeyUpdated)
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
func (it *AggchainFEPAggregationVkeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPAggregationVkeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPAggregationVkeyUpdated represents a AggregationVkeyUpdated event raised by the AggchainFEP contract.
type AggchainFEPAggregationVkeyUpdated struct {
	OldAggregationVkey [32]byte
	NewAggregationVkey [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAggregationVkeyUpdated is a free log retrieval operation binding the contract event 0x390b73b2b067afcef04d30b573e4590c6e565519e370927dd777ca0ce8a55db0.
//
// Solidity: event AggregationVkeyUpdated(bytes32 indexed oldAggregationVkey, bytes32 indexed newAggregationVkey)
func (_AggchainFEP *AggchainFEPFilterer) FilterAggregationVkeyUpdated(opts *bind.FilterOpts, oldAggregationVkey [][32]byte, newAggregationVkey [][32]byte) (*AggchainFEPAggregationVkeyUpdatedIterator, error) {

	var oldAggregationVkeyRule []interface{}
	for _, oldAggregationVkeyItem := range oldAggregationVkey {
		oldAggregationVkeyRule = append(oldAggregationVkeyRule, oldAggregationVkeyItem)
	}
	var newAggregationVkeyRule []interface{}
	for _, newAggregationVkeyItem := range newAggregationVkey {
		newAggregationVkeyRule = append(newAggregationVkeyRule, newAggregationVkeyItem)
	}

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "AggregationVkeyUpdated", oldAggregationVkeyRule, newAggregationVkeyRule)
	if err != nil {
		return nil, err
	}
	return &AggchainFEPAggregationVkeyUpdatedIterator{contract: _AggchainFEP.contract, event: "AggregationVkeyUpdated", logs: logs, sub: sub}, nil
}

// WatchAggregationVkeyUpdated is a free log subscription operation binding the contract event 0x390b73b2b067afcef04d30b573e4590c6e565519e370927dd777ca0ce8a55db0.
//
// Solidity: event AggregationVkeyUpdated(bytes32 indexed oldAggregationVkey, bytes32 indexed newAggregationVkey)
func (_AggchainFEP *AggchainFEPFilterer) WatchAggregationVkeyUpdated(opts *bind.WatchOpts, sink chan<- *AggchainFEPAggregationVkeyUpdated, oldAggregationVkey [][32]byte, newAggregationVkey [][32]byte) (event.Subscription, error) {

	var oldAggregationVkeyRule []interface{}
	for _, oldAggregationVkeyItem := range oldAggregationVkey {
		oldAggregationVkeyRule = append(oldAggregationVkeyRule, oldAggregationVkeyItem)
	}
	var newAggregationVkeyRule []interface{}
	for _, newAggregationVkeyItem := range newAggregationVkey {
		newAggregationVkeyRule = append(newAggregationVkeyRule, newAggregationVkeyItem)
	}

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "AggregationVkeyUpdated", oldAggregationVkeyRule, newAggregationVkeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPAggregationVkeyUpdated)
				if err := _AggchainFEP.contract.UnpackLog(event, "AggregationVkeyUpdated", log); err != nil {
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

// ParseAggregationVkeyUpdated is a log parse operation binding the contract event 0x390b73b2b067afcef04d30b573e4590c6e565519e370927dd777ca0ce8a55db0.
//
// Solidity: event AggregationVkeyUpdated(bytes32 indexed oldAggregationVkey, bytes32 indexed newAggregationVkey)
func (_AggchainFEP *AggchainFEPFilterer) ParseAggregationVkeyUpdated(log types.Log) (*AggchainFEPAggregationVkeyUpdated, error) {
	event := new(AggchainFEPAggregationVkeyUpdated)
	if err := _AggchainFEP.contract.UnpackLog(event, "AggregationVkeyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPDisableOptimisticModeIterator is returned from FilterDisableOptimisticMode and is used to iterate over the raw logs and unpacked data for DisableOptimisticMode events raised by the AggchainFEP contract.
type AggchainFEPDisableOptimisticModeIterator struct {
	Event *AggchainFEPDisableOptimisticMode // Event containing the contract specifics and raw log

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
func (it *AggchainFEPDisableOptimisticModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPDisableOptimisticMode)
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
		it.Event = new(AggchainFEPDisableOptimisticMode)
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
func (it *AggchainFEPDisableOptimisticModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPDisableOptimisticModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPDisableOptimisticMode represents a DisableOptimisticMode event raised by the AggchainFEP contract.
type AggchainFEPDisableOptimisticMode struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableOptimisticMode is a free log retrieval operation binding the contract event 0x334fa04f09bf04163481cd42794a867682f0b5ccb521db4fc4dbcca8a1e755ac.
//
// Solidity: event DisableOptimisticMode()
func (_AggchainFEP *AggchainFEPFilterer) FilterDisableOptimisticMode(opts *bind.FilterOpts) (*AggchainFEPDisableOptimisticModeIterator, error) {

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "DisableOptimisticMode")
	if err != nil {
		return nil, err
	}
	return &AggchainFEPDisableOptimisticModeIterator{contract: _AggchainFEP.contract, event: "DisableOptimisticMode", logs: logs, sub: sub}, nil
}

// WatchDisableOptimisticMode is a free log subscription operation binding the contract event 0x334fa04f09bf04163481cd42794a867682f0b5ccb521db4fc4dbcca8a1e755ac.
//
// Solidity: event DisableOptimisticMode()
func (_AggchainFEP *AggchainFEPFilterer) WatchDisableOptimisticMode(opts *bind.WatchOpts, sink chan<- *AggchainFEPDisableOptimisticMode) (event.Subscription, error) {

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "DisableOptimisticMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPDisableOptimisticMode)
				if err := _AggchainFEP.contract.UnpackLog(event, "DisableOptimisticMode", log); err != nil {
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

// ParseDisableOptimisticMode is a log parse operation binding the contract event 0x334fa04f09bf04163481cd42794a867682f0b5ccb521db4fc4dbcca8a1e755ac.
//
// Solidity: event DisableOptimisticMode()
func (_AggchainFEP *AggchainFEPFilterer) ParseDisableOptimisticMode(log types.Log) (*AggchainFEPDisableOptimisticMode, error) {
	event := new(AggchainFEPDisableOptimisticMode)
	if err := _AggchainFEP.contract.UnpackLog(event, "DisableOptimisticMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPDisableUseDefaultGatewayFlagIterator is returned from FilterDisableUseDefaultGatewayFlag and is used to iterate over the raw logs and unpacked data for DisableUseDefaultGatewayFlag events raised by the AggchainFEP contract.
type AggchainFEPDisableUseDefaultGatewayFlagIterator struct {
	Event *AggchainFEPDisableUseDefaultGatewayFlag // Event containing the contract specifics and raw log

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
func (it *AggchainFEPDisableUseDefaultGatewayFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPDisableUseDefaultGatewayFlag)
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
		it.Event = new(AggchainFEPDisableUseDefaultGatewayFlag)
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
func (it *AggchainFEPDisableUseDefaultGatewayFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPDisableUseDefaultGatewayFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPDisableUseDefaultGatewayFlag represents a DisableUseDefaultGatewayFlag event raised by the AggchainFEP contract.
type AggchainFEPDisableUseDefaultGatewayFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableUseDefaultGatewayFlag is a free log retrieval operation binding the contract event 0x67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e.
//
// Solidity: event DisableUseDefaultGatewayFlag()
func (_AggchainFEP *AggchainFEPFilterer) FilterDisableUseDefaultGatewayFlag(opts *bind.FilterOpts) (*AggchainFEPDisableUseDefaultGatewayFlagIterator, error) {

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "DisableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainFEPDisableUseDefaultGatewayFlagIterator{contract: _AggchainFEP.contract, event: "DisableUseDefaultGatewayFlag", logs: logs, sub: sub}, nil
}

// WatchDisableUseDefaultGatewayFlag is a free log subscription operation binding the contract event 0x67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e.
//
// Solidity: event DisableUseDefaultGatewayFlag()
func (_AggchainFEP *AggchainFEPFilterer) WatchDisableUseDefaultGatewayFlag(opts *bind.WatchOpts, sink chan<- *AggchainFEPDisableUseDefaultGatewayFlag) (event.Subscription, error) {

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "DisableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPDisableUseDefaultGatewayFlag)
				if err := _AggchainFEP.contract.UnpackLog(event, "DisableUseDefaultGatewayFlag", log); err != nil {
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

// ParseDisableUseDefaultGatewayFlag is a log parse operation binding the contract event 0x67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e.
//
// Solidity: event DisableUseDefaultGatewayFlag()
func (_AggchainFEP *AggchainFEPFilterer) ParseDisableUseDefaultGatewayFlag(log types.Log) (*AggchainFEPDisableUseDefaultGatewayFlag, error) {
	event := new(AggchainFEPDisableUseDefaultGatewayFlag)
	if err := _AggchainFEP.contract.UnpackLog(event, "DisableUseDefaultGatewayFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPEnableOptimisticModeIterator is returned from FilterEnableOptimisticMode and is used to iterate over the raw logs and unpacked data for EnableOptimisticMode events raised by the AggchainFEP contract.
type AggchainFEPEnableOptimisticModeIterator struct {
	Event *AggchainFEPEnableOptimisticMode // Event containing the contract specifics and raw log

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
func (it *AggchainFEPEnableOptimisticModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPEnableOptimisticMode)
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
		it.Event = new(AggchainFEPEnableOptimisticMode)
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
func (it *AggchainFEPEnableOptimisticModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPEnableOptimisticModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPEnableOptimisticMode represents a EnableOptimisticMode event raised by the AggchainFEP contract.
type AggchainFEPEnableOptimisticMode struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableOptimisticMode is a free log retrieval operation binding the contract event 0x26cf5e39429c85f7657b1e1f24aa2eb5a5882942a3f4a0dcd42844579bf7850a.
//
// Solidity: event EnableOptimisticMode()
func (_AggchainFEP *AggchainFEPFilterer) FilterEnableOptimisticMode(opts *bind.FilterOpts) (*AggchainFEPEnableOptimisticModeIterator, error) {

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "EnableOptimisticMode")
	if err != nil {
		return nil, err
	}
	return &AggchainFEPEnableOptimisticModeIterator{contract: _AggchainFEP.contract, event: "EnableOptimisticMode", logs: logs, sub: sub}, nil
}

// WatchEnableOptimisticMode is a free log subscription operation binding the contract event 0x26cf5e39429c85f7657b1e1f24aa2eb5a5882942a3f4a0dcd42844579bf7850a.
//
// Solidity: event EnableOptimisticMode()
func (_AggchainFEP *AggchainFEPFilterer) WatchEnableOptimisticMode(opts *bind.WatchOpts, sink chan<- *AggchainFEPEnableOptimisticMode) (event.Subscription, error) {

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "EnableOptimisticMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPEnableOptimisticMode)
				if err := _AggchainFEP.contract.UnpackLog(event, "EnableOptimisticMode", log); err != nil {
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

// ParseEnableOptimisticMode is a log parse operation binding the contract event 0x26cf5e39429c85f7657b1e1f24aa2eb5a5882942a3f4a0dcd42844579bf7850a.
//
// Solidity: event EnableOptimisticMode()
func (_AggchainFEP *AggchainFEPFilterer) ParseEnableOptimisticMode(log types.Log) (*AggchainFEPEnableOptimisticMode, error) {
	event := new(AggchainFEPEnableOptimisticMode)
	if err := _AggchainFEP.contract.UnpackLog(event, "EnableOptimisticMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPEnableUseDefaultGatewayFlagIterator is returned from FilterEnableUseDefaultGatewayFlag and is used to iterate over the raw logs and unpacked data for EnableUseDefaultGatewayFlag events raised by the AggchainFEP contract.
type AggchainFEPEnableUseDefaultGatewayFlagIterator struct {
	Event *AggchainFEPEnableUseDefaultGatewayFlag // Event containing the contract specifics and raw log

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
func (it *AggchainFEPEnableUseDefaultGatewayFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPEnableUseDefaultGatewayFlag)
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
		it.Event = new(AggchainFEPEnableUseDefaultGatewayFlag)
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
func (it *AggchainFEPEnableUseDefaultGatewayFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPEnableUseDefaultGatewayFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPEnableUseDefaultGatewayFlag represents a EnableUseDefaultGatewayFlag event raised by the AggchainFEP contract.
type AggchainFEPEnableUseDefaultGatewayFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableUseDefaultGatewayFlag is a free log retrieval operation binding the contract event 0xb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84.
//
// Solidity: event EnableUseDefaultGatewayFlag()
func (_AggchainFEP *AggchainFEPFilterer) FilterEnableUseDefaultGatewayFlag(opts *bind.FilterOpts) (*AggchainFEPEnableUseDefaultGatewayFlagIterator, error) {

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "EnableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainFEPEnableUseDefaultGatewayFlagIterator{contract: _AggchainFEP.contract, event: "EnableUseDefaultGatewayFlag", logs: logs, sub: sub}, nil
}

// WatchEnableUseDefaultGatewayFlag is a free log subscription operation binding the contract event 0xb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84.
//
// Solidity: event EnableUseDefaultGatewayFlag()
func (_AggchainFEP *AggchainFEPFilterer) WatchEnableUseDefaultGatewayFlag(opts *bind.WatchOpts, sink chan<- *AggchainFEPEnableUseDefaultGatewayFlag) (event.Subscription, error) {

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "EnableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPEnableUseDefaultGatewayFlag)
				if err := _AggchainFEP.contract.UnpackLog(event, "EnableUseDefaultGatewayFlag", log); err != nil {
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

// ParseEnableUseDefaultGatewayFlag is a log parse operation binding the contract event 0xb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84.
//
// Solidity: event EnableUseDefaultGatewayFlag()
func (_AggchainFEP *AggchainFEPFilterer) ParseEnableUseDefaultGatewayFlag(log types.Log) (*AggchainFEPEnableUseDefaultGatewayFlag, error) {
	event := new(AggchainFEPEnableUseDefaultGatewayFlag)
	if err := _AggchainFEP.contract.UnpackLog(event, "EnableUseDefaultGatewayFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AggchainFEP contract.
type AggchainFEPInitializedIterator struct {
	Event *AggchainFEPInitialized // Event containing the contract specifics and raw log

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
func (it *AggchainFEPInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPInitialized)
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
		it.Event = new(AggchainFEPInitialized)
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
func (it *AggchainFEPInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPInitialized represents a Initialized event raised by the AggchainFEP contract.
type AggchainFEPInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AggchainFEP *AggchainFEPFilterer) FilterInitialized(opts *bind.FilterOpts) (*AggchainFEPInitializedIterator, error) {

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AggchainFEPInitializedIterator{contract: _AggchainFEP.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AggchainFEP *AggchainFEPFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AggchainFEPInitialized) (event.Subscription, error) {

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPInitialized)
				if err := _AggchainFEP.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AggchainFEP *AggchainFEPFilterer) ParseInitialized(log types.Log) (*AggchainFEPInitialized, error) {
	event := new(AggchainFEPInitialized)
	if err := _AggchainFEP.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPOutputProposedIterator is returned from FilterOutputProposed and is used to iterate over the raw logs and unpacked data for OutputProposed events raised by the AggchainFEP contract.
type AggchainFEPOutputProposedIterator struct {
	Event *AggchainFEPOutputProposed // Event containing the contract specifics and raw log

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
func (it *AggchainFEPOutputProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPOutputProposed)
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
		it.Event = new(AggchainFEPOutputProposed)
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
func (it *AggchainFEPOutputProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPOutputProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPOutputProposed represents a OutputProposed event raised by the AggchainFEP contract.
type AggchainFEPOutputProposed struct {
	OutputRoot    [32]byte
	L2OutputIndex *big.Int
	L2BlockNumber *big.Int
	L1Timestamp   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOutputProposed is a free log retrieval operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_AggchainFEP *AggchainFEPFilterer) FilterOutputProposed(opts *bind.FilterOpts, outputRoot [][32]byte, l2OutputIndex []*big.Int, l2BlockNumber []*big.Int) (*AggchainFEPOutputProposedIterator, error) {

	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "OutputProposed", outputRootRule, l2OutputIndexRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &AggchainFEPOutputProposedIterator{contract: _AggchainFEP.contract, event: "OutputProposed", logs: logs, sub: sub}, nil
}

// WatchOutputProposed is a free log subscription operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_AggchainFEP *AggchainFEPFilterer) WatchOutputProposed(opts *bind.WatchOpts, sink chan<- *AggchainFEPOutputProposed, outputRoot [][32]byte, l2OutputIndex []*big.Int, l2BlockNumber []*big.Int) (event.Subscription, error) {

	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "OutputProposed", outputRootRule, l2OutputIndexRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPOutputProposed)
				if err := _AggchainFEP.contract.UnpackLog(event, "OutputProposed", log); err != nil {
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

// ParseOutputProposed is a log parse operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_AggchainFEP *AggchainFEPFilterer) ParseOutputProposed(log types.Log) (*AggchainFEPOutputProposed, error) {
	event := new(AggchainFEPOutputProposed)
	if err := _AggchainFEP.contract.UnpackLog(event, "OutputProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPRangeVkeyCommitmentUpdatedIterator is returned from FilterRangeVkeyCommitmentUpdated and is used to iterate over the raw logs and unpacked data for RangeVkeyCommitmentUpdated events raised by the AggchainFEP contract.
type AggchainFEPRangeVkeyCommitmentUpdatedIterator struct {
	Event *AggchainFEPRangeVkeyCommitmentUpdated // Event containing the contract specifics and raw log

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
func (it *AggchainFEPRangeVkeyCommitmentUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPRangeVkeyCommitmentUpdated)
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
		it.Event = new(AggchainFEPRangeVkeyCommitmentUpdated)
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
func (it *AggchainFEPRangeVkeyCommitmentUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPRangeVkeyCommitmentUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPRangeVkeyCommitmentUpdated represents a RangeVkeyCommitmentUpdated event raised by the AggchainFEP contract.
type AggchainFEPRangeVkeyCommitmentUpdated struct {
	OldRangeVkeyCommitment [32]byte
	NewRangeVkeyCommitment [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRangeVkeyCommitmentUpdated is a free log retrieval operation binding the contract event 0xbf8cab6317796bfa97fea82b6d27c9542a08fa0821813cf2a57e7cff7fdc8156.
//
// Solidity: event RangeVkeyCommitmentUpdated(bytes32 indexed oldRangeVkeyCommitment, bytes32 indexed newRangeVkeyCommitment)
func (_AggchainFEP *AggchainFEPFilterer) FilterRangeVkeyCommitmentUpdated(opts *bind.FilterOpts, oldRangeVkeyCommitment [][32]byte, newRangeVkeyCommitment [][32]byte) (*AggchainFEPRangeVkeyCommitmentUpdatedIterator, error) {

	var oldRangeVkeyCommitmentRule []interface{}
	for _, oldRangeVkeyCommitmentItem := range oldRangeVkeyCommitment {
		oldRangeVkeyCommitmentRule = append(oldRangeVkeyCommitmentRule, oldRangeVkeyCommitmentItem)
	}
	var newRangeVkeyCommitmentRule []interface{}
	for _, newRangeVkeyCommitmentItem := range newRangeVkeyCommitment {
		newRangeVkeyCommitmentRule = append(newRangeVkeyCommitmentRule, newRangeVkeyCommitmentItem)
	}

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "RangeVkeyCommitmentUpdated", oldRangeVkeyCommitmentRule, newRangeVkeyCommitmentRule)
	if err != nil {
		return nil, err
	}
	return &AggchainFEPRangeVkeyCommitmentUpdatedIterator{contract: _AggchainFEP.contract, event: "RangeVkeyCommitmentUpdated", logs: logs, sub: sub}, nil
}

// WatchRangeVkeyCommitmentUpdated is a free log subscription operation binding the contract event 0xbf8cab6317796bfa97fea82b6d27c9542a08fa0821813cf2a57e7cff7fdc8156.
//
// Solidity: event RangeVkeyCommitmentUpdated(bytes32 indexed oldRangeVkeyCommitment, bytes32 indexed newRangeVkeyCommitment)
func (_AggchainFEP *AggchainFEPFilterer) WatchRangeVkeyCommitmentUpdated(opts *bind.WatchOpts, sink chan<- *AggchainFEPRangeVkeyCommitmentUpdated, oldRangeVkeyCommitment [][32]byte, newRangeVkeyCommitment [][32]byte) (event.Subscription, error) {

	var oldRangeVkeyCommitmentRule []interface{}
	for _, oldRangeVkeyCommitmentItem := range oldRangeVkeyCommitment {
		oldRangeVkeyCommitmentRule = append(oldRangeVkeyCommitmentRule, oldRangeVkeyCommitmentItem)
	}
	var newRangeVkeyCommitmentRule []interface{}
	for _, newRangeVkeyCommitmentItem := range newRangeVkeyCommitment {
		newRangeVkeyCommitmentRule = append(newRangeVkeyCommitmentRule, newRangeVkeyCommitmentItem)
	}

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "RangeVkeyCommitmentUpdated", oldRangeVkeyCommitmentRule, newRangeVkeyCommitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPRangeVkeyCommitmentUpdated)
				if err := _AggchainFEP.contract.UnpackLog(event, "RangeVkeyCommitmentUpdated", log); err != nil {
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

// ParseRangeVkeyCommitmentUpdated is a log parse operation binding the contract event 0xbf8cab6317796bfa97fea82b6d27c9542a08fa0821813cf2a57e7cff7fdc8156.
//
// Solidity: event RangeVkeyCommitmentUpdated(bytes32 indexed oldRangeVkeyCommitment, bytes32 indexed newRangeVkeyCommitment)
func (_AggchainFEP *AggchainFEPFilterer) ParseRangeVkeyCommitmentUpdated(log types.Log) (*AggchainFEPRangeVkeyCommitmentUpdated, error) {
	event := new(AggchainFEPRangeVkeyCommitmentUpdated)
	if err := _AggchainFEP.contract.UnpackLog(event, "RangeVkeyCommitmentUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPRollupConfigHashUpdatedIterator is returned from FilterRollupConfigHashUpdated and is used to iterate over the raw logs and unpacked data for RollupConfigHashUpdated events raised by the AggchainFEP contract.
type AggchainFEPRollupConfigHashUpdatedIterator struct {
	Event *AggchainFEPRollupConfigHashUpdated // Event containing the contract specifics and raw log

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
func (it *AggchainFEPRollupConfigHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPRollupConfigHashUpdated)
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
		it.Event = new(AggchainFEPRollupConfigHashUpdated)
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
func (it *AggchainFEPRollupConfigHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPRollupConfigHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPRollupConfigHashUpdated represents a RollupConfigHashUpdated event raised by the AggchainFEP contract.
type AggchainFEPRollupConfigHashUpdated struct {
	OldRollupConfigHash [32]byte
	NewRollupConfigHash [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRollupConfigHashUpdated is a free log retrieval operation binding the contract event 0x5d9ebe9f09b0810b3546b30781ba9a51092b37dd6abada4b830ce54a41ac6a4b.
//
// Solidity: event RollupConfigHashUpdated(bytes32 indexed oldRollupConfigHash, bytes32 indexed newRollupConfigHash)
func (_AggchainFEP *AggchainFEPFilterer) FilterRollupConfigHashUpdated(opts *bind.FilterOpts, oldRollupConfigHash [][32]byte, newRollupConfigHash [][32]byte) (*AggchainFEPRollupConfigHashUpdatedIterator, error) {

	var oldRollupConfigHashRule []interface{}
	for _, oldRollupConfigHashItem := range oldRollupConfigHash {
		oldRollupConfigHashRule = append(oldRollupConfigHashRule, oldRollupConfigHashItem)
	}
	var newRollupConfigHashRule []interface{}
	for _, newRollupConfigHashItem := range newRollupConfigHash {
		newRollupConfigHashRule = append(newRollupConfigHashRule, newRollupConfigHashItem)
	}

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "RollupConfigHashUpdated", oldRollupConfigHashRule, newRollupConfigHashRule)
	if err != nil {
		return nil, err
	}
	return &AggchainFEPRollupConfigHashUpdatedIterator{contract: _AggchainFEP.contract, event: "RollupConfigHashUpdated", logs: logs, sub: sub}, nil
}

// WatchRollupConfigHashUpdated is a free log subscription operation binding the contract event 0x5d9ebe9f09b0810b3546b30781ba9a51092b37dd6abada4b830ce54a41ac6a4b.
//
// Solidity: event RollupConfigHashUpdated(bytes32 indexed oldRollupConfigHash, bytes32 indexed newRollupConfigHash)
func (_AggchainFEP *AggchainFEPFilterer) WatchRollupConfigHashUpdated(opts *bind.WatchOpts, sink chan<- *AggchainFEPRollupConfigHashUpdated, oldRollupConfigHash [][32]byte, newRollupConfigHash [][32]byte) (event.Subscription, error) {

	var oldRollupConfigHashRule []interface{}
	for _, oldRollupConfigHashItem := range oldRollupConfigHash {
		oldRollupConfigHashRule = append(oldRollupConfigHashRule, oldRollupConfigHashItem)
	}
	var newRollupConfigHashRule []interface{}
	for _, newRollupConfigHashItem := range newRollupConfigHash {
		newRollupConfigHashRule = append(newRollupConfigHashRule, newRollupConfigHashItem)
	}

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "RollupConfigHashUpdated", oldRollupConfigHashRule, newRollupConfigHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPRollupConfigHashUpdated)
				if err := _AggchainFEP.contract.UnpackLog(event, "RollupConfigHashUpdated", log); err != nil {
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

// ParseRollupConfigHashUpdated is a log parse operation binding the contract event 0x5d9ebe9f09b0810b3546b30781ba9a51092b37dd6abada4b830ce54a41ac6a4b.
//
// Solidity: event RollupConfigHashUpdated(bytes32 indexed oldRollupConfigHash, bytes32 indexed newRollupConfigHash)
func (_AggchainFEP *AggchainFEPFilterer) ParseRollupConfigHashUpdated(log types.Log) (*AggchainFEPRollupConfigHashUpdated, error) {
	event := new(AggchainFEPRollupConfigHashUpdated)
	if err := _AggchainFEP.contract.UnpackLog(event, "RollupConfigHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the AggchainFEP contract.
type AggchainFEPSetTrustedSequencerIterator struct {
	Event *AggchainFEPSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *AggchainFEPSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPSetTrustedSequencer)
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
		it.Event = new(AggchainFEPSetTrustedSequencer)
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
func (it *AggchainFEPSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPSetTrustedSequencer represents a SetTrustedSequencer event raised by the AggchainFEP contract.
type AggchainFEPSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_AggchainFEP *AggchainFEPFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*AggchainFEPSetTrustedSequencerIterator, error) {

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &AggchainFEPSetTrustedSequencerIterator{contract: _AggchainFEP.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_AggchainFEP *AggchainFEPFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *AggchainFEPSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPSetTrustedSequencer)
				if err := _AggchainFEP.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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

// ParseSetTrustedSequencer is a log parse operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_AggchainFEP *AggchainFEPFilterer) ParseSetTrustedSequencer(log types.Log) (*AggchainFEPSetTrustedSequencer, error) {
	event := new(AggchainFEPSetTrustedSequencer)
	if err := _AggchainFEP.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the AggchainFEP contract.
type AggchainFEPSetTrustedSequencerURLIterator struct {
	Event *AggchainFEPSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *AggchainFEPSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPSetTrustedSequencerURL)
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
		it.Event = new(AggchainFEPSetTrustedSequencerURL)
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
func (it *AggchainFEPSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the AggchainFEP contract.
type AggchainFEPSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_AggchainFEP *AggchainFEPFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*AggchainFEPSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &AggchainFEPSetTrustedSequencerURLIterator{contract: _AggchainFEP.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_AggchainFEP *AggchainFEPFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *AggchainFEPSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPSetTrustedSequencerURL)
				if err := _AggchainFEP.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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

// ParseSetTrustedSequencerURL is a log parse operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_AggchainFEP *AggchainFEPFilterer) ParseSetTrustedSequencerURL(log types.Log) (*AggchainFEPSetTrustedSequencerURL, error) {
	event := new(AggchainFEPSetTrustedSequencerURL)
	if err := _AggchainFEP.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPSubmissionIntervalUpdatedIterator is returned from FilterSubmissionIntervalUpdated and is used to iterate over the raw logs and unpacked data for SubmissionIntervalUpdated events raised by the AggchainFEP contract.
type AggchainFEPSubmissionIntervalUpdatedIterator struct {
	Event *AggchainFEPSubmissionIntervalUpdated // Event containing the contract specifics and raw log

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
func (it *AggchainFEPSubmissionIntervalUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPSubmissionIntervalUpdated)
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
		it.Event = new(AggchainFEPSubmissionIntervalUpdated)
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
func (it *AggchainFEPSubmissionIntervalUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPSubmissionIntervalUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPSubmissionIntervalUpdated represents a SubmissionIntervalUpdated event raised by the AggchainFEP contract.
type AggchainFEPSubmissionIntervalUpdated struct {
	OldSubmissionInterval *big.Int
	NewSubmissionInterval *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSubmissionIntervalUpdated is a free log retrieval operation binding the contract event 0xc1bf9abfb57ea01ed9ecb4f45e9cefa7ba44b2e6778c3ce7281409999f1af1b2.
//
// Solidity: event SubmissionIntervalUpdated(uint256 oldSubmissionInterval, uint256 newSubmissionInterval)
func (_AggchainFEP *AggchainFEPFilterer) FilterSubmissionIntervalUpdated(opts *bind.FilterOpts) (*AggchainFEPSubmissionIntervalUpdatedIterator, error) {

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "SubmissionIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return &AggchainFEPSubmissionIntervalUpdatedIterator{contract: _AggchainFEP.contract, event: "SubmissionIntervalUpdated", logs: logs, sub: sub}, nil
}

// WatchSubmissionIntervalUpdated is a free log subscription operation binding the contract event 0xc1bf9abfb57ea01ed9ecb4f45e9cefa7ba44b2e6778c3ce7281409999f1af1b2.
//
// Solidity: event SubmissionIntervalUpdated(uint256 oldSubmissionInterval, uint256 newSubmissionInterval)
func (_AggchainFEP *AggchainFEPFilterer) WatchSubmissionIntervalUpdated(opts *bind.WatchOpts, sink chan<- *AggchainFEPSubmissionIntervalUpdated) (event.Subscription, error) {

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "SubmissionIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPSubmissionIntervalUpdated)
				if err := _AggchainFEP.contract.UnpackLog(event, "SubmissionIntervalUpdated", log); err != nil {
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

// ParseSubmissionIntervalUpdated is a log parse operation binding the contract event 0xc1bf9abfb57ea01ed9ecb4f45e9cefa7ba44b2e6778c3ce7281409999f1af1b2.
//
// Solidity: event SubmissionIntervalUpdated(uint256 oldSubmissionInterval, uint256 newSubmissionInterval)
func (_AggchainFEP *AggchainFEPFilterer) ParseSubmissionIntervalUpdated(log types.Log) (*AggchainFEPSubmissionIntervalUpdated, error) {
	event := new(AggchainFEPSubmissionIntervalUpdated)
	if err := _AggchainFEP.contract.UnpackLog(event, "SubmissionIntervalUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the AggchainFEP contract.
type AggchainFEPTransferAdminRoleIterator struct {
	Event *AggchainFEPTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *AggchainFEPTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPTransferAdminRole)
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
		it.Event = new(AggchainFEPTransferAdminRole)
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
func (it *AggchainFEPTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPTransferAdminRole represents a TransferAdminRole event raised by the AggchainFEP contract.
type AggchainFEPTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_AggchainFEP *AggchainFEPFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*AggchainFEPTransferAdminRoleIterator, error) {

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &AggchainFEPTransferAdminRoleIterator{contract: _AggchainFEP.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_AggchainFEP *AggchainFEPFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *AggchainFEPTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPTransferAdminRole)
				if err := _AggchainFEP.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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

// ParseTransferAdminRole is a log parse operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_AggchainFEP *AggchainFEPFilterer) ParseTransferAdminRole(log types.Log) (*AggchainFEPTransferAdminRole, error) {
	event := new(AggchainFEPTransferAdminRole)
	if err := _AggchainFEP.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPTransferAggchainManagerRoleIterator is returned from FilterTransferAggchainManagerRole and is used to iterate over the raw logs and unpacked data for TransferAggchainManagerRole events raised by the AggchainFEP contract.
type AggchainFEPTransferAggchainManagerRoleIterator struct {
	Event *AggchainFEPTransferAggchainManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainFEPTransferAggchainManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPTransferAggchainManagerRole)
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
		it.Event = new(AggchainFEPTransferAggchainManagerRole)
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
func (it *AggchainFEPTransferAggchainManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPTransferAggchainManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPTransferAggchainManagerRole represents a TransferAggchainManagerRole event raised by the AggchainFEP contract.
type AggchainFEPTransferAggchainManagerRole struct {
	CurrentAggchainManager    common.Address
	NewPendingAggchainManager common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterTransferAggchainManagerRole is a free log retrieval operation binding the contract event 0xa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab6.
//
// Solidity: event TransferAggchainManagerRole(address currentAggchainManager, address newPendingAggchainManager)
func (_AggchainFEP *AggchainFEPFilterer) FilterTransferAggchainManagerRole(opts *bind.FilterOpts) (*AggchainFEPTransferAggchainManagerRoleIterator, error) {

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "TransferAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainFEPTransferAggchainManagerRoleIterator{contract: _AggchainFEP.contract, event: "TransferAggchainManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferAggchainManagerRole is a free log subscription operation binding the contract event 0xa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab6.
//
// Solidity: event TransferAggchainManagerRole(address currentAggchainManager, address newPendingAggchainManager)
func (_AggchainFEP *AggchainFEPFilterer) WatchTransferAggchainManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainFEPTransferAggchainManagerRole) (event.Subscription, error) {

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "TransferAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPTransferAggchainManagerRole)
				if err := _AggchainFEP.contract.UnpackLog(event, "TransferAggchainManagerRole", log); err != nil {
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

// ParseTransferAggchainManagerRole is a log parse operation binding the contract event 0xa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab6.
//
// Solidity: event TransferAggchainManagerRole(address currentAggchainManager, address newPendingAggchainManager)
func (_AggchainFEP *AggchainFEPFilterer) ParseTransferAggchainManagerRole(log types.Log) (*AggchainFEPTransferAggchainManagerRole, error) {
	event := new(AggchainFEPTransferAggchainManagerRole)
	if err := _AggchainFEP.contract.UnpackLog(event, "TransferAggchainManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPTransferOptimisticModeManagerRoleIterator is returned from FilterTransferOptimisticModeManagerRole and is used to iterate over the raw logs and unpacked data for TransferOptimisticModeManagerRole events raised by the AggchainFEP contract.
type AggchainFEPTransferOptimisticModeManagerRoleIterator struct {
	Event *AggchainFEPTransferOptimisticModeManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainFEPTransferOptimisticModeManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPTransferOptimisticModeManagerRole)
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
		it.Event = new(AggchainFEPTransferOptimisticModeManagerRole)
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
func (it *AggchainFEPTransferOptimisticModeManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPTransferOptimisticModeManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPTransferOptimisticModeManagerRole represents a TransferOptimisticModeManagerRole event raised by the AggchainFEP contract.
type AggchainFEPTransferOptimisticModeManagerRole struct {
	CurrentOptimisticModeManager    common.Address
	NewPendingOptimisticModeManager common.Address
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterTransferOptimisticModeManagerRole is a free log retrieval operation binding the contract event 0xf67c2e74a956fb061c1a9c17172d5a9197efc33c180fac0319ce5cd90702af79.
//
// Solidity: event TransferOptimisticModeManagerRole(address currentOptimisticModeManager, address newPendingOptimisticModeManager)
func (_AggchainFEP *AggchainFEPFilterer) FilterTransferOptimisticModeManagerRole(opts *bind.FilterOpts) (*AggchainFEPTransferOptimisticModeManagerRoleIterator, error) {

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "TransferOptimisticModeManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainFEPTransferOptimisticModeManagerRoleIterator{contract: _AggchainFEP.contract, event: "TransferOptimisticModeManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferOptimisticModeManagerRole is a free log subscription operation binding the contract event 0xf67c2e74a956fb061c1a9c17172d5a9197efc33c180fac0319ce5cd90702af79.
//
// Solidity: event TransferOptimisticModeManagerRole(address currentOptimisticModeManager, address newPendingOptimisticModeManager)
func (_AggchainFEP *AggchainFEPFilterer) WatchTransferOptimisticModeManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainFEPTransferOptimisticModeManagerRole) (event.Subscription, error) {

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "TransferOptimisticModeManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPTransferOptimisticModeManagerRole)
				if err := _AggchainFEP.contract.UnpackLog(event, "TransferOptimisticModeManagerRole", log); err != nil {
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

// ParseTransferOptimisticModeManagerRole is a log parse operation binding the contract event 0xf67c2e74a956fb061c1a9c17172d5a9197efc33c180fac0319ce5cd90702af79.
//
// Solidity: event TransferOptimisticModeManagerRole(address currentOptimisticModeManager, address newPendingOptimisticModeManager)
func (_AggchainFEP *AggchainFEPFilterer) ParseTransferOptimisticModeManagerRole(log types.Log) (*AggchainFEPTransferOptimisticModeManagerRole, error) {
	event := new(AggchainFEPTransferOptimisticModeManagerRole)
	if err := _AggchainFEP.contract.UnpackLog(event, "TransferOptimisticModeManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPTransferVKeyManagerRoleIterator is returned from FilterTransferVKeyManagerRole and is used to iterate over the raw logs and unpacked data for TransferVKeyManagerRole events raised by the AggchainFEP contract.
type AggchainFEPTransferVKeyManagerRoleIterator struct {
	Event *AggchainFEPTransferVKeyManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainFEPTransferVKeyManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPTransferVKeyManagerRole)
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
		it.Event = new(AggchainFEPTransferVKeyManagerRole)
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
func (it *AggchainFEPTransferVKeyManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPTransferVKeyManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPTransferVKeyManagerRole represents a TransferVKeyManagerRole event raised by the AggchainFEP contract.
type AggchainFEPTransferVKeyManagerRole struct {
	CurrentVKeyManager    common.Address
	NewPendingVKeyManager common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTransferVKeyManagerRole is a free log retrieval operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_AggchainFEP *AggchainFEPFilterer) FilterTransferVKeyManagerRole(opts *bind.FilterOpts) (*AggchainFEPTransferVKeyManagerRoleIterator, error) {

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainFEPTransferVKeyManagerRoleIterator{contract: _AggchainFEP.contract, event: "TransferVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferVKeyManagerRole is a free log subscription operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_AggchainFEP *AggchainFEPFilterer) WatchTransferVKeyManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainFEPTransferVKeyManagerRole) (event.Subscription, error) {

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPTransferVKeyManagerRole)
				if err := _AggchainFEP.contract.UnpackLog(event, "TransferVKeyManagerRole", log); err != nil {
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

// ParseTransferVKeyManagerRole is a log parse operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_AggchainFEP *AggchainFEPFilterer) ParseTransferVKeyManagerRole(log types.Log) (*AggchainFEPTransferVKeyManagerRole, error) {
	event := new(AggchainFEPTransferVKeyManagerRole)
	if err := _AggchainFEP.contract.UnpackLog(event, "TransferVKeyManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainFEPUpdateAggchainVKeyIterator is returned from FilterUpdateAggchainVKey and is used to iterate over the raw logs and unpacked data for UpdateAggchainVKey events raised by the AggchainFEP contract.
type AggchainFEPUpdateAggchainVKeyIterator struct {
	Event *AggchainFEPUpdateAggchainVKey // Event containing the contract specifics and raw log

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
func (it *AggchainFEPUpdateAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainFEPUpdateAggchainVKey)
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
		it.Event = new(AggchainFEPUpdateAggchainVKey)
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
func (it *AggchainFEPUpdateAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainFEPUpdateAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainFEPUpdateAggchainVKey represents a UpdateAggchainVKey event raised by the AggchainFEP contract.
type AggchainFEPUpdateAggchainVKey struct {
	Selector             [4]byte
	PreviousAggchainVKey [32]byte
	NewAggchainVKey      [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateAggchainVKey is a free log retrieval operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_AggchainFEP *AggchainFEPFilterer) FilterUpdateAggchainVKey(opts *bind.FilterOpts) (*AggchainFEPUpdateAggchainVKeyIterator, error) {

	logs, sub, err := _AggchainFEP.contract.FilterLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AggchainFEPUpdateAggchainVKeyIterator{contract: _AggchainFEP.contract, event: "UpdateAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUpdateAggchainVKey is a free log subscription operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_AggchainFEP *AggchainFEPFilterer) WatchUpdateAggchainVKey(opts *bind.WatchOpts, sink chan<- *AggchainFEPUpdateAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _AggchainFEP.contract.WatchLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainFEPUpdateAggchainVKey)
				if err := _AggchainFEP.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
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

// ParseUpdateAggchainVKey is a log parse operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_AggchainFEP *AggchainFEPFilterer) ParseUpdateAggchainVKey(log types.Log) (*AggchainFEPUpdateAggchainVKey, error) {
	event := new(AggchainFEPUpdateAggchainVKey)
	if err := _AggchainFEP.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

