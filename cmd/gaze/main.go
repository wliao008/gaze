package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wliao008/gaze/algos"
)

// default maze argument
const (
	defaultHeight uint16 = 10
	defaultWidth  uint16 = 30
)

// maze options
type argument struct {
	height uint16
	width  uint16
}

// parses and return command line arguments
func parseArguments() *argument {
	var height = flag.Uint("h", uint(defaultHeight), fmt.Sprintf("maze height, e.g. -h %d", defaultHeight))
	var width = flag.Uint("w", uint(defaultWidth), fmt.Sprintf("maze width, e.g. -w %d", defaultWidth))

	flag.Parse()

	return &argument{height: uint16(*height), width: uint16(*width)}
}

func main() {
	arg := parseArguments()
	k := algos.NewPrimTriangle(arg.height, arg.width)
	k.Generate()
	k.Board.Write3a(os.Stdout)
	fmt.Printf("\n")
	k.Board.WriteVisited(os.Stdout)
	fmt.Printf("\n")
	k.Board.Write3(os.Stdout)
}
