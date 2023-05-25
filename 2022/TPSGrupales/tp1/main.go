package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	Cola "tp1/cola"
	Errores "tp1/errores"
	Procesamiento "tp1/procesamiento"
	. "tp1/votos" // tip de Dato
)

const (
	_PARAMETROS_NECESARIOS_PROGRAMA = 3
	_PARAMETROS_NECESARIOS_INGRESAR = 2
	_PARAMETROS_NECESARIOS_VOTAR    = 3
	_PRIMER_PARAMETRO               = 1
	_SEGUNDO_PARAMETRO              = 2
	_POSICION_PARTIDOS              = 1
	_POSICION_PADRON                = 2
	_POSICION_COMANDO               = 0
	_COMANDO_VOTAR                  = "votar"
	_COMANDO_INGRESAR               = "ingresar"
	_COMANDO_DESHACER               = "deshacer"
	_COMANDO_FIN_VOTAR              = "fin-votar"
)

func validarParametros(args []string, argv int) ([]Partido, []Votante, error) {
	if argv != _PARAMETROS_NECESARIOS_PROGRAMA { //chequeo q haya suficientes parametros
		return nil, nil, Errores.ErrorParametros{}
	}
	partidos, err := Procesamiento.LeerCsvPartidos(args[_POSICION_PARTIDOS]) //lee el csv de los partidos y lo guarda en un slice de partidos, donde [0] es partido en blanco y el resto partidos comunes
	if err != nil {
		return nil, nil, err
	}
	padron, err := Procesamiento.LeerTxtPadron(args[_POSICION_PADRON]) //lee el .txt de dnis y devuelve en un slice ordenado de votantes
	if err != nil {
		return nil, nil, err
	}
	return partidos, padron, nil
}

func ingresar(palabras []string, partidos []Partido, padron []Votante, colaVotacion Cola.Cola[Votante]) {
	if len(palabras) != _PARAMETROS_NECESARIOS_INGRESAR {
		fmt.Println(Errores.ErrorParametros{}.Error())
		return
	}
	pos, err := Procesamiento.PadronValidar(padron, palabras[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	colaVotacion.Encolar(padron[pos])
	fmt.Println("OK")
}

func votar(palabras []string, partidos []Partido, padron []Votante, colaVotacion Cola.Cola[Votante]) {
	if len(palabras) != _PARAMETROS_NECESARIOS_VOTAR {
		fmt.Println(Errores.ErrorParametros{}.Error())
		return
	}
	if colaVotacion.EstaVacia() {
		fmt.Println(Errores.FilaVacia{}.Error())
		return
	}
	tipoVoto, listaVotada, err := Procesamiento.VotoValidar(len(partidos), palabras[_PRIMER_PARAMETRO], palabras[_SEGUNDO_PARAMETRO])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	votante := colaVotacion.VerPrimero()
	err = votante.Votar(tipoVoto, listaVotada)
	if err != nil {
		fmt.Println(err.Error())
		colaVotacion.Desencolar()
		return
	}
	fmt.Println("OK")
}

func deshacer(colaVotacion Cola.Cola[Votante]) {

	if colaVotacion.EstaVacia() {
		fmt.Println(Errores.FilaVacia{}.Error())
		return
	}
	votante := colaVotacion.VerPrimero()
	if votante.YaVoto() {
		colaVotacion.Desencolar()
		fmt.Println(Errores.ErrorVotanteFraudulento{Dni: votante.LeerDNI()}.Error())
		return
	}
	err := votante.Deshacer()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("OK")
}

func finVotar(partidos []Partido, colaVotacion Cola.Cola[Votante], impugnados *int) {
	if colaVotacion.EstaVacia() {
		fmt.Println(Errores.FilaVacia{}.Error())
		return
	}
	votante := colaVotacion.VerPrimero()
	voto, err := votante.FinVoto()
	if err != nil {
		fmt.Println(err.Error())
		colaVotacion.Desencolar()
		return
	}
	if voto.Impugnado {
		(*impugnados)++
	} else {
		partidos[voto.VotoPorTipo[PRESIDENTE]].VotadoPara(PRESIDENTE)
		partidos[voto.VotoPorTipo[GOBERNADOR]].VotadoPara(GOBERNADOR)
		partidos[voto.VotoPorTipo[INTENDENTE]].VotadoPara(INTENDENTE)
		//hacer esto me dio un acv
	}
	colaVotacion.Desencolar()
	fmt.Println("OK")
}

func imprimirResultados(partidos []Partido, padron []Votante, colaVotacion Cola.Cola[Votante], impugnados int) {
	if !colaVotacion.EstaVacia() {
		fmt.Println(Errores.ErrorCiudadanosSinVotar{}.Error())
	}
	fmt.Printf("Presidente:\n")
	for i := 0; i < len(partidos); i++ {
		fmt.Println(partidos[i].ObtenerResultado(PRESIDENTE))
	}
	fmt.Printf("\nGobernador:\n")
	for i := 0; i < len(partidos); i++ {
		fmt.Println(partidos[i].ObtenerResultado(GOBERNADOR))
	}
	fmt.Printf("\nIntendente:\n")
	for i := 0; i < len(partidos); i++ {
		fmt.Println(partidos[i].ObtenerResultado(INTENDENTE))
	}
	resultadoImpugnados := "\nVotos Impugnados: " + strconv.Itoa(impugnados) + " voto"
	if impugnados != 1 {
		fmt.Println(resultadoImpugnados + "s")
	} else {
		fmt.Println(resultadoImpugnados)
	}
}
func main() {
	partidos, padron, err := validarParametros(os.Args, len(os.Args))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	colaVotacion := Cola.CrearColaEnlazada[Votante]() //cola de espera para votar
	entradaStdin := bufio.NewScanner(os.Stdin)
	impugnados := 0
	for entradaStdin.Scan() { //mientras q stdin este abierta, leo lineas de comandos
		palabras := Procesamiento.ProcesarLinea(entradaStdin)
		switch palabras[_POSICION_COMANDO] {
		case _COMANDO_INGRESAR:
			ingresar(palabras, partidos, padron, colaVotacion)
		case _COMANDO_VOTAR:
			votar(palabras, partidos, padron, colaVotacion)
		case _COMANDO_DESHACER:
			deshacer(colaVotacion)
		case _COMANDO_FIN_VOTAR:
			finVotar(partidos, colaVotacion, &impugnados)
		default:
			fmt.Println(Errores.ErrorParametros{}.Error())
		}
	}
	imprimirResultados(partidos, padron, colaVotacion, impugnados) //se llega aca cuando se cierra stdin
}
