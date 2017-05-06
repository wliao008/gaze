package main

import (
	"fmt"
	"github.com/wliao008/mazing/algos"
	"github.com/wliao008/mazing/structs"
	"os"
)

func main() {
	/*
	bt := algos.Kruskal{40, 20, nil}
	err := bt.Generate()
	if err != nil {
		fmt.Println("ERROR")
	}
	bt.Write(os.Stdout)
	*/
	bt := algos.NewKruskal(10, 5)
	err := bt.Generate()
	if err != nil {
		fmt.Println("ERROR")
	}
	bt.Write(os.Stdout)
	//fmt.Println(&bt.Board)
	//Display(&bt.Board)
	//Display(&bt, cells)
}

func PrintBoard(board *structs.Board) {
	for i := uint16(0); i < board.Width; i++ {
		for j := uint16(0); j < board.Height; j++ {
			fmt.Printf("[%d,%d,%d] ", i,j,board.Cells[i][j].Flag)
		}
		fmt.Println("")
	}
}

func Display(board *structs.Board) {
	fmt.Printf("  ")
	for i := uint16(1); i < board.Width; i++ {
		fmt.Printf(" _")
	}
	fmt.Printf("\n")

	for j := uint16(0); j < board.Height; j++ {
		fmt.Printf("|")
		for h := uint16(0); h < board.Width; h++ {
			c := board.Cells[h][j]
			if h == board.Width-1 && j == board.Height-1 {
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
