package post

import (
	"strconv"
	"strings"
	arbol "tp2/abb"
)

type publicacion struct {
	id, usrId      int
	mensaje        string
	usuarioCreador string
	likes          arbol.DiccionarioOrdenado[string, int]
}

func CrearPost(id, usrId int, mensaje, usr string) Post {
	return &publicacion{id: id, usrId: usrId, mensaje: mensaje, usuarioCreador: usr, likes: arbol.CrearABB[string, int](strings.Compare)}
}

func (post *publicacion) LikearPost(usuario string) {
	if !post.likes.Pertenece(usuario) {
		(*post).likes.Guardar(usuario, 1)
	}
}

func (post publicacion) VerPost() string {
	return "Post ID " + strconv.Itoa(post.id) + "\n" + post.usuarioCreador + " dijo: " + post.mensaje + "\nLikes: " + strconv.Itoa(post.likes.Cantidad())
}

func (post publicacion) MostrarLikes() (int, []string) {
	var nombres []string
	post.likes.Iterar(func(clave string, _ int) bool {
		nombres = append(nombres, clave)
		return true
	})
	return post.likes.Cantidad(), nombres
}
