package lista

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}

type iterador[T any] struct {
	anterior *nodo[T]
	actual   *nodo[T]
	lista    *listaEnlazada[T]
}

func nodoCrear[T any](elem T) *nodo[T] {
	return &nodo[T]{dato: elem}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elem T) {
	nuevoNodo := nodoCrear(elem)
	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	} else {
		nuevoNodo.siguiente = lista.primero
	}
	lista.primero = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elem T) {
	nuevoNodo := nodoCrear(elem)
	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}
	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	lista.largo--
	dato := lista.primero.dato
	lista.primero = lista.primero.siguiente
	return dato
}

func (lista listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	nodoAux := lista.primero
	for nodoAux != nil {
		if !visitar(nodoAux.dato) {
			return
		}
		nodoAux = nodoAux.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterador[T]{actual: lista.primero, lista: lista}
}
func (iter iterador[T]) HaySiguiente() bool {
	return iter.actual != nil
}
func (iter iterador[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter *iterador[T]) Siguiente() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.anterior, iter.actual = iter.actual, iter.actual.siguiente
	return iter.anterior.dato
}

func (iter *iterador[T]) Insertar(dato T) {
	if iter.anterior == nil {
		iter.lista.InsertarPrimero(dato)
		iter.actual = iter.lista.primero
		return
	}
	if !iter.HaySiguiente() {
		iter.lista.InsertarUltimo(dato)
		iter.actual = iter.lista.ultimo
		return
	}
	nuevo := nodoCrear(dato)
	nuevo.siguiente = iter.actual
	iter.anterior.siguiente = nuevo
	iter.actual = nuevo
	iter.lista.largo++
}
func (iter *iterador[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	dato := iter.actual.dato
	if iter.anterior != nil {
		iter.anterior.siguiente = iter.actual.siguiente
	} else {
		iter.lista.primero = iter.lista.primero.siguiente
	}
	iter.actual = iter.actual.siguiente
	if !iter.HaySiguiente() {
		iter.lista.ultimo = iter.anterior
	}
	iter.lista.largo--
	return dato
}
