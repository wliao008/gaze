package main

import (
	"os"

	"github.com/wliao008/gaze/algos"
)

func main() {
	k := algos.NewKruskalWeave(3, 3)
	k.Generate()
	k.Board.Write(os.Stdout)
}
