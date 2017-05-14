package algos

import (
	"testing"
	"github.com/wliao008/mazing/structs"
	"math/rand"
	"time"
)

func TestKruskalAlgo_Generate_AllCellsVisited(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	count := rand.Intn(50) + 1
	for c := 0; c < count; c++ {
		k := NewKruskal(uint16(rand.Intn(100)+1), uint16(rand.Intn(100)+1))
		k.Generate()
		for i := uint16(0); i < k.Board.Width; i++ {
			for j := uint16(0); j < k.Board.Height; j++ {
				if !k.Board.Cells[i][j].IsSet(structs.VISITED) {
					t.Errorf("Every cell should be visited, but not [%d,%d]", i, j)
				}
			}
		}
	}
}

func TestNewKruskal(t *testing.T) {
	k := NewKruskal(10, 10)
	for _, row := range k.Board.Cells {
		for _, cell := range row {
			if cell.Flag != 15 {
				t.Errorf("NewKruskal(), every celll should have flag set to 15, got %d", cell.Flag)
			}
		}
	}
}

func BenchmarkKruskalAlgo50x25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := NewKruskal(50, 25)
		k.Generate()
	}
}

func BenchmarkKruskalAlgo100x50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := NewKruskal(100, 50)
		k.Generate()
	}
}

func BenchmarkKruskalAlgo1000x500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := NewKruskal(1000, 500)
		k.Generate()
	}
}