from grafo import Grafo

#   Implementar una función que reciba un grafo no dirigido y no pesado implementado
#   con listas de adyacencia (diccionario de diccionarios) y devuelva una matriz que
#   sea equivalente a la representación de matriz de adyacencia del mismo grafo.
#   Indicar y justificar el orden del algoritmo implementado.


def matrizDeAdy(grafo):
    dic = {}
    matriz = []
    vertices = []
    for v in grafo:
        vertices.append(v)
    for i in range(len(vertices)):
        dic[vertices[i]] = i
    for i in range(len(dic)):
        lista = []
        for w in range(len(dic)):
            lista.append(0)
        matriz.append(lista)
    for v in grafo:
        for w in grafo.adyacentes(v):
            matriz[dic[v]][dic[w]] = 1
    return matriz
# O(V^2) -> pq es lo que se tarda en inicializar la matriz en 0s


gr = Grafo()
gr.insertar_arista('A', 'B')
gr.insertar_arista('C', 'B')
gr.insertar_arista('D', 'C')
gr.insertar_arista('D', 'E')
gr.insertar_arista('H', 'D')
gr.insertar_arista('F', 'E')
gr.insertar_arista('G', 'F')
gr.insertar_arista('H', 'G')

print(matrizDeAdy(gr))
