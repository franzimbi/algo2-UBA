# Implementar el Algoritmo de Kruskal para obtener el Árbol de Tendido Mínimo de un grafo pesado. Aplicar dicho algoritmo al grafo del dorso. Responder: ¿Para qué se utiliza la estructura Union-Find (o conjunto disjunto) en este algoritmo? ¿qué problema viene a resolver?

def obtener_aristas(grafo):
    aristas = []
    for v in grafo:
        for w in grafo.adyacentes(v):
            arista_append(v,w,grafo.peso(v,w))
    return aristas

def mst_kruskal(grafo):
    conjunto = UnionFind(grafo.obtener_vertices())
    aristas = sorted(obtener_aristas(grafo), key=lamda arista[peso])
    arbol= grafo(es_dirigido=false, obtener_vertices(grafo))
    for a in aristas:
        v, w, peso = a
        if conjunto.find(v) == conjunto.find(w)
            continue
        arbol.agregar_arista(v, w, peso)
        conjunto.union(v, w)
    return arbol

