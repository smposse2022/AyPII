package ejercicios_avl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTreeSet(t *testing.T) {
	set := NewTreeSet[int]()

	assert.NotNil(t, set)
	assert.Equal(t, 0, set.Size())
}

func TestTreeSetAdd(t *testing.T) {
	set := NewTreeSet[int]()

	set.Add(1)
	assert.Equal(t, 1, set.Size())
}

func TestTreeSetAddMultiple(t *testing.T) {
	set := NewTreeSet[int]()

	set.Add(1, 2, 3)
	assert.Equal(t, 3, set.Size())
}

func TestTreeSetAddExistenteNoRepite(t *testing.T) {
	set := NewTreeSet[int]()

	set.Add(1)
	set.Add(1)
	assert.Equal(t, 1, set.Size())
}

func TestTreeSetContains(t *testing.T) {
	set := NewTreeSet[int]()
	set.Add(1)

	assert.True(t, set.Contains(1))
	assert.False(t, set.Contains(2))
}

func TestTreeSetRemove(t *testing.T) {
	set := NewTreeSet[int]()
	set.Add(1)
	set.Add(2)
	assert.True(t, set.Contains(1))
	assert.Equal(t, 2, set.Size())

	set.Remove(1)
	assert.False(t, set.Contains(1))
	assert.Equal(t, 1, set.Size())
}

func TestTreeSetRemoveNonExistent(t *testing.T) {
	set := NewTreeSet[int]()
	set.Add(1)
	assert.Equal(t, 1, set.Size())

	set.Remove(2)
	assert.Equal(t, 1, set.Size())
}

func TestTreeSetSize(t *testing.T) {
	set := NewTreeSet[int]()
	assert.Equal(t, 0, set.Size())

	set.Add(1)
	assert.Equal(t, 1, set.Size())

	set.Add(2)
	assert.Equal(t, 2, set.Size())
}

func TestTreeSetaluesOnAnEmptySet(t *testing.T) {
	set := NewTreeSet[int]()
	values := set.Values()

	assert.Empty(t, values)
}

func TestTreeSetValuesOnANonEmptySet(t *testing.T) {
	set := NewTreeSet(1, 2)
	values := set.Values()
	assert.Len(t, values, 2)
	assert.ElementsMatch(t, []int{1, 2}, values)
}

func TestTreeSetStringEnSetVacio(t *testing.T) {
	set := NewTreeSet[int]()
	assert.Equal(t, "Set: {}", set.String())
}

func TestTreeSetStringEnSetNoVacio(t *testing.T) {
	set := NewTreeSet(1, 7)
	assert.Contains(t, "Set: {1, 7}", set.String())
}
