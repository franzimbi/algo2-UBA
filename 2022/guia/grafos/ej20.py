from grafo import Grafo
from collections import deque

#   Implementar un algoritmo que reciba un grafo dirigido
#   y nos devuelva la cantidad de componentes débilmente conexas de este.
#   Indicar y justificar la complejidad del algoritmo implementado.


def bfs(grafo, v, visitados):
    cola = deque()
    cola.append(v)
    visitados.add(v)
    while len(cola) != 0:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                cola.append(w)


def componentesDebilmenteConexas(grafo):
    visitados = set()
    cantidadComp = 0
    v = grafo.vertice_aleatorio()
    for ady in grafo:
        if ady not in visitados:
            bfs(grafo, v, visitados)
            cantidadComp += 1

    return cantidadComp


gr = Grafo(True)
gr.insertar_vertice('a')
gr.insertar_vertice('b')
gr.insertar_arista('a', 'b')
gr.insertar_vertice('c')

print(componentesDebilmenteConexas(gr))
