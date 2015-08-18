package pfds

import "fmt"

type orderedString struct {
	s string
}

// Ordered functor for string
func OrderedString(s string) Ordered {
	return &orderedString{s}
}

func (a *orderedString) lt(b Ordered) bool {
	return a.s < b.(*orderedString).s
}

func (a *orderedString) gt(b Ordered) bool {
	return a.s > b.(*orderedString).s
}

func (a *orderedString) eq(b Ordered) bool {
	return a.s == b.(*orderedString).s
}

func (a *orderedString) String() string {
	return fmt.Sprintf("%s", a.s)
}
