package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeapCrearVacio(t *testing.T) {
	m := NewMaxHeap[int]()
	assert.Equal(t, 0, m.Size())
}

func TestMaxHeapRemoveMaxVacio(t *testing.T) {
	m := NewMaxHeap[int]()
	_, err := m.Remove()
	assert.NotNil(t, err)
}

// Gracias a visualgo.net/en/heap
// por la ayuda para preparar este caso de prueba.
//
// Insertando los siguientes elementos en orden:
// 44, 29, 58, 2, 98, 11, 65, 3, 68, 99
//
// El arbol resultante debería ser:
//
//	[99]
//	├── [98]
//	│   ├── [58]
//	│   │   ├── [2]
//	│   │   └── [3]
//	│   └── [68]
//	│       └── [29]
//	└── [65]
//	    ├── [11]
//	    └── [44]
//
// Como arreglo:
// [99, 98, 65, 58, 68, 11, 44, 2, 3, 29].
func TestMaxHeapCrearInsertarYExtraer(t *testing.T) {
	secuenciaDeInsercion := []int{44, 29, 58, 2, 98, 11, 65, 3, 68, 99}

	ordenEsperadoDespuesDeInsertar := [][]int{
		{44},
		{44, 29},
		{58, 29, 44},
		{58, 29, 44, 2},
		{98, 58, 44, 2, 29},
		{98, 58, 44, 2, 29, 11},
		{98, 58, 65, 2, 29, 11, 44},
		{98, 58, 65, 3, 29, 11, 44, 2},
		{98, 68, 65, 58, 29, 11, 44, 2, 3},
		{99, 98, 65, 58, 68, 11, 44, 2, 3, 29},
	}

	// Verificaciones iniciales
	m := NewMaxHeap[int]()
	assert.Equal(t, 0, m.Size())

	// Verificaciones a medida que vamos insertando
	for i := 0; i < len(secuenciaDeInsercion); i++ {
		m.Insert(secuenciaDeInsercion[i])
		assert.Equal(t, ordenEsperadoDespuesDeInsertar[i], m.elements)
	}

	ordenEsperadoDespuesDeEliminar := [][]int{
		{98, 68, 65, 58, 29, 11, 44, 2, 3},
		{68, 58, 65, 3, 29, 11, 44, 2},
		{65, 58, 44, 3, 29, 11, 2},
		{58, 29, 44, 3, 2, 11},
		{44, 29, 11, 3, 2},
		{29, 3, 11, 2},
		{11, 3, 2},
		{3, 2},
		{2},
		{},
	}

	for i := 0; i < len(secuenciaDeInsercion); i++ {
		_, err := m.Remove()
		assert.Equal(t, ordenEsperadoDespuesDeEliminar[i], m.elements)
		assert.NoError(t, err)
	}
}
// Test para verificar que el heap contiene todos los elementos del arreglo de entrada.
func TestNuevoMonticuloMaxDesdeArreglo_ContieneTodosLosElementos(t *testing.T) {
	arr := []int{3, 1, 6, 5, 2, 4}
	heap := NuevoMonticuloMaxDesdeArreglo(arr)

	assert.Equal(t, len(arr), heap.Size(), "El heap debe contener el mismo número de elementos que el arreglo de entrada")
}

// Test para verificar que los elementos están organizados correctamente según las propiedades de un max-heap.
func TestNuevoMonticuloMaxDesdeArreglo_PropiedadesMaxHeap(t *testing.T) {
	arr := []int{3, 1, 6, 5, 2, 4}
	heap := NuevoMonticuloMaxDesdeArreglo(arr)

	// Verificar que cada padre es mayor o igual a sus hijos
	for i := 0; i < heap.Size()/2; i++ {
		left := 2*i + 1
		right := 2*i + 2

		if left < heap.Size() {
			assert.True(t, heap.compare(heap.elements[i], heap.elements[left]) >= 0, "El padre debe ser mayor o igual que el hijo izquierdo")
		}

		if right < heap.Size() {
			assert.True(t, heap.compare(heap.elements[i], heap.elements[right]) >= 0, "El padre debe ser mayor o igual que el hijo derecho")
		}
	}
}

