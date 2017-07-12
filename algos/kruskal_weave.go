package algos

import (
	"github.com/wliao008/mazing/structs"
	"math/rand"
	"time"
	"fmt"
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
	k.preprocess()
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
					}else {
						fmt.Printf("\tthis cell is marked as CROSS V\n")
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

