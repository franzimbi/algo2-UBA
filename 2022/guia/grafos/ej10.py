from grafo import Grafo
from collections import deque


def esBipartito(grafo):
    color = {}
    q = deque()
    v = grafo.vertice_aleatorio()
    color[v] = 0
    q.append(v)
    while len(q) != 0:
        v = q.popleft()
        for w in grafo.adyacentes(v):
            if w in color:
                if color[w] == color[v]:
                    return False
            else:
                q.append(w)
                if color[v] == 0:
                    color[w] = 1
                if color[v] == 1:
                    color[w] = 0
    return True
# O(v + E) -> recorrido bfs


gr = Grafo()
gr.insertar_arista('A', 'B')
gr.insertar_arista('C', 'B')
gr.insertar_arista('D', 'C')
gr.insertar_arista('E', 'C')
gr.insertar_arista('E', 'F')
gr.insertar_arista('G', 'D')
gr.insertar_arista('H', 'D')
gr.insertar_arista('I', 'H')
gr.insertar_arista('J', 'I')
gr.insertar_arista('K', 'I')

print(gr)
print(esBipartito(gr))
