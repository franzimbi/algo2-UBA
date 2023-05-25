#include "heap.h"
#include "lista.h"

lista_t* los_mejors_k_superheroes(lista_t* heroes, size_t k, cmp_f cmp){
	heap_t* heap = heap_crear(cmp);
	if(heap == NULL) return NULL;

	while(!lista_esta_vacia(heroes)){ //O(n)
		heap_encolar(heap, lista_borrar_primero(heroes)); //O(log n)
	}
	lista_t* k_heroes = lista_crear();
	if(k_heroes == NULL){
		heap_destruir(heap, NULL); // podria destruirlos o devolverlos a la lista.
		return NULL;
	}
	for(size_t i =0; i<k; i++){ // O(k)
		heroe_t* heroe = heap_desencolar(heap);
		if (heroe == NULL) // si n < k devuelve n heroes
			break;
		lista_insertar_primero(k_heroes, heroe); // O(log n)
	}
	heap_destruir(heap, NULL);
	return k_heroes;
}

/* esta implementacion da O((n+k) log n). la unica forma de hacerlo O(n + k log n) es usando heapify para hacer un heap en O(n) con la lista. pero heapify segun las teoricas
es solo para arreglos. */

//seria reemplazar el while y la creacion del heap por un heap_t* heap = heapify (lista_t* heroes) -> O(n + k log n)