package skiplist

import (
	"testing"
)

type iStackNode struct {
	v    int
	next StackNode
}

func (i *iStackNode) Next() StackNode {
	return i.next
}
func (i *iStackNode) SetNext(n StackNode) {
	i.next = n
}

func TestStack(t *testing.T) {
	tcs := map[string]struct {
		input  []int
		expect []int
		empty  bool
		size   int
	}{
		"normal": {
			input:  []int{1, 2, 3, 4, 5},
			expect: []int{5, 4, 3, 2, 1},
			empty:  true,
			size:   0,
		},
	}

	for k, v := range tcs {
		t.Run(k, func(t *testing.T) {
			stack := NewStack()
			for i := 0; i < len(v.input); i++ {
				stack.Push(&iStackNode{v: v.input[i]})
			}
			for i := 0; i < len(v.expect); i++ {
				inode := stack.Pop()
				isnode, ok := inode.(*iStackNode)
				if !ok {
					t.Error("type assert failed")
					return
				}
				if v.expect[i] != isnode.v {
					t.Errorf("Stack.Poo expect : %d , but got : %d", v.expect[i], isnode.v)
					return
				}
			}
			if stack.Size() != v.size {
				t.Errorf("Stack.Size expect : %d , but got : %d", v.size, stack.Size())
				return
			}

			if stack.Empty() != v.empty {
				t.Errorf("Stack.Empty expect : %t , but got : %t", v.empty, stack.Empty())
			}
		})
	}
}
