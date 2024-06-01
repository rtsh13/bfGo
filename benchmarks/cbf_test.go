package benchmarks

import (
	"fmt"
	"testing"

	bloom "github.com/rtsh13/bfGo/cbf"
)

func BenchmarkCBF_Insert(b *testing.B) {
	benchmarks := []struct {
		inputSize  int
		filterSize int
	}{
		// scenario 1: keep the filter size consistent but input increasing 10x
		{inputSize: 1000, filterSize: 1000},
		{inputSize: 10000, filterSize: 1000},
		{inputSize: 100000, filterSize: 1000},
		{inputSize: 1000000, filterSize: 1000},

		// scenario 2: keep the input size consistent but the filter size increasing 10x
		{inputSize: 1000, filterSize: 1000},
		{inputSize: 1000, filterSize: 10000},
		{inputSize: 1000, filterSize: 100000},
		{inputSize: 1000, filterSize: 1000000},
	}

	for _, bm := range benchmarks {
		b.Run(fmt.Sprintf("InputSize:[%d]FilterSize:[%d]", bm.inputSize, bm.filterSize), func(b *testing.B) {
			f, _ := bloom.New(bloom.WithSize(uint(bm.filterSize)))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				f.Insert(generateRandBytes(bm.inputSize))
			}
		})
	}
}

func Benchmark_CBFMembership(b *testing.B) {
	benchmarks := []struct {
		inputSize  int
		filterSize int
	}{
		// scenario 1: keep the filter size consistent but input increasing 10x
		{inputSize: 1000, filterSize: 1000},
		{inputSize: 10000, filterSize: 1000},

		// scenario 2: keep the input size consistent but the filter size increasing 10x
		{inputSize: 1000, filterSize: 1000},
		{inputSize: 1000, filterSize: 10000},
		{inputSize: 1000, filterSize: 100000},
		{inputSize: 1000, filterSize: 1000000},
	}

	for _, bm := range benchmarks {
		f, _ := bloom.New(bloom.WithSize(uint(bm.filterSize)))

		bfInputs := make([][]byte, bm.inputSize)

		for i := 0; i < bm.inputSize; i++ {
			bfInputs[i] = generateRandBytes(bm.inputSize)
			f.Insert(bfInputs[i])
		}

		b.Run(fmt.Sprintf("InputSize:[%d]FilterSize:[%d]", bm.inputSize, bm.filterSize), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				f.MemberOf([]byte(bfInputs[i%bm.inputSize]))
			}
		})
	}
}
