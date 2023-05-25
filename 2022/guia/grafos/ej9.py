from grafo import Grafo


def _dfs(grafo, v, orden, visitados):
    orden.append(v)
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            _dfs(grafo, w, orden, visitados)


def mapaMuseo(grafo):
    orden = []
    visitados = set()
    v = grafo.vertice_aleatorio()
    visitados.add(v)
    _dfs(grafo, v, orden, visitados)
    return orden
# O(V + E) -> recorrido dfs


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
print(mapaMuseo(gr))
