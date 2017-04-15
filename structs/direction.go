package structs

type Direction struct{}

func (d *Direction) XDirection(dir FlagPosition) int {
	switch dir {
	case NORTH:
		return 0
	case SOUTH:
		return 0
	case EAST:
		return 1
	case WEST:
		return -1
	default:
		return 0
	}
}

func (d *Direction) YDirection(dir FlagPosition) int {
	switch dir {
	case NORTH:
		return -1
	case SOUTH:
		return 1
	case EAST:
		return 0
	case WEST:
		return 0
	default:
		return 0
	}
}
