package observer

import "strings"

// CanonicalAddress renders an Ethereum-style address as 0x-prefixed lowercase:
// the spelling rpc_missed_checkpoint_signature already used, and so the one
// that needs no downstream change.
//
// Three upstreams spell the same validator three ways, and all three reach
// metric labels, where case is significant:
//
//	heimdall-api    0xeedba2484aaf940f37cd3cd21a5d7c4a7dafbfc0
//	tendermint-api  EEDBA2484AAF940F37CD3CD21A5D7C4A7DAFBFC0
//	Address.Hex()   0xeEDBa2484aAF940f37cd3CD21a5D7C4A7DAfbfC0
//
// Idempotent, so it is applied at every address label site in the heimdall
// observers, including those whose source is already canonical. The rpc_ and
// spn_ families are not converted yet and still carry the Address.Hex()
// spelling.
//
// An empty address stays empty: Tendermint commit signatures for absent
// validators carry validator_address "".
func CanonicalAddress(addr string) string {
	if addr == "" {
		return ""
	}

	bare := strings.TrimPrefix(strings.TrimPrefix(addr, "0x"), "0X")

	return "0x" + strings.ToLower(bare)
}
