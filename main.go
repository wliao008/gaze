package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/wliao008/mazing/algos"
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
	//bt.Board.Write(os.Stdout)

	model := struct {
		Name string
		Algo *algos.Kruskal
	}{
		"Wei",
		bt,
	}
	err = tpl.ExecuteTemplate(w, "index.html", model)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
