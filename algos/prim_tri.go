package algos

import (
	"math/rand"
	"time"

	"github.com/wliao008/gaze"
)

type PrimTriangle struct {
	Name  string
	Board gaze.BoardTri
}

func NewPrimTriangle(height, width uint16) *PrimTriangle {
	p := &PrimTriangle{Name: "prim algorithm", Board: gaze.BoardTri{height, width, nil}}
	p.Board.Init()
	return p
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (p *PrimTriangle) Generate() error {
	stack := gaze.Stack{}
	stack.Push(&p.Board.Cells[0][0])
	p.Board.Cells[0][0].SetBit(gaze.VISITED)
	for !stack.IsEmpty() {
		item := stack.Peek()
		cell := item.(*gaze.Cell)
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

func (p *PrimTriangle) GenerateNew() error {
	stack := gaze.Stack{}
	stack.Push(&p.Board.Cells[0][0])
	p.Board.Cells[0][0].SetBit(gaze.VISITED)
	var item interface{}
	cell := &gaze.Cell{}
	for !stack.IsEmpty() {
		item = stack.Peek()
		cell = item.(*gaze.Cell)
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
