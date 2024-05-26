package bloom

import (
	"github.com/rtsh13/bfGo/hash"
	"github.com/rtsh13/bfGo/indexing"
)

// inserts the input to the buckets
func (f *filter) Insert(v []byte) {
	f.mu.Lock()
	defer f.mu.Unlock()

	fnvH := hash.Fnv1a(v)
	seg1, seg2 := hash.Murmum3_128(v)

	indexes := indexing.ModuloBiasing([]uint{fnvH, seg1, seg2}, f.size)

	for _, idx := range indexes {
		// set bit to true for the given idx
		f.bucket.Set(idx)
	}
}

/*
verifies the membership of the input in the buckets.
To avoid false +ve, since the likelihood of the collision is
unknown due to variable space, we defer to total membership check.
*/
func (f *filter) MemberOf(v []byte) bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	fnvH := hash.Fnv1a(v)
	seg1, seg2 := hash.Murmum3_128(v)

	indexes := indexing.ModuloBiasing([]uint{fnvH, seg1, seg2}, f.size)

	for _, idx := range indexes {
		if !f.bucket.Test(idx) {
			return false
		}
	}

	return true
}

// flushes all the set bits in the buckets
func (f *filter) Flush() {
	f.bucket = f.bucket.ClearAll()
}

// count (number of set bits). Also known as "popcount" or "population count".
func (f *filter) PopCnt() uint {
	return f.bucket.Count()
}
