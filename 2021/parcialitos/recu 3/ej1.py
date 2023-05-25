import grafo

def calcular_conexion(grafo, vertice):
	cant_vertices = len(grafo.adyacentes(vertice)) # O(V)
	return ((cant_vertices - 1) * cant_vertices)/2 # O(1)

def es_altamente_conectado(grafo):
	maximo_posible = (len(grafo) * (len(grafo)-1))/2 #la cantida maxima de E que puede tener un grafo no dirigido es max(E) = V (V-1)/2 
	maximo_posible = (maximo_posible*40)/100
	for v in grafo: # O(V)
		aux = calcular_conexion(grafo, v) # O(V)
		if aux < maximo_posible: # si aux es menor que el 40% del maximo posible
			return False
	return True

""" Ver todos los adyacentes en una matriz de adyacencias es O(V), ya que tiene q ir a la fila del vector especifico y avanzar columna por columna.
	despues, calcular_conexiones es O(V) por lo mismo que antes. y como hay un for O(V) que adentro hace el calculo en O(V) -> O(V^2) es el orden de es_altamente_conectado"""

