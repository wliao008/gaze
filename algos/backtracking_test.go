package algos

import (
	"testing"
)

func BenchmarkBackTrackingAlgo(b *testing.B) {
	bt := BackTracking{50, 25, nil}
	for i := 0; i < b.N; i++ {
		bt.Generate()
	}
}
