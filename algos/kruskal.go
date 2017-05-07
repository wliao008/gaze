package algos

import (
	"github.com/wliao008/mazing/structs"
	"github.com/wliao008/mazing/util"
	"math/rand"
	"time"
)

type Kruskal struct {
	Board structs.Board
}

func NewKruskal(width, height uint16) *Kruskal {
	k := &Kruskal{Board: structs.Board{width, height, nil}}
	k.Board.Init()
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
