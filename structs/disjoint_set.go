package structs

import (
	"io"
	"fmt"
)

type DisjointSet struct {
	Items []*Item
}

type Item struct {
	Data *Cell
	Next *Item
	Prev *Item
}

func (i *Item) String() string {
	return fmt.Sprintf("%v", i.Data)
}

func (ds *DisjointSet) FindItem(c *Cell) (bool, *Item) {
	for _, item := range ds.Items {
		if item.Data.X == c.X && item.Data.Y == c.Y {
			return true, item
		}
	}
	return false, nil
}

//Find goes up the tree and find the root
func (ds *DisjointSet) Find(item *Item) *Item {
	if item.Next == nil {
		return item 
	}
	return ds.Find(item.Next)
}

func (ds *DisjointSet) FindPrev(item *Item) *Item {
	if item.Prev == nil {
		return item 
	}
	return ds.FindPrev(item.Prev)
}

func (ds *DisjointSet) Union(item1, item2 *Item) *Item {
	last := ds.Find(item1)
	root := ds.FindPrev(item2)
	last.Next = root
	root.Prev = last
	return ds.Find(item2)
}

func (ds *DisjointSet) Write(writer io.Writer) {
	for _, item := range ds.Items {
		ds.WriteItem(item, writer)
		writer.Write([]byte("\n"))
	}
}

func (ds *DisjointSet) WriteItem(item *Item, writer io.Writer) {
	str := fmt.Sprintf("%v", item.Data)
	writer.Write([]byte(str))
	if item.Next != nil {
		writer.Write([]byte(" --> "))
		ds.WriteItem(item.Next, writer)
	} else {
		writer.Write([]byte(" --> nil "))
	}
}
