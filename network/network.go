// Package network defines objects that we'll use for understanding which network we are observing
package network

import (
	"fmt"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/log"
)

// Network is the interface that we need to identify a network.
type Network interface {
	GetName() string
}

const (
	PolygonMainnetName = "Polygon Mainnet"
	PolygonMumbaiName  = "Polygon Mumbai"
	PolygonAmoyName    = "Polygon Amoy"

	HeimdallMainnetName = "Heimdall Mainnet"
	HeimdallMumbaiName  = "Heimdall Mumbai"
	HeimdallAmoyName    = "Heimdall Amoy"

	ZkEVMMainnetName = "zkEVM Mainnet"
	ZkEVMTestnetName = "zkEVM Testnet"
	ZkEVMCardonaName = "zkEVM Cardona"
	ZkEVMBaliName    = "zkEVM Bali"

	EthereumName = "Ethereum"
	SepoliaName  = "Sepolia"
	GoerliName   = "Goerli"
)

var PolygonMainnet = EVMNetwork{Name: PolygonMainnetName, ChainID: 137}
var PolygonMumbai = EVMNetwork{Name: PolygonMumbaiName, ChainID: 80001}
var PolygonAmoy = EVMNetwork{Name: PolygonAmoyName, ChainID: 80002}

var HeimdallMainnet = TendermintNetwork{Name: HeimdallMainnetName, ChainID: "heimdall-137"}
var HeimdallMumbai = TendermintNetwork{Name: HeimdallMumbaiName, ChainID: "heimdall-80001"}
var HeimdallAmoy = TendermintNetwork{Name: HeimdallAmoyName, ChainID: "heimdall-80002"}

var ZkEVMMainnet = EVMNetwork{Name: ZkEVMMainnetName, ChainID: 1101}
var ZkEVMTestnet = EVMNetwork{Name: ZkEVMTestnetName, ChainID: 1442}
var ZkEVMCardona = EVMNetwork{Name: ZkEVMCardonaName, ChainID: 2442}
var ZkEVMBali = EVMNetwork{Name: ZkEVMBaliName, ChainID: 2440}

var Ethereum = EVMNetwork{Name: EthereumName, ChainID: 1}
var Sepolia = EVMNetwork{Name: SepoliaName, ChainID: 11155111}
var Goerli = EVMNetwork{Name: GoerliName, ChainID: 5}

var KnownNetworks = []Network{
	&PolygonMainnet,
	&PolygonMumbai,
	&PolygonAmoy,

	&HeimdallMainnet,
	&HeimdallMumbai,
	&HeimdallAmoy,

	&ZkEVMMainnet,
	&ZkEVMTestnet,
	&ZkEVMCardona,
	&ZkEVMBali,

	&Ethereum,
	&Goerli,
	&Sepolia,
}

// EVMNetwork is a specific type of network that is assumed to generally follow
// the norms of an Ethereum based network.
type EVMNetwork struct {
	Name    string
	ChainID uint64
}

// GetName returns the network name. It doesn't have to be canonical in anyway.
// Just something to match configs.
func (n *EVMNetwork) GetName() string {
	return n.Name
}

// GetChainID returns the configured chain ID for the network.
func (n *EVMNetwork) GetChainID() uint64 {
	return n.ChainID
}

// TendermintNetwork is close to an EVMNetwork, but generally, the chain
// ID is a string, and there won't be typical JSON RPC providers.
type TendermintNetwork struct {
	Name    string
	ChainID string
}

// GetName returns the configured name for this network.
func (n *TendermintNetwork) GetName() string {
	return n.Name
}

// GetChainID returns the chain ID for the network.
func (n *TendermintNetwork) GetChainID() string {
	return n.ChainID
}

// GetNetworkByName is a convenience method to convert a name like "Ethereum"
// into a Network object.
func GetNetworkByName(name string) (Network, error) {
	for _, n := range config.Config().Networks {
		if n.GetName() == name {
			return n, nil
		}
	}

	for _, n := range KnownNetworks {
		if n.GetName() == name {
			return n, nil
		}
	}

	return nil, fmt.Errorf("the network name %s is not recognized", name)
}

func init() {
	names := make(map[string]struct{}, 0)
	for _, k := range KnownNetworks {
		if _, ok := names[k.GetName()]; ok {
			log.Fatal().Str("name", k.GetName()).Msg("Two networks with the same name detected")
		}

		names[k.GetName()] = struct{}{}
	}
}
