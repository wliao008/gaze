package algos

import (
	"github.com/wliao008/mazing/structs"
	"math/rand"
	"time"
	"fmt"
	"strconv"
)

type Kruskal struct {
	Board structs.Board
	Set *structs.DisjointSet
	List []*ListItem
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
	k.Set = &structs.DisjointSet{}
	k.Set.Items = make(map[string]*structs.Item)
	// ~8ms

	height2 := int(k.Board.Height)
	width2 := int(k.Board.Width)
	for h := 0; h < height2; h++ {
		for w := 0; w < width2; w++ {
			//k.Board.Cells[h][w].SetBit(structs.VISITED)
			item := &structs.Item{&k.Board.Cells[h][w], nil}
			k.Set.Items[strconv.Itoa(h) + "_" + strconv.Itoa(w)] = item
		}
	}

	for h := 0; h < height2; h++ {
		for w := 0; w < width2; w++ {
			c := &k.Board.Cells[h][w]
			_, fromItem := k.Set.FindItem(c)
			cells := k.Board.Neighbors(c)
			for _, cell := range cells {
				_, toItem := k.Set.FindItem(cell)
				li := &ListItem{From: fromItem, To: toItem}
				k.List = append(k.List, li)
			}
		}
	}

	return k
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (k *Kruskal) Generate() error {
	// ~60ms
	k.Shuffle()
	for _, item := range k.List {
		root1 := k.Set.Find(item.From)
		root2 := k.Set.Find(item.To)
		if root1.Data.X == root2.Data.X &&
			root1.Data.Y == root2.Data.Y {
			continue
		}

		dir := k.Board.GetDirection(item.From.Data, item.To.Data)
		k.Board.BreakWall(item.From.Data, item.To.Data, dir)

		_ = k.Set.Union(root1, root2)
		item.From.Data.SetBit(structs.VISITED)
		item.To.Data.SetBit(structs.VISITED)
	}

	return nil
}

func (k *Kruskal) Shuffle() {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	for i := len(k.List) - 1; i > 0; i-- {
		j := rand.Intn(i)
		k.List[i], k.List[j] = k.List[j], k.List[i]
	}
}

