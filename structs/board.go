package structs

import (
	"io"
)

type Board struct {
	Width, Height uint16
	Cells         [][]Cell
}

func (b *Board) Init() {
	b.Cells = make([][]Cell, b.Width)
	for i := uint16(0); i < b.Width; i++ {
		b.Cells[i] = make([]Cell, b.Height)
	}

	for i := uint16(0); i < b.Width; i++ {
		for j := uint16(0); j < b.Height; j++ {
			// set the flag field with 15, which in binary will be 0b00001111,
			// the 4 bits indicates that all 4 walls are up, so the cells are
			// isolated/sealed from each other initially. ex:
			//  _ _ 
			// |_|_| 
			// |_|_|  
			// 
			b.Cells[i][j].Flag = 15

			// set the relative [x,y] position of the cell on the board
			b.Cells[i][j].X = i 
			b.Cells[i][j].Y = j
		}
	}
}

// Neighbors find the neighboring cells of the current cell
func (b *Board) Neighbors(c *Cell) []*Cell {
	result := []*Cell{}
	if ok, cell := b.getNeighbor(c.X+1, c.Y); ok {
		result = append(result, cell)
	}

	if ok, cell := b.getNeighbor(c.X-1, c.Y); ok {
		result = append(result, cell)
	}

	if ok, cell := b.getNeighbor(c.X, c.Y+1); ok {
		result = append(result, cell)
	}

	if ok, cell := b.getNeighbor(c.X, c.Y-1); ok {
		result = append(result, cell)
	}
	return result
}

func (b *Board) GetDirection(from, to *Cell) FlagPosition {
	if from.X < to.X {
		return EAST
	}
	if from.X > to.X {
		return WEST
	}
	if from.Y < to.Y {
		return SOUTH
	}
	if from.Y > to.Y {
		return NORTH
	}
	//TODO: This is really an error case here
	return EAST
}

func (b *Board) BreakWall(from, to *Cell, direction FlagPosition) {
	from.SetBit(VISITED)
	to.SetBit(VISITED)

	switch direction {
	case EAST:
		from.ClearBit(EAST)
		to.ClearBit(WEST)
	case SOUTH:
		from.ClearBit(SOUTH)
		to.ClearBit(NORTH)
	case WEST:
		from.ClearBit(WEST)
		to.ClearBit(EAST)
	case NORTH:
		from.ClearBit(NORTH)
		to.ClearBit(SOUTH)
	}
}

func (b *Board) Write(w io.Writer) {
	w.Write([]byte("  "))
	for i := uint16(1); i < b.Width; i++ {
		w.Write([]byte(" _"))
	}
	w.Write([]byte("\n"))

	for j := uint16(0); j < b.Height; j++ {
		w.Write([]byte("|"))
		for h := uint16(0); h < b.Width; h++ {
			c := b.Cells[h][j]
			if h == b.Width-1 && j == b.Height-1 {
				w.Write([]byte(" |"))
				break
			}
			if c.IsSet(SOUTH) {
				w.Write([]byte("_"))
			} else {
				w.Write([]byte(" "))
			}

			if c.IsSet(EAST) {
				w.Write([]byte("|"))
			} else {
				w.Write([]byte(" "))
			}
		}
		w.Write([]byte("\n"))
	}
}

func (b *Board) getNeighbor(x, y uint16) (bool, *Cell) {
	if x >= 0 && x < b.Width &&
		y >= 0 && y < b.Height &&
		!b.Cells[x][y].IsSet(VISITED) {
		return true, &b.Cells[x][y]
	}
	return false, nil
}
