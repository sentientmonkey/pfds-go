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
	Update(int, Item) (List, error)
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

	head, _ := l.Head()
	tail, _ := l.Tail()
	return tail.Append(ys).Cons(head)
}

func (l *list) Update(i int, y Item) (List, error) {
	if l.IsEmpty() {
		return nil, errors.New("i is out of bounds")
	}

	x, _ := l.Head()
	xs, _ := l.Tail()

	if i == 0 {
		return xs.Cons(y), nil
	}
	update, err := xs.Update(i-1, y)
	if err != nil {
		return nil, err
	}

	return update.Cons(x), nil
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
