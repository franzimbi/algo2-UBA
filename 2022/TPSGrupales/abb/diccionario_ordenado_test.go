package diccionario_test

import (
	Arbol "diccionario"
	"math/rand"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const _TEST_VOLUMEN = 10000

func cmpInts(a, b int) int {
	return a - b
}

type alumno struct {
	padron int
	nombre string
}

func cmpAlumnos(a, b alumno) int {
	return a.padron - b.padron
}
func generarAbb[K comparable](arr []K, cmp func(K, K) int) Arbol.DiccionarioOrdenado[K, K] {
	arbol := Arbol.CrearABB[K, K](cmp)
	for i := 0; i < len(arr); i++ {
		arbol.Guardar(arr[i], arr[i])
	}
	return arbol
}
func TestDiccionarioVacio(t *testing.T) {
	t.Log("Hacemos pruebas con un abb vacio")
	arbol := Arbol.CrearABB[string, any](strings.Compare)
	require.EqualValues(t, arbol.Cantidad(), 0)
	require.EqualValues(t, arbol.Pertenece("gdfrds4e"), false)
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { arbol.Borrar("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { arbol.Obtener("") })
	require.EqualValues(t, arbol.Pertenece(""), false)
	require.EqualValues(t, arbol.Cantidad(), 0)
}
func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	arbol := Arbol.CrearABB[string, any](strings.Compare)
	arbol.Guardar("A", nil)
	//
	require.EqualValues(t, arbol.Cantidad(), 1)
	require.EqualValues(t, arbol.Pertenece("A"), true)
	require.EqualValues(t, arbol.Pertenece("B"), false)
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { arbol.Obtener("B") })
}
func TestDiccionarioGuardar(t *testing.T) {
	t.Log("Hacemos pruebas guardando un elemento en un abb")
	arbol := Arbol.CrearABB[string, int](strings.Compare)
	arbol.Guardar("martin", 10)
	arbol.Guardar("jorge", 1)
	arbol.Guardar("julian", 2)
	//
	require.EqualValues(t, arbol.Cantidad(), 3)
	require.EqualValues(t, arbol.Pertenece("jorge"), true)
	require.EqualValues(t, arbol.Pertenece("boca"), false)
	require.EqualValues(t, arbol.Obtener("julian"), 2)
	arbol.Guardar("jazmina", 4)
	arbol.Guardar("boca", 12)
	require.EqualValues(t, arbol.Cantidad(), 5)
	require.EqualValues(t, arbol.Pertenece("martin"), true)
	require.EqualValues(t, arbol.Pertenece("julian"), true)
	require.EqualValues(t, arbol.Obtener("boca"), 12)
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { arbol.Obtener("racing") })
	require.EqualValues(t, arbol.Pertenece("river"), false)
	require.EqualValues(t, arbol.Obtener("martin"), 10)
}
func TestReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	arbol := Arbol.CrearABB[int, string](cmpInts)
	arbol.Guardar(6, "liber")
	arbol.Guardar(86, "diego")
	//
	require.EqualValues(t, arbol.Pertenece(6), true)
	require.EqualValues(t, arbol.Pertenece(86), true)
	require.EqualValues(t, arbol.Obtener(6), "liber")
	require.EqualValues(t, arbol.Obtener(86), "diego")
	require.EqualValues(t, arbol.Cantidad(), 2)
	//
	arbol.Guardar(6, "marcos rojo")
	arbol.Guardar(86, "bilardo")
	//
	require.EqualValues(t, arbol.Pertenece(6), true)
	require.EqualValues(t, arbol.Pertenece(86), true)
	require.EqualValues(t, arbol.Obtener(6), "marcos rojo")
	require.EqualValues(t, arbol.Obtener(86), "bilardo")
	require.EqualValues(t, arbol.Cantidad(), 2)
}
func TestDiccionarioBorrar(t *testing.T) {
	t.Log("Hacemos pruebas borrando un elemento en un abb")
	arbol := Arbol.CrearABB[int, string](cmpInts)
	arbol.Guardar(6, "seis")
	arbol.Guardar(3, "tres")
	arbol.Guardar(10, "d10s")
	arbol.Guardar(1, "boca puntero")
	arbol.Guardar(4, "cuatrochi")
	arbol.Guardar(8, "no voy a decir nada...")
	arbol.Guardar(13, "jajaja sh")
	//
	require.EqualValues(t, arbol.Borrar(13), "jajaja sh")
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { arbol.Obtener(13) })
	require.EqualValues(t, arbol.Cantidad(), 6)
	//
	require.EqualValues(t, arbol.Borrar(10), "d10s")
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { arbol.Obtener(10) })
	require.EqualValues(t, arbol.Cantidad(), 5)
	//
	require.EqualValues(t, arbol.Borrar(3), "tres")
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { arbol.Obtener(3) })
	require.EqualValues(t, arbol.Cantidad(), 4)
	//
	require.EqualValues(t, arbol.Borrar(4), "cuatrochi")
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { arbol.Obtener(4) })
	require.EqualValues(t, arbol.Cantidad(), 3)
	//
	require.EqualValues(t, arbol.Borrar(6), "seis")
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { arbol.Obtener(6) })
	require.EqualValues(t, arbol.Cantidad(), 2)
	//
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { arbol.Borrar(10) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { arbol.Borrar(3) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { arbol.Borrar(6) })
	require.EqualValues(t, arbol.Cantidad(), 2)
}
func TestReutilizarBorrados(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	arbol := Arbol.CrearABB[string, int](strings.Compare)
	arbol.Guardar("hola", 3)
	arbol.Guardar("como", 4)
	arbol.Guardar("estas", 0)
	arbol.Guardar("todo", 29)
	arbol.Guardar("bien", 40)

	require.EqualValues(t, arbol.Pertenece("todo"), true)
	require.EqualValues(t, arbol.Obtener("como"), 4)
	require.EqualValues(t, arbol.Cantidad(), 5)
	//
	require.EqualValues(t, arbol.Borrar("como"), 4)
	require.EqualValues(t, arbol.Cantidad(), 4)
	require.EqualValues(t, arbol.Borrar("todo"), 29)
	require.EqualValues(t, arbol.Cantidad(), 3)
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { arbol.Obtener("como") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { arbol.Borrar("todo") })
	//
	arbol.Guardar("como", 33)
	require.EqualValues(t, arbol.Pertenece("como"), true)
	require.EqualValues(t, arbol.Obtener("como"), 33)
	require.EqualValues(t, arbol.Cantidad(), 4)
	//
	arbol.Guardar("todo", -1)
	require.EqualValues(t, arbol.Pertenece("todo"), true)
	require.EqualValues(t, arbol.Obtener("todo"), -1)
	require.EqualValues(t, arbol.Cantidad(), 5)
}
func TestConClavesStruct(t *testing.T) {
	t.Log("Valida que tambien funcione con estructuras mas complejas")
	alumno1 := alumno{padron: 103295, nombre: "francisco"}
	alumno2 := alumno{padron: 112343, nombre: "pepe"}
	alumno3 := alumno{padron: 000001, nombre: "madero"}
	//
	arbol := Arbol.CrearABB[alumno, int](cmpAlumnos)
	arbol.Guardar(alumno1, 10)
	arbol.Guardar(alumno2, 8)
	arbol.Guardar(alumno3, 9)
	//
	require.EqualValues(t, arbol.Pertenece(alumno1), true)
	require.EqualValues(t, arbol.Pertenece(alumno2), true)
	require.EqualValues(t, arbol.Pertenece(alumno3), true)
	require.EqualValues(t, arbol.Cantidad(), 3)
	//
	require.EqualValues(t, arbol.Obtener(alumno1), 10)
	require.EqualValues(t, arbol.Obtener(alumno2), 8)
	require.EqualValues(t, arbol.Obtener(alumno3), 9)
	//
	require.EqualValues(t, arbol.Borrar(alumno1), 10)
	require.EqualValues(t, arbol.Cantidad(), 2)
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { arbol.Obtener(alumno1) })
	require.EqualValues(t, arbol.Obtener(alumno3), 9)
}
func TestValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	arbol := Arbol.CrearABB[string, *int](strings.Compare)
	arbol.Guardar("yo", nil)
	//
	require.EqualValues(t, arbol.Pertenece("yo"), true)
	require.EqualValues(t, arbol.Obtener("yo"), (*int)(nil))
	require.EqualValues(t, arbol.Cantidad(), 1)
	//
	require.EqualValues(t, arbol.Borrar("yo"), (*int)(nil))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { arbol.Obtener("yo") })
}
func TestVolumen(t *testing.T) {
	t.Log("Probamos guardar y borrar", _TEST_VOLUMEN, " elementos")
	// crea ints randoms
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 2147483646
	numeros := new([_TEST_VOLUMEN]int)
	for i := 0; i < _TEST_VOLUMEN; i++ {
		numeros[i] = rand.Intn(max-min+1) + min
	}
	//
	arbol := Arbol.CrearABB[int, int](cmpInts)
	for i := 0; i < _TEST_VOLUMEN; i++ {
		arbol.Guardar(numeros[i], i)
		require.EqualValues(t, arbol.Pertenece(numeros[i]), true)
		require.EqualValues(t, arbol.Obtener(numeros[i]), i)
	}
	//
	for i := 0; i < _TEST_VOLUMEN; i++ {
		arbol.Borrar(numeros[i])
	}
	require.EqualValues(t, arbol.Cantidad(), 0)
}
func TestIteradorExterno(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas con el iterador interno y en orden")
	var arr = []string{"e", "c", "h", "b", "d", "f", "g", "i"}
	arbol := generarAbb[string](arr, strings.Compare)
	//
	iter := arbol.Iterador()
	require.EqualValues(t, iter.Siguiente(), "b")
	auxS, auxI := iter.VerActual()
	require.EqualValues(t, auxS, "c")
	require.EqualValues(t, auxI, "c")
	require.EqualValues(t, iter.Siguiente(), "c")
	auxS, auxI = iter.VerActual()
	require.EqualValues(t, auxS, "d")
	require.EqualValues(t, auxI, "d")
	//
	require.EqualValues(t, iter.Siguiente(), "d")
	require.EqualValues(t, iter.Siguiente(), "e")
	require.EqualValues(t, iter.Siguiente(), "f")
	require.EqualValues(t, iter.Siguiente(), "g")
	require.EqualValues(t, iter.Siguiente(), "h")
	require.EqualValues(t, iter.HaySiguiente(), true)
	require.EqualValues(t, iter.Siguiente(), "i")
	//
	require.EqualValues(t, iter.HaySiguiente(), false)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
func TestIterarDiccionarioVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	arbol := Arbol.CrearABB[string, int](strings.Compare)
	iter := arbol.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
func TestIteradorPorRango(t *testing.T) {
	t.Log("Valida que el iterador en un rango corte")
	var arr = []int{100, 50, 150, 25, 75, 125, 200, 10, 30, 60, 80, 110, 130, 190, 300}
	arbol := generarAbb[int](arr, cmpInts)
	//
	desde := 60
	hasta := 120
	iter := arbol.IteradorRango(&desde, &hasta)
	auxS, auxI := iter.VerActual()
	require.EqualValues(t, auxS, 60)
	require.EqualValues(t, auxI, 60)
	require.EqualValues(t, iter.Siguiente(), 60)
	//
	require.EqualValues(t, iter.Siguiente(), 75)
	require.EqualValues(t, iter.Siguiente(), 80)
	require.EqualValues(t, iter.Siguiente(), 100)
	//
	auxS, auxI = iter.VerActual()
	require.EqualValues(t, auxS, 110)
	require.EqualValues(t, auxI, 110)
	//
	require.EqualValues(t, iter.Siguiente(), 110)
	require.False(t, iter.HaySiguiente())
	//
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	//
	require.EqualValues(t, arbol.Cantidad(), 15)
}
func TestIteradorPorRangoNil(t *testing.T) {
	t.Log("Valida que el iterador en un rango de nil, nil sea igual q el iterador comun")
	var arr = []int{100, 50, 150, 25, 75, 125, 200, 130}
	arbol := generarAbb[int](arr, cmpInts)
	//
	iter := arbol.IteradorRango(nil, nil)
	auxS, auxI := iter.VerActual()
	require.EqualValues(t, auxS, 25)
	require.EqualValues(t, auxI, 25)
	require.EqualValues(t, iter.Siguiente(), 25)
	//
	require.EqualValues(t, iter.Siguiente(), 50)
	require.EqualValues(t, iter.Siguiente(), 75)
	require.EqualValues(t, iter.Siguiente(), 100)
	//
	require.EqualValues(t, iter.Siguiente(), 125)
	require.EqualValues(t, iter.Siguiente(), 130)
	require.EqualValues(t, iter.Siguiente(), 150)
	//
	require.EqualValues(t, iter.Siguiente(), 200)
	require.False(t, iter.HaySiguiente())
	//
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	//
	require.EqualValues(t, arbol.Cantidad(), 8)
}
func TestIteradorInterno(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas con el iterador interno en orden")
	var arr = []string{"e", "c", "h", "b", "d", "f", "g", "i"}
	arbol := generarAbb[string](arr, strings.Compare)
	//
	pos := 0
	arbol.Iterar(func(clave, dato string) bool {
		posibles := []string{"b", "c", "d", "e", "f", "g", "h", "i"}
		if strings.Compare(posibles[pos], clave) != 0 {
			return false
		}
		require.EqualValues(t, strings.Compare(posibles[pos], clave), 0)
		pos++
		return true
	})
	require.EqualValues(t, arbol.Cantidad(), 8)
}
func TestIteradorInternoPorRango(t *testing.T) {
	t.Log("Valida que el iterador interno itere en el rango")
	var arr = []int{100, 50, 150, 25, 75, 125, 200, 10, 30, 60, 80, 110, 130, 190, 300}
	arbol := generarAbb[int](arr, cmpInts)
	//
	desde := 60
	hasta := 120
	arbol.IterarRango(&desde, &hasta, func(clave, dato int) bool {
		require.EqualValues(t, clave >= desde, true)
		require.EqualValues(t, clave <= hasta, true)
		return true
	})
	require.EqualValues(t, arbol.Cantidad(), 15)
}
func TestIteradorInternoCorte(t *testing.T) {
	t.Log("Valida que el iterador interno corte cuando visitar devuelve false")
	var arr = []int{1, 5, 15, 2, 7, 12, 20, 1, 0, 6, 8, 11, 13, 19, 30}
	arbol := generarAbb[int](arr, cmpInts)
	//
	desde := 5
	var ultimoVisitado int
	arbol.IterarRango(&desde, nil, func(clave, dato int) bool {
		ultimoVisitado = clave
		if clave == 15 {
			return false
		}
		require.EqualValues(t, clave >= desde, true)
		return true
	})
	require.EqualValues(t, ultimoVisitado, 15)
}
func TestIteradorSinDesde(t *testing.T) {
	t.Log("Valida que el iterador interno sin desde funcione")
	var arr = []string{"e", "c", "h", "b", "d", "f", "g", "i"}
	arbol := generarAbb[string](arr, strings.Compare)
	//
	pos := 0
	hasta := "g"
	arbol.IterarRango(nil, &hasta, func(clave, dato string) bool {
		posibles := []string{"b", "c", "d", "e", "f", "g", "h", "i"}
		if strings.Compare(posibles[pos], clave) != 0 {
			return false
		}
		require.EqualValues(t, strings.Compare(posibles[pos], clave), 0)
		pos++
		return true
	})
	require.EqualValues(t, pos, 6)
	require.EqualValues(t, arbol.Cantidad(), 8)
}
func TestIteradorSinHasta(t *testing.T) {
	t.Log("Valida que el iterador interno sin hasta funcione")
	var arr = []string{"eee", "cc", "hh", "bbbb", "dd", "fff", "gggg", "iiiii"}
	arbol := generarAbb[string](arr, strings.Compare)
	//
	pos := 2
	desde := "dd"
	posibles := []string{"bbbb", "cc", "dd", "eee", "fff", "gggg", "hh", "iiiii"}
	sort.Strings(posibles)
	arbol.IterarRango(&desde, nil, func(clave, dato string) bool {
		if strings.Compare(posibles[pos], clave) != 0 {
			return false
		}
		require.EqualValues(t, strings.Compare(posibles[pos], clave), 0)
		pos++
		return true
	})
	require.EqualValues(t, pos, 8)
	require.EqualValues(t, arbol.Cantidad(), 8)
}
func TestIterarSinDesde(t *testing.T) {
	t.Log("Valida que el iterador por rango sin desde funcione")
	var arr = []int{47, 20, 80, 13, 51, 200, 150, 23, 18, 69, 82, 59, 27, 60, 432}
	arbol := generarAbb[int](arr, cmpInts)
	//
	hasta := 47
	iter := arbol.IteradorRango(nil, &hasta)
	auxS, auxI := iter.VerActual()
	require.EqualValues(t, auxS, 13)
	require.EqualValues(t, auxI, 13)
	require.EqualValues(t, iter.Siguiente(), 13)
	//
	require.EqualValues(t, iter.Siguiente(), 18)
	require.EqualValues(t, iter.Siguiente(), 20)
	require.EqualValues(t, iter.Siguiente(), 23)
	//
	auxS, auxI = iter.VerActual()
	require.EqualValues(t, auxS, 27)
	require.EqualValues(t, auxI, 27)
	//
	require.EqualValues(t, iter.Siguiente(), 27)
	//
	require.EqualValues(t, iter.Siguiente(), 47)
	//
	require.False(t, iter.HaySiguiente())
	//
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	//
	require.EqualValues(t, arbol.Cantidad(), 15)
}
func TestIterarSinHasta(t *testing.T) {
	t.Log("Valida que el iterador por rango sin hasta funcione")
	var arr = []int{20, 15, 25, 10, 30, 8, 22, 26, 31, 9}
	arbol := generarAbb[int](arr, cmpInts)
	//
	desde := 23
	iter := arbol.IteradorRango(&desde, nil)
	auxS, auxI := iter.VerActual()
	require.EqualValues(t, auxS, 25)
	require.EqualValues(t, auxI, 25)
	require.EqualValues(t, iter.Siguiente(), 25)
	//
	require.EqualValues(t, iter.Siguiente(), 26)
	require.EqualValues(t, iter.Siguiente(), 30)
	require.EqualValues(t, iter.Siguiente(), 31)
	//
	require.False(t, iter.HaySiguiente())
	//
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	//
	require.EqualValues(t, arbol.Cantidad(), 10)
}
func TestIterarSinHastaOrdenado(t *testing.T) {
	t.Log("Valida que el iterador por rango sin hasta funcione")
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	arbol := generarAbb[int](arr, cmpInts)
	//
	//
	desde := 5
	iter := arbol.IteradorRango(&desde, nil)
	auxS, auxI := iter.VerActual()
	require.EqualValues(t, auxS, 5)
	require.EqualValues(t, auxI, 5)
	require.EqualValues(t, iter.Siguiente(), 5)
	//
	require.EqualValues(t, iter.Siguiente(), 6)
	require.EqualValues(t, iter.Siguiente(), 7)
	require.EqualValues(t, iter.Siguiente(), 8)
	//
	auxS, auxI = iter.VerActual()
	require.EqualValues(t, auxS, 9)
	require.EqualValues(t, auxI, 9)
	//
	require.EqualValues(t, iter.Siguiente(), 9)
	//
	require.EqualValues(t, iter.Siguiente(), 10)
	require.EqualValues(t, iter.Siguiente(), 11)
	//
	require.False(t, iter.HaySiguiente())
	//
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	//
	require.EqualValues(t, arbol.Cantidad(), 11)
}
