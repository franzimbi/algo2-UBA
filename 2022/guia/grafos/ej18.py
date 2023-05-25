from grafo import Grafo
from collections import deque
#   Dado un número inicial X se pueden realizar dos tipos de operaciones sobre el número:

#   Multiplicar por 2
#   Restarle 1.

#   Implementar un algoritmo que encuentra la menor cantidad de operaciones a realizar
#   para convertir el número X en el número Y, con tan solo las operaciones mencionadas
#   arriba (podemos aplicarlas la cantidad de veces que querramos).


def construirGrafoDeOperaciones(inicio, fin):
    grafo = Grafo(True)
    grafo.insertar_vertice(inicio)
    cola = deque()
    cola.append(inicio)
    while len(cola) != 0:
        v = cola.popleft()
        multiplicacion = int(v) * 2
        sustraccion = int(v) - 1
        grafo.insertar_arista(v, str(multiplicacion))
        grafo.insertar_arista(v, str(sustraccion))
        if str(multiplicacion) == fin or str(sustraccion) == fin:
            return grafo
        else:
            cola.append(str(multiplicacion))
            cola.append(str(sustraccion))


def bfsCaminoMinimo(grafo, inicio, fin):
    orden = {}
    visitados = set()
    orden[inicio] = 0
    visitados.add(inicio)
    cola = deque()
    cola.append(inicio)
    while len(cola) != 0:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                orden[w] = orden[v] + 1
                cola.append(w)
            if w == fin:
                return orden[w]
    return None


def minimaCantidadOperaciones(inicio, fin):
    grafo = construirGrafoDeOperaciones(str(inicio), str(fin))
    return bfsCaminoMinimo(grafo, str(inicio), str(fin))


print(minimaCantidadOperaciones(8, 14))
