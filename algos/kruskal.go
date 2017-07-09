package algos

import (
	"github.com/wliao008/mazing/structs"
	"math/rand"
	"time"
	"fmt"
	//"strconv"
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
	ds.Items = make(map[string]*structs.Item)
	tmp := make(map[string]int)
	var list []*ListItem
	height := int(k.Board.Height)
	width := int(k.Board.Width)
	mapcount := 0
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			//k.Board.Cells[h][w].SetBit(structs.VISITED)
			item := &structs.Item{&k.Board.Cells[h][w], nil}
			key := fmt.Sprintf("%d_%d", h, w)
			if key == "143" {
				fmt.Printf("%d,%d\n", h, w)
			}
			ds.Items[key] = item
			tmp[fmt.Sprintf("%d_%d", h, w)] += 1
			mapcount += 1
		}
	}

	for k, v := range tmp {
		if v > 1 {
			fmt.Printf("%s=%d\n", k, v)
		}
	}

	fmt.Printf("ds.Items = %d, mapcount=%d\n", len(ds.Items), mapcount)

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
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
	idx := 0
	for _, item := range list {
		root1 := ds.Find(item.From)
		root2 := ds.Find(item.To)
		if root1.Data.X == root2.Data.X &&
			root1.Data.Y == root2.Data.Y {
			idx += 1
			continue
		}

		dir := k.Board.GetDirection(item.From.Data, item.To.Data)
		k.Board.BreakWall(&k.Board.Cells[item.From.Data.X][item.From.Data.Y], 
				&k.Board.Cells[item.To.Data.X][item.To.Data.Y], dir)

		_ = ds.Union(root1, root2)
		k.Board.Cells[item.From.Data.X][item.From.Data.Y].SetBit(structs.VISITED)
		k.Board.Cells[item.To.Data.X][item.To.Data.Y].SetBit(structs.VISITED)
		idx += 1
	}
	/*
	*/
	fmt.Printf("list = %d, idx=%d\n", len(list), idx)

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

