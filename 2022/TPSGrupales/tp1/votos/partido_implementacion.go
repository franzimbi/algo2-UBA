package votos

import (
	"strconv"
)

// votos[0] = votos presidente, votos[1]= votos gobernador, votos[2] = votos intendente
type partidoImplementacion struct {
	partido    string
	postulante [CANT_VOTACION]string
	votos      [CANT_VOTACION]int
}

type partidoEnBlanco struct {
	votos [CANT_VOTACION]int
}

func CrearPartido(nombre string, candidatos [3]string) Partido {
	return &partidoImplementacion{partido: nombre, postulante: candidatos}
}

func CrearVotosEnBlanco() Partido {
	return &partidoEnBlanco{}
}

func (partido *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	(*partido).votos[tipo]++
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	(*partido).votos[tipo]++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	cant := strconv.Itoa(partido.votos[tipo])
	resultado := partido.partido + " - " + partido.postulante[tipo] + ": " + cant + " voto"
	if partido.votos[tipo] != 1 {
		return resultado + "s"
	} else {
		return resultado
	}
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	resultado := "Votos en Blanco: " + strconv.Itoa(blanco.votos[tipo]) + " voto"
	if blanco.votos[tipo] != 1 {
		return resultado + "s"
	}
	return resultado
}
