package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/wliao008/mazing/algos"
	"github.com/wliao008/mazing/models"
	"github.com/wliao008/mazing/structs"
	"strings"
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
	bt := algos.NewKruskal(10, 20)
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
			if h==0 {
				model.Cells[0][w].CssClasses +="north "
				//model.Cells[0][w].Note += "north "
			}

			if bt.Board.Cells[h][w].IsSet(structs.EAST) {
				model.Cells[h][w].CssClasses += "east "
			}
			if bt.Board.Cells[h][w].IsSet(structs.WEST) {
				model.Cells[h][w].CssClasses += "west "
			}
			if bt.Board.Cells[h][w].IsSet(structs.NORTH) {
				model.Cells[h][w].CssClasses += "north "
			}
			if bt.Board.Cells[h][w].IsSet(structs.SOUTH) {
				model.Cells[h][w].CssClasses += "south "
			}
			/*
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
	model.Cells[0][0].CssClasses = strings.Replace(model.Cells[0][0].CssClasses, "north ","",-1)
	model.Cells[bt.Board.Height-1][bt.Board.Width-1].CssClasses = strings.Replace(model.Cells[bt.Board.Height-1][bt.Board.Width-1].CssClasses, "south ","",-1)

	err = tpl.ExecuteTemplate(w, "index.html", model)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
