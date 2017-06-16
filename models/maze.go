package models

import "github.com/wliao008/mazing/structs"
type BoardModel struct {
	Cells [][]CellModel
	RawCells [][]structs.Cell
}

type CellModel struct {
	X, Y     uint16
	CssClasses string
	Note string
	Flag string
}
