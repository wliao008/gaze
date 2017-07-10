package models

import "github.com/wliao008/mazing/structs"
type BoardModel struct {
	Name string
	Cells [][]CellModel
	RawCells [][]structs.Cell
	Height, Width uint16
}

type CellModel struct {
	X, Y     uint16
	CssClasses string
}
