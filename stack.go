package skiplist

// ---------------------------------------------

type StackNode interface {
	Next() StackNode
	SetNext(StackNode)
}

type Stack struct {
	Header StackNode
	size   int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(v StackNode) {
	v.SetNext(s.Header)
	s.Header = v
	s.size += 1
}

func (s *Stack) Pop() StackNode {
	r := s.Header
	if r != nil {
		s.Header = s.Header.Next()
	}
	s.size -= 1
	return r
}

func (s *Stack) Empty() bool {
	return s.Header == nil
}

func (s *Stack) Size() int {
	return s.size
}

// ---------------------------------------------

type SLNStackNode struct {
	value Node
	next  StackNode
}

func (s *SLNStackNode) Next() StackNode {
	return s.next
}
func (s *SLNStackNode) SetNext(v StackNode) {
	s.next = v
}

// ---------------------------------------------
