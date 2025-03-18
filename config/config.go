// Package confg is a placeholder package for settings that will
// probably be config values at some point in the future. I couldn't
// decide on a config package or library. It's hard coded here for now
// in a way that I hope is convenient for refactoring in the future.
package config

import (
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Runner holds the config values that are needed to manage the job system.
type Runner struct {
	Interval uint `mapstructure:"interval" validate:"required"`
}

type Providers struct {
	RPCs              []RPC              `mapstructure:"rpc" validate:"dive"`
	HeimdallEndpoints []HeimdallEndpoint `mapstructure:"heimdall" validate:"dive"`
	SensorNetworks    []SensorNetwork    `mapstructure:"sensor_network" validate:"dive"`
	HashDivergence    *HashDivergence    `mapstructure:"hash_divergence"`
	System            *System            `mapstructure:"system"`
	ExchangeRates     *ExchangeRates     `mapstructure:"exchange_rates"`
}

// RPC defines the various RPC providers that will be monitored.
type RPC struct {
	Name       string            `mapstructure:"name"`
	URL        string            `mapstructure:"url" validate:"url,required_with=Name"`
	Label      string            `mapstructure:"label" validate:"required_with=Name"`
	Interval   uint              `mapstructure:"interval"`
	Contracts  ContractAddresses `mapstructure:"contracts"`
	TimeToMine *TimeToMine       `mapstructure:"time_to_mine"`
	Accounts   []string          `mapstructure:"accounts"`
}

type ContractAddresses struct {
	// PoS
	StateSyncSenderAddress   *string `mapstructure:"state_sync_sender_address"`
	StateSyncReceiverAddress *string `mapstructure:"state_sync_receiver_address"`
	CheckpointAddress        *string `mapstructure:"checkpoint_address"`

	// zkEVM
	GlobalExitRootL2Address *string `mapstructure:"global_exit_root_l2_address"`
	ZkEVMBridgeAddress      *string `mapstructure:"zkevm_bridge_address"`
	RollupManagerAddress    *string `mapstructure:"rollup_manager_address"`
}

type TimeToMine struct {
	Sender           string `mapstructure:"sender" validate:"required"`
	SenderPrivateKey string `mapstructure:"sender_private_key" validate:"required"`
	Receiver         string `mapstructure:"receiver" validate:"required"`
	Value            int64  `mapstructure:"value" validate:"required"`
	Data             string `mapstructure:"data"`
	GasPriceFactor   int64  `mapstructure:"gas_price_factor"`
	GasLimit         uint64 `mapstructure:"gas_limit" validate:"required"`
}

type HashDivergence struct {
	Interval uint `mapstructure:"interval"`
}

type System struct {
	Interval uint `mapstructure:"interval"`
}

type ExchangeRates struct {
	CoinbaseURL string              `mapstructure:"coinbase_url" validate:"required"`
	Tokens      map[string][]string `mapstructure:"tokens"`
	Interval    uint                `mapstructure:"interval"`
}

type HeimdallEndpoint struct {
	Name          string `mapstructure:"name"`
	TendermintURL string `mapstructure:"tendermint_url" validate:"url,required_with=Name"`
	HeimdallURL   string `mapstructure:"heimdall_url" validate:"url,required_with=Name"`
	Label         string `mapstructure:"label" validate:"required_with=Name"`
	Interval      uint   `mapstructure:"interval"`
}

type SensorNetwork struct {
	Name     string `mapstructure:"name"`
	Label    string `mapstructure:"label" validate:"required_with=Name"`
	Project  string `mapstructure:"project" validate:"required_with=Name"`
	Database string `mapstructure:"database"`
	Interval uint   `mapstructure:"interval"`
}

type Observers struct {
	Enabled  []string `mapstructure:"enabled"`
	Disabled []string `mapstructure:"disabled"`
}

// HTTP defines the HTTP properties that we'll use for exposing metrics.
type HTTP struct {
	Port    int    `mapstructure:"port" validate:"required"`
	Address string `mapstructure:"address" validate:"required"`
	Path    string `mapstructure:"path" validate:"required"`
}

type CustomNetwork struct {
	Name    string `mapstructure:"name"`
	ChainID uint64 `mapstructure:"chain_id"`
}

// GetName returns the network name.
func (n CustomNetwork) GetName() string {
	return n.Name
}

// GetChainID returns the network chain ID.
func (n CustomNetwork) GetChainID() uint64 {
	return n.ChainID
}

type Logs struct {
	Pretty    bool   `mapstructure:"pretty"`
	Verbosity string `mapstructure:"verbosity"`
}

type config struct {
	Namespace string          `mapstructure:"namespace" validate:"required"`
	Runner    Runner          `mapstructure:"runner"`
	HTTP      HTTP            `mapstructure:"http"`
	Providers *Providers      `mapstructure:"providers"`
	Observers Observers       `mapstructure:"observers"`
	Networks  []CustomNetwork `mapstructure:"networks"`
	Logs      Logs            `mapstructure:"logs"`
}

var c *config

// Config will return the current configuration for the entire application
// (assuming it's been initialized).
func Config() *config {
	return c
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/panoptichain/")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("panoptichain")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Failed to read config: %v", err)
	}

	if err := viper.Unmarshal(&c); err != nil {
		log.Panicf("Failed to load config: %v", err)
	}

	validate := validator.New()
	if err := validate.Struct(c); err != nil {
		log.Panicf("Failed to validate config: %v", err)
	}
}
