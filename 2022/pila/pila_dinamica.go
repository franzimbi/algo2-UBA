package pila

const (
	_TAMANO_PILA_INICIAL            = 10
	_CONSTANTE_REDIMESION_DESAPILAR = 4
)

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, _TAMANO_PILA_INICIAL)
	pila.cantidad = 0
	return pila
}
func pilaRedimensionar[T any](pila *pilaDinamica[T], NuevoTam int) {
	pila_nueva := make([]T, NuevoTam) //creo un slice del tamano nuevo
	copy(pila_nueva, pila.datos)      // copio los datos del viejo slice
	pila.datos = pila_nueva           // reemplazo la pila vieja por la nueva        // la capacidad aumento a NuevoTam
}
func (pila pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}
func (pila pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return (pila).datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(dato T) {
	if pila.cantidad == len(pila.datos) {
		pilaRedimensionar(pila, len(pila.datos)*2)
	}
	(*pila).datos[pila.cantidad] = dato
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	var extraido T = pila.VerTope()
	if pila.cantidad*_CONSTANTE_REDIMESION_DESAPILAR <= len(pila.datos) && len(pila.datos) != _TAMANO_PILA_INICIAL {
		if len(pila.datos)/2 <= _TAMANO_PILA_INICIAL {
			pilaRedimensionar(pila, _TAMANO_PILA_INICIAL)
		} else {
			pilaRedimensionar(pila, len(pila.datos)/2)
		}
	}
	pila.cantidad--
	return extraido
}
