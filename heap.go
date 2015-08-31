package pfds

import "errors"

type Heap interface {
	IsEmpty() bool
	Insert(Ordered) Heap
	Merge(Heap) Heap
	FindMin() (Ordered, error)
	DeleteMin() (Heap, error)
	makeT(Ordered, Heap) Heap
}

type leftistHeap struct {
	rank        int
	value       Ordered
	left, right Heap
}

func NewHeap() Heap {
	return &leftistHeap{}
}

func (h *leftistHeap) IsEmpty() bool {
	return h.value == nil
}

func (h *leftistHeap) Insert(x Ordered) Heap {
	return h.Merge(&leftistHeap{1, x, nil, nil})
}

func (a *leftistHeap) makeT(x Ordered, h Heap) Heap {
	if h == nil {
		return &leftistHeap{1, x, a, nil}
	}

	b := h.(*leftistHeap)

	if a.rank >= b.rank {
		return &leftistHeap{b.rank + 1, x, a, b}
	}

	return &leftistHeap{a.rank + 1, x, b, a}
}

func (h1 *leftistHeap) Merge(h Heap) Heap {
	if h == nil {
		return h1
	}

	h2 := h.(*leftistHeap)

	if h1.IsEmpty() {
		return h2
	}

	if h2.IsEmpty() {
		return h1
	}

	if h1.value.Lt(h2.value) || h2.value.Eq(h2.value) {
		return h2.Merge(h1.right).makeT(h1.value, h1.left)
	}

	return h1.Merge(h2.right).makeT(h2.value, h2.left)
}

func (h *leftistHeap) FindMin() (Ordered, error) {
	if h.IsEmpty() {
		return nil, errors.New("Empty")
	}

	return h.value, nil
}

func (h *leftistHeap) DeleteMin() (Heap, error) {
	if h.IsEmpty() {
		return nil, errors.New("Empty")
	}

	return h.left.Merge(h.right), nil
}
