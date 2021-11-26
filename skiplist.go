package skiplist

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// ---------------------------------------------------------

// ---------------------------------------------------------

// Node skiplist node
type Node interface {
	Next() Node
	Down() Node
	SetNext(Node)
	SetDown(Node)

	Key() int64
	CopyValue() Node
}

// SkipList skiplist
type SkipList struct {
	Header       Node
	MaxLevel     int8
	CurrentLevel int8
	r            *rand.Rand
}

func NewSkipList(MaxLeave int8) *SkipList {
	return &SkipList{
		MaxLevel:     MaxLeave,
		CurrentLevel: 0,
		r:            rand.New(rand.NewSource(time.Now().Unix())),
	}
}

// Search search node
func (s *SkipList) Search(v int64) Node {
	n := s.Header
	for n != nil {
		if n.Key() == v {
			return n
		}
		if n.Next() == nil {
			n = n.Down()
			continue
		}
		if n.Next().Key() > v {
			n = n.Down()
			continue
		}
		n = n.Next()
	}
	return nil
}

// Delete remove node from skiplist
func (s *SkipList) Delete(v int64) Node {
	n := s.Header
	var result Node
	for n != nil {
		if n.Next() == nil {
			n = n.Down()
			continue
		}
		if n.Next().Key() == v {
			result = n.Next()
			n.SetNext(n.Next().Next())
			n = n.Down()
			continue
		}
		if n.Next().Key() > v {
			n = n.Down()
			continue
		}
		n = n.Next()
	}
	return result
}

// Insert insert node into skiplist
func (s *SkipList) Insert(v Node) {
	k := v.Key()
	exists := s.Search(k)
	if exists != nil {
		return
	}

	if s.Header == nil {
		s.Header = v
		s.CurrentLevel = 1
		return
	}

	// process v less then header
	if v.Key() < s.Header.Key() {
		var stack = NewStack()
		inode := s.Header
		for inode != nil {
			stack.Push(&SLNStackNode{value: inode})
			inode = inode.Down()
		}
		var private Node
		for !stack.Empty() {
			n, _ := stack.Pop().(*SLNStackNode)
			newnode := v.CopyValue()
			newnode.SetNext(n.value)
			newnode.SetDown(private)
			s.Header = newnode
			private = newnode
		}
		if s.CurrentLevel >= s.MaxLevel {
			return
		}

		if s.r.Intn(100)%2 == 0 {
			newnode := v.CopyValue()
			newnode.SetDown(private)
			s.CurrentLevel += 1
			s.Header = newnode
		}
		return
	}

	// process v more then header
	stack := NewStack()
	// find insert node
	inode := s.Header
	for inode != nil {
		if inode.Next() == nil {
			stack.Push(&SLNStackNode{value: inode})
			inode = inode.Down()
			continue
		}
		if inode.Next().Key() > k {
			stack.Push(&SLNStackNode{value: inode})
			inode = inode.Down()
			continue
		}
		inode = inode.Next()
	}

	var downNode Node
	var level int8 = 1

	for !stack.Empty() {
		pre, _ := stack.Pop().(*SLNStackNode)

		newNode := v.CopyValue()
		newNode.SetDown(downNode)
		downNode = newNode

		if pre.value.Next() != nil {
			newNode.SetNext(pre.value.Next())
		}
		pre.value.SetNext(newNode)

		// break if leave more the MaxLeave
		if level >= s.MaxLevel {
			break
		}
		// do random grow leave
		if s.r.Intn(100)%2 != 0 {
			break
		}
		level++
		if level > s.CurrentLevel {
			s.CurrentLevel = level
			nhead := s.Header.CopyValue()
			nhead.SetDown(s.Header)
			s.Header = nhead
			stack.Push(&SLNStackNode{value: nhead})
		}
	}
}
func (s SkipList) String() string {
	buf := new(bytes.Buffer)
	n := s.Header
	for n != nil {
		v := n
		l1 := new(bytes.Buffer)
		l2 := new(bytes.Buffer)
		l3 := new(bytes.Buffer)
		for v != nil {
			k := fmt.Sprintf("%d", v.Key())
			nx := ""
			if v.Next() != nil {
				nx = fmt.Sprintf("%d", v.Next().Key())
			}
			dv := ""
			if v.Down() != nil {
				dv = fmt.Sprintf("%d", v.Down().Key())
			}

			l1.WriteString(fmt.Sprintf("| %s - %s |", k, nx))

			l2.WriteString("| ")
			for i := 0; i < (len(k)-1)/2; i++ {
				l2.WriteString(" ")
			}
			if len(k)%2 == 0 {
				l2.WriteString(" ")
			}
			l2.WriteString("^")
			for i := 0; i < (len(k)-1)/2; i++ {
				l2.WriteString(" ")
			}
			l2.WriteString("   ")
			for i := 0; i < len(nx); i++ {
				l2.WriteString(" ")
			}
			l2.WriteString(" |")

			l3.WriteString(fmt.Sprintf("| %s", dv))
			if dv == "" {
				for i := 0; i < len(k); i++ {
					l3.WriteString(" ")
				}
			}
			l3.WriteString("   ")
			for i := 0; i < len(nx); i++ {
				l3.WriteString(" ")
			}
			l3.WriteString(" |")
			v = v.Next()
		}
		buf.WriteString(l1.String())
		buf.WriteString("\n")
		buf.WriteString(l2.String())
		buf.WriteString("\n")
		buf.WriteString(l3.String())
		buf.WriteString("\n")
		buf.WriteString("\n")
		n = n.Down()
	}
	return buf.String()
}
