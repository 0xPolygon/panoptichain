package api

import (
	"errors"
	"net/url"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/consensus/clique"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/network"
)

// refreshInterval is how long validators will be cached for.
const refreshInterval = time.Hour

// Validator represents a Polygon PoS validator from Heimdall.
type Validator struct {
	ID               uint64 `json:"val_id,string"`
	Signer           string `json:"signer"`
	StartEpoch       uint64 `json:"start_epoch,string"`
	EndEpoch         uint64 `json:"end_epoch,string"`
	Nonce            uint64 `json:"nonce,string"`
	Power            uint64 `json:"voting_power,string"`
	PubKey           string `json:"pub_key"`
	LastUpdated      string `json:"last_updated"`
	Jailed           bool   `json:"jailed"`
	ProposerPriority int64  `json:"proposer_priority,string"`
}

// ValidatorSet is a set of Polygon PoS validators from Heimdall.
type ValidatorSet struct {
	ValidatorSet struct {
		Validators       []Validator `json:"validators"`
		Proposer         *Validator  `json:"proposer"`
		TotalVotingPower uint64      `json:"total_voting_power,string"`
	} `json:"validator_set"`
}

// ValidatorsCache holds a cache of validators with a time-to-live (TTL).
type ValidatorsCache struct {
	validators []Validator
	ttl        time.Time
}

// cache maps network.Network to ValidatorsCache.
var cache sync.Map

// getCachedValidators attempts to retrieve a cached set of validators for the
// given network.
func getCachedValidators(n network.Network) ([]Validator, bool) {
	value, ok := cache.Load(n)
	if !ok {
		return nil, false
	}

	vc := value.(ValidatorsCache)
	if time.Now().After(vc.ttl) {
		return nil, false
	}

	return vc.validators, true
}

// Validators queries the Heimdall API for the validator set. The validator set
// is cached based on the refreshInterval.
func Validators(n network.Network) ([]Validator, error) {
	validators, ok := getCachedValidators(n)
	if ok {
		return validators, nil
	}

	var path *string
	for _, heimdall := range config.Config().Providers.HeimdallEndpoints {
		if heimdall.Name == n.GetName() {
			path = &heimdall.HeimdallURL
			break
		}
	}

	if path == nil {
		return nil, errors.New("no validators for this network")
	}

	validators, err := getValidators(*path)
	if err != nil {
		return nil, err
	}

	cache.Store(n, ValidatorsCache{
		validators: validators,
		ttl:        time.Now().Add(refreshInterval),
	})

	return validators, nil
}

func getValidators(path string) ([]Validator, error) {
	path, err := url.JoinPath(path, "stake", "validators-set")
	if err != nil {
		return nil, err
	}

	var body ValidatorSet
	err = GetJSON(path, &body)
	if err != nil {
		return nil, err
	}

	if body.ValidatorSet.Validators == nil {
		return nil, errors.New("empty validator body response")
	}

	validators := make([]Validator, len(body.ValidatorSet.Validators))
	for i, v := range body.ValidatorSet.Validators {
		validators[i] = v
	}

	return validators, nil
}

// Signers maps the validator signer to the validator.
func Signers(n network.Network) (map[string]Validator, error) {
	validators, err := Validators(n)
	if err != nil {
		return nil, err
	}

	signers := make(map[string]Validator)
	for _, validator := range validators {
		signers[validator.Signer] = validator
	}

	return signers, nil
}

// Ecrecover recovers the block signer given the block header.
func Ecrecover(header *types.Header) ([]byte, error) {
	// These values will cause clique.SealHash to panic.
	if header.WithdrawalsHash != nil ||
		header.BlobGasUsed != nil ||
		header.ExcessBlobGas != nil ||
		header.ParentBeaconRoot != nil {
		return nil, errors.New("unable to encode clique header")
	}

	start := len(header.Extra) - crypto.SignatureLength
	if start < 0 || start > len(header.Extra) {
		return nil, errors.New("unable to recover signature")
	}
	signature := header.Extra[start:]
	pubkey, err := crypto.Ecrecover(clique.SealHash(header).Bytes(), signature)
	if err != nil {
		return nil, err
	}
	signer := crypto.Keccak256(pubkey[1:])[12:]

	return signer, nil
}
