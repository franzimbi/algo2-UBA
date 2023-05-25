package main

import (
	//abb "ejercicios/abb"
	"fmt"
	"sort"
)

func cmpInts(a, b int) int {
	return a - b
}
func _mitadRecursiva(arr []int, orden *[]int) {
	tam := len(arr)
	if tam <= 1 {
		return
	}
	*orden = append(*orden, arr[tam/2])
	_mitadRecursiva(arr[tam/2:], orden)
	_mitadRecursiva(arr[:tam/2], orden)
}
func arregloOrdenadoParaABBBalanceado(arr []int) []int {
	sort.Ints(arr)
	fmt.Println(arr)
	var orden []int
	_mitadRecursiva(arr, &orden)
	return orden
}
func main() {
	arr := []int{4, 5, 6, 1, 2, 0, 3}
	res := arregloOrdenadoParaABBBalanceado(arr)
	fmt.Println(res)
}
