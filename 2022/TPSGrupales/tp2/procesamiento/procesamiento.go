package procesamiento

import (
	"bufio"
	"os"
	"strings"
	. "tp2/errores"
)

const (
	_POSICION_COMANDO = 0
	_PRIMERA_POSICION = 1
	_SEGUNDA_POSICION = 2
)

// recibe el archivo txt y devuelve un slice con todos los usuarios.
// en caso de error al leer el archivo devuelve ErrorLeerArchivo. caso contrario, error = nil
func LeerUsuarios(archivoTxt string) ([]string, error) {
	archivo, error := os.Open(archivoTxt)
	if error != nil {
		return nil, &ErrorLeerArchivo{}
	}
	fileScanner := bufio.NewScanner(archivo)
	fileScanner.Split(bufio.ScanLines)
	var usuarios []string
	for fileScanner.Scan() {
		usuarios = append(usuarios, fileScanner.Text())
	}
	archivo.Close()
	return usuarios, nil
}

// lee la linea obtenida del stdin y la porcesa dividiendola entre el comando y sus parametros
func ProcesarEntrada(linea *bufio.Scanner) (string, string) {
	lineaLeida := linea.Text()
	aux := strings.Split(lineaLeida, " ")
	var parametro string
	tam := len(aux)
	for i := _PRIMERA_POSICION; i < tam; i++ {
		if tam != _SEGUNDA_POSICION && i != _PRIMERA_POSICION {
			parametro += " "
		}
		parametro += aux[i]
	}
	return aux[_POSICION_COMANDO], parametro
}
