package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	. "tp2/errores"
	"tp2/procesamiento"
	. "tp2/sistema"
)

type funcComando func(Algogram, string)

const (
	_CANTIDAD_PARAMETROS_NECESARIOS_PROGRAMA = 2
	_PRIMER_PARAMETRO                        = 0
)

// - - - - - - - - - - - - - - - - - - - - - - - COMANDOS - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
func login(a Algogram, nombreLeido string) {
	a.Login(nombreLeido)
}
func logout(a Algogram, _ string) {
	a.Logout()
}
func publicar(a Algogram, msj string) {
	a.Publicar(msj)
}
func verSiguienteFeed(a Algogram, _ string) {
	a.VerSiguienteFeed()
}
func likearPost(a Algogram, idStr string) {
	idPost, err := strconv.Atoi(idStr)
	if err != nil {
		idPost = -1
	}
	a.LikearPost(idPost)
}
func mostrarLikes(a Algogram, idStr string) {
	idPost, err := strconv.Atoi(idStr)
	if err != nil {
		idPost = -1
	}
	a.MostrarLikes(idPost)
}

// - - - - - - - - - - - - - - - - - - - FUNCION MAIN - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
func main() {
	if len(os.Args) != _CANTIDAD_PARAMETROS_NECESARIOS_PROGRAMA { // chequeo que el programa empiece con 2 parametros. ./programa archivoUsuarios.txt
		fmt.Println(ErrorLeerArchivo{}.Error())
		return
	}
	nombresUsuarios, error := procesamiento.LeerUsuarios(os.Args[_PRIMER_PARAMETRO+1]) // leo archivo txt de usuarios
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	algogram := InicioAlgogram((nombresUsuarios))
	var comandos map[string]funcComando = map[string]funcComando{ // creo un mapa de los comandos del programa
		"login":              login,
		"logout":             logout,
		"publicar":           publicar,
		"ver_siguiente_feed": verSiguienteFeed,
		"likear_post":        likearPost,
		"mostrar_likes":      mostrarLikes}

	entradaStdin := bufio.NewScanner(os.Stdin) // abro un scanner para el input
	for entradaStdin.Scan() {                  // mientras lo leido no sea EOF corre
		comando, parametro := procesamiento.ProcesarEntrada(entradaStdin) // procesa la linea leida en itdin
		if funcionLeida, ok := comandos[comando]; ok {
			funcionLeida(algogram, parametro)
		} else {
			fmt.Println(ErrorComando{}.Error())
		}
	}
}
