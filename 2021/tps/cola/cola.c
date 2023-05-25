#include "cola.h"
#include <stdio.h>
#include <stdlib.h>

typedef struct nodo{
    struct nodo* siguiente;
    void* dato;
}nodo_t;

struct cola{
    struct nodo* primero;
    struct nodo* ultimo;
};

static nodo_t *nodo_crear(void){
    nodo_t* nuevo=malloc(sizeof(nodo_t));
    if(nuevo==NULL)
        return NULL;

    nuevo->siguiente=NULL;
    nuevo->dato=NULL;

    return nuevo;
}

cola_t *cola_crear(void){

    cola_t* nuevo=malloc(sizeof(cola_t));
    if(nuevo==NULL)
        return NULL;

    nuevo->primero=NULL;
    nuevo->ultimo=NULL;

    return nuevo;
}

void cola_destruir(cola_t *cola, void (*destruir_dato)(void *)){
    while(!cola_esta_vacia(cola)){
        if(destruir_dato!=NULL)
            destruir_dato(cola_desencolar(cola));
        else{
            cola_desencolar(cola);
        }
    }
    free(cola);
}

bool cola_esta_vacia(const cola_t *cola){
    return cola->primero==NULL;
}

bool cola_encolar(cola_t *cola, void *valor){
    nodo_t* nuevo_nodo = nodo_crear();
    if(nuevo_nodo==NULL)
        return false;

    nuevo_nodo->dato=valor;
    nuevo_nodo->siguiente=NULL;
    if(cola_esta_vacia(cola)){
        cola->primero=nuevo_nodo;
        cola->ultimo=nuevo_nodo;

        return true;
    }
    cola->ultimo->siguiente=nuevo_nodo;
    cola->ultimo=nuevo_nodo;
    
    return true;
}

void *cola_ver_primero(const cola_t *cola){
    return cola_esta_vacia(cola) ? NULL : cola->primero->dato;
}

void *cola_desencolar(cola_t *cola){
    if(cola_esta_vacia(cola))
        return NULL;

    nodo_t* nodo_primero=cola->primero;
    if(cola->primero==cola->ultimo)
        cola->ultimo=NULL;

    cola->primero=cola->primero->siguiente;

    void* dato=nodo_primero->dato;
    free(nodo_primero);

    return dato;
}
