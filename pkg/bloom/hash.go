package bloom

import (
	"hash/fnv"

	"github.com/twmb/murmur3"
)

func fnv1a(v []byte) uint64 {
	h := fnv.New64a()
	_, _ = h.Write(v)

	return h.Sum64()
}

func murmum3_128(v []byte) (uint64, uint64) {
	h := murmur3.New128()
	_, _ = h.Write(v)

	return h.Sum128()
}

func (f *Filter) hashToBitsetIdx(hashes ...uint64) (indexes []uint) {
	for _, h := range hashes {
		indexes = append(indexes, uint(h%uint64(f.size)))
	}

	return
}
