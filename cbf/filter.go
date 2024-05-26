package countingbloom

import (
	"sync"
	"time"
)

const (
	cbfSize = 10 ^ 5
)

type filter struct {
	mu             sync.RWMutex
	size           uint
	freqHashMap    map[uint]int
	lastVacuumedAt time.Time
}

func New(options ...func(*filter)) *filter {
	bf := &filter{mu: sync.RWMutex{}, size: 0, freqHashMap: make(map[uint]int, 0)}

	for _, apply := range options {
		apply(bf)
	}

	return bf
}

func WithSize(n uint) func(*filter) {
	return func(f *filter) {
		if n <= 1 {
			n = cbfSize
		}

		f.size = n
	}
}
