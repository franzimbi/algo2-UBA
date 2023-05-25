#!/usr/bin/python3
import procesamiento
import comandos
import sys

_CANTIDAD_PARAMETROS_INICIO_PROGRAMA = 2
_PARAMETRO_PAJEK_MAPA = 1
_ESPACIOS = ' '
_STR_VACIO = ''
_CANTIDAD_PARAMETROS_NECESARIOS = 3
_PRIMER_PARAMETRO = 0

if __name__ == "__main__":
    if len(sys.argv) != _CANTIDAD_PARAMETROS_INICIO_PROGRAMA:
        print("Sin archivo pajek")
        sys.exit(1)

    mapa = procesamiento.CrearGrafoConPajek(sys.argv[_PARAMETRO_PAJEK_MAPA])
    sys.setrecursionlimit(100000000)

    diccComandos = {"ir": comandos.Ir, "itinerario": comandos.Itinerario,
                    "viaje": comandos.Viaje, "reducir_caminos": comandos.reducirCaminos}
    diccParametros = {"ir": 3, "itinerario": 1,
                      "viaje": 2, "reducir_caminos": 1}
    for line in sys.stdin:
        line = line.replace('\n', _STR_VACIO)  # elimina los \n
        palabras = line.split(', ')
        for p in palabras:
            p = p.replace(',', _STR_VACIO)
        funcion = palabras[_PRIMER_PARAMETRO].split(_ESPACIOS, 1)
        palabras.pop(_PRIMER_PARAMETRO)
        palabras.insert(_PRIMER_PARAMETRO, funcion[_PRIMER_PARAMETRO+1])
        funcion = funcion[_PRIMER_PARAMETRO]
        if funcion not in diccComandos:
            continue
        comando = diccComandos[funcion]
        cantParametros = diccParametros[funcion]
        if len(palabras) != cantParametros:
            continue
        while len(palabras) != _CANTIDAD_PARAMETROS_NECESARIOS:
            palabras.insert(_PRIMER_PARAMETRO, " ")
        try:
            print(comando(palabras[_PRIMER_PARAMETRO], palabras[_PRIMER_PARAMETRO + 1],
                          palabras[_PRIMER_PARAMETRO + 2], mapa))
        except:
            continue
        sys.stdout.flush()
