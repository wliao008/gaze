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
	bt := algos.NewKruskal(40, 20)
	err := bt.Generate()
	if err != nil {
		fmt.Println("ERROR")
	}

	// create model
	model := &models.BoardModel{}
	model.Cells = make([][]models.CellModel, bt.Board.Width)
	for i := uint16(0); i < bt.Board.Width; i++ {
		model.Cells[i] = make([]models.CellModel, bt.Board.Height)
	}

	// initialize model
	for i := uint16(0); i < bt.Board.Width; i++ {
		for j := uint16(0); j < bt.Board.Height; j++ {
			model.Cells[i][j].X = i;
			model.Cells[i][j].Y = j
			if bt.Board.Cells[i][j].IsSet(structs.EAST) {
				model.Cells[i][j].CssClasses = "east "
			}
			if bt.Board.Cells[i][j].IsSet(structs.WEST) {
				model.Cells[i][j].CssClasses += "west "
			}
			if bt.Board.Cells[i][j].IsSet(structs.NORTH) {
				model.Cells[i][j].CssClasses += "north "
			}
			if bt.Board.Cells[i][j].IsSet(structs.SOUTH) {
				model.Cells[i][j].CssClasses += "south "
			}
		}
	}

	err = tpl.ExecuteTemplate(w, "index.html", model)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
