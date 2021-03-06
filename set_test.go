package pfds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	set := NewSet(OrderedInts(1, 2, 3)...)
	assert.NotNil(t, set)
	assert.True(t, set.Member(OrderedInt(1)))
	assert.True(t, set.Member(OrderedInt(2)))
	assert.True(t, set.Member(OrderedInt(3)))
	assert.False(t, set.Member(OrderedInt(4)))
	assert.Equal(t, "1 2 3", set.String())

	anotherSet := set.Insert(OrderedInt(10))
	assert.True(t, anotherSet.Member(OrderedInt(10)))
	assert.False(t, set.Member(OrderedInt(10)))

	sameSet := anotherSet.Insert(OrderedInt(10))
	assert.Equal(t, anotherSet, sameSet)

	unbalancedSet := anotherSet.Insert(OrderedInt(5))
	assert.Equal(t, "1 2 3 5 10", unbalancedSet.String())
	assert.True(t, unbalancedSet.Member(OrderedInt(5)))
}

func TestSetStringTest(t *testing.T) {
	set := NewSet(OrderedStrings("a", "b", "c")...)
	assert.NotNil(t, set)
	assert.Equal(t, "a b c", set.String())
}
