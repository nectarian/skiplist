package skiplist

import (
	"testing"
)

// -------------------------------------------------------

type iNode struct {
	v    int64
	next Node
	down Node
}

func (n *iNode) Next() Node {
	return n.next
}
func (n *iNode) Down() Node {
	return n.down
}
func (n *iNode) SetNext(nx Node) {
	n.next = nx
}
func (n *iNode) SetDown(d Node) {
	n.down = d
}
func (n *iNode) Key() int64 {
	return n.v
}
func (n *iNode) CopyValue() Node {
	return &iNode{
		v: n.v,
	}
}

// -------------------------------------------------------

func TestSkipList(t *testing.T) {
	var tcs = map[string]struct {
		input         []int64
		delete        []int64
		delete_expect []int64
		search        []int64
		search_expect []int64
	}{
		"normal": {
			input:         []int64{1, 2, 3, 4, 5, 5, 6, 7},
			delete:        []int64{4, 5, 6, 7},
			delete_expect: []int64{4, 5, 6, 7},
			search:        []int64{1, 2, 3},
			search_expect: []int64{1, 2, 3},
		},
		"more": {
			input:         []int64{12123, 9487, 53, 21, 12, 3, 34, 52, 23, 234, 12, 3123},
			delete:        []int64{4, 5, 6, 7, 21},
			delete_expect: []int64{0, 0, 0, 0, 21},
			search:        []int64{1, 2, 3},
			search_expect: []int64{0, 0, 3},
		},
	}

	for k, v := range tcs {
		t.Run(k, func(t *testing.T) {
			sl := NewSkipList(5)
			for i := 0; i < len(v.input); i++ {
				sl.Insert(&iNode{v: v.input[i]})
			}
			for i := 0; i < len(v.delete); i++ {
				deleted := sl.Delete(v.delete[i])
				if v.delete_expect[i] == 0 && deleted != nil {
					t.Errorf("SkipList.Delete expect : nil , but got : %d", deleted.Key())
					return
				}
				if v.delete_expect[i] != 0 {
					if deleted == nil {
						t.Errorf("SkipList.Delete expect : %d , but got : nil", v.delete_expect[i])
						return
					}
					if deleted.Key() != v.delete_expect[i] {
						t.Errorf("SkipList.Delete expect : %d , but got : %d", v.delete_expect[i], deleted.Key())
						return
					}
				}
			}

			for i := 0; i < len(v.search); i++ {
				sresult := sl.Search(v.search[i])
				if v.search_expect[i] == 0 && sresult != nil {
					t.Errorf("SkipList.Search expect : nil , but got : %d", sresult.Key())
				}
				if v.search_expect[i] != 0 {
					if sresult == nil {
						t.Errorf("SkipList.Search expect : %d , but got : nil", v.search_expect[i])
						return
					}
					if sresult.Key() != v.search_expect[i] {
						t.Errorf("SkipList.Search expect : %d , but got : %d", v.search_expect[i], sresult.Key())
						return
					}
				}
			}
		})
	}

}
