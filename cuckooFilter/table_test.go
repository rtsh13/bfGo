package cuckoo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Peek(t *testing.T) {
	tests := []struct {
		name string
		pos  uint
		want uint
	}{
		{"peek filled slot", 1, 2},
		{"peek another filled slot", 2, 3},
		{"peek edge slot", 0, 1},
		{"peek another edge slot", 3, 4},
		{"peek empty slot", 4, 0}, // Assuming slot 4 is initially empty
	}

	h := hashTable{slots: []uint{1, 2, 3, 4, 0}} // Adding an extra slot to test the empty slot case

	for _, tt := range tests {
		got := h.Peek(tt.pos)
		assert.Equalf(t, tt.want, got, "%s: hashtable->peek", tt.name)
	}

	// Peek at multiple slots
	h = hashTable{slots: []uint{10, 20, 30, 40, 50}}
	assert.Equal(t, uint(10), h.Peek(0), "Slot 0 should be 10")
	assert.Equal(t, uint(20), h.Peek(1), "Slot 1 should be 20")
	assert.Equal(t, uint(30), h.Peek(2), "Slot 2 should be 30")
	assert.Equal(t, uint(40), h.Peek(3), "Slot 3 should be 40")
	assert.Equal(t, uint(50), h.Peek(4), "Slot 4 should be 50")

	// Peek at an empty slot
	h = hashTable{slots: []uint{0, 0, 0, 0, 0}}
	assert.Equal(t, uint(0), h.Peek(0), "Slot 0 should be empty (0)")
	assert.Equal(t, uint(0), h.Peek(1), "Slot 1 should be empty (0)")
	assert.Equal(t, uint(0), h.Peek(2), "Slot 2 should be empty (0)")
	assert.Equal(t, uint(0), h.Peek(3), "Slot 3 should be empty (0)")
	assert.Equal(t, uint(0), h.Peek(4), "Slot 4 should be empty (0)")
}

func Test_Occupy(t *testing.T) {
	tests := []struct {
		name   string
		fprint uint
		pos    uint
	}{
		{"occupy empty slot", 10, 1},
		{"occupy filled slot", 20, 2},
		{"occupy edge slot", 30, 0},
		{"occupy another edge slot", 40, 3},
	}

	h := hashTable{slots: []uint{0, 2, 3, 0}}

	for _, tt := range tests {
		preValue := h.Peek(tt.pos)
		h.Occupy(tt.fprint, tt.pos)
		postValue := h.Peek(tt.pos)
		assert.Equalf(t, tt.fprint, postValue, "%s: hashtable->occupy", tt.name)
		assert.NotEqual(t, preValue, postValue, "%s: pre and post value should not be equal", tt.name)
	}

	// Occupy multiple slots
	h = hashTable{slots: []uint{0, 0, 0, 0}}
	h.Occupy(50, 0)
	h.Occupy(60, 1)
	h.Occupy(70, 2)
	h.Occupy(80, 3)
	assert.Equal(t, uint(50), h.Peek(0), "Slot 0 should be occupied by 50")
	assert.Equal(t, uint(60), h.Peek(1), "Slot 1 should be occupied by 60")
	assert.Equal(t, uint(70), h.Peek(2), "Slot 2 should be occupied by 70")
	assert.Equal(t, uint(80), h.Peek(3), "Slot 3 should be occupied by 80")

	// Occupy already filled slots
	preValue := h.Peek(2)
	h.Occupy(90, 2)
	postValue := h.Peek(2)
	assert.NotEqual(t, preValue, postValue, "Slot 2 should have a new value after re-occupation")
	assert.Equal(t, uint(90), postValue, "Slot 2 should be occupied by 90 after re-occupation")
}

func Test_IsSlotAvailable(t *testing.T) {
	tests := []struct {
		name string
		pos  uint
		want bool
	}{
		{"slot available", 1, true},
		{"slot not available", 2, false},
		{"edge slot available", 0, true},
		{"edge slot not available", 3, false},
	}

	h := newTable(4)

	// check availability of an empty slot
	got := h.IsSlotAvailable(tests[0].pos)
	assert.Equalf(t, tests[0].want, got, tests[0].name)

	// fill a slot and check availability
	h.Occupy(100, tests[1].pos)
	got = h.IsSlotAvailable(tests[1].pos)
	assert.Equalf(t, tests[1].want, got, tests[1].name)

	// check availability of an edge slot
	got = h.IsSlotAvailable(tests[2].pos)
	assert.Equalf(t, tests[2].want, got, tests[2].name)

	// fill an edge slot and check availability
	h.Occupy(200, tests[3].pos)
	got = h.IsSlotAvailable(tests[3].pos)
	assert.Equalf(t, tests[3].want, got, tests[3].name)

	// check availability of multiple slots after filling some
	h.Occupy(300, 1)
	h.Occupy(400, 2)
	assert.False(t, h.IsSlotAvailable(1), "Slot 1 should not be available after occupation")
	assert.False(t, h.IsSlotAvailable(2), "Slot 2 should not be available after occupation")
	assert.True(t, h.IsSlotAvailable(0), "Slot 0 should be available")
	assert.False(t, h.IsSlotAvailable(3), "Slot 3 should not be available after occupation")
}

func Test_Clear(t *testing.T) {
	h := newTable(4)

	h.Occupy(100, 1)
	preValue := h.Peek(1)
	h.Clear(1)
	postValue := h.Peek(1)
	assert.Lessf(t, postValue, preValue, "hashtable->clear")

	// clear a slot that is already empty
	h.Clear(1)
	postValue = h.Peek(1)
	assert.Equal(t, uint(0), postValue, "Clearing an already empty slot should leave it empty")

	// clear multiple slots
	h.Occupy(200, 0)
	h.Occupy(300, 2)
	h.Clear(0)
	h.Clear(2)
	assert.Equal(t, uint(0), h.Peek(0), "Slot 0 should be empty after clearing")
	assert.Equal(t, uint(0), h.Peek(2), "Slot 2 should be empty after clearing")

	// clear slots at the edges of the hashtable
	h.Occupy(400, 0)
	h.Occupy(500, 3)
	h.Clear(0)
	h.Clear(3)
	assert.Equal(t, uint(0), h.Peek(0), "Slot 0 should be empty after clearing")
	assert.Equal(t, uint(0), h.Peek(3), "Slot 3 should be empty after clearing")

	// clear a slot and verify other slots are unaffected
	h.Occupy(600, 1)
	h.Occupy(700, 2)
	h.Clear(1)
	assert.Equal(t, uint(0), h.Peek(1), "Slot 1 should be empty after clearing")
	assert.Equal(t, uint(700), h.Peek(2), "Slot 2 should remain unaffected")
}
