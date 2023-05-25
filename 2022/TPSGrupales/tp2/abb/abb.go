package diccionario

import (
	p "tp2/abb/pila"
)

type nodoAbb[K comparable, V any] struct {
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
	clave K
	dato  V
}
type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}
type iterAbb[K comparable, V any] struct {
	pila         p.Pila[nodoAbb[K, V]]
	arbol        abb[K, V]
	desde, hasta *K
}

func encontrarNodo[K comparable, V any](raiz **nodoAbb[K, V], clave K, cmp func(K, K) int) **nodoAbb[K, V] {
	if *raiz == nil || cmp((*raiz).clave, clave) == 0 {
		return raiz
	}
	if cmp((*raiz).clave, clave) > 0 {
		return encontrarNodo(&(*raiz).izq, clave, cmp)
	} else {
		return encontrarNodo(&(*raiz).der, clave, cmp)
	}
}
func reemplazoTodoDerecha[K comparable, V any](raiz **nodoAbb[K, V]) *nodoAbb[K, V] {
	if (*raiz).der == nil {
		aux := *raiz
		(*raiz) = (*raiz).izq
		return aux
	}
	return reemplazoTodoDerecha(&(*raiz).der)
}

// - - - - - - - - - - - - - - - - - - PRIMITIVAS ABB - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{cmp: funcion_cmp}
}
func (arbol *abb[K, V]) Guardar(clave K, dato V) {
	posicion := encontrarNodo(&(arbol).raiz, clave, arbol.cmp)
	if (*posicion) == nil {
		*posicion = &nodoAbb[K, V]{clave: clave, dato: dato}
		arbol.cantidad++
		return
	} else {
		(*posicion).dato = dato
	}
}
func (arbol *abb[K, V]) Pertenece(clave K) bool {
	posicion := encontrarNodo(&(arbol).raiz, clave, arbol.cmp)
	return (*posicion) != nil
}
func (arbol *abb[K, V]) Obtener(clave K) V {
	posicion := encontrarNodo(&(arbol).raiz, clave, arbol.cmp)
	if (*posicion) == nil {
		panic("La clave no pertenece al diccionario")
	} else {
		return ((*posicion).dato)
	}
}
func (arbol *abb[K, V]) Cantidad() int {
	return arbol.cantidad
}
func (arbol *abb[K, V]) Borrar(clave K) V {
	posicion := encontrarNodo(&(arbol).raiz, clave, arbol.cmp)
	if (*posicion) == nil {
		panic("La clave no pertenece al diccionario")
	}
	dato := (*posicion).dato
	arbol.cantidad--
	if (*posicion).izq == nil {
		(*posicion) = (*posicion).der
		return dato
	}
	if (*posicion).der == nil {
		(*posicion) = (*posicion).izq
		return dato
	} else {
		reemplazo := reemplazoTodoDerecha(&(*posicion).izq)
		izq := (*posicion).izq
		der := (*posicion).der
		(*posicion) = reemplazo
		if reemplazo != izq {
			(*reemplazo).izq = izq
		}
		if reemplazo != der {
			(*reemplazo).der = der
		}
		return dato
	}
}

// ---------- iteradores --------------------------

func iterarRango[K comparable, V any](cmp func(K, K) int, raiz *nodoAbb[K, V], desde, hasta *K, visitar func(clave K, dato V) bool) bool {
	if raiz == nil {
		return true
	}
	cmpDesde := 1
	cmpHasta := -1
	if desde != nil {
		cmpDesde = cmp(raiz.clave, *desde)
	}
	if hasta != nil {
		cmpHasta = cmp(raiz.clave, *hasta)
	}
	if cmpDesde > 0 {
		if !iterarRango(cmp, raiz.izq, desde, hasta, visitar) {
			return false
		}
	}
	if cmpDesde >= 0 && cmpHasta <= 0 {
		if !visitar(raiz.clave, raiz.dato) {
			return false
		}
	}
	if cmpHasta < 0 {
		if !iterarRango(cmp, raiz.der, desde, hasta, visitar) {
			return false
		}
	}
	return true
}
func (arbol *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	iterarRango(arbol.cmp, arbol.raiz, desde, hasta, visitar)
}
func apiloRaizEIzq[K comparable, V any](raiz *nodoAbb[K, V], iter *iterAbb[K, V], cmp func(K, K) int, desde *K) {
	if raiz == nil {
		return
	}
	if desde != nil {
		cmpDesde := cmp(raiz.clave, *desde)
		if cmpDesde == 0 {
			(*iter).pila.Apilar(*raiz)
		} else if cmpDesde > 0 {
			(*iter).pila.Apilar(*raiz)
			apiloRaizEIzq(raiz.izq, iter, cmp, desde)
		} else {
			apiloRaizEIzq(raiz.der, iter, cmp, desde)
		}
	} else {
		(*iter).pila.Apilar(*raiz)
		apiloRaizEIzq(raiz.izq, iter, cmp, desde)
	}
}
func (arbol *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := &iterAbb[K, V]{pila: p.CrearPilaDinamica[nodoAbb[K, V]](), arbol: *arbol, desde: desde, hasta: hasta}
	apiloRaizEIzq(arbol.raiz, iter, arbol.cmp, desde)
	return iter
}
func (iter iterAbb[K, V]) HaySiguiente() bool {
	if iter.pila.EstaVacia() {
		return false
	}
	if iter.hasta != nil {
		if iter.arbol.cmp(iter.pila.VerTope().clave, *iter.hasta) > 0 {
			return false
		}
	}
	return true
}
func (iter iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	return iter.pila.VerTope().clave, iter.pila.VerTope().dato
}

func (iter *iterAbb[K, V]) Siguiente() K {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	anterior := (*iter).pila.Desapilar()
	apiloRaizEIzq(anterior.der, iter, iter.arbol.cmp, iter.desde)
	return anterior.clave
}
func (arbol *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	arbol.IterarRango(nil, nil, visitar)
}
func (arbol *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return arbol.IteradorRango(nil, nil)
}
