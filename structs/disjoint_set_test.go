package structs

import (
	"testing"
	"os"
)

func TestFind(t *testing.T) {
	ds := &DisjointSet{}
	ds.Items = make([]*Item, 2)
	c1 := &Cell{15, 0, 0}
	c2 := &Cell{15, 0, 1}
	item1 := &Item{c1, nil, nil}
	item2 := &Item{c2, nil, nil}
	ds.Items[0] = item1
	ds.Items[1] = item2
	result := ds.Find(item1)
	if result.Data.X != 0 && result.Data.Y != 0 {
		t.Errorf("Find(), want [0,0], got [%d,%d]", result.Data.X, result.Data.Y)
	}
}

func TestUnion(t *testing.T) {
	ds := &DisjointSet{}
	ds.Items = make([]*Item, 2)
	c1 := &Cell{15, 0, 0}
	c2 := &Cell{15, 0, 1}
	item1 := &Item{c1, nil, nil}
	item2 := &Item{c2, nil, nil}
	ds.Items[0] = item1
	ds.Items[1] = item2
	result := ds.Union(item1, item2)
	if result.Data.X != 0 && result.Data.Y != 1 {
		t.Errorf("Union(), want [0,1], got [%d,%d]", result.Data.X, result.Data.Y)
	}
	root := ds.Find(item1)
	if result.Data.X != 0 && result.Data.Y != 1 {
		t.Errorf("Union(), merged item1 & item2, Find(item1) should return item2 as root, got %v", root.Data)
	}
	ds.Write(os.Stdout)
}
