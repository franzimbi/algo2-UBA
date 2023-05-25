package votos

import (
	Errores "tp1/errores"
	Pila "tp1/votos/pila"
)

type paqueteVoto struct {
	tipo        TipoVoto
	alternativa int
}

type votanteImplementacion struct {
	dni    int
	votos  Pila.Pila[paqueteVoto]
	yaVoto bool
}

func CrearVotante(dni int) Votante {
	return &votanteImplementacion{dni: dni, votos: Pila.CrearPilaDinamica[paqueteVoto](), yaVoto: false}
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante votanteImplementacion) YaVoto() bool {
	return votante.yaVoto
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	if votante.YaVoto() {
		return &Errores.ErrorVotanteFraudulento{Dni: votante.dni}
	}
	paquete := paqueteVoto{tipo: tipo, alternativa: alternativa}
	(*votante).votos.Apilar(paquete)
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	if votante.votos.EstaVacia() {
		return &Errores.ErrorNoHayVotosAnteriores{}
	}
	if votante.YaVoto() {
		return nil
	}
	(*votante).votos.Desapilar()
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	if votante.YaVoto() {
		return Voto{}, &Errores.ErrorVotanteFraudulento{Dni: votante.dni}
	}
	retorno := Voto{VotoPorTipo: [3]int{-1, -1, -1}, Impugnado: false}
	for !votante.votos.EstaVacia() {
		aux := (*votante).votos.Desapilar()
		if aux.alternativa == 0 {
			retorno.Impugnado = true
			break
		}
		if retorno.VotoPorTipo[aux.tipo] == -1 {
			retorno.VotoPorTipo[aux.tipo] = aux.alternativa
		}
	}
	if retorno.VotoPorTipo[PRESIDENTE] == -1 {
		retorno.VotoPorTipo[PRESIDENTE] = 0
	}
	if retorno.VotoPorTipo[GOBERNADOR] == -1 {
		retorno.VotoPorTipo[GOBERNADOR] = 0
	}
	if retorno.VotoPorTipo[INTENDENTE] == -1 {
		retorno.VotoPorTipo[INTENDENTE] = 0
	}
	(*votante).yaVoto = true
	return retorno, nil
}
