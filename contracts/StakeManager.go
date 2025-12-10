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

// StakeManagerMetaData contains all meta data concerning the StakeManager contract.
var StakeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousRootChain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRootChain\",\"type\":\"address\"}],\"name\":\"RootChainChanged\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CHECKPOINT_REWARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NFTContract\",\"outputs\":[{\"internalType\":\"contractStakingNFT\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NFTCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WITHDRAWAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"auctionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRootChain\",\"type\":\"address\"}],\"name\":\"changeRootChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkPointBlockInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockInterval\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"voteHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256[3][]\",\"name\":\"sigs\",\"type\":\"uint256[3][]\"}],\"name\":\"checkSignatures\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkpointRewardDelta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"accumFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"claimFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"heimdallFee\",\"type\":\"uint256\"}],\"name\":\"confirmAuctionBid\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"convertMaticToPOL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentValidatorSetSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentValidatorSetTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseValidatorDelegatedAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"delegatedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"delegationDeposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"delegationDepositPOL\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"delegatorsReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"auctionUser\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"heimdallFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auctionAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"acceptDelegation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"signerPubkey\",\"type\":\"bytes\"}],\"name\":\"dethroneAndStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dynasty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"eventsHub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"extensionCode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"forceUnstake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"forceUnstakePOL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"getValidatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getValidatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"contractIGovernance\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_NFTContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingLogger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorShareFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_extensionCode\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenNew\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_migration\",\"type\":\"address\"}],\"name\":\"initializePOL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"}],\"name\":\"insertSigners\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"latestSignerUpdateEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"logger\",\"outputs\":[{\"internalType\":\"contractStakingInfo\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxRewardedCheckpoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromValidatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"migrateDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorIdFrom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorIdTo\",\"type\":\"uint256\"}],\"name\":\"migrateValidatorsData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"migration\",\"outputs\":[{\"internalType\":\"contractIPolygonMigration\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minHeimdallFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"prevBlockInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proposerBonus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_NFTContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingLogger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorShareFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_extensionCode\",\"type\":\"address\"}],\"name\":\"reinitialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"replacementCoolDown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stakeRewards\",\"type\":\"bool\"}],\"name\":\"restake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stakeRewards\",\"type\":\"bool\"}],\"name\":\"restakePOL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardDecreasePerCheckpoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootChain\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_currentEpoch\",\"type\":\"uint256\"}],\"name\":\"setCurrentEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setDelegationEnabled\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"setStakingToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerToValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signerUpdateLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_slashingInfoList\",\"type\":\"bytes\"}],\"name\":\"slash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"heimdallFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"acceptDelegation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"signerPubkey\",\"type\":\"bytes\"}],\"name\":\"stakeFor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"heimdallFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"acceptDelegation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"signerPubkey\",\"type\":\"bytes\"}],\"name\":\"stakeForPOL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_acceptDelegation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_signerPubkey\",\"type\":\"bytes\"}],\"name\":\"startAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"forNCheckpoints\",\"type\":\"uint256\"}],\"name\":\"stopAuctions\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenMatic\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"heimdallFee\",\"type\":\"uint256\"}],\"name\":\"topUpForFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalHeimdallFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalRewardsLiquidated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"totalStakedFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"transferFunds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"transferFundsPOL\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"unjail\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"unstakeClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"unstakeClaimPOL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"unstakePOL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blocks\",\"type\":\"uint256\"}],\"name\":\"updateCheckPointBlockInterval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReward\",\"type\":\"uint256\"}],\"name\":\"updateCheckpointReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardDecreasePerCheckpoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxRewardedCheckpoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_checkpointRewardDelta\",\"type\":\"uint256\"}],\"name\":\"updateCheckpointRewardParams\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newCommissionRate\",\"type\":\"uint256\"}],\"name\":\"updateCommissionRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDynasty\",\"type\":\"uint256\"}],\"name\":\"updateDynastyValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minHeimdallFee\",\"type\":\"uint256\"}],\"name\":\"updateMinAmounts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newProposerBonus\",\"type\":\"uint256\"}],\"name\":\"updateProposerBonus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signerPubkey\",\"type\":\"bytes\"}],\"name\":\"updateSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"updateSignerUpdateLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newContractAddress\",\"type\":\"address\"}],\"name\":\"updateValidatorContractAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"delegation\",\"type\":\"bool\"}],\"name\":\"updateValidatorDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"updateValidatorState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"updateValidatorThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userFeeExit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"acceptDelegation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"signerPubkey\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"validatorReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorShareFactory\",\"outputs\":[{\"internalType\":\"contractValidatorShareFactory\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"validatorStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakerCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorStateChanges\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"stakerCount\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activationEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivationEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"enumStakeManagerStorage.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastCommissionUpdate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorsReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialRewardPerStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"withdrawDelegatorsReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"withdrawRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"withdrawRewardsPOL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawalDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeManagerMetaData.ABI instead.
var StakeManagerABI = StakeManagerMetaData.ABI

// StakeManager is an auto generated Go binding around an Ethereum contract.
type StakeManager struct {
	StakeManagerCaller     // Read-only binding to the contract
	StakeManagerTransactor // Write-only binding to the contract
	StakeManagerFilterer   // Log filterer for contract events
}

// StakeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeManagerSession struct {
	Contract     *StakeManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeManagerCallerSession struct {
	Contract *StakeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StakeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeManagerTransactorSession struct {
	Contract     *StakeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StakeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeManagerRaw struct {
	Contract *StakeManager // Generic contract binding to access the raw methods on
}

// StakeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeManagerCallerRaw struct {
	Contract *StakeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeManagerTransactorRaw struct {
	Contract *StakeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakeManager creates a new instance of StakeManager, bound to a specific deployed contract.
func NewStakeManager(address common.Address, backend bind.ContractBackend) (*StakeManager, error) {
	contract, err := bindStakeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakeManager{StakeManagerCaller: StakeManagerCaller{contract: contract}, StakeManagerTransactor: StakeManagerTransactor{contract: contract}, StakeManagerFilterer: StakeManagerFilterer{contract: contract}}, nil
}

// NewStakeManagerCaller creates a new read-only instance of StakeManager, bound to a specific deployed contract.
func NewStakeManagerCaller(address common.Address, caller bind.ContractCaller) (*StakeManagerCaller, error) {
	contract, err := bindStakeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeManagerCaller{contract: contract}, nil
}

// NewStakeManagerTransactor creates a new write-only instance of StakeManager, bound to a specific deployed contract.
func NewStakeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeManagerTransactor, error) {
	contract, err := bindStakeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeManagerTransactor{contract: contract}, nil
}

// NewStakeManagerFilterer creates a new log filterer instance of StakeManager, bound to a specific deployed contract.
func NewStakeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeManagerFilterer, error) {
	contract, err := bindStakeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeManagerFilterer{contract: contract}, nil
}

// bindStakeManager binds a generic wrapper to an already deployed contract.
func bindStakeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeManager *StakeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeManager.Contract.StakeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeManager *StakeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeManager *StakeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeManager *StakeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeManager *StakeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeManager *StakeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeManager.Contract.contract.Transact(opts, method, params...)
}

// CHECKPOINTREWARD is a free data retrieval call binding the contract method 0x7d669752.
//
// Solidity: function CHECKPOINT_REWARD() view returns(uint256)
func (_StakeManager *StakeManagerCaller) CHECKPOINTREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "CHECKPOINT_REWARD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHECKPOINTREWARD is a free data retrieval call binding the contract method 0x7d669752.
//
// Solidity: function CHECKPOINT_REWARD() view returns(uint256)
func (_StakeManager *StakeManagerSession) CHECKPOINTREWARD() (*big.Int, error) {
	return _StakeManager.Contract.CHECKPOINTREWARD(&_StakeManager.CallOpts)
}

// CHECKPOINTREWARD is a free data retrieval call binding the contract method 0x7d669752.
//
// Solidity: function CHECKPOINT_REWARD() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) CHECKPOINTREWARD() (*big.Int, error) {
	return _StakeManager.Contract.CHECKPOINTREWARD(&_StakeManager.CallOpts)
}

// NFTContract is a free data retrieval call binding the contract method 0x31c2273b.
//
// Solidity: function NFTContract() view returns(address)
func (_StakeManager *StakeManagerCaller) NFTContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "NFTContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NFTContract is a free data retrieval call binding the contract method 0x31c2273b.
//
// Solidity: function NFTContract() view returns(address)
func (_StakeManager *StakeManagerSession) NFTContract() (common.Address, error) {
	return _StakeManager.Contract.NFTContract(&_StakeManager.CallOpts)
}

// NFTContract is a free data retrieval call binding the contract method 0x31c2273b.
//
// Solidity: function NFTContract() view returns(address)
func (_StakeManager *StakeManagerCallerSession) NFTContract() (common.Address, error) {
	return _StakeManager.Contract.NFTContract(&_StakeManager.CallOpts)
}

// NFTCounter is a free data retrieval call binding the contract method 0x5508d8e1.
//
// Solidity: function NFTCounter() view returns(uint256)
func (_StakeManager *StakeManagerCaller) NFTCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "NFTCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NFTCounter is a free data retrieval call binding the contract method 0x5508d8e1.
//
// Solidity: function NFTCounter() view returns(uint256)
func (_StakeManager *StakeManagerSession) NFTCounter() (*big.Int, error) {
	return _StakeManager.Contract.NFTCounter(&_StakeManager.CallOpts)
}

// NFTCounter is a free data retrieval call binding the contract method 0x5508d8e1.
//
// Solidity: function NFTCounter() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) NFTCounter() (*big.Int, error) {
	return _StakeManager.Contract.NFTCounter(&_StakeManager.CallOpts)
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() view returns(uint256)
func (_StakeManager *StakeManagerCaller) WITHDRAWALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "WITHDRAWAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() view returns(uint256)
func (_StakeManager *StakeManagerSession) WITHDRAWALDELAY() (*big.Int, error) {
	return _StakeManager.Contract.WITHDRAWALDELAY(&_StakeManager.CallOpts)
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) WITHDRAWALDELAY() (*big.Int, error) {
	return _StakeManager.Contract.WITHDRAWALDELAY(&_StakeManager.CallOpts)
}

// AccountStateRoot is a free data retrieval call binding the contract method 0x17c2b910.
//
// Solidity: function accountStateRoot() view returns(bytes32)
func (_StakeManager *StakeManagerCaller) AccountStateRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "accountStateRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AccountStateRoot is a free data retrieval call binding the contract method 0x17c2b910.
//
// Solidity: function accountStateRoot() view returns(bytes32)
func (_StakeManager *StakeManagerSession) AccountStateRoot() ([32]byte, error) {
	return _StakeManager.Contract.AccountStateRoot(&_StakeManager.CallOpts)
}

// AccountStateRoot is a free data retrieval call binding the contract method 0x17c2b910.
//
// Solidity: function accountStateRoot() view returns(bytes32)
func (_StakeManager *StakeManagerCallerSession) AccountStateRoot() ([32]byte, error) {
	return _StakeManager.Contract.AccountStateRoot(&_StakeManager.CallOpts)
}

// AuctionPeriod is a free data retrieval call binding the contract method 0x0cccfc58.
//
// Solidity: function auctionPeriod() view returns(uint256)
func (_StakeManager *StakeManagerCaller) AuctionPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "auctionPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionPeriod is a free data retrieval call binding the contract method 0x0cccfc58.
//
// Solidity: function auctionPeriod() view returns(uint256)
func (_StakeManager *StakeManagerSession) AuctionPeriod() (*big.Int, error) {
	return _StakeManager.Contract.AuctionPeriod(&_StakeManager.CallOpts)
}

// AuctionPeriod is a free data retrieval call binding the contract method 0x0cccfc58.
//
// Solidity: function auctionPeriod() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) AuctionPeriod() (*big.Int, error) {
	return _StakeManager.Contract.AuctionPeriod(&_StakeManager.CallOpts)
}

// CheckPointBlockInterval is a free data retrieval call binding the contract method 0x25316411.
//
// Solidity: function checkPointBlockInterval() view returns(uint256)
func (_StakeManager *StakeManagerCaller) CheckPointBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "checkPointBlockInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckPointBlockInterval is a free data retrieval call binding the contract method 0x25316411.
//
// Solidity: function checkPointBlockInterval() view returns(uint256)
func (_StakeManager *StakeManagerSession) CheckPointBlockInterval() (*big.Int, error) {
	return _StakeManager.Contract.CheckPointBlockInterval(&_StakeManager.CallOpts)
}

// CheckPointBlockInterval is a free data retrieval call binding the contract method 0x25316411.
//
// Solidity: function checkPointBlockInterval() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) CheckPointBlockInterval() (*big.Int, error) {
	return _StakeManager.Contract.CheckPointBlockInterval(&_StakeManager.CallOpts)
}

// CheckpointRewardDelta is a free data retrieval call binding the contract method 0x7c7eaf1a.
//
// Solidity: function checkpointRewardDelta() view returns(uint256)
func (_StakeManager *StakeManagerCaller) CheckpointRewardDelta(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "checkpointRewardDelta")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckpointRewardDelta is a free data retrieval call binding the contract method 0x7c7eaf1a.
//
// Solidity: function checkpointRewardDelta() view returns(uint256)
func (_StakeManager *StakeManagerSession) CheckpointRewardDelta() (*big.Int, error) {
	return _StakeManager.Contract.CheckpointRewardDelta(&_StakeManager.CallOpts)
}

// CheckpointRewardDelta is a free data retrieval call binding the contract method 0x7c7eaf1a.
//
// Solidity: function checkpointRewardDelta() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) CheckpointRewardDelta() (*big.Int, error) {
	return _StakeManager.Contract.CheckpointRewardDelta(&_StakeManager.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_StakeManager *StakeManagerCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_StakeManager *StakeManagerSession) CurrentEpoch() (*big.Int, error) {
	return _StakeManager.Contract.CurrentEpoch(&_StakeManager.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) CurrentEpoch() (*big.Int, error) {
	return _StakeManager.Contract.CurrentEpoch(&_StakeManager.CallOpts)
}

// CurrentValidatorSetSize is a free data retrieval call binding the contract method 0x7f952d95.
//
// Solidity: function currentValidatorSetSize() view returns(uint256)
func (_StakeManager *StakeManagerCaller) CurrentValidatorSetSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "currentValidatorSetSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentValidatorSetSize is a free data retrieval call binding the contract method 0x7f952d95.
//
// Solidity: function currentValidatorSetSize() view returns(uint256)
func (_StakeManager *StakeManagerSession) CurrentValidatorSetSize() (*big.Int, error) {
	return _StakeManager.Contract.CurrentValidatorSetSize(&_StakeManager.CallOpts)
}

// CurrentValidatorSetSize is a free data retrieval call binding the contract method 0x7f952d95.
//
// Solidity: function currentValidatorSetSize() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) CurrentValidatorSetSize() (*big.Int, error) {
	return _StakeManager.Contract.CurrentValidatorSetSize(&_StakeManager.CallOpts)
}

