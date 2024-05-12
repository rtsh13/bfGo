package bloom

import (
	"sync"

	"github.com/bits-and-blooms/bitset"
)

type Filter struct {
	mu     sync.Mutex     // exclusive lock for reads and writes
	size   uint           // length of the bucket
	bucket *bitset.BitSet // equivalent hashmap
}

func New(options ...func(*Filter)) *Filter {
	bf := &Filter{mu: sync.Mutex{}, size: 0, bucket: nil}

	for _, apply := range options {
		apply(bf)
	}

	return bf
}

func WithSize(n uint) func(*Filter) {
	return func(f *Filter) {
		f.size = n
		bSet := bitset.New(n)
		f.bucket = bSet
	}
}
