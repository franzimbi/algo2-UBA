from tdaGrafo import Grafo
from collections import deque


def caminoMinimoBfs(grafo, origen):
    cola = deque()
    visitados = set()
    distancia = {}
    cola.append(origen)
    distancia[origen] = 0
    visitados.add(origen)

    while len(cola) != 0:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                distancia[w] = distancia[v] + 1
                cola.append(w)
    return distancia


def diametro(grafo):
    maxDistancia = 0
    for v in grafo:
        distancia = caminoMinimoBfs(grafo, v)
        for d in distancia:
            if distancia[d] > maxDistancia:
                maxDistancia = distancia[d]
    return maxDistancia
    # O(V*(V+E))


gr = Grafo()
gr.insertarArista('A', 'B')
gr.insertarArista('C', 'B')
gr.insertarArista('D', 'C')
gr.insertarArista('E', 'D')

print(diametro(gr))
