import grafo
from collections import deque

# Implementar un algoritmo que, dado un grafo no dirigido,
# nos devuelva un ciclo dentro del mismo, si es que los tiene.
# Indicar el orden del algoritmo.


def _bfs(origen, grafo, padres, visitados, orden):
    if origen in visitados:
        return
    padres[origen] = None
    orden[origen] = 0
    q = deque()  # cola
    visitados.add(origen)
    q.append(origen)
    while len(q) != 0:
        v = q.popleft()  # desencolar
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                orden[w] = orden[v]+1
                visitados.add(w)
                q.append(w)  # encolar


def reconstruirCamino(destino, padres):
    recorrido = []
    while destino is not None:
        recorrido.append(destino)
        destino = padres[destino]
    return recorrido[::-1]


def reconstruirCiclo(padre, inicio, fin):
    v = fin
    camino = []
    while v != inicio:
        camino.append(v)
        v = padre[v]
    camino.append(inicio)
    return camino[::-1]


def bfs(grafo):
    padres = {}
    visitados = set()
    orden = {}
    for v in grafo:
        _bfs(v, grafo, padres, visitados, orden)
    return padres, orden


def _cicloBfs(v, grafo, visitados):
    q = deque()
    q.append(v)
    visitados.add(v)
    padre = {}
    padre[v] = None
    while not len(q) != 0:
        v = q.popleft()
        for w in grafo.adyacentes(v):
            if w in visitados:
                if w != padre[v]:
                    return reconstruirCiclo(padre, w, v)
            else:
                q.append(w)
                visitados.add(v)
                padre[w] = v
    return None


def cicloBfs(grafo):
    visitados = set()
    for v in grafo:
        if v not in visitados:
            ciclo = _cicloBfs(v, grafo, visitados)
            if ciclo is not None:
                return ciclo
    return None


def _dfs(v, grafo, padres, orden, visitados):
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            padres[w] = v
            orden[w] = orden[v] + 1
            _dfs(w, grafo, padres, orden, visitados)


def dfsCompleto(grafo):
    padres = {}
    orden = {}
    visitados = set()
    for v in grafo:
        if v not in visitados:
            orden[v] = 1
            visitados.add(v)
            padres[v] = None
            _dfs(v, grafo, padres, orden, visitados)
    return padres, orden


def _cicloDfs(v, grafo, padres, visitados):
    visitados.add(v)
    for w in grafo.adyacentes(v):
        if w in visitados:
            if padres[v] != w:
                return reconstruirCiclo(padres, w, v)
        else:
            padres[w] = v
            ciclo = _cicloDfs(w, grafo, padres, visitados)
            if ciclo is not None:
                return ciclo
    return None


def cicloDfs(grafo):
    visitados = set()
    padres = {}
    for v in grafo:
        if v not in visitados:
            padres[v] = None
            ciclo = _cicloDfs(v, grafo, padres, visitados)
            if ciclo is not None:
                return ciclo
    return None


def construirGrafo():
    g = grafo.Grafo()
    g.insertar_arista('A', 'H')
    g.insertar_arista('A', 'B')
    g.insertar_arista('H', 'F')
    g.insertar_arista('B', 'F')
    #g.insertar_arista('F', 'D')
    #g.insertar_arista('C', 'D')
    return g


ej = construirGrafo()
print(ej)
print("con dfs: ", cicloDfs(ej))
print("con bfs: ", cicloBfs(ej))
