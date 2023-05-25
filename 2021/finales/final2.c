#include <stdio.h>
#include <string.h>
#include <stdbool.h>

typedef struct arbol {
void* dato;
struct arbol* izq;
struct arbol* der;
} ab_t;

/*Implementar en C una primitiva para el árbol binario que determine si el mismo es un árbol completo. Indicar y
justificar la complejidad de la primitiva.*/

bool ab_es_completo(ab_t* ab){
    if(ab==NULL)
        return true;
    if(ab->der!=NULL && ab->izq==NULL)
        return false;
    if(ab->izq==NULL)
        return true;

    bool izq = ab_es_completo(ab->izq);
    bool der = ab_es_completo(ab->der);

    return izq && der;
}

/* t(n) = AT(n/B) + O(n^C)

A= 2
B = 2
C = 0 -> O(n^C) = O(1)

log2(2) > 0 -> T(n) = T(n^1) -> T(n)

*/

/* Implementar en C una primitiva para la lista enlazada lista_t* lista_slice(const lista_t*, size_t inicio,
size_t fin) que dada una lista devuelva una nueva lista con los elementos de la primera comprendidos entre la
posición incio (incluyendo) y fin (sin incluir). La lista original no debe ser modificada. Indicar y justificar la complejidad
de la primitiva implementada. */

typedef struct nodo{
    void* dato;
    struct nodo* proximo;
}nodo_t;

struct lista{
    struct nodo* primero;
    struct nodo* ultimo;
    size_t largo;
};

lista_t* lista_slice(const lista_t* lista, size_t inicio, size_t fin){
    lista_t* nueva = lista_crear();
    if(lista==NULL)
        return nueva;
    nodo_t* actual = lista->primero;
    for( size_t i=0; i<inicio; i++){
        if(actual==NULL){
            lista_destruir(nueva);
            return NULL;
        }
        actual= actual->proximo;
    }
    for(size_t j=inicio; j<fin; i++){
        if(actual==NULL){
            lista_destruir(nueva);
            return NULL;
        }
        lista_insertar_ultimo(actual->dato);
        actual = actual->proximo;
    }
    return nueva;
}
# O(fin)
