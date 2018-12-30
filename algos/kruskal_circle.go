package algos

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/wliao008/gaze"
)

type KruskalCircle struct {
	Name  string
	Board gaze.BoardCircle
	Set   *gaze.DisjointSet
	List  []*ListItem
}

// type ListItem struct {
// 	From *gaze.Item
// 	To   *gaze.Item
// }

// func (li *ListItem) String() string {
// 	from := li.From.Data
// 	to := li.To.Data
// 	return fmt.Sprintf("[%d,%d] to [%d,%d]", from.X, from.Y, to.X, to.Y)
// }

func NewKruskalCircle(height, width uint16) *KruskalCircle {
	k := &KruskalCircle{Board: gaze.BoardCircle{height, width, nil}}
	k.Name = "kruskal algorithm (circle)"
	k.Board.Init()
	k.Set = &gaze.DisjointSet{}
	k.Set.Items = make(map[string]*gaze.Item)
	// ~8ms, ~80k allocations, ~2mb
	h := uint16(0)
	w := uint16(0)
	for h = uint16(0); h < height; h++ {
		for w = uint16(0); w < width; w++ {
			item := &gaze.Item{&k.Board.Cells[h][w], nil}
			k.Set.Items[fmt.Sprintf("%d_%d", h, w)] = item
		}
	}

	for h = uint16(0); h < height; h++ {
		for w = uint16(0); w < width; w++ {
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

	k.Shuffle()
	return k
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (k *KruskalCircle) Generate() error {
	// ~60ms, ~40k allocation, ~1mb
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
		item.From.Data.SetBit(gaze.VISITED)
		item.To.Data.SetBit(gaze.VISITED)
	}

	return nil
}

func (k *KruskalCircle) Shuffle() {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	for i := len(k.List) - 1; i > 0; i-- {
		j := rand.Intn(i)
		k.List[i], k.List[j] = k.List[j], k.List[i]
	}
}
