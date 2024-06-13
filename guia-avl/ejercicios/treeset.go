package ejercicios_avl

import (
	"github.com/untref-ayp2/data-structures/types"
)

type TreeSet[T types.Ordered] struct {
}

func NewTreeSet[T types.Ordered](elements ...T) *TreeSet[T] {
	return nil
}

func (s *TreeSet[T]) Contains(element T) bool {
	return false
}

func (s *TreeSet[T]) Add(elements ...T) {
}

func (s *TreeSet[T]) Remove(element T) {
}

func (s *TreeSet[T]) Size() int {
	return 0
}

func (s *TreeSet[T]) Values() []T {
	return nil
}

func (s *TreeSet[T]) String() string {
	return ""
}
