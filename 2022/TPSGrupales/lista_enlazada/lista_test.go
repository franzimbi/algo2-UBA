package lista_test

import (
	Lista "lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_TEST_VOLUMEN = 100000
)

func TestListaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con una lista vacia")
	lista := Lista.CrearListaEnlazada[int]()
	require.EqualValues(t, lista.EstaVacia(), true)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.EqualValues(t, lista.EstaVacia(), true)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, lista.EstaVacia(), true)
	require.EqualValues(t, lista.Largo(), 0)
	require.EqualValues(t, lista.EstaVacia(), true)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.EqualValues(t, lista.Largo(), 0)
	require.EqualValues(t, lista.EstaVacia(), true)
}
func TestListaBorrarPrimero(t *testing.T) {
	t.Log("Hacemos pruebas borrando en una lista")
	lista := Lista.CrearListaEnlazada[any]()
	require.EqualValues(t, lista.EstaVacia(), true)
	lista.InsertarPrimero(22)
	require.EqualValues(t, lista.EstaVacia(), false)
	require.EqualValues(t, lista.Largo(), 1)
	require.EqualValues(t, lista.VerPrimero(), 22)
	require.EqualValues(t, lista.EstaVacia(), false)
	require.EqualValues(t, lista.Largo(), 1)
	require.EqualValues(t, lista.VerUltimo(), 22)
	require.EqualValues(t, lista.EstaVacia(), false)
	require.EqualValues(t, lista.Largo(), 1)
	require.EqualValues(t, lista.BorrarPrimero(), 22)
	require.EqualValues(t, lista.EstaVacia(), true)
	require.EqualValues(t, lista.Largo(), 0)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.EqualValues(t, lista.EstaVacia(), true)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	lista.InsertarPrimero("hola")
	lista.InsertarUltimo("como")
	lista.InsertarUltimo("estas")
	require.EqualValues(t, lista.EstaVacia(), false)
	require.EqualValues(t, lista.Largo(), 3)
	require.EqualValues(t, lista.VerPrimero(), "hola")
	require.EqualValues(t, lista.VerUltimo(), "estas")
	require.EqualValues(t, lista.EstaVacia(), false)
	require.EqualValues(t, lista.Largo(), 3)
	require.EqualValues(t, lista.BorrarPrimero(), "hola")
	require.EqualValues(t, lista.EstaVacia(), false)
	require.EqualValues(t, lista.Largo(), 2)
	require.EqualValues(t, lista.VerPrimero(), "como")
	require.EqualValues(t, lista.VerUltimo(), "estas")
	require.EqualValues(t, lista.EstaVacia(), false)
	require.EqualValues(t, lista.Largo(), 2)
	require.EqualValues(t, lista.BorrarPrimero(), "como")
	require.EqualValues(t, lista.VerPrimero(), "estas")
	require.EqualValues(t, lista.VerUltimo(), "estas")
	require.EqualValues(t, lista.EstaVacia(), false)
	require.EqualValues(t, lista.Largo(), 1)
	require.EqualValues(t, lista.BorrarPrimero(), "estas")
	require.EqualValues(t, lista.EstaVacia(), true)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.EqualValues(t, lista.EstaVacia(), true)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, lista.EstaVacia(), true)
	require.EqualValues(t, lista.Largo(), 0)
	require.EqualValues(t, lista.EstaVacia(), true)
}

func TestVolumen(t *testing.T) {
	t.Log("Hacemos pruebas de volumen de ", _TEST_VOLUMEN, " elementos")
	//creo un arreglo de ints
	var cosas [_TEST_VOLUMEN]int
	for i := 0; i < _TEST_VOLUMEN; i++ {
		cosas[i] = i + 1
	}
	lista := Lista.CrearListaEnlazada[int]()
	for i := 0; i < _TEST_VOLUMEN; i++ {
		lista.InsertarUltimo(cosas[i])
		require.EqualValues(t, lista.Largo(), i+1)
		require.EqualValues(t, lista.VerUltimo(), cosas[i])
		require.EqualValues(t, lista.EstaVacia(), false)
	}
	for i := 0; i < _TEST_VOLUMEN; i++ {
		require.EqualValues(t, lista.VerPrimero(), cosas[i])
		require.EqualValues(t, lista.BorrarPrimero(), cosas[i])
		require.EqualValues(t, lista.Largo(), _TEST_VOLUMEN-i-1)
	}
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.EqualValues(t, lista.EstaVacia(), true)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, lista.Largo(), 0)
}

func TestIterInterno(t *testing.T) {
	t.Log("Hacemos pruebas con el iterador interno de lista")
	lista := Lista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i + 1) //1,2,3,4,5,6,7,8,9,10
	}
	var total int
	lista.Iterar(func(dato int) bool {
		total += dato
		return true
	})
	require.EqualValues(t, total, 55)
	acum := 0
	lista.Iterar(func(dato int) bool {
		acum++
		return dato < 5
	})
	require.EqualValues(t, acum, 5)
	require.EqualValues(t, lista.VerPrimero(), 1)
	require.EqualValues(t, lista.VerUltimo(), 10)
	require.EqualValues(t, lista.Largo(), 10)
	require.EqualValues(t, lista.EstaVacia(), false)
}

