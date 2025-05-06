// Package config handles application configuration by loading values from files
// and environment variables.
package config

import (
	"flag"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

const DefaultBlockLookBack uint64 = 1000

// Runner configures the execution interval of the job system.
type Runner struct {
	Interval *time.Duration `mapstructure:"interval" validate:"required"`
}

// Providers encloses the different providers configurations. Providers are
// responsible for fetching data.
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
	Name          string         `mapstructure:"name" validate:"required"`
	URL           string         `mapstructure:"url" validate:"url,required"`
	Label         string         `mapstructure:"label" validate:"required"`
	Interval      *time.Duration `mapstructure:"interval"`
	Contracts     Contracts      `mapstructure:"contracts"`
	TimeToMine    *TimeToMine    `mapstructure:"time_to_mine"`
	Accounts      []string       `mapstructure:"accounts"`
	BlockLookBack *uint64        `mapstructure:"block_look_back"`
	TxPool        bool           `mapstructure:"txpool"`
}

// Contracts maps specific contracts to their addresses. This is used to
// fetch on-chain data from these contracts.
type Contracts struct {
	// PoS
	StateSyncSenderAddress   *string `mapstructure:"state_sync_sender_address"`
	StateSyncReceiverAddress *string `mapstructure:"state_sync_receiver_address"`
	CheckpointAddress        *string `mapstructure:"checkpoint_address"`

	// zkEVM
	GlobalExitRootL2Address *string       `mapstructure:"global_exit_root_l2_address"`
	ZkEVMBridgeAddress      *string       `mapstructure:"zkevm_bridge_address"`
	RollupManagerAddress    *string       `mapstructure:"rollup_manager_address"`
	RollupManager           RollupManager `mapstructure:"rollup_manager"`
}

// RollupManager defines the configuration for managing a set of rollups.
type RollupManager struct {
	Rollups  map[uint32]Rollup `mapstructure:"rollups"`
	Enabled  []uint32          `mapstructure:"enabled"`
	Disabled []uint32          `mapstructure:"disabled"`
}

// Rollup represents the configuration for a single rollup. Every field of the
// `RPC` provider can be overridden here.
type Rollup struct {
	RPC   `mapstructure:",squash"`
	Name  *string `mapstructure:"name"`
	URL   *string `mapstructure:"url" validate:"url"`
	Label *string `mapstructure:"label"`
}

// TimeToMine configures the time to mine provider. This periodically sends
// transactions on the network and records how long they took to be recorded in
// a block.
type TimeToMine struct {
	Sender           string `mapstructure:"sender" validate:"required"`
	SenderPrivateKey string `mapstructure:"sender_private_key" validate:"required"`
	Receiver         string `mapstructure:"receiver" validate:"required"`
	Value            int64  `mapstructure:"value" validate:"required"`
	Data             string `mapstructure:"data"`
	GasPriceFactor   int64  `mapstructure:"gas_price_factor"`
	GasLimit         uint64 `mapstructure:"gas_limit" validate:"required"`
}

// HashDivergence configures the hash divergence provider. This tracks whether
// blocks with the same block number have different hashes.
type HashDivergence struct {
	Interval *time.Duration `mapstructure:"interval"`
}

// System configures the system provider. This keeps system diagnostic metrics
// such as uptime.
type System struct {
	Interval *time.Duration `mapstructure:"interval"`
}

// ExchangeRates configures the exchange rates provider. This fetches exchange
// rate data from the Coinbase API.
type ExchangeRates struct {
	CoinbaseURL string              `mapstructure:"coinbase_url" validate:"required"`
	Tokens      map[string][]string `mapstructure:"tokens"`
	Interval    *time.Duration      `mapstructure:"interval"`
}

// HeimdallEndpoint configures the heimdall provider. This provider fetches data
// from the consensus layer endpoints for Polygon PoS chains.
type HeimdallEndpoint struct {
	Name          string         `mapstructure:"name" validate:"required"`
	TendermintURL string         `mapstructure:"tendermint_url" validate:"url,required"`
	HeimdallURL   string         `mapstructure:"heimdall_url" validate:"url,required"`
	Label         string         `mapstructure:"label" validate:"required"`
	Interval      *time.Duration `mapstructure:"interval"`
	Version       *uint          `mapstructure:"version" validate:"omitempty,oneof=1 2"`
}

