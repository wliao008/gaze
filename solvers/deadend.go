package solvers

import (
	"fmt"
	"github.com/wliao008/mazing/structs"
	"github.com/wliao008/mazing/util"
)

type DeadEndFiller struct {
	Board *structs.Board
}

func NewDeadEndFiller() *DeadEndFiller {
	def := &DeadEndFiller{}
	return def
}

func (def *DeadEndFiller) Solve() {
	fmt.Println("solving")
	// get initial list of dead ends
	// while there are more dead ends
	//    mark current cell as a dead end
	//    find dead neighbors of the current cell and add to list
	stack := def.InitialDeadEnds()
	fmt.Println(stack)
	c := stack.Pop()
	for c != nil {	
		c.(*structs.Cell).SetBit(structs.DEAD)
		fmt.Println(c)
		c = stack.Pop()
	}
}

func (def *DeadEndFiller) InitialDeadEnds() *util.Stack {
	stack := &util.Stack{}
	c := &structs.Cell{}
	walls := 0
	for h := uint16(0); h < def.Board.Height; h++ {
		for w := uint16(0); w < def.Board.Width; w++ {
			c = &def.Board.Cells[h][w]	
			walls = 0
			if c.IsSet(structs.EAST) {
				walls += 1
			}
			if c.IsSet(structs.SOUTH) {
				walls += 1
			}
			if c.IsSet(structs.WEST) {
				walls += 1
			}
			if c.IsSet(structs.NORTH) {
				walls += 1
			}
			if walls == 3 {
				stack.Push(c)
			}
		}
	}
	return stack
}

func (def *DeadEndFiller) DeadNeighbors(c *structs.Cell, stack *util.Stack) {

}
