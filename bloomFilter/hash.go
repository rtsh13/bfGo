package bloom

import (
	"hash/fnv"

	"github.com/twmb/murmur3"
)

func fnv1a(v []byte) uint64 {
	h := fnv.New64a()
	h.Write(v)

	return h.Sum64()
}

func murmum3_128(v []byte) (uint64, uint64) {
	h := murmur3.New128()
	h.Write(v)

	return h.Sum128()
}

/*
Once the hashes are generated by FNV-1a and Murmur3, we need to find one/many
index position(s) that will flush/seek the data. This is deliberately done
to minimize the false postives resulted by collisions.

	index = {hash(i) % bitset_size}

For each hash h, it calculates the modulo operation to ensure that the index
falls within the range of the bitset size
*/
func (f *Filter) hashToBitsetIdx(v []byte) []uint {
	var (
		indexes        = make([]uint, 0)
		m3seg1, m3seg2 = murmum3_128(v)
	)

	indexes = append(indexes, uint(fnv1a(v)%uint64(f.size)),
		uint(m3seg1%uint64(f.size)), uint(m3seg2%uint64(f.size)))

	return indexes
}