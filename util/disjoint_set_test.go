package util

import (
	"testing"
	"fmt"
	"os"
)

func TestFind(t *testing.T) {
	ds := &DisjointSet{}
	ds.Items = make([]*Item, 2)
	item1 := &Item{"1", nil}
	item2 := &Item{"2", nil}
	ds.Items[0] = item1
	ds.Items[1] = item2
	result := ds.Find(item1)
	fmt.Println(result)
	if result.Data != "1" {
		t.Errorf("Find(), want 1, got %d", result.Data)
	}
}

func TestUnion(t *testing.T) {
	ds := &DisjointSet{}
	ds.Items = make([]*Item, 2)
	item1 := &Item{"1", nil}
	item2 := &Item{"2", nil}
	ds.Items[0] = item1
	ds.Items[1] = item2
	result := ds.Union(item1, item2)
	if result.Data != "2" {
		t.Errorf("Union(), want 2, got %d", result.Data)
	}
	root := ds.Find(item1)
	if root.Data != "2" {
		t.Errorf("Union(), merged item1 & item2, Find(item1) should return item2 as root, got %d", root.Data)
	}
	ds.Write(os.Stdout)
}
