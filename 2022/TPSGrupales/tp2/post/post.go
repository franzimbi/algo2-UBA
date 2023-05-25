package post

type Post interface {
	// agrega el like al post del usuario recibido,
	// si este ya habia likeado no lo vuelve a likear
	LikearPost(usuario string)

	// devuelve un string con "post id, mensaje, usuario, ..."
	VerPost() string

	// devuelve la cantidad de likes y un []string con los q likearon en orden alfabetico
	MostrarLikes() (int, []string)
}
