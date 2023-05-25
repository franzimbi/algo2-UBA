from grafo import Grafo


# Un autor decidió escribir un libro con varias tramas que se puede leer de forma no lineal.
# Es decir, por ejemplo, después del capítulo 1 puede leer el 2 o el 73; pero la historia
#  no tiene sentido si se abordan estos últimos antes que el 1.

# Siendo un aficionado de la computación, el autor ahora necesita un orden para publicar
# su obra, y decidió modelar este problema como un grafo dirigido, en dónde los capítulos
#  son los vértices y sus dependencias las aristas. Así existen, por ejemplo, las aristas
#  (v1, v2) y (v1, v73).

# Escribir un algoritmo que devuelva un orden en el que se puede leer la historia sin
# obviar ningún capítulo. Indicar la complejidad del algoritmo.


def _dfs(grafo, v, visitados, lista):
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            lista.append(w)
            _dfs(grafo, w, visitados, lista)


def ordenLibro(grafo, cap1):
    visitados = set()
    lista = []
    visitados.add(cap1)
    lista.append(cap1)
    _dfs(grafo, cap1, visitados, lista)
    return lista


gr = Grafo(True)
gr.insertar_arista('1', '2')
gr.insertar_arista('1', '75')
gr.insertar_arista('2', '8')
gr.insertar_arista('2', '16')
gr.insertar_arista('75', '88')
gr.insertar_arista('75', '94')

print(ordenLibro(gr, '1'))
