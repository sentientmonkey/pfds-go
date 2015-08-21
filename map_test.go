package pfds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	m1 := NewMap()
	r1, err := m1.Lookup(OrderedString("i"))
	assert.Nil(t, r1)
	assert.Error(t, err)

	m2 := m1.Bind(OrderedString("i"), 1)
	r2, err := m1.Lookup(OrderedString("i"))
	assert.Nil(t, r2)
	assert.Error(t, err)
	r3, err := m2.Lookup(OrderedString("i"))
	assert.Equal(t, 1, r3)
	assert.NoError(t, err)
	assert.Equal(t, "i:1", m2.String())

	m3 := m2.Bind(OrderedString("i"), 2)
	r4, err := m2.Lookup(OrderedString("i"))
	assert.Equal(t, 1, r4)
	assert.NoError(t, err)
	r5, err := m3.Lookup(OrderedString("i"))
	assert.Equal(t, 2, r5)
	assert.NoError(t, err)
	assert.Equal(t, "i:2", m3.String())

	m4 := m3.Bind(OrderedString("k"), 5)
	assert.Equal(t, "i:2 k:5", m4.String())
	assert.Equal(t, "i:2", m3.String())

	r6, err := m4.Lookup(OrderedString("k"))
	assert.Equal(t, 5, r6)
	assert.NoError(t, err)

	r8, err := m4.Lookup(OrderedString("a"))
	assert.Nil(t, r8)
	assert.Error(t, err)

	m5 := m4.Bind(OrderedString("a"), 3)
	assert.Equal(t, "a:3 i:2 k:5", m5.String())
}
