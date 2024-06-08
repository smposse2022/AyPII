// Package heap provee una implementación de un heap binario.
package heap

import (
	"errors"

	"github.com/untref-ayp2/data-structures/types"
	"github.com/untref-ayp2/data-structures/utils"
)

type Heap[T any] struct {
	// contenedor de datos
	elements []T
	// Función de comparación. Para un heap de mínimo,
	// devuelve -1 si a < b, 0 si a == b, 1 si a > b
	// Para un heap de máximo, devuelve 1 si a < b, 0 si a == b, -1 si a > b
	compare func(a T, b T) int
}

// NewMinHeap crea un nuevo heap binario de mínimos.
//
// Uso:
//
//	heap := heap.NewMinHeap[int]()
//
// Retorna:
//   - un puntero a un heap binario de mínimos.
func NewMinHeap[T types.Ordered]() *Heap[T] {
	return &Heap[T]{compare: utils.Compare[T], elements: make([]T, 0)}
}

// NewMaxHeap crea un nuevo heap binario de máximos.
//
// Uso:
//
//	heap := heap.NewMaxHeap[int]()
//
// Retorna:
//   - un puntero a un heap binario de máximos.
func NewMaxHeap[T types.Ordered]() *Heap[T] {
	comp := func(a T, b T) int {
		return utils.Compare[T](b, a)
	}

	return &Heap[T]{compare: comp, elements: make([]T, 0)}
}

// NewGenericHeap crea un nuevo heap binario con una función de comparación personalizada.
//
// Uso:
//
//	heap := heap.NewGenericHeap[int](func(a int, b int) int {
//		if a < b {
//			return -1
//		}
//		if a > b {
//			return 1
//		}
//		return 0
//	})
//
// Parámetros:
//   - `comp` función de comparación personalizada.
//
// Retorna:
//   - un puntero a un heap binario con una función de comparación personalizada.
func NewGenericHeap[T any](comp func(a T, b T) int) *Heap[T] {
	return &Heap[T]{compare: comp, elements: make([]T, 0)}
}

// Size retorna la cantidad de elementos en el heap.
//
// Uso:
//
//	size := heap.Size()
//
// Retorna:
//   - la cantidad de elementos en el heap.
func (m *Heap[T]) Size() int {
	return len(m.elements)
}

// Insert agrega un elemento al heap.
//
// Uso:
//
//	heap := heap.NewMinHeap[int]()
//	heap.Insert(5)
//
// Parámetros:
//
//	element: elemento a agregar al heap.
func (m *Heap[T]) Insert(element T) {
	m.elements = append(m.elements, element)
	m.upHeap(len(m.elements) - 1)
}

// upHeap reordena el heap hacia arriba.
//
// Parámetros:
//   - `i` índice del elemento a reordenar.
func (m *Heap[T]) upHeap(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if m.compare(m.elements[i], m.elements[parent]) > 0 {
			break
		}
		m.elements[i], m.elements[parent] = m.elements[parent], m.elements[i]
		i = parent
	}
}

// Remove elimina y retorna el elemento en la cima del heap.
//
// Uso:
//
//	heap := heap.NewMinHeap[int]()
//	heap.Insert(5)
//	element, _ := heap.Remove()
//
// Retorna:
//   - el elemento en la cima del heap.
func (m *Heap[T]) Remove() (T, error) {
	var element T
	if m.Size() == 0 {
		return element, errors.New("heap vacío")
	}
	element = m.elements[0]
	m.elements[0] = m.elements[m.Size()-1]
	m.elements = m.elements[:m.Size()-1]
	m.downHeap(0)

	return element, nil
}

// downHeap reordena el heap hacia abajo.
//
// Parámetros:
//   - `i` índice del elemento a reordenar.
func (m *Heap[T]) downHeap(i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < m.Size() && m.compare(m.elements[left], m.elements[smallest]) < 0 {
			smallest = left
		}

		if right < m.Size() && m.compare(m.elements[right], m.elements[smallest]) < 0 {
			smallest = right
		}

		if smallest == i {
			break
		}

		m.elements[i], m.elements[smallest] = m.elements[smallest], m.elements[i]
		i = smallest
	}
}

func NuevoMonticuloMaxDesdeArreglo[T types.Ordered](arr []T) *Heap[T] {
    // Crear un nuevo heap de máximos
    heap := NewMaxHeap[T]()
    
    // Insertar cada elemento del arreglo en el heap
    for _, element := range arr {
        heap.Insert(element)
    }
    
    return heap
}

func EnesimoMaximo[T types.Ordered](heap *Heap[T], n int) (T, error) {
	var maximo T
	var err error
	if n < 1 || n > heap.Size() {
		return maximo, errors.New("n debe estar en el rango de 1 a M")
	}

	// Cre una copia del heap para no modificar el original
	copiaHeap := &Heap[T]{compare: heap.compare, elements: make([]T, len(heap.elements))}
	copy(copiaHeap.elements, heap.elements)

	for i := 0; i < n; i++ {
		maximo, err = copiaHeap.Remove()
		if err != nil {
			return maximo, err
		}
	}

	return maximo, nil
}

func CombinarMonticulos[T types.Ordered](heap1, heap2 *Heap[T]) *Heap[T] {
	// Determinar el tipo de heap
	var combinedHeap *Heap[T]
	if heap1.Size() > 1 && heap1.compare(heap1.elements[0], heap1.elements[1]) > 0 {
		// Si el primer elemento es mayor que el segundo, crear un nuevo max-heap
		combinedHeap = NewMaxHeap[T]()
	} else {
		// Si el primer elemento es menor que el segundo o si hay solo un elemento, crear un nuevo min-heap
		combinedHeap = NewMinHeap[T]()
	}

	// Insertar todos los elementos del primer heap en el combinado
	for _, element := range heap1.elements {
		combinedHeap.Insert(element)
	}

	// Insertar todos los elementos del segundo heap en el combinado
	for _, element := range heap2.elements {
		combinedHeap.Insert(element)
	}

	return combinedHeap
}