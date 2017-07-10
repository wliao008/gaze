package algos

import (
	"github.com/wliao008/mazing/structs"
	"github.com/wliao008/mazing/util"
	"math/rand"
	"time"
)

type Prim struct {
	Name string
	Board structs.Board
}

func NewPrim(height, width uint16) *Prim {
	p := &Prim{Name: "prim algorithm", Board: structs.Board{height, width, nil}}
	p.Board.Init()
	return p
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (p *Prim) Generate() error {
	stack := util.Stack{}
	stack.Push(&p.Board.Cells[0][0])
	p.Board.Cells[0][0].SetBit(structs.VISITED)
	for !stack.IsEmpty() {
		item := stack.Peek()
		cell := item.(*structs.Cell)
		neighbors := p.Board.Neighbors(cell)
		if len(neighbors) > 0 {
			var idx int = rand.Intn(len(neighbors))
			to := neighbors[idx]
			dir := p.Board.GetDirection(cell, to)
			p.Board.BreakWall(&p.Board.Cells[cell.X][cell.Y], 
				&p.Board.Cells[to.X][to.Y], dir)
			stack.Push(&p.Board.Cells[to.X][to.Y])
		} else {
			_ = stack.Pop()
		}
	}
	return nil
}

func (p *Prim) GenerateNew() error {
	stack := util.Stack{}
	stack.Push(&p.Board.Cells[0][0])
	p.Board.Cells[0][0].SetBit(structs.VISITED)
	var item interface{}
	cell := &structs.Cell{}
	for !stack.IsEmpty() {
		item = stack.Peek()
		cell = item.(*structs.Cell)
		neighbors := p.Board.Neighbors(cell)
		if len(neighbors) > 0 {
			var idx int = rand.Intn(len(neighbors))
			to := neighbors[idx]
			dir := p.Board.GetDirection(cell, to)
			p.Board.BreakWall(&p.Board.Cells[cell.X][cell.Y], 
				&p.Board.Cells[to.X][to.Y], dir)
			stack.Push(&p.Board.Cells[to.X][to.Y])
		} else {
			_ = stack.Pop()
		}
	}
	return nil
}
