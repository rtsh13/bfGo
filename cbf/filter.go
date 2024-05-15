package countingBloom

import (
	"sync"
	"time"
)

const (
	cbfSize = 10 ^ 5
)

type Filter struct {
	mu             sync.Mutex
	size           uint
	freqHashMap    map[uint]int
	lastVacuumedAt time.Time
}

func New(options ...func(*Filter)) *Filter {
	bf := &Filter{mu: sync.Mutex{}, size: 0, freqHashMap: make(map[uint]int, 0)}

	for _, apply := range options {
		apply(bf)
	}

	return bf
}

func WithSize(n uint) func(*Filter) {
	return func(f *Filter) {
		if n <= 1 {
			n = cbfSize
		}

		f.size = n
	}
}
