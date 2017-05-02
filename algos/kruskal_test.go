package algos

import (
	"testing"
)

func BenchmarkKruskalAlgo(b *testing.B) {
	bt := Kruskal{100, 50, nil}
	for i := 0; i < b.N; i++ {
		bt.Generate()
	}
}