// CurrentValidatorSetTotalStake is a free data retrieval call binding the contract method 0xa4769071.
//
// Solidity: function currentValidatorSetTotalStake() view returns(uint256)
func (_StakeManager *StakeManagerCaller) CurrentValidatorSetTotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "currentValidatorSetTotalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentValidatorSetTotalStake is a free data retrieval call binding the contract method 0xa4769071.
//
// Solidity: function currentValidatorSetTotalStake() view returns(uint256)
func (_StakeManager *StakeManagerSession) CurrentValidatorSetTotalStake() (*big.Int, error) {
	return _StakeManager.Contract.CurrentValidatorSetTotalStake(&_StakeManager.CallOpts)
}

// CurrentValidatorSetTotalStake is a free data retrieval call binding the contract method 0xa4769071.
//
// Solidity: function currentValidatorSetTotalStake() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) CurrentValidatorSetTotalStake() (*big.Int, error) {
	return _StakeManager.Contract.CurrentValidatorSetTotalStake(&_StakeManager.CallOpts)
}

// DelegatedAmount is a free data retrieval call binding the contract method 0x7f4b4323.
//
// Solidity: function delegatedAmount(uint256 validatorId) view returns(uint256)
func (_StakeManager *StakeManagerCaller) DelegatedAmount(opts *bind.CallOpts, validatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "delegatedAmount", validatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatedAmount is a free data retrieval call binding the contract method 0x7f4b4323.
//
// Solidity: function delegatedAmount(uint256 validatorId) view returns(uint256)
func (_StakeManager *StakeManagerSession) DelegatedAmount(validatorId *big.Int) (*big.Int, error) {
	return _StakeManager.Contract.DelegatedAmount(&_StakeManager.CallOpts, validatorId)
}

// DelegatedAmount is a free data retrieval call binding the contract method 0x7f4b4323.
//
// Solidity: function delegatedAmount(uint256 validatorId) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) DelegatedAmount(validatorId *big.Int) (*big.Int, error) {
	return _StakeManager.Contract.DelegatedAmount(&_StakeManager.CallOpts, validatorId)
}

// DelegationEnabled is a free data retrieval call binding the contract method 0x54b8c601.
//
// Solidity: function delegationEnabled() view returns(bool)
func (_StakeManager *StakeManagerCaller) DelegationEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "delegationEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DelegationEnabled is a free data retrieval call binding the contract method 0x54b8c601.
//
// Solidity: function delegationEnabled() view returns(bool)
func (_StakeManager *StakeManagerSession) DelegationEnabled() (bool, error) {
	return _StakeManager.Contract.DelegationEnabled(&_StakeManager.CallOpts)
}

// DelegationEnabled is a free data retrieval call binding the contract method 0x54b8c601.
//
// Solidity: function delegationEnabled() view returns(bool)
func (_StakeManager *StakeManagerCallerSession) DelegationEnabled() (bool, error) {
	return _StakeManager.Contract.DelegationEnabled(&_StakeManager.CallOpts)
}

// DelegatorsReward is a free data retrieval call binding the contract method 0x39610f78.
//
// Solidity: function delegatorsReward(uint256 validatorId) view returns(uint256)
func (_StakeManager *StakeManagerCaller) DelegatorsReward(opts *bind.CallOpts, validatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "delegatorsReward", validatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatorsReward is a free data retrieval call binding the contract method 0x39610f78.
//
// Solidity: function delegatorsReward(uint256 validatorId) view returns(uint256)
func (_StakeManager *StakeManagerSession) DelegatorsReward(validatorId *big.Int) (*big.Int, error) {
	return _StakeManager.Contract.DelegatorsReward(&_StakeManager.CallOpts, validatorId)
}

// DelegatorsReward is a free data retrieval call binding the contract method 0x39610f78.
//
// Solidity: function delegatorsReward(uint256 validatorId) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) DelegatorsReward(validatorId *big.Int) (*big.Int, error) {
	return _StakeManager.Contract.DelegatorsReward(&_StakeManager.CallOpts, validatorId)
}

// Dynasty is a free data retrieval call binding the contract method 0x7060054d.
//
// Solidity: function dynasty() view returns(uint256)
func (_StakeManager *StakeManagerCaller) Dynasty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "dynasty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dynasty is a free data retrieval call binding the contract method 0x7060054d.
//
// Solidity: function dynasty() view returns(uint256)
func (_StakeManager *StakeManagerSession) Dynasty() (*big.Int, error) {
	return _StakeManager.Contract.Dynasty(&_StakeManager.CallOpts)
}

// Dynasty is a free data retrieval call binding the contract method 0x7060054d.
//
// Solidity: function dynasty() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) Dynasty() (*big.Int, error) {
	return _StakeManager.Contract.Dynasty(&_StakeManager.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_StakeManager *StakeManagerCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_StakeManager *StakeManagerSession) Epoch() (*big.Int, error) {
	return _StakeManager.Contract.Epoch(&_StakeManager.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) Epoch() (*big.Int, error) {
	return _StakeManager.Contract.Epoch(&_StakeManager.CallOpts)
}

// EventsHub is a free data retrieval call binding the contract method 0x883b455f.
//
// Solidity: function eventsHub() view returns(address)
func (_StakeManager *StakeManagerCaller) EventsHub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "eventsHub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EventsHub is a free data retrieval call binding the contract method 0x883b455f.
//
// Solidity: function eventsHub() view returns(address)
func (_StakeManager *StakeManagerSession) EventsHub() (common.Address, error) {
	return _StakeManager.Contract.EventsHub(&_StakeManager.CallOpts)
}

// EventsHub is a free data retrieval call binding the contract method 0x883b455f.
//
// Solidity: function eventsHub() view returns(address)
func (_StakeManager *StakeManagerCallerSession) EventsHub() (common.Address, error) {
	return _StakeManager.Contract.EventsHub(&_StakeManager.CallOpts)
}

// ExtensionCode is a free data retrieval call binding the contract method 0xf8a3176c.
//
// Solidity: function extensionCode() view returns(address)
func (_StakeManager *StakeManagerCaller) ExtensionCode(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "extensionCode")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExtensionCode is a free data retrieval call binding the contract method 0xf8a3176c.
//
// Solidity: function extensionCode() view returns(address)
func (_StakeManager *StakeManagerSession) ExtensionCode() (common.Address, error) {
	return _StakeManager.Contract.ExtensionCode(&_StakeManager.CallOpts)
}

// ExtensionCode is a free data retrieval call binding the contract method 0xf8a3176c.
//
// Solidity: function extensionCode() view returns(address)
func (_StakeManager *StakeManagerCallerSession) ExtensionCode() (common.Address, error) {
	return _StakeManager.Contract.ExtensionCode(&_StakeManager.CallOpts)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_StakeManager *StakeManagerCaller) GetRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_StakeManager *StakeManagerSession) GetRegistry() (common.Address, error) {
	return _StakeManager.Contract.GetRegistry(&_StakeManager.CallOpts)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_StakeManager *StakeManagerCallerSession) GetRegistry() (common.Address, error) {
	return _StakeManager.Contract.GetRegistry(&_StakeManager.CallOpts)
}

