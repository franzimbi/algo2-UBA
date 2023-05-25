from tdaGrafo import Grafo
import heapq


def prim(grafo):
    arbol = Grafo(False, True)
    visitados = set()
    h = []
    v = grafo.verticeAleatorio()
    visitados.add(v)
    for w in grafo.adyacentes(v):  # O(v)
        p = grafo.pesoArista(v, w)
        heapq.heappush(h, (p, v, w))
    while len(h) != 0:  # O(V+E)
        p, v, w = heapq.heappop(h)
        if w not in visitados:
            visitados.add(w)
            arbol.insertarArista(v, w, p)  # O(V)
            for j in grafo.adyacentes(w):
                if j not in visitados:
                    p = grafo.pesoArista(w, j)
                    heapq.heappush(h, (p, w, j))
    return arbol
    # O(E log V)
    # O(E log(V+E)) -> lista de ady
    # O(E log (V*E)) -> matriz de ady


gr = Grafo(False, True)
gr.insertarArista('A', 'B', 3)
gr.insertarArista('A', 'C', 4)
gr.insertarArista('C', 'B', 5)
gr.insertarArista('D', 'C', 2)
gr.insertarArista('D', 'B', 3)
gr.insertarArista('E', 'B', 3)
gr.insertarArista('E', 'D', 4)
gr.insertarArista('E', 'F', 6)
gr.insertarArista('F', 'D', 2)
gr.insertarArista('G', 'F', 1)
gr.insertarArista('G', 'C', 6)

print(prim(gr))
