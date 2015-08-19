package pfds

import (
	"fmt"
	"strings"
)

type Map interface {
	Bind(Ordered, Item) Map
	Lookup(Ordered) Item
	String() string
}

type orderedMap struct {
	key         Ordered
	value       Item
	left, right Map
}

func NewMap() Map {
	return &orderedMap{}
}

func (m *orderedMap) Bind(k Ordered, v Item) Map {
	if m.key == nil {
		return &orderedMap{k, v, NewMap(), NewMap()}
	}
	if k.Lt(m.key) {
		return &orderedMap{m.key, m.value, m.left.Bind(k, v), m.right}
	}
	if k.Gt(m.key) {
		return &orderedMap{m.key, m.value, m.left, m.right.Bind(k, v)}
	}
	return &orderedMap{m.key, v, m.left, m.right}
}

func (m *orderedMap) Lookup(k Ordered) Item {
	if m.value == nil {
		return nil
	}
	if k.Lt(m.key) {
		return m.left.Lookup(k)
	}
	if k.Gt(m.key) {
		return m.right.Lookup(k)
	}

	return m.value
}

func (m *orderedMap) String() string {
	if m.value == nil {
		return ""
	}
	return strings.TrimSpace(fmt.Sprintf("%v %v:%v %v", m.left, m.key, m.value, m.right))
}
