package cola_prioridad_test

import (
	TDAheap "cola_prioridad"
	//"crypto/rand"
	"math/rand"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const _CANTIDADVOLUMEN = 1000000

func cmpInts(a, b int) int {
	return a - b
}

func TestHeapVacio(t *testing.T) {
	t.Log("Hacemos pruebas con un heap vacio")
	heap := TDAheap.CrearHeap[string](strings.Compare)
	require.EqualValues(t, heap.EstaVacia(), true)
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.EqualValues(t, heap.EstaVacia(), true)
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, heap.EstaVacia(), true)
}

func TestHeapDeInts(t *testing.T) {
	t.Log("Hacemos pruebas encolando y desencolando ints en el heap")
	heap := TDAheap.CrearHeap[int](cmpInts)
	heap.Encolar(2)
	//
	require.EqualValues(t, heap.EstaVacia(), false)
	require.EqualValues(t, heap.VerMax(), 2)
	require.EqualValues(t, heap.Cantidad(), 1)
	//
	heap.Encolar(10)
	require.EqualValues(t, heap.VerMax(), 10)
	require.EqualValues(t, heap.Cantidad(), 2)
	//
	heap.Encolar(1)
	require.EqualValues(t, heap.VerMax(), 10)
	require.EqualValues(t, heap.Cantidad(), 3)
	//
	heap.Encolar(13)
	require.EqualValues(t, heap.VerMax(), 13)
	require.EqualValues(t, heap.Cantidad(), 4)
	//
	require.EqualValues(t, heap.Desencolar(), 13)
	require.EqualValues(t, heap.Cantidad(), 3)
	//
	require.EqualValues(t, heap.Desencolar(), 10)
	require.EqualValues(t, heap.Cantidad(), 2)
	//
	require.EqualValues(t, heap.Desencolar(), 2)
	require.EqualValues(t, heap.Cantidad(), 1)
	//
	require.EqualValues(t, heap.Desencolar(), 1)
	require.EqualValues(t, heap.Cantidad(), 0)
	//
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.EqualValues(t, heap.EstaVacia(), true)
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestEncolarUnIntNegativoEnElMedio(t *testing.T) {
	t.Log("Hacemos pruebas encolando un -1(int) en el heap")
	heap := TDAheap.CrearHeap[int](cmpInts)
	//
	heap.Encolar(1)
	require.EqualValues(t, heap.EstaVacia(), false)
	require.EqualValues(t, heap.VerMax(), 1)
	require.EqualValues(t, heap.Cantidad(), 1)
	//
	heap.Encolar(2)
	heap.Encolar(3)
	require.EqualValues(t, heap.VerMax(), 3)
	require.EqualValues(t, heap.Cantidad(), 3)
	heap.Encolar(4)
	heap.Encolar(-1)
	//
	require.EqualValues(t, heap.Desencolar(), 4)
	require.EqualValues(t, heap.VerMax(), 3)
	//
	require.EqualValues(t, heap.Desencolar(), 3)
	require.EqualValues(t, heap.VerMax(), 2)
	//
	require.EqualValues(t, heap.Desencolar(), 2)
	require.EqualValues(t, heap.VerMax(), 1)
	//
	require.EqualValues(t, heap.Desencolar(), 1)
	require.EqualValues(t, heap.VerMax(), -1)
	require.EqualValues(t, heap.Desencolar(), -1)
	//
	require.EqualValues(t, heap.Cantidad(), 0)
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.EqualValues(t, heap.EstaVacia(), true)
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestEncolarDesencolarStrings(t *testing.T) {
	t.Log("Hacemos pruebas encolando strings en el heap")
	heap := TDAheap.CrearHeap[string](strings.Compare)
	heap.Encolar("f")
	//
	require.EqualValues(t, heap.EstaVacia(), false)
	require.EqualValues(t, heap.VerMax(), "f")
	require.EqualValues(t, heap.Cantidad(), 1)
	//
	heap.Encolar("g")
	require.EqualValues(t, heap.VerMax(), "g")
	require.EqualValues(t, heap.Cantidad(), 2)
	//
	heap.Encolar("a")
	require.EqualValues(t, heap.VerMax(), "g")
	require.EqualValues(t, heap.Cantidad(), 3)
	//
	heap.Encolar("i")
	require.EqualValues(t, heap.VerMax(), "i")
	require.EqualValues(t, heap.Cantidad(), 4)
	//
	require.EqualValues(t, heap.Desencolar(), "i")
	require.EqualValues(t, heap.Cantidad(), 3)
	//
	require.EqualValues(t, heap.Desencolar(), "g")
	require.EqualValues(t, heap.Cantidad(), 2)
	//
	require.EqualValues(t, heap.Desencolar(), "f")
	require.EqualValues(t, heap.Cantidad(), 1)
	//
	require.EqualValues(t, heap.Desencolar(), "a")
	require.EqualValues(t, heap.Cantidad(), 0)
	//
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.EqualValues(t, heap.EstaVacia(), true)
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestVolumen(t *testing.T) {
	t.Log("Hacemos pruebas de volumen")
	heap := TDAheap.CrearHeap[int](cmpInts)
	//
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 2147483646
	var elementos []int
	for i := 0; i < _CANTIDADVOLUMEN; i++ {
		elementos = append(elementos, rand.Intn(max-min+1)+min)
	}
	for i := 0; i < _CANTIDADVOLUMEN; i++ {
		heap.Encolar(elementos[i])
		require.EqualValues(t, heap.Cantidad(), i+1)
	}
	sort.Ints(elementos)
	require.EqualValues(t, heap.Cantidad(), _CANTIDADVOLUMEN)
	require.EqualValues(t, heap.VerMax(), elementos[_CANTIDADVOLUMEN-1])
	for i := _CANTIDADVOLUMEN - 1; i >= 0; i-- {
		require.EqualValues(t, heap.Desencolar(), elementos[i])
	}
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.EqualValues(t, heap.EstaVacia(), true)
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestCrearHeapConArreglo(t *testing.T) {
	t.Log("Hacemos pruebas creando un heap con un arreglo desordenado")
	arr := []int{5, 3, 10, 8, 0, 9, -1, 13, 2, 2}
	heap := TDAheap.CrearHeapArr[int](arr, cmpInts)
	//
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		require.EqualValues(t, heap.Desencolar(), arr[len(arr)-i-1])
	}
}

func TestHeapify(t *testing.T) {
	t.Log("ondenando con heapsort un arreglo desordenado")
	arr := []string{"f", "r", "g", "h", "a", "w", "q", "z", "s", "l", "h", "b", "w", "y"}
	TDAheap.HeapSort[string](arr, strings.Compare)
	require.EqualValues(t, sort.StringsAreSorted(arr), true)
}
