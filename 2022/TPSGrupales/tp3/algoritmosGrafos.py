import heapq
from collections import deque
from tdaGrafo import Grafo


def dijkstra(grafo, origen):
    dist = {}
    padres = {}
    for v in grafo:
        dist[v] = float("inf")
    padres[origen] = None
    dist[origen] = 0
    heap = []
    heapq.heappush(heap, (0, origen))
    while len(heap) != 0:
        p, v = heapq.heappop(heap)
        for w in grafo.adyacentes(v):
            if dist[w] > dist[v] + grafo.pesoArista(v, w):
                dist[w] = dist[v] + grafo.pesoArista(v, w)
                padres[w] = v
                heapq.heappush(heap, (dist[w], w))
    return dist, padres


def CaminoMinimo(grafo, origen, destino):
    distancias, padres = dijkstra(grafo, origen)
    v = padres[destino]
    camino = []
    camino.append(destino)
    if destino not in distancias:
        return None
    while v != None:
        camino.append(v)
        v = padres[v]
    return (camino[::-1], distancias[destino])


def obtenerGradosEntrada(grafo):
    gEntrada = {}
    for v in grafo:
        gEntrada[v] = 0
    for v in grafo:
        for w in grafo.adyacentes(v):
            gEntrada[w] += 1
    return gEntrada


def ordenTopologico(grafo):
    gEntrada = obtenerGradosEntrada(grafo)
    cola = deque()
    resultado = []
    for v in grafo:
        if gEntrada[v] == 0:
            cola.append(v)
    while len(cola) != 0:
        v = cola.popleft()
        resultado.append(v)
        for w in grafo.adyacentes(v):
            gEntrada[w] -= 1
            if gEntrada[w] == 0:
                cola.append(w)
    return resultado


def Recomendaciones(grafoOriginal, grafoTopo):
    orden = ordenTopologico(grafoTopo)
    if len(orden) == 0:
        return None
    for v in grafoOriginal:
        if v not in grafoTopo:
            orden.append(v)
    return orden


def ArbolTendidoMinimoPrim(grafo):
    visitados = set()
    pesoTotal = 0
    arbol = []
    v = grafo.verticeAleatorio()
    visitados.add(v)
    heap = []
    for w in grafo.adyacentes(v):
        heapq.heappush(heap, (grafo.pesoArista(v, w), v, w))
    while len(heap) != 0:
        p, o, d = heapq.heappop(heap)
        if d not in visitados:
            arbol.append((o, d, p))
            pesoTotal += p
            visitados.add(d)
            for w in grafo.adyacentes(d):
                if w not in visitados:
                    heapq.heappush(heap, (grafo.pesoArista(d, w), d, w))
    return arbol, pesoTotal


def cantidadVerticesDeGradoPar(grafo):
    cant = 0
    for v in grafo:
        tam = len(grafo.adyacentes(v))
        if tam % 2 == 1:
            cant += 1
    return cant


def esConexo(grafo):
    visitados = set()
    q = deque()
    v = grafo.verticeAleatorio()
    q.append(v)
    visitados.add(v)
    while len(q) != 0:
        v = q.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                q.append(w)
    return len(visitados) == len(grafo)


def copiarGrafo(grafo):
    copia = Grafo(False, False)
    visitados = set()
    for v in grafo:
        for w in grafo.adyacentes(v):
            if v+w not in visitados and w+v not in visitados:
                visitados.add(v+w)
                copia.insertarArista(v, w)
    return copia


def dfs_aristas(v, mapa, camino, origen):
    camino.append(v)
    for w in mapa.adyacentes(v):
        if w == origen:
            mapa.borrarArista(v, w)
            camino.append(w)
            return camino
        else:
            mapa.borrarArista(v, w)
            c = dfs_aristas(w, mapa, camino, origen)
            if c != None:
                return c
    return None


def Hierholzer(grafo, origen):
    if cantidadVerticesDeGradoPar(grafo) != 0:
        return None
    mapa = copiarGrafo(grafo)
    if not esConexo(mapa):
        return None
    caminoC = dfs_aristas(origen, mapa, [], origen)
    if caminoC == None:
        return None
    hayCaminoU = True
    while hayCaminoU:
        hayCaminoU = False
        for c in range(len(caminoC)-1):
            caminoU = dfs_aristas(
                caminoC[c], mapa, [], caminoC[c])
            if caminoU != None:
                hayCaminoU = True
                pos = c
                caminoC.pop(pos)
                for u in caminoU:
                    caminoC.insert(pos, u)
                    pos += 1
    tiempoTotal = 0
    for e in range(1, len(caminoC)):
        tiempoTotal += grafo.pesoArista(caminoC[e-1], caminoC[e])
    return caminoC, tiempoTotal
