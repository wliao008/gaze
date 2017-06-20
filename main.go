package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/wliao008/mazing/algos"
	"github.com/wliao008/mazing/models"
	"github.com/wliao008/mazing/structs"
	"strings"
	"io"
	"time"
	"strconv"
	_ "os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("web/templates/*.tmpl"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/home", homeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, req *http.Request){
	http.Redirect(w, req, "/home", http.StatusSeeOther)
}

func homeHandler(w http.ResponseWriter, req *http.Request){
	height, width := getSize(w, req)
	bt := algos.NewKruskal(height, width)
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
			if h==0 {
				model.Cells[0][w].CssClasses +="north "
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
		}
	}

	//set the openning and ending cell
	model.Cells[0][0].CssClasses = strings.Replace(model.Cells[0][0].CssClasses, "north ","",-1)
	model.Cells[bt.Board.Height-1][bt.Board.Width-1].CssClasses = strings.Replace(model.Cells[bt.Board.Height-1][bt.Board.Width-1].CssClasses, "south ","",-1)

	err = tpl.ExecuteTemplate(w, "index.tmpl", model)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func staticHandler(w http.ResponseWriter, req *http.Request) {
	static_file := req.URL.Path[len("/static/css/"):]
	fmt.Println(static_file)
	f, err := http.Dir("/web/static/css/").Open("style.css")
	if err == nil {
		content := io.ReadSeeker(f)
		http.ServeContent(w, req, "/web/static/css/style.css", time.Now(), content)
		return
	}
	http.NotFound(w, req)
}

func getSize(w http.ResponseWriter, req *http.Request) (uint16, uint16) {
	req.ParseForm()
	height := uint16(20)
	width := uint16(40)
	if val, ok := req.Form["height"]; ok {
		h, _ := strconv.ParseInt(val[0], 10, 0)
		height = uint16(h)
	}
	if val, ok := req.Form["width"]; ok {
		w, _ := strconv.ParseInt(val[0], 10, 0)
		width = uint16(w)
	}
	return height, width
}

