package bloomfilter

import (
	"sync"

	"github.com/bits-and-blooms/bitset"
)

type Filter struct {
	mu   sync.Mutex
	size uint
	bSet *bitset.BitSet
}

func New(options ...func(*Filter)) *Filter {
	bf := &Filter{}

	for _, apply := range options {
		apply(bf)
	}

	return bf
}

func WithSize(n uint) func(*Filter) {
	return func(f *Filter) {
		f.size = n
		f.mu = sync.Mutex{}
		bSet := bitset.New(n)
		f.bSet = bSet
	}
}
