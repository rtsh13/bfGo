package bloom

import (
	"sync"

	"github.com/bits-and-blooms/bitset"
)

type filter struct {
	mu     sync.RWMutex   // exclusive lock for reads and writes
	size   uint           // count of buckets
	bucket *bitset.BitSet // equivalent hashmap
}

func New(options ...func(*filter)) *filter {
	bf := &filter{mu: sync.RWMutex{}, size: 0, bucket: nil}

	for _, apply := range options {
		apply(bf)
	}

	return bf
}

func WithSize(n uint) func(*filter) {
	return func(f *filter) {
		f.size = n
		bSet := bitset.New(n)
		f.bucket = bSet
	}
}
