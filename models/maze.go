package models

type BoardModel struct {
	Cells [][]CellModel
}

type CellModel struct {
	X, Y     uint16
	CssClasses string
}
