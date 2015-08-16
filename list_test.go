package pfds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	list := NewList()
	assert.True(t, list.IsEmpty())

	newList := list.Cons(1)
	assert.True(t, list.IsEmpty())
	assert.False(t, newList.IsEmpty())
	assert.Equal(t, "[]", list.String())

	head, err := newList.Head()
	assert.Equal(t, 1, head)
	assert.Nil(t, err)
	tail, err := newList.Tail()
	assert.Nil(t, err)
	assert.NotNil(t, tail)
	assert.True(t, tail.IsEmpty())
	assert.Equal(t, "[1]", newList.String())

	anotherList := newList.Cons(2)
	anotherHead, err := anotherList.Head()
	assert.Equal(t, 2, anotherHead)
	anotherTail, err := anotherList.Tail()
	assert.Nil(t, err)
	assert.Equal(t, newList, anotherTail)

	assert.Equal(t, "[2 1]", anotherList.String())
}

func TestListUpdate(t *testing.T) {
	xs := NewList().Cons(2).Cons(1).Cons(0)
	ys := NewList().Cons(5).Cons(4).Cons(3)

	assert.Equal(t, "[0 1 2]", xs.String())
	assert.Equal(t, "[3 4 5]", ys.String())
	zs := xs.Append(ys)
	assert.Equal(t, "[0 1 2 3 4 5]", zs.String())
	assert.Equal(t, "[0 1 2]", xs.String())
	assert.Equal(t, "[3 4 5]", ys.String())
}