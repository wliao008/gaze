package gaze

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {
	ds := &DisjointSet{}
	ds.Items = make(map[string]*Item)
	c1 := &Cell{15, 0, 0, nil, nil, nil, 0.0, 0.0}
	c2 := &Cell{15, 0, 1, nil, nil, nil, 0.0, 0.0}
	item1 := &Item{c1, nil}
	item2 := &Item{c2, nil}
	ds.Items["00"] = item1
	ds.Items["01"] = item2
	result := ds.Find(item1)
	if result.Data.X != 0 && result.Data.Y != 0 {
		t.Errorf("Find(), want [0,0], got [%d,%d]", result.Data.X, result.Data.Y)
	}
}

func TestUnion(t *testing.T) {
	ds := &DisjointSet{}
	ds.Items = make(map[string]*Item)
	c1 := &Cell{15, 0, 0, nil, nil, nil, 0.0, 0.0}
	c2 := &Cell{15, 0, 1, nil, nil, nil, 0.0, 0.0}
	item1 := &Item{c1, nil}
	item2 := &Item{c2, nil}
	ds.Items["00"] = item1
	ds.Items["01"] = item2
	result := ds.Union(item1, item2)
	if result.Data.X != 0 && result.Data.Y != 1 {
		t.Errorf("Union(), want [0,1], got [%d,%d]", result.Data.X, result.Data.Y)
	}
	root := ds.Find(item1)
	if result.Data.X != 0 && result.Data.Y != 1 {
		t.Errorf("Union(), merged item1 & item2, Find(item1) should return item2 as root, got %v", root.Data)
	}
}

func BenchmarkDisjointSetFind(b *testing.B) {
	size := uint16(1000)
	ds := &DisjointSet{}
	ds.Items = make(map[string]*Item)
	for i := uint16(0); i < size; i++ {
		cell := &Cell{15, 0, i, nil, nil, nil, 0.0, 0.0}
		item := &Item{cell, nil}
		ds.Items[fmt.Sprintf("%d", i)] = item
	}
	for i := uint16(0); i < size-1; i++ {
		ds.Items[fmt.Sprintf("%d", i)].Parent = ds.Items[fmt.Sprintf("%d", i+1)]
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = ds.Find(ds.Items["0"])
	}
}

func BenchmarkDisjointSetUnion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		size := uint16(1000)
		ds := &DisjointSet{}
		ds.Items = make(map[string]*Item)
		for i := uint16(0); i < size; i++ {
			cell := &Cell{15, 0, i, nil, nil, nil, 0.0, 0.0}
			item := &Item{cell, nil}
			ds.Items[fmt.Sprintf("%d", i)] = item
		}
		for i := uint16(0); i < size/2; i++ {
			ds.Items[fmt.Sprintf("%d", i)].Parent = ds.Items[fmt.Sprintf("%d", i+1)]
		}
		root1 := ds.Find(ds.Items["0"])
		root2 := ds.Find(ds.Items["501"])
		_ = ds.Union(root1, root2)
	}
}
