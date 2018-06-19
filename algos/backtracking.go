package algos

import (
	"github.com/wliao008/gaze/structs"
	"github.com/wliao008/gaze/util"
)

type BackTracking struct {
	Board structs.Board
}

var directions []interface{}

func NewBackTracking(height, width uint16) *BackTracking {
	bt := &BackTracking{Board: structs.Board{height, width, nil}}
	bt.Board.Init()
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
		if nextX >= 0 && nextX < int(bt.Board.Height) &&
			nextY >= 0 && nextY < int(bt.Board.Width) &&
			!bt.Board.Cells[nextX][nextY].IsSet(structs.VISITED) {
			bt.Board.BreakWall(&bt.Board.Cells[x][y], &bt.Board.Cells[nextX][nextY], dir)
			bt.doWork(nextX, nextY)
		}
	}
}
