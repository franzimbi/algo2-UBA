package usuario

import (
	heap "tp2/heap"
)

type postPrioridad struct {
	prioridad, id int
}
type usuario struct {
	nombre string
	id     int
	feed   heap.ColaPrioridad[postPrioridad]
}

func calculoPrioridad(idLector, idCreador int) int {
	prioridad := idCreador - idLector
	if prioridad < 0 {
		prioridad *= -1
	}
	return prioridad
}
func cmpPost(a, b postPrioridad) int {
	if a.prioridad == b.prioridad {
		return b.id - a.id
	}
	return b.prioridad - a.prioridad
}
func CrearUsuario(nombre string, id int) Usuario {
	return &usuario{nombre: nombre, id: id, feed: heap.CrearHeap(cmpPost)}
}

func (usr *usuario) AgregarPostAlFeed(idPost, idCreador int) {
	post := postPrioridad{prioridad: calculoPrioridad(usr.id, idCreador), id: idPost}
	(*usr).feed.Encolar(post)
}

func (usr *usuario) ProximoEnFeed() int {
	if usr.feed.EstaVacia() {
		return -1
	}
	return (*usr).feed.Desencolar().id
}

func (usr usuario) VerDatosUsuario() (string, int) {
	return usr.nombre, usr.id
}
