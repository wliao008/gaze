package main

import (
	"fmt"
	"github.com/wliao008/mazing/algos"
	"github.com/wliao008/mazing/structs"
	"os"
)

func main() {
	bt := algos.Kruskal{40, 20, nil}
	err := bt.Generate()
	if err != nil {
		fmt.Println("ERROR")
	}
	bt.Write(os.Stdout)
	//Display(&bt, cells)
}


func Display(bt *algos.BackTracking, cells [][]structs.Cell) {
	fmt.Printf("  ")
	for i := 1; i < bt.Width; i++ {
		fmt.Printf(" _")
	}
	fmt.Printf("\n")

	for j := int(0); j < bt.Height; j++ {
		fmt.Printf("|")
		for h := int(0); h < bt.Width; h++ {
			c := cells[h][j]
			if h == bt.Width-1 && j == bt.Height-1 {
				fmt.Printf(" |")
				break
			}
			if c.IsSet(structs.SOUTH) {
				fmt.Printf("_")
			} else {
				fmt.Printf(" ")
			}

			if c.IsSet(structs.EAST) {
				fmt.Printf("|")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
