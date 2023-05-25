#include "pila.h" //considero el TDA pila del tp

/* Dada una pila, implementar una función que devuelva el largo de la misma, la pila debe
quedar en su estado original. Justificar el orden del algoritmo. */

int tamano_pila(pila_t* pila){

    pila_t* pila_aux=pila_crear(); // O(1)
    if(pila_aux==NULL) //o(1)
        return -1;

    int tam=0;
    while(pila_esta_vacia(pila)==false){ //o(n)

        pila_apilar(pila_aux, pila_desapilar(pila));
        tam++;
    }

    while(pila_esta_vacia(pila_aux)==false){ //o(n)
        pila_apilar(pila, pila_desapilar(pila_aux));
    }

    return tam; //o(1)
}

/* definiendo como "n" a la cantidad de elementos que puedo encontrar en la pila: el orden de este algoritmo es de o(2n).
considero que la creacion de una pila auxiliar es de tiempo constante (o(1)) y que las dos desapilacions/apilaciones son de
orden o(n) cada una. o(n) + o(n) = o(2n)