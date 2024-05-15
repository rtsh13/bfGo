package countingBloom

import (
	"time"

	"github.com/rtsh13/bfGo/hash"
	idx "github.com/rtsh13/bfGo/indexing"
)

// inserts the input to the CBF
func (f *Filter) Insert(v []byte) {
	f.mu.Lock()
	defer f.mu.Unlock()

	keys := f.hashToFMap(v)

	for _, k := range keys {
		if _, ok := f.freqHashMap[k]; !ok {
			f.freqHashMap[k] = 1
			continue
		}

		f.freqHashMap[k] = f.freqHashMap[k] + 1
	}
}

/*
verifies the membership of the input in the CBF.
To avoid false +ve, since the likelihood of the collision is
unknown due to variable space, we defer to total membership check.
*/
func (f *Filter) MemberOf(v []byte) bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	keys := f.hashToFMap(v)

	for _, idx := range keys {
		if _, ok := f.freqHashMap[idx]; !ok {
			return false
		}
	}

	return true
}

/*
Delete from CBF leads to decrement of corresponding frequencies
of the keys calculated via hashing the input. To avoid Negative
Counts, we verify the last frequency for the key and delete it
if the count is approaching 0.
*/
func (f *Filter) Delete(v []byte) {
	f.mu.Lock()
	defer f.mu.Unlock()

	keys := f.hashToFMap(v)

	for _, k := range keys {
		if count, ok := f.freqHashMap[k]; ok {
			switch {
			case count == 1:
				delete(f.freqHashMap, k)
			default:
				f.freqHashMap[k] = count - 1
			}
		}
	}
}

// flushes all the keys and their frequencies and updates the vacuum time
func (f *Filter) Flush() {
	f.mu.Lock()
	f.freqHashMap = map[uint]int{}
	f.lastVacuumedAt = time.Now().UTC()
	f.mu.Unlock()
}

/*
Once the hashes are generated by FNV-1a and Murmur3, we need to find
one/many index position(s) that will insert/seek/purge the data/counter.
Since 3 hashes are generated, there will be:

	a. at best 3 corresponding keys
	b. at worst 1 corresponding keys

We use Modulo Biasing to compress the hash value(s) to be within the CBF size.
This is deliberately done to minimize the false +ve resulted by collisions.

	NOTE: However they are still prone to false +ve if the size is too small
*/
func (f *Filter) hashToFMap(v []byte) []uint {
	var (
		m3seg1, m3seg2 = hash.Murmum3_128(v)
		hashes         = make([]uint, 0)
	)

	hashes = append(hashes, hash.Fnv1a(v), m3seg1, m3seg2)

	return idx.ModuloBiasing(hashes, f.size)
}
