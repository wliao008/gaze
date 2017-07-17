package algos

import (
	"github.com/wliao008/mazing/structs"
	"math/rand"
	"time"
	"fmt"
	"os"
)

type KruskalWeave struct {
	Name string
	Board structs.Board
	Set *structs.DisjointSet
	List []*ListItem
}

func NewKruskalWeave(height, width uint16) *KruskalWeave {
	k := &KruskalWeave{Board: structs.Board{height, width, nil}}
	k.Name = "kruskal weave algorithm"
	k.Board.Init()
	k.Set = &structs.DisjointSet{}
	k.Set.Items = make(map[string]*structs.Item)
	// ~8ms, ~80k allocations, ~2mb
	h := uint16(0)
	w := uint16(0)
	for h = uint16(0); h < height; h++ {
		for w = uint16(0) ; w < width; w++ {
			//k.Board.Cells[h][w].SetBit(structs.VISITED)
			item := &structs.Item{&k.Board.Cells[h][w], nil, nil}
			k.Set.Items[fmt.Sprintf("%d_%d", h, w)] = item
		}
	}
	//fmt.Printf("sets created: %d\n", len(k.Set.Items))
	k.preprocess()
	fmt.Println("sets")
	k.Set.Write(os.Stdout)
	//fmt.Printf("after pre-process\n")
	k.Board.Write(os.Stdout)

	for h = uint16(0); h < height; h++ {
		for w = uint16(0); w < width; w++ {
			c := &k.Board.Cells[h][w]
			if c.IsSet(structs.CROSS) {
				c.SetBit(structs.VISITED)
				continue
			}

			_, fromItem := k.Set.FindItem(c)
			cells := k.Board.Neighbors(c)
			for _, cell := range cells {
				if cell.IsSet(structs.CROSS) {
					continue
				}
				_, toItem := k.Set.FindItem(cell)
				li := &ListItem{From: fromItem, To: toItem}
				k.List = append(k.List, li)
			}
		}
	}
	//fmt.Printf("k.List=%d\n", len(k.List))
	//for _, li := range k.List {
	//	fmt.Println(li)
	//}

	return k
}

func (k *KruskalWeave) preprocess() {
	//fmt.Println("preprocessing")
	h := uint16(0)
	w := uint16(0)
	for h = uint16(1); h < k.Board.Height-1; h++ {
		for w = uint16(1) ; w < k.Board.Width-1; w++ {
			c := &k.Board.Cells[h][w]
			fmt.Printf("%v -----------\n", c)
			neighbors := k.Board.Neighbors(c)
			crossed := 0
			for _, neighbor := range neighbors {
				fmt.Printf("\t%v\n", neighbor)
				if neighbor.IsSet(structs.CROSS) {
					crossed += 1
				}
			}
			if crossed == 0 {
				c.SetBit(structs.CROSS)
				fmt.Printf("\tmarked as CROSS\n")
				var idx int = rand.Intn(2)
				if idx == 0 {
					//fmt.Printf("\tthis cell is marked as CROSS H\n")
					left := &k.Board.Cells[c.X][c.Y-1]
					right := &k.Board.Cells[c.X][c.Y+1]
					_, fromItem := k.Set.FindItem(left)
					_, toItem := k.Set.FindItem(c)
					root1 := k.Set.FindTail(fromItem)
					root2 := k.Set.Find(toItem)
					_ = k.Set.Union(root1, root2)

					_, toItem2 := k.Set.FindItem(right)
					root1b := k.Set.FindTail(toItem)
					root2b := k.Set.Find(toItem2)
					_ = k.Set.Union(root1b, root2b)

					up := &k.Board.Cells[c.X-1][c.Y]
					down := &k.Board.Cells[c.X+1][c.Y]
					_, upItem := k.Set.FindItem(up)
					_, downItem := k.Set.FindItem(down)
					rootUp := k.Set.FindTail(upItem)
					rootDown := k.Set.Find(downItem)
					_ = k.Set.Union(rootUp, rootDown)

				}else {
					//fmt.Printf("\tthis cell is marked as CROSS V\n")
					up := &k.Board.Cells[c.X-1][c.Y]
					down := &k.Board.Cells[c.X+1][c.Y]
					_, fromItem := k.Set.FindItem(up)
					_, toItem := k.Set.FindItem(c)
					root1 := k.Set.FindTail(fromItem)
					root2 := k.Set.Find(toItem)
					_ = k.Set.Union(root1, root2)

					_, toItem2 := k.Set.FindItem(down)
					root1b := k.Set.FindTail(toItem)
					root2b := k.Set.Find(toItem2)
					_ = k.Set.Union(root1b, root2b)

					left := &k.Board.Cells[c.X][c.Y-1]
					right := &k.Board.Cells[c.X][c.Y+1]
					_, leftItem := k.Set.FindItem(left)
					_, rightItem := k.Set.FindItem(right)
					rootLeft := k.Set.FindTail(leftItem)
					rootRight := k.Set.Find(rightItem)
					_ = k.Set.Union(rootLeft, rootRight)
				}

				k.Board.Break2Walls(c, idx)
			}
		}
	}
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (k *KruskalWeave) Generate() error {
	// ~60ms, ~40k allocation, ~1mb
	/*
	*/
	for _, item := range k.List {
		fmt.Printf("%v to %v\n", item.From, item.To)
		root1 := k.Set.Find(item.From)
		root2 := k.Set.Find(item.To)
		if root1.Data.X == root2.Data.X &&
			root1.Data.Y == root2.Data.Y {
			fmt.Printf("ignoring [%d,%d]\n", root1.Data.X, root1.Data.Y)
			continue
		}

		dir := k.Board.GetDirection(item.From.Data, item.To.Data)
		k.Board.BreakWall(item.From.Data, item.To.Data, dir)

		_, tailItem := k.Set.FindItem(item.From.Data)
		tail := k.Set.FindTail(tailItem)
		_ = k.Set.Union(tail, root2)
		fmt.Printf("\tconnect set %v (tail) -> %v (root)\n", tail, root2)
		k.Board.Write(os.Stdout)
		k.Set.Write(os.Stdout)
		fmt.Println("")
		item.From.Data.SetBit(structs.VISITED)
		item.To.Data.SetBit(structs.VISITED)
	}
	//fmt.Printf("after Generate\n")
	//k.Set.Write(os.Stdout)
	return nil
}

func (k *KruskalWeave) Shuffle() {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	for i := len(k.List) - 1; i > 0; i-- {
		j := rand.Intn(i)
		k.List[i], k.List[j] = k.List[j], k.List[i]
	}
}

