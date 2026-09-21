package observer

import (
	"encoding/json"
	"testing"
)

// The exact body heimdall-api-amoy.polygon.technology returns when nothing is
// buffered. bor_chain_id is "" because it is a string on the wire, and this
// used to fail the whole decode with:
//
//	json: invalid use of ,string struct tag, trying to unmarshal "" into uint64
const emptyBufferedCheckpoint = `{"checkpoint":{"id":"0","proposer":"","start_block":"0","end_block":"0","root_hash":null,"bor_chain_id":"","timestamp":"0"}}`

// The exact body heimdall-api.polygon.technology returns with one buffered.
const populatedBufferedCheckpoint = `{"checkpoint":{"id":"112112","proposer":"0x9aa1e7971a94cbc400d95e593e82fb6e79fc99d8","start_block":"94206374","end_block":"94206629","root_hash":"NBCJ1ne6u1vP+AWKTLfsAc7iIc09YMZKYeAZshlKw0Y=","bor_chain_id":"137","timestamp":"1790015967"}}`

func TestBufferedCheckpoint_EmptyDecodes(t *testing.T) {
	var v2 HeimdallCheckpointV2
	if err := json.Unmarshal([]byte(emptyBufferedCheckpoint), &v2); err != nil {
		t.Fatalf("empty buffered checkpoint must decode, got: %v", err)
	}

	// ID 0 is what refreshBufferedCheckpoint keys "nothing buffered" off, so it
	// has to survive the decode rather than be lost with the error.
	if v2.Checkpoint.ID != 0 {
		t.Errorf("ID = %d, want 0", v2.Checkpoint.ID)
	}
	if v2.Checkpoint.BorChainID != 0 {
		t.Errorf("BorChainID = %d, want 0", v2.Checkpoint.BorChainID)
	}
}

func TestBufferedCheckpoint_PopulatedDecodes(t *testing.T) {
	var v2 HeimdallCheckpointV2
	if err := json.Unmarshal([]byte(populatedBufferedCheckpoint), &v2); err != nil {
		t.Fatalf("populated buffered checkpoint must decode, got: %v", err)
	}

	for _, c := range []struct {
		name string
		got  uint64
		want uint64
	}{
		{"ID", v2.Checkpoint.ID, 112112},
		{"StartBlock", v2.Checkpoint.StartBlock, 94206374},
		{"EndBlock", v2.Checkpoint.EndBlock, 94206629},
		{"BorChainID", uint64(v2.Checkpoint.BorChainID), 137},
		{"Timestamp", v2.Checkpoint.Timestamp, 1790015967},
	} {
		if c.got != c.want {
			t.Errorf("%s = %d, want %d", c.name, c.got, c.want)
		}
	}
}

// A genuinely malformed chain id must still be an error rather than a silent 0.
func TestChainID_RejectsGarbage(t *testing.T) {
	var c ChainID
	if err := json.Unmarshal([]byte(`"not-a-number"`), &c); err == nil {
		t.Fatal("expected an error for a non-numeric bor_chain_id")
	}
}

func TestChainID_AcceptsNullAndEmpty(t *testing.T) {
	for _, in := range []string{`""`, `null`} {
		var c ChainID
		if err := json.Unmarshal([]byte(in), &c); err != nil {
			t.Errorf("Unmarshal(%s): %v", in, err)
		}
		if c != 0 {
			t.Errorf("Unmarshal(%s) = %d, want 0", in, c)
		}
	}
}
