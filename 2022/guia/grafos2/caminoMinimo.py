from tdaGrafo import Grafo
import heapq
import math


def dijkstra(grafo, origen):
    dist = {}
    for v in grafo:
        dist[v] = float("inf")
    padres = {}
    padres[origen] = None
    dist[origen] = 0
    h = []
    heapq.heappush(h, (0, origen))
    while len(h) != 0:  # O(E)
        p, v = heapq.heappop(h)  # O(log V)
        for w in grafo.adyacentes(v):
            if dist[w] > dist[v] + grafo.pesoArista(v, w):
                dist[w] = dist[v] + grafo.pesoArista(v, w)
                padres[w] = v
                heapq.heappush(h, (dist[w], w))
    return dist, padres
# O(E log V)


def obtenerAristasNoDir(grafo):
    aristas = []
    dic = set()
    for v in grafo:
        for w in grafo.adyacentes(v):
            if v+w not in dic and w+v not in dic:
                aristas.append((v, w))
                dic.add(v+w)
    return aristas


def bellmanFord(grafo, origen):
    aristas = obtenerAristasNoDir(grafo)
    dist = {}
    padres = {}
    for v in grafo:
        dist[v] = float("inf")
    dist[origen] = 0
    padres[origen] = None
    for i in range(len(grafo)):
        for a in aristas:
            o, d = a
            if dist[d] > dist[o] + grafo.pesoArista(o, d):
                dist[d] = dist[o] + grafo.pesoArista(o, d)
                padres[d] = o
    for a in aristas:
        o, d = a
        if dist[d] > dist[o] + grafo.pesoArista(o, d):
            return False
    return dist, padres
# O(VxE)


gr = Grafo(False, True)
gr.insertarArista('A', 'B', 5)
gr.insertarArista('A', 'D', 1)
gr.insertarArista('D', 'B', 2)
gr.insertarArista('C', 'B', 3)
gr.insertarArista('C', 'D', 9)
gr.insertarArista('C', 'E', 5)
gr.insertarArista('D', 'E', 3)
gr.insertarArista('D', 'F', 8)
gr.insertarArista('E', 'F', 4)

print(dijkstra(gr, 'A'))
print(bellmanFord(gr, 'A'))
