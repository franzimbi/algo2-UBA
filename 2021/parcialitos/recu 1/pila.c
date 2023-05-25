#include "pila.h"

#include <stdlib.h>

#define TAMANO_PILA_INICIAL 10
#define CONSTANTE_REDIMENSION_DESAPILAR 4


/* Definición del struct pila proporcionado por la cátedra.
 */
struct pila {
    void **datos;
    size_t cantidad;   // Cantidad de elementos almacenados.
    size_t capacidad;  // Capacidad del arreglo 'datos'.
};

/* *****************************************************************
 *                    PRIMITIVAS DE LA PILA
 * *****************************************************************/

// ...
pila_t *pila_crear(void){
    pila_t* pila = malloc(sizeof(pila_t));  // creas el struct pila en  el heap
    if(pila==NULL)
        return NULL;

    pila->datos = malloc(sizeof(void*) * TAMANO_PILA_INICIAL); //creas el vector para apilar, que apunta a voids*
    if(pila->datos==NULL){
        free(pila);
        return NULL;
    }

    pila->cantidad=0;
    pila->capacidad=TAMANO_PILA_INICIAL;

    return pila;
}

void pila_destruir(pila_t *pila){
    free(pila->datos);
    free(pila);
}

//devuelve: la pila a capacidad x2 o NULL si no pudo agrandarse
//pre: la pila esta creada.
/* static pila_t* agrandar_pila(pila_t* pila){
    void** aux=realloc(pila->datos, sizeof(void *) * (pila->capacidad *2));
    if(aux==NULL)
        return NULL;

    pila->datos = aux;
    pila->capacidad *= 2;

    return pila;
}

//devuelve: la pila a capaciodad x 1/2 o NULL si no pudo achicarse.
//pre: la pila esta creada.
static pila_t* achicar_pila(pila_t* pila){
    void** aux=NULL;

    if((pila->capacidad)/2 <= TAMANO_PILA_INICIAL){
        aux=realloc(pila->datos, sizeof(void*) * TAMANO_PILA_INICIAL);
        if(aux==NULL)
            return NULL;

        pila->datos = aux;
        pila->capacidad = TAMANO_PILA_INICIAL;

        return pila;
    }

    else{
        aux=realloc(pila->datos, sizeof(void*) * ((pila->capacidad)/2));
        if(aux==NULL)
            return NULL;

        pila->datos = aux;
        pila->capacidad /= 2;

        return pila;
    }
} */

static bool pila_redimensionar(pila_t* pila, size_t nuevo_tam){
    
    void** aux=realloc(pila->datos, sizeof(void *) * nuevo_tam);
    if(aux==NULL)
        return false;

    pila->datos = aux;
    pila->capacidad = nuevo_tam;

    return true;
}

bool pila_esta_vacia(const pila_t *pila){

    return pila->cantidad == 0 ? true : false;
}

bool pila_apilar(pila_t *pila, void *valor){
    bool status=true;

    if(pila->capacidad == (pila->cantidad))
        status = pila_redimensionar(pila, pila->capacidad *2);
    if(status == false)
        return false;

    pila->datos[pila->cantidad] = valor;
    pila->cantidad++;

    return true;
}

void *pila_ver_tope(const pila_t *pila){

    return pila_esta_vacia(pila) ? NULL : pila->datos[pila->cantidad -1];
}

void *pila_desapilar(pila_t *pila){

    if(pila_esta_vacia(pila) == true)
        return NULL;

    if( (pila->cantidad * CONSTANTE_REDIMENSION_DESAPILAR <= pila->capacidad) ){

        if((pila->capacidad)/2 <= TAMANO_PILA_INICIAL){
            pila_redimensionar(pila, TAMANO_PILA_INICIAL);
    }  else{
        pila_redimensionar(pila, (pila->capacidad)/2);
    }
    }
    pila->cantidad --;
    return pila->datos[pila->cantidad];
}
