package solvers

import (
	"github.com/wliao008/gaze"
)

type DeadEndFiller struct {
	Board *gaze.Board
}

func NewDeadEndFiller() *DeadEndFiller {
	def := &DeadEndFiller{}
	return def
}

func (def *DeadEndFiller) Solve() {
	// get initial list of dead ends
	// while there are more dead ends
	//    mark current cell as a dead end
	//    find dead neighbors of the current cell and add to list
	stack := &gaze.Stack{}
	def.Board.DeadEnds(stack)
	c := stack.Pop()
	for c != nil {
		cell := c.(*gaze.Cell)
		cell.SetBit(gaze.DEAD)
		def.Board.DeadNeighbors(cell, stack)
		c = stack.Pop()
	}
}