// GetValidatorContract is a free data retrieval call binding the contract method 0x56342d8c.
//
// Solidity: function getValidatorContract(uint256 validatorId) view returns(address)
func (_StakeManager *StakeManagerCaller) GetValidatorContract(opts *bind.CallOpts, validatorId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getValidatorContract", validatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetValidatorContract is a free data retrieval call binding the contract method 0x56342d8c.
//
// Solidity: function getValidatorContract(uint256 validatorId) view returns(address)
func (_StakeManager *StakeManagerSession) GetValidatorContract(validatorId *big.Int) (common.Address, error) {
	return _StakeManager.Contract.GetValidatorContract(&_StakeManager.CallOpts, validatorId)
}

// GetValidatorContract is a free data retrieval call binding the contract method 0x56342d8c.
//
// Solidity: function getValidatorContract(uint256 validatorId) view returns(address)
func (_StakeManager *StakeManagerCallerSession) GetValidatorContract(validatorId *big.Int) (common.Address, error) {
	return _StakeManager.Contract.GetValidatorContract(&_StakeManager.CallOpts, validatorId)
}

// GetValidatorId is a free data retrieval call binding the contract method 0x174e6832.
//
// Solidity: function getValidatorId(address user) view returns(uint256)
func (_StakeManager *StakeManagerCaller) GetValidatorId(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getValidatorId", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorId is a free data retrieval call binding the contract method 0x174e6832.
//
// Solidity: function getValidatorId(address user) view returns(uint256)
func (_StakeManager *StakeManagerSession) GetValidatorId(user common.Address) (*big.Int, error) {
	return _StakeManager.Contract.GetValidatorId(&_StakeManager.CallOpts, user)
}

// GetValidatorId is a free data retrieval call binding the contract method 0x174e6832.
//
// Solidity: function getValidatorId(address user) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) GetValidatorId(user common.Address) (*big.Int, error) {
	return _StakeManager.Contract.GetValidatorId(&_StakeManager.CallOpts, user)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_StakeManager *StakeManagerCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_StakeManager *StakeManagerSession) Governance() (common.Address, error) {
	return _StakeManager.Contract.Governance(&_StakeManager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_StakeManager *StakeManagerCallerSession) Governance() (common.Address, error) {
	return _StakeManager.Contract.Governance(&_StakeManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_StakeManager *StakeManagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_StakeManager *StakeManagerSession) IsOwner() (bool, error) {
	return _StakeManager.Contract.IsOwner(&_StakeManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_StakeManager *StakeManagerCallerSession) IsOwner() (bool, error) {
	return _StakeManager.Contract.IsOwner(&_StakeManager.CallOpts)
}

// IsValidator is a free data retrieval call binding the contract method 0x2649263a.
//
// Solidity: function isValidator(uint256 validatorId) view returns(bool)
func (_StakeManager *StakeManagerCaller) IsValidator(opts *bind.CallOpts, validatorId *big.Int) (bool, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "isValidator", validatorId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0x2649263a.
//
// Solidity: function isValidator(uint256 validatorId) view returns(bool)
func (_StakeManager *StakeManagerSession) IsValidator(validatorId *big.Int) (bool, error) {
	return _StakeManager.Contract.IsValidator(&_StakeManager.CallOpts, validatorId)
}

// IsValidator is a free data retrieval call binding the contract method 0x2649263a.
//
// Solidity: function isValidator(uint256 validatorId) view returns(bool)
func (_StakeManager *StakeManagerCallerSession) IsValidator(validatorId *big.Int) (bool, error) {
	return _StakeManager.Contract.IsValidator(&_StakeManager.CallOpts, validatorId)
}

// LatestSignerUpdateEpoch is a free data retrieval call binding the contract method 0xd7f5549d.
//
// Solidity: function latestSignerUpdateEpoch(uint256 ) view returns(uint256)
func (_StakeManager *StakeManagerCaller) LatestSignerUpdateEpoch(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "latestSignerUpdateEpoch", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestSignerUpdateEpoch is a free data retrieval call binding the contract method 0xd7f5549d.
//
// Solidity: function latestSignerUpdateEpoch(uint256 ) view returns(uint256)
func (_StakeManager *StakeManagerSession) LatestSignerUpdateEpoch(arg0 *big.Int) (*big.Int, error) {
	return _StakeManager.Contract.LatestSignerUpdateEpoch(&_StakeManager.CallOpts, arg0)
}

// LatestSignerUpdateEpoch is a free data retrieval call binding the contract method 0xd7f5549d.
//
// Solidity: function latestSignerUpdateEpoch(uint256 ) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) LatestSignerUpdateEpoch(arg0 *big.Int) (*big.Int, error) {
	return _StakeManager.Contract.LatestSignerUpdateEpoch(&_StakeManager.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_StakeManager *StakeManagerCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "locked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_StakeManager *StakeManagerSession) Locked() (bool, error) {
	return _StakeManager.Contract.Locked(&_StakeManager.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_StakeManager *StakeManagerCallerSession) Locked() (bool, error) {
	return _StakeManager.Contract.Locked(&_StakeManager.CallOpts)
}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_StakeManager *StakeManagerCaller) Logger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "logger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_StakeManager *StakeManagerSession) Logger() (common.Address, error) {
	return _StakeManager.Contract.Logger(&_StakeManager.CallOpts)
}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_StakeManager *StakeManagerCallerSession) Logger() (common.Address, error) {
	return _StakeManager.Contract.Logger(&_StakeManager.CallOpts)
}

// MaxRewardedCheckpoints is a free data retrieval call binding the contract method 0x451b5985.
//
// Solidity: function maxRewardedCheckpoints() view returns(uint256)
func (_StakeManager *StakeManagerCaller) MaxRewardedCheckpoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "maxRewardedCheckpoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRewardedCheckpoints is a free data retrieval call binding the contract method 0x451b5985.
//
// Solidity: function maxRewardedCheckpoints() view returns(uint256)
func (_StakeManager *StakeManagerSession) MaxRewardedCheckpoints() (*big.Int, error) {
	return _StakeManager.Contract.MaxRewardedCheckpoints(&_StakeManager.CallOpts)
}

// MaxRewardedCheckpoints is a free data retrieval call binding the contract method 0x451b5985.
//
// Solidity: function maxRewardedCheckpoints() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) MaxRewardedCheckpoints() (*big.Int, error) {
	return _StakeManager.Contract.MaxRewardedCheckpoints(&_StakeManager.CallOpts)
}

// Migration is a free data retrieval call binding the contract method 0x1705a3bd.
//
// Solidity: function migration() view returns(address)
func (_StakeManager *StakeManagerCaller) Migration(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "migration")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migration is a free data retrieval call binding the contract method 0x1705a3bd.
//
// Solidity: function migration() view returns(address)
func (_StakeManager *StakeManagerSession) Migration() (common.Address, error) {
	return _StakeManager.Contract.Migration(&_StakeManager.CallOpts)
}

// Migration is a free data retrieval call binding the contract method 0x1705a3bd.
//
// Solidity: function migration() view returns(address)
func (_StakeManager *StakeManagerCallerSession) Migration() (common.Address, error) {
	return _StakeManager.Contract.Migration(&_StakeManager.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_StakeManager *StakeManagerCaller) MinDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "minDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_StakeManager *StakeManagerSession) MinDeposit() (*big.Int, error) {
	return _StakeManager.Contract.MinDeposit(&_StakeManager.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) MinDeposit() (*big.Int, error) {
	return _StakeManager.Contract.MinDeposit(&_StakeManager.CallOpts)
}

// MinHeimdallFee is a free data retrieval call binding the contract method 0xfba58f34.
//
// Solidity: function minHeimdallFee() view returns(uint256)
func (_StakeManager *StakeManagerCaller) MinHeimdallFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "minHeimdallFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinHeimdallFee is a free data retrieval call binding the contract method 0xfba58f34.
//
// Solidity: function minHeimdallFee() view returns(uint256)
func (_StakeManager *StakeManagerSession) MinHeimdallFee() (*big.Int, error) {
	return _StakeManager.Contract.MinHeimdallFee(&_StakeManager.CallOpts)
}

// MinHeimdallFee is a free data retrieval call binding the contract method 0xfba58f34.
//
// Solidity: function minHeimdallFee() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) MinHeimdallFee() (*big.Int, error) {
	return _StakeManager.Contract.MinHeimdallFee(&_StakeManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakeManager *StakeManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakeManager *StakeManagerSession) Owner() (common.Address, error) {
	return _StakeManager.Contract.Owner(&_StakeManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakeManager *StakeManagerCallerSession) Owner() (common.Address, error) {
	return _StakeManager.Contract.Owner(&_StakeManager.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_StakeManager *StakeManagerCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_StakeManager *StakeManagerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _StakeManager.Contract.OwnerOf(&_StakeManager.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_StakeManager *StakeManagerCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _StakeManager.Contract.OwnerOf(&_StakeManager.CallOpts, tokenId)
}

// PrevBlockInterval is a free data retrieval call binding the contract method 0x91f1a3a5.
//
// Solidity: function prevBlockInterval() view returns(uint256)
func (_StakeManager *StakeManagerCaller) PrevBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "prevBlockInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrevBlockInterval is a free data retrieval call binding the contract method 0x91f1a3a5.
//
// Solidity: function prevBlockInterval() view returns(uint256)
func (_StakeManager *StakeManagerSession) PrevBlockInterval() (*big.Int, error) {
	return _StakeManager.Contract.PrevBlockInterval(&_StakeManager.CallOpts)
}

// PrevBlockInterval is a free data retrieval call binding the contract method 0x91f1a3a5.
//
// Solidity: function prevBlockInterval() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) PrevBlockInterval() (*big.Int, error) {
	return _StakeManager.Contract.PrevBlockInterval(&_StakeManager.CallOpts)
}

// ProposerBonus is a free data retrieval call binding the contract method 0x34274586.
//
// Solidity: function proposerBonus() view returns(uint256)
func (_StakeManager *StakeManagerCaller) ProposerBonus(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "proposerBonus")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposerBonus is a free data retrieval call binding the contract method 0x34274586.
//
// Solidity: function proposerBonus() view returns(uint256)
func (_StakeManager *StakeManagerSession) ProposerBonus() (*big.Int, error) {
	return _StakeManager.Contract.ProposerBonus(&_StakeManager.CallOpts)
}

// ProposerBonus is a free data retrieval call binding the contract method 0x34274586.
//
// Solidity: function proposerBonus() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) ProposerBonus() (*big.Int, error) {
	return _StakeManager.Contract.ProposerBonus(&_StakeManager.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_StakeManager *StakeManagerCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_StakeManager *StakeManagerSession) Registry() (common.Address, error) {
	return _StakeManager.Contract.Registry(&_StakeManager.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_StakeManager *StakeManagerCallerSession) Registry() (common.Address, error) {
	return _StakeManager.Contract.Registry(&_StakeManager.CallOpts)
}

// ReplacementCoolDown is a free data retrieval call binding the contract method 0x77939d10.
//
// Solidity: function replacementCoolDown() view returns(uint256)
func (_StakeManager *StakeManagerCaller) ReplacementCoolDown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "replacementCoolDown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReplacementCoolDown is a free data retrieval call binding the contract method 0x77939d10.
//
// Solidity: function replacementCoolDown() view returns(uint256)
func (_StakeManager *StakeManagerSession) ReplacementCoolDown() (*big.Int, error) {
	return _StakeManager.Contract.ReplacementCoolDown(&_StakeManager.CallOpts)
}

// ReplacementCoolDown is a free data retrieval call binding the contract method 0x77939d10.
//
// Solidity: function replacementCoolDown() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) ReplacementCoolDown() (*big.Int, error) {
	return _StakeManager.Contract.ReplacementCoolDown(&_StakeManager.CallOpts)
}

// RewardDecreasePerCheckpoint is a free data retrieval call binding the contract method 0xe568959a.
//
// Solidity: function rewardDecreasePerCheckpoint() view returns(uint256)
func (_StakeManager *StakeManagerCaller) RewardDecreasePerCheckpoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "rewardDecreasePerCheckpoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardDecreasePerCheckpoint is a free data retrieval call binding the contract method 0xe568959a.
//
// Solidity: function rewardDecreasePerCheckpoint() view returns(uint256)
func (_StakeManager *StakeManagerSession) RewardDecreasePerCheckpoint() (*big.Int, error) {
	return _StakeManager.Contract.RewardDecreasePerCheckpoint(&_StakeManager.CallOpts)
}

// RewardDecreasePerCheckpoint is a free data retrieval call binding the contract method 0xe568959a.
//
// Solidity: function rewardDecreasePerCheckpoint() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) RewardDecreasePerCheckpoint() (*big.Int, error) {
	return _StakeManager.Contract.RewardDecreasePerCheckpoint(&_StakeManager.CallOpts)
}

// RewardPerStake is a free data retrieval call binding the contract method 0xa8dc889b.
//
// Solidity: function rewardPerStake() view returns(uint256)
func (_StakeManager *StakeManagerCaller) RewardPerStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "rewardPerStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerStake is a free data retrieval call binding the contract method 0xa8dc889b.
//
// Solidity: function rewardPerStake() view returns(uint256)
func (_StakeManager *StakeManagerSession) RewardPerStake() (*big.Int, error) {
	return _StakeManager.Contract.RewardPerStake(&_StakeManager.CallOpts)
}

// RewardPerStake is a free data retrieval call binding the contract method 0xa8dc889b.
//
// Solidity: function rewardPerStake() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) RewardPerStake() (*big.Int, error) {
	return _StakeManager.Contract.RewardPerStake(&_StakeManager.CallOpts)
}

// RootChain is a free data retrieval call binding the contract method 0x987ab9db.
//
// Solidity: function rootChain() view returns(address)
func (_StakeManager *StakeManagerCaller) RootChain(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "rootChain")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RootChain is a free data retrieval call binding the contract method 0x987ab9db.
//
// Solidity: function rootChain() view returns(address)
func (_StakeManager *StakeManagerSession) RootChain() (common.Address, error) {
	return _StakeManager.Contract.RootChain(&_StakeManager.CallOpts)
}

// RootChain is a free data retrieval call binding the contract method 0x987ab9db.
//
// Solidity: function rootChain() view returns(address)
func (_StakeManager *StakeManagerCallerSession) RootChain() (common.Address, error) {
	return _StakeManager.Contract.RootChain(&_StakeManager.CallOpts)
}

// SignerToValidator is a free data retrieval call binding the contract method 0x3862da0b.
//
// Solidity: function signerToValidator(address ) view returns(uint256)
func (_StakeManager *StakeManagerCaller) SignerToValidator(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "signerToValidator", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignerToValidator is a free data retrieval call binding the contract method 0x3862da0b.
//
// Solidity: function signerToValidator(address ) view returns(uint256)
func (_StakeManager *StakeManagerSession) SignerToValidator(arg0 common.Address) (*big.Int, error) {
	return _StakeManager.Contract.SignerToValidator(&_StakeManager.CallOpts, arg0)
}

// SignerToValidator is a free data retrieval call binding the contract method 0x3862da0b.
//
// Solidity: function signerToValidator(address ) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) SignerToValidator(arg0 common.Address) (*big.Int, error) {
	return _StakeManager.Contract.SignerToValidator(&_StakeManager.CallOpts, arg0)
}

// SignerUpdateLimit is a free data retrieval call binding the contract method 0x4e3c83f1.
//
// Solidity: function signerUpdateLimit() view returns(uint256)
func (_StakeManager *StakeManagerCaller) SignerUpdateLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "signerUpdateLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignerUpdateLimit is a free data retrieval call binding the contract method 0x4e3c83f1.
//
// Solidity: function signerUpdateLimit() view returns(uint256)
func (_StakeManager *StakeManagerSession) SignerUpdateLimit() (*big.Int, error) {
	return _StakeManager.Contract.SignerUpdateLimit(&_StakeManager.CallOpts)
}

// SignerUpdateLimit is a free data retrieval call binding the contract method 0x4e3c83f1.
//
// Solidity: function signerUpdateLimit() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) SignerUpdateLimit() (*big.Int, error) {
	return _StakeManager.Contract.SignerUpdateLimit(&_StakeManager.CallOpts)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_StakeManager *StakeManagerCaller) Signers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "signers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_StakeManager *StakeManagerSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _StakeManager.Contract.Signers(&_StakeManager.CallOpts, arg0)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_StakeManager *StakeManagerCallerSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _StakeManager.Contract.Signers(&_StakeManager.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StakeManager *StakeManagerCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StakeManager *StakeManagerSession) Token() (common.Address, error) {
	return _StakeManager.Contract.Token(&_StakeManager.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StakeManager *StakeManagerCallerSession) Token() (common.Address, error) {
	return _StakeManager.Contract.Token(&_StakeManager.CallOpts)
}

// TokenMatic is a free data retrieval call binding the contract method 0xf1674fdf.
//
// Solidity: function tokenMatic() view returns(address)
func (_StakeManager *StakeManagerCaller) TokenMatic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "tokenMatic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMatic is a free data retrieval call binding the contract method 0xf1674fdf.
//
// Solidity: function tokenMatic() view returns(address)
func (_StakeManager *StakeManagerSession) TokenMatic() (common.Address, error) {
	return _StakeManager.Contract.TokenMatic(&_StakeManager.CallOpts)
}

// TokenMatic is a free data retrieval call binding the contract method 0xf1674fdf.
//
// Solidity: function tokenMatic() view returns(address)
func (_StakeManager *StakeManagerCallerSession) TokenMatic() (common.Address, error) {
	return _StakeManager.Contract.TokenMatic(&_StakeManager.CallOpts)
}

// TotalHeimdallFee is a free data retrieval call binding the contract method 0x9a8a6243.
//
// Solidity: function totalHeimdallFee() view returns(uint256)
func (_StakeManager *StakeManagerCaller) TotalHeimdallFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "totalHeimdallFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalHeimdallFee is a free data retrieval call binding the contract method 0x9a8a6243.
//
// Solidity: function totalHeimdallFee() view returns(uint256)
func (_StakeManager *StakeManagerSession) TotalHeimdallFee() (*big.Int, error) {
	return _StakeManager.Contract.TotalHeimdallFee(&_StakeManager.CallOpts)
}

// TotalHeimdallFee is a free data retrieval call binding the contract method 0x9a8a6243.
//
// Solidity: function totalHeimdallFee() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) TotalHeimdallFee() (*big.Int, error) {
	return _StakeManager.Contract.TotalHeimdallFee(&_StakeManager.CallOpts)
}

// TotalRewards is a free data retrieval call binding the contract method 0x0e15561a.
//
// Solidity: function totalRewards() view returns(uint256)
func (_StakeManager *StakeManagerCaller) TotalRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "totalRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRewards is a free data retrieval call binding the contract method 0x0e15561a.
//
// Solidity: function totalRewards() view returns(uint256)
func (_StakeManager *StakeManagerSession) TotalRewards() (*big.Int, error) {
	return _StakeManager.Contract.TotalRewards(&_StakeManager.CallOpts)
}

// TotalRewards is a free data retrieval call binding the contract method 0x0e15561a.
//
// Solidity: function totalRewards() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) TotalRewards() (*big.Int, error) {
	return _StakeManager.Contract.TotalRewards(&_StakeManager.CallOpts)
}

// TotalRewardsLiquidated is a free data retrieval call binding the contract method 0xcd6b8388.
//
// Solidity: function totalRewardsLiquidated() view returns(uint256)
func (_StakeManager *StakeManagerCaller) TotalRewardsLiquidated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "totalRewardsLiquidated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRewardsLiquidated is a free data retrieval call binding the contract method 0xcd6b8388.
//
// Solidity: function totalRewardsLiquidated() view returns(uint256)
func (_StakeManager *StakeManagerSession) TotalRewardsLiquidated() (*big.Int, error) {
	return _StakeManager.Contract.TotalRewardsLiquidated(&_StakeManager.CallOpts)
}

// TotalRewardsLiquidated is a free data retrieval call binding the contract method 0xcd6b8388.
//
// Solidity: function totalRewardsLiquidated() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) TotalRewardsLiquidated() (*big.Int, error) {
	return _StakeManager.Contract.TotalRewardsLiquidated(&_StakeManager.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_StakeManager *StakeManagerCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_StakeManager *StakeManagerSession) TotalStaked() (*big.Int, error) {
	return _StakeManager.Contract.TotalStaked(&_StakeManager.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) TotalStaked() (*big.Int, error) {
	return _StakeManager.Contract.TotalStaked(&_StakeManager.CallOpts)
}

// TotalStakedFor is a free data retrieval call binding the contract method 0x4b341aed.
//
// Solidity: function totalStakedFor(address user) view returns(uint256)
func (_StakeManager *StakeManagerCaller) TotalStakedFor(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "totalStakedFor", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakedFor is a free data retrieval call binding the contract method 0x4b341aed.
//
// Solidity: function totalStakedFor(address user) view returns(uint256)
func (_StakeManager *StakeManagerSession) TotalStakedFor(user common.Address) (*big.Int, error) {
	return _StakeManager.Contract.TotalStakedFor(&_StakeManager.CallOpts, user)
}

// TotalStakedFor is a free data retrieval call binding the contract method 0x4b341aed.
//
// Solidity: function totalStakedFor(address user) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) TotalStakedFor(user common.Address) (*big.Int, error) {
	return _StakeManager.Contract.TotalStakedFor(&_StakeManager.CallOpts, user)
}

// UserFeeExit is a free data retrieval call binding the contract method 0x78f84a44.
//
// Solidity: function userFeeExit(address ) view returns(uint256)
func (_StakeManager *StakeManagerCaller) UserFeeExit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "userFeeExit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserFeeExit is a free data retrieval call binding the contract method 0x78f84a44.
//
// Solidity: function userFeeExit(address ) view returns(uint256)
func (_StakeManager *StakeManagerSession) UserFeeExit(arg0 common.Address) (*big.Int, error) {
	return _StakeManager.Contract.UserFeeExit(&_StakeManager.CallOpts, arg0)
}

// UserFeeExit is a free data retrieval call binding the contract method 0x78f84a44.
//
// Solidity: function userFeeExit(address ) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) UserFeeExit(arg0 common.Address) (*big.Int, error) {
	return _StakeManager.Contract.UserFeeExit(&_StakeManager.CallOpts, arg0)
}

// ValidatorAuction is a free data retrieval call binding the contract method 0x5325e144.
//
// Solidity: function validatorAuction(uint256 ) view returns(uint256 amount, uint256 startEpoch, address user, bool acceptDelegation, bytes signerPubkey)
func (_StakeManager *StakeManagerCaller) ValidatorAuction(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount           *big.Int
	StartEpoch       *big.Int
	User             common.Address
	AcceptDelegation bool
	SignerPubkey     []byte
}, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "validatorAuction", arg0)

	outstruct := new(struct {
		Amount           *big.Int
		StartEpoch       *big.Int
		User             common.Address
		AcceptDelegation bool
		SignerPubkey     []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartEpoch = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.User = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.AcceptDelegation = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.SignerPubkey = *abi.ConvertType(out[4], new([]byte)).(*[]byte)

	return *outstruct, err

}

// ValidatorAuction is a free data retrieval call binding the contract method 0x5325e144.
//
// Solidity: function validatorAuction(uint256 ) view returns(uint256 amount, uint256 startEpoch, address user, bool acceptDelegation, bytes signerPubkey)
func (_StakeManager *StakeManagerSession) ValidatorAuction(arg0 *big.Int) (struct {
	Amount           *big.Int
	StartEpoch       *big.Int
	User             common.Address
	AcceptDelegation bool
	SignerPubkey     []byte
}, error) {
	return _StakeManager.Contract.ValidatorAuction(&_StakeManager.CallOpts, arg0)
}

// ValidatorAuction is a free data retrieval call binding the contract method 0x5325e144.
//
// Solidity: function validatorAuction(uint256 ) view returns(uint256 amount, uint256 startEpoch, address user, bool acceptDelegation, bytes signerPubkey)
func (_StakeManager *StakeManagerCallerSession) ValidatorAuction(arg0 *big.Int) (struct {
	Amount           *big.Int
	StartEpoch       *big.Int
	User             common.Address
	AcceptDelegation bool
	SignerPubkey     []byte
}, error) {
	return _StakeManager.Contract.ValidatorAuction(&_StakeManager.CallOpts, arg0)
}

// ValidatorReward is a free data retrieval call binding the contract method 0xb65de35e.
//
// Solidity: function validatorReward(uint256 validatorId) view returns(uint256)
func (_StakeManager *StakeManagerCaller) ValidatorReward(opts *bind.CallOpts, validatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "validatorReward", validatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorReward is a free data retrieval call binding the contract method 0xb65de35e.
//
// Solidity: function validatorReward(uint256 validatorId) view returns(uint256)
func (_StakeManager *StakeManagerSession) ValidatorReward(validatorId *big.Int) (*big.Int, error) {
	return _StakeManager.Contract.ValidatorReward(&_StakeManager.CallOpts, validatorId)
}

// ValidatorReward is a free data retrieval call binding the contract method 0xb65de35e.
//
// Solidity: function validatorReward(uint256 validatorId) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) ValidatorReward(validatorId *big.Int) (*big.Int, error) {
	return _StakeManager.Contract.ValidatorReward(&_StakeManager.CallOpts, validatorId)
}

// ValidatorShareFactory is a free data retrieval call binding the contract method 0x1ae4818f.
//
// Solidity: function validatorShareFactory() view returns(address)
func (_StakeManager *StakeManagerCaller) ValidatorShareFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "validatorShareFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorShareFactory is a free data retrieval call binding the contract method 0x1ae4818f.
//
// Solidity: function validatorShareFactory() view returns(address)
func (_StakeManager *StakeManagerSession) ValidatorShareFactory() (common.Address, error) {
	return _StakeManager.Contract.ValidatorShareFactory(&_StakeManager.CallOpts)
}

// ValidatorShareFactory is a free data retrieval call binding the contract method 0x1ae4818f.
//
// Solidity: function validatorShareFactory() view returns(address)
func (_StakeManager *StakeManagerCallerSession) ValidatorShareFactory() (common.Address, error) {
	return _StakeManager.Contract.ValidatorShareFactory(&_StakeManager.CallOpts)
}

// ValidatorStake is a free data retrieval call binding the contract method 0xeceec1d3.
//
// Solidity: function validatorStake(uint256 validatorId) view returns(uint256)
func (_StakeManager *StakeManagerCaller) ValidatorStake(opts *bind.CallOpts, validatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "validatorStake", validatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorStake is a free data retrieval call binding the contract method 0xeceec1d3.
//
// Solidity: function validatorStake(uint256 validatorId) view returns(uint256)
func (_StakeManager *StakeManagerSession) ValidatorStake(validatorId *big.Int) (*big.Int, error) {
	return _StakeManager.Contract.ValidatorStake(&_StakeManager.CallOpts, validatorId)
}

// ValidatorStake is a free data retrieval call binding the contract method 0xeceec1d3.
//
// Solidity: function validatorStake(uint256 validatorId) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) ValidatorStake(validatorId *big.Int) (*big.Int, error) {
	return _StakeManager.Contract.ValidatorStake(&_StakeManager.CallOpts, validatorId)
}

// ValidatorState is a free data retrieval call binding the contract method 0xe59ee0c6.
//
// Solidity: function validatorState() view returns(uint256 amount, uint256 stakerCount)
func (_StakeManager *StakeManagerCaller) ValidatorState(opts *bind.CallOpts) (struct {
	Amount      *big.Int
	StakerCount *big.Int
}, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "validatorState")

	outstruct := new(struct {
		Amount      *big.Int
		StakerCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StakerCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ValidatorState is a free data retrieval call binding the contract method 0xe59ee0c6.
//
// Solidity: function validatorState() view returns(uint256 amount, uint256 stakerCount)
func (_StakeManager *StakeManagerSession) ValidatorState() (struct {
	Amount      *big.Int
	StakerCount *big.Int
}, error) {
	return _StakeManager.Contract.ValidatorState(&_StakeManager.CallOpts)
}

// ValidatorState is a free data retrieval call binding the contract method 0xe59ee0c6.
//
// Solidity: function validatorState() view returns(uint256 amount, uint256 stakerCount)
func (_StakeManager *StakeManagerCallerSession) ValidatorState() (struct {
	Amount      *big.Int
	StakerCount *big.Int
}, error) {
	return _StakeManager.Contract.ValidatorState(&_StakeManager.CallOpts)
}

// ValidatorStateChanges is a free data retrieval call binding the contract method 0x25726df2.
//
// Solidity: function validatorStateChanges(uint256 ) view returns(int256 amount, int256 stakerCount)
func (_StakeManager *StakeManagerCaller) ValidatorStateChanges(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount      *big.Int
	StakerCount *big.Int
}, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "validatorStateChanges", arg0)

	outstruct := new(struct {
		Amount      *big.Int
		StakerCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StakerCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ValidatorStateChanges is a free data retrieval call binding the contract method 0x25726df2.
//
// Solidity: function validatorStateChanges(uint256 ) view returns(int256 amount, int256 stakerCount)
func (_StakeManager *StakeManagerSession) ValidatorStateChanges(arg0 *big.Int) (struct {
	Amount      *big.Int
	StakerCount *big.Int
}, error) {
	return _StakeManager.Contract.ValidatorStateChanges(&_StakeManager.CallOpts, arg0)
}

// ValidatorStateChanges is a free data retrieval call binding the contract method 0x25726df2.
//
// Solidity: function validatorStateChanges(uint256 ) view returns(int256 amount, int256 stakerCount)
func (_StakeManager *StakeManagerCallerSession) ValidatorStateChanges(arg0 *big.Int) (struct {
	Amount      *big.Int
	StakerCount *big.Int
}, error) {
	return _StakeManager.Contract.ValidatorStateChanges(&_StakeManager.CallOpts, arg0)
}

// ValidatorThreshold is a free data retrieval call binding the contract method 0x4fd101d7.
//
// Solidity: function validatorThreshold() view returns(uint256)
func (_StakeManager *StakeManagerCaller) ValidatorThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "validatorThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorThreshold is a free data retrieval call binding the contract method 0x4fd101d7.
//
// Solidity: function validatorThreshold() view returns(uint256)
func (_StakeManager *StakeManagerSession) ValidatorThreshold() (*big.Int, error) {
	return _StakeManager.Contract.ValidatorThreshold(&_StakeManager.CallOpts)
}

// ValidatorThreshold is a free data retrieval call binding the contract method 0x4fd101d7.
//
// Solidity: function validatorThreshold() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) ValidatorThreshold() (*big.Int, error) {
	return _StakeManager.Contract.ValidatorThreshold(&_StakeManager.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(uint256 amount, uint256 reward, uint256 activationEpoch, uint256 deactivationEpoch, uint256 jailTime, address signer, address contractAddress, uint8 status, uint256 commissionRate, uint256 lastCommissionUpdate, uint256 delegatorsReward, uint256 delegatedAmount, uint256 initialRewardPerStake)
func (_StakeManager *StakeManagerCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount                *big.Int
	Reward                *big.Int
	ActivationEpoch       *big.Int
	DeactivationEpoch     *big.Int
	JailTime              *big.Int
	Signer                common.Address
	ContractAddress       common.Address
	Status                uint8
	CommissionRate        *big.Int
	LastCommissionUpdate  *big.Int
	DelegatorsReward      *big.Int
	DelegatedAmount       *big.Int
	InitialRewardPerStake *big.Int
}, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		Amount                *big.Int
		Reward                *big.Int
		ActivationEpoch       *big.Int
		DeactivationEpoch     *big.Int
		JailTime              *big.Int
		Signer                common.Address
		ContractAddress       common.Address
		Status                uint8
		CommissionRate        *big.Int
		LastCommissionUpdate  *big.Int
		DelegatorsReward      *big.Int
		DelegatedAmount       *big.Int
		InitialRewardPerStake *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ActivationEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DeactivationEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.JailTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Signer = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.ContractAddress = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.CommissionRate = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.LastCommissionUpdate = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.DelegatorsReward = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.DelegatedAmount = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.InitialRewardPerStake = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(uint256 amount, uint256 reward, uint256 activationEpoch, uint256 deactivationEpoch, uint256 jailTime, address signer, address contractAddress, uint8 status, uint256 commissionRate, uint256 lastCommissionUpdate, uint256 delegatorsReward, uint256 delegatedAmount, uint256 initialRewardPerStake)
func (_StakeManager *StakeManagerSession) Validators(arg0 *big.Int) (struct {
	Amount                *big.Int
	Reward                *big.Int
	ActivationEpoch       *big.Int
	DeactivationEpoch     *big.Int
	JailTime              *big.Int
	Signer                common.Address
	ContractAddress       common.Address
	Status                uint8
	CommissionRate        *big.Int
	LastCommissionUpdate  *big.Int
	DelegatorsReward      *big.Int
	DelegatedAmount       *big.Int
	InitialRewardPerStake *big.Int
}, error) {
	return _StakeManager.Contract.Validators(&_StakeManager.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(uint256 amount, uint256 reward, uint256 activationEpoch, uint256 deactivationEpoch, uint256 jailTime, address signer, address contractAddress, uint8 status, uint256 commissionRate, uint256 lastCommissionUpdate, uint256 delegatorsReward, uint256 delegatedAmount, uint256 initialRewardPerStake)
func (_StakeManager *StakeManagerCallerSession) Validators(arg0 *big.Int) (struct {
	Amount                *big.Int
	Reward                *big.Int
	ActivationEpoch       *big.Int
	DeactivationEpoch     *big.Int
	JailTime              *big.Int
	Signer                common.Address
	ContractAddress       common.Address
	Status                uint8
	CommissionRate        *big.Int
	LastCommissionUpdate  *big.Int
	DelegatorsReward      *big.Int
	DelegatedAmount       *big.Int
	InitialRewardPerStake *big.Int
}, error) {
	return _StakeManager.Contract.Validators(&_StakeManager.CallOpts, arg0)
}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xa7ab6961.
//
// Solidity: function withdrawalDelay() view returns(uint256)
func (_StakeManager *StakeManagerCaller) WithdrawalDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "withdrawalDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xa7ab6961.
//
// Solidity: function withdrawalDelay() view returns(uint256)
func (_StakeManager *StakeManagerSession) WithdrawalDelay() (*big.Int, error) {
	return _StakeManager.Contract.WithdrawalDelay(&_StakeManager.CallOpts)
}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xa7ab6961.
//
// Solidity: function withdrawalDelay() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) WithdrawalDelay() (*big.Int, error) {
	return _StakeManager.Contract.WithdrawalDelay(&_StakeManager.CallOpts)
}

// ChangeRootChain is a paid mutator transaction binding the contract method 0xe8afa8e8.
//
// Solidity: function changeRootChain(address newRootChain) returns()
func (_StakeManager *StakeManagerTransactor) ChangeRootChain(opts *bind.TransactOpts, newRootChain common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "changeRootChain", newRootChain)
}

// ChangeRootChain is a paid mutator transaction binding the contract method 0xe8afa8e8.
//
// Solidity: function changeRootChain(address newRootChain) returns()
func (_StakeManager *StakeManagerSession) ChangeRootChain(newRootChain common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.ChangeRootChain(&_StakeManager.TransactOpts, newRootChain)
}

// ChangeRootChain is a paid mutator transaction binding the contract method 0xe8afa8e8.
//
// Solidity: function changeRootChain(address newRootChain) returns()
func (_StakeManager *StakeManagerTransactorSession) ChangeRootChain(newRootChain common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.ChangeRootChain(&_StakeManager.TransactOpts, newRootChain)
}

// CheckSignatures is a paid mutator transaction binding the contract method 0x2fa9d18b.
//
// Solidity: function checkSignatures(uint256 blockInterval, bytes32 voteHash, bytes32 stateRoot, address proposer, uint256[3][] sigs) returns(uint256)
func (_StakeManager *StakeManagerTransactor) CheckSignatures(opts *bind.TransactOpts, blockInterval *big.Int, voteHash [32]byte, stateRoot [32]byte, proposer common.Address, sigs [][3]*big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "checkSignatures", blockInterval, voteHash, stateRoot, proposer, sigs)
}

// CheckSignatures is a paid mutator transaction binding the contract method 0x2fa9d18b.
//
// Solidity: function checkSignatures(uint256 blockInterval, bytes32 voteHash, bytes32 stateRoot, address proposer, uint256[3][] sigs) returns(uint256)
func (_StakeManager *StakeManagerSession) CheckSignatures(blockInterval *big.Int, voteHash [32]byte, stateRoot [32]byte, proposer common.Address, sigs [][3]*big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.CheckSignatures(&_StakeManager.TransactOpts, blockInterval, voteHash, stateRoot, proposer, sigs)
}

// CheckSignatures is a paid mutator transaction binding the contract method 0x2fa9d18b.
//
// Solidity: function checkSignatures(uint256 blockInterval, bytes32 voteHash, bytes32 stateRoot, address proposer, uint256[3][] sigs) returns(uint256)
func (_StakeManager *StakeManagerTransactorSession) CheckSignatures(blockInterval *big.Int, voteHash [32]byte, stateRoot [32]byte, proposer common.Address, sigs [][3]*big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.CheckSignatures(&_StakeManager.TransactOpts, blockInterval, voteHash, stateRoot, proposer, sigs)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x68cb812a.
//
// Solidity: function claimFee(uint256 accumFeeAmount, uint256 index, bytes proof) returns()
func (_StakeManager *StakeManagerTransactor) ClaimFee(opts *bind.TransactOpts, accumFeeAmount *big.Int, index *big.Int, proof []byte) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "claimFee", accumFeeAmount, index, proof)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x68cb812a.
//
// Solidity: function claimFee(uint256 accumFeeAmount, uint256 index, bytes proof) returns()
func (_StakeManager *StakeManagerSession) ClaimFee(accumFeeAmount *big.Int, index *big.Int, proof []byte) (*types.Transaction, error) {
	return _StakeManager.Contract.ClaimFee(&_StakeManager.TransactOpts, accumFeeAmount, index, proof)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x68cb812a.
//
// Solidity: function claimFee(uint256 accumFeeAmount, uint256 index, bytes proof) returns()
func (_StakeManager *StakeManagerTransactorSession) ClaimFee(accumFeeAmount *big.Int, index *big.Int, proof []byte) (*types.Transaction, error) {
	return _StakeManager.Contract.ClaimFee(&_StakeManager.TransactOpts, accumFeeAmount, index, proof)
}

// ConfirmAuctionBid is a paid mutator transaction binding the contract method 0x99d18f6f.
//
// Solidity: function confirmAuctionBid(uint256 validatorId, uint256 heimdallFee) returns()
func (_StakeManager *StakeManagerTransactor) ConfirmAuctionBid(opts *bind.TransactOpts, validatorId *big.Int, heimdallFee *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "confirmAuctionBid", validatorId, heimdallFee)
}

// ConfirmAuctionBid is a paid mutator transaction binding the contract method 0x99d18f6f.
//
// Solidity: function confirmAuctionBid(uint256 validatorId, uint256 heimdallFee) returns()
func (_StakeManager *StakeManagerSession) ConfirmAuctionBid(validatorId *big.Int, heimdallFee *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.ConfirmAuctionBid(&_StakeManager.TransactOpts, validatorId, heimdallFee)
}

// ConfirmAuctionBid is a paid mutator transaction binding the contract method 0x99d18f6f.
//
// Solidity: function confirmAuctionBid(uint256 validatorId, uint256 heimdallFee) returns()
func (_StakeManager *StakeManagerTransactorSession) ConfirmAuctionBid(validatorId *big.Int, heimdallFee *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.ConfirmAuctionBid(&_StakeManager.TransactOpts, validatorId, heimdallFee)
}

// ConvertMaticToPOL is a paid mutator transaction binding the contract method 0x86a22666.
//
// Solidity: function convertMaticToPOL(uint256 amount) returns()
func (_StakeManager *StakeManagerTransactor) ConvertMaticToPOL(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "convertMaticToPOL", amount)
}

// ConvertMaticToPOL is a paid mutator transaction binding the contract method 0x86a22666.
//
// Solidity: function convertMaticToPOL(uint256 amount) returns()
func (_StakeManager *StakeManagerSession) ConvertMaticToPOL(amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.ConvertMaticToPOL(&_StakeManager.TransactOpts, amount)
}

// ConvertMaticToPOL is a paid mutator transaction binding the contract method 0x86a22666.
//
// Solidity: function convertMaticToPOL(uint256 amount) returns()
func (_StakeManager *StakeManagerTransactorSession) ConvertMaticToPOL(amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.ConvertMaticToPOL(&_StakeManager.TransactOpts, amount)
}

// DecreaseValidatorDelegatedAmount is a paid mutator transaction binding the contract method 0x858a7c03.
//
// Solidity: function decreaseValidatorDelegatedAmount(uint256 validatorId, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactor) DecreaseValidatorDelegatedAmount(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "decreaseValidatorDelegatedAmount", validatorId, amount)
}

// DecreaseValidatorDelegatedAmount is a paid mutator transaction binding the contract method 0x858a7c03.
//
// Solidity: function decreaseValidatorDelegatedAmount(uint256 validatorId, uint256 amount) returns()
func (_StakeManager *StakeManagerSession) DecreaseValidatorDelegatedAmount(validatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.DecreaseValidatorDelegatedAmount(&_StakeManager.TransactOpts, validatorId, amount)
}

// DecreaseValidatorDelegatedAmount is a paid mutator transaction binding the contract method 0x858a7c03.
//
// Solidity: function decreaseValidatorDelegatedAmount(uint256 validatorId, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactorSession) DecreaseValidatorDelegatedAmount(validatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.DecreaseValidatorDelegatedAmount(&_StakeManager.TransactOpts, validatorId, amount)
}

// DelegationDeposit is a paid mutator transaction binding the contract method 0x6901b253.
//
// Solidity: function delegationDeposit(uint256 validatorId, uint256 amount, address delegator) returns(bool)
func (_StakeManager *StakeManagerTransactor) DelegationDeposit(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "delegationDeposit", validatorId, amount, delegator)
}

