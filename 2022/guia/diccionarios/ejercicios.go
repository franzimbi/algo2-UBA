package main

import (
	TDAhash "ejercicios/hash"
	"fmt"
)

/*
Implementar una función de orden O(n)que dado un arreglo de n números enteros
devuelva true o false según si existe algún elemento que aparezca más de la mitad
de las veces. Justificar el orden de la solución. Ejemplos:

[1, 2, 1, 2, 3] -> false
[1, 1, 2, 3] -> false
[1, 2, 3, 1, 1, 1] -> true
[1] -> true
*/
func elementoMasDeLaMitad(arr []int) bool {
	hash := TDAhash.CrearHash[int, int]()
	mitad := len(arr) / 2
	for i := 0; i < len(arr); i++ {
		if hash.Pertenece(arr[i]) {
			aux := hash.Obtener(arr[i])
			hash.Guardar(arr[i], aux+1)
		} else {
			hash.Guardar(arr[i], 1)
		}
		if hash.Obtener(arr[i]) > mitad {
			return true
		}
	}
	return false
}

/*func mismoDiccionario(a, b TDAhash.Diccionario[string, int]) bool {
	if a.Cantidad() != b.Cantidad() {
		return false
	}
	for i := a.Iterador(); i.HaySiguiente(); i.Siguiente() {
		c, d := i.VerActual()
		if !b.Pertenece(c) {
			return false
		}
		if b.Obtener(c) != d {
			return false
		}
	}
	return true
}*/

/*
Implementar el TDA MultiConjunto. Este es un Conjunto que permite más de una aparición
de un elemento, por lo que eliminando una aparición, el elemento puede seguir perteneciendo.
Dicho TDA debe tener como primitivas:
CrearMulticonj[K](): crea un multiconjunto.
Guardar(elem K): guarda un elemento en el multiconjunto.
Pertence(elem K) bool: devuelve true si el elemento aparece al menos una vez en el
conjunto.
Borrar(elem K) bool: elimina una aparición del elemento dentro del conjunto.
Devuelve true si se eliminó una aparición del elemento.
Dar la estructura del TDA y la implementación de las 4 primitivas marcadas,
de forma tal que todas sean O(1).
*/
func CrearMulticonj[K comparable]() TDAhash.Diccionario[K, int] {
	return TDAhash.CrearHash[K, int]()
}

func MultiGuardar[K comparable](dic TDAhash.Diccionario[K, int], elem K) {
	if dic.Pertenece(elem) {
		dic.Guardar(elem, dic.Obtener(elem)+1)
	} else {
		dic.Guardar(elem, 1)
	}
}
func MultiPertenece[K comparable](dic TDAhash.Diccionario[K, int], elem K) bool {
	return dic.Pertenece(elem)
}
func MultiBorrar[K comparable](dic TDAhash.Diccionario[K, int], elem k) bool {
	if !dic.Pertenece(elem) {
		return false
	}
	cant := dic.Obtener(elem)
	if cant == 1 {
		dic.Borrar(elem)
	} else {
		dic.Guardar(elem, cant-1)
	}
	return true
}

func main() {
	var arr = []int{1, 1, 2, 3}
	fmt.Print(elementoMasDeLaMitad(arr), "\n")

}
