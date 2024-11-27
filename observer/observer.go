package observer

import (
	"context"
	"math"
	"reflect"
	"sort"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/log"
	"github.com/0xPolygon/panoptichain/network"
)

// Observer defines the functioned required for adding a new observer. This
// (ideally) will be the most commonly implemented interface. Each metric be
// managed by an observer. An observer can manage multiple metrics. It's also
// worth noting that a single observer will be triggered by multiple providers
// so the observer needs to be aware of networks and providers.
type Observer interface {
	// Notify is called when a message is published.
	Notify(context.Context, Message)

	// Register allows the observer to substribe to whatever topics it cares about
	// and setup any metrics.
	Register(*EventBus)

	// GetCollectors will return a slice of the collectors that are
	// being tracked in this observer.
	GetCollectors() []prometheus.Collector
}

// ObserverSet is an abstraction for a collection of observers that is
// itself an observer.
type ObserverSet []Observer

// CoreMessage implements the shared functionality that will be common to all
// messages.
type CoreMessage struct {
	time     time.Time
	network  network.Network
	provider string
	data     any
}

// Message is the basic abstraction that we'll use for pub/sub
// here. Every message should be specific to a network and for a given
// provider.
type Message interface {
	Time() time.Time
	Network() network.Network
	Provider() string
	Data() any
}

// NewMessage will create a new core message type.
func NewMessage(n network.Network, label string, data any) *CoreMessage {
	return &CoreMessage{
		time:     time.Now(),
		network:  n,
		provider: label,
		data:     data,
	}
}

// Time returns the time that the message was created.
func (cm *CoreMessage) Time() time.Time {
	return cm.time
}

// Network returns the network that this particular message relates to.
func (cm *CoreMessage) Network() network.Network {
	return cm.network
}

// Provider returns the label for the provider.
func (cm *CoreMessage) Provider() string {
	return cm.provider
}

// Data returns the generic data that was stored inside CoreMessage.
func (cm *CoreMessage) Data() any {
	return cm.data
}

// Topic is a basic abstraction for events that are going to be published or
// subscribed.
type Topic interface {
	String() string
}

// EventBus is the object that will be responsible for passing messages between
// providers and observers.
type EventBus struct {
	observers map[string]ObserverSet
	jobs      chan struct{}
}

// Subscribe will configure the given observer to be notified whenever the given
// topic is published. It's up to the observer to filter and handle the variety
// of topics that could be created.
func (eb *EventBus) Subscribe(topic Topic, o Observer) {
	eb.observers[topic.String()] = append(eb.observers[topic.String()], o)
}

// Publish is called but the providers when they want to send a
// message to all subscribers.
func (eb *EventBus) Publish(ctx context.Context, topic Topic, m Message) {
	if len(eb.observers[topic.String()]) == 0 {
		log.Warn().Str("topic", topic.String()).Msg("Topic published to empty subscriber set")
	}

	for _, s := range eb.observers[topic.String()] {
		eb.jobs <- struct{}{}
		go func(o Observer) {
			o.Notify(ctx, m)
			<-eb.jobs
		}(s)
	}
}

func (eb *EventBus) Jobs() int {
	return len(eb.jobs)
}

// NewEventBus is convenience constructor for the EventBus.
func NewEventBus() *EventBus {
	return &EventBus{
		observers: make(map[string]ObserverSet),
		jobs:      make(chan struct{}, 1024),
	}
}