// DelegationDeposit is a paid mutator transaction binding the contract method 0x6901b253.
//
// Solidity: function delegationDeposit(uint256 validatorId, uint256 amount, address delegator) returns(bool)
func (_StakeManager *StakeManagerSession) DelegationDeposit(validatorId *big.Int, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.DelegationDeposit(&_StakeManager.TransactOpts, validatorId, amount, delegator)
}

// DelegationDeposit is a paid mutator transaction binding the contract method 0x6901b253.
//
// Solidity: function delegationDeposit(uint256 validatorId, uint256 amount, address delegator) returns(bool)
func (_StakeManager *StakeManagerTransactorSession) DelegationDeposit(validatorId *big.Int, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.DelegationDeposit(&_StakeManager.TransactOpts, validatorId, amount, delegator)
}

// DelegationDepositPOL is a paid mutator transaction binding the contract method 0x4c025500.
//
// Solidity: function delegationDepositPOL(uint256 validatorId, uint256 amount, address delegator) returns(bool)
func (_StakeManager *StakeManagerTransactor) DelegationDepositPOL(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "delegationDepositPOL", validatorId, amount, delegator)
}

// DelegationDepositPOL is a paid mutator transaction binding the contract method 0x4c025500.
//
// Solidity: function delegationDepositPOL(uint256 validatorId, uint256 amount, address delegator) returns(bool)
func (_StakeManager *StakeManagerSession) DelegationDepositPOL(validatorId *big.Int, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.DelegationDepositPOL(&_StakeManager.TransactOpts, validatorId, amount, delegator)
}