func TestIterInterno2(t *testing.T) {
	t.Log("Hacemos pruebas con el iterador interno de lista")
	lista := Lista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i + 1) //1,2,3,4,5,6,7,8,9,10
	}
	for i := 0; i < 10; i++ {
		require.EqualValues(t, lista.BorrarPrimero(), 1+i)
	}
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.EqualValues(t, lista.EstaVacia(), true)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, lista.Largo(), 0)
	acum := 0

	lista.Iterar(func(dato int) bool {
		acum++
		return true
	})
	require.EqualValues(t, acum, 0)
	require.EqualValues(t, lista.Largo(), 0)
}

func TestIterExternoListaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con un iterador externo de una lista vacia")
	lista := Lista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
	require.EqualValues(t, iter.HaySiguiente(), false)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.EqualValues(t, lista.EstaVacia(), true)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, lista.Largo(), 0)
}

func TestIterExternoCiclo(t *testing.T) {
	t.Log("Hacemos pruebas de ciclo con un iterador externo de una lista")
	lista := Lista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(1 + i)
	}
	iter := lista.Iterador()
	var i, total int
	//
	for iter.HaySiguiente() {
		if i == 5 {
			break
		}
		i++
		total += iter.VerActual()
		iter.Siguiente()
	}
	require.EqualValues(t, iter.VerActual(), 6)
	require.EqualValues(t, total, 15)
	require.EqualValues(t, iter.HaySiguiente(), true)
	require.EqualValues(t, iter.VerActual(), 6)
	require.EqualValues(t, iter.Siguiente(), 6)
	require.EqualValues(t, iter.VerActual(), 7)
	for iter.HaySiguiente() {
		i++
		iter.Siguiente()
	}
	require.EqualValues(t, iter.HaySiguiente(), false)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.EqualValues(t, i, 9)
	require.EqualValues(t, lista.VerPrimero(), 1)
	require.EqualValues(t, lista.VerUltimo(), 10)
	require.EqualValues(t, lista.Largo(), 10)
	require.EqualValues(t, lista.EstaVacia(), false)
}

