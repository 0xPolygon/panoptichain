package topics

//go:generate stringer -type=ObservableTopic
type ObservableTopic int

const (
	AccountBalances             ObservableTopic = iota // observer.AccountBalances
	AggchainEvent                                      // *observer.AggchainEvent
	BlockInterval                                      // uint64
	BorMissedBlockProposal                             // observer.MissedBlockProposal
	BorStateSync                                       // *observer.StateSync
	BridgeEvent                                        // *contracts.PolygonZkEVMBridgeV2BridgeEvent
	BridgeEventTimes                                   // *observer.BridgeEventTimes
	Checkpoint                                         // *observer.HeimdallCheckpoint
	CheckpointSignatures                               // *observer.CheckpointSignatures
	ClaimEvent                                         // *contracts.PolygonZkEVMBridgeV2ClaimEvent
	ClaimEventTimes                                    // *observer.ClaimEventTimes
	DepositCounts                                      // *observer.DepositCounts
	ExchangeRate                                       // observer.ExchangeRate
	ExitRoots                                          // *observer.ExitRoots
	FinalizedHeight                                    // uint64
	HashDivergence                                     // *observer.HashDivergence
	HeimdallBlockInterval                              // uint64
	HeimdallMissedBlockProposal                        // observer.HeimdallMissedBlockProposal
	Milestone                                          // *observer.HeimdallMilestone
	MissedCheckpointProposal                           // []string
	NewEVMBlock                                        // *types.Block
	NewHeimdallBlock                                   // *observer.HeimdallBlock
	ProofRequest                                       // *proto.ProofRequest
	RefreshStateTime                                   // *time.Duration
	Reorg                                              // *observer.DatastoreReorg
	RollupManager                                      // *observer.RollupManager
	SensorBlockEvents                                  // *observer.SensorBlockEvents
	SensorBlocks                                       // *observer.SensorBlocks
	Span                                               // *observer.HeimdallSpan
	StolenBlock                                        // *types.Block
	System                                             // *observer.System
	TimeToFinalized                                    // uint64
	TimeToMine                                         // float64
	TransactionPool                                    // *observer.TransactionPool
	TrustedBatch                                       // *util.Batch
	ValidatorWallet                                    // observer.ValidatorWalletBalances
	ZkEVMBatches                                       // observer.ZkEVMBatches
)
