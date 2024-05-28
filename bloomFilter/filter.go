package bloom

import (
	"sync"

	"github.com/bits-and-blooms/bitset"
	bfgo "github.com/rtsh13/bfGo"
	"github.com/rtsh13/bfGo/errors"
)

type Options func(*filter) error

type filter struct {
	mu     sync.RWMutex   // exclusive lock for reads and writes
	m      uint           // count of buckets
	bucket *bitset.BitSet // equivalent hashmap
}

func New(options ...Options) (*filter, error) {
	bf := &filter{mu: sync.RWMutex{}, m: 0, bucket: nil}

	for _, apply := range options {
		if err := apply(bf); err != nil {
			return nil, err
		}
	}

	return bf, nil
}

func WithSize(m uint) Options {
	return func(f *filter) error {
		if m <= bfgo.MinimumBloomFSize {
			return errors.BloomSize{Size: m}
		}

		f.m = m
		f.bucket = bitset.New(m)

		return nil
	}
}
