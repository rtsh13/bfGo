package benchmarks

import (
	"fmt"
	"testing"

	bloom "github.com/rtsh13/bfGo/cuckooFilter"
)

func BenchmarkCF_Insert(b *testing.B) {
	benchmarks := []struct {
		inputSize  int
		filterSize int
		slots      int
	}{
		// scenario 1: keep the filter and slot size consistent but input increasing 10x
		{inputSize: 1000, filterSize: 1000, slots: 10},
		{inputSize: 10000, filterSize: 1000, slots: 10},
		{inputSize: 100000, filterSize: 1000, slots: 10},
		{inputSize: 1000000, filterSize: 1000, slots: 10},

		// scenario 2: keep the input and slot size consistent but the filter size increasing 10x
		{inputSize: 1000, filterSize: 1000, slots: 10},
		{inputSize: 1000, filterSize: 10000, slots: 10},
		{inputSize: 1000, filterSize: 100000, slots: 10},
		{inputSize: 1000, filterSize: 1000000, slots: 10},

		// scenario 3: keep the input and filter size consistent but the slot size increasing 10x
		{inputSize: 1000, filterSize: 1000, slots: 10},
		{inputSize: 1000, filterSize: 1000, slots: 100},
		{inputSize: 1000, filterSize: 1000, slots: 1000},
		{inputSize: 1000, filterSize: 1000, slots: 10000},
	}

	for _, bm := range benchmarks {
		f, _ := bloom.New(bloom.WithSize(uint(bm.filterSize), uint(bm.slots)))

		b.Run(fmt.Sprintf("InputSize:[%d]FilterSize:[%d]SlotSize:[%d]", bm.inputSize, bm.filterSize, bm.slots), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				f.Insert(generateRandBytes(bm.inputSize))
			}
		})
	}
}

func Benchmark_CFMembership(b *testing.B) {
	benchmarks := []struct {
		inputSize  int
		filterSize int
		slots      int
	}{
		// scenario 1: keep the filter size consistent but input increasing 10x
		{inputSize: 1000, filterSize: 1000, slots: 10},
		{inputSize: 10000, filterSize: 1000, slots: 10},

		// scenario 2: keep the input size consistent but the filter size increasing 10x
		{inputSize: 1000, filterSize: 1000, slots: 10},
		{inputSize: 1000, filterSize: 10000, slots: 10},
		{inputSize: 1000, filterSize: 100000, slots: 10},
		{inputSize: 1000, filterSize: 1000000, slots: 10},

		// scenario 3: keep the input and filter size consistent but the slot size increasing 10x
		{inputSize: 1000, filterSize: 1000, slots: 10},
		{inputSize: 1000, filterSize: 1000, slots: 100},
		{inputSize: 1000, filterSize: 1000, slots: 1000},
		{inputSize: 1000, filterSize: 1000, slots: 10000},
	}

	for _, bm := range benchmarks {
		f, _ := bloom.New(bloom.WithSize(uint(bm.filterSize), uint(bm.slots)))

		bfInputs := make([][]byte, bm.inputSize)

		for i := 0; i < bm.inputSize; i++ {
			bfInputs[i] = generateRandBytes(bm.inputSize)
			f.Insert(bfInputs[i])
		}

		b.Run(fmt.Sprintf("InputSize:[%d]FilterSize:[%d]SlotSize:[%d]", bm.inputSize, bm.filterSize, bm.slots), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				f.MemberOf([]byte(bfInputs[i%bm.inputSize]))
			}
		})
	}
}