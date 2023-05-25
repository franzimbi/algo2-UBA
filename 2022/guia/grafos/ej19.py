from grafo import Grafo
from collections import deque

#   Se tiene un arreglo de palabras de un lenguaje alienigena. Dicho arreglo se
#   encuentra ordenado para dicho idioma (no conocemos el orden de su abecedario).
#   Implementar un algoritmo que reciba dicho arreglo y determine un orden posible
#   para las letras del abecedario en dicho idioma. Por ejemplo:
#       {"caa", "acbd", "acba", "bac", "bad"} --> ['c', 'd', 'a', 'b']


def crearGrafoPalabras(palabras):
    grafo = Grafo(True)
    for p in range(len(palabras)-1):
        palabra1 = palabras[p]
        palabra2 = palabras[p+1]
        for i in range(len(palabra1)):
            if palabra1[i] != palabra2[i]:
                grafo.insertar_arista(palabra1[i], palabra2[i])
                break
    return grafo


def ordenPalabras(palabras):
    grafo = crearGrafoPalabras(palabras)
    grados = {}
    for v in grafo:
        for w in grafo.adyacentes(v):
            if w not in grados:
                grados[w] = 1
            else:
                grados[w] += 1
    cola = deque()
    orden = []
    for v in grafo:
        if v not in grados:
            cola.append(v)

    while len(cola) != 0:
        v = cola.popleft()
        orden.append(v)
        for w in grafo.adyacentes(v):
            grados[w] -= 1
            if grados[w] == 0:
                cola.append(w)
    return orden


print(crearGrafoPalabras(list({"caa", "acbd", "acba", "bac", "bad"})))
