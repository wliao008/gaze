package gaze

import (
	"fmt"
	"math/rand"
	"time"
	//"fmt"
	"io"
)

type BoardTri struct {
	Height uint16 // row
	Width  uint16 // col
	Cells  [][]Cell
}

func (b *BoardTri) Init() {
	rand.Seed(time.Now().UTC().UnixNano())
	b.Cells = make([][]Cell, b.Height)
	for i := uint16(0); i < b.Height; i++ {
		b.Cells[i] = make([]Cell, b.Width)
	}

	var idx int = rand.Intn(2)
	fmt.Println(idx)
	for h := uint16(0); h < b.Height; h++ {
		start := h

		for w := uint16(0); w < b.Width; w++ {
			if start%2 == uint16(idx) {
				b.Cells[h][w].Flag = 526 // triangle pointed up: 1000001110
			} else {
				b.Cells[h][w].Flag = 13 // triangle pointed down: 0000001101
			}
			start++

			// set the relative [x,y] position of the cell on the board
			b.Cells[h][w].X = h
			b.Cells[h][w].Y = w
		}
	}
}

// Neighbors find the (3 at most) neighboring cells of the current cell
func (b *BoardTri) Neighbors(c *Cell) []*Cell {
	result := []*Cell{}

	if c.IsSet(TRIANGLE_UP) {
		//ignore north since both points are at each other
		if ok, cell := b.getNeighbor(c.X+1, c.Y); ok {
			result = append(result, cell)
		}
	} else {
		if ok, cell := b.getNeighbor(c.X-1, c.Y); ok {
			result = append(result, cell)
		}
	}

	if ok, cell := b.getNeighbor(c.X, c.Y+1); ok {
		result = append(result, cell)
	}

	if ok, cell := b.getNeighbor(c.X, c.Y-1); ok {
		result = append(result, cell)
	}

	return result
}

