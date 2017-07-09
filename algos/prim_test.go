package algos

import (
	"testing"
	"github.com/wliao008/mazing/structs"
	"math/rand"
	"time"
)

func TestPrim(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	count := rand.Intn(50) + 1
	for c := 0; c < count; c++ {
		k := NewPrim(uint16(rand.Intn(100)+1), uint16(rand.Intn(100)+1))
		k.Generate()
		for h := uint16(0); h < k.Board.Height; h++ {
			for w := uint16(0); w < k.Board.Width; w++ {
				if !k.Board.Cells[h][w].IsSet(structs.VISITED) {
					t.Errorf("Prim, every cell should be visited, but not [%d,%d]", h, w)
				}
			}
		}
	}
}

func TestNewPrim(t *testing.T) {
	k := NewPrim(10, 10)
	for _, row := range k.Board.Cells {
		for _, cell := range row {
			if cell.Flag != 15 {
				t.Errorf("NewPrim(), every celll should have flag set to 15, got %d", cell.Flag)
			}
		}
	}
}

func BenchmarkPrimAlgo_1000x500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := NewPrim(1000, 500)
		k.Generate()
	}
}

func BenchmarkPrimAlgo_100x50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := NewPrim(100, 50)
		k.GenerateNew()
	}
}

