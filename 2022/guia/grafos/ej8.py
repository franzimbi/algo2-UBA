from grafo import Grafo
from collections import deque

# La teoría de los 6 grados de separación dice que cualquiera en la Tierra puede estar
#  conectado a cualquier otra persona del planeta a través de una cadena de conocidos
# que no tiene más de cinco intermediarios (conectando a ambas personas con solo seis
# enlaces). Suponiendo que se tiene un grafo G en el que cada vértice es una persona y
# cada arista conecta gente que se conoce (el grafo es no dirigido):

# a. Implementar un algoritmo para comprobar si se cumple tal teoría para todo el
# conjunto de personas representadas en el grafo G. Indicar el orden del algoritmo.

# b. Suponiendo que en el grafo G no habrán altas ni bajas de vértices, pero podrían
# haberla de aristas (la gente se va conociendo), explicar las ventajas y desventajas
# que tendría implementar al grafo G con una matriz de adyacencia.


def bfsGrado6(grafo, v):
    visitados = set()
    orden = {}
    q = deque()
    q.append(v)
    orden[v] = 0
    visitados.add(v)
    while len(q) != 0:
        v = q.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                orden[w] = orden[v] + 1
                q.append(w)
            if orden[w] > 5:
                return False
    return len(visitados) == len(grafo)


def cumpleTeoriaDeSeisGrados(grafo):
    v = grafo.vertice_aleatorio()
    return bfsGrado6(grafo, v)

# si lo implemento con matriz de adyacencia puedo ver si las filas son LI, en caso
# de serlo significaria que cumpliria


gr = Grafo()
gr.insertar_arista('A', 'B')
gr.insertar_arista('B', 'C')
gr.insertar_arista('D', 'C')
gr.insertar_arista('E', 'D')
gr.insertar_arista('F', 'E')
gr.insertar_arista('F', 'G')
gr.insertar_arista('H', 'G')
gr.insertar_arista('H', 'I')
gr.insertar_arista('I', 'C')

print(cumpleTeoriaDeSeisGrados(gr))
