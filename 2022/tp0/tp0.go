package tp0

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) == 0 {
		return -1
	}
	var pos int
	for i := range vector {
		if vector[pos] < vector[i] {
			pos = i
		}
	}
	return pos
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	for i := 0; i < len(vector1) && i < len(vector2); i++ {
		var aux int = vector1[i] - vector2[i]
		if aux < 0 {
			return -1
		} else if aux > 0 {
			return 1
		}
	}
	var dif int = len(vector1) - len(vector2)
	if dif == 0 {
		return 0
	}
	if dif < 0 {
		return -1
	}
	return 1
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	n := len(vector)
	for i := n - 1; i > 0; i-- {
		var aux int = Maximo(vector[:i+1])
		Swap(&vector[i], &vector[aux])
	}
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	tam := len(vector)
	if tam == 0 {
		return 0
	}
	return vector[tam-1] + Suma(vector[:tam-1])
}

// EsPalindromo devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA.

func EsPalindromo(cadena string) bool {
	tam := len(cadena)
	if tam == 1 || tam == 0 {
		return true
	}
	if cadena[0] == cadena[tam-1] {
		return EsPalindromo(cadena[1 : tam-1])
	} else {
		return false
	}
}
