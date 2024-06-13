package ejercicios_avl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPQueue(t *testing.T) {
	q := NewPQueue[int]()

	assert.NotNil(t, q)
	assert.True(t, q.IsEmpty())
}

func TestPQueueEnqueue(t *testing.T) {
	q := NewPQueue[int]()

	q.Enqueue(1)
	assert.False(t, q.IsEmpty())
}

func TestPQueueDequeue(t *testing.T) {
	q := NewPQueue[int]()

	q.Enqueue(1)
	q.Enqueue(2)

	v, _ := q.Dequeue()
	assert.Equal(t, 2, v)

	v, _ = q.Dequeue()
	assert.Equal(t, 1, v)

	assert.True(t, q.IsEmpty())
}

func TestPQueueDequeueOnEmptyQueue(t *testing.T) {
	q := NewPQueue[int]()

	_, err := q.Dequeue()
	assert.EqualError(t, err, "cola vacía")
}

func TestPQueueFrontOnEmptyQueue(t *testing.T) {
	q := NewPQueue[int]()

	_, err := q.Front()
	assert.EqualError(t, err, "cola vacía")
}

func TestPQueueFront(t *testing.T) {
	q := NewPQueue[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	v, _ := q.Front()
	assert.Equal(t, 3, v)

	q.Dequeue()

	v, _ = q.Front()
	assert.Equal(t, 2, v)

	q.Dequeue()

	v, _ = q.Front()
	assert.Equal(t, 1, v)
}
