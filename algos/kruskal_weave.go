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
			item := &structs.Item{&k.Board.Cells[h][w], nil}
			k.Set.Items[fmt.Sprintf("%d_%d", h, w)] = item
		}
	}
	fmt.Printf("sets created: %d\n", len(k.Set.Items))
	k.preprocess()
	k.Set.Write(os.Stdout)
	return k
}

func (k *KruskalWeave) preprocess() {
	fmt.Println("preprocessing")
	h := uint16(0)
	w := uint16(0)
	for h = uint16(1); h < k.Board.Height-1; h++ {
		for w = uint16(1) ; w < k.Board.Width-1; w++ {
			c := &k.Board.Cells[h][w]
			fmt.Printf("%v -----------\n", c)
			neighbors := k.Board.Neighbors(c)
			crossed := 0
			for _, neighbor := range neighbors {
				//fmt.Printf("\t%v\n", neighbor)
				if neighbor.IsSet(structs.CROSS) {
					crossed += 1
				}
			}
			corners := k.Board.CornerNeighbors(c)
			fmt.Println("\t-----------")
			for _, corner := range corners {
				//fmt.Printf("\t%v\n", corner)
				if corner.IsSet(structs.CROSS) {
					crossed += 1
				}
			}
			if crossed == 0 {
				var idx int = rand.Intn(2)
				//ignore the walls breaking if 0
				var idx2 int = rand.Intn(2)
				if idx2 == 1 {
					c.SetBit(structs.CROSS)
					if idx == 0 {
						fmt.Printf("\tthis cell is marked as CROSS H\n")
						left := &k.Board.Cells[c.X][c.Y-1]
						right := &k.Board.Cells[c.X][c.Y+1]
						_, fromItem := k.Set.FindItem(left)
						_, toItem := k.Set.FindItem(c)
						root1 := k.Set.Find(fromItem)
						root2 := k.Set.Find(toItem)
						_ = k.Set.Union(root1, root2)
						_, toItem2 := k.Set.FindItem(right)
						root1b := k.Set.Find(toItem)
						root2b := k.Set.Find(toItem2)
						_ = k.Set.Union(root1b, root2b)

						up := &k.Board.Cells[c.X-1][c.Y]
						down := &k.Board.Cells[c.X+1][c.Y]
						_, upItem := k.Set.FindItem(up)
						_, downItem := k.Set.FindItem(down)
						rootUp := k.Set.Find(upItem)
						rootDown := k.Set.Find(downItem)
						_ = k.Set.Union(root2b, rootUp)
						_ = k.Set.Union(rootUp, rootDown)

					}else {
						fmt.Printf("\tthis cell is marked as CROSS V\n")
						up := &k.Board.Cells[c.X-1][c.Y]
						down := &k.Board.Cells[c.X+1][c.Y]
						_, fromItem := k.Set.FindItem(up)
						_, toItem := k.Set.FindItem(c)
						root1 := k.Set.Find(fromItem)
						root2 := k.Set.Find(toItem)
						_ = k.Set.Union(root1, root2)
						_, toItem2 := k.Set.FindItem(down)
						root1b := k.Set.Find(toItem)
						root2b := k.Set.Find(toItem2)
						_ = k.Set.Union(root1b, root2b)

						left := &k.Board.Cells[c.X][c.Y-1]
						right := &k.Board.Cells[c.X][c.Y+1]
						_, leftItem := k.Set.FindItem(left)
						_, rightItem := k.Set.FindItem(right)
						rootLeft := k.Set.Find(leftItem)
						rootRight := k.Set.Find(rightItem)
						_ = k.Set.Union(root2b, rootLeft)
						_ = k.Set.Union(rootLeft, rootRight)
					}
					k.Board.Break2Walls(c, idx)


				}
			}
		}
	}
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (k *KruskalWeave) Generate() error {

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

