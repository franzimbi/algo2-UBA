package main

import (
	Heap "ejercicios/colaP"
	"fmt"
	"strings"
)

/*
Implementar en lenguaje Go una función recursiva con la firma func esHeap(arr []int).
Esta función debe devolver true o false de acuerdo a si el arreglo que recibe como
parámetro cumple la propiedad de heap (de mínimos).
*/

func esHeap(arr []int) bool {
	tam := len(arr)
	if tam <= 1 {
		return true
	}
	PosPadre := (tam - 2) / 2
	if arr[PosPadre] <= arr[tam-1] {
		return esHeap(arr[:tam-1])
	} else {
		return false
	}
}

/*
	Se tienen k arreglos de enteros previamente ordenados y se quiere obtener un arreglo

ordenado que contenga a todos los elementos de los k arreglos. Sabiendo que cada
arreglo tiene tamaño h, definimos como n a la sumatoria de la cantidad de elementos de
todos los arreglos, es decir, n=k×h.
Escribir en Go una función func KMerge(arr [][]int) que reciba los k arreglos y devuelva
uno nuevo con los n elementos ordenados entre sí. La función debe ser de orden O(nlogk).
Justificar el orden del algoritmo propuesto.
*/
type tipo struct {
	elem string
	arr  int
	pos  int
}

func cmp(a, b tipo) int {
	return strings.Compare(b.elem, a.elem)
}
func KMerge(arr [][]string) []string {
	heap := Heap.CrearHeap[tipo](cmp)
	cantidadArreglos := len(arr)
	cantidadLetras := len(arr[0])
	var result []string
	for i := 0; i < cantidadArreglos; i++ {
		heap.Encolar(tipo{elem: arr[i][0], arr: i, pos: 0})
	}
	for !heap.EstaVacia() {
		aux := heap.Desencolar()
		result = append(result, aux.elem)
		if aux.pos < cantidadLetras-1 {
			heap.Encolar(tipo{elem: arr[aux.arr][aux.pos+1], arr: aux.arr, pos: aux.pos + 1})
		}
	}
	return result
}
func main() {
	//arr := []int{1, 7, 2, 8, 7, 6, 3, 3, 9, 10}
	//fmt.Println(esHeap(arr))
	st1 := [][]string{{"a", "c"}, {"b", "d"}, {"e", "f"}}
	fmt.Println(KMerge(st1))
}
