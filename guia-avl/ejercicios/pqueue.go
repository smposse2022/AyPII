package ejercicios_avl

import (
	"github.com/untref-ayp2/data-structures/types"
)

type PQueue[T types.Ordered] struct {
}

func NewPQueue[T types.Ordered]() *PQueue[T] {
	return nil
}

func (q *PQueue[T]) Enqueue(v T) {
}

func (q *PQueue[T]) Dequeue() (T, error) {
	var head T
	return head, nil
}

func (q *PQueue[T]) Front() (T, error) {
	var head T
	return head, nil
}

func (q *PQueue[T]) IsEmpty() bool {
	return false
}