// DelegationDepositPOL is a paid mutator transaction binding the contract method 0x4c025500.
//
// Solidity: function delegationDepositPOL(uint256 validatorId, uint256 amount, address delegator) returns(bool)
func (_StakeManager *StakeManagerTransactorSession) DelegationDepositPOL(validatorId *big.Int, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.DelegationDepositPOL(&_StakeManager.TransactOpts, validatorId, amount, delegator)
}

// DethroneAndStake is a paid mutator transaction binding the contract method 0x52b8115d.
//
// Solidity: function dethroneAndStake(address auctionUser, uint256 heimdallFee, uint256 validatorId, uint256 auctionAmount, bool acceptDelegation, bytes signerPubkey) returns()
func (_StakeManager *StakeManagerTransactor) DethroneAndStake(opts *bind.TransactOpts, auctionUser common.Address, heimdallFee *big.Int, validatorId *big.Int, auctionAmount *big.Int, acceptDelegation bool, signerPubkey []byte) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "dethroneAndStake", auctionUser, heimdallFee, validatorId, auctionAmount, acceptDelegation, signerPubkey)
}

// DethroneAndStake is a paid mutator transaction binding the contract method 0x52b8115d.
//
// Solidity: function dethroneAndStake(address auctionUser, uint256 heimdallFee, uint256 validatorId, uint256 auctionAmount, bool acceptDelegation, bytes signerPubkey) returns()
func (_StakeManager *StakeManagerSession) DethroneAndStake(auctionUser common.Address, heimdallFee *big.Int, validatorId *big.Int, auctionAmount *big.Int, acceptDelegation bool, signerPubkey []byte) (*types.Transaction, error) {
	return _StakeManager.Contract.DethroneAndStake(&_StakeManager.TransactOpts, auctionUser, heimdallFee, validatorId, auctionAmount, acceptDelegation, signerPubkey)
}

// DethroneAndStake is a paid mutator transaction binding the contract method 0x52b8115d.
//
// Solidity: function dethroneAndStake(address auctionUser, uint256 heimdallFee, uint256 validatorId, uint256 auctionAmount, bool acceptDelegation, bytes signerPubkey) returns()
func (_StakeManager *StakeManagerTransactorSession) DethroneAndStake(auctionUser common.Address, heimdallFee *big.Int, validatorId *big.Int, auctionAmount *big.Int, acceptDelegation bool, signerPubkey []byte) (*types.Transaction, error) {
	return _StakeManager.Contract.DethroneAndStake(&_StakeManager.TransactOpts, auctionUser, heimdallFee, validatorId, auctionAmount, acceptDelegation, signerPubkey)
}

// ForceUnstake is a paid mutator transaction binding the contract method 0x91460149.
//
// Solidity: function forceUnstake(uint256 validatorId) returns()
func (_StakeManager *StakeManagerTransactor) ForceUnstake(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "forceUnstake", validatorId)
}

// ForceUnstake is a paid mutator transaction binding the contract method 0x91460149.
//
// Solidity: function forceUnstake(uint256 validatorId) returns()
func (_StakeManager *StakeManagerSession) ForceUnstake(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.ForceUnstake(&_StakeManager.TransactOpts, validatorId)
}

// ForceUnstake is a paid mutator transaction binding the contract method 0x91460149.
//
// Solidity: function forceUnstake(uint256 validatorId) returns()
func (_StakeManager *StakeManagerTransactorSession) ForceUnstake(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.ForceUnstake(&_StakeManager.TransactOpts, validatorId)
}

// ForceUnstakePOL is a paid mutator transaction binding the contract method 0x7854f354.
//
// Solidity: function forceUnstakePOL(uint256 validatorId) returns()
func (_StakeManager *StakeManagerTransactor) ForceUnstakePOL(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "forceUnstakePOL", validatorId)
}

// ForceUnstakePOL is a paid mutator transaction binding the contract method 0x7854f354.
//
// Solidity: function forceUnstakePOL(uint256 validatorId) returns()
func (_StakeManager *StakeManagerSession) ForceUnstakePOL(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.ForceUnstakePOL(&_StakeManager.TransactOpts, validatorId)
}

// ForceUnstakePOL is a paid mutator transaction binding the contract method 0x7854f354.
//
// Solidity: function forceUnstakePOL(uint256 validatorId) returns()
func (_StakeManager *StakeManagerTransactorSession) ForceUnstakePOL(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.ForceUnstakePOL(&_StakeManager.TransactOpts, validatorId)
}

// Initialize is a paid mutator transaction binding the contract method 0xf5e95acb.
//
// Solidity: function initialize(address _registry, address _rootchain, address _token, address _NFTContract, address _stakingLogger, address _validatorShareFactory, address _governance, address _owner, address _extensionCode) returns()
func (_StakeManager *StakeManagerTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _rootchain common.Address, _token common.Address, _NFTContract common.Address, _stakingLogger common.Address, _validatorShareFactory common.Address, _governance common.Address, _owner common.Address, _extensionCode common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "initialize", _registry, _rootchain, _token, _NFTContract, _stakingLogger, _validatorShareFactory, _governance, _owner, _extensionCode)
}

// Initialize is a paid mutator transaction binding the contract method 0xf5e95acb.
//
// Solidity: function initialize(address _registry, address _rootchain, address _token, address _NFTContract, address _stakingLogger, address _validatorShareFactory, address _governance, address _owner, address _extensionCode) returns()
func (_StakeManager *StakeManagerSession) Initialize(_registry common.Address, _rootchain common.Address, _token common.Address, _NFTContract common.Address, _stakingLogger common.Address, _validatorShareFactory common.Address, _governance common.Address, _owner common.Address, _extensionCode common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.Initialize(&_StakeManager.TransactOpts, _registry, _rootchain, _token, _NFTContract, _stakingLogger, _validatorShareFactory, _governance, _owner, _extensionCode)
}

// Initialize is a paid mutator transaction binding the contract method 0xf5e95acb.
//
// Solidity: function initialize(address _registry, address _rootchain, address _token, address _NFTContract, address _stakingLogger, address _validatorShareFactory, address _governance, address _owner, address _extensionCode) returns()
func (_StakeManager *StakeManagerTransactorSession) Initialize(_registry common.Address, _rootchain common.Address, _token common.Address, _NFTContract common.Address, _stakingLogger common.Address, _validatorShareFactory common.Address, _governance common.Address, _owner common.Address, _extensionCode common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.Initialize(&_StakeManager.TransactOpts, _registry, _rootchain, _token, _NFTContract, _stakingLogger, _validatorShareFactory, _governance, _owner, _extensionCode)
}

// InitializePOL is a paid mutator transaction binding the contract method 0x3c343556.
//
// Solidity: function initializePOL(address _tokenNew, address _migration) returns()
func (_StakeManager *StakeManagerTransactor) InitializePOL(opts *bind.TransactOpts, _tokenNew common.Address, _migration common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "initializePOL", _tokenNew, _migration)
}

// InitializePOL is a paid mutator transaction binding the contract method 0x3c343556.
//
// Solidity: function initializePOL(address _tokenNew, address _migration) returns()
func (_StakeManager *StakeManagerSession) InitializePOL(_tokenNew common.Address, _migration common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.InitializePOL(&_StakeManager.TransactOpts, _tokenNew, _migration)
}

// InitializePOL is a paid mutator transaction binding the contract method 0x3c343556.
//
// Solidity: function initializePOL(address _tokenNew, address _migration) returns()
func (_StakeManager *StakeManagerTransactorSession) InitializePOL(_tokenNew common.Address, _migration common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.InitializePOL(&_StakeManager.TransactOpts, _tokenNew, _migration)
}

// InsertSigners is a paid mutator transaction binding the contract method 0x2cf44a43.
//
// Solidity: function insertSigners(address[] _signers) returns()
func (_StakeManager *StakeManagerTransactor) InsertSigners(opts *bind.TransactOpts, _signers []common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "insertSigners", _signers)
}

// InsertSigners is a paid mutator transaction binding the contract method 0x2cf44a43.
//
// Solidity: function insertSigners(address[] _signers) returns()
func (_StakeManager *StakeManagerSession) InsertSigners(_signers []common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.InsertSigners(&_StakeManager.TransactOpts, _signers)
}

// InsertSigners is a paid mutator transaction binding the contract method 0x2cf44a43.
//
// Solidity: function insertSigners(address[] _signers) returns()
func (_StakeManager *StakeManagerTransactorSession) InsertSigners(_signers []common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.InsertSigners(&_StakeManager.TransactOpts, _signers)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_StakeManager *StakeManagerTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_StakeManager *StakeManagerSession) Lock() (*types.Transaction, error) {
	return _StakeManager.Contract.Lock(&_StakeManager.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_StakeManager *StakeManagerTransactorSession) Lock() (*types.Transaction, error) {
	return _StakeManager.Contract.Lock(&_StakeManager.TransactOpts)
}

// MigrateDelegation is a paid mutator transaction binding the contract method 0xfb1ef52c.
//
// Solidity: function migrateDelegation(uint256 fromValidatorId, uint256 toValidatorId, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactor) MigrateDelegation(opts *bind.TransactOpts, fromValidatorId *big.Int, toValidatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "migrateDelegation", fromValidatorId, toValidatorId, amount)
}

// MigrateDelegation is a paid mutator transaction binding the contract method 0xfb1ef52c.
//
// Solidity: function migrateDelegation(uint256 fromValidatorId, uint256 toValidatorId, uint256 amount) returns()
func (_StakeManager *StakeManagerSession) MigrateDelegation(fromValidatorId *big.Int, toValidatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.MigrateDelegation(&_StakeManager.TransactOpts, fromValidatorId, toValidatorId, amount)
}

