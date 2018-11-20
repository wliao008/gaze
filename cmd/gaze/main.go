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
	help   bool
}

// parses and return command line arguments
func parseArguments() *argument {
	var height = flag.Uint("h", uint(defaultHeight), fmt.Sprintf("maze height, e.g. -h %d", defaultHeight))
	var width = flag.Uint("w", uint(defaultWidth), fmt.Sprintf("maze width, e.g. -w %d", defaultWidth))
	var help = flag.Bool("help", false, "This is a Help Subcommand")

	flag.Parse()

	return &argument{height: uint16(*height), width: uint16(*width), help: bool(*help)}
}

func main() {
	arg := parseArguments()
	if arg.help {
		fmt.Println(helpReturn())
		os.Exit(0)
	}
	k := algos.NewPrimTriangle(arg.height, arg.width)
	k.Generate()
	k.Board.Write(os.Stdout)
	fmt.Println("")
	k.Board.WriteSvg(os.Stdout)
}

func helpReturn() string {
	return `
	Use params to improve your experience.
	- To set Heigth: -h <SIZE_IN_INTEGER>
	- To set Weigth: -w <SIZE_IN_INTEGER>

	Example: 

	$ go run main.go -h 12 -w 12
	_ _ _ _ _ _ _ _ _ _ _
	| |   |  _ _ _ _    |   |
	| | | |_ _|  _ _ _|_ _| |
	|_ _|_ _ _ _| |  _ _ _ _|
	|     |  _ _  |_ _  |_  |
	|_| | | |   |_  | |_  | |
	|  _|_ _| |  _| |  _| | |
	|_  | |  _| |  _ _|  _| |
	|   |_ _|  _| | |  _|  _|
	| |_|   |_ _ _  | |  _  |
	|  _ _|_ _ _  | | | |  _|
	| |   |_ _  | |_| |_|_  |
	|_ _|_ _ _ _|_ _ _ _ _  |
	`
}
