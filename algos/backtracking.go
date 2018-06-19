package algos

import (
	"github.com/wliao008/gaze"
)

type BackTracking struct {
	Board gaze.Board
}

var directions []interface{}

func NewBackTracking(height, width uint16) *BackTracking {
	bt := &BackTracking{Board: gaze.Board{height, width, nil}}
	bt.Board.Init()
	return bt
}

func init() {
	directions = append(directions, gaze.NORTH, gaze.SOUTH, gaze.EAST, gaze.WEST)
}

func (bt *BackTracking) Generate() error {
	//start at cell 0,0
	bt.doWork(0, 0)
	return nil
}

//doWork: the recrusive backtracking algorithm
func (bt *BackTracking) doWork(x, y int) {
	d := gaze.Direction{}
	gaze.Shuffle(directions)
	for _, direction := range directions {
		dir := direction.(gaze.FlagPosition)
		var nextX int = x + d.XDirection(dir)
		var nextY int = y + d.YDirection(dir)
		if nextX >= 0 && nextX < int(bt.Board.Height) &&
			nextY >= 0 && nextY < int(bt.Board.Width) &&
			!bt.Board.Cells[nextX][nextY].IsSet(gaze.VISITED) {
			bt.Board.BreakWall(&bt.Board.Cells[x][y], &bt.Board.Cells[nextX][nextY], dir)
			bt.doWork(nextX, nextY)
		}
	}
}
