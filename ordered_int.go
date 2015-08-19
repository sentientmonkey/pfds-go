package pfds

import "fmt"

type orderedInt struct {
	i int
}

// Ordered functor for int
func OrderedInt(i int) Ordered {
	return &orderedInt{i}
}

func OrderedInts(ints ...int) []Ordered {
	r := make([]Ordered, len(ints))
	for i, anInt := range ints {
		r[i] = OrderedInt(anInt)
	}
	return r
}

func (a *orderedInt) Lt(b Ordered) bool {
	return a.i < b.(*orderedInt).i
}

func (a *orderedInt) Gt(b Ordered) bool {
	return a.i > b.(*orderedInt).i
}

func (a *orderedInt) Eq(b Ordered) bool {
	return a.i == b.(*orderedInt).i
}

func (a *orderedInt) String() string {
	return fmt.Sprintf("%d", a.i)
}
