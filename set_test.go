package pfds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	set := NewSet(OrderedInt(1), OrderedInt(2), OrderedInt(3))
	assert.NotNil(t, set)
	assert.True(t, set.Member(OrderedInt(1)))
	assert.True(t, set.Member(OrderedInt(2)))
	assert.True(t, set.Member(OrderedInt(3)))
	assert.False(t, set.Member(OrderedInt(4)))
	assert.Equal(t, "( 1 ( 2 ( 3 )))", set.String())

	anotherSet := set.Insert(OrderedInt(10))
	assert.True(t, anotherSet.Member(OrderedInt(10)))
	assert.False(t, set.Member(OrderedInt(10)))

	assert.Equal(t, "( 1 ( 2 ( 3 (( 5 ) 10 ))))", anotherSet.Insert(OrderedInt(5)).String())
}

func TestSetStringTest(t *testing.T) {
	set := NewSet(OrderedString("a"), OrderedString("b"), OrderedString("c"))
	assert.NotNil(t, set)
	assert.Equal(t, "( a ( b ( c )))", set.String())
}
