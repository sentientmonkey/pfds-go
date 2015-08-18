package pfds

import "fmt"

type orderedInt struct {
	i int
}

// Ordered functor for int
func OrderedInt(i int) Ordered {
	return &orderedInt{i}
}

func (a *orderedInt) lt(b Ordered) bool {
	return a.i < b.(*orderedInt).i
}

func (a *orderedInt) gt(b Ordered) bool {
	return a.i > b.(*orderedInt).i
}

func (a *orderedInt) eq(b Ordered) bool {
	return a.i == b.(*orderedInt).i
}

func (a *orderedInt) String() string {
	return fmt.Sprintf("%d", a.i)
}
