package ejercicios_avl

import (
	"guiaAvl/avl"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVLPosOrderIteratorVacio(t *testing.T) {
	avl := avl.NewAVLTree[int]()
	iterator := NewAVLPosOrderIterator[int](avl.GetRoot())

	assert.False(t, iterator.HasNext())
}

func TestAVLPosOrderIterator(t *testing.T) {
	avl := avl.NewAVLTree[int]()

	avl.Insert(4)
	avl.Insert(2)
	avl.Insert(6)
	avl.Insert(1)
	avl.Insert(3)
	avl.Insert(5)
	avl.Insert(7)

	iterator := NewAVLPosOrderIterator[int](avl.GetRoot())

	expected := []int{1, 3, 2, 5, 7, 6, 4}

	for _, expVal := range expected {
		assert.True(t, iterator.HasNext())
		val, _ := iterator.Next()
		assert.Equal(t, expVal, val)
	}
	assert.False(t, iterator.HasNext())
}

func TestAVLPosOrderIteratorNextOverflow(t *testing.T) {
	avl := avl.NewAVLTree[int]()

	avl.Insert(1)
	avl.Insert(2)
	avl.Insert(3)

	iterator := NewAVLPosOrderIterator[int](avl.GetRoot())

	_, _ = iterator.Next()
	_, _ = iterator.Next()
	_, _ = iterator.Next()

	assert.False(t, iterator.HasNext())
	_, err := iterator.Next()
	assert.Error(t, err)
}
