package ejercicios_avl

import (
	"guiaAvl/avl"

	"github.com/untref-ayp2/data-structures/types"
)

type AVLPreOrderIterator[T types.Ordered] struct {
}

func NewAVLPreOrderIterator[T types.Ordered](root *avl.AVLNode[T]) *AVLPreOrderIterator[T] {
	return nil
}

func (it *AVLPreOrderIterator[T]) HasNext() bool {
	return false
}

func (it *AVLPreOrderIterator[T]) Next() (T, error) {
	var data T
	return data, nil
}
