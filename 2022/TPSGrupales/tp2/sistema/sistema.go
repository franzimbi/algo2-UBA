package sistema

type Algogram interface {
	Login(nombreLeido string)
	Logout()
	Publicar(mensaje string)
	VerSiguienteFeed()
	LikearPost(idPost int)
	MostrarLikes(idPost int)
}
