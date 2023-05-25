import grafo
from collections import deque


def proyeccion_v(grafo, v):
	grupo_de_V = []
	grupo_opuesto= []
	proyeccion = grafo.GraFO();
	q = deque()
	for w in grafo:
		grupo_de_V.append(v)
	q.append(v) #encolar
	while len(q) != 0:
		v = q.popleft() #desencolar
		for w in grafo.adyacentes(v):
			if w not in grupo_opuesto or w not in grupo_de_V:
				if v in grupo_de_V:
					grupo_opuesto.append(w)
					q.append(w)
				else:
					grupo_de_V.append(w)
					q.append(w)
					proyeccion.insertar_arista(v, w) # inserta la arista de (V,W) en O(1)
	return proyeccion

"""este algoritmo recorre en forma de BFS el grafo partiendo desde el vertice V. BFS es de orden O(V+E) """
