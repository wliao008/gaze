package structs

type Board struct {
	Width, Height uint16
	Cells [][]Cell
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

	switch direction{
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

func (b *Board) getNeighbor(x, y uint16) (bool, *Cell) {
	if x >= 0 && x < b.Width &&
		y >= 0 && y < b.Height &&
		!b.Cells[x][y].IsSet(VISITED) {
			return true, &b.Cells[x][y]
		}
	return false, nil
}
