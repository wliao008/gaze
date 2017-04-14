package main

import (
	"fmt"
	"github.com/wliao008/mazing/algos"
	"github.com/wliao008/mazing/structs"
)

func main() {
	c := structs.Cell{}
	c.SetBit(structs.EAST)
	fmt.Println("%+v", c)
	c.ClearBit(structs.EAST)
	fmt.Println("%+v", c)

	bt := algos.BackTracking{50, 10}
	bt.Generate()
}
