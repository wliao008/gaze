package algos

import (
	"github.com/wliao008/mazing/structs"
	"github.com/wliao008/mazing/util"
	"math/rand"
	"time"
	"io"
)

type BackTracking struct {
	Width  int
	Height int
	Cells [][]structs.Cell
}

func (b *BackTracking) Generate() error {
	//fmt.Println("generating %d x %d", b.Width, b.Height)
	b.Cells = make([][]structs.Cell, b.Width)
	for i := 0; i < b.Width; i++ {
		b.Cells[i] = make([]structs.Cell, b.Height)
	}

	//init the flag value to have the bits of 4 directions set.
	for i := 0; i < b.Width; i++ {
		for j := 0; j < b.Height; j++ {
			b.Cells[i][j].Flag = 15
		}
	}

	//start at cell 0,0
	b.doWork(0, 0)
	return nil
}

//Write displays the maze
func (b *BackTracking) Write(w io.Writer) {
	w.Write([]byte("  "))
	for i := 1; i < b.Width; i++ {
		w.Write([]byte(" _"))
	}
	w.Write([]byte("\n"))

	for j := int(0); j < b.Height; j++ {
		w.Write([]byte("|"))
		for h := int(0); h < b.Width; h++ {
			c := b.Cells[h][j]
			if h == b.Width-1 && j == b.Height-1 {
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

//doWork: the recrusive backtracking algorithm
func (b *BackTracking) doWork(x, y int) {
	rand.Seed(time.Now().UTC().UnixNano())
	var directions []interface{}
	directions = append(directions, structs.NORTH)
	directions = append(directions, structs.SOUTH)
	directions = append(directions, structs.EAST)
	directions = append(directions, structs.WEST)
	d := structs.Direction{}
	util.Shuffle(directions)
	for _, direction := range directions {
		dir := direction.(structs.FlagPosition)
		var nextX int = x + d.XDirection(dir)
		var nextY int = y + d.YDirection(dir)
		if nextX >= 0 && nextX < b.Width &&
			nextY >= 0 && nextY < b.Height &&
			!b.Cells[nextX][nextY].IsSet(structs.VISITED) {
			b.carvePassage(dir, &b.Cells[x][y], &b.Cells[nextX][nextY])
			b.doWork(nextX, nextY)
		}
	}
}

func (b *BackTracking) carvePassage(dir structs.FlagPosition, from, to *structs.Cell) {
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
