package structs

import (
	"testing"
	"fmt"
)

func TestNeighbors(t *testing.T) {
	b := initBoard()
	var tests = []struct{
		x uint16
		y uint16
		want int
	}{
		{0,0,2},
		{0,1,3},
		{0,2,2},
		{1,0,3},
		{1,1,4},
		{1,2,3},
		{2,0,2},
		{2,1,3},
		{2,2,2},
	}
	for _, test := range tests {
		c := &b.Cells[test.x][test.y]
		got := b.Neighbors(c)
		if len(got) != test.want {
			t.Errorf("Neighbors(%+v) got %d, want %d", *c, len(got), test.want)
		}
	}
}

func initBoard() *Board {
	b := Board{3, 3, nil}
	b.Init()
	return &b
}

func print(b *Board) {
	for i:=uint16(0); i<b.Width; i++ {
		for j:=uint16(0);j<b.Height;j++{
			fmt.Printf("[%d,%d,%d] ", i,j,b.Cells[i][j].Flag)
		}
		fmt.Println("")
	}
}
