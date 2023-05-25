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
