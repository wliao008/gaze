package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/wliao008/mazing/algos"
	"github.com/wliao008/mazing/models"
	"github.com/wliao008/mazing/structs"
	"github.com/wliao008/mazing/solvers"
	"strings"
	"io"
	"math/rand"
	"time"
	"strconv"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("web/templates/*.tmpl"))
}

func main_con() {
	k := algos.NewKruskalWeave(3, 3)
	k.Generate()
	k.Board.Write(os.Stdout)
}

func main2() {
	bt := algos.NewPrim(3, 3)
	err := bt.Generate()
	if err != nil {
		fmt.Println("ERROR")
	}
	bt.Board.Cells[0][0].ClearBit(structs.NORTH)
	bt.Board.Cells[bt.Board.Height-1][bt.Board.Width-1].ClearBit(structs.SOUTH)
	bt.Board.Write(os.Stdout)
	def := solvers.DeadEndFiller{}
	def.Board = &bt.Board
	def.Solve()
	bt.Board.Write2(os.Stdout)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))
	http.ListenAndServe(":8080", nil)
}

func faviconHandler(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "web/favicon.ico")
}

func indexHandler(w http.ResponseWriter, req *http.Request){
	http.Redirect(w, req, "/home", http.StatusSeeOther)
}

func aboutHandler(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w, "about.tmpl", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func getBoard(height, width uint16) (*structs.Board, string) {
	rand.Seed(time.Now().UTC().UnixNano())
	var idx int = rand.Intn(3)
	board := &structs.Board{}
	idx=0
	if idx == 0 {
		k := algos.NewKruskalWeave(height, width)
		start := time.Now()
		_ = k.Generate()
		board = &k.Board
		elasped := time.Since(start)
		return board, fmt.Sprintf("%s, took %s", k.Name, elasped)
	} else if idx == 1 {
		k := algos.NewKruskal(height, width)
		start := time.Now()
		_ = k.Generate()
		board = &k.Board
		elasped := time.Since(start)
		return board, fmt.Sprintf("%s, took %s", k.Name, elasped)
	} else {
		p := algos.NewPrim(height, width)
		start := time.Now()
		_ = p.Generate()
		board = &p.Board
		elasped := time.Since(start)
		return board, fmt.Sprintf("prim algorithm, took %s", elasped)
	}
}

func homeHandler(w http.ResponseWriter, req *http.Request){
	height, width := getSize(w, req)

	board, name := getBoard(height, width)
	/*
	bt := algos.NewKruskal(height, width)
	err := bt.Generate()
	if err != nil {
		fmt.Println("ERROR")
	}
	*/
	board.Cells[0][0].ClearBit(structs.NORTH)
	board.Cells[height-1][width-1].ClearBit(structs.SOUTH)
	def := solvers.DeadEndFiller{}
	def.Board = board
	def.Solve()
	//bt.Board.Write(os.Stdout)

	// create model
	model := &models.BoardModel{}
	model.TableCss = "cb2"
	if !strings.Contains(name, "weave") {
		model.TableCss = "cb"
	}
	model.Name = name
	model.Height = height
	model.Width = width
	model.Cells = make([][]models.CellModel, height)
	model.RawCells = board.Cells
	for i := uint16(0); i < height; i++ {
		model.Cells[i] = make([]models.CellModel, width)
	}

	// initialize model
	for w := uint16(0); w < width; w++ {
		model.Cells[0][w].CssClasses += "north "
		model.Cells[height-1][w].CssClasses += "south "
	}

	for h := uint16(0); h < height; h++ {
		model.Cells[h][0].CssClasses +="west "
		for w := uint16(0); w < width; w++ {
			model.Cells[h][w].X = w;
			model.Cells[h][w].Y = h
			if !board.Cells[h][w].IsSet(structs.DEAD){
				model.Cells[h][w].CssClasses += "p "
			}
			if w == width - 1 {
				model.Cells[h][w].CssClasses +="east "
			}
			if h==0 {
				model.Cells[0][w].CssClasses +="north "
			}

			if board.Cells[h][w].IsSet(structs.EAST) {
				model.Cells[h][w].CssClasses += "east "
			}
			if board.Cells[h][w].IsSet(structs.WEST) {
				model.Cells[h][w].CssClasses += "west "
			}
			if board.Cells[h][w].IsSet(structs.NORTH) {
				model.Cells[h][w].CssClasses += "north "
			}
			if board.Cells[h][w].IsSet(structs.SOUTH) {
				model.Cells[h][w].CssClasses += "south "
			}

			if board.Cells[h][w].IsSet(structs.CROSS) {
				if board.Cells[h][w].IsSet(structs.EAST) {
					model.Cells[h][w+1].CssClasses += "west2 "
				}
				if board.Cells[h][w].IsSet(structs.WEST) {
					model.Cells[h][w-1].CssClasses += "east2 "
				}
				if board.Cells[h][w].IsSet(structs.NORTH) {
					model.Cells[h-1][w].CssClasses += "south2 "
				}
				if board.Cells[h][w].IsSet(structs.SOUTH) {
					model.Cells[h+1][w].CssClasses += "north2 "
				}
			}
		}
	}

	//set the openning and ending cell
	model.Cells[0][0].CssClasses = strings.Replace(model.Cells[0][0].CssClasses, "north ","",-1)
	model.Cells[height-1][width-1].CssClasses = strings.Replace(model.Cells[height-1][width-1].CssClasses, "south ","",-1)


	model = processWeaveMaze(board, name)
	err := tpl.ExecuteTemplate(w, "index.tmpl", model)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func processWeaveMaze(board *structs.Board, name string) *models.BoardModel {
	// create model
	height := board.Height + (board.Height - 1)
	width := board.Width + (board.Width - 1)
	model := &models.BoardModel{}
	model.TableCss = "cb"
	model.Name = name
	model.Height = height
	model.Width = width
	model.Cells = make([][]models.CellModel, height)
	model.RawCells = board.Cells
	for i := uint16(0); i < height; i++ {
		model.Cells[i] = make([]models.CellModel, width)
	}

	// initialize model
	for h := uint16(0); h < height; h++ {
		for w := uint16(0); w < width; w++ {
			model.Cells[h][w].X = w;
			model.Cells[h][w].Y = h
			model.Cells[h][w].CssClasses = " north south east west "
			if h % 2 == 1 || w % 2 == 1 {
				model.Cells[h][w].CssClasses += "sp "
			} else {
				model.Cells[h][w].CssClasses += "td "
			}
		}
	}

	/*
	for h := uint16(0); h < board.Height; h++ {
		for w := uint16(0); w < board.Width; w++ {
			c := board.Cells[h][w]
			//mcr := model.Cells[h][w*2+1]
			if !c.IsSet(structs.EAST) {
				model.Cells[h][w*2].CssClasses += strings.Replace(model.Cells[h][w*2].CssClasses, "east", "", -1)
				//mcr.CssClasses += strings.Replace(mcr.CssClasses, "west", "", -1)
				fmt.Println("process")
			}
		}
	}
	*/

	//set the openning and ending cell
	model.Cells[0][0].CssClasses = strings.Replace(model.Cells[0][0].CssClasses, "north ","",-1)
	model.Cells[height-1][width-1].CssClasses = strings.Replace(model.Cells[height-1][width-1].CssClasses, "south ","",-1)
	return model
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
	
	if len(req.Form) == 0 {
		// no new size specified by user
		cookieSize, err := req.Cookie("size")
		if err == nil {
			size := strings.Split(cookieSize.Value, ",")
			heightNew, _ := strconv.ParseUint(size[0], 10, 16)
			widthNew, _ := strconv.ParseUint(size[1], 10, 16)
			height = uint16(heightNew)
			width = uint16(widthNew)
		}
	} else {
		if val, ok := req.Form["height"]; ok {
			h, _ := strconv.ParseInt(val[0], 10, 0)
			height = uint16(h)
		}
		if val, ok := req.Form["width"]; ok {
			w, _ := strconv.ParseInt(val[0], 10, 0)
			width = uint16(w)
		}
	}
	expiration := time.Now().Add(365 * 24 * time.Hour)
	value := fmt.Sprintf("%d,%d", height, width)
	cookieSize := &http.Cookie{Name: "size", Value: value, Expires: expiration}
	http.SetCookie(w, cookieSize)
	return height, width
}

