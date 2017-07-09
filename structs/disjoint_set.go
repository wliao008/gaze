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
	Parent *Item
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
	if item.Parent == nil {
		return item 
	}
	return ds.Find(item.Parent)
}

func (ds *DisjointSet) Union(item1, item2 *Item) *Item {
	root1 := ds.Find(item1)
	root2 := ds.Find(item2)
	root1.Parent = root2
	return root2
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
	if item.Parent != nil {
		writer.Write([]byte(" --> "))
		ds.WriteItem(item.Parent, writer)
	} else {
		writer.Write([]byte(" --> nil "))
	}
}
