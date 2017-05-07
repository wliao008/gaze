package algos

import (
	"testing"
)

func BenchmarkBackTrackingAlgo50x25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bt := NewBackTracking(50, 25)
		bt.Generate()
	}
}

func BenchmarkBackTrackingAlgo100x50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bt := NewBackTracking(100, 50)
		bt.Generate()
	}
}

func BenchmarkBackTrackingAlgo1000x500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bt := NewBackTracking(1000, 500)
		bt.Generate()
	}
}
