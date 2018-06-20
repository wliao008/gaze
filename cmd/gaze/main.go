package main

import (
	"os"

	"github.com/wliao008/gaze/algos"
)

func main() {
	k := algos.NewPrim(10, 30)
	k.Generate()
	k.Board.Write(os.Stdout)
}
