from grafo import Grafo
from collections import deque
from ejemplo import cicloDfs

# Implementar un algoritmo que, dado un grafo dirigido, nos devuelva un ciclo
# dentro del mismo, si es que lo tiene. Indicar el orden del algoritmo.


def reconstruirCamino(padres, fin):
    camino = []
    while fin != None:
        camino.append(fin)
        fin = padres[fin]
    return camino[::-1]


def _cicloDfsDir(grafo, v, visitados, padres):
    for w in grafo.adyacentes(v):  # O(e1 + e2 + ... + ek)
        if w in visitados:
            return reconstruirCamino(padres, v)
        else:
            visitados.add(w)
            padres[w] = v
            ciclo = _cicloDfsDir(grafo, w, visitados, padres)
            if ciclo is not None:
                return ciclo
    return None


def cicloGrafoDirigido(grafo):
    visitados = set()
    for v in grafo:  # O(V)
        if v not in visitados:
            visitados.add(v)
            padres = {}
            padres[v] = None
            ciclo = _cicloDfsDir(grafo, v, visitados, padres)
            if ciclo is not None:
                return ciclo
    return None
# O(V+E)


def esConexo(grafo):
    visitados = set()
    padres = {}
    cola = deque()
    v = grafo.vertice_aleatorio()
    visitados.add(v)
    cola.append(v)
    padres[v] = None
    while len(cola) != 0:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                padres[w] = v
                cola.append(w)
    return len(cola) == len(visitados)


def esArbol(grafo):
    ciclo = cicloDfs(grafo)
    if ciclo is not None:
        return False
    return esConexo(grafo)


ej = Grafo()

ej.insertar_arista('A', 'B')
ej.insertar_arista('C', 'B')
ej.insertar_arista('B', 'D')

print(ej)
print(esArbol(ej))
