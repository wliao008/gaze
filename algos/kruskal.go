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
	for h := uint16(0); h < k.Board.Height; h++ {
		for w := uint16(0); w < k.Board.Width; w++ {
			cells := k.Board.Neighbors(&k.Board.Cells[h][w])
			for _, cell := range cells {
				fmt.Printf("[%d,%d] - [%d,%d]\n", h, w, cell.X, cell.Y)
				item := &util.Item{From: &k.Board.Cells[h][w], To: cell, Parent: nil}
				ds.Items = append(ds.Items, item)
			}
		}
	}
	k.Shuffle(ds.Items)
	fmt.Println("----------")
	for _, item := range ds.Items {
		from := item.From.(*structs.Cell)
		to := item.To.(*structs.Cell)
		fmt.Printf("[%d,%d] - [%d,%d]\n", from.X, from.Y, to.X, to.Y)
	}

}

func (k *Kruskal) Shuffle(arr []*util.Item) {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

