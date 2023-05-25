#include "pila.h"

int cantidad_pila(pila_t* pila){
    
    pila_t* aux=pila_crear();

    int i=0;

    while(!pila_esta_vacia(pila)){

        pila_apilar(aux, pila_desapilar(pila));
        i++;
    }
    while(!pila_esta_vacia(aux)){

        pila_apilar(pila, pila_desapilar(aux));
    }

    return i;
}

void ordenar_pila(pila_t* pila) {
    
    pila_t* aux=pila_crear();
    if(aux==NULL)
        return;

    int tam = cantidad_pila(pila);

    for(size_t i=0; i<tam; i++){
        
    }
}