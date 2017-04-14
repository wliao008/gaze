package algos

import (
	"fmt"
	"github.com/wliao008/mazing/structs"
	"math/rand"
	"time"
)

type BackTracking struct {
	Width  int
	Height int
}

func shuffle(arr []structs.FlagPosition) {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond())) // no shuffling without this line

	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func (b *BackTracking) doWork(cells [][]structs.Cell, x, y int) {
	rand.Seed(time.Now().UTC().UnixNano())
	directions := []structs.FlagPosition{structs.NORTH, structs.SOUTH, structs.EAST, structs.WEST}
	d := structs.Direction{}
	shuffle(directions)
	for _, direction := range directions {
		var nextX int = x + d.XDirection(direction)
		var nextY int = y + d.YDirection(direction)
		if nextX >= 0 && nextX < b.Width &&
			nextY >= 0 && nextY < b.Height &&
			!cells[nextX][nextY].IsSet(structs.VISITED) {
			fmt.Println("direction: %v, nextX=%v, nextY=%v", direction, nextX, nextY)
			b.carvePassage(direction, &cells[x][y], &cells[nextX][nextY])
			b.doWork(cells, nextX, nextY)
		}
	}
}

func (b *BackTracking) Generate() ([][]structs.Cell, error) {
	fmt.Println("generating %d x %d", b.Width, b.Height)
	cells := make([][]structs.Cell, b.Width)
	for i := 0; i < b.Width; i++ {
		cells[i] = make([]structs.Cell, b.Height)
	}

	//init
	for i := 0; i < b.Width; i++ {
		for j := 0; j < b.Height; j++ {
			cells[i][j].Flag = 15
		}
	}

	b.doWork(cells, 0, 0)
	b.Display(cells)
	return cells, nil
}

func (b *BackTracking) carvePassage(dir structs.FlagPosition, from, to *structs.Cell) {
	from.SetBit(structs.VISITED)
	to.SetBit(structs.VISITED)
	//fmt.Println("dir=%v, going from %v to %v", dir, from, to)

	switch dir {
	case structs.NORTH:
		from.ClearBit(structs.NORTH)
		to.ClearBit(structs.SOUTH)
	case structs.SOUTH:
		from.ClearBit(structs.SOUTH)
		to.ClearBit(structs.NORTH)
	case structs.EAST:
		from.ClearBit(structs.EAST)
		to.ClearBit(structs.WEST)
	case structs.WEST:
		from.ClearBit(structs.WEST)
		to.ClearBit(structs.EAST)
	}
}

func (b *BackTracking) Display(cells [][]structs.Cell) {
	fmt.Println(cells)

	fmt.Printf("  ")
	for i := 1; i < b.Width; i++ {
		fmt.Printf(" _")
	}
	fmt.Printf("\n")

	for j := int(0); j < b.Height; j++ {
		fmt.Printf("|")
		for h := int(0); h < b.Width; h++ {
			c := cells[h][j]
			if h == b.Width -1 && j == b.Height -1 {
				fmt.Printf(" |")
				break
			}
			if c.IsSet(structs.SOUTH){
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

func (b *BackTracking) Test() {
	directions := []structs.FlagPosition{structs.NORTH, structs.SOUTH, structs.EAST, structs.WEST}
	n := rand.Intn(len(directions))
	fmt.Println(directions[n])
}