// MigrateDelegation is a paid mutator transaction binding the contract method 0xfb1ef52c.
//
// Solidity: function migrateDelegation(uint256 fromValidatorId, uint256 toValidatorId, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactorSession) MigrateDelegation(fromValidatorId *big.Int, toValidatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.MigrateDelegation(&_StakeManager.TransactOpts, fromValidatorId, toValidatorId, amount)
}

// MigrateValidatorsData is a paid mutator transaction binding the contract method 0x9ddbbf85.
//
// Solidity: function migrateValidatorsData(uint256 validatorIdFrom, uint256 validatorIdTo) returns()
func (_StakeManager *StakeManagerTransactor) MigrateValidatorsData(opts *bind.TransactOpts, validatorIdFrom *big.Int, validatorIdTo *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "migrateValidatorsData", validatorIdFrom, validatorIdTo)
}

// MigrateValidatorsData is a paid mutator transaction binding the contract method 0x9ddbbf85.
//
// Solidity: function migrateValidatorsData(uint256 validatorIdFrom, uint256 validatorIdTo) returns()
func (_StakeManager *StakeManagerSession) MigrateValidatorsData(validatorIdFrom *big.Int, validatorIdTo *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.MigrateValidatorsData(&_StakeManager.TransactOpts, validatorIdFrom, validatorIdTo)
}

// MigrateValidatorsData is a paid mutator transaction binding the contract method 0x9ddbbf85.
//
// Solidity: function migrateValidatorsData(uint256 validatorIdFrom, uint256 validatorIdTo) returns()
func (_StakeManager *StakeManagerTransactorSession) MigrateValidatorsData(validatorIdFrom *big.Int, validatorIdTo *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.MigrateValidatorsData(&_StakeManager.TransactOpts, validatorIdFrom, validatorIdTo)
}

// Reinitialize is a paid mutator transaction binding the contract method 0x078a13b1.
//
// Solidity: function reinitialize(address _NFTContract, address _stakingLogger, address _validatorShareFactory, address _extensionCode) returns()
func (_StakeManager *StakeManagerTransactor) Reinitialize(opts *bind.TransactOpts, _NFTContract common.Address, _stakingLogger common.Address, _validatorShareFactory common.Address, _extensionCode common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "reinitialize", _NFTContract, _stakingLogger, _validatorShareFactory, _extensionCode)
}

// Reinitialize is a paid mutator transaction binding the contract method 0x078a13b1.
//
// Solidity: function reinitialize(address _NFTContract, address _stakingLogger, address _validatorShareFactory, address _extensionCode) returns()
func (_StakeManager *StakeManagerSession) Reinitialize(_NFTContract common.Address, _stakingLogger common.Address, _validatorShareFactory common.Address, _extensionCode common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.Reinitialize(&_StakeManager.TransactOpts, _NFTContract, _stakingLogger, _validatorShareFactory, _extensionCode)
}

// Reinitialize is a paid mutator transaction binding the contract method 0x078a13b1.
//
// Solidity: function reinitialize(address _NFTContract, address _stakingLogger, address _validatorShareFactory, address _extensionCode) returns()
func (_StakeManager *StakeManagerTransactorSession) Reinitialize(_NFTContract common.Address, _stakingLogger common.Address, _validatorShareFactory common.Address, _extensionCode common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.Reinitialize(&_StakeManager.TransactOpts, _NFTContract, _stakingLogger, _validatorShareFactory, _extensionCode)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakeManager *StakeManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakeManager *StakeManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakeManager.Contract.RenounceOwnership(&_StakeManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakeManager *StakeManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakeManager.Contract.RenounceOwnership(&_StakeManager.TransactOpts)
}

// Restake is a paid mutator transaction binding the contract method 0x28cc4e41.
//
// Solidity: function restake(uint256 validatorId, uint256 amount, bool stakeRewards) returns()
func (_StakeManager *StakeManagerTransactor) Restake(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int, stakeRewards bool) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "restake", validatorId, amount, stakeRewards)
}

// Restake is a paid mutator transaction binding the contract method 0x28cc4e41.
//
// Solidity: function restake(uint256 validatorId, uint256 amount, bool stakeRewards) returns()
func (_StakeManager *StakeManagerSession) Restake(validatorId *big.Int, amount *big.Int, stakeRewards bool) (*types.Transaction, error) {
	return _StakeManager.Contract.Restake(&_StakeManager.TransactOpts, validatorId, amount, stakeRewards)
}

// Restake is a paid mutator transaction binding the contract method 0x28cc4e41.
//
// Solidity: function restake(uint256 validatorId, uint256 amount, bool stakeRewards) returns()
func (_StakeManager *StakeManagerTransactorSession) Restake(validatorId *big.Int, amount *big.Int, stakeRewards bool) (*types.Transaction, error) {
	return _StakeManager.Contract.Restake(&_StakeManager.TransactOpts, validatorId, amount, stakeRewards)
}

// RestakePOL is a paid mutator transaction binding the contract method 0x96a93224.
//
// Solidity: function restakePOL(uint256 validatorId, uint256 amount, bool stakeRewards) returns()
func (_StakeManager *StakeManagerTransactor) RestakePOL(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int, stakeRewards bool) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "restakePOL", validatorId, amount, stakeRewards)
}

// RestakePOL is a paid mutator transaction binding the contract method 0x96a93224.
//
// Solidity: function restakePOL(uint256 validatorId, uint256 amount, bool stakeRewards) returns()
func (_StakeManager *StakeManagerSession) RestakePOL(validatorId *big.Int, amount *big.Int, stakeRewards bool) (*types.Transaction, error) {
	return _StakeManager.Contract.RestakePOL(&_StakeManager.TransactOpts, validatorId, amount, stakeRewards)
}

// RestakePOL is a paid mutator transaction binding the contract method 0x96a93224.
//
// Solidity: function restakePOL(uint256 validatorId, uint256 amount, bool stakeRewards) returns()
func (_StakeManager *StakeManagerTransactorSession) RestakePOL(validatorId *big.Int, amount *big.Int, stakeRewards bool) (*types.Transaction, error) {
	return _StakeManager.Contract.RestakePOL(&_StakeManager.TransactOpts, validatorId, amount, stakeRewards)
}

// SetCurrentEpoch is a paid mutator transaction binding the contract method 0x1dd6b9b1.
//
// Solidity: function setCurrentEpoch(uint256 _currentEpoch) returns()
func (_StakeManager *StakeManagerTransactor) SetCurrentEpoch(opts *bind.TransactOpts, _currentEpoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "setCurrentEpoch", _currentEpoch)
}

// SetCurrentEpoch is a paid mutator transaction binding the contract method 0x1dd6b9b1.
//
// Solidity: function setCurrentEpoch(uint256 _currentEpoch) returns()
func (_StakeManager *StakeManagerSession) SetCurrentEpoch(_currentEpoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetCurrentEpoch(&_StakeManager.TransactOpts, _currentEpoch)
}

// SetCurrentEpoch is a paid mutator transaction binding the contract method 0x1dd6b9b1.
//
// Solidity: function setCurrentEpoch(uint256 _currentEpoch) returns()
func (_StakeManager *StakeManagerTransactorSession) SetCurrentEpoch(_currentEpoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetCurrentEpoch(&_StakeManager.TransactOpts, _currentEpoch)
}

// SetDelegationEnabled is a paid mutator transaction binding the contract method 0xf28699fa.
//
// Solidity: function setDelegationEnabled(bool enabled) returns()
func (_StakeManager *StakeManagerTransactor) SetDelegationEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "setDelegationEnabled", enabled)
}

// SetDelegationEnabled is a paid mutator transaction binding the contract method 0xf28699fa.
//
// Solidity: function setDelegationEnabled(bool enabled) returns()
func (_StakeManager *StakeManagerSession) SetDelegationEnabled(enabled bool) (*types.Transaction, error) {
	return _StakeManager.Contract.SetDelegationEnabled(&_StakeManager.TransactOpts, enabled)
}

// SetDelegationEnabled is a paid mutator transaction binding the contract method 0xf28699fa.
//
// Solidity: function setDelegationEnabled(bool enabled) returns()
func (_StakeManager *StakeManagerTransactorSession) SetDelegationEnabled(enabled bool) (*types.Transaction, error) {
	return _StakeManager.Contract.SetDelegationEnabled(&_StakeManager.TransactOpts, enabled)
}

// SetStakingToken is a paid mutator transaction binding the contract method 0x1e9b12ef.
//
// Solidity: function setStakingToken(address _token) returns()
func (_StakeManager *StakeManagerTransactor) SetStakingToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "setStakingToken", _token)
}

// SetStakingToken is a paid mutator transaction binding the contract method 0x1e9b12ef.
//
// Solidity: function setStakingToken(address _token) returns()
func (_StakeManager *StakeManagerSession) SetStakingToken(_token common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.SetStakingToken(&_StakeManager.TransactOpts, _token)
}

// SetStakingToken is a paid mutator transaction binding the contract method 0x1e9b12ef.
//
// Solidity: function setStakingToken(address _token) returns()
func (_StakeManager *StakeManagerTransactorSession) SetStakingToken(_token common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.SetStakingToken(&_StakeManager.TransactOpts, _token)
}

// Slash is a paid mutator transaction binding the contract method 0x5e47655f.
//
// Solidity: function slash(bytes _slashingInfoList) returns(uint256)
func (_StakeManager *StakeManagerTransactor) Slash(opts *bind.TransactOpts, _slashingInfoList []byte) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "slash", _slashingInfoList)
}

// Slash is a paid mutator transaction binding the contract method 0x5e47655f.
//
// Solidity: function slash(bytes _slashingInfoList) returns(uint256)
func (_StakeManager *StakeManagerSession) Slash(_slashingInfoList []byte) (*types.Transaction, error) {
	return _StakeManager.Contract.Slash(&_StakeManager.TransactOpts, _slashingInfoList)
}

// Slash is a paid mutator transaction binding the contract method 0x5e47655f.
//
// Solidity: function slash(bytes _slashingInfoList) returns(uint256)
func (_StakeManager *StakeManagerTransactorSession) Slash(_slashingInfoList []byte) (*types.Transaction, error) {
	return _StakeManager.Contract.Slash(&_StakeManager.TransactOpts, _slashingInfoList)
}

// StakeFor is a paid mutator transaction binding the contract method 0x4fdd20f1.
//
// Solidity: function stakeFor(address user, uint256 amount, uint256 heimdallFee, bool acceptDelegation, bytes signerPubkey) returns()
func (_StakeManager *StakeManagerTransactor) StakeFor(opts *bind.TransactOpts, user common.Address, amount *big.Int, heimdallFee *big.Int, acceptDelegation bool, signerPubkey []byte) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "stakeFor", user, amount, heimdallFee, acceptDelegation, signerPubkey)
}

// StakeFor is a paid mutator transaction binding the contract method 0x4fdd20f1.
//
// Solidity: function stakeFor(address user, uint256 amount, uint256 heimdallFee, bool acceptDelegation, bytes signerPubkey) returns()
func (_StakeManager *StakeManagerSession) StakeFor(user common.Address, amount *big.Int, heimdallFee *big.Int, acceptDelegation bool, signerPubkey []byte) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeFor(&_StakeManager.TransactOpts, user, amount, heimdallFee, acceptDelegation, signerPubkey)
}

// StakeFor is a paid mutator transaction binding the contract method 0x4fdd20f1.
//
// Solidity: function stakeFor(address user, uint256 amount, uint256 heimdallFee, bool acceptDelegation, bytes signerPubkey) returns()
func (_StakeManager *StakeManagerTransactorSession) StakeFor(user common.Address, amount *big.Int, heimdallFee *big.Int, acceptDelegation bool, signerPubkey []byte) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeFor(&_StakeManager.TransactOpts, user, amount, heimdallFee, acceptDelegation, signerPubkey)
}

// StakeForPOL is a paid mutator transaction binding the contract method 0x81254e12.
//
// Solidity: function stakeForPOL(address user, uint256 amount, uint256 heimdallFee, bool acceptDelegation, bytes signerPubkey) returns()
func (_StakeManager *StakeManagerTransactor) StakeForPOL(opts *bind.TransactOpts, user common.Address, amount *big.Int, heimdallFee *big.Int, acceptDelegation bool, signerPubkey []byte) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "stakeForPOL", user, amount, heimdallFee, acceptDelegation, signerPubkey)
}

// StakeForPOL is a paid mutator transaction binding the contract method 0x81254e12.
//
// Solidity: function stakeForPOL(address user, uint256 amount, uint256 heimdallFee, bool acceptDelegation, bytes signerPubkey) returns()
func (_StakeManager *StakeManagerSession) StakeForPOL(user common.Address, amount *big.Int, heimdallFee *big.Int, acceptDelegation bool, signerPubkey []byte) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeForPOL(&_StakeManager.TransactOpts, user, amount, heimdallFee, acceptDelegation, signerPubkey)
}

// StakeForPOL is a paid mutator transaction binding the contract method 0x81254e12.
//
// Solidity: function stakeForPOL(address user, uint256 amount, uint256 heimdallFee, bool acceptDelegation, bytes signerPubkey) returns()
func (_StakeManager *StakeManagerTransactorSession) StakeForPOL(user common.Address, amount *big.Int, heimdallFee *big.Int, acceptDelegation bool, signerPubkey []byte) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeForPOL(&_StakeManager.TransactOpts, user, amount, heimdallFee, acceptDelegation, signerPubkey)
}

// StartAuction is a paid mutator transaction binding the contract method 0xa6854877.
//
// Solidity: function startAuction(uint256 validatorId, uint256 amount, bool _acceptDelegation, bytes _signerPubkey) returns()
func (_StakeManager *StakeManagerTransactor) StartAuction(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int, _acceptDelegation bool, _signerPubkey []byte) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "startAuction", validatorId, amount, _acceptDelegation, _signerPubkey)
}

// StartAuction is a paid mutator transaction binding the contract method 0xa6854877.
//
// Solidity: function startAuction(uint256 validatorId, uint256 amount, bool _acceptDelegation, bytes _signerPubkey) returns()
func (_StakeManager *StakeManagerSession) StartAuction(validatorId *big.Int, amount *big.Int, _acceptDelegation bool, _signerPubkey []byte) (*types.Transaction, error) {
	return _StakeManager.Contract.StartAuction(&_StakeManager.TransactOpts, validatorId, amount, _acceptDelegation, _signerPubkey)
}

// StartAuction is a paid mutator transaction binding the contract method 0xa6854877.
//
// Solidity: function startAuction(uint256 validatorId, uint256 amount, bool _acceptDelegation, bytes _signerPubkey) returns()
func (_StakeManager *StakeManagerTransactorSession) StartAuction(validatorId *big.Int, amount *big.Int, _acceptDelegation bool, _signerPubkey []byte) (*types.Transaction, error) {
	return _StakeManager.Contract.StartAuction(&_StakeManager.TransactOpts, validatorId, amount, _acceptDelegation, _signerPubkey)
}

// StopAuctions is a paid mutator transaction binding the contract method 0xf771fc87.
//
// Solidity: function stopAuctions(uint256 forNCheckpoints) returns()
func (_StakeManager *StakeManagerTransactor) StopAuctions(opts *bind.TransactOpts, forNCheckpoints *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "stopAuctions", forNCheckpoints)
}

