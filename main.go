package main

import (
	"github.com/wliao008/mazing/algos"
	"os"
	"fmt"
)

func main() {
	bt := algos.NewBackTracking(10, 5)
	err := bt.Generate()
	if err != nil {
		fmt.Println("ERROR")
	}
	bt.Board.Write(os.Stdout)
}
