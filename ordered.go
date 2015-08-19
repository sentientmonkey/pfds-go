package pfds

type Ordered interface {
	Gt(Ordered) bool
	Lt(Ordered) bool
	Eq(Ordered) bool
	String() string
}
