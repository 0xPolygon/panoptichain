
## AccountBalancesObserver


### panoptichain_rpc_account_balance
The account balance (wei)

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- address
- token

## BaseFeePerGasObserver


### panoptichain_rpc_base_fee_per_gas
The base fee per gas (gwei)

Metric Type: GaugeVec

Variable Labels:
- network
- provider

## BlockObserver


### panoptichain_rpc_height
The latest known block height

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_rpc_block
The total number of blocks observed

Metric Type: CounterVec

Variable Labels:
- network
- provider

### panoptichain_rpc_difficulty
The difficulty of the block

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_rpc_block_size
The block size per block (bytes)

Metric Type: HistogramVec

Variable Labels:
- network
- provider

### panoptichain_rpc_extra_size
The size of the extra data (bytes)

Metric Type: HistogramVec

Variable Labels:
- network
- provider

## BlockIntervalObserver


### panoptichain_rpc_block_interval
the number of seconds between blocks

Metric Type: HistogramVec

Variable Labels:
- network
- provider

## BogonBlockObserver


### panoptichain_rpc_bogon_block
The total number of bogon blocks observed

Metric Type: CounterVec

Variable Labels:
- network
- provider

## BridgeEventObserver


### panoptichain_rpc_time_since_last_bridge_event
The time since the last zkEVM bridge event (in seconds)

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- origin_network
- destination_network

### panoptichain_rpc_bridge_event_deposit_count
The deposit count of the latest bridge event

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- origin_network
- destination_network

### panoptichain_rpc_bridge_event_amount
The amount in bridged (gwei)

Metric Type: HistogramVec

Variable Labels:
- network
- provider
- origin_network
- destination_network

## CheckpointObserver


### panoptichain_rpc_checkpoint_id
The last header block ID

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- finalized

### panoptichain_rpc_checkpoint_signatures
The number of validators that signed the latest checkpoint

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- finalized

### panoptichain_rpc_time_since_last_checkpoint
The elapsed time since the last checkpoint

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- finalized

### panoptichain_rpc_signed_checkpoint
Counts the number of times a validator has signed a checkpoint

Metric Type: CounterVec

Variable Labels:
- network
- provider
- signer

## ClaimEventObserver


### panoptichain_rpc_time_since_last_claim_event
The time since the last zkEVM claim event (in seconds)

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- origin_network

### panoptichain_rpc_claim_event_amount
The amount in claimed (gwei)

Metric Type: HistogramVec

Variable Labels:
- network
- provider
- origin_network

### panoptichain_rpc_observed_claim_events
The number of claim events observed

Metric Type: CounterVec

Variable Labels:
- network
- provider
- origin_network

## DepositCountObserver


### panoptichain_rpc_bridge_deposit_count
zkEVM bridge deposit count

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_rpc_bridge_last_updated_deposit_count
zkEVM bridge last updated deposit count

Metric Type: GaugeVec

Variable Labels:
- network
- provider

## DoubleSignObserver


### panoptichain_sensor_double_sign
Number of double sign events detected

Metric Type: CounterVec

Variable Labels:
- network
- provider
- signer_address

## EmptyBlockObserver


### panoptichain_rpc_empty_block
The total number of empty blocks observed

Metric Type: CounterVec

Variable Labels:
- network
- provider

## ExchangeRatesObserver


### panoptichain_exchange_rates
The exchange rate between the base and quote currencies

Metric Type: GaugeVec

Variable Labels:
- base
- quote

## ExitRootsObserver


### panoptichain_rpc_time_since_last_global_exit_root
The elapsed time since the last global exit root (in seconds)

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_rpc_time_since_last_mainnet_exit_root
The elapsed time since the last mainnet exit root (in seconds)

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_rpc_time_since_last_rollup_exit_root
The elapsed time since the last rollup exit root (in seconds)

Metric Type: GaugeVec

Variable Labels:
- network
- provider

## FinalizedHeightObserver


### panoptichain_rpc_finalized_height
The latest known finalized block height

Metric Type: GaugeVec

Variable Labels:
- network
- provider

## GasLimitObserver


### panoptichain_rpc_gas_limit
The gas limit of the block

