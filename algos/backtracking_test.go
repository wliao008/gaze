package algos

import (
	"testing"

	"github.com/wliao008/gaze"
)

func TestBackTrakcingAlgo_Generate_AllCellsVisited(t *testing.T) {
	k := NewBackTracking(10, 10)
	k.Generate()
	for i := uint16(0); i < k.Board.Width; i++ {
		for j := uint16(0); j < k.Board.Height; j++ {
			if !k.Board.Cells[i][j].IsSet(gaze.VISITED) {
				t.Errorf("Every cell should be visited, but not [%d,%d]", i, j)
			}
		}
	}
}

func TestNewBackTracking(t *testing.T) {
	bt := NewBackTracking(10, 10)
	for _, row := range bt.Board.Cells {
		for _, cell := range row {
			if cell.Flag != 15 {
				t.Errorf("NewBackTrakcing(10,10) should have all flag set to 15, got %d", cell.Flag)
			}
		}
	}
}

func BenchmarkBackTrackingAlgo1000x500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bt := NewBackTracking(1000, 500)
		bt.Generate()
	}
}
