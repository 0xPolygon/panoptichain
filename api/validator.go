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

const refreshInterval = time.Hour

type Validator struct {
	ID          uint64 `json:"ID"`
	Signer      string `json:"signer"`
	StartEpoch  uint64 `json:"startEpoch"`
	EndEpoch    uint64 `json:"endEpoch"`
	Nonce       uint64 `json:"nonce"`
	Power       uint64 `json:"power"`
	PubKey      string `json:"pubKey"`
	LastUpdated string `json:"last_updated"`
	Jailed      bool   `json:"jailed"`
	Accum       int64  `json:"accum"`
}

type ValidatorSet struct {
	Height string `json:"height"`
	Result struct {
		Validators []*Validator `json:"validators"`
		Proposer   *Validator   `json:"proposer"`
	} `json:"result"`
}

type ValidatorsCache struct {
	validators []*Validator
	ttl        time.Time
}

var cache sync.Map

// Validators queries the Heimdall API for the validator set. The validator set
// is cached based on the refreshInterval.
func Validators(n network.Network) ([]*Validator, error) {
	var heimdallURL *string
	for _, heimdall := range config.Config().Providers.HeimdallEndpoints {
		if heimdall.Name == n.GetName() {
			heimdallURL = &heimdall.HeimdallURL
			break
		}
	}

	if heimdallURL == nil {
		return nil, errors.New("no validators for this network")
	}

	value, ok := cache.Load(n)
	if ok {
		vc, ok := value.(ValidatorsCache)
		if !ok {
			return nil, errors.New("validator cache type assertion failed")
		}

		if time.Now().Before(vc.ttl) {
			return vc.validators, nil
		}
	}

	path, err := url.JoinPath(*heimdallURL, "staking/validator-set")
	if err != nil {
		return nil, err
	}

	var body *ValidatorSet
	err = GetJSON(path, &body)
	if err != nil {
		return nil, err
	}

	if body == nil || body.Result.Validators == nil {
		return nil, errors.New("empty validator body response")
	}

	cache.Store(n, ValidatorsCache{
		validators: body.Result.Validators,
		ttl:        time.Now().Add(refreshInterval),
	})

	return body.Result.Validators, nil
}

// Signers maps the validator signer to the validator.
func Signers(n network.Network) (map[string]*Validator, error) {
	validators, err := Validators(n)
	if err != nil {
		return nil, err
	}

	signers := make(map[string]*Validator)
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