Metric Type: GaugeVec

Variable Labels:
- network
- provider

## GasUsedObserver


### panoptichain_rpc_gas_used
The gas used in the block (million gas)

Metric Type: HistogramVec

Variable Labels:
- network
- provider

## HashDivergenceObserver


### panoptichain_rpc_hash_divergence
The number of blocks that have different hashes across different RPC providers

Metric Type: CounterVec

Variable Labels:
- network
- provider

## HeimdallBlockObserver


### panoptichain_heimdall_height
The block height for Heimdall

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_heimdall_transactions_per_block
The number of transactions per Heimdall block

Metric Type: HistogramVec

Variable Labels:
- network
- provider

### panoptichain_heimdall_total_transaction_count
The number of total transactions observed for Heimdall

Metric Type: CounterVec

Variable Labels:
- network
- provider

## HeimdallBlockIntervalObserver


### panoptichain_heimdall_block_interval
The time interval (in seconds) between Heimdall blocks

Metric Type: HistogramVec

Variable Labels:
- network
- provider

## HeimdallCheckpointObserver


### panoptichain_heimdall_checkpoint_start_block
The checkpoint start block

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_heimdall_checkpoint_end_block
The checkpoint end block

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_heimdall_checkpoint_id
The checkpoint id

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_heimdall_time_since_last_checkpoint
The time since last checkpoint

Metric Type: GaugeVec

Variable Labels:
- network
- provider

## HeimdallMissedBlockProposalObserver


### panoptichain_heimdall_missed_block_proposal
Missed block proposals

Metric Type: CounterVec

Variable Labels:
- network
- provider
- signer_address

## HeimdallMissedCheckpointProposalObserver


### panoptichain_heimdall_missed_checkpoint_proposal
Missed checkpoint proposals

Metric Type: CounterVec

Variable Labels:
- network
- provider
- signer_address

## HeimdallMissedMilestoneProposal


### panoptichain_heimdall_missed_milestone_proposal
Missed milestone proposals

Metric Type: CounterVec

Variable Labels:
- network
- provider
- signer_address

## HeimdallSignatureCountObserver


### panoptichain_heimdall_signatures
The number of signatures on block

Metric Type: GaugeVec

Variable Labels:
- network
- provider

## MilestoneObserver


### panoptichain_heimdall_time_since_last_milestone
The time since last milestone

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_heimdall_milestone_count
The milestone count

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_heimdall_milestone_start_block
The milestone start block

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_heimdall_milestone_end_block
The milestone end block

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_heimdall_milestone_observed
The number of milestones observed

Metric Type: CounterVec

Variable Labels:
- network
- provider

### panoptichain_heimdall_milestone_block_range
The number of blocks in the milestone

Metric Type: HistogramVec

Variable Labels:
- network
- provider

## MissedBlockProposalObserver


### panoptichain_rpc_missed_block_proposal
Missed block proposals

Metric Type: CounterVec

Variable Labels:
- network
- provider
- signer_address

## RefreshStateTimeObserver


### panoptichain_system_refresh_state_time
The amount of time it took to refresh the state in milliseconds

Metric Type: HistogramVec

Variable Labels:
- network
- provider

## ReorgObserver


### panoptichain_sensor_reorg_depth
The number of blocks that were reorganized

Metric Type: HistogramVec

Variable Labels:
- network
- provider

## RollupManagerObserver


### panoptichain_rpc_zkevm_last_batch_sequenced
The last batch sequenced number

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- rollup

### panoptichain_rpc_zkevm_time_since_last_sequenced
The time since the last sequenced batch (in seconds)

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- rollup

### panoptichain_rpc_zkevm_total_sequenced_batches
The total number of sequenced batches

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_rpc_zkevm_time_between_sequenced_batches
The time between sequenced batches (in seconds)

Metric Type: HistogramVec

Variable Labels:
- network
- provider
- rollup

### panoptichain_rpc_zkevm_sequenced_batches_tx_fee
The transaction fee of OnSequencedBatches events (in gwei)

Metric Type: HistogramVec

Variable Labels:
- network
- provider
- rollup
- address

### panoptichain_rpc_zkevm_observed_sequenced_batches
The number of sequenced batches observed

