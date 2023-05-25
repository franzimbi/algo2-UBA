from grafo import Grafo
from collections import deque

#  Implementar un algoritmo que reciba un grafo dirigido, un vértice V
# y un número N, y devuelva una lista con todos los vértices que se encuentren
#  a exactamente N aristas de distancia del vértice V. Indicar el tipo de recorrido
#  utilizado y el orden del algoritmo. Justificar.


def aDistanciaN(grafo, v, n):
    orden = {}
    visitados = set()
    visitados.add(v)
    orden[v] = 0
    vertices = []
    q = deque()
    q.append(v)
    while len(q) != 0:
        v = q.popleft()
        if orden[v] == n:
            vertices.append(v)
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                orden[w] = orden[v] + 1
                if orden[w] <= n:
                    q.append(w)
    return vertices
# O(V + E) -> BFS


gr = Grafo(True)
gr.insertar_arista('A', 'B')
gr.insertar_arista('A', 'C')
gr.insertar_arista('B', 'D')
gr.insertar_arista('C', 'E')
gr.insertar_arista('E', 'F')
gr.insertar_arista('E', 'G')
gr.insertar_arista('D', 'H')
gr.insertar_arista('H', 'I')
gr.insertar_arista('I', 'J')
gr.insertar_arista('J', 'K')
gr.insertar_arista('J', 'M')
gr.insertar_arista('G', 'N')
gr.insertar_arista('N', 'O')

print(gr)
print(aDistanciaN(gr, 'A', 3))
