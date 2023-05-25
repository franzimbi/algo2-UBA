from tdaGrafo import Grafo
from pila import Pila


def dfsTarjanPA(grafo, v, orden, masBajo, padres, puntosArticulacion):
    for w in grafo.adyacentes(v):
        if w not in orden:
            orden[w] = orden[v] + 1
            masBajo[w] = orden[w]
            padres[w] = v
            dfsTarjanPA(grafo, w, orden, masBajo, padres, puntosArticulacion)
            if masBajo[w] >= orden[v] and padres[v] != None:
                puntosArticulacion.add(v)
        if padres[v] != w:
            if masBajo[w] < masBajo[v]:
                masBajo[v] = masBajo[w]


def puntosArticulacion(grafo):
    orden = {}
    masBajo = {}
    padres = {}
    puntosArticulacion = set()
    for v in grafo:
        if v not in orden:
            orden[v] = 0
            masBajo[v] = 0
            padres[v] = None
            dfsTarjanPA(grafo, v, orden, masBajo, padres, puntosArticulacion)
    return puntosArticulacion


def cFC(pila, apilados, origen):
    componente = []
    v = pila.desapilar()
    while v != origen:
        componente.append(v)
        apilados.remove(v)
        v = pila.desapilar()
    componente.append(v)
    apilados.remove(v)
    return componente


def dfsTarjanCFC(grafo, v, orden, masBajo, apilados, pila, componentesFuertCon, ordenGlobal):
    pila.apilar(v)
    apilados.add(v)
    for w in grafo.adyacentes(v):
        if w not in orden:
            orden[w] = ordenGlobal
            masBajo[w] = orden[w]
            ordenGlobal[0] += 1
            dfsTarjanCFC(grafo, w, orden, masBajo, apilados, pila,
                         componentesFuertCon, ordenGlobal)
        else:
            if masBajo[w] < masBajo[v] and v in apilados:
                masBajo[v] = masBajo[w]
    if masBajo[w] == orden[w] and w in apilados:
        componente = cFC(pila, apilados, w)
        componentesFuertCon.append(componente)


def componentesFuertementeConexas(grafo):
    orden = {}
    masBajo = {}
    apilados = set()
    pila = Pila()
    componentes = []
    ordenGlobal = [0]
    for v in grafo:
        if v not in orden:
            orden[v] = ordenGlobal
            ordenGlobal[0] += 1
            masBajo[v] = orden[v]
            dfsTarjanCFC(grafo, v, orden, masBajo, apilados,
                         pila, componentes, ordenGlobal)
            if not pila.es_vacia():
                c = cFC(pila, apilados, v)
                componentes.append(c)
    return componentes


gr = Grafo()
gr.insertarArista('A', 'B')
gr.insertarArista('C', 'B')
gr.insertarArista('C', 'G')
gr.insertarArista('F', 'C')
gr.insertarArista('D', 'B')
gr.insertarArista('E', 'D')
gr.insertarArista('H', 'E')
# print(puntosArticulacion(gr))

grD = Grafo(True)
grD.insertarArista('A', 'B')
grD.insertarArista('B', 'C')
grD.insertarArista('C', 'D')
grD.insertarArista('D', 'B')
grD.insertarArista('A', 'F')
grD.insertarArista('F', 'G')
grD.insertarArista('G', 'H')
grD.insertarArista('G', 'I')
grD.insertarArista('I', 'G')
grD.insertarArista('H', 'J')
grD.insertarArista('J', 'K')
grD.insertarArista('K', 'H')
#grD.insertarArista('Z', 'W')
print(grD)
print(componentesFuertementeConexas(grD))
