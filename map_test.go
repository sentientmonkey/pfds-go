package pfds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	m1 := NewMap()
	assert.Nil(t, m1.Lookup(OrderedString("a")))

	m2 := m1.Bind(OrderedString("a"), 1)
	assert.Nil(t, m1.Lookup(OrderedString("a")))
	assert.Equal(t, 1, m2.Lookup(OrderedString("a")))
	assert.Equal(t, "a:1", m2.String())

	m3 := m2.Bind(OrderedString("a"), 2)
	assert.Equal(t, 1, m2.Lookup(OrderedString("a")))
	assert.Equal(t, 2, m3.Lookup(OrderedString("a")))
	assert.Equal(t, "a:2", m3.String())

	m4 := m3.Bind(OrderedString("b"), 5)
	assert.Equal(t, "a:2 b:5", m4.String())
	assert.Equal(t, "a:2", m3.String())
}
