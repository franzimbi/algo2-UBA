package cola

type nodoCola[T any] struct {
	dato    T
	proximo *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func nodoCrear[T any](dato T) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.dato = dato
	return nodo
}
func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	return cola
}

func (cola colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.dato
}

func (cola *colaEnlazada[T]) Encolar(dato T) {
	nuevoNodo := nodoCrear(dato)

	if cola.EstaVacia() {
		(*cola).primero = nuevoNodo
		(*cola).ultimo = nuevoNodo
		return
	}
	(*cola).ultimo.proximo = nuevoNodo
	(*cola).ultimo = nuevoNodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := cola.primero.dato
	if cola.primero == cola.ultimo {
		(*cola).ultimo = nil
	}
	(*cola).primero = cola.primero.proximo
	return dato

}

func (cola *colaEnlazada[T]) Multiprimeros(k int) []T {
	var arr []T
	actual := cola.primero
	for i := 0; i < k; i++ {
		if actual == nil {
			//return arr[:i]
			break
		}
		//arr[i] = actual.dato
		arr = append(arr, actual.dato)
		actual = actual.proximo
	}
	return arr
}

//O(k)

/*Implementar en Go una primitiva func que invierta la cola, sin utilizar estructuras auxiliares. Indicar y justificar el orden de la primitiva. */

func (cola *colaEnlazada[T]) Invertir() {
	if cola.EstaVacia() || cola.primero == cola.ultimo {
		return
	}
	var anterior *nodoCola[T] = (*cola).primero
	var actual *nodoCola[T] = anterior.proximo
	var proximo *nodoCola[T] = actual.proximo

	for anterior != cola.ultimo {
		actual.proximo = anterior
		anterior = actual
		if proximo == nil {
			break
		}
		actual = proximo
		proximo = proximo.proximo
	}
	cola.primero, cola.ultimo = cola.ultimo, cola.primero
	cola.ultimo.proximo = nil
}
