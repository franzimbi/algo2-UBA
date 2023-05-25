import grafo
from collections import deque

# El grado de todos los vértices de un grafo no dirigido


def gradosVertices(grafo):
    grados = {}
    for v in grafo:
        grados[v] = len(grafo.adyacentes(v))
    return grados

# El grado de salida (y entrada) de todos los vértices de un grafo dirigido.


def gradosEntradaYSalida(grafo):
    entrada = {}
    salida = {}
    for v in grafo:
        entrada[v] = 0
        salida[v] = 0
    for v in grafo:
        for w in grafo.adyacentes(v):
            entrada[w] = entrada[w] + 1
            salida[v] = salida[v] + 1
    return entrada, salida


def esConexo(grafo):
    visitados = set()
    q = deque()
    v = grafo.vertice_aleatorio()
    q.append(v)
    visitados.add(v)
    while len(q) != 0:  # O(v)
        v = q.popleft()
        for w in grafo.adyacentes(v):  # O(V)
            if w not in visitados:
                visitados.add(w)
                q.append(w)
    return len(visitados) == len(grafo)
# para matriz de adyacencia es O(V^2)
# matriz de incidencia es O(V*E)


ej = grafo.Grafo()
ej.insertar_arista('A', 'B')
ej.insertar_arista('C', 'B')
ej.insertar_arista('A', 'D')
ej.insertar_arista('A', 'E')
ej.insertar_arista('F', 'E')
ej.insertar_arista('F', 'A')
ej.insertar_arista('W', 'C')

print(esConexo(ej))
