package gaze

import (
	"fmt"
	"math"
)

type BoardCircle struct {
	Height uint16 // row
	Width  uint16 // col
	Cells  [][]Cell
}

func (b *BoardCircle) Init() {
	b.Cells = make([][]Cell, b.Height)
	for i := uint16(0); i < b.Height; i++ {
		b.Cells[i] = make([]Cell, b.Width)
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

	b.SplitCells(12)
}

func (b *BoardCircle) SplitCells(offset int) {
	centerx := int(b.Width)*offset*2/2 + offset
	centery := int(b.Height)*offset*2/2 + offset
	theta := math.Pi * 2 / float64(b.Width)
	rowCount := 1
	for row := int(b.Height); row > 0; row-- {
		for i := int(b.Width); i > 0; i-- {
			c := &b.Cells[row-1][i-1]
			b.DoSplit(c, offset, rowCount, i, centerx, centery, theta, int(b.Width))
		}
		rowCount = rowCount + 1
	}
}

func (b *BoardCircle) DoSplit(cell *Cell, offset, rowCount, i, centerx, centery int, theta float64, width int) {
	radius := float64(offset * rowCount)
	thetaccw := float64(theta) * float64(i)
	thetacw := float64(theta) * float64(i+1)
	ax := float64(centerx) + (radius * math.Cos(thetaccw))
	ay := float64(centery) + (radius * math.Sin(thetaccw))

	cx := float64(centerx) + ((radius) * math.Cos(thetacw))
	cy := float64(centery) + ((radius) * math.Sin(thetacw))

	south := math.Sqrt(math.Pow(cx-ax, 2) + math.Pow(cy-ay, 2))
	cell.ThetaFrom = thetaccw
	cell.ThetaTo = thetacw
	//fmt.Printf("[%d,%d] %f, %f\n", cell.X, cell.Y, cell.ThetaFrom, cell.ThetaTo)
	b.split(cell, offset, south)
}

func (b *BoardCircle) split(cell *Cell, offset int, south float64) {
	half := south / float64(offset)
	if half >= 1.6 {
		//split
		cell.Left = &Cell{Flag: 15, X: cell.X, Y: cell.Y}
		cell.Left.Parent = cell
		cell.Left.ThetaTo = cell.ThetaTo
		cell.Left.ThetaFrom = cell.ThetaFrom + math.Abs(cell.ThetaFrom-cell.ThetaTo)/2
		//fmt.Printf("\t[%d,%d] left\t %f, %f\n", cell.Left.X, cell.Left.Y, cell.Left.ThetaFrom, cell.Left.ThetaTo)

		cell.Right = &Cell{Flag: 15, X: cell.X, Y: cell.Y}
		cell.Right.Parent = cell
		cell.Right.ThetaTo = cell.Left.ThetaFrom
		cell.Right.ThetaFrom = cell.ThetaFrom
		//fmt.Printf("\t[%d,%d] right\t %f, %f\n", cell.Right.X, cell.Right.Y, cell.Right.ThetaFrom, cell.Right.ThetaTo)
		//fmt.Printf("ax=%f ay=%f, cx=%f cy=%f, cx-ax=%f, south=%f\n", ax, ay, cx, cy, cx-ax, south)
		b.split(cell.Left, offset, south/2)
		b.split(cell.Right, offset, south/2)
		//fmt.Printf("half, southwall=%f\n", south/2)
	}
}

func (b *BoardCircle) print() {
	for h := 0; h < int(b.Height); h++ {
		for w := 0; w < int(b.Width); w++ {
			c := &b.Cells[h][w]
			b.doprint(c, 0)
		}
	}
}

func (b *BoardCircle) doprint(cell *Cell, spacers int) {
	str := ""
	for i := 0; i < spacers; i++ {
		str += " "
	}

	if cell.Left == nil {
		fmt.Printf("%s[%d,%d] %f, %f\n", str, cell.X, cell.Y, cell.ThetaFrom, cell.ThetaTo)
		return
	}

	spacers += 4
	fmt.Printf("%s[%d,%d] %f, %f\n", str, cell.X, cell.Y, cell.ThetaFrom, cell.ThetaTo)
	b.doprint(cell.Left, spacers)
	b.doprint(cell.Right, spacers)
}

// Neighbors find the neighboring cells of the current cell
func (b *BoardCircle) Neighbors(c *Cell) []*Cell {
	result := []*Cell{}
	//north
	x := c.X - 1
	if x >= 0 && x < b.Height &&
		c.Y >= 0 && c.Y < b.Width {
		to := b.Cells[x][c.Y]
		if ok, cells := b.GetCircleNeighbor(c, &to, NORTH); ok {
			for _, cell := range cells {
				if !cell.IsSet(VISITED) {
					result = append(result, cell)
				}
			}
		}
	}

	//south
	x = c.X + 1
	if x >= 0 && x < b.Height &&
		c.Y >= 0 && c.Y < b.Width {
		to := b.Cells[x][c.Y]
		if ok, cells := b.GetCircleNeighbor(c, &to, SOUTH); ok {
			for _, cell := range cells {
				if !cell.IsSet(VISITED) {
					result = append(result, cell)
				}
			}
		}
	}

	//east
	if ok, cell := b.GetCircleNeighborEW(c, EAST); ok {
		if !cell.IsSet(VISITED) {
			result = append(result, cell)
		}
	}

	//west
	if ok, cell := b.GetCircleNeighborEW(c, WEST); ok {
		if !cell.IsSet(VISITED) {
			result = append(result, cell)
		}
	}
	return result
}

func (b *BoardCircle) GetDirection(from, to *Cell) FlagPosition {
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
	if from.Y == b.Width-1 && to.Y == 0 {
		return EASTWESTWRAP
	}
	if from.Y == 0 && to.Y == b.Width-1 {
		return WESTEASTWRAP
	}
	if from.Y < to.Y {
		return EAST
	}
	if from.Y > to.Y {
		return WEST
	}
	if from.Y == to.Y {
		//split cells
		parent := b.getParentCell(from)
		cells := b.FlattenCells(parent)
		fromIdx := b.getIndex(from, cells)
		toIdx := b.getIndex(to, cells)
		if fromIdx > toIdx {
			return WEST
		}
		return EAST
	}
	//TODO: This is really an error case here
	return EAST
}

func (b *BoardCircle) BreakWall(from, to *Cell, direction FlagPosition) {
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
	case EASTWESTWRAP:
		from.ClearBit(EAST)
		to.ClearBit(WEST)
	case WESTEASTWRAP:
		from.ClearBit(WEST)
		to.ClearBit(EAST)
	}
}

