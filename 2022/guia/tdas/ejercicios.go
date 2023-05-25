package main

import (
	Cola "ejercicios/cola"
	Pila "ejercicios/pila"
	"fmt"
)

func pilaOrdenada(pila Pila.Pila[int]) bool {
	auxiliarPila := Pila.CrearPilaDinamica[int]()
	auxiliarPila.Apilar(pila.Desapilar())
	for !pila.EstaVacia() { // O(n)
		if pila.VerTope() > auxiliarPila.VerTope() {
			for !auxiliarPila.EstaVacia() {
				pila.Apilar(auxiliarPila.Desapilar())
			}
			return false
		}
		auxiliarPila.Apilar(pila.Desapilar())
	}
	for !auxiliarPila.EstaVacia() {
		pila.Apilar(auxiliarPila.Desapilar()) //o(n)
	}
	return true
}

//O(n)

func Multiprimeros[T any](cola Cola.Cola[T], k int) []T {
	colaAux := Cola.CrearColaEnlazada[T]()
	var arr []T
	for i := 0; !cola.EstaVacia(); i++ {
		if i < k {
			arr = append(arr, cola.VerPrimero())
		}
		colaAux.Encolar(cola.Desencolar())
	}
	for !colaAux.EstaVacia() {
		cola.Encolar(colaAux.Desencolar())
	}
	return arr
}

// O(n)

func ordenarAscendente(pila Pila.Pila[int]) {
	pila_aux := Pila.CrearPilaDinamica[int]()
	aux := pila.Desapilar()
	var orden bool = true
	for orden { //O(n)
		if pila.VerTope() > aux {
			pila_aux.Apilar(pila.Desapilar()) //O(1)
		} else {
			pila_aux.Apilar(aux)
			aux = pila.Desapilar() //O(1)
		}
		if pila.EstaVacia() {
			orden = false
			for !pila_aux.EstaVacia() { //O(n)
				if aux > pila_aux.VerTope() {
					orden = true
				}
				pila.Apilar(aux)
				aux = pila_aux.Desapilar()
			}
		}
	}
	pila.Apilar(aux)
}

//O(n^2)

func MergePilas(pila1, pila2 Pila.Pila[int]) []int {
	var arr []int
	for !pila1.EstaVacia() && !pila2.EstaVacia() { //O(n)+O(m)
		if pila1.VerTope() < pila1.VerTope() {
			arr = append(arr, pila1.Desapilar())
		} else {
			arr = append(arr, pila2.Desapilar())
		}
	}
	for !pila1.EstaVacia() { //O(n)
		arr = append(arr, pila1.Desapilar())
	}
	for !pila2.EstaVacia() { //(m)
		arr = append(arr, pila2.Desapilar())
	}
	//fmt.Print(arr) O(n+m)
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return arr
}

// O(n+m)
func main() {
	pila := Pila.CrearPilaDinamica[int]()
	pila2 := Pila.CrearPilaDinamica[int]()
	pila2.Apilar(4)
	pila2.Apilar(6)
	pila2.Apilar(7)
	pila2.Apilar(8)
	pila2.Apilar(9)
	pila2.Apilar(10)
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(4)
	pila.Apilar(4)
	//fmt.Print(pilaOrdenada(pila), "\n")
	//fmt.Print(pila.VerTope(), "\n")
	/*cola := Cola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Encolar(21)
	cola.Encolar(3232)
	cola.Encolar(23)
	cola.Encolar(-11)
	cola.Encolar(25)
	cola.Encolar(13)
	cola.Invertir()
	cola.Invertir()
	fmt.Print(Multiprimeros(cola, 7), "\n") */
	//ordenarAscendente(pila)
	fmt.Print(MergePilas(pila, pila2))
}