// SensorNetwork configures the sensor network provider. This fetches data from
// GCP Datastore where the sensors write their data.
type SensorNetwork struct {
	Name     string         `mapstructure:"name" validate:"required"`
	Label    string         `mapstructure:"label" validate:"required"`
	Project  string         `mapstructure:"project" validate:"required"`
	Database string         `mapstructure:"database"`
	Interval *time.Duration `mapstructure:"interval"`
}

// Observers defines which observers should be enabled or disabled. Observers
// are responsible for emitting the Prometheus metrics.
type Observers struct {
	Enabled  []string `mapstructure:"enabled"`
	Disabled []string `mapstructure:"disabled"`
}

// HTTP defines the properties that used for exposing metrics.
type HTTP struct {
	PromPort  int    `mapstructure:"port" validate:"required"`
	PprofPort int    `mapstructure:"pprof_port" validate:"required"`
	Address   string `mapstructure:"address" validate:"required"`
	Path      string `mapstructure:"path" validate:"required"`
}

// Network defines metadata about a blockchain network.
type Network struct {
	Name         string `mapstructure:"name" validate:"required"`
	ChainID      uint64 `mapstructure:"chain_id"`
	PolygonPoS   bool   `mapstructure:"polygon_pos"`
	PolygonZkEVM bool   `mapstructure:"polygon_zkevm"`
}

// GetName returns the network name.
func (n Network) GetName() string {
	return n.Name
}

// GetChainID returns the network chain ID.
func (n Network) GetChainID() uint64 {
	return n.ChainID
}

// IsPolygonPoS returns if this is a Polygon PoS chain.
func (n Network) IsPolygonPoS() bool {
	return n.PolygonPoS
}

// IsPolygonZkEVM returns if the network is a Polygon zkEVM chain.
func (n Network) IsPolygonZkEVM() bool {
	return n.PolygonZkEVM
}

// Logs configures logging format and verbosity options.
type Logs struct {
	Pretty    bool   `mapstructure:"pretty"`
	Verbosity string `mapstructure:"verbosity"`
}

type config struct {
	Namespace string    `mapstructure:"namespace" validate:"required"`
	Runner    Runner    `mapstructure:"runner"`
	HTTP      HTTP      `mapstructure:"http"`
	Providers Providers `mapstructure:"providers"`
	Observers Observers `mapstructure:"observers"`
	Networks  []Network `mapstructure:"networks"`
	Logs      Logs      `mapstructure:"logs"`
}

var c *config

// Config returns the configuration. `Init()` should be called before this.
func Config() *config {
	return c
}

// expandEnvHookFunc expands environment variables when the viper is decoding
// into the `config` struct.
func expandEnvHookFunc() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data any) (any, error) {
		if f.Kind() == reflect.String && t.Kind() == reflect.String {
			return os.ExpandEnv(data.(string)), nil
		}
		return data, nil
	}
}

// Init initializes the config. This should be called before using `Config()`.
func Init() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/panoptichain/")

	if len(flag.Args()) > 0 {
		viper.SetConfigFile(flag.Args()[0])
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("panoptichain")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetDefault("namespace", "panoptichain")
	viper.SetDefault("runner.interval", "30s")
	viper.SetDefault("http.port", 9090)
	viper.SetDefault("http.pprof_port", 6060)
	viper.SetDefault("http.address", "localhost")
	viper.SetDefault("http.path", "/metrics")
	viper.SetDefault("logs.pretty", false)
	viper.SetDefault("logs.verbosity", "info")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	opts := viper.DecodeHook(
		mapstructure.ComposeDecodeHookFunc(
			expandEnvHookFunc(),
			mapstructure.StringToTimeDurationHookFunc(),
		),
	)

	if err := viper.Unmarshal(&c, opts); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(c); err != nil {
		return err
	}

	return nil
}
