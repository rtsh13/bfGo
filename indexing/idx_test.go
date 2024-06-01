package indexing

import (
	"reflect"
	"testing"
)

func TestModuloBiasing(t *testing.T) {
	tests := []struct {
		name     string
		hashes   []uint
		n        uint
		expected []uint
	}{
		{"Basic Functionality", []uint{10, 20, 30}, 3, []uint{1, 2, 0}},
		{"Single Hash Value", []uint{15}, 4, []uint{3}},
		{"Multiple Hash Values", []uint{5, 25, 125}, 6, []uint{5, 1, 5}},
		{"Large Hash Values", []uint{1000000007, 1000000009}, 100, []uint{7, 9}},
		{"Identical Hash Values", []uint{7, 7, 7}, 5, []uint{2, 2, 2}},
		{"Empty Input", []uint{}, 5, []uint{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ModuloBiasing(tt.hashes, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
