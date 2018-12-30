package algos

import (
	"testing"
)

func TestPrimCircle(t *testing.T) {
	for c := 0; c < 1; c++ {
		k := NewPrimCircle(6, 6)
		//k.Generate()
		c := k.Board.Cells[0][0].Left.Left
		neighbors := k.Board.Neighbors(c)
		if len(neighbors) != 3 {
			t.Errorf("should have 3 neighbors")
		}
		// for h := uint16(0); h < k.Board.Height; h++ {
		// 	for w := uint16(0); w < k.Board.Width; w++ {
		// 		if !k.Board.Cells[h][w].IsSet(gaze.VISITED) {
		// 			t.Errorf("Prim, every cell should be visited, but not [%d,%d]", h, w)
		// 		}
		// 	}
		// }
	}
}
