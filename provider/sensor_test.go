package provider

import "testing"

// newSensorProvider creates a SensorNetworkProvider with the block-tracking
// fields set, suitable for exercising the pure range-clamping logic.
func newSensorProvider(prev, curr uint64) *SensorNetworkProvider {
	return &SensorNetworkProvider{
		prevBlockNumber: prev,
		blockNumber:     curr,
		logger:          NewLogger(nil, "test"),
	}
}

func TestClampStart(t *testing.T) {
	tests := []struct {
		name  string
		prev  uint64
		curr  uint64
		start uint64
		want  uint64
	}{
		{
			name:  "small gap within buffer size is unchanged",
			prev:  blockBufferSize + 1000,
			curr:  blockBufferSize + 1003,
			start: blockBufferSize + 1000,
			want:  blockBufferSize + 1000,
		},
		{
			name:  "large gap is clamped to head minus buffer size",
			prev:  1000,
			curr:  500000,
			start: 1000,
			want:  500000 - blockBufferSize,
		},
		{
			name:  "head below buffer size never underflows",
			prev:  5,
			curr:  blockBufferSize - 100,
			start: 5,
			want:  5,
		},
		{
			name:  "exactly at threshold is unchanged",
			prev:  1000,
			curr:  1000 + blockBufferSize,
			start: 1000,
			want:  1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newSensorProvider(tt.prev, tt.curr)
			got := s.clampStart(tt.start)
			if got != tt.want {
				t.Errorf("clampStart(%d) = %d, want %d", tt.start, got, tt.want)
			}

			// The clamped start must never produce an out-of-range query: it
			// stays within [start, blockNumber] and within blockBufferSize of head.
			if got > tt.curr {
				t.Errorf("clampStart(%d) = %d exceeds head %d", tt.start, got, tt.curr)
			}
			if tt.curr > blockBufferSize && tt.curr-got > blockBufferSize {
				t.Errorf("clampStart(%d) = %d is more than %d behind head %d",
					tt.start, got, blockBufferSize, tt.curr)
			}
		})
	}
}
