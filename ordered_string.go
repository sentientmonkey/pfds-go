package pfds

import "fmt"

type orderedString struct {
	s string
}

// Ordered functor for string
func OrderedString(s string) Ordered {
	return &orderedString{s}
}

func OrderedStrings(strings ...string) []Ordered {
	r := make([]Ordered, len(strings))
	for i, aString := range strings {
		r[i] = OrderedString(aString)
	}
	return r
}

func (a *orderedString) Lt(b Ordered) bool {
	return a.s < b.(*orderedString).s
}

func (a *orderedString) Gt(b Ordered) bool {
	return a.s > b.(*orderedString).s
}

func (a *orderedString) Eq(b Ordered) bool {
	return a.s == b.(*orderedString).s
}

func (a *orderedString) String() string {
	return fmt.Sprintf("%s", a.s)
}
