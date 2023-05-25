
def numerar_vertices(grafo, vertice, numerados, nros_disponibles):
    if(len(numerados) == len(grafo)):
        return numerados
    v = vertice + 1 #avanzo de vertice
    for n in nros_disponibles:
        numerados[v] = n
        if (not es_valido(grafo, v, numerados)): #es_valido chequea q la solucion parcial a la q se llego en este punto sea valida o no
            numerados.remove(v)
            continue
        if numerar_vertices(grafo, v, numerados, nros_disponibles)!= None:
            return numerados
    numerados.remove(v)
    return None
