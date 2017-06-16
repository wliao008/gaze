package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/wliao008/mazing/algos"
	"github.com/wliao008/mazing/models"
	"github.com/wliao008/mazing/structs"
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
	bt := algos.NewKruskal(5, 5)
	err := bt.Generate()
	if err != nil {
		fmt.Println("ERROR")
	}

	// create model
	model := &models.BoardModel{}
	model.Cells = make([][]models.CellModel, bt.Board.Width)
	model.RawCells = bt.Board.Cells
	for i := uint16(0); i < bt.Board.Width; i++ {
		model.Cells[i] = make([]models.CellModel, bt.Board.Height)
	}
	// initialize model
	for i := uint16(0); i < bt.Board.Height; i++ {
		for j := uint16(0); j < bt.Board.Width; j++ {
			model.Cells[j][i].X = i;
			model.Cells[j][i].Y = j
			/*if j==0{
				model.Cells[j][i].CssClasses +="north "
			}*/
			if bt.Board.Cells[j][i].IsSet(structs.EAST) {
				model.Cells[j][i].CssClasses += "east "
			}
			if bt.Board.Cells[j][i].IsSet(structs.WEST) {
				model.Cells[j][i].CssClasses += "west "
			}
			if bt.Board.Cells[j][i].IsSet(structs.NORTH) {
				model.Cells[j][i].CssClasses += "north "
			}
			if bt.Board.Cells[j][i].IsSet(structs.SOUTH) {
				model.Cells[j][i].CssClasses += "south "
			}
		}
	}

	err = tpl.ExecuteTemplate(w, "index.html", model)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}