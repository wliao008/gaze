package algos

import (
	"github.com/wliao008/mazing/structs"
	"math/rand"
	"time"
	"fmt"
	"os"
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
	return fmt.Sprintf("[%d,%d] (%p) to [%d,%d] (%p)", from.X, from.Y, &li.From, to.X, to.Y, &li.To)
}

func NewKruskal(height, width uint16) *Kruskal {
	k := &Kruskal{Board: structs.Board{height, width, nil}}
	k.Board.Init()
	return k
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (k *Kruskal) Test2() {
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
			ok, fromItem := ds.FindItem(c)
			if !ok {
				fmt.Printf("fromItem not found\n")
			}
			cells := k.Board.Neighbors(c)
			for _, cell := range cells {
				ok, toItem := ds.FindItem(cell)
				if !ok {
					fmt.Printf("toItem not found\n")
				}
				li := &ListItem{From: fromItem, To: toItem}
				list = append(list, li)
			}
		}
	}

	for _, li := range list {
		fmt.Println(li)
	}
	fmt.Println("-----------")
	k.Shuffle(list)
	for _, item := range list {
		from := item.From.Data
		to := item.To.Data
		fmt.Printf("[%d,%d] (item %p) - [%d,%d] (item %p)\n", from.X, from.Y, &item, to.X, to.Y, &item)
	}
	for _, item := range list {
		root1 := ds.Find(item.From)
		root2 := ds.Find(item.To)
		if root1.Data.IsSet(structs.VISITED) &&
			root2.Data.IsSet(structs.VISITED) {
			continue
		}

		newroot := ds.Union(root1, root2)
		fmt.Printf("new root is %v\n", newroot)
		root1.Data.SetBit(structs.VISITED)
		root2.Data.SetBit(structs.VISITED)
		ds.Write(os.Stdout)
		fmt.Println("-------------")
	}

	ds.Write(os.Stdout)
	fmt.Println("-----------")
}

func (k *Kruskal) Generate() error {
	return nil
}

//func (k *Kruskal) Test() {
//	fmt.Println("Test()")
//	var list []*ListItem
//	ds := &structs.DisjointSet{}
//	for h := uint16(0); h < k.Board.Height; h++ {
//		for w := uint16(0); w < k.Board.Width; w++ {
//			c := &k.Board.Cells[h][w]
//			itm := &structs.Item{c, nil}
//			ds.Items = append(ds.Items, itm)
//			fmt.Printf("[%d,%d] set created, item addr=%p\n", c.X, c.Y, &itm)
//			cells := k.Board.Neighbors(c)
//			for _, cell := range cells {
//				cel := &structs.Item{cell, nil}
//				item := &ListItem{From: itm, To: cel}
//				list = append(list, item)
//			}
//		}
//	}
//	//k.Shuffle(list)
//	fmt.Println("----------")
//	for _, item := range list {
//		from := item.From.Data.(*structs.Cell)
//		to := item.To.Data.(*structs.Cell)
//		fmt.Printf("[%d,%d] (item %p) - [%d,%d] (item %p)\n", from.X, from.Y, &item, to.X, to.Y, &item)
//	}
//	fmt.Println("----------")
//	for _, item := range list {
//		/*
//		from := item.From.Data.(*structs.Cell)
//		to := item.To.Data.(*structs.Cell)
//		root1 := ds.Find(item.From)
//		root2 := ds.Find(item.To)
//		rootCell := root1.Data.(*structs.Cell)
//		rootCell2 := root2.Data.(*structs.Cell)
//		fmt.Printf("From [%d,%d] -> [%d,%d], To [%d,%d] -> [%d,%d]\n", from.X, from.Y, rootCell.X, rootCell.Y, to.X, to.Y, rootCell2.X, rootCell2.Y)
//		if rootCell.X != rootCell2.X || rootCell.Y != rootCell2.Y {
//			ds.Union(root1, root2)
//		}*/
//		fmt.Println(item)
//		root1 := ds.Find(item.From)
//		root2 := ds.Find(item.To)
//		fmt.Printf("\troot1: %v, item addr=%p\n", root1, &root1)
//		fmt.Printf("\troot2: %v, item addr=%p\n", root2, &root2)
//		rootCell1 := root1.Data.(*structs.Cell)
//		rootCell2 := root2.Data.(*structs.Cell)
//		if rootCell1.IsSet(structs.VISITED) &&
//			rootCell2.IsSet(structs.VISITED) {
//			continue
//		}
//
//		newroot := ds.Union(root1, root2)
//		fmt.Printf("new root is %v\n", newroot)
//		rootCell1.SetBit(structs.VISITED)
//		rootCell2.SetBit(structs.VISITED)
//		ds.Write(os.Stdout)
//		fmt.Println("-------------")
//	}
//	fmt.Println("-------------")
//	ds.Write(os.Stdout)
//
//	tmpItem := ds.Items[0]
//	tmpRoot := ds.Find(tmpItem)
//	fmt.Printf("tmpRoot: %v\n", tmpRoot)
//}

func (k *Kruskal) Shuffle(arr []*ListItem) {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