// StopAuctions is a paid mutator transaction binding the contract method 0xf771fc87.
//
// Solidity: function stopAuctions(uint256 forNCheckpoints) returns()
func (_StakeManager *StakeManagerSession) StopAuctions(forNCheckpoints *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.StopAuctions(&_StakeManager.TransactOpts, forNCheckpoints)
}

// StopAuctions is a paid mutator transaction binding the contract method 0xf771fc87.
//
// Solidity: function stopAuctions(uint256 forNCheckpoints) returns()
func (_StakeManager *StakeManagerTransactorSession) StopAuctions(forNCheckpoints *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.StopAuctions(&_StakeManager.TransactOpts, forNCheckpoints)
}

// TopUpForFee is a paid mutator transaction binding the contract method 0x63656798.
//
// Solidity: function topUpForFee(address user, uint256 heimdallFee) returns()
func (_StakeManager *StakeManagerTransactor) TopUpForFee(opts *bind.TransactOpts, user common.Address, heimdallFee *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "topUpForFee", user, heimdallFee)
}

// TopUpForFee is a paid mutator transaction binding the contract method 0x63656798.
//
// Solidity: function topUpForFee(address user, uint256 heimdallFee) returns()
func (_StakeManager *StakeManagerSession) TopUpForFee(user common.Address, heimdallFee *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.TopUpForFee(&_StakeManager.TransactOpts, user, heimdallFee)
}

// TopUpForFee is a paid mutator transaction binding the contract method 0x63656798.
//
// Solidity: function topUpForFee(address user, uint256 heimdallFee) returns()
func (_StakeManager *StakeManagerTransactorSession) TopUpForFee(user common.Address, heimdallFee *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.TopUpForFee(&_StakeManager.TransactOpts, user, heimdallFee)
}

// TransferFunds is a paid mutator transaction binding the contract method 0xbc8756a9.
//
// Solidity: function transferFunds(uint256 validatorId, uint256 amount, address delegator) returns(bool)
func (_StakeManager *StakeManagerTransactor) TransferFunds(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "transferFunds", validatorId, amount, delegator)
}

// TransferFunds is a paid mutator transaction binding the contract method 0xbc8756a9.
//
// Solidity: function transferFunds(uint256 validatorId, uint256 amount, address delegator) returns(bool)
func (_StakeManager *StakeManagerSession) TransferFunds(validatorId *big.Int, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.TransferFunds(&_StakeManager.TransactOpts, validatorId, amount, delegator)
}

// TransferFunds is a paid mutator transaction binding the contract method 0xbc8756a9.
//
// Solidity: function transferFunds(uint256 validatorId, uint256 amount, address delegator) returns(bool)
func (_StakeManager *StakeManagerTransactorSession) TransferFunds(validatorId *big.Int, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.TransferFunds(&_StakeManager.TransactOpts, validatorId, amount, delegator)
}

// TransferFundsPOL is a paid mutator transaction binding the contract method 0xc7f067cb.
//
// Solidity: function transferFundsPOL(uint256 validatorId, uint256 amount, address delegator) returns(bool)
func (_StakeManager *StakeManagerTransactor) TransferFundsPOL(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "transferFundsPOL", validatorId, amount, delegator)
}

// TransferFundsPOL is a paid mutator transaction binding the contract method 0xc7f067cb.
//
// Solidity: function transferFundsPOL(uint256 validatorId, uint256 amount, address delegator) returns(bool)
func (_StakeManager *StakeManagerSession) TransferFundsPOL(validatorId *big.Int, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.TransferFundsPOL(&_StakeManager.TransactOpts, validatorId, amount, delegator)
}

// TransferFundsPOL is a paid mutator transaction binding the contract method 0xc7f067cb.
//
// Solidity: function transferFundsPOL(uint256 validatorId, uint256 amount, address delegator) returns(bool)
func (_StakeManager *StakeManagerTransactorSession) TransferFundsPOL(validatorId *big.Int, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.TransferFundsPOL(&_StakeManager.TransactOpts, validatorId, amount, delegator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakeManager *StakeManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakeManager *StakeManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.TransferOwnership(&_StakeManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakeManager *StakeManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.TransferOwnership(&_StakeManager.TransactOpts, newOwner)
}

// Unjail is a paid mutator transaction binding the contract method 0x178c2c83.
//
// Solidity: function unjail(uint256 validatorId) returns()
func (_StakeManager *StakeManagerTransactor) Unjail(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "unjail", validatorId)
}

// Unjail is a paid mutator transaction binding the contract method 0x178c2c83.
//
// Solidity: function unjail(uint256 validatorId) returns()
func (_StakeManager *StakeManagerSession) Unjail(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Unjail(&_StakeManager.TransactOpts, validatorId)
}

// Unjail is a paid mutator transaction binding the contract method 0x178c2c83.
//
// Solidity: function unjail(uint256 validatorId) returns()
func (_StakeManager *StakeManagerTransactorSession) Unjail(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Unjail(&_StakeManager.TransactOpts, validatorId)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_StakeManager *StakeManagerTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_StakeManager *StakeManagerSession) Unlock() (*types.Transaction, error) {
	return _StakeManager.Contract.Unlock(&_StakeManager.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_StakeManager *StakeManagerTransactorSession) Unlock() (*types.Transaction, error) {
	return _StakeManager.Contract.Unlock(&_StakeManager.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 validatorId) returns()
func (_StakeManager *StakeManagerTransactor) Unstake(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "unstake", validatorId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 validatorId) returns()
func (_StakeManager *StakeManagerSession) Unstake(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Unstake(&_StakeManager.TransactOpts, validatorId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 validatorId) returns()
func (_StakeManager *StakeManagerTransactorSession) Unstake(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Unstake(&_StakeManager.TransactOpts, validatorId)
}

// UnstakeClaim is a paid mutator transaction binding the contract method 0xd86d53e7.
//
// Solidity: function unstakeClaim(uint256 validatorId) returns()
func (_StakeManager *StakeManagerTransactor) UnstakeClaim(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "unstakeClaim", validatorId)
}

// UnstakeClaim is a paid mutator transaction binding the contract method 0xd86d53e7.
//
// Solidity: function unstakeClaim(uint256 validatorId) returns()
func (_StakeManager *StakeManagerSession) UnstakeClaim(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UnstakeClaim(&_StakeManager.TransactOpts, validatorId)
}

// UnstakeClaim is a paid mutator transaction binding the contract method 0xd86d53e7.
//
// Solidity: function unstakeClaim(uint256 validatorId) returns()
func (_StakeManager *StakeManagerTransactorSession) UnstakeClaim(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UnstakeClaim(&_StakeManager.TransactOpts, validatorId)
}

// UnstakeClaimPOL is a paid mutator transaction binding the contract method 0xff3db8e3.
//
// Solidity: function unstakeClaimPOL(uint256 validatorId) returns()
func (_StakeManager *StakeManagerTransactor) UnstakeClaimPOL(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "unstakeClaimPOL", validatorId)
}

// UnstakeClaimPOL is a paid mutator transaction binding the contract method 0xff3db8e3.
//
// Solidity: function unstakeClaimPOL(uint256 validatorId) returns()
func (_StakeManager *StakeManagerSession) UnstakeClaimPOL(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UnstakeClaimPOL(&_StakeManager.TransactOpts, validatorId)
}

// UnstakeClaimPOL is a paid mutator transaction binding the contract method 0xff3db8e3.
//
// Solidity: function unstakeClaimPOL(uint256 validatorId) returns()
func (_StakeManager *StakeManagerTransactorSession) UnstakeClaimPOL(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UnstakeClaimPOL(&_StakeManager.TransactOpts, validatorId)
}

// UnstakePOL is a paid mutator transaction binding the contract method 0x30cbe571.
//
// Solidity: function unstakePOL(uint256 validatorId) returns()
func (_StakeManager *StakeManagerTransactor) UnstakePOL(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "unstakePOL", validatorId)
}

// UnstakePOL is a paid mutator transaction binding the contract method 0x30cbe571.
//
// Solidity: function unstakePOL(uint256 validatorId) returns()
func (_StakeManager *StakeManagerSession) UnstakePOL(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UnstakePOL(&_StakeManager.TransactOpts, validatorId)
}

// UnstakePOL is a paid mutator transaction binding the contract method 0x30cbe571.
//
// Solidity: function unstakePOL(uint256 validatorId) returns()
func (_StakeManager *StakeManagerTransactorSession) UnstakePOL(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UnstakePOL(&_StakeManager.TransactOpts, validatorId)
}

// UpdateCheckPointBlockInterval is a paid mutator transaction binding the contract method 0xa440ab1e.
//
// Solidity: function updateCheckPointBlockInterval(uint256 _blocks) returns()
func (_StakeManager *StakeManagerTransactor) UpdateCheckPointBlockInterval(opts *bind.TransactOpts, _blocks *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "updateCheckPointBlockInterval", _blocks)
}

// UpdateCheckPointBlockInterval is a paid mutator transaction binding the contract method 0xa440ab1e.
//
// Solidity: function updateCheckPointBlockInterval(uint256 _blocks) returns()
func (_StakeManager *StakeManagerSession) UpdateCheckPointBlockInterval(_blocks *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateCheckPointBlockInterval(&_StakeManager.TransactOpts, _blocks)
}

// UpdateCheckPointBlockInterval is a paid mutator transaction binding the contract method 0xa440ab1e.
//
// Solidity: function updateCheckPointBlockInterval(uint256 _blocks) returns()
func (_StakeManager *StakeManagerTransactorSession) UpdateCheckPointBlockInterval(_blocks *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateCheckPointBlockInterval(&_StakeManager.TransactOpts, _blocks)
}

// UpdateCheckpointReward is a paid mutator transaction binding the contract method 0xcbf383d5.
//
// Solidity: function updateCheckpointReward(uint256 newReward) returns()
func (_StakeManager *StakeManagerTransactor) UpdateCheckpointReward(opts *bind.TransactOpts, newReward *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "updateCheckpointReward", newReward)
}

// UpdateCheckpointReward is a paid mutator transaction binding the contract method 0xcbf383d5.
//
// Solidity: function updateCheckpointReward(uint256 newReward) returns()
func (_StakeManager *StakeManagerSession) UpdateCheckpointReward(newReward *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateCheckpointReward(&_StakeManager.TransactOpts, newReward)
}

// UpdateCheckpointReward is a paid mutator transaction binding the contract method 0xcbf383d5.
//
// Solidity: function updateCheckpointReward(uint256 newReward) returns()
func (_StakeManager *StakeManagerTransactorSession) UpdateCheckpointReward(newReward *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateCheckpointReward(&_StakeManager.TransactOpts, newReward)
}

// UpdateCheckpointRewardParams is a paid mutator transaction binding the contract method 0x60c8d122.
//
// Solidity: function updateCheckpointRewardParams(uint256 _rewardDecreasePerCheckpoint, uint256 _maxRewardedCheckpoints, uint256 _checkpointRewardDelta) returns()
func (_StakeManager *StakeManagerTransactor) UpdateCheckpointRewardParams(opts *bind.TransactOpts, _rewardDecreasePerCheckpoint *big.Int, _maxRewardedCheckpoints *big.Int, _checkpointRewardDelta *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "updateCheckpointRewardParams", _rewardDecreasePerCheckpoint, _maxRewardedCheckpoints, _checkpointRewardDelta)
}

// UpdateCheckpointRewardParams is a paid mutator transaction binding the contract method 0x60c8d122.
//
// Solidity: function updateCheckpointRewardParams(uint256 _rewardDecreasePerCheckpoint, uint256 _maxRewardedCheckpoints, uint256 _checkpointRewardDelta) returns()
func (_StakeManager *StakeManagerSession) UpdateCheckpointRewardParams(_rewardDecreasePerCheckpoint *big.Int, _maxRewardedCheckpoints *big.Int, _checkpointRewardDelta *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateCheckpointRewardParams(&_StakeManager.TransactOpts, _rewardDecreasePerCheckpoint, _maxRewardedCheckpoints, _checkpointRewardDelta)
}

// UpdateCheckpointRewardParams is a paid mutator transaction binding the contract method 0x60c8d122.
//
// Solidity: function updateCheckpointRewardParams(uint256 _rewardDecreasePerCheckpoint, uint256 _maxRewardedCheckpoints, uint256 _checkpointRewardDelta) returns()
func (_StakeManager *StakeManagerTransactorSession) UpdateCheckpointRewardParams(_rewardDecreasePerCheckpoint *big.Int, _maxRewardedCheckpoints *big.Int, _checkpointRewardDelta *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateCheckpointRewardParams(&_StakeManager.TransactOpts, _rewardDecreasePerCheckpoint, _maxRewardedCheckpoints, _checkpointRewardDelta)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0xdcd962b2.
//
// Solidity: function updateCommissionRate(uint256 validatorId, uint256 newCommissionRate) returns()
func (_StakeManager *StakeManagerTransactor) UpdateCommissionRate(opts *bind.TransactOpts, validatorId *big.Int, newCommissionRate *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "updateCommissionRate", validatorId, newCommissionRate)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0xdcd962b2.
//
// Solidity: function updateCommissionRate(uint256 validatorId, uint256 newCommissionRate) returns()
func (_StakeManager *StakeManagerSession) UpdateCommissionRate(validatorId *big.Int, newCommissionRate *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateCommissionRate(&_StakeManager.TransactOpts, validatorId, newCommissionRate)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0xdcd962b2.
//
// Solidity: function updateCommissionRate(uint256 validatorId, uint256 newCommissionRate) returns()
func (_StakeManager *StakeManagerTransactorSession) UpdateCommissionRate(validatorId *big.Int, newCommissionRate *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateCommissionRate(&_StakeManager.TransactOpts, validatorId, newCommissionRate)
}

// UpdateDynastyValue is a paid mutator transaction binding the contract method 0xe6692f49.
//
// Solidity: function updateDynastyValue(uint256 newDynasty) returns()
func (_StakeManager *StakeManagerTransactor) UpdateDynastyValue(opts *bind.TransactOpts, newDynasty *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "updateDynastyValue", newDynasty)
}

// UpdateDynastyValue is a paid mutator transaction binding the contract method 0xe6692f49.
//
// Solidity: function updateDynastyValue(uint256 newDynasty) returns()
func (_StakeManager *StakeManagerSession) UpdateDynastyValue(newDynasty *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateDynastyValue(&_StakeManager.TransactOpts, newDynasty)
}

// UpdateDynastyValue is a paid mutator transaction binding the contract method 0xe6692f49.
//
// Solidity: function updateDynastyValue(uint256 newDynasty) returns()
func (_StakeManager *StakeManagerTransactorSession) UpdateDynastyValue(newDynasty *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateDynastyValue(&_StakeManager.TransactOpts, newDynasty)
}

// UpdateMinAmounts is a paid mutator transaction binding the contract method 0xb1d23f02.
//
// Solidity: function updateMinAmounts(uint256 _minDeposit, uint256 _minHeimdallFee) returns()
func (_StakeManager *StakeManagerTransactor) UpdateMinAmounts(opts *bind.TransactOpts, _minDeposit *big.Int, _minHeimdallFee *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "updateMinAmounts", _minDeposit, _minHeimdallFee)
}