func TestIterExternoInsertarVacio(t *testing.T) {
	t.Log("Hacemos pruebas con un iterador externo de una lista vacia insertando")
	lista := Lista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(1)
	require.EqualValues(t, lista.VerPrimero(), 1)
	require.EqualValues(t, lista.VerUltimo(), 1)
	require.EqualValues(t, lista.Largo(), 1)
	require.EqualValues(t, lista.EstaVacia(), false)
	require.EqualValues(t, iter.HaySiguiente(), true)
	require.EqualValues(t, iter.VerActual(), 1)
	iter.Insertar(3)
	require.EqualValues(t, lista.VerPrimero(), 3)
	require.EqualValues(t, lista.VerUltimo(), 1)
	require.EqualValues(t, lista.Largo(), 2)
	require.EqualValues(t, lista.EstaVacia(), false)
	require.EqualValues(t, iter.HaySiguiente(), true)
	require.EqualValues(t, iter.VerActual(), 3)
	require.EqualValues(t, iter.Siguiente(), 3)
	require.EqualValues(t, iter.VerActual(), 1)
	require.EqualValues(t, iter.HaySiguiente(), true)
	require.EqualValues(t, lista.VerPrimero(), 3)
	require.EqualValues(t, lista.VerUltimo(), 1)
	iter.Insertar(2)
	require.EqualValues(t, lista.VerPrimero(), 3)
	require.EqualValues(t, lista.VerUltimo(), 1)
	require.EqualValues(t, lista.Largo(), 3)
	require.EqualValues(t, lista.EstaVacia(), false)
	require.EqualValues(t, iter.HaySiguiente(), true)
	require.EqualValues(t, iter.Siguiente(), 2)
	require.EqualValues(t, iter.VerActual(), 1)
	require.EqualValues(t, iter.Siguiente(), 1)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIterExternoBorrar(t *testing.T) {
	t.Log("Hacemos pruebas con un iterador externo de una lista llena borrando")
	lista := Lista.CrearListaEnlazada[int]()
	for i := 0; i < 5; i++ {
		lista.InsertarUltimo(i + 1)
	}
	iter := lista.Iterador()
	require.EqualValues(t, iter.Borrar(), 1)
	require.EqualValues(t, iter.VerActual(), 2)
	for iter.HaySiguiente() {
		if iter.VerActual() == 3 {
			break
		}
		iter.Siguiente()
	}
	require.EqualValues(t, iter.Borrar(), 3)
	require.EqualValues(t, iter.VerActual(), 4)
	require.EqualValues(t, iter.Siguiente(), 4)
	require.EqualValues(t, iter.VerActual(), 5)
	require.EqualValues(t, iter.Borrar(), 5)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.EqualValues(t, lista.VerPrimero(), 2)
	require.EqualValues(t, lista.VerUltimo(), 4)
	require.EqualValues(t, lista.Largo(), 2)
	require.EqualValues(t, lista.EstaVacia(), false)
}

func TestIterExternoInsertarBorrar(t *testing.T) {
	t.Log("Hacemos pruebas con un iterador externo de una lista insertando y borrando")
	lista := Lista.CrearListaEnlazada[any]()
	iter := lista.Iterador()
	iter.Insertar("estas")
	require.EqualValues(t, iter.VerActual(), "estas")
	iter.Insertar("como")
	require.EqualValues(t, lista.VerPrimero(), "como")
	require.EqualValues(t, lista.VerUltimo(), "estas")
	require.EqualValues(t, lista.Largo(), 2)
	require.EqualValues(t, iter.VerActual(), "como")
	require.EqualValues(t, iter.Borrar(), "como")
	require.EqualValues(t, iter.VerActual(), "estas")
	require.EqualValues(t, lista.VerPrimero(), "estas")
	require.EqualValues(t, lista.VerUltimo(), "estas")
	require.EqualValues(t, lista.Largo(), 1)
	iter.Insertar("como")
	require.EqualValues(t, lista.VerPrimero(), "como")
	require.EqualValues(t, lista.VerUltimo(), "estas")
	require.EqualValues(t, lista.Largo(), 2)
	iter.Insertar(nil)
	require.EqualValues(t, lista.VerPrimero(), nil)
	require.EqualValues(t, lista.VerUltimo(), "estas")
	require.EqualValues(t, lista.Largo(), 3)
	require.EqualValues(t, iter.VerActual(), nil)
	require.EqualValues(t, iter.Borrar(), nil)
	require.EqualValues(t, lista.VerPrimero(), "como")
	require.EqualValues(t, lista.VerUltimo(), "estas")
	require.EqualValues(t, lista.Largo(), 2)
	require.EqualValues(t, iter.VerActual(), "como")
	require.EqualValues(t, iter.Siguiente(), "como")
	require.EqualValues(t, iter.VerActual(), "estas")
	require.EqualValues(t, iter.Borrar(), "estas")
	require.EqualValues(t, lista.VerPrimero(), "como")
	require.EqualValues(t, lista.VerUltimo(), "como")
	require.EqualValues(t, lista.Largo(), 1)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	lista.BorrarPrimero()
	require.EqualValues(t, lista.Largo(), 0)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.EqualValues(t, lista.EstaVacia(), true)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, lista.EstaVacia(), true)
	require.EqualValues(t, lista.Largo(), 0)
}

func TestIterExternoInsertarBorrar2(t *testing.T) {
	t.Log("Hacemos pruebas con un iterador externo de una lista insertando y borrando")
	lista := Lista.CrearListaEnlazada[any]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i + 1)
	}
	iter := lista.Iterador()
	impar := 1
	for iter.HaySiguiente() {
		if (impar % 2) == 1 {
			require.EqualValues(t, iter.Borrar(), impar)
			impar++
		}
		impar++
		iter.Siguiente()
	}
	require.EqualValues(t, lista.VerPrimero(), 2)
	require.EqualValues(t, lista.VerUltimo(), 10)
	require.EqualValues(t, lista.Largo(), 5)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	iter2 := lista.Iterador()
	i := 2
	for iter2.HaySiguiente() {
		require.EqualValues(t, iter2.VerActual(), i)
		require.EqualValues(t, iter2.Borrar(), i)
		i += 2
	}
	require.EqualValues(t, lista.Largo(), 0)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.EqualValues(t, lista.EstaVacia(), true)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, lista.EstaVacia(), true)
}

func TestIterExternoVolume(t *testing.T) {
	t.Log("Hacemos pruebas de volumen de ", _TEST_VOLUMEN, " elementos con iterador externo")
	lista := Lista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	for i := 0; i < _TEST_VOLUMEN; i++ {
		iter.Insertar(_TEST_VOLUMEN - i)
		require.EqualValues(t, iter.VerActual(), _TEST_VOLUMEN-i)
	}
	require.EqualValues(t, lista.VerPrimero(), 1)
	require.EqualValues(t, lista.VerUltimo(), _TEST_VOLUMEN)
	require.EqualValues(t, lista.Largo(), _TEST_VOLUMEN)

	iter2 := lista.Iterador()
	for i := 0; i < _TEST_VOLUMEN; i++ {
		require.EqualValues(t, iter2.Borrar(), i+1)
		if i < _TEST_VOLUMEN-1 {
			require.EqualValues(t, iter2.VerActual(), i+2)
		}
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter2.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter2.Siguiente() })
	require.EqualValues(t, lista.Largo(), 0)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.EqualValues(t, lista.EstaVacia(), true)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, lista.EstaVacia(), true)
}
