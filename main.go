package main

import (
	"fmt"
	"github.com/wliao008/mazing/algos"
	"os"
)

func main() {
	bt := algos.NewKruskal(40, 20)
	err := bt.Generate()
	if err != nil {
		fmt.Println("ERROR")
	}
	bt.Board.Write(os.Stdout)
}