Metric Type: CounterVec

Variable Labels:
- network
- provider
- rollup
- address

### panoptichain_rpc_zkevm_last_verified_batch
The last verified batch number

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- rollup

### panoptichain_rpc_zkevm_time_since_last_verified
The time since the last verified batch (in seconds)

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- rollup
- pessimistic

### panoptichain_rpc_zkevm_total_verified_batches
The total number of verified batches

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_rpc_zkevm_time_between_verified_batches
The time between verified batches (in seconds)

Metric Type: HistogramVec

Variable Labels:
- network
- provider
- rollup

### panoptichain_rpc_zkevm_verified_batches_tx_fee
The transaction fee of verifyBatches and verifyBatchesTrustedAggregator events (in gwei)

Metric Type: HistogramVec

Variable Labels:
- network
- provider
- rollup
- address

### panoptichain_rpc_zkevm_observed_verified_batches
The number of verified batches observed

Metric Type: CounterVec

Variable Labels:
- network
- provider
- rollup
- address

### panoptichain_rpc_zkevm_last_force_batch
The last force batch number

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- rollup

### panoptichain_rpc_zkevm_last_force_batch_sequenced
The last force batch sequenced number

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- rollup

### panoptichain_rpc_zkevm_rollup_chain_id
The rollup chain ID

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- rollup

### panoptichain_rpc_zkevm_batch_fee
The batch fee (gwei)

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_rpc_zkevm_reward_per_batch
The reward per batch (gwei)

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_rpc_zkevm_trusted_sequencer_balance
The trusted sequencer balance (wei)

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- rollup
- token

### panoptichain_rpc_zkevm_aggregator_balance
The aggregator balance (wei)

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- address
- token

### panoptichain_rpc_zkevm_rollup_count
The number of rollups

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_rpc_zkevm_rollup_type_count
The number of rollup types

Metric Type: GaugeVec

Variable Labels:
- network
- provider

## SealedOutOfTurnObserver


### panoptichain_sensor_sealed_out_of_turn
The number of blocks that were sealed out of turn

Metric Type: CounterVec

Variable Labels:
- network
- provider
- signer_address

## BlockEventsObserver


### panoptichain_sensor_block_latency
The difference between the block time and the time seen for all blocks by a sensor (in milliseconds)

Metric Type: HistogramVec

Variable Labels:
- network
- provider
- sensor

### panoptichain_sensor_block_latency_diff
The difference between the first and last time a block was received by a sensor (in milliseconds)

Metric Type: HistogramVec

Variable Labels:
- network
- provider
- sensor

### panoptichain_sensor_first_block_latency
The difference between the block time and the time first seen by a sensor (in milliseconds)

Metric Type: HistogramVec

Variable Labels:
- network
- provider
- sensor

### panoptichain_sensor_last_block_latency
The difference between the block time and the time last seen by a sensor (in milliseconds)

Metric Type: HistogramVec

Variable Labels:
- network
- provider
- sensor

### panoptichain_sensor_peers
The number of unique block propagators observed by a sensor

Metric Type: HistogramVec

Variable Labels:
- network
- provider
- sensor

### panoptichain_sensor_block_events
The number of block events that occurred for a block by a sensor

Metric Type: HistogramVec

Variable Labels:
- network
- provider
- sensor

### panoptichain_sensor_connections
The number of sensors a peer propagated a block to

Metric Type: HistogramVec

Variable Labels:
- network
- provider

### panoptichain_sensor_latency_rank
The rank order in which a sensor received a block relative to other sensors

Metric Type: HistogramVec

Variable Labels:
- network
- provider
- sensor

## SensorBlocksObserver


### panoptichain_sensor_forks_per_block_number
The cardinality of forks for a block number

Metric Type: HistogramVec

Variable Labels:
- network
- provider

### panoptichain_sensor_total_blocks
The total number of blocks observed by sensors including bogons

Metric Type: CounterVec

Variable Labels:
- network
- provider

## SensorBogonBlockObserver


### panoptichain_sensor_bogon_block
The total number of bogon blocks observed

Metric Type: CounterVec

Variable Labels:
- network
- provider
- signer_address

