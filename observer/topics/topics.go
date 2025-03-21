package topics

//go:generate stringer -type=ObservableTopic
type ObservableTopic int

const (
	NewEVMBlock                 ObservableTopic = iota // *types.Block
	BorStateSync                                       // *observer.StateSync
	BlockInterval                                      // uint64
	CheckpointSignatures                               // *observer.CheckpointSignatures
	ValidatorWallet                                    // observer.ValidatorWalletBalances
	HeimdallBlockInterval                              // uint64
	NewHeimdallBlock                                   // *observer.HeimdallBlock
	Milestone                                          // *observer.HeimdallMilestone
	Reorg                                              // *observer.DatastoreReorg
	SensorBlocks                                       // *observer.SensorBlocks
	SensorBlockEvents                                  // *observer.SensorBlockEvents
	BorMissedBlockProposal                             // observer.MissedBlockProposal
	HeimdallMissedBlockProposal                        // *observer.HeimdallMissedBlockProposal
	Checkpoint                                         // *observer.HeimdallCheckpoint
	MissedCheckpointProposal                           // []string
	MissedMilestoneProposal                            // []string
	TransactionPool                                    // *observer.TransactionPool
	StolenBlock                                        // *types.Block
	HashDivergence                                     // *observer.HashDivergence
	System                                             // *observer.System
	RefreshStateTime                                   // *time.Duration
	ZkEVMBatches                                       // observer.ZkEVMBatches
	ExitRoots                                          // *observer.ExitRoots
	BridgeEvent                                        // *contracts.PolygonZkEVMBridgeV2BridgeEvent
	ClaimEvent                                         // *contracts.PolygonZkEVMBridgeV2ClaimEvent
	DepositCounts                                      // *observer.DepositCounts
	BridgeEventTimes                                   // *observer.BridgeEventTimes
	ClaimEventTimes                                    // *observer.ClaimEventTimes
	RollupManager                                      // *observer.RollupManager
	Span                                               // *observer.HeimdallSpan
	TimeToMine                                         // float64
	AccountBalances                                    // observer.AccountBalances
	TrustedBatch                                       // *zkevmtypes.Batch
	ExchangeRate                                       // observer.ExchangeRate
	TimeToFinalized                                    // uint64
)
