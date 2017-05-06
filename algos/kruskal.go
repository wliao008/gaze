package algos

import (
	"github.com/wliao008/mazing/structs"
	"github.com/wliao008/mazing/util"
	"math/rand"
	"time"
)

type Kruskal struct {
	Width  uint16
	Height uint16
	Cells  [][]structs.Cell
	Board structs.Board
}

func NewKruskal(width, height uint16) *Kruskal {
	k := &Kruskal{Board: structs.Board{width, height, nil}}
	k.Board.Cells = make([][]structs.Cell, width)
	for i := uint16(0); i < width; i++ {
		k.Board.Cells[i] = make([]structs.Cell, height)
	}

	for i := uint16(0); i < width; i++ {
		for j := uint16(0); j < height; j++ {
			// set the flag field with 15, which in binary will be 0b00001111,
			// the 4 bits indicates that all 4 walls are up, so the cells are
			// isolated/sealed from each other initially. ex:
			//  _ _
			// |_|_|
			// |_|_|
			//
			k.Board.Cells[i][j].Flag = 15

			// set the relative [x,y] position of the cell on the board
			k.Board.Cells[i][j].X = i
			k.Board.Cells[i][j].Y = j
		}
	}
	return k
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (k *Kruskal) Generate() error {
	stack := util.Stack{}
	stack.Push(&k.Board.Cells[0][0])
	k.Board.Cells[0][0].SetBit(structs.VISITED)

	for !stack.IsEmpty() {
		item := stack.Peek()
		cell := item.(*structs.Cell)
		neighbors := k.Board.Neighbors(cell)
		if len(neighbors) > 0 {
			var idx int = rand.Intn(len(neighbors))
			to := neighbors[idx]
			dir := k.Board.GetDirection(cell, to)
			k.Board.BreakWall(&k.Board.Cells[cell.X][cell.Y], 
				&k.Board.Cells[to.X][to.Y], dir)
			stack.Push(&k.Board.Cells[to.X][to.Y])
		} else {
			_ = stack.Pop()
		}
	}
	return nil
}
