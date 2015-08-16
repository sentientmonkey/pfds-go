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

func TestListAppend(t *testing.T) {
	xs := NewList(0, 1, 2)
	ys := NewList(3, 4, 5)

	assert.Equal(t, "[0 1 2]", xs.String())
	assert.Equal(t, "[3 4 5]", ys.String())
	zs := xs.Append(ys)
	assert.Equal(t, "[0 1 2 3 4 5]", zs.String())
	assert.Equal(t, "[0 1 2]", xs.String())
	assert.Equal(t, "[3 4 5]", ys.String())
}

func TestListUpdate(t *testing.T) {
	xs := NewList(0, 1, 2, 3, 4)

	assert.Equal(t, "[0 1 2 3 4]", xs.String())

	ys, err := xs.Update(2, 7)
	assert.Nil(t, err)
	assert.Equal(t, "[0 1 7 3 4]", ys.String())
	assert.Equal(t, "[0 1 2 3 4]", xs.String())

	_, err = xs.Update(5, 7)
	assert.NotNil(t, err)
	assert.Equal(t, "[0 1 2 3 4]", xs.String())
}

func TestListFromArray(t *testing.T) {
	xs := NewList(0, 1, 2, 3, 4)
	assert.Equal(t, "[0 1 2 3 4]", xs.String())

	xs = NewList(0)
	assert.Equal(t, "[0]", xs.String())

	xs = NewList()
	assert.Equal(t, "[]", xs.String())
}
