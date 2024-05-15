package hash

import (
	"hash/fnv"

	"github.com/twmb/murmur3"
)

func Fnv1a(v []byte) uint {
	h := fnv.New64a()
	h.Write(v)

	return uint(h.Sum64())
}

func Murmum3_128(v []byte) (uint, uint) {
	h := murmur3.New128()
	h.Write(v)

	h1, h2 := h.Sum128()

	return uint(h1), uint(h2)
}