func (b *BoardTri) GetDirection(from, to *Cell) FlagPosition {
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

func (b *BoardTri) BreakWall(from, to *Cell, direction FlagPosition) {
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

func (b *BoardTri) BreakWall2(from, to *Cell, direction FlagPosition) {
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

func (b *BoardTri) Break2Walls(c *Cell, idx int) {
	if idx == 0 {
		left := &b.Cells[c.X][c.Y-1]
		right := &b.Cells[c.X][c.Y+1]
		b.BreakWall2(left, c, EAST)
		b.BreakWall2(c, right, EAST)
	} else {
		up := &b.Cells[c.X-1][c.Y]
		down := &b.Cells[c.X+1][c.Y]
		b.BreakWall2(up, c, SOUTH)
		b.BreakWall2(c, down, SOUTH)
	}
}

func (b *BoardTri) Write3a(writer io.Writer) {
	for h := uint16(0); h < b.Height; h++ {
		for w := uint16(0); w < b.Width; w++ {
			c := b.Cells[h][w]
			if c.IsSet(TRIANGLE_UP) {
				writer.Write([]byte("∆"))
			} else {
				writer.Write([]byte("∇"))
			}
		}
		writer.Write([]byte("\n"))
	}
}

func (b *BoardTri) Write(writer io.Writer) {
	for h := uint16(0); h < b.Height; h++ {
		//left border
		/*
			c0 := b.Cells[h][0]
			if c0.IsSet(TRIANGLE_UP) {
				writer.Write([]byte("∕"))
			} else {
				writer.Write([]byte("∖"))
			}
		*/

		//first pass
		for w := uint16(0); w < b.Width; w++ {
			c := b.Cells[h][w]

			if c.IsSet(TRIANGLE_UP) {
				if w == 0 {
					if c.IsSet(WEST) {
						writer.Write([]byte(" /"))
					} else {
						writer.Write([]byte("  "))
					}
				} else {
					if c.IsSet(WEST) {
						writer.Write([]byte("/"))
					} else {
						writer.Write([]byte(" "))
					}
				}

				if c.IsSet(EAST) {
					writer.Write([]byte("\\"))
				} else {
					writer.Write([]byte(" "))
				}
			} else {
				//triangle pointing down
				if c.IsSet(WEST) {
					writer.Write([]byte("\\"))
				} else {
					writer.Write([]byte(" "))
				}

				if c.IsSet(NORTH) {
					writer.Write([]byte("--"))
				} else {
					writer.Write([]byte("  "))
				}

				if c.IsSet(EAST) {
					writer.Write([]byte("/"))
				} else {
					writer.Write([]byte(" "))
				}
			}
		}
		writer.Write([]byte("\n"))

		//2nd pass
		for w := uint16(0); w < b.Width; w++ {
			c := b.Cells[h][w]

			if c.IsSet(TRIANGLE_UP) {
				if c.IsSet(WEST) {
					writer.Write([]byte("/"))
				} else {
					writer.Write([]byte(" "))
				}

				if c.IsSet(SOUTH) {
					writer.Write([]byte("__"))
				} else {
					writer.Write([]byte("  "))
				}

				if c.IsSet(EAST) {
					writer.Write([]byte("\\"))
				} else {
					writer.Write([]byte(" "))
				}
			} else {
				//triangle pointing down
				if w == 0 {
					if c.IsSet(WEST) {
						writer.Write([]byte(" \\"))
					} else {
						writer.Write([]byte("  "))
					}
				} else {
					if c.IsSet(WEST) {
						writer.Write([]byte("\\"))
					} else {
						writer.Write([]byte(" "))
					}
				}

				if c.IsSet(EAST) {
					writer.Write([]byte("/"))
				} else {
					writer.Write([]byte(" "))
				}
			}
		}
		writer.Write([]byte("\n"))
	}
}

func (b *BoardTri) WriteVisited(writer io.Writer) {
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

func (b *BoardTri) DeadEnds(stack *Stack) {
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

func (b *BoardTri) DeadNeighbors(c *Cell, stack *Stack) {
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

func (b *BoardTri) getLiveNeighbor(x, y uint16) (bool, *Cell) {
	if x >= 0 && x < b.Height &&
		y >= 0 && y < b.Width &&
		!b.Cells[x][y].IsSet(DEAD) {
		return true, &b.Cells[x][y]
	}
	return false, nil
}

func (b *BoardTri) getNeighbor(x, y uint16) (bool, *Cell) {
	if x >= 0 && x < b.Height &&
		y >= 0 && y < b.Width &&
		!b.Cells[x][y].IsSet(VISITED) {
		return true, &b.Cells[x][y]
	}
	return false, nil
}

func (b *BoardTri) WriteSvg(writer io.Writer) {
	startx := 10
	starty := 10
	for h := uint16(0); h < b.Height; h++ {
		if h%2 == 0 {
			startx = 10
		} else {
			startx = 0
		}
		for w := uint16(0); w < b.Width; w++ {
			c := b.Cells[h][w]
			if c.IsSet(TRIANGLE_UP) {
				if c.IsSet(WEST) {
					fmt.Printf("<path d=\"M%d %d l -10 10\" fill=\"none\" stroke=\"blue\" stroke-width=\"1\"/>\n", startx, starty)
				}
				if c.IsSet(EAST) {
					fmt.Printf("<path d=\"M%d %d l 10 10\" fill=\"none\" stroke=\"blue\" stroke-width=\"1\"/>\n", startx, starty)
				}
				if c.IsSet(SOUTH) {
					fmt.Printf("<path d=\"M%d %d h 20\" fill=\"none\" stroke=\"blue\" stroke-width=\"1\"/>\n", startx-10, starty+10)
				}
			} else {
				if c.IsSet(NORTH) {
					fmt.Printf("<path d=\"M%d %d h 20\" fill=\"none\" stroke=\"blue\" stroke-width=\"1\"/>\n", startx, starty)
				}
				if c.IsSet(WEST) {
					fmt.Printf("<path d=\"M%d %d l 10 10\" fill=\"none\" stroke=\"blue\" stroke-width=\"1\"/>\n", startx, starty)
				}
				if c.IsSet(EAST) {
					fmt.Printf("<path d=\"M%d %d l -10 10\" fill=\"none\" stroke=\"blue\" stroke-width=\"1\"/>\n", startx+20, starty)
				}
				startx += 20
			}
		}
		starty += 10
	}
}
