package algos

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/wliao008/gaze"
)

type PrimCircle struct {
	Name  string
	Board gaze.BoardCircle
}

func NewPrimCircle(height, width uint16) *PrimCircle {
	p := &PrimCircle{Name: "prim algorithm (circle)", Board: gaze.BoardCircle{height, width, nil}}
	p.Board.Init()
	return p
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (p *PrimCircle) Generate() error {
	stack := gaze.Stack{}
	c := &p.Board.Cells[0][0]
	flats := p.Board.FlattenCells(c)
	first := flats[0]
	stack.Push(first)
	first.SetBit(gaze.VISITED)
	fmt.Printf("flats of %+v\n", c)
	for _, n := range flats {
		fmt.Printf("%+v\n", n)
	}
	//p.Board.Cells[0][0].SetBit(gaze.VISITED)
	count := 0
	fmt.Printf("[%d] stack: %+v\n", count, stack)
	for !stack.IsEmpty() {
		item := stack.Peek()
		cell := item.(*gaze.Cell)
		neighbors := p.Board.Neighbors(cell)
		fmt.Printf("neighbors of %+v\n", cell)
		for _, n := range neighbors {
			fmt.Printf("%+v\n", n)
		}
		if len(neighbors) > 0 {
			var idx int = rand.Intn(len(neighbors))
			to := neighbors[idx]
			dir := p.Board.GetDirection(cell, to)
			p.Board.BreakWall(cell, to, dir)
			fmt.Printf("from %+v to %+v\n", cell, to)
			stack.Push(to)
		} else {
			_ = stack.Pop()
		}
		count++
		fmt.Printf("[%d] stack: %+v\n", count, stack)
		if count == 100 {
			break
		}
	}
	return nil
}

func (p *PrimCircle) GenerateNew() error {
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
