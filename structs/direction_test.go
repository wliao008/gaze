package structs

import "testing"

func TestXDirection(t *testing.T) {
	dir := &Direction{}
	tests := []struct{
		pos FlagPosition
		want int
	}{
		{NORTH, -1},
		{SOUTH, 1},
		{EAST, 0},
		{WEST, 0},
	}

	for _, test := range tests{
		got := dir.XDirection(test.pos)
		if got != test.want {
			t.Errorf("XDirection(%d) want %d, got %d", test.pos, test.want, got)
		}
	}
}

func TestYDirection(t *testing.T) {
	dir := &Direction{}
	tests := []struct{
		pos FlagPosition
		want int
	}{
		{NORTH, 0},
		{SOUTH, 0},
		{EAST, 1},
		{WEST, -1},
	}

	for _, test := range tests{
		got := dir.YDirection(test.pos)
		if got != test.want {
			t.Errorf("XDirection(%d) want %d, got %d", test.pos, test.want, got)
		}
	}
}
