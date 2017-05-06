package algos

import (
	"fmt"
	"github.com/wliao008/mazing/structs"
	"github.com/wliao008/mazing/util"
	"io"
	"math/rand"
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

func (k *Kruskal) Write2(w io.Writer) {
	for i := uint16(0); i < k.Width; i++ {
		for j := uint16(0); j < k.Height; j++ {
			w.Write([]byte(fmt.Sprintf("[%d, %d, %t] ", i, j, k.Cells[i][j].IsSet(structs.VISITED))))
		}
		w.Write([]byte("\n"))
	}
}

func (k *Kruskal) Write(w io.Writer) {
	w.Write([]byte("  "))
	for i := uint16(1); i < k.Width; i++ {
		w.Write([]byte(" _"))
	}
	w.Write([]byte("\n"))

	for j := uint16(0); j < k.Height; j++ {
		w.Write([]byte("|"))
		for h := uint16(0); h < k.Width; h++ {
			c := k.Cells[h][j]
			if h == k.Width-1 && j == k.Height-1 {
				w.Write([]byte(" |"))
				break
			}
			if c.IsSet(structs.SOUTH) {
				w.Write([]byte("_"))
			} else {
				w.Write([]byte(" "))
			}

			if c.IsSet(structs.EAST) {
				w.Write([]byte("|"))
			} else {
				w.Write([]byte(" "))
			}
		}
		w.Write([]byte("\n"))
	}
}

//neighbors return the neighboring unvisted cells of a cell
func (k *Kruskal) neighbors(c *structs.Cell) []*structs.Cell {
	cells := []*structs.Cell{}
	x, y, ok := k.getNeighborPos(c.X+1, c.Y)
	if ok {
		//fmt.Println("south: %v", *cell)
		cells = append(cells, &k.Cells[x][y])
	}
	x, y, ok = k.getNeighborPos(c.X-1, c.Y)
	if ok {
		//fmt.Println("south: %v", *cell)
		cells = append(cells, &k.Cells[x][y])
	}
	x, y, ok = k.getNeighborPos(c.X, c.Y+1)
	if ok {
		//fmt.Println("south: %v", *cell)
		cells = append(cells, &k.Cells[x][y])
	}
	x, y, ok = k.getNeighborPos(c.X, c.Y-1)
	if ok {
		//fmt.Println("south: %v", *cell)
		cells = append(cells, &k.Cells[x][y])
	}
	return cells
}

func (k *Kruskal) getNeighborPos(x, y uint16) (uint16, uint16, bool) {
	if x >= 0 && x < k.Width &&
		y >= 0 && y < k.Height &&
		!k.Cells[x][y].IsSet(structs.VISITED) {
		return x, y, true
	}
	return 0, 0, false
}

func (k *Kruskal) Generate() error {
	stack := util.Stack{}
	k.initCells()
	stack.Push(&k.Cells[0][0])
	k.Cells[0][0].SetBit(structs.VISITED)
	i := 25
	for !stack.IsEmpty() {
		item := stack.Peek()
		cell := item.(*structs.Cell)
		neighbors := k.neighbors(cell)
		//fmt.Println(neighbors)
		if len(neighbors) > 0 {
			var idx int = rand.Intn(len(neighbors))
			to := neighbors[idx]
			dir := k.getDirection(cell, to)
			k.carvePassage(dir, &k.Cells[cell.X][cell.Y], &k.Cells[to.X][to.Y])
			//fmt.Println("%v from %v to %v, stack count: %d", dir, cell, to, stack.Count)
			//toNeighbors := k.neighbors(&to)
			//fmt.Println("to neighbors: %v", toNeighbors)
			stack.Push(&k.Cells[to.X][to.Y])
			//fmt.Println("pushed %v, count=%d", k.Cells[to.X][to.Y], stack.Count)
			i -= 1
		} else {
			_ = stack.Pop()
			//fmt.Println("popped %v, count=%d", poppedItem, stack.Count)
		}
		//fmt.Println(cell)
	}
	//fmt.Println(i)
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
