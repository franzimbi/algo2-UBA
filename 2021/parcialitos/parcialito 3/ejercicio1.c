#include "heap.h"
#include "lista.h"
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

int cmp_int(void* a, void* b){
    return (int)b - (int)a;
}

lista_t* lista_mayores_k(int* arr, size_t tam, size_t k){
    heap_t* heap = heap_crear_arr((void*) &arr, tam, cmp_int); // la funcion cmp_int hace que el heap sea de minimos
    
    while(heap_ver_max(heap) <= k){
        int nro_menor = heap_desencolar(heap);
        if(nro_menor > k)
            break;
        int nro_menor2 = heap_desencolar(heap);
        if(nro_menor>k){
            heap_destruir(heap, NULL);
            return NULL;
        }
        int aux = nro_menor + nro_menor2 * 2;
        heap_encolar(heap, (void*) &aux);
    }
    lista_t* lista = lista_crear();
    while(!heap_esta_vacio(heap))
        lista_insertar_ultimo(lista, heap_desencolar(heap));
    heap_destruir(heap, NULL);
    return lista;
}