package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/wliao008/mazing/algos"
	"github.com/wliao008/mazing/models"
	_ "github.com/wliao008/mazing/structs"
	_ "os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("web/templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request){
	bt := algos.NewKruskal(5, 4)
	err := bt.Generate()
	if err != nil {
		fmt.Println("ERROR")
	}

	// create model
	model := &models.BoardModel{}
	model.Cells = make([][]models.CellModel, bt.Board.Height)
	model.RawCells = bt.Board.Cells
	for i := uint16(0); i < bt.Board.Height; i++ {
		model.Cells[i] = make([]models.CellModel, bt.Board.Width)
	}
	// initialize model
	for w := uint16(0); w < bt.Board.Width; w++ {
		model.Cells[0][w].CssClasses += "north "
		model.Cells[bt.Board.Height-1][w].CssClasses += "south "
	}

	for h := uint16(0); h < bt.Board.Height; h++ {
		model.Cells[h][0].CssClasses +="west "
		for w := uint16(0); w < bt.Board.Width; w++ {
			model.Cells[h][w].X = w;
			model.Cells[h][w].Y = h
			if w == bt.Board.Width - 1 {
				model.Cells[h][w].CssClasses +="east "
			}
			//cell := bt.Board.Cells[j][i]
			/*
			model.Cells[i][j].X = i;
			model.Cells[i][j].Y = j
			
			if i==0 {
				model.Cells[0][j].CssClasses +="north "
				model.Cells[0][j].Note += "north "
			}

			model.Cells[i][j].Flag = fmt.Sprint(bt.Board.Cells[i][j].Flag)

			if bt.Board.Cells[i][j].IsSet(structs.EAST) {
				model.Cells[i][j].CssClasses += "east "
				model.Cells[i][j].Note += "east "
			}
			if bt.Board.Cells[i][j].IsSet(structs.WEST) {
				model.Cells[i][j].CssClasses += "west "
				model.Cells[i][j].Note += "west "
			}
			if bt.Board.Cells[i][j].IsSet(structs.NORTH) {
				model.Cells[i][j].CssClasses += "north "
				model.Cells[i][j].Note += "north "
			}
			if bt.Board.Cells[i][j].IsSet(structs.SOUTH) {
				model.Cells[i][j].CssClasses += "south "
				model.Cells[i][j].Note += "south "
			}*/
		}
	}

	err = tpl.ExecuteTemplate(w, "index.html", model)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
