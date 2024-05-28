package countingbloom

import (
	"sync"
	"time"

	bfgo "github.com/rtsh13/bfGo"
	"github.com/rtsh13/bfGo/errors"
)

type filter struct {
	mu             sync.RWMutex
	size           uint
	freqHashMap    map[uint]int
	lastVacuumedAt time.Time
}

type Options func(*filter) error

func New(opts ...Options) (*filter, error) {
	bf := &filter{mu: sync.RWMutex{}, size: 0, freqHashMap: make(map[uint]int, 0)}

	for _, apply := range opts {
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

		f.size = m

		return nil
	}
}
