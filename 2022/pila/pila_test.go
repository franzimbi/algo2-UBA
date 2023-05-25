package pila_test

import (
	//"math/rand"

	Pila "pila"
	"testing"

	//"time"

	"github.com/stretchr/testify/require"
)

const (
	_CONSTANTETESTVOLUMEN = 10000
)

func TestPilaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con una pila vacia")
	pila := Pila.CrearPilaDinamica[int]()
	require.EqualValues(t, pila.EstaVacia(), true)
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.EqualValues(t, pila.EstaVacia(), true)
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.EqualValues(t, pila.EstaVacia(), true)
}

func TestApilarDesapilarInt(t *testing.T) {
	t.Log("Hacemos pruebas apilando ints en la pila")
	pila := Pila.CrearPilaDinamica[int]()
	pila.Apilar(15)
	require.EqualValues(t, pila.EstaVacia(), false)
	require.EqualValues(t, pila.VerTope(), 15)
	require.EqualValues(t, pila.EstaVacia(), false)
	//
	pila.Apilar(3)
	require.EqualValues(t, pila.VerTope(), 3)
	require.EqualValues(t, pila.EstaVacia(), false)
	//
	pila.Apilar(1243)
	require.EqualValues(t, pila.EstaVacia(), false)
	require.EqualValues(t, pila.VerTope(), 1243)
	//
	require.EqualValues(t, pila.Desapilar(), 1243)
	require.EqualValues(t, pila.VerTope(), 3)
	require.EqualValues(t, pila.EstaVacia(), false)
	require.EqualValues(t, pila.Desapilar(), 3)
	require.EqualValues(t, pila.VerTope(), 15)
	require.EqualValues(t, pila.EstaVacia(), false)
	require.EqualValues(t, pila.Desapilar(), 15)
	require.EqualValues(t, pila.EstaVacia(), true)
	//
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}
func TestApilarDesapilarMucho(t *testing.T) {
	t.Log("Hacemos pruebas apilando y desapilando mucho")
	pila := Pila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	require.EqualValues(t, pila.EstaVacia(), false)
	require.EqualValues(t, pila.VerTope(), 1)
	require.EqualValues(t, pila.Desapilar(), 1)
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	//
	pila.Apilar(2)
	require.EqualValues(t, pila.EstaVacia(), false)
	require.EqualValues(t, pila.VerTope(), 2)
	require.EqualValues(t, pila.Desapilar(), 2)
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	//
	pila.Apilar(2)
	pila.Apilar(3)
	pila.Apilar(4)
	require.EqualValues(t, pila.EstaVacia(), false)
	require.EqualValues(t, pila.VerTope(), 4)
	require.EqualValues(t, pila.Desapilar(), 4)
	require.EqualValues(t, pila.EstaVacia(), false)
	require.EqualValues(t, pila.VerTope(), 3)
	require.EqualValues(t, pila.Desapilar(), 3)
	require.EqualValues(t, pila.EstaVacia(), false)
	require.EqualValues(t, pila.VerTope(), 2)
	require.EqualValues(t, pila.Desapilar(), 2)
	//
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func TestApilarDesapilarString(t *testing.T) {
	t.Log("Hacemos pruebas apilando strings en la pila")
	pila := Pila.CrearPilaDinamica[string]()
	pila.Apilar("hola ")
	require.EqualValues(t, pila.EstaVacia(), false)
	require.EqualValues(t, pila.VerTope(), "hola ")
	require.EqualValues(t, pila.EstaVacia(), false)
	//
	pila.Apilar("como")
	require.EqualValues(t, pila.VerTope(), "como")
	require.EqualValues(t, pila.EstaVacia(), false)
	//
	pila.Apilar("estas ?")
	require.EqualValues(t, pila.EstaVacia(), false)
	require.EqualValues(t, pila.VerTope(), "estas ?")
	require.EqualValues(t, pila.Desapilar(), "estas ?")
	require.EqualValues(t, pila.VerTope(), "como")
	require.EqualValues(t, pila.EstaVacia(), false)
	require.EqualValues(t, pila.Desapilar(), "como")
	require.EqualValues(t, pila.VerTope(), "hola ")
	require.EqualValues(t, pila.EstaVacia(), false)
	require.EqualValues(t, pila.Desapilar(), "hola ")
	require.EqualValues(t, pila.EstaVacia(), true)
	//
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}
func TestApilarDesapilarPunteros(t *testing.T, pila Pila) {
	t.Log("Hacemos pruebas apilando *ints en la pila")
	pila := Pila.CrearPilaDinamica[*int]()
	var nro int = 40
	var a *int = &nro
	pila.Apilar(a)
	require.EqualValues(t, pila.EstaVacia(), false)
	require.EqualValues(t, pila.VerTope(), a)
	require.EqualValues(t, pila.EstaVacia(), false)
	//
	var nro2 int = 13
	var b *int = &nro2
	pila.Apilar(b)
	require.EqualValues(t, pila.VerTope(), b)
	require.EqualValues(t, pila.EstaVacia(), false)
	require.EqualValues(t, pila.Desapilar(), b)
	require.EqualValues(t, pila.VerTope(), a)
	require.EqualValues(t, pila.EstaVacia(), false)
	require.EqualValues(t, pila.Desapilar(), a)
	require.EqualValues(t, pila.EstaVacia(), true)
	//
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func TestVariosInts(t *testing.T) {
	t.Log("Hacemos pruebas con varios ints")
	pila := Pila.CrearPilaDinamica[int]()
	nro := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	//
	for i := 0; i < 20; i++ {
		pila.Apilar(nro[i])
		require.EqualValues(t, pila.VerTope(), nro[i])
		require.EqualValues(t, pila.EstaVacia(), false)
	}
	//
	for i := 19; !(pila.EstaVacia()); i-- {
		require.EqualValues(t, pila.Desapilar(), nro[i])
	}
	//
	require.EqualValues(t, pila.EstaVacia(), true)
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func TestVolumen(t *testing.T) {
	t.Log("Hacemos pruebas de volumen")
	// Creo un arreglo de ints para apilar y desapilar
	var cosas [_CONSTANTETESTVOLUMEN]int
	for i := 0; i < _CONSTANTETESTVOLUMEN; i++ {
		cosas[i] = i + 1
	}
	//
	pila := Pila.CrearPilaDinamica[int]()
	//
	for i := 0; i < _CONSTANTETESTVOLUMEN; i++ {
		pila.Apilar(cosas[i])
		require.EqualValues(t, pila.VerTope(), cosas[i])
		require.EqualValues(t, pila.EstaVacia(), false)
	}
	//
	for i := _CONSTANTETESTVOLUMEN - 1; !(pila.EstaVacia()); i-- {
		require.EqualValues(t, pila.Desapilar(), cosas[i])
	}
	require.EqualValues(t, pila.EstaVacia(), true)
	//
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}
