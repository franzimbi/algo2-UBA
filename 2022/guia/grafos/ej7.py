from grafo import Grafo


def traspuesto(grafo):
    tras = Grafo(True)
    for v in grafo:  # O(V) | O(V) | O(V)
        for w in grafo.adyacentes(v):  # O(e) | O(v) | O(v)
            tras.insertar_arista(w, v)  # O(V*E) | O(1) | O(v)
    return tras
    # O((V*E)^2) matriz incidencia
    # O(V^2) matriz de adyacencia
    # O(V * E) lista de ady


grafo = Grafo(True)
grafo.insertar_arista('A', 'B')
grafo.insertar_arista('A', 'C')
grafo.insertar_arista('B', 'D')
grafo.insertar_arista('D', 'F')
grafo.insertar_arista('F', 'A')

print(grafo)
print(traspuesto(grafo))
