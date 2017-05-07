package algos

import (
	"github.com/wliao008/mazing/structs"
	"github.com/wliao008/mazing/util"
)

type BackTracking struct {
	Board structs.Board
}

var directions []interface{}

func NewBackTracking(width, height uint16) *BackTracking {
	bt := &BackTracking{Board: structs.Board{width, height, nil}}
	bt.Board.Cells = make([][]structs.Cell, width)
	for i := uint16(0); i < width; i++ {
		bt.Board.Cells[i] = make([]structs.Cell, height)
	}

	for i := uint16(0); i < width; i++ {
		for j := uint16(0); j < height; j++ {
			bt.Board.Cells[i][j].Flag = 15
			bt.Board.Cells[i][j].X = i
			bt.Board.Cells[i][j].Y = j
		}
	}
	return bt
}

func init() {
	directions = append(directions, structs.NORTH, structs.SOUTH, structs.EAST, structs.WEST)
}

func (bt *BackTracking) Generate() error {
	//start at cell 0,0
	bt.doWork(0, 0)
	return nil
}

//doWork: the recrusive backtracking algorithm
func (bt *BackTracking) doWork(x, y int) {
	d := structs.Direction{}
	util.Shuffle(directions)
	for _, direction := range directions {
		dir := direction.(structs.FlagPosition)
		var nextX int = x + d.XDirection(dir)
		var nextY int = y + d.YDirection(dir)
		if nextX >= 0 && nextX < int(bt.Board.Width) &&
			nextY >= 0 && nextY < int(bt.Board.Height) &&
			!bt.Board.Cells[nextX][nextY].IsSet(structs.VISITED) {
			bt.Board.BreakWall(&bt.Board.Cells[x][y], &bt.Board.Cells[nextX][nextY], dir)
			bt.doWork(nextX, nextY)
		}
	}
}
