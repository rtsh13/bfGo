package cuckoo

import (
	"math/rand"

	"github.com/rtsh13/bfGo/compress"
	"github.com/rtsh13/bfGo/hash"
	"github.com/rtsh13/bfGo/indexing"
)

// insert in CF occurs by compressing and calculating an available
// slot within the boundaries of buckets. Incase the position is not found,
// we defer open addressing to find alternative slots
func (f *filter) Insert(input []byte) bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	fpVal := compress.CRC64(input)
	fpByte := compress.ToByteArray(fpVal)

	h1 := hash.Fnv1a(fpByte)
	h2 := h1 ^ hash.Fnv1a(fpByte)
	indexes := indexing.ModuloBiasing([]uint{h1, h2}, f.bCount)

	for _, idx := range indexes {
		for j := uint(0); j < f.slotCap; j++ {
			if f.buckets[idx].IsSlotAvailable(j) {
				f.buckets[idx].Occupy(fpVal, j)
				return true
			}
		}
	}

	return f.openAddressing(fpVal, indexes)
}

// verifies if the input is part of the CF.
// exact check is executed by comparing the fingerprints
// to avoid false positives
func (f *filter) MemberOf(input []byte) bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	fpVal := compress.CRC64(input)
	fpByte := compress.ToByteArray(fpVal)

	h1 := hash.Fnv1a(fpByte)
	h2 := h1 ^ hash.Fnv1a(fpByte)
	indexes := indexing.ModuloBiasing([]uint{h1, h2}, f.bCount)

	for _, idx := range indexes {
		for j := uint(0); j < f.slotCap; j++ {
			if fpVal == f.buckets[idx].Peek(j) {
				return true
			}
		}
	}

	return false
}

// purges an input from the CF. before an item is removed,
// the exact fingerprint match is determined to avoid purge-miss
func (f *filter) Delete(input []byte) {
	f.mu.Lock()
	defer f.mu.Unlock()

	fpVal := compress.CRC64(input)
	fpByte := compress.ToByteArray(fpVal)

	h1 := hash.Fnv1a(fpByte)
	h2 := h1 ^ hash.Fnv1a(fpByte)
	indexes := indexing.ModuloBiasing([]uint{h1, h2}, f.bCount)

	for _, idx := range indexes {
		for j := uint(0); j < f.slotCap; j++ {
			if fpVal == f.buckets[idx].Peek(j) {
				f.buckets[idx].Clear(j)
			}
		}
	}
}

// max open addressing iterations for the CF
func (f *filter) Kicks() uint {
	return f.kicks
}

// count of buckets in CF
func (f *filter) BucketPop() uint {
	return f.bCount
}

// during collision upon insertion, we shall reallocate the existing fingerprint
// to a new slot of a new bucket.
// Open addressing will only happen a limited number of times(defined by kick size)
// to avoid excessive computation. Incase an empty slot is not found we return false,
// else true
//
//nolint:gosec // we are picking a deterministic bucket position and slot
func (f *filter) openAddressing(fpVal uint, indexes []uint) bool {
	for i := uint(0); i < f.Kicks(); i++ {
		randIdx := rand.Intn(len(indexes))
		randSlot := rand.Intn(int(f.slotCap))

		idx := indexes[randIdx]

		eFpVal := f.buckets[idx].Peek(uint(randSlot))
		eFpVal, fpVal = fpVal, eFpVal
		f.buckets[idx].Occupy(eFpVal, uint(randSlot))

		fpByte := compress.ToByteArray(fpVal)
		h1 := hash.Fnv1a(fpByte)
		h2 := h1 ^ hash.Fnv1a(fpByte)
		newIndexes := indexing.ModuloBiasing([]uint{h1, h2}, f.bCount)

		for _, idx := range newIndexes {
			for j := uint(0); j < f.slotCap; j++ {
				if f.buckets[idx].IsSlotAvailable(j) {
					f.buckets[idx].Occupy(fpVal, j)
					return true
				}
			}
		}
	}

	return false
}
