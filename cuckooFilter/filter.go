package cuckoo

import (
	"sync"

	bfgo "github.com/rtsh13/bfGo"
	"github.com/rtsh13/bfGo/errors"
)

const (
	maxKicks        = 50
	bCount          = 10 ^ 4
	bSlots          = 4
	bucketSizeError = "incorrect bucket or slot size provided"
	maxKickError    = "incorrect max kicks provided"
)

type Options func(*filter) error

type filter struct {
	bCount  uint         // count of buckets in the filter
	slotCap uint         // maximum allocated slots for a bucket
	buckets []table      // hashtable
	kicks   uint         // max allowed reallocation
	mu      sync.RWMutex // concurrency management
}

func New(opts ...Options) (*filter, error) {
	cf := &filter{
		bCount: bCount, slotCap: bSlots, kicks: maxKicks,
		mu: sync.RWMutex{}, buckets: []table{},
	}

	for _, apply := range opts {
		if err := apply(cf); err != nil {
			return nil, err
		}
	}

	return cf, nil
}

/*
1. m represents number of buckets in the filter

2. n represents number of slots for each bucket
*/
func WithSize(m uint, n uint) Options {
	return func(f *filter) error {
		if m <= bfgo.MinimumBloomFSize {
			return errors.BloomSize{Size: m}
		}

		if n < 1 {
			return errors.BloomSlots{Slot: n}
		}

		f.buckets = make([]table, m)
		for i := range f.buckets {
			f.buckets[i] = newTable(n)
		}

		f.slotCap = n
		f.bCount = m

		return nil
	}
}

// custom defined CF kick that allow maximum possible open addressing iterations/reallocations
func WithKicks(k uint) Options {
	return func(f *filter) error {
		if k < 1 {
			return errors.BloomKicks{Kick: k}
		}

		f.kicks = k

		return nil
	}
}
