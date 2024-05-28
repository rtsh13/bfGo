package cuckoo

type hashTable struct {
	slots []uint
}

type table interface {
	Peek(uint) uint
	Occupy(uint, uint)
	IsSlotAvailable(uint) bool
	Clear(uint)
}

func newTable(n uint) table {
	p := make([]uint, n)
	return &hashTable{slots: p}
}

// returns the fingerprint as specific slot of the table/bucket
func (t *hashTable) Peek(pos uint) uint {
	return t.slots[pos]
}

// occupies the fingerprint in a specific slot of the table/bucket
func (t *hashTable) Occupy(fprint, pos uint) {
	t.slots[pos] = fprint
}

// verifies if the slot is available
func (t *hashTable) IsSlotAvailable(pos uint) bool {
	return t.slots[pos] == 0
}

// clears the fingerprint from the slot of the table
func (t *hashTable) Clear(pos uint) {
	t.slots[pos] = 0
}
