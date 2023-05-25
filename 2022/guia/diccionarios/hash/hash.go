package diccionario

import (
	"fmt"
)

const (
	_CAPACIDAD_INICIAL = 991
	_AUMENTO_TABLA     = 10
	_DISMINUCION_TABLA = 2
)

const (
	VACIO = iota
	OCUPADO
	BORRADO
)

type elemento[K comparable, V any] struct {
	clave  K
	dato   V
	estado int
}

type tablaHash[K comparable, V any] struct {
	tabla                           []elemento[K, V]
	capacidad, acumulados, cantidad uint
}

//acumulados = borrados + ocupados
//cantidad = ocupados

type iteradorHash[K comparable, V any] struct {
	posicion uint
	hash     *tablaHash[K, V]
}

// https://en.wikipedia.org/wiki/Jenkins_hash_function
// jenkins_one_at_a_time_hash traducida a golang
func jenkinsOneAtATime(str []byte) uint {
	i := uint(0)
	hash := uint(0)
	tam := uint(len(str))
	for i != tam {
		hash += uint(str[i])
		i++
		hash += hash << 10
		hash ^= hash >> 6
	}
	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15
	return hash
}
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func funcHashing[K comparable](clave K, largo uint) uint {
	return jenkinsOneAtATime(convertirABytes(clave)) % largo
}

func encontrarPosicion[K comparable, V any](tabla []elemento[K, V], posicion uint, clave K) uint {
	for tabla[posicion].estado != VACIO {
		if tabla[posicion].estado == OCUPADO && tabla[posicion].clave == clave {
			return posicion
		}
		posicion++
		if posicion == uint((len(tabla))) {
			posicion = 0
		}
	}
	return posicion
}
func (diccionario *tablaHash[K, V]) redimensionar(tamanoNuevo uint) {
	tablaNueva := make([]elemento[K, V], tamanoNuevo)
	(*diccionario).acumulados = 0
	(*diccionario).cantidad = 0
	for i := uint(0); i < diccionario.capacidad; i++ {
		if diccionario.tabla[i].estado == OCUPADO {
			posicionValida := encontrarPosicion(tablaNueva, funcHashing(diccionario.tabla[i].clave, tamanoNuevo), diccionario.tabla[i].clave)
			tablaNueva[posicionValida].estado = OCUPADO
			tablaNueva[posicionValida].clave = diccionario.tabla[i].clave
			tablaNueva[posicionValida].dato = diccionario.tabla[i].dato
			(*diccionario).acumulados++
			(*diccionario).cantidad++
		}
	}
	(*diccionario).capacidad = tamanoNuevo
	(*diccionario).tabla = tablaNueva
}

// -------------------------- Funciones del Hash -----------------------------
func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(tablaHash[K, V])
	hash.tabla = make([]elemento[K, V], _CAPACIDAD_INICIAL)
	hash.capacidad = _CAPACIDAD_INICIAL
	return hash
}

func (diccionario *tablaHash[K, V]) Guardar(clave K, dato V) {
	if float64(diccionario.acumulados/diccionario.capacidad) > float64(0.7) {
		diccionario.redimensionar(diccionario.capacidad * _AUMENTO_TABLA)
	}
	posicionValida := encontrarPosicion(diccionario.tabla, funcHashing(clave, diccionario.capacidad), clave)
	(*diccionario).tabla[posicionValida].dato = dato
	if diccionario.tabla[posicionValida].estado == OCUPADO {
		return
	} else {
		(*diccionario).tabla[posicionValida].clave = clave
		diccionario.tabla[posicionValida].estado = OCUPADO
		(*diccionario).acumulados++
		(*diccionario).cantidad++
	}
}

func (diccionario tablaHash[K, V]) Pertenece(clave K) bool {
	posicionValida := encontrarPosicion(diccionario.tabla, funcHashing(clave, diccionario.capacidad), clave)
	return diccionario.tabla[posicionValida].estado == OCUPADO
}

func (diccionario *tablaHash[K, V]) Obtener(clave K) V {
	posicionValida := encontrarPosicion(diccionario.tabla, funcHashing(clave, diccionario.capacidad), clave)
	if diccionario.tabla[posicionValida].estado != OCUPADO {
		panic("La clave no pertenece al diccionario")
	}
	return diccionario.tabla[posicionValida].dato
}

func (diccionario *tablaHash[K, V]) Borrar(clave K) V {
	posicionValida := encontrarPosicion(diccionario.tabla, funcHashing(clave, diccionario.capacidad), clave)
	if diccionario.tabla[posicionValida].estado != OCUPADO {
		panic("La clave no pertenece al diccionario")
	}
	dato := (*diccionario).tabla[posicionValida].dato
	(*diccionario).tabla[posicionValida].estado = BORRADO
	var claveNil K                                        // pq hago esto ? para dejar en limpio el espacio y q en caso de que
	var datoNil V                                         // la clave o el dato sea muy grandes ocupe menos memoria
	(*diccionario).tabla[posicionValida].clave = claveNil //por ejemplo: si la clave es un string "hola",
	(*diccionario).tabla[posicionValida].dato = datoNil   //al borrarse la clave pasara a ser "" y el recolector liberara la vieja clave
	(*diccionario).cantidad--
	if diccionario.acumulados-diccionario.cantidad > diccionario.capacidad/3 {
		diccionario.redimensionar(diccionario.capacidad / _DISMINUCION_TABLA)
	}
	return dato

}

func (diccionario tablaHash[K, V]) Cantidad() int {
	return int(diccionario.cantidad)
}

func (diccionario *tablaHash[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for i := int64(0); i < int64(diccionario.capacidad); i++ {
		if diccionario.tabla[i].estado != OCUPADO {
			continue
		} else {
			if !visitar(diccionario.tabla[i].clave, diccionario.tabla[i].dato) {
				return
			}
		}
	}
}

func (diccionario *tablaHash[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iteradorHash[K, V])
	i := uint(0)
	for ; i < diccionario.capacidad; i++ {
		if diccionario.tabla[i].estado == OCUPADO {
			break
		}
	}
	iter.posicion = i
	iter.hash = diccionario
	return iter
}

func (iterador *iteradorHash[K, V]) HaySiguiente() bool {
	for i := iterador.posicion; i < iterador.hash.capacidad; i++ {
		if iterador.hash.tabla[i].estado == OCUPADO {
			return true
		}
	}
	return false
}

func (iterador iteradorHash[K, V]) VerActual() (K, V) {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iterador.hash.tabla[iterador.posicion].clave, iterador.hash.tabla[iterador.posicion].dato

}

func (iterador *iteradorHash[K, V]) Siguiente() K {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	clave := iterador.hash.tabla[iterador.posicion].clave
	iterador.posicion++
	for iterador.posicion < iterador.hash.capacidad && iterador.hash.tabla[iterador.posicion].estado != OCUPADO {
		iterador.posicion++
	}
	return clave
}
