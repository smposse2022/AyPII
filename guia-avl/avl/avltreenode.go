package avl

import (
	"fmt"

	"github.com/untref-ayp2/data-structures/types"
	"github.com/untref-ayp2/data-structures/utils"
)

type AVLNode[T types.Ordered] struct {
	data   T           // dato
	height int         // altura
	left   *AVLNode[T] // hijo izquierdo
	right  *AVLNode[T] // hijo derecho
}

func newAVLNode[T types.Ordered](data T, left *AVLNode[T], right *AVLNode[T]) *AVLNode[T] {
	return &AVLNode[T]{left: left, right: right, data: data, height: 0}
}

func (n *AVLNode[T]) GetData() T {
	return n.data
}

func (n *AVLNode[T]) string() string {
	return fmt.Sprintf("%v", n.data)
}

func (n *AVLNode[T]) GetLeft() *AVLNode[T] {
	return n.left
}

func (n *AVLNode[T]) GetRight() *AVLNode[T] {
	return n.right
}

func (n *AVLNode[T]) GetHeight() int {
	if n == nil {
		return -1
	}

	return n.height
}

func (n *AVLNode[T]) GetBalance() int {
	if n == nil {
		return 0
	}

	return n.left.GetHeight() - n.right.GetHeight()
}

func (n *AVLNode[T]) updateHeight() {
	n.height = 1 + utils.Max(n.left.GetHeight(), n.right.GetHeight())
}

func (n *AVLNode[T]) insert(value T) *AVLNode[T] {
	if n == nil {
		return newAVLNode[T](value, nil, nil)
	}
	switch {
	case value < n.data:
		n.left = n.left.insert(value)
	case value > n.data:
		n.right = n.right.insert(value)
	default:
		return n
	}
	n.updateHeight()

	return n.applyRotation()
}

func (n *AVLNode[T]) rotateRight() *AVLNode[T] {
	y := n.left
	t2 := y.right

	y.right = n
	n.left = t2

	n.updateHeight()
	y.updateHeight()

	return y
}

func (n *AVLNode[T]) rotateLeft() *AVLNode[T] {
	x := n.right
	t2 := x.left

	x.left = n
	n.right = t2

	n.updateHeight()
	x.updateHeight()

	return x
}

func (n *AVLNode[T]) remove(value T) *AVLNode[T] {
	if n == nil {
		return n
	}

	switch {
	case value < n.data:
		n.left = n.left.remove(value)
	case value > n.data:
		n.right = n.right.remove(value)
	default:
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		temp := n.right.findMin()
		n.data = temp.data
		n.right = n.right.remove(temp.data)
	}

	n.updateHeight()

	return n.applyRotation()
}

func (n *AVLNode[T]) applyRotation() *AVLNode[T] {
	balance := n.GetBalance()

	if balance > 1 {
		if n.left.GetBalance() < 0 {
			n.left = n.left.rotateLeft()
		}

		return n.rotateRight()
	}

	if balance < -1 {
		if n.right.GetBalance() > 0 {
			n.right = n.right.rotateRight()
		}

		return n.rotateLeft()
	}

	return n
}

func (n *AVLNode[T]) findMin() *AVLNode[T] {
	if n.left == nil {
		return n
	}

	return n.left.findMin()
}

func (n *AVLNode[T]) findMax() *AVLNode[T] {
	if n.right == nil {
		return n
	}

	return n.right.findMax()
}

func (n *AVLNode[T]) search(k T) bool {
	if n == nil {
		return false
	}

	if n.data > k {
		return n.left.search(k)
	}

	if n.data < k {
		return n.right.search(k)
	}

	return true
}

func (n *AVLNode[T]) inOrder() string {
	if n == nil {
		return ""
	}

	return n.left.inOrder() + " " + n.string() + " " + n.right.inOrder()
}

func (n *AVLNode[T]) Size() int {
	if n == nil {
		return 0
	}

	return 1 + n.left.Size() + n.right.Size()
}
