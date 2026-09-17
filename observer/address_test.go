package observer

import "testing"

// One validator, spelled as each live upstream spells it.
func TestCanonicalAddress_AllUpstreamSpellingsAgree(t *testing.T) {
	const want = "0xeedba2484aaf940f37cd3cd21a5d7c4a7dafbfc0"

	for name, in := range map[string]string{
		"heimdall app (0x lowercase)": "0xeedba2484aaf940f37cd3cd21a5d7c4a7dafbfc0",
		"tendermint (bare uppercase)": "EEDBA2484AAF940F37CD3CD21A5D7C4A7DAFBFC0",
		"go-ethereum (EIP-55 mixed)":  "0xeEDBa2484aAF940f37cd3CD21a5D7C4A7DAfbfC0",
		"bare lowercase":              "eedba2484aaf940f37cd3cd21a5d7c4a7dafbfc0",
		"uppercase 0X prefix":         "0XEEDBA2484AAF940F37CD3CD21A5D7C4A7DAFBFC0",
	} {
		if got := CanonicalAddress(in); got != want {
			t.Errorf("%s: CanonicalAddress(%q) = %q, want %q", name, in, got, want)
		}
	}
}

// Applied to already-canonical sources too, so it must be a no-op on its own
// output.
func TestCanonicalAddress_Idempotent(t *testing.T) {
	for _, in := range []string{
		"EEDBA2484AAF940F37CD3CD21A5D7C4A7DAFBFC0",
		"0xeEDBa2484aAF940f37cd3CD21a5D7C4A7DAfbfC0",
		"",
	} {
		once := CanonicalAddress(in)
		if twice := CanonicalAddress(once); twice != once {
			t.Errorf("CanonicalAddress(%q): %q then %q", in, once, twice)
		}
	}
}

// Absent validators carry validator_address ""; "0x" would be a worse label.
func TestCanonicalAddress_EmptyStaysEmpty(t *testing.T) {
	if got := CanonicalAddress(""); got != "" {
		t.Fatalf("CanonicalAddress(%q) = %q, want %q", "", got, "")
	}
}

// Prefix and case only: a malformed address must stay visibly malformed rather
// than be padded or coerced into something plausible.
func TestCanonicalAddress_DoesNotReinterpretDigits(t *testing.T) {
	for in, want := range map[string]string{
		"0xABC":     "0xabc",
		"ABC":       "0xabc",
		"0x":        "0x",
		"not-hex":   "0xnot-hex",
		"0xNOT-HEX": "0xnot-hex",
	} {
		if got := CanonicalAddress(in); got != want {
			t.Errorf("CanonicalAddress(%q) = %q, want %q", in, got, want)
		}
	}
}
