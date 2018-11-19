package algos

import (
	"math/rand"
	"testing"
	"time"

	"github.com/wliao008/gaze"
)

func TestPrimTriangle(t *testing.T) {
	t.Skip("skipping this test for now")
	rand.Seed(time.Now().UTC().UnixNano())
	count := rand.Intn(50) + 1
	for c := 0; c < count; c++ {
		k := NewPrimTriangle(uint16(rand.Intn(100)+1), uint16(rand.Intn(100)+1))
		k.Generate()
		for h := uint16(0); h < k.Board.Height; h++ {
			for w := uint16(0); w < k.Board.Width; w++ {
				if !k.Board.Cells[h][w].IsSet(gaze.VISITED) {
					t.Errorf("Prim, every cell should be visited, but not [%d,%d]", h, w)
				}
			}
		}
	}
}

func TestNewPrimTriangle(t *testing.T) {
	k := NewPrimTriangle(10, 10)
	for _, row := range k.Board.Cells {
		for _, cell := range row {
			if cell.Flag != 15 && cell.Flag != 527 {
				t.Errorf("NewPrim(), every celll should have flag set to either 15 or 527, got %d", cell.Flag)
			}
		}
	}
}

func BenchmarkPrimTriangleAlgo_1000x500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := NewPrimTriangle(1000, 500)
		k.Generate()
	}
}

func BenchmarkPrimTriangleAlgo_100x50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := NewPrimTriangle(100, 50)
		k.Generate()
	}
}
