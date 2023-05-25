from tdaGrafo import Grafo
import heapq


def obtenerAristasNoDir(grafo):
    dic = set()
    aristas = []
    for v in grafo:  # O(V)
        for w in grafo.adyacentes(v):  # O(e)
            if v+w not in dic and w+v not in dic:
                dic.add(v+w)
                aristas.append((grafo.pesoArista(v, w), v, w))
    return aristas


def prim(grafo):
    arbol = Grafo(False, True)
    v = grafo.verticeAleatorio()
    visitados = set()
    h = []
    visitados.add(v)
    for d in grafo.adyacentes(v):
        heapq.heappush(h, (grafo.pesoArista(v, d), v, d))
    while len(h) != 0:  # O(E)
        p, o, d = heapq.heappop(h)  # O(log V)
        if d not in visitados:
            visitados.add(d)
            arbol.insertarArista(o, d, grafo.pesoArista(o, d))
            for w in grafo.adyacentes(d):
                heapq.heappush(h, (grafo.pesoArista(d, w), d, w))
    return arbol
# O(E log V)


def kruskal(grafo):
    aristas = obtenerAristasNoDir(grafo)
    arbol = Grafo(False, True)
    conjunto = UnionFind(grafo.listaVertices())
    aristas.sort()
    for i in range(len(aristas)):
        p, o, d = i
        if conjunto.find(o) == conjunto.find(d):
            continue
        conjunto.union(o, d)
        arbol.insertarArista(o, d, p)
    return arbol


gr = Grafo(False, True)
gr.insertarArista('A', 'B', 3)
gr.insertarArista('A', 'D', 6)
gr.insertarArista('A', 'E', 4)
gr.insertarArista('D', 'B', 4)
gr.insertarArista('D', 'E', 3)
gr.insertarArista('C', 'B', 2)
gr.insertarArista('D', 'C', 4)
gr.insertarArista('E', 'C', 5)

print(kruskal(gr))
