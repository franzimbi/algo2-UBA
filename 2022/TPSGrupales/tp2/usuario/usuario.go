package usuario

type Usuario interface {
	// agrega el id del post al feed del usuario.
	// el id del post y del creador deben ser  validos
	AgregarPostAlFeed(idPost, idCreador int)

	// devuelve el id del siguiente post con mayor afinidad y en caso de igual afinidad,
	// devuelve el q se publico primero.
	// si no hay proximo en el feed se devuelve un -1
	ProximoEnFeed() int

	//devuelve el nombre y id de un usuario valido
	VerDatosUsuario() (string, int)
}
