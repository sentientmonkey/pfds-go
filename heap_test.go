package pfds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHeap(t *testing.T) {
	h := NewHeap()
	assert.True(t, h.IsEmpty())

	min, err := h.FindMin()
	assert.Nil(t, min)
	assert.Error(t, err)
}

func TestHeapInsert(t *testing.T) {
	h := NewHeap().Insert(OrderedString("x"))
	assert.False(t, h.IsEmpty())

	min, err := h.FindMin()
	assert.NoError(t, err)
	assert.Equal(t, OrderedString("x"), min)
}

func TestHeapInserts(t *testing.T) {
	h := NewHeap().Insert(OrderedString("x")).Insert(OrderedString("y"))
	min, err := h.FindMin()
	assert.NoError(t, err)
	assert.Equal(t, OrderedString("x"), min)
}

func TestHeapDeleteMin(t *testing.T) {
	h := NewHeap().Insert(OrderedString("x")).Insert(OrderedString("y"))

	min, err := h.FindMin()
	assert.NoError(t, err)
	assert.Equal(t, OrderedString("x"), min)

	newHeap, err := h.DeleteMin()
	assert.NoError(t, err)

	newMin, err := newHeap.FindMin()
	assert.NoError(t, err)
	assert.Equal(t, OrderedString("y"), newMin)
}

func TestHeapDeleteRight(t *testing.T) {
	h := NewHeap().Insert(OrderedString("y")).Insert(OrderedString("x")).Insert(OrderedString("a"))

	newHeap, err := h.DeleteMin()
	assert.NoError(t, err)

	newMin, err := newHeap.FindMin()
	assert.NoError(t, err)
	assert.Equal(t, OrderedString("x"), newMin)
}

func TestHeapDeleteEmpty(t *testing.T) {
	h := NewHeap()

	_, err := h.DeleteMin()
	assert.Error(t, err)
}
