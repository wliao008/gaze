package util

import (
	"io"
)

type DisjointSet struct {
	Items []*Item
}

type Item struct {
	From interface{}
	To interface{}
	Parent *Item
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
		writer.Write([]byte("nil\n"))
	}
}

func (ds *DisjointSet) WriteItem(item *Item, writer io.Writer) {
	writer.Write([]byte(item.From.(string) + " --> "))
	if item.Parent != nil {
		ds.WriteItem(item.Parent, writer)
	}
}
