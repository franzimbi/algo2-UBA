package cola_test

import (
	Cola "cola"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	_CONSTANTE_TEST_VOLUMEN = 100000
)

func TestColaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con una cola vacia")
	cola := Cola.CrearColaEnlazada[int]()
	require.EqualValues(t, cola.EstaVacia(), true)
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.EqualValues(t, cola.EstaVacia(), true)
	//
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.EqualValues(t, cola.EstaVacia(), true)
}

func TestColaEncolarDesencolarInt(t *testing.T) {
	t.Log("Hacemos pruebas con una cola de ints")
	cola := Cola.CrearColaEnlazada[int]()
	cola.Encolar(34)
	require.EqualValues(t, cola.EstaVacia(), false)
	require.EqualValues(t, cola.VerPrimero(), 34)
	require.EqualValues(t, cola.EstaVacia(), false)
	//
	cola.Encolar(-121)
	require.EqualValues(t, cola.EstaVacia(), false)
	require.EqualValues(t, cola.VerPrimero(), 34)
	require.EqualValues(t, cola.EstaVacia(), false)
	//
	cola.Encolar(1)
	require.EqualValues(t, cola.EstaVacia(), false)
	require.EqualValues(t, cola.VerPrimero(), 34)
	require.EqualValues(t, cola.Desencolar(), 34)
	require.EqualValues(t, cola.VerPrimero(), -121)
	require.EqualValues(t, cola.EstaVacia(), false)
	require.EqualValues(t, cola.Desencolar(), -121)
	require.EqualValues(t, cola.EstaVacia(), false)
	require.EqualValues(t, cola.VerPrimero(), 1)
	require.EqualValues(t, cola.Desencolar(), 1)
	require.EqualValues(t, cola.EstaVacia(), true)
	//
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.EqualValues(t, cola.EstaVacia(), true)
}

func TestColaEncolarDesencolarMucho(t *testing.T) {
	t.Log("Hacemos pruebas encolando y desencolando mucho")
	cola := Cola.CrearColaEnlazada[any]()
	var a int = 3
	var b *int = &a
	//
	cola.Encolar(b)
	require.EqualValues(t, cola.EstaVacia(), false)
	require.EqualValues(t, cola.VerPrimero(), b)
	require.EqualValues(t, cola.Desencolar(), b)
	require.EqualValues(t, cola.EstaVacia(), true)
	//
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	//
	cola.Encolar("hola")
	require.EqualValues(t, cola.EstaVacia(), false)
	require.EqualValues(t, cola.VerPrimero(), "hola")
	require.EqualValues(t, cola.Desencolar(), "hola")
	require.EqualValues(t, cola.EstaVacia(), true)
	//
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	//
	cola.Encolar(nil)
	cola.Encolar(&a)
	cola.Encolar(true)
	require.EqualValues(t, cola.VerPrimero(), nil)
	require.EqualValues(t, cola.Desencolar(), nil)
	require.EqualValues(t, cola.VerPrimero(), &a)
	require.EqualValues(t, cola.Desencolar(), &a)
	require.EqualValues(t, cola.VerPrimero(), true)
	require.EqualValues(t, cola.Desencolar(), true)
	require.EqualValues(t, cola.EstaVacia(), true)
	//
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}
func TestVolumen(t *testing.T) {
	t.Log("Hacemos pruebas de volumen")
	//creo un arreglo de ints
	var cosas [_CONSTANTE_TEST_VOLUMEN]int
	for i := 0; i < _CONSTANTE_TEST_VOLUMEN; i++ {
		cosas[i] = i + 1
	}
	//
	cola := Cola.CrearColaEnlazada[int]()
	//
	for i := 0; i < _CONSTANTE_TEST_VOLUMEN; i++ {
		cola.Encolar(cosas[i])
		require.EqualValues(t, cola.EstaVacia(), false)
		require.EqualValues(t, cola.VerPrimero(), cosas[0])
	}
	//
	for i := 0; i < _CONSTANTE_TEST_VOLUMEN; i++ {
		require.EqualValues(t, cola.VerPrimero(), cosas[i])
		require.EqualValues(t, cola.Desencolar(), cosas[i])
		if i < _CONSTANTE_TEST_VOLUMEN-1 {
			require.EqualValues(t, cola.VerPrimero(), cosas[i+1])
		}
	}
	//
	require.EqualValues(t, cola.EstaVacia(), true)
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}
