// Package config handles application configuration by loading values from files
// and environment variables.
package config

import (
	"os"
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
	Name          string            `mapstructure:"name"`
	URL           string            `mapstructure:"url" validate:"url,required_with=Name"`
	Label         string            `mapstructure:"label" validate:"required_with=Name"`
	Interval      uint              `mapstructure:"interval"`
	Contracts     ContractAddresses `mapstructure:"contracts"`
	TimeToMine    *TimeToMine       `mapstructure:"time_to_mine"`
	Accounts      []string          `mapstructure:"accounts"`
	BlockLookBack *uint64           `mapstructure:"block_look_back"`
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

// HTTP defines the properties that used for exposing metrics.
type HTTP struct {
	PromPort  int    `mapstructure:"port"`
	PprofPort int    `mapstructure:"pprof_port"`
	Address   string `mapstructure:"address"`
	Path      string `mapstructure:"path"`
}

type CustomNetwork struct {
	Name    string `mapstructure:"name" validate:"required"`
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
	Providers Providers       `mapstructure:"providers"`
	Observers Observers       `mapstructure:"observers"`
	Networks  []CustomNetwork `mapstructure:"networks"`
	Logs      Logs            `mapstructure:"logs"`
}

var c *config

// Config returns the configuration. `Init()` should be called before this.
func Config() *config {
	return c
}

// Init initializes the config. This should be called before using `Config()`.
func Init(args []string) error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/panoptichain/")

	if len(args) > 1 {
		viper.SetConfigFile(args[1])
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("panoptichain")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetDefault("namespace", "panoptichain")
	viper.SetDefault("runner.interval", 30)
	viper.SetDefault("http.port", 9090)
	viper.SetDefault("http.pprof_port", 6060)
	viper.SetDefault("http.address", "localhost")
	viper.SetDefault("http.path", "/metrics")
	viper.SetDefault("logs.pretty", false)
	viper.SetDefault("logs.verbosity", "info")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	for _, k := range viper.AllKeys() {
		v := viper.GetString(k)
		viper.Set(k, os.ExpandEnv(v))
	}

	if err := viper.Unmarshal(&c); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(c); err != nil {
		return err
	}

	return nil
}
