package main

import "fmt"

// el minimo de un arreglo por div y conquista
func minimo(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	mitad := len(arr) / 2

	izq := minimo(arr[:mitad])
	der := minimo(arr[mitad:])
	if izq <= der {
		return izq
	} else {
		return der
	}
}

/*
T(n) = AT(n/B) + O(n^C)
A: cantidad llamados recursivos
B: proporcion del tamaño original con el q llamamos recursivamente
C: O(n^c) todo lo q no es recursivo

-> A: 2, B: 2, C: 0 -> log.B(A) > 0 -> O(n)
*/

func estaOrdenado(arr []int) bool {
	tam := len(arr)
	if tam == 1 || tam == 0 {
		return true
	}
	if arr[tam/2-1] >= arr[tam/2] {
		return false
	} else {
		return estaOrdenado(arr[:tam/2]) && estaOrdenado(arr[tam/2:])
	}
}

/*
T(n) = A T(n/B) . O(n^C)
A: cantidad de llamados recursivos -> 2
B: proporcion del tamaño con el q se hacen los llamados recurs -> 2
C: todo lo q no es recursivo -> O(n^C) = O(1) -> C=0

como log b(a)=1 > C=0 -> T(n) = O(n^1) = O(n)
*/

/* implementar, por división y conquista, una función que dado un arreglo sin elementos repetidos y casi ordenado (todos los elementos se encuentran ordenados, salvo uno), obtenga el elemento fuera de lugar. Indicar y justificar el orden. */

func elementoFueraDeLugar(arr []int) int {
	tam := len(arr)

	if tam == 1 {
		return arr[0]
	}
	izq := elementoFueraDeLugar(arr[:tam/2])
	der := elementoFueraDeLugar(arr[tam/2:])
	if izq < der {
		return der
	} else {
		return izq
	}
}

/*
	T(n) = A.T(n/B) + O(n^C)
	A: cantidad de llamados recursivos -> 2
	B: proporcion respecto del tamaño original con el q se hace cada llamado rec -> 2
	C: todo lo q no sea rec O(n^C) -> 0
	log.B(A)=1 > C=0 --> O(n) = O(n^log.B(A)) -> O(n)
*/

/*
Se tiene un arreglo tal que [1, 1, 1, ..., 0, 0, ...] (es decir, unos seguidos de ceros). Se pide:
una función de orden
O(log n),que encuentre el índice del primer 0. Si no hay ningún 0 (solo hay unos), debe devolver -1.
demostrar con el Teorema Maestro que la función es, en efecto,
*/
func primerCero(arr []int, ini, fin int) int {
	if ini == fin {
		return -1
	}
	medio := (ini + fin) / 2
	if arr[medio] == 1 && arr[medio+1] == 0 {
		return medio + 1
	} else if arr[medio] == 0 {
		return primerCero(arr, ini, medio)
	} else {
		return primerCero(arr, medio+1, fin)
	}
}

/*
	T(n) = A. T(n/B) + O(n^C)
	A: cantidad de llamados recursivos -> 1
	B: proporcion del tamaño originial del dato con el q se hace recursividad -> 2
	C: todo lo q no es recursivo -> 0
	log B(A)= 0 = C=0 -> T(n) = O^c . log N -> O(log n)
*/

/* Implementar un algoritmo que, por división y conquista, permita obtener la parte entera de la raíz cuadrada de un número
n, en tiempo O(log n). Por ejemplo, para n=10 debe devolver 3, y para n=25 debe devolver 5. */

func parteEntera(nro int, nroFijo int) int {
	if nro*nro == nroFijo {
		return nro
	}
	raiz := (nro / 2) * (nro / 2)
	if raiz < nroFijo {
		return parteEntera(nro+1, nroFijo)
	}
	return parteEntera(nro/2, nroFijo)
}

func esMagico(arr []int, ini, fin int) bool {
	if ini > fin {
		return false
	}
	medio := ini + fin/2
	if arr[medio] == medio {
		return true
	}
	if arr[medio] < medio {
		return esMagico(arr, medio+1, fin)
	} else {
		return esMagico(arr, ini, medio-1)
	}
}

/*
T(n) = A.T(n) + O(n^c)
A: la cantidad de llamados recursivos -> 1
B:proporcion respecto al tamaño original en las recursividades -> 2
C: lo no recursivo O(1) -> 0

log.B(A)=0 == C=0 -> O(n^c . log n) -> O(log n)
*/
func main() {
	numeros := []int{1, 2, 4, 6, 7, 9}
	fmt.Print(esMagico(numeros, 0, len(numeros)-1), "\n")
}