## HeimdallSpanObserver


### panoptichain_heimdall_span_id
The span id

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_heimdall_span_start_block
The span start block

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_heimdall_span_end_block
The span end block

Metric Type: GaugeVec

Variable Labels:
- network
- provider

## StateSyncObserver


### panoptichain_rpc_state_sync_id
the latest observed state sync id

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- finalized

### panoptichain_rpc_time_since_last_state_sync
The elapsed time since the last state sync

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- finalized

## StolenBlockObserver


### panoptichain_sensor_stolen_block
The number blocks stolen from a validator

Metric Type: CounterVec

Variable Labels:
- network
- provider
- signer_address

## SystemObserver


### panoptichain_system_uptime
How long the system has been running in seconds

Metric Type: gauge

## TimeToFinalizedObserver


### panoptichain_rpc_time_to_finalized
The time difference between the latest block and the last finalized block (in seconds)

Metric Type: GaugeVec

Variable Labels:
- network
- provider

## TimeToMineObserver


### panoptichain_rpc_time_to_mine
Time it takes for sent transaction to be included in a block (in seconds)

Metric Type: HistogramVec

Variable Labels:
- network
- provider
- gas_price_factor

### panoptichain_rpc_time_to_mine_gas_price
The gas price for the time to mine transactions (in gwei)

Metric Type: HistogramVec

Variable Labels:
- network
- provider
- gas_price_factor

## TransactionCostObserver


### panoptichain_rpc_transaction_cost
The transaction cost (ether)

Metric Type: HistogramVec

Variable Labels:
- network
- provider

## TransactionCountObserver


### panoptichain_rpc_transactions_per_block
The number of transactions per block

Metric Type: HistogramVec

Variable Labels:
- network
- provider

## TransactionGasFeeCapObserver


### panoptichain_rpc_transaction_gas_fee_cap
The transaction gas fee cap (gwei)

Metric Type: HistogramVec

Variable Labels:
- network
- provider

## TransactionGasLimitObserver


### panoptichain_rpc_transaction_gas_limit
The transaction gas limit (gas)

Metric Type: HistogramVec

Variable Labels:
- network
- provider

## TransactionGasPriceObserver


### panoptichain_rpc_transaction_gas_price
The transaction gas price (gwei)

Metric Type: HistogramVec

Variable Labels:
- network
- provider

## TransactionGasTipCapObserver


### panoptichain_rpc_transaction_gas_tip_cap
The transaction gas tip cap (gwei)

Metric Type: HistogramVec

Variable Labels:
- network
- provider

## TransactionPoolObserver


### panoptichain_rpc_pending_tx_size
Number of pending transactions

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_rpc_queued_tx_size
Number of queued transactions

Metric Type: GaugeVec

Variable Labels:
- network
- provider

## TransactionValueObserver


### panoptichain_rpc_transaction_value
The value of the transactions (ether)

Metric Type: HistogramVec

Variable Labels:
- network
- provider

## TrustedBatchObserver


### panoptichain_rpc_transactions_per_batch
The number of transactions per trusted batch

Metric Type: HistogramVec

Variable Labels:
- network
- provider

## UnclesObserver


### panoptichain_rpc_uncles
The number of uncles for the block

Metric Type: CounterVec

Variable Labels:
- network
- provider

## ValidatorWalletBalanceObserver


### panoptichain_rpc_validator_wallet_balance
PoS validator wallet balance

Metric Type: GaugeVec

Variable Labels:
- network
- provider
- signer_address

## ZkEVMBatchObserver


### panoptichain_rpc_trusted_batch
zkEVM trusted batch number

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_rpc_virtual_batch
zkEVM virtual batch number

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_rpc_verified_batch
zkEVM verified batch number

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_rpc_time_since_last_trusted_batch
time since last zkEVM trusted batch (in seconds)

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_rpc_time_since_last_virtual_batch
time since last zkEVM virtual batch (in seconds)

Metric Type: GaugeVec

Variable Labels:
- network
- provider

### panoptichain_rpc_time_since_last_verified_batch
time since last zkEVM verified batch (in seconds)

Metric Type: GaugeVec

Variable Labels:
- network
- provider
