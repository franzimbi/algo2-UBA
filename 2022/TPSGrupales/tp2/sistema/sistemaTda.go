package sistema

import (
	"fmt"
	. "tp2/errores"
	. "tp2/hash"
	. "tp2/post"
	. "tp2/usuario"
)

const (
	_MENSAJE_HOLA                   = "Hola "
	_MENSAJE_ADIOS                  = "Adios"
	_MENSAJE_POST_PUBLICADO         = "Post publicado"
	_MENSAJE_NO_MAS_POSTS           = "Usuario no loggeado o no hay mas posts para ver"
	_MENSAJE_POST_LIKEADO           = "Post likeado"
	_MENSAJE_CANTIDAD_LIKES_1_PARTE = "El post tiene"
	_MENSAJE_CANTIDAD_LIKES_2_PARTE = "likes:"
)

type algogram struct {
	diccionarioUsuarios Diccionario[string, int]
	usuarios            []Usuario
	publicaciones       []Post
	usuarioActual       Usuario
	idPostGlobal        int
}

// crea los usuarios y los devuelve en un slice, donde su posicion coincide con su id
func crearUsuarios(usuarios []string) []Usuario {
	tam := len(usuarios)
	sliceUsr := make([]Usuario, tam)
	for i := 0; i < tam; i++ {
		sliceUsr[i] = CrearUsuario(usuarios[i], i)
	}
	return sliceUsr
}

// crea un diccionario del tipo K con las posiciones de cada elemento en el slice recibido
func crearDiccionario(usuarios []string) Diccionario[string, int] {
	tam := len(usuarios)
	dic := CrearHash[string, int]()
	for i := 0; i < tam; i++ {
		dic.Guardar(usuarios[i], i)
	}
	return dic
}
func obtenerUsuario(dic Diccionario[string, int], usuarios []Usuario, nombre string) (Usuario, error) {
	if dic.Pertenece(nombre) {
		return usuarios[dic.Obtener(nombre)], nil
	}
	return nil, &ErrorUsuarioInexistente{}
}

// Obtiene la publicacion siguiente del post
func obtenerPublicacion(usuario Usuario, publicaciones []Post) (string, error) {
	if usuario == nil {
		return "", &ErrorPost{}
	}
	pos := usuario.ProximoEnFeed()
	if pos == -1 {
		return "", &ErrorPost{}
	}
	return publicaciones[pos].VerPost(), nil
}

// Chequea que el post exista y devuelve true.
// caso contrario devuelve false
func publicacionExiste(cantidad, nroId int) bool {
	if nroId < 0 || nroId >= cantidad {
		return false
	}
	return true
}
func InicioAlgogram(nombresUsuarios []string) Algogram {
	return &algogram{diccionarioUsuarios: crearDiccionario(nombresUsuarios),
		usuarios: crearUsuarios((nombresUsuarios)), usuarioActual: nil, idPostGlobal: 0}
}

func (a *algogram) Login(nombreLeido string) {
	if (*a).usuarioActual != nil {
		fmt.Println(ErrorUsuarioLoggeado{}.Error())
		return
	}
	var e error
	(*a).usuarioActual, e = obtenerUsuario(a.diccionarioUsuarios, a.usuarios, nombreLeido)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(_MENSAJE_HOLA + nombreLeido)
}
func (a *algogram) Logout() {
	if (*a).usuarioActual == nil {
		fmt.Println(ErrorUsuarioNoLoggeado{}.Error())
		return
	}
	(*a).usuarioActual = nil
	fmt.Println(_MENSAJE_ADIOS)
}
func (a *algogram) Publicar(mensaje string) {
	if (*a).usuarioActual == nil {
		fmt.Println(ErrorUsuarioNoLoggeado{}.Error())
		return
	}
	nombre, idUsr := ((*a).usuarioActual).VerDatosUsuario()
	nuevoPost := CrearPost((*a).idPostGlobal, idUsr, mensaje, nombre)
	(*a).publicaciones = append((*a).publicaciones, nuevoPost)
	tam := len(a.usuarios)
	for i := 0; i < tam; i++ {
		if i == idUsr {
			continue
		}
		(*a).usuarios[i].AgregarPostAlFeed((*a).idPostGlobal, idUsr)
	}
	(*a).idPostGlobal++
	fmt.Println(_MENSAJE_POST_PUBLICADO)
}
func (a *algogram) VerSiguienteFeed() {
	msj, error := obtenerPublicacion((*a).usuarioActual, (*a).publicaciones)
	if error != nil {
		fmt.Println(_MENSAJE_NO_MAS_POSTS)
		return
	}
	fmt.Println(msj)
}
func (a *algogram) LikearPost(idPost int) {
	if (*a).usuarioActual == nil || !publicacionExiste(len(a.publicaciones), idPost) {
		fmt.Println(ErrorPost{}.Error())
		return
	}
	nombre, _ := (*a).usuarioActual.VerDatosUsuario()
	(*a).publicaciones[idPost].LikearPost(nombre)
	fmt.Println(_MENSAJE_POST_LIKEADO)
}
func (a algogram) MostrarLikes(idPost int) {
	if !publicacionExiste(len(a.publicaciones), idPost) {
		fmt.Println(ErrorMostrarLikes{}.Error())
		return
	}
	cantidad, likes := a.publicaciones[idPost].MostrarLikes()
	if cantidad == 0 { // si no tiene likes
		fmt.Println(ErrorMostrarLikes{}.Error())
		return
	}
	fmt.Printf("%s %d %s\n", _MENSAJE_CANTIDAD_LIKES_1_PARTE, cantidad, _MENSAJE_CANTIDAD_LIKES_2_PARTE)
	for i := 0; i < len(likes); i++ {
		fmt.Printf("\t%s\n", likes[i])
	}
}