/*
func (b *BoardCircle) BreakWall2(from, to *Cell, direction FlagPosition) {
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

func (b *BoardCircle) Break2Walls(c *Cell, idx int) {
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

func (b *BoardCircle) Write(writer io.Writer) {
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

func (b *BoardCircle) Write2(writer io.Writer) {
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

func (b *BoardCircle) WriteVisited(writer io.Writer) {
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
}*/

func (b *BoardCircle) DeadEnds(stack *Stack) {
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

func (b *BoardCircle) DeadNeighbors(c *Cell, stack *Stack) {
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

func (b *BoardCircle) getLiveNeighbor(x, y uint16) (bool, *Cell) {
	if x >= 0 && x < b.Height &&
		y >= 0 && y < b.Width &&
		!b.Cells[x][y].IsSet(DEAD) {
		return true, &b.Cells[x][y]
	}
	return false, nil
}

func (b *BoardCircle) getNeighbor(x, y uint16) (bool, *Cell) {
	if x >= 0 && x < b.Height &&
		y >= 0 && y < b.Width &&
		!b.Cells[x][y].IsSet(VISITED) {
		return true, &b.Cells[x][y]
	}
	return false, nil
}

func (b *BoardCircle) GetCircleNeighbor(from, to *Cell, direction FlagPosition) (bool, []*Cell) {
	var cells []*Cell
	if direction == NORTH { //from south to north
		parentTo := b.getParentCell(to)
		flatCellsTo := b.FlattenCells(parentTo)

		for _, cell := range flatCellsTo {
			if cell.ThetaFrom >= from.ThetaFrom &&
				cell.ThetaTo <= from.ThetaTo &&
				!cell.IsSet(VISITED) {
				cells = append(cells, cell)
			}
		}
	} else if direction == SOUTH {
		parentFrom := b.getParentCell(from)
		flatCellsFrom := b.FlattenCells(parentFrom)
		parentTo := b.getParentCell(to)
		flatCellsTo := b.FlattenCells(parentTo)

		if len(flatCellsFrom) == len(flatCellsTo) {
			idx := b.getIndex(from, flatCellsFrom)
			cells = append(cells, flatCellsTo[idx])
		} else {
			for _, cell := range flatCellsTo {
				if from.ThetaFrom >= cell.ThetaFrom &&
					from.ThetaTo <= cell.ThetaTo &&
					!cell.IsSet(VISITED) {
					cells = append(cells, cell)
				}
			}
		}
	}

	return true, cells
}

func (b *BoardCircle) GetCircleNeighborEW(cell *Cell, direction FlagPosition) (bool, *Cell) {
	parent := b.getParentCell(cell)
	cells := b.FlattenCells(parent)
	if len(cells) == 0 {
		return false, nil
	}

	if direction == EAST {
		//no split cells
		if len(cells) == 1 {
			col := cell.Y + 1
			if cell.Y == b.Width-1 {
				col = 0
			}
			c := b.Cells[cell.X][col]
			if !c.IsSet(VISITED) {
				return true, &c
			}
			return false, nil
		}

		idx := b.getIndex(cell, cells)
		if idx == len(cells)-1 {
			//should bleed into the east (next cell)
			col := cell.Y + 1
			if cell.Y == b.Width-1 {
				col = 0
			}
			c := b.Cells[cell.X][col]
			nextCells := b.FlattenCells(&c)
			if !nextCells[0].IsSet(VISITED) {
				return true, nextCells[0]
			}
			return false, nil
		}

		if !cells[idx+1].IsSet(VISITED) {
			return true, cells[idx+1]
		}
		return false, nil
	} else if direction == WEST {
		//no split cells
		if len(cells) == 1 {
			col := cell.Y - 1
			if cell.Y == 0 {
				col = b.Width - 1
			}
			c := b.Cells[cell.X][col]
			if !c.IsSet(VISITED) {
				return true, &c
			}
			return false, nil
		}

		idx := b.getIndex(cell, cells)
		if idx == 0 {
			//should bleed into the west (prev cell)
			col := cell.Y - 1
			if cell.Y == 0 {
				col = b.Width - 1
			}
			c := b.Cells[cell.X][col]
			nextCells := b.FlattenCells(&c)
			if !nextCells[len(nextCells)-1].IsSet(VISITED) {
				return true, nextCells[len(nextCells)-1]
			}
			return false, nil
		}

		if !cells[idx-1].IsSet(VISITED) {
			return true, cells[idx-1]
		}
		return false, nil
	}
	return false, nil
}

func (b *BoardCircle) getIndex(cell *Cell, cells []*Cell) int {
	for i, c := range cells {
		if c == cell {
			return i
		}
	}

	return -1
}

func (b *BoardCircle) GetLeftMostCell(cell *Cell) *Cell {
	if cell.Left == nil {
		return cell
	}

	return b.GetLeftMostCell(cell.Left)
}

func (b *BoardCircle) getRightMostCell(cell *Cell) *Cell {
	if cell.Right == nil {
		return cell
	}

	return b.getRightMostCell(cell.Right)
}

func (b *BoardCircle) getParentCell(cell *Cell) *Cell {
	if cell.Parent == nil {
		return cell
	}

	return b.getParentCell(cell.Parent)
}

func (b *BoardCircle) FlattenCells(cell *Cell) []*Cell {
	var cells []*Cell
	if cell.Left == nil {
		cells = append(cells, cell)
		return cells
	}

	cells = append(cells, b.FlattenCells(cell.Left)...)
	cells = append(cells, b.FlattenCells(cell.Right)...)
	return cells
}
