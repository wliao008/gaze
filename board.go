package gaze

import (
	//"fmt"
	"fmt"
	"io"
)

type Board struct {
	Height uint16 // row
	Width  uint16 // col
	Cells  [][]Cell
	Masks  [][]int
}

func (b *Board) Init() {
	b.Cells = make([][]Cell, b.Height)
	b.Masks = make([][]int, b.Height)
	for i := uint16(0); i < b.Height; i++ {
		b.Cells[i] = make([]Cell, b.Width)
		b.Masks[i] = make([]int, b.Width)
	}

	for h := uint16(0); h < b.Height; h++ {
		for w := uint16(0); w < b.Width; w++ {
			// set the flag field with 15, which in binary will be 0b00001111,
			// the 4 bits indicates that all 4 walls are up, so the cells are
			// isolated/sealed from each other initially. ex:
			//  _ _
			// |_|_|
			// |_|_|
			//
			b.Cells[h][w].Flag = 15

			// set the relative [x,y] position of the cell on the board
			b.Cells[h][w].X = h
			b.Cells[h][w].Y = w

			b.Cells[h][w].Parent = nil
			b.Cells[h][w].Left = nil
			b.Cells[h][w].Right = nil
		}
	}
}

func (b *Board) ApplyMasks() {
	for h := uint16(0); h < b.Height; h++ {
		for w := uint16(0); w < b.Width; w++ {
			if b.Masks[h][w] == 1 {
				b.Cells[h][w].SetBit(DEAD)
			}
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

// Neighbors find the neighboring cells of the current cell
func (b *Board) CornerNeighbors(c *Cell) []*Cell {
	result := []*Cell{}
	if ok, cell := b.getNeighbor(c.X+1, c.Y+1); ok {
		result = append(result, cell)
	}

	if ok, cell := b.getNeighbor(c.X+1, c.Y-1); ok {
		result = append(result, cell)
	}

	if ok, cell := b.getNeighbor(c.X-1, c.Y+1); ok {
		result = append(result, cell)
	}

	if ok, cell := b.getNeighbor(c.X-1, c.Y-1); ok {
		result = append(result, cell)
	}
	return result
}

func (b *Board) GetDirection(from, to *Cell) FlagPosition {
	// X denotes row, Y denotes col
	//
	//        col 0  | col 1 | col 2
	// --------------------------------
	// row 0  [x0,y0] [x0,y1] [x0,y2]
	// row 1  [x1,y0] [x1,y1] [x1,y2]
	if from.X < to.X {
		return SOUTH
	}
	if from.X > to.X {
		return NORTH
	}
	if from.Y < to.Y {
		return EAST
	}
	if from.Y > to.Y {
		return WEST
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

func (b *Board) BreakWall2(from, to *Cell, direction FlagPosition) {
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

func (b *Board) Break2Walls(c *Cell, idx int) {
	if idx == 0 {
		left := &b.Cells[c.X][c.Y-1]
		right := &b.Cells[c.X][c.Y+1]
		b.BreakWall2(left, c, EAST)
		b.BreakWall2(c, right, EAST)
		//fmt.Printf("\t%v - %v\n", left, c)
		//fmt.Printf("\t%v - %v\n", c, right)
	} else {
		up := &b.Cells[c.X-1][c.Y]
		down := &b.Cells[c.X+1][c.Y]
		b.BreakWall2(up, c, SOUTH)
		b.BreakWall2(c, down, SOUTH)
		//fmt.Printf("\t%v - %v\n", up, c)
		//fmt.Printf("\t%v - %v\n", c, down)
	}
}

func (b *Board) Write(writer io.Writer) {
	writer.Write([]byte("  "))
	for i := uint16(1); i < b.Width; i++ {
		writer.Write([]byte(" _"))
	}
	writer.Write([]byte("\n"))

	for h := uint16(0); h < b.Height; h++ {
		writer.Write([]byte("|"))
		for w := uint16(0); w < b.Width; w++ {
			c := b.Cells[h][w]
			if w == b.Width-1 && h == b.Height-1 {
				writer.Write([]byte(" |"))
				break
			}
			if c.IsSet(SOUTH) {
				writer.Write([]byte("_"))
			} else {
				writer.Write([]byte(" "))
			}

			if c.IsSet(EAST) {
				writer.Write([]byte("|"))
			} else {
				writer.Write([]byte(" "))
			}
		}
		writer.Write([]byte("\n"))
	}
}

func (b *Board) Write2(writer io.Writer) {
	for h := uint16(0); h < b.Height; h++ {
		for w := uint16(0); w < b.Width; w++ {
			c := b.Cells[h][w]
			if c.IsSet(DEAD) {
				writer.Write([]byte("*"))
			} else {
				writer.Write([]byte("-"))
			}
		}
		writer.Write([]byte("\n"))
	}
}

func (b *Board) WriteVisited(writer io.Writer) {
	for h := uint16(0); h < b.Height; h++ {
		for w := uint16(0); w < b.Width; w++ {
			c := b.Cells[h][w]
			if c.IsSet(VISITED) {
				writer.Write([]byte("*"))
			} else {
				writer.Write([]byte(" "))
			}
		}
		writer.Write([]byte("\n"))
	}
}

func (b *Board) DeadEnds(stack *Stack) {
	//this function is a memory optimzation, declaring h, w etc outside of
	//the for loops reduces allocations.
	flag := uint16(0)
	walls := uint16(0)
	h := uint16(0)
	c := &Cell{}

	for ; h < b.Height; h++ {
		for w := uint16(0); w < b.Width; w++ {
			c = &b.Cells[h][w]
			if c.IsSet(DEAD) {
				continue
			}

			// check for solid walls
			flag = c.Flag & 15
			walls = 0
			walls += flag & 1
			walls += (flag >> 1) & 1
			walls += (flag >> 2) & 1
			walls += (flag >> 3) & 1

			if walls >= 3 {
				stack.Push(c)
			}
		}
	}
}

func (b *Board) DeadNeighbors(c *Cell, stack *Stack) {
	result := []*Cell{}
	if ok, cell := b.getLiveNeighbor(c.X+1, c.Y); ok {
		result = append(result, cell)
	}

	if ok, cell := b.getLiveNeighbor(c.X-1, c.Y); ok {
		result = append(result, cell)
	}

	if ok, cell := b.getLiveNeighbor(c.X, c.Y+1); ok {
		result = append(result, cell)
	}

	if ok, cell := b.getLiveNeighbor(c.X, c.Y-1); ok {
		result = append(result, cell)
	}

	for _, item := range result {
		// check for solid walls
		flag := item.Flag & 15
		walls := uint16(0)
		walls += flag & 1
		walls += (flag >> 1) & 1
		walls += (flag >> 2) & 1
		walls += (flag >> 3) & 1

		// check for surrounding cells that are dead ends
		if !item.IsSet(WEST) && item.Y-1 >= 0 &&
			b.Cells[item.X][item.Y-1].IsSet(DEAD) {
			walls += 1
		}
		if !item.IsSet(EAST) && item.Y+1 < b.Width &&
			b.Cells[item.X][item.Y+1].IsSet(DEAD) {
			walls += 1
		}
		if !item.IsSet(NORTH) && item.X != 0 &&
			b.Cells[item.X-1][item.Y].IsSet(DEAD) {
			walls += 1
		}
		if !item.IsSet(SOUTH) && item.X+1 < b.Height &&
			b.Cells[item.X+1][item.Y].IsSet(DEAD) {
			walls += 1
		}

		if walls >= 3 {
			stack.Push(item)
		}
	}
}

func (b *Board) getLiveNeighbor(x, y uint16) (bool, *Cell) {
	if x >= 0 && x < b.Height &&
		y >= 0 && y < b.Width &&
		!b.Cells[x][y].IsSet(DEAD) {
		return true, &b.Cells[x][y]
	}
	return false, nil
}

func (b *Board) getNeighbor(x, y uint16) (bool, *Cell) {
	if x >= 0 && x < b.Height &&
		y >= 0 && y < b.Width &&
		b.Cells[x][y].IsSet(DEAD) {
		return false, nil
	}

	if x >= 0 && x < b.Height &&
		y >= 0 && y < b.Width &&
		!b.Cells[x][y].IsSet(VISITED) {
		return true, &b.Cells[x][y]
	}
	return false, nil
}

func (b *Board) GetCircleNeighbor(from, to *Cell, direction FlagPosition) (bool, []*Cell) {
	var cells []*Cell
	if direction == NORTH { //from south to north
		flatCellsTo := b.FlattenCells(to)
		for _, cell := range flatCellsTo {
			fmt.Printf("**[%d,%d]: %f, %f\n", cell.X, cell.Y, cell.ThetaFrom, cell.ThetaTo)
		}
		for _, cell := range flatCellsTo {
			if cell.ThetaFrom >= from.ThetaFrom &&
				cell.ThetaTo <= from.ThetaTo {
				cells = append(cells, cell)
				fmt.Printf("....%f, %f added to cells\n", cell.ThetaFrom, cell.ThetaTo)
			}
		}
	} else if direction == SOUTH {
		flatCellsFrom := b.FlattenCells(from)
		flatCellsTo := b.FlattenCells(to)
		if len(flatCellsFrom) == len(flatCellsTo) {
			cells = append(cells, flatCellsTo[0])
		} else {
			for _, cell := range flatCellsTo {
				if from.ThetaFrom <= cell.ThetaFrom &&
					from.ThetaTo <= cell.ThetaTo {
					cells = append(cells, cell)
				}
			}
		}
	}

	return false, cells
}

func (b *Board) GetCircleNeighborEW(cell *Cell, direction FlagPosition) *Cell {
	if direction == EAST {
		if cell.Parent == nil {
			//this cell has no splits, so by definition the next cell
			//has no splits as well.
			col := cell.Y + 1
			if cell.Y == b.Width-1 {
				col = 0
				fmt.Printf("cell.Y=%d, resetting to 0\n", cell.Y)
			}
			c := b.Cells[cell.X][col]
			fmt.Printf("EAST: %+v\n", c)
			return &c
		}

		if cell.Parent.Right != cell {
			return cell.Parent.Right
		}

		return cell.Parent.Parent.Right.Left
	} else if direction == WEST {
		if cell.Parent == nil {
			//this cell has no splits, so by definition the cell to the left
			//has no splits as well.
			col := cell.Y - 1
			if cell.Y == 0 {
				col = b.Width - 1
				fmt.Printf("cell.Y=%d, resetting to %d\n", cell.Y, b.Width-1)
			}
			c := b.Cells[cell.X][col]
			fmt.Printf("WEST: %+v\n", c)
			return &c
		}

		if cell.Parent.Left != cell {
			return cell.Parent.Left
		} else {
			c := b.Cells[cell.X][cell.Y-1]
			rightmost := b.getRightMostCell(&c)
			return rightmost
		}

		return cell.Parent.Parent.Left.Right
	}
	return nil
}

func (b *Board) getLeftMostCell(cell *Cell) *Cell {
	if cell.Left == nil {
		return cell
	}

	return b.getLeftMostCell(cell.Left)
}

func (b *Board) getRightMostCell(cell *Cell) *Cell {
	if cell.Right == nil {
		return cell
	}

	return b.getRightMostCell(cell.Right)
}

func (b *Board) FlattenCells(cell *Cell) []*Cell {
	fmt.Printf("FlattenCells called\n")
	var cells []*Cell
	if cell.Left == nil {
		cells = append(cells, cell)
		return cells
	}

	cells = append(cells, b.FlattenCells(cell.Left)...)
	cells = append(cells, b.FlattenCells(cell.Right)...)
	return cells
}
