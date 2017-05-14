package structs

import (
	"testing"
	"bytes"
	"strings"
)

func TestInit(t *testing.T) {
	b := &Board{10,10,nil}
	for _, row := range b.Cells {
		for _, cell := range row {
			if cell.Flag != 15 {
				t.Errorf("Init() should init all Flag to 15, got %d for %+v", cell.Flag, cell)
			}
		}
	}
}

func TestNeighbors(t *testing.T) {
	b := &Board{3,3,nil}
	b.Init()
	var tests = []struct{
		x uint16
		y uint16
		want int
	}{
		{0,0,2},
		{0,1,3},
		{0,2,2},
		{1,0,3},
		{1,1,4},
		{1,2,3},
		{2,0,2},
		{2,1,3},
		{2,2,2},
	}
	for _, test := range tests {
		c := &b.Cells[test.x][test.y]
		got := b.Neighbors(c)
		if len(got) != test.want {
			t.Errorf("Neighbors(%+v) got %d, want %d", *c, len(got), test.want)
		}
	}
}

func TestGetDirection(t *testing.T) {
	b := &Board{10,10,nil}
	b.Init()
	var tests = []struct{
		from *Cell
		to *Cell
		want FlagPosition
	}{
		{&b.Cells[1][1], &b.Cells[0][1], WEST},
		{&b.Cells[1][1], &b.Cells[1][0], NORTH},
		{&b.Cells[1][1], &b.Cells[1][2], SOUTH},
		{&b.Cells[1][1], &b.Cells[2][1], EAST},
	}

	for _, test := range tests {
		dir := b.GetDirection(test.from, test.to)
		if dir != test.want {
			t.Errorf("GetDirection(%v, %v), should be %v, got %v", test.from, test.to, dir, test.want)
		}
	}
}

func TestBreakWall(t *testing.T) {
	b := &Board{10,10,nil}
	b.Init()
	var tests = []struct{
		from, to *Cell
		dir FlagPosition
		want1, want2 uint8
	}{
		{&b.Cells[1][1], &b.Cells[0][1], WEST, 23, 27},
		{&b.Cells[1][1], &b.Cells[1][0], NORTH, 30, 29},
		{&b.Cells[1][1], &b.Cells[1][2], SOUTH, 29, 30},
		{&b.Cells[1][1], &b.Cells[2][1], EAST, 27, 23},
	}

	for _, test := range tests {
		b.BreakWall(test.from, test.to, test.dir)
		if test.from.Flag != test.want1 && test.to.Flag != test.want2 {
			t.Error("BreakWall(%v, %v), flags should be %d, %d, got %d, %d", 
				test.want1, test.want2, test.from.Flag, test.to.Flag)
		}
	}
}

func TestWrite(t *testing.T) {
	b := &Board{3,3,nil}
	b.Init()
	var buf bytes.Buffer
	b.Write(&buf)
	got := strings.TrimRight(buf.String(), string(10)) // remove trailing line feed
	want := "   _ _\n|_|_|_|\n|_|_|_|\n|_|_| |"
	if got != want {
		t.Errorf("Write(), \nwant \n%s \ngot \n%s", want, got)
	}
}