package pfds

import (
	"fmt"
	"strings"
)

type Set interface {
	Insert(Ordered) Set
	Member(Ordered) bool
	String() string
}

type tree struct {
	value       Ordered
	left, right Set
}

// Builds a new set. If ordered items provided, inserts in order.
func NewSet(ordered ...Ordered) Set {
	if len(ordered) == 0 {
		return &tree{}
	}

	last := len(ordered) - 1
	return NewSet(ordered[:last]...).Insert(ordered[last])
}

func Element(value Ordered) Set {
	return &tree{value, NewSet(), NewSet()}
}

func Tree(value Ordered, left, right Set) Set {
	return &tree{value, left, right}
}

func (t *tree) Insert(x Ordered) Set {
	if t.value == nil {
		return Element(x)
	}
	if x.Lt(t.value) {
		return Tree(t.value, t.left.Insert(x), t.right)
	}
	if x.Gt(t.value) {
		return Tree(t.value, t.left, t.right.Insert(x))
	}

	return t
}

func (t *tree) Member(x Ordered) bool {
	if t.value == nil {
		return false
	}
	if x.Lt(t.value) {
		return t.left.Member(x)
	}
	if x.Gt(t.value) {
		return t.right.Member(x)
	}

	return true
}

func (t *tree) String() string {
	if t.value == nil {
		return ""
	}
	return strings.TrimSpace(fmt.Sprintf("%v %v %v", t.left, t.value, t.right))
}
