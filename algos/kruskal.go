package algos

import (
	"fmt"
	"github.com/wliao008/mazing/structs"
	"github.com/wliao008/mazing/util"
	"math/rand"
	"io"
	"time"
)

type Kruskal struct {
	Width  uint16
	Height uint16
	Cells  [][]structs.Cell
}

var directions2 []interface{}

func init() {
	directions2 = append(directions2, structs.NORTH, structs.SOUTH, structs.EAST, structs.WEST)
	rand.Seed(time.Now().UTC().UnixNano())
}

func (k *Kruskal) initCells() {
	k.Cells = make([][]structs.Cell, k.Width)
	for i := uint16(0); i < k.Width; i++ {
		k.Cells[i] = make([]structs.Cell, k.Height)
	}

	for i := uint16(0); i < k.Width; i++ {
		for j := uint16(0); j < k.Height; j++ {
			k.Cells[i][j].Flag = 15 //set 4 walls
			k.Cells[i][j].X = i
			k.Cells[i][j].Y = j
		}
	}
}

func (k *Kruskal) Write(w io.Writer) {
	for i := uint16(0); i < k.Width; i++ {
		for j := uint16(0); j < k.Height; j++ {
			w.Write([]byte(fmt.Sprintf("[%d, %d] ", i, j)))
		}
		w.Write([]byte("\n"))
	}

	//test neighbor of a cell
	cells := k.neighbors(&k.Cells[0][0])
	fmt.Println("cells: %v", cells)
	fmt.Println(cells[0])
	for _, cell := range cells {
		w.Write([]byte(fmt.Sprintf("[%d, %d]", cell.X, cell.Y)))
	}
}

//neighbors return the neighboring unvisted cells of a cell
func (k *Kruskal) neighbors(c *structs.Cell) []structs.Cell {
	cells := []structs.Cell{}
	cell := &structs.Cell{}
	cell = k.getNeighbor(c.X+1, c.Y)
	if cell != nil {
		//fmt.Println("south: %v", *cell)
		cells = append(cells, *cell)
	}
	cell = k.getNeighbor(c.X-1, c.Y)
	if cell != nil {
		//fmt.Println("north: %v", *cell)
		cells = append(cells, *cell)
	}
	cell = k.getNeighbor(c.X, c.Y+1)
	if cell != nil {
		//fmt.Println("east: %v", *cell)
		cells = append(cells, *cell)
	}
	cell = k.getNeighbor(c.X, c.Y-1)
	if cell != nil {
		//fmt.Println("west: %v", *cell)
		cells = append(cells, *cell)
	}
	return cells
}

func (k *Kruskal) getNeighbor(x, y uint16) *structs.Cell {
	if x >= 0 && x < k.Width && y >= 0 && y < k.Height && !k.Cells[x][y].IsSet(structs.VISITED) {
		return &k.Cells[x][y]
	}
	return nil
}

func (k *Kruskal) Generate() error {
	stack := util.Stack{}
	k.initCells()
	stack.Push(k.Cells[0][0])
	k.Cells[0][0].SetBit(structs.VISITED)
	i := 25
	for !stack.IsEmpty() && i != 0 {
		item := stack.Peek()
		cell := item.(structs.Cell)
		neighbors := k.neighbors(&cell)
		if len(neighbors) > 0 {
			var idx int = rand.Intn(len(neighbors))
			to := neighbors[idx]
			dir := k.getDirection(&cell, &to)
			k.carvePassage(dir, &cell, &to)
			fmt.Println("%v from %v to %v, stack count: %d", dir, cell, to, stack.Count)
			//toNeighbors := k.neighbors(&to)
			//fmt.Println("to neighbors: %v", toNeighbors)
			stack.Push(to)
			i -= 1
		} else {
			stack.Pop()
		}
		//fmt.Println(cell)
	}
	return nil
}

func (k *Kruskal) getDirection(from, to *structs.Cell) structs.FlagPosition {
	if from.X < to.X {
		return structs.EAST
	} 
	if from.X > to.X {
		return structs.WEST
	}
	if from.Y < to.Y {
		return structs.SOUTH
	}
	if from.Y > to.Y {
		return structs.NORTH
	}
	return structs.EAST
}

func (k *Kruskal) carvePassage(dir structs.FlagPosition, from, to *structs.Cell) {
	from.SetBit(structs.VISITED)
	to.SetBit(structs.VISITED)

	switch dir {
	case structs.NORTH:
		from.ClearBit(structs.NORTH)
		to.ClearBit(structs.SOUTH)
	case structs.SOUTH:
		from.ClearBit(structs.SOUTH)
		to.ClearBit(structs.NORTH)
	case structs.EAST:
		from.ClearBit(structs.EAST)
		to.ClearBit(structs.WEST)
	case structs.WEST:
		from.ClearBit(structs.WEST)
		to.ClearBit(structs.EAST)
	}
}
