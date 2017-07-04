package algos

import (
	"github.com/wliao008/mazing/structs"
	"github.com/wliao008/mazing/util"
	"math/rand"
	"time"
	"fmt"
)

type Kruskal struct {
	Board structs.Board
}

func NewKruskal(height, width uint16) *Kruskal {
	k := &Kruskal{Board: structs.Board{height, width, nil}}
	k.Board.Init()
	return k
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (k *Kruskal) Test() {
	fmt.Println("Test()")
	ds := &util.DisjointSet{}
	fmt.Println(ds)
	for h := uint16(0); h < k.Board.Height; h++ {
		for w := uint16(0); w < k.Board.Width; w++ {
			cells := k.Board.Neighbors(&k.Board.Cells[h][w])
			for _, cell := range cells {
				fmt.Printf("[%d,%d] - [%d,%d]\n", h, w, cell.X, cell.Y)
				
			}
		}
	}
}
