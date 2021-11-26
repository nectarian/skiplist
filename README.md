# skiplist
skiplist implement by Go


# USAGE


```go
package main

import (
	"fmt"

	"github.com/nectarian/skiplist"
)

// ----------------------------------------------
// define a struct , implement skiplist.Node

type iNode struct {
	v    int64
	next skiplist.Node
	down skiplist.Node
}

func (n *iNode) Next() skiplist.Node {
	return n.next
}
func (n *iNode) Down() skiplist.Node {
	return n.down
}
func (n *iNode) SetNext(nx skiplist.Node) {
	n.next = nx
}
func (n *iNode) SetDown(d skiplist.Node) {
	n.down = d
}
func (n *iNode) Key() int64 {
	return n.v
}
func (n *iNode) CopyValue() skiplist.Node {
	return &iNode{
		v: n.v,
	}
}

// ----------------------------------------------
// do what you want to do

func main() {
	sl := skiplist.NewSkipList(5)
	// insert
	for i := 0; i < 100; i++ {
		sl.Insert(&iNode{v: int64(i * 2)})
	}
	// delete
	for i := 0; i < 15; i++ {
		sl.Delete(int64(i * 2))
	}
	// search
	fmt.Println(sl.Search(98).Key())
	fmt.Println(sl)
}
```