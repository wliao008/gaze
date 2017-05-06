package structs

type FlagPosition uint8

const (
	NORTH FlagPosition = iota
	SOUTH
	EAST
	WEST
	VISITED
	START
	END
	DEAD
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
	Flag uint8
	X uint16
	Y uint16
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