// Test para verificar el comportamiento con un arreglo vacío.
func TestNuevoMonticuloMaxDesdeArreglo_ArregloVacio(t *testing.T) {
	arr := []int{}
	heap := NuevoMonticuloMaxDesdeArreglo(arr)

	assert.Equal(t, 0, heap.Size(), "El heap debe estar vacío cuando el arreglo de entrada está vacío")
}

// TestEnesimoMaximo_Valido verifica el enésimo máximo válido
func TestEnesimoMaximo_Valido(t *testing.T) {
	heap := NewMaxHeap[int]()
	heap.Insert(3)
	heap.Insert(1)
	heap.Insert(6)
	heap.Insert(5)
	heap.Insert(2)
	heap.Insert(4)

	// Obtener el 3er máximo
	tercerMaximo, err := EnesimoMaximo(heap, 3)
	assert.NoError(t, err)
	assert.Equal(t, 4, tercerMaximo)
}

// TestEnesimoMaximo_FueraDeRango verifica cuando n está fuera del rango
func TestEnesimoMaximo_FueraDeRango(t *testing.T) {
	heap := NewMaxHeap[int]()
	heap.Insert(3)
	heap.Insert(1)
	heap.Insert(6)
	heap.Insert(5)
	heap.Insert(2)
	heap.Insert(4)

	// Intentar obtener el 7mo máximo de un heap con solo 6 elementos
	_, err := EnesimoMaximo(heap, 7)
	assert.Error(t, err)
}

// TestEnesimoMaximo_HeapVacio verifica cuando el heap está vacío
func TestEnesimoMaximo_HeapVacio(t *testing.T) {
	heap := NewMaxHeap[int]()

	// Intentar obtener el 1er máximo de un heap vacío
	_, err := EnesimoMaximo(heap, 1)
	assert.Error(t, err)
}

func TestCombinarMonticulos_MinHeapYMinHeap(t *testing.T) {
	// Crear dos min-heaps
	heap1 := NewMinHeap[int]()
	heap2 := NewMinHeap[int]()

	heap1.Insert(3)
	heap1.Insert(5)
	heap1.Insert(7)

	heap2.Insert(2)
	heap2.Insert(4)
	heap2.Insert(6)

	// Combinar los dos montículos
	combinedHeap := CombinarMonticulos(heap1, heap2)

	// Verificar que el montículo combinado es un min-heap
	assert.True(t, combinedHeap.compare(combinedHeap.elements[0], combinedHeap.elements[1]) <= 0)
}

func TestCombinarMonticulos_MaxHeapYMaxHeap(t *testing.T) {
	// Crear dos max-heaps
	heap1 := NewMaxHeap[int]()
	heap2 := NewMaxHeap[int]()

	heap1.Insert(7)
	heap1.Insert(5)
	heap1.Insert(3)

	heap2.Insert(6)
	heap2.Insert(4)
	heap2.Insert(2)

	// Combinar los dos montículos
	combinedHeap := CombinarMonticulos(heap1, heap2)

	// Verificar que el montículo combinado es un max-heap
	assert.True(t, combinedHeap.compare(combinedHeap.elements[0], combinedHeap.elements[1]) >= 0)
}

func TestCombinarMonticulos_MinHeapYMaxHeap(t *testing.T) {
	// Crear un min-heap y un max-heap
	heap1 := NewMinHeap[int]()
	heap2 := NewMaxHeap[int]()

	heap1.Insert(1)
	heap1.Insert(2)
	heap1.Insert(3)

	heap2.Insert(6)
	heap2.Insert(5)
	heap2.Insert(4)

	// Combinar los dos montículos
	combinedHeap := CombinarMonticulos(heap1, heap2)

	// Verificar que el primer elemento del montículo combinado sea menor que el segundo para un min-heap
	// y mayor para un max-heap
	assert.True(t, combinedHeap.compare(combinedHeap.elements[0], combinedHeap.elements[1]) <= 0) // Para un min-heap
}