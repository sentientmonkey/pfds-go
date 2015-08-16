package pfds

import (
	"errors"
	"fmt"
)

type Item interface{}

type List interface {
	IsEmpty() bool
	Cons(Item) List
	Head() (Item, error)
	Tail() (List, error)
	Append(List) List
	String() string
}

type list struct {
	value Item
	next  *list
}

func NewList() List {
	return &list{}
}

func (l *list) IsEmpty() bool {
	return l.value == nil && l.next == nil
}

func (l *list) Cons(item Item) List {
	return &list{item, l}
}

func (l *list) Head() (Item, error) {
	if l.IsEmpty() {
		return nil, errors.New("List is empty")
	}

	return l.value, nil
}

func (l *list) Tail() (List, error) {
	if l.IsEmpty() {
		return nil, errors.New("List is empty")
	}

	return l.next, nil
}

func (l *list) Append(ys List) List {
	if l.IsEmpty() {
		return ys
	}

	var curr List = l
	next, _ := l.Tail()
	for !next.IsEmpty() {
		curr = next
		next, _ = next.Tail()
	}

	curr.(*list).next = ys.(*list)

	return l
}

func (l *list) String() string {
	r := "["
	curr := l
	for !curr.IsEmpty() {
		head, _ := curr.Head()
		r += fmt.Sprintf("%v", head)
		next, _ := curr.Tail()
		curr = next.(*list)
		if !curr.IsEmpty() {
			r += " "
		}
	}

	r += "]"
	return r
}
