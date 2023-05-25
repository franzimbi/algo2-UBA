from grafo import Grafo


def puedeSerNoDirigido(grafo):
    for v in grafo:
        for w in grafo.adyacentes(v):  # O(e1 + e2 + e3 + ...)
            if not grafo.ver_dos_vertices_unidos(v, w):
                return False
            if not grafo.ver_dos_vertices_unidos(w, v):
                return False
    return True
# O(V+E)


gr = Grafo(True)

gr.insertar_arista('A', 'B')
gr.insertar_arista('B', 'A')
gr.insertar_arista('A', 'C')
gr.insertar_arista('C', 'A')

print(puedeSerNoDirigido(gr))
