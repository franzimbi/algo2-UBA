from grafo import Grafo
from collections import deque


def topo_dfs(grafo, v, visitados, pila):
    visitados.add(v)
    for w in grafo.adyacentes(v):
        if w not in visitados:
            topo_dfs(grafo, w, visitados, pila)
    pila.append(v)


def topologicoDfs(grafo):
    visitados = set()
    pila = []
    for v in grafo:
        if v not in visitados:
            topo_dfs(grafo, v, visitados, pila)
    return pila[::-1]


def topologicoBfs(grafo):
    gradosEntrada = {}
    resultado = []
    for v in grafo:
        for w in grafo.adyacentes(v):
            if w not in gradosEntrada:
                gradosEntrada[w] = 1
            else:
                gradosEntrada[w] += 1
    cola = deque()
    for v in grafo:
        if v not in gradosEntrada:
            cola.append(v)
    while len(cola) != 0:
        v = cola.popleft()
        resultado.append(v)
        for w in grafo.adyacentes(v):
            gradosEntrada[w] -= 1
            if gradosEntrada[w] == 0:
                cola.append(w)
    return resultado


gr = Grafo(True)
gr.insertar_arista('fisica1', 'fisica2')
gr.insertar_arista('analisis2', 'fisica2')
gr.insertar_arista('analisis2', 'proba')
gr.insertar_arista('analisis2', 'numerico')
gr.insertar_arista('algebra2', 'proba')
gr.insertar_arista('algebra2', 'estructuras')
gr.insertar_arista('algebra2', 'numerico')
gr.insertar_arista('algo1', 'algo2')
gr.insertar_arista('fisica2', 'estructuras')
gr.insertar_arista('algo2', 'estructuras')
gr.insertar_arista('algo2', 'numerico')
gr.insertar_arista('algo2', 'algo3')
gr.insertar_arista('estructuras', 'taller')
gr.insertar_arista('estructuras', 'datos')
gr.insertar_arista('numerico', 'taller')

print(topologicoDfs(gr))
print(topologicoBfs(gr))
