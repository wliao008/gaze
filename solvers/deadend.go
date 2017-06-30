package solvers

import (
	"fmt"
	"github.com/wliao008/mazing/structs"
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
}

func (def *DeadEndFiller) InitialDeadEnds() {
	for h := uint16(0); h < def.Board.Height; h++ {
		for w := uint16(0); w < def.Board.Width; w++ {
			
		}
	}
}
