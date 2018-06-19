package gaze

type BoardModel struct {
	Name          string
	Cells         [][]CellModel
	RawCells      [][]Cell
	Height, Width uint16
	TableCss      string
	WeaveChecked  string
}

type CellModel struct {
	X, Y       uint16
	CssClasses string
}
