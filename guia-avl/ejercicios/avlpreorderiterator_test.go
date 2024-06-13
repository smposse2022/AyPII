package ejercicios_avl

import (
	"guiaAvl/avl"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVLPreOrderIteratorVacio(t *testing.T) {
	avl := avl.NewAVLTree[int]()
	iterator := NewAVLPreOrderIterator[int](avl.GetRoot())

	assert.False(t, iterator.HasNext())
}

func TestAVLPreOrderIterator(t *testing.T) {
	avl := avl.NewAVLTree[int]()

	avl.Insert(4)
	avl.Insert(2)
	avl.Insert(6)
	avl.Insert(1)
	avl.Insert(3)
	avl.Insert(5)
	avl.Insert(7)

	iterator := NewAVLPreOrderIterator[int](avl.GetRoot())

	expected := []int{4, 2, 1, 3, 6, 5, 7}

	for _, expVal := range expected {
		assert.True(t, iterator.HasNext())
		val, _ := iterator.Next()
		assert.Equal(t, expVal, val)
	}
	assert.False(t, iterator.HasNext())
}

func TestAVLPreOrderIteratorNextOverflow(t *testing.T) {
	avl := avl.NewAVLTree[int]()

	avl.Insert(1)
	avl.Insert(2)
	avl.Insert(3)

	iterator := NewAVLPreOrderIterator[int](avl.GetRoot())

	_, _ = iterator.Next()
	_, _ = iterator.Next()
	_, _ = iterator.Next()

	assert.False(t, iterator.HasNext())
	_, err := iterator.Next()
	assert.Error(t, err)
}
