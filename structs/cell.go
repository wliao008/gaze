package structs

import "fmt"

type FlagPosition uint16

const (
	NORTH FlagPosition = iota
	SOUTH
	EAST
	WEST
	VISITED
	START
	END
	DEAD 	//a dead cell is not part of the solution path
	CROSS	//for weave maze
)

type Cell struct {
	/* bits that has the following structure:
	 *
	 * 7|6|5|4|3|2|1|0
	 *
	 * 0: north door
	 * 1: south
	 * 2: east
	 * 3: west
	 * 4: visited
	 * 5: start
	 * 6: end
	 * 7: dead
	 */
	Flag uint16
	X uint16 // row
	Y uint16 // col
}

func (c *Cell) SetBit(pos FlagPosition) {
	c.Flag |= 1 << pos
}

func (c *Cell) ClearBit(pos FlagPosition) {
	c.Flag &= ^(1 << pos)
}

func (c *Cell) IsSet(pos FlagPosition) bool {
	return c.Flag&(1<<pos) != 0
}

func (c *Cell) String() string {
	return fmt.Sprintf("[%d,%d]", c.X, c.Y)
}
