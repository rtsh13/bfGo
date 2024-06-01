package errors

import (
	"testing"
)

func TestBloomSizeError(t *testing.T) {
	tests := []struct {
		size     uint
		expected string
	}{
		{0, "invalid filter size : [0] provided"},
		{100, "invalid filter size : [100] provided"},
		{1000000, "invalid filter size : [1000000] provided"},
	}

	for _, tt := range tests {
		err := BloomSize{Size: tt.size}
		if err.Error() != tt.expected {
			t.Errorf("expected %q, got %q", tt.expected, err.Error())
		}
	}
}

func TestBloomKicks(t *testing.T) {
	tests := []struct {
		kicks    uint
		expected string
	}{
		{0, "invalid kick size : [0] provided"},
		{100, "invalid kick size : [100] provided"},
		{1000000, "invalid kick size : [1000000] provided"},
	}

	for _, tt := range tests {
		err := BloomKicks{Kick: tt.kicks}
		if err.Error() != tt.expected {
			t.Errorf("expected %q, got %q", tt.expected, err.Error())
		}
	}
}

func Test_BloomSlots(t *testing.T) {
	tests := []struct {
		slot     uint
		expected string
	}{
		{0, "invalid slots : [0] provided"},
		{100, "invalid slots : [100] provided"},
		{1000000, "invalid slots : [1000000] provided"},
	}

	for _, tt := range tests {
		err := BloomSlots{Slot: tt.slot}
		if err.Error() != tt.expected {
			t.Errorf("expected %q, got %q", tt.expected, err.Error())
		}
	}
}
