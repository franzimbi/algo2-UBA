import procesamiento
import algoritmosGrafos

_MSJ_NO_SE_ENCONTRO_RECORRIDO = "No se encontro recorrido"
_MSJ_TIEMPO_TOTAL = "\nTiempo total: "
_MSJ_PESO_MST = "Peso total: "


def nombreDocumentoKML(desde, hasta):
    return "Camino desde " + desde + " hacia " + hasta


def Ir(desde, hasta, archivo, grafo):
    if desde not in grafo or hasta not in grafo:
        return _MSJ_NO_SE_ENCONTRO_RECORRIDO
    resultado = algoritmosGrafos.CaminoMinimo(grafo, desde, hasta)
    if resultado == None:
        return _MSJ_NO_SE_ENCONTRO_RECORRIDO
    camino, tiempo = resultado
    procesamiento.CrearKMLCamino(nombreDocumentoKML(
        desde, hasta), archivo, grafo, camino)
    return procesamiento.FormatoDeSalidaDeCamino(camino) + _MSJ_TIEMPO_TOTAL + str(tiempo)


def Itinerario(_, __, archivo, grafo):
    grafoTopologico = procesamiento.CrearGrafoTopologicoConCSV(archivo)
    resultado = algoritmosGrafos.Recomendaciones(grafo, grafoTopologico)
    if resultado == None:
        return _MSJ_NO_SE_ENCONTRO_RECORRIDO
    return procesamiento.FormatoDeSalidaDeCamino(resultado)


def Viaje(_, origen, archivo, grafo):
    if origen not in grafo:
        return _MSJ_NO_SE_ENCONTRO_RECORRIDO
    resultado = algoritmosGrafos.Hierholzer(grafo, origen)
    if resultado == None:
        return _MSJ_NO_SE_ENCONTRO_RECORRIDO
    camino, tiempo = resultado
    procesamiento.CrearKMLCamino(nombreDocumentoKML(
        origen, origen), archivo, grafo, camino)
    return procesamiento.FormatoDeSalidaDeCamino(camino) + _MSJ_TIEMPO_TOTAL + str(tiempo)


def reducirCaminos(_, __, archivo, grafo):
    aristas, peso = algoritmosGrafos.ArbolTendidoMinimoPrim(grafo)
    if peso != 0:
        procesamiento.CrearPajekArbolTendidoMinimo(grafo, aristas, archivo)
    return _MSJ_PESO_MST + str(peso)
