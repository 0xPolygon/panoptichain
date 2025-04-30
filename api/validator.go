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

type Validator interface {
	GetID() uint64
	GetSigner() string
	GetStartEpoch() uint64
	GetEndEpoch() uint64
	GetNonce() uint64
	GetPower() uint64
	GetPubKey() string
	GetLastUpdated() string
	IsJailed() bool
}

// ValidatorV1 represents a Polygon PoS validator from Heimdall v1.
type ValidatorV1 struct {
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

func (v ValidatorV1) GetID() uint64          { return v.ID }
func (v ValidatorV1) GetSigner() string      { return v.Signer }
func (v ValidatorV1) GetStartEpoch() uint64  { return v.StartEpoch }
func (v ValidatorV1) GetEndEpoch() uint64    { return v.EndEpoch }
func (v ValidatorV1) GetNonce() uint64       { return v.Nonce }
func (v ValidatorV1) GetPower() uint64       { return v.Power }
func (v ValidatorV1) GetPubKey() string      { return v.PubKey }
func (v ValidatorV1) GetLastUpdated() string { return v.LastUpdated }
func (v ValidatorV1) IsJailed() bool         { return v.Jailed }
func (v ValidatorV1) GetAccum() int64        { return v.Accum }

// ValidatorV2 represents a Polygon PoS validator from Heimdall v2.
type ValidatorV2 struct {
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

func (v ValidatorV2) GetID() uint64              { return v.ID }
func (v ValidatorV2) GetSigner() string          { return v.Signer }
func (v ValidatorV2) GetStartEpoch() uint64      { return v.StartEpoch }
func (v ValidatorV2) GetEndEpoch() uint64        { return v.EndEpoch }
func (v ValidatorV2) GetNonce() uint64           { return v.Nonce }
func (v ValidatorV2) GetPower() uint64           { return v.Power }
func (v ValidatorV2) GetPubKey() string          { return v.PubKey }
func (v ValidatorV2) GetLastUpdated() string     { return v.LastUpdated }
func (v ValidatorV2) IsJailed() bool             { return v.Jailed }
func (v ValidatorV2) GetProposerPriority() int64 { return v.ProposerPriority }

// ValidatorSetV1 is a set of Polygon PoS validators from Heimdall v1.
type ValidatorSetV1 struct {
	Height string `json:"height"`
	Result struct {
		Validators []ValidatorV1 `json:"validators"`
		Proposer   *ValidatorV1  `json:"proposer"`
	} `json:"result"`
}

// ValidatorSetV2 is a set of Polygon PoS validators from Heimdall v2.
type ValidatorSetV2 struct {
	ValidatorSet struct {
		Validators       []ValidatorV2 `json:"validators"`
		Proposer         *ValidatorV2  `json:"proposer"`
		TotalVotingPower uint64        `json:"total_voting_power,string"`
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
	var version uint = 1

	for _, heimdall := range config.Config().Providers.HeimdallEndpoints {
		if heimdall.Name == n.GetName() {
			path = &heimdall.HeimdallURL
			if heimdall.Version != nil {
				version = *heimdall.Version
			}
			break
		}
	}

	if path == nil {
		return nil, errors.New("no validators for this network")
	}

	var err error
	switch version {
	case 1:
		validators, err = getValidatorsV1(*path)
	case 2:
		validators, err = getValidatorsV2(*path)
	}

	if err != nil {
		return nil, err
	}

	cache.Store(n, ValidatorsCache{
		validators: validators,
		ttl:        time.Now().Add(refreshInterval),
	})

	return validators, nil
}

func getValidatorsV1(path string) ([]Validator, error) {
	path, err := url.JoinPath(path, "staking/validator-set")
	if err != nil {
		return nil, err
	}

	var body ValidatorSetV1
	err = GetJSON(path, &body)
	if err != nil {
		return nil, err
	}

	if body.Result.Validators == nil {
		return nil, errors.New("empty validator body response")
	}

	validators := make([]Validator, len(body.Result.Validators))
	for i, v := range body.Result.Validators {
		validators[i] = v
	}

	return validators, nil
}

func getValidatorsV2(path string) ([]Validator, error) {
	path, err := url.JoinPath(path, "stake/validator-set")
	if err != nil {
		return nil, err
	}

	var body ValidatorSetV2
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
		signers[validator.GetSigner()] = validator
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
