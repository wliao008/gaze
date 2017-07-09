package algos

import (
	"github.com/wliao008/mazing/structs"
	"math/rand"
	"time"
	"fmt"
)

type Kruskal struct {
	Board structs.Board
}

type ListItem struct {
	From *structs.Item
	To *structs.Item
}

func (li *ListItem) String() string {
	from := li.From.Data
	to := li.To.Data
	return fmt.Sprintf("[%d,%d] to [%d,%d]", from.X, from.Y, to.X, to.Y)
}

func NewKruskal(height, width uint16) *Kruskal {
	k := &Kruskal{Board: structs.Board{height, width, nil}}
	k.Board.Init()
	return k
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (k *Kruskal) Generate() error {
	ds := &structs.DisjointSet{}
	var list []*ListItem
	for h := uint16(0); h < k.Board.Height; h++ {
		for w := uint16(0); w < k.Board.Width; w++ {
			item := &structs.Item{&k.Board.Cells[h][w], nil}
			ds.Items = append(ds.Items, item)
		}
	}

	for h := uint16(0); h < k.Board.Height; h++ {
		for w := uint16(0); w < k.Board.Width; w++ {
			c := &k.Board.Cells[h][w]
			_, fromItem := ds.FindItem(c)
			cells := k.Board.Neighbors(c)
			for _, cell := range cells {
				_, toItem := ds.FindItem(cell)
				li := &ListItem{From: fromItem, To: toItem}
				list = append(list, li)
			}
		}
	}

	k.Shuffle(list)
	for _, item := range list {
		root1 := ds.Find(item.From)
		root2 := ds.Find(item.To)
		if root1.Data.X == root2.Data.X &&
			root1.Data.Y == root2.Data.Y {
			continue
		}

		dir := k.Board.GetDirection(item.From.Data, item.To.Data)
		k.Board.BreakWall(&k.Board.Cells[item.From.Data.X][item.From.Data.Y], 
				&k.Board.Cells[item.To.Data.X][item.To.Data.Y], dir)

		_ = ds.Union(root1, root2)
		root1.Data.SetBit(structs.VISITED)
		root2.Data.SetBit(structs.VISITED)
	}

	return nil
}

func (k *Kruskal) Shuffle(arr []*ListItem) {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

