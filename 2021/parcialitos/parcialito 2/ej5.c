/*
Dado un árbol binario que tiene como datos números enteros, implementar una primitiva que reemplace el valor de los
nodos internos del árbol de manera tal que cada nodo interno tenga como valor la suma de los valores de sus nodos hijos. No
se debe modificar el valor de las hojas. ¿Cuál es el orden de la primitiva implementada? ¿Qué tipo de recorrido utilizaste?
La estructura de cada nodo es la siguiente:
typedef struct ab_t {
int valor;
struct ab_t* izq;
struct ab_t* der
} ab_t;
*/

#include "abb.h"

typedef struct ab_t {
int valor;
struct ab_t* izq;
struct ab_t* der
} ab_t;

void reemplazar_internos_por_suma_de_hijos(ab_t* raiz){
    if(raiz==NULL)
        return;

    reemplazar_internos_por_suma_de_hijos(raiz->izq);
    reemplazar_internos_por_suma_de_hijos(raiz->der);

    int izq = 0;
    int der = 0;

    if(raiz->izq==NULL && raiz->der==NULL)
        return;
    if(raiz->izq!=NULL)
        izq = raiz->izq->valor;
    if(raiz->der != NULL)
        der = raiz->der->valor;
    raiz->valor = izq + der;
}

// LA PRIMITIVA DESARROLLA REALIZA UN RECORRIDO POSTORDER, YA QUE PRIMERO VA A LA IZQUIERDA, LUEGO A LA DERECHA Y FINALMENTE MIRA EL DATO ACTUAL.
//  EL ORDEN DE ESTA PRIMITIVA ES DE O(n), SIENDO n LA CANTIDAD DE NODOS. ESTO ES ASI PUESTO Q RECORRE CADA NODO UNA SOLA VEZ, SIENDO TODAS LAS 
// OTRAS ASIGNACIONES DE ORDEN O(1).
