package structs

import (
	"io"
	"fmt"
)

type DisjointSet struct {
	Items map[string]*Item
}

type Item struct {
	Data *Cell
	Prev *Item
	Next *Item
}

func (i *Item) String() string {
	return fmt.Sprintf("%v", i.Data)
}

func (ds *DisjointSet) FindItem(c *Cell) (bool, *Item) {
	if item, ok := ds.Items[fmt.Sprintf("%d_%d", c.X, c.Y)]; ok {
		return true, item
	}
	return false, nil
}

//Find goes up the tree and find the root
func (ds *DisjointSet) Find(item *Item) *Item {
	if item.Prev == nil {
		return item 
	}
	return ds.Find(item.Prev)
}

func (ds *DisjointSet) FindTail(item *Item) *Item {
	if item.Next == nil {
		return item 
	}
	return ds.FindTail(item.Next)
}

func (ds *DisjointSet) Union(item1, item2 *Item) *Item {
	item1.Next = item2
	item2.Prev = item1
	return item2
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
	if item.Next!= nil {
		writer.Write([]byte(" --> "))
		ds.WriteItem(item.Next, writer)
	} else {
		writer.Write([]byte(" --> nil "))
	}
}
