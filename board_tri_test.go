package gaze

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestBoardTriInit(t *testing.T) {
	b := &BoardTri{10, 10, nil}
	for _, row := range b.Cells {
		for _, cell := range row {
			if cell.Flag != 13 || cell.Flag != 526 {
				t.Errorf("Init() should init all Flag to either 13 or 526, got %d for %+v", cell.Flag, cell)
			}
		}
	}
}

func TestBoardTriNeighbors(t *testing.T) {
	b := &BoardTri{3, 3, nil}
	b.Init()
	var tests = []struct {
		x    uint16
		y    uint16
		want int
	}{
		{0, 0, 2},
		{0, 1, 2},
		{0, 2, 2},
		{1, 0, 2},
		{1, 1, 3},
		{1, 2, 2},
		{2, 0, 1},
		{2, 1, 3},
		{2, 2, 1},
	}
	for _, test := range tests {
		c := &b.Cells[test.x][test.y]
		got := b.Neighbors(c)
		if len(got) != test.want {
			t.Errorf("Neighbors(%+v) got %d, want %d", *c, len(got), test.want)
		}
	}
}

func TestBoardTriGetDirection(t *testing.T) {
	b := &BoardTri{10, 10, nil}
	b.Init()
	var tests = []struct {
		from *Cell
		to   *Cell
		want FlagPosition
	}{
		{&b.Cells[1][1], &b.Cells[1][0], WEST},
		{&b.Cells[1][1], &b.Cells[1][2], EAST},
		{&b.Cells[1][1], &b.Cells[2][1], SOUTH},
	}

	for _, test := range tests {
		dir := b.GetDirection(test.from, test.to)
		if dir != test.want {
			t.Errorf("GetDirection(%v, %v), should be %v, got %v", test.from, test.to, dir, test.want)
		}
	}
}

func TestBoardTriBreakWall(t *testing.T) {
	b := &BoardTri{10, 10, nil}
	b.Init()
	var tests = []struct {
		from, to     *Cell
		dir          FlagPosition
		want1, want2 uint16
	}{
		{&b.Cells[1][2], &b.Cells[0][2], NORTH, 30, 541},
		{&b.Cells[1][2], &b.Cells[1][1], WEST, 22, 539},
		{&b.Cells[1][2], &b.Cells[1][3], EAST, 18, 535},
	}

	for _, test := range tests {
		b.BreakWall(test.from, test.to, test.dir)
		fmt.Printf("from.Flag=%d, to.Flag=%d", test.from.Flag, test.to.Flag)
		if test.from.Flag != test.want1 && test.to.Flag != test.want2 {
			t.Errorf("BreakWall(,), flags should be %d, %d, got %d, %d",
				test.want1, test.want2, test.from.Flag, test.to.Flag)
		}
	}
}

func TestBoardTriWrite(t *testing.T) {
	b := &BoardTri{3, 3, nil}
	b.Init()
	var buf bytes.Buffer
	b.Write(&buf)
	got := strings.TrimRight(buf.String(), string(10)) // remove trailing line feed
	want := " ________\n /\\\\--//\\\n/__\\\\//__\\\n\\--//\\\\--/\n \\//__\\\\/\n /\\\\--//\\\n/__\\\\//__\\"
	if got != want {
		t.Errorf("Write(), \nwant \n%s \ngot \n%s", want, got)
	}
}

func TestBoardTriDeadEnds(t *testing.T) {
	b := &BoardTri{3, 3, nil}
	b.Init()
	stack := &Stack{}
	b.DeadEnds(stack)
	count := 0
	c := stack.Pop()
	for c != nil {
		count += 1
		c = stack.Pop()
	}
	if count != 9 {
		t.Error("DeadEnds(), want 9, got %d", count)
	}
}

func TestBoardTriDeadEnds2(t *testing.T) {
	b := &BoardTri{1, 10, nil}
	b.Init()
	b.Cells[0][0].ClearBit(NORTH)
	b.Cells[0][9].ClearBit(SOUTH)
	for i := 0; i < 9; i++ {
		b.Cells[0][i].ClearBit(EAST)
		b.Cells[0][i+1].ClearBit(WEST)
	}
	stack := &Stack{}
	b.DeadEnds(stack)
	count := 0
	c := stack.Pop()
	for c != nil {
		count += 1
		c = stack.Pop()
	}
	if count != 0 {
		t.Error("DeadEnds(), want 0, got %d", count)
	}
}
