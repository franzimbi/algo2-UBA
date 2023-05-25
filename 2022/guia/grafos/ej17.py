from grafo import Grafo
from collections import deque

#   Implementar una función que reciba un grafo no dirigido,
#   y que compruebe la siguiente afirmación: “La cantidad de vértices
#   de grado IMPAR es PAR”. Indicar y justificar el orden del algoritmo
#   si el grafo está implementado como matriz de adyacencia.


def esParLaCantidadDeVerticesDeGradoImpar(grafo):
    visitados = set()
    verticesImpares = 0
    v = grafo.vertice_aleatorio()
    q = deque()
    visitados.add(v)
    q.append(v)
    for v in grafo:
        if v not in visitados:
            while len(q) != 0:
                v = q.popleft()
                if len(grafo.adyacentes(v)) % 2 == 1:
                    verticesImpares += 1
                for w in grafo.adyacentes(v):
                    if w not in visitados:
                        visitados.add(w)
                        q.append(w)
    return verticesImpares % 2 == 0


gr = Grafo()
gr.insertar_arista('A', 'B')
gr.insertar_arista('A', 'C')
gr.insertar_arista('C', 'D')
gr.insertar_arista('E', 'C')


print(esParLaCantidadDeVerticesDeGradoImpar(gr))
