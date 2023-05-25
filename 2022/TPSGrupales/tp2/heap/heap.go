package cola_prioridad

const (
	_tamanoInicial        = 15
	_constanteRedimension = 4
)

type heap[T comparable] struct {
	arr      []T
	cantidad int
	cmp      func(T, T) int
}

func swap[T comparable](a, b *T) {
	*a, *b = *b, *a
}
func max[T comparable](arreglo []T, pos1, pos2, cantidad int, cmp func(T, T) int) int {
	if pos2 >= cantidad {
		return pos1
	}
	if cmp(arreglo[pos1], arreglo[pos2]) >= 0 {
		return pos1
	}
	return pos2
}

func upheap[T comparable](arreglo []T, pos int, cmp func(T, T) int) {
	if pos == 0 {
		return
	}
	padre := (pos - 1) / 2 // padre = i - 1/ 2
	if cmp(arreglo[pos], arreglo[padre]) > 0 {
		swap(&arreglo[pos], &arreglo[padre])
		upheap(arreglo, padre, cmp)
	}
}
func downheap[T comparable](arreglo []T, pos, cantidad int, cmp func(T, T) int) {
	if pos == cantidad {
		return
	}
	hijoIzq := 2*pos + 1
	hijoDer := 2*pos + 2
	if hijoIzq >= cantidad {
		return
	}
	hijoMayor := max(arreglo, hijoIzq, hijoDer, cantidad, cmp) // hijo izq = 2i + 1  |  hijo der = 2i + 2
	if cmp(arreglo[pos], arreglo[hijoMayor]) < 0 {
		swap(&arreglo[pos], &arreglo[hijoMayor])
		downheap(arreglo, hijoMayor, cantidad, cmp)
	}
}
func heapRedimensionar[T comparable](heap *heap[T], tamNuevo int) {
	if tamNuevo < _tamanoInicial {
		tamNuevo = _tamanoInicial
	}
	newArr := make([]T, tamNuevo)
	copy(newArr, heap.arr)
	(*heap).arr = newArr
}
func heapify[T comparable](arreglo []T, cantidad int, cmp func(T, T) int) {
	for i := 0; i < cantidad; i++ {
		downheap(arreglo, cantidad-i-1, cantidad, cmp)
	}
}

// - - - - - - - - - - - - - - - - PRIMITIVAS HEAP - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{arr: make([]T, _tamanoInicial), cmp: funcion_cmp}
}
func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := &heap[T]{arr: make([]T, len(arreglo)), cmp: funcion_cmp, cantidad: len(arreglo)}
	copy(heap.arr, arreglo)
	heapify(heap.arr, heap.cantidad, funcion_cmp)
	return heap
}
func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, len(elementos), funcion_cmp)
	tam := len(elementos)
	for i := 0; i < tam; i++ {
		swap(&elementos[0], &elementos[tam-i-1])
		downheap(elementos, 0, tam-i-1, funcion_cmp)
	}
}
func (heap heap[T]) Cantidad() int {
	return heap.cantidad
}
func (heap heap[T]) EstaVacia() bool {
	return heap.Cantidad() == 0
}
func (heap heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.arr[0]
}
func (heap *heap[T]) Encolar(elem T) {
	if heap.cantidad == len(heap.arr) {
		heapRedimensionar(heap, heap.cantidad*2)
	}
	(*heap).arr[heap.cantidad] = elem
	(*heap).cantidad++
	upheap(heap.arr, heap.cantidad-1, heap.cmp)
}

func (heap *heap[T]) Desencolar() T {
	maximo := heap.VerMax()
	if heap.cantidad*_constanteRedimension <= len(heap.arr) && len(heap.arr) != _tamanoInicial {
		heapRedimensionar(heap, len(heap.arr)/2)
	}
	var nulo T
	(*heap).arr[0] = nulo
	(*heap).cantidad--
	swap(&heap.arr[0], &heap.arr[heap.cantidad])
	downheap(heap.arr, 0, heap.cantidad, heap.cmp)
	return maximo
}