// UpdateMinAmounts is a paid mutator transaction binding the contract method 0xb1d23f02.
//
// Solidity: function updateMinAmounts(uint256 _minDeposit, uint256 _minHeimdallFee) returns()
func (_StakeManager *StakeManagerSession) UpdateMinAmounts(_minDeposit *big.Int, _minHeimdallFee *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateMinAmounts(&_StakeManager.TransactOpts, _minDeposit, _minHeimdallFee)
}

// UpdateMinAmounts is a paid mutator transaction binding the contract method 0xb1d23f02.
//
// Solidity: function updateMinAmounts(uint256 _minDeposit, uint256 _minHeimdallFee) returns()
func (_StakeManager *StakeManagerTransactorSession) UpdateMinAmounts(_minDeposit *big.Int, _minHeimdallFee *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateMinAmounts(&_StakeManager.TransactOpts, _minDeposit, _minHeimdallFee)
}

// UpdateProposerBonus is a paid mutator transaction binding the contract method 0x9b33f434.
//
// Solidity: function updateProposerBonus(uint256 newProposerBonus) returns()
func (_StakeManager *StakeManagerTransactor) UpdateProposerBonus(opts *bind.TransactOpts, newProposerBonus *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "updateProposerBonus", newProposerBonus)
}

// UpdateProposerBonus is a paid mutator transaction binding the contract method 0x9b33f434.
//
// Solidity: function updateProposerBonus(uint256 newProposerBonus) returns()
func (_StakeManager *StakeManagerSession) UpdateProposerBonus(newProposerBonus *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateProposerBonus(&_StakeManager.TransactOpts, newProposerBonus)
}

// UpdateProposerBonus is a paid mutator transaction binding the contract method 0x9b33f434.
//
// Solidity: function updateProposerBonus(uint256 newProposerBonus) returns()
func (_StakeManager *StakeManagerTransactorSession) UpdateProposerBonus(newProposerBonus *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateProposerBonus(&_StakeManager.TransactOpts, newProposerBonus)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xf41a9642.
//
// Solidity: function updateSigner(uint256 validatorId, bytes signerPubkey) returns()
func (_StakeManager *StakeManagerTransactor) UpdateSigner(opts *bind.TransactOpts, validatorId *big.Int, signerPubkey []byte) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "updateSigner", validatorId, signerPubkey)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xf41a9642.
//
// Solidity: function updateSigner(uint256 validatorId, bytes signerPubkey) returns()
func (_StakeManager *StakeManagerSession) UpdateSigner(validatorId *big.Int, signerPubkey []byte) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateSigner(&_StakeManager.TransactOpts, validatorId, signerPubkey)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xf41a9642.
//
// Solidity: function updateSigner(uint256 validatorId, bytes signerPubkey) returns()
func (_StakeManager *StakeManagerTransactorSession) UpdateSigner(validatorId *big.Int, signerPubkey []byte) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateSigner(&_StakeManager.TransactOpts, validatorId, signerPubkey)
}

// UpdateSignerUpdateLimit is a paid mutator transaction binding the contract method 0x06cfb104.
//
// Solidity: function updateSignerUpdateLimit(uint256 _limit) returns()
func (_StakeManager *StakeManagerTransactor) UpdateSignerUpdateLimit(opts *bind.TransactOpts, _limit *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "updateSignerUpdateLimit", _limit)
}

// UpdateSignerUpdateLimit is a paid mutator transaction binding the contract method 0x06cfb104.
//
// Solidity: function updateSignerUpdateLimit(uint256 _limit) returns()
func (_StakeManager *StakeManagerSession) UpdateSignerUpdateLimit(_limit *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateSignerUpdateLimit(&_StakeManager.TransactOpts, _limit)
}

// UpdateSignerUpdateLimit is a paid mutator transaction binding the contract method 0x06cfb104.
//
// Solidity: function updateSignerUpdateLimit(uint256 _limit) returns()
func (_StakeManager *StakeManagerTransactorSession) UpdateSignerUpdateLimit(_limit *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateSignerUpdateLimit(&_StakeManager.TransactOpts, _limit)
}

// UpdateValidatorContractAddress is a paid mutator transaction binding the contract method 0xc710e922.
//
// Solidity: function updateValidatorContractAddress(uint256 validatorId, address newContractAddress) returns()
func (_StakeManager *StakeManagerTransactor) UpdateValidatorContractAddress(opts *bind.TransactOpts, validatorId *big.Int, newContractAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "updateValidatorContractAddress", validatorId, newContractAddress)
}

// UpdateValidatorContractAddress is a paid mutator transaction binding the contract method 0xc710e922.
//
// Solidity: function updateValidatorContractAddress(uint256 validatorId, address newContractAddress) returns()
func (_StakeManager *StakeManagerSession) UpdateValidatorContractAddress(validatorId *big.Int, newContractAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateValidatorContractAddress(&_StakeManager.TransactOpts, validatorId, newContractAddress)
}

// UpdateValidatorContractAddress is a paid mutator transaction binding the contract method 0xc710e922.
//
// Solidity: function updateValidatorContractAddress(uint256 validatorId, address newContractAddress) returns()
func (_StakeManager *StakeManagerTransactorSession) UpdateValidatorContractAddress(validatorId *big.Int, newContractAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateValidatorContractAddress(&_StakeManager.TransactOpts, validatorId, newContractAddress)
}

// UpdateValidatorDelegation is a paid mutator transaction binding the contract method 0xd6de07d0.
//
// Solidity: function updateValidatorDelegation(bool delegation) returns()
func (_StakeManager *StakeManagerTransactor) UpdateValidatorDelegation(opts *bind.TransactOpts, delegation bool) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "updateValidatorDelegation", delegation)
}

// UpdateValidatorDelegation is a paid mutator transaction binding the contract method 0xd6de07d0.
//
// Solidity: function updateValidatorDelegation(bool delegation) returns()
func (_StakeManager *StakeManagerSession) UpdateValidatorDelegation(delegation bool) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateValidatorDelegation(&_StakeManager.TransactOpts, delegation)
}

// UpdateValidatorDelegation is a paid mutator transaction binding the contract method 0xd6de07d0.
//
// Solidity: function updateValidatorDelegation(bool delegation) returns()
func (_StakeManager *StakeManagerTransactorSession) UpdateValidatorDelegation(delegation bool) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateValidatorDelegation(&_StakeManager.TransactOpts, delegation)
}

// UpdateValidatorState is a paid mutator transaction binding the contract method 0x9ff11500.
//
// Solidity: function updateValidatorState(uint256 validatorId, int256 amount) returns()
func (_StakeManager *StakeManagerTransactor) UpdateValidatorState(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "updateValidatorState", validatorId, amount)
}

// UpdateValidatorState is a paid mutator transaction binding the contract method 0x9ff11500.
//
// Solidity: function updateValidatorState(uint256 validatorId, int256 amount) returns()
func (_StakeManager *StakeManagerSession) UpdateValidatorState(validatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateValidatorState(&_StakeManager.TransactOpts, validatorId, amount)
}

// UpdateValidatorState is a paid mutator transaction binding the contract method 0x9ff11500.
//
// Solidity: function updateValidatorState(uint256 validatorId, int256 amount) returns()
func (_StakeManager *StakeManagerTransactorSession) UpdateValidatorState(validatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateValidatorState(&_StakeManager.TransactOpts, validatorId, amount)
}

// UpdateValidatorThreshold is a paid mutator transaction binding the contract method 0x16827b1b.
//
// Solidity: function updateValidatorThreshold(uint256 newThreshold) returns()
func (_StakeManager *StakeManagerTransactor) UpdateValidatorThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "updateValidatorThreshold", newThreshold)
}

// UpdateValidatorThreshold is a paid mutator transaction binding the contract method 0x16827b1b.
//
// Solidity: function updateValidatorThreshold(uint256 newThreshold) returns()
func (_StakeManager *StakeManagerSession) UpdateValidatorThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateValidatorThreshold(&_StakeManager.TransactOpts, newThreshold)
}

// UpdateValidatorThreshold is a paid mutator transaction binding the contract method 0x16827b1b.
//
// Solidity: function updateValidatorThreshold(uint256 newThreshold) returns()
func (_StakeManager *StakeManagerTransactorSession) UpdateValidatorThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateValidatorThreshold(&_StakeManager.TransactOpts, newThreshold)
}

// WithdrawDelegatorsReward is a paid mutator transaction binding the contract method 0x7ed4b27c.
//
// Solidity: function withdrawDelegatorsReward(uint256 validatorId) returns(uint256)
func (_StakeManager *StakeManagerTransactor) WithdrawDelegatorsReward(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "withdrawDelegatorsReward", validatorId)
}

// WithdrawDelegatorsReward is a paid mutator transaction binding the contract method 0x7ed4b27c.
//
// Solidity: function withdrawDelegatorsReward(uint256 validatorId) returns(uint256)
func (_StakeManager *StakeManagerSession) WithdrawDelegatorsReward(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawDelegatorsReward(&_StakeManager.TransactOpts, validatorId)
}

// WithdrawDelegatorsReward is a paid mutator transaction binding the contract method 0x7ed4b27c.
//
// Solidity: function withdrawDelegatorsReward(uint256 validatorId) returns(uint256)
func (_StakeManager *StakeManagerTransactorSession) WithdrawDelegatorsReward(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawDelegatorsReward(&_StakeManager.TransactOpts, validatorId)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x9342c8f4.
//
// Solidity: function withdrawRewards(uint256 validatorId) returns()
func (_StakeManager *StakeManagerTransactor) WithdrawRewards(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "withdrawRewards", validatorId)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x9342c8f4.
//
// Solidity: function withdrawRewards(uint256 validatorId) returns()
func (_StakeManager *StakeManagerSession) WithdrawRewards(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawRewards(&_StakeManager.TransactOpts, validatorId)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x9342c8f4.
//
// Solidity: function withdrawRewards(uint256 validatorId) returns()
func (_StakeManager *StakeManagerTransactorSession) WithdrawRewards(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawRewards(&_StakeManager.TransactOpts, validatorId)
}

// WithdrawRewardsPOL is a paid mutator transaction binding the contract method 0x1c404cec.
//
// Solidity: function withdrawRewardsPOL(uint256 validatorId) returns()
func (_StakeManager *StakeManagerTransactor) WithdrawRewardsPOL(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "withdrawRewardsPOL", validatorId)
}

// WithdrawRewardsPOL is a paid mutator transaction binding the contract method 0x1c404cec.
//
// Solidity: function withdrawRewardsPOL(uint256 validatorId) returns()
func (_StakeManager *StakeManagerSession) WithdrawRewardsPOL(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawRewardsPOL(&_StakeManager.TransactOpts, validatorId)
}

// WithdrawRewardsPOL is a paid mutator transaction binding the contract method 0x1c404cec.
//
// Solidity: function withdrawRewardsPOL(uint256 validatorId) returns()
func (_StakeManager *StakeManagerTransactorSession) WithdrawRewardsPOL(validatorId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawRewardsPOL(&_StakeManager.TransactOpts, validatorId)
}

// StakeManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakeManager contract.
type StakeManagerOwnershipTransferredIterator struct {
	Event *StakeManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakeManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerOwnershipTransferred)
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
		it.Event = new(StakeManagerOwnershipTransferred)
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
func (it *StakeManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerOwnershipTransferred represents a OwnershipTransferred event raised by the StakeManager contract.
type StakeManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakeManager *StakeManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakeManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerOwnershipTransferredIterator{contract: _StakeManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakeManager *StakeManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakeManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerOwnershipTransferred)
				if err := _StakeManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakeManager *StakeManagerFilterer) ParseOwnershipTransferred(log types.Log) (*StakeManagerOwnershipTransferred, error) {
	event := new(StakeManagerOwnershipTransferred)
	if err := _StakeManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerRootChainChangedIterator is returned from FilterRootChainChanged and is used to iterate over the raw logs and unpacked data for RootChainChanged events raised by the StakeManager contract.
type StakeManagerRootChainChangedIterator struct {
	Event *StakeManagerRootChainChanged // Event containing the contract specifics and raw log

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
func (it *StakeManagerRootChainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerRootChainChanged)
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
		it.Event = new(StakeManagerRootChainChanged)
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
func (it *StakeManagerRootChainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerRootChainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerRootChainChanged represents a RootChainChanged event raised by the StakeManager contract.
type StakeManagerRootChainChanged struct {
	PreviousRootChain common.Address
	NewRootChain      common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRootChainChanged is a free log retrieval operation binding the contract event 0x211c9015fc81c0dbd45bd99f0f29fc1c143bfd53442d5ffd722bbbef7a887fe9.
//
// Solidity: event RootChainChanged(address indexed previousRootChain, address indexed newRootChain)
func (_StakeManager *StakeManagerFilterer) FilterRootChainChanged(opts *bind.FilterOpts, previousRootChain []common.Address, newRootChain []common.Address) (*StakeManagerRootChainChangedIterator, error) {

	var previousRootChainRule []interface{}
	for _, previousRootChainItem := range previousRootChain {
		previousRootChainRule = append(previousRootChainRule, previousRootChainItem)
	}
	var newRootChainRule []interface{}
	for _, newRootChainItem := range newRootChain {
		newRootChainRule = append(newRootChainRule, newRootChainItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "RootChainChanged", previousRootChainRule, newRootChainRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerRootChainChangedIterator{contract: _StakeManager.contract, event: "RootChainChanged", logs: logs, sub: sub}, nil
}

// WatchRootChainChanged is a free log subscription operation binding the contract event 0x211c9015fc81c0dbd45bd99f0f29fc1c143bfd53442d5ffd722bbbef7a887fe9.
//
// Solidity: event RootChainChanged(address indexed previousRootChain, address indexed newRootChain)
func (_StakeManager *StakeManagerFilterer) WatchRootChainChanged(opts *bind.WatchOpts, sink chan<- *StakeManagerRootChainChanged, previousRootChain []common.Address, newRootChain []common.Address) (event.Subscription, error) {

	var previousRootChainRule []interface{}
	for _, previousRootChainItem := range previousRootChain {
		previousRootChainRule = append(previousRootChainRule, previousRootChainItem)
	}
	var newRootChainRule []interface{}
	for _, newRootChainItem := range newRootChain {
		newRootChainRule = append(newRootChainRule, newRootChainItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "RootChainChanged", previousRootChainRule, newRootChainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerRootChainChanged)
				if err := _StakeManager.contract.UnpackLog(event, "RootChainChanged", log); err != nil {
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

// ParseRootChainChanged is a log parse operation binding the contract event 0x211c9015fc81c0dbd45bd99f0f29fc1c143bfd53442d5ffd722bbbef7a887fe9.
//
// Solidity: event RootChainChanged(address indexed previousRootChain, address indexed newRootChain)
func (_StakeManager *StakeManagerFilterer) ParseRootChainChanged(log types.Log) (*StakeManagerRootChainChanged, error) {
	event := new(StakeManagerRootChainChanged)
	if err := _StakeManager.contract.UnpackLog(event, "RootChainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