var observersMap = map[string]Observer{
	"account_balances":                    new(AccountBalancesObserver),
	"base_fee_per_gas":                    new(BaseFeePerGasObserver),
	"block":                               new(BlockObserver),
	"block_interval":                      new(BlockIntervalObserver),
	"bogon_block":                         new(BogonBlockObserver),
	"bridge_event":                        new(BridgeEventObserver),
	"checkpoint":                          new(CheckpointObserver),
	"claim_event":                         new(ClaimEventObserver),
	"deposit_counts":                      new(DepositCountObserver),
	"double_sign":                         new(DoubleSignObserver),
	"empty_block":                         new(EmptyBlockObserver),
	"exchange_rates":                      new(ExchangeRatesObserver),
	"exit_roots":                          new(ExitRootsObserver),
	"gas_limit":                           new(GasLimitObserver),
	"gas_used":                            new(GasUsedObserver),
	"hash_divergence":                     new(HashDivergenceObserver),
	"heimdall_block_interval":             new(HeimdallBlockIntervalObserver),
	"heimdall_checkpoint":                 new(HeimdallCheckpointObserver),
	"heimdall_height":                     new(HeimdallHeightObserver),
	"heimdall_missed_block_proposal":      new(HeimdallMissedBlockProposalObserver),
	"heimdall_missed_checkpoint_proposal": new(HeimdallMissedCheckpointProposalObserver),
	"heimdall_missed_milestone_proposal":  new(HeimdallMissedMilestoneProposal),
	"heimdall_signature_count":            new(HeimdallSignatureCountObserver),
	"heimdall_total_transaction_count":    new(HeimdallTotalTransactionCountObserver),
	"heimdall_transaction_count":          new(HeimdallTransactionCountObserver),
	"milestone":                           new(MilestoneObserver),
	"missed_block_proposal":               new(MissedBlockProposalObserver),
	"refresh_state_time":                  new(RefreshStateTimeObserver),
	"reorg":                               new(ReorgObserver),
	"sealed_out_of_turn":                  new(SealedOutOfTurnObserver),
	"sensor_block_events":                 new(BlockEventsObserver),
	"sensor_blocks":                       new(SensorBlocksObserver),
	"sensor_bogon_block":                  new(SensorBogonBlockObserver),
	"state_sync":                          new(StateSyncObserver),
	"stolen_block":                        new(StolenBlockObserver),
	"system":                              new(SystemObserver),
	"time_to_finalized":                   new(TimeToFinalizedObserver),
	"time_to_mine":                        new(TimeToMineObserver),
	"transaction_cost":                    new(TransactionCostObserver),
	"transaction_count":                   new(TransactionCountObserver),
	"transaction_gas_fee_cap":             new(TransactionGasFeeCapObserver),
	"transaction_gas_limit":               new(TransactionGasLimitObserver),
	"transaction_gas_price":               new(TransactionGasPriceObserver),
	"transaction_gas_tip_cap":             new(TransactionGasTipCapObserver),
	"transaction_pool":                    new(TransactionPoolObserver),
	"transaction_value":                   new(TransactionValueObserver),
	"trusted_batch":                       new(TrustedBatchObserver),
	"uncles":                              new(UnclesObserver),
	"validator_wallet_balance":            new(ValidatorWalletBalanceObserver),
	"zkevm_batches":                       new(ZkEVMBatchObserver),
	"rollup_manager":                      new(RollupManagerObserver),
	"span":                                new(HeimdallSpanObserver),
}

func GetEnabledObserverSet() ObserverSet {
	observers := make(ObserverSet, 0, len(observersMap))

	for _, name := range config.Config().Observers {
		observer, ok := observersMap[name]
		if !ok {
			log.Fatal().Msgf("Observer %s does not exist", name)
		}

		observers = append(observers, observer)
	}

	return observers
}

// GetCompleteObserverSet collects all of the known observers for use in
// different packages.
func GetCompleteObserverSet() ObserverSet {
	names := make([]string, 0, len(observersMap))
	observers := make(ObserverSet, 0, len(observersMap))

	for name := range observersMap {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		observers = append(observers, observersMap[name])
	}

	return observers
}

// Register handled register methods for the observer set on the whole.
func (os *ObserverSet) Register(eb *EventBus) {
	for _, o := range *os {
		o.Register(eb)
	}
}

// Notify handles notifications for the observer set on the whole.
func (os *ObserverSet) Notify(ctx context.Context, msg Message) {
	for _, o := range *os {
		o.Notify(ctx, msg)
	}
}

// newExponentialBuckets generates a float64 slice starting from zero to the
// base^exp value (inclusive). The returned slice will have exp+2 buckets.
func newExponentialBuckets(base int, exp int) []float64 {
	buckets := make([]float64, exp+2)
	for i := 1; i < len(buckets); i++ {
		buckets[i] = math.Pow(float64(base), float64(i-1))
	}

	return buckets
}

// newLinearBuckets creates a slice from start to end going from step
// (inclusive).
func newLinearBuckets(start int, end int, step int) []float64 {
	if step <= 0 || end < start {
		return []float64{}
	}

	buckets := make([]float64, 0, 1+(end-start)/step)
	for start <= end {
		buckets = append(buckets, float64(start))
		start += step
	}

	return buckets
}

func NewLogger(o Observer, m Message) zerolog.Logger {
	observer := reflect.ValueOf(o).Elem().Type().Name()

	return log.With().
		Str("observer", observer).
		Str("network", m.Network().GetName()).
		Str("provider", m.Provider()).Time("time", m.Time()).
		Logger()
}
