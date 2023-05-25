#include "pila.h"
#include "testing.h"
#include <stdio.h>
#include <stdlib.h>

static void prueba_pila_vacia(void){
    printf("\nINICIO DE PRUEBAS CON PILA VACIA\n");

    //creo una pila
    pila_t *pila = pila_crear();
    print_test("creo la pila", pila!=NULL);

    print_test("ver si pila esta vacia", pila_esta_vacia(pila)==true);
    print_test("ver si pila tiene algo para desapilar", pila_desapilar(pila)==NULL);

    print_test("destruyo la pila", pila!=NULL);
    pila_destruir(pila);
}

static void prueba_pila_apilar_desapilar(void){
    printf("\nINICIO DE PRUEBAS APILANDO Y DESAPILANDO\n");

    pila_t *pila = pila_crear();
    print_test("creo la pila", pila!=NULL);

    //creo elementos para apilar y los uso
    int* dato1=malloc(sizeof(int));
    if(dato1==NULL){
        printf("\nFALLO DE MEMORIA\n");
        return;
    }
    *dato1 =  45;

    float* dato2=malloc(sizeof(float));
    if(dato2==NULL){
        free(dato1);
        printf("\nFALLO DE MEMORIA\n");
        return;
    }
    *dato2 = 234.576f;

    char* dato3=malloc(sizeof(char));
    if(dato2==NULL){
        free(dato1);
        free(dato2);
        printf("\nFALLO DE MEMORIA\n");
        return;
    }
    *dato3='A';

    print_test("apilo un elemento", pila_apilar(pila, &dato1)==true);
    print_test("desapilo elemento", pila_desapilar(pila)==&dato1);
    print_test("ver si pila vuelve a estar vacia", pila_esta_vacia(pila)==true);
    print_test("apilo nuevamente el elemento", pila_apilar(pila, &dato1)==true);
    print_test("apilo otro elemento", pila_apilar(pila, &dato2)==true);
    print_test("apilo un 3er elemento", pila_apilar(pila, &dato3)==true);
    print_test("desapilo uno", pila_desapilar(pila)==&dato3);
    print_test("verifico si esta vacia", pila_esta_vacia(pila)==false);
    print_test("desapilo otro", pila_desapilar(pila)==&dato2);
    print_test("verifico si esta vacia", pila_esta_vacia(pila)==false);
    print_test("desapilo el ultimo", pila_desapilar(pila)==&dato1);
    print_test("verifico si esta vacia", pila_esta_vacia(pila)==true);
    print_test("apilo un elemento", pila_apilar(pila, &dato1)==true);
    print_test("veo el tope", pila_ver_tope(pila)==&dato1);
    print_test("verifico si esta vacia", pila_esta_vacia(pila)==false);

    free(dato1);
    free(dato2);
    free(dato3);

    print_test("destruyo la pila", pila!=NULL);
    pila_destruir(pila);
}


#define CANTIDAD_ELEMENTOS 10000 //siempre que sean mas de 10

static void prueba_pila_redimension(){
    printf("\nINICIO DE PRUEBAS DE DINAMICA DE LA PILA\n");

    pila_t *pila = pila_crear();
    print_test("creo la pila", pila!=NULL);

    size_t* arr=malloc(sizeof(size_t)*CANTIDAD_ELEMENTOS);
    if(arr==NULL){
        printf("\nFALLO DE MEMORIA\n");
        return;
    }

    print_test("verifico si esta vacia", pila_esta_vacia(pila)==true);
    
    for(size_t i=0; i<CANTIDAD_ELEMENTOS; i++)
        arr[i]= i+1;

    for(size_t i=0; i<CANTIDAD_ELEMENTOS; i++){
        print_test("apilo nuevo elemento", pila_apilar(pila, &arr[i])==true);

        if(i==10 || i==20 || i==CANTIDAD_ELEMENTOS - 2)
            print_test("verifico si esta vacia", pila_esta_vacia(pila)==false);
    }

    print_test("verifico si esta vacia", pila_esta_vacia(pila)==false);
    print_test("desapilo un elemento", pila_desapilar(pila)==(void*) &arr[CANTIDAD_ELEMENTOS-1]);
    print_test("verifico si esta vacia", pila_esta_vacia(pila)==false);
    print_test("desapilo otro elemento", pila_desapilar(pila)== (void*) &arr[CANTIDAD_ELEMENTOS-2]);
    print_test("apilo nuevamente",  pila_apilar(pila, &arr[CANTIDAD_ELEMENTOS-2])==true);
    print_test("veo el tope", pila_ver_tope(pila)== &arr[CANTIDAD_ELEMENTOS-2]);
    print_test("apilo nuevamente",  pila_apilar(pila, &arr[CANTIDAD_ELEMENTOS-1])==true);
    print_test("veo el tope", pila_ver_tope(pila)==&arr[CANTIDAD_ELEMENTOS-1]);


    for(size_t i=CANTIDAD_ELEMENTOS; i>0; i--){
        print_test("desapilo el ultimo", pila_desapilar(pila) ==  (void*) &arr[i-1]);

        if(i>3)
            print_test("veo el tope", pila_ver_tope(pila)== (void*) &arr[i-2]);

        if(i==10 || i==20 || i==CANTIDAD_ELEMENTOS - 2)
            print_test("verifico si esta vacia", pila_esta_vacia(pila)==false);
    }
    print_test("verifico si esta vacia", pila_esta_vacia(pila)==true);
    print_test("veo el tope", pila_ver_tope(pila)==NULL);

    free(arr);
    
    print_test("destruyo la pila", pila!=NULL);
    pila_destruir(pila);
}



void pruebas_pila_estudiante() {
    prueba_pila_vacia();
    prueba_pila_apilar_desapilar();
    prueba_pila_redimension();

}




#ifndef CORRECTOR  // Para que no dé conflicto con el main() del corrector.

int main(void) {
    pruebas_pila_estudiante();
    return failure_count() > 0;  // Indica si falló alguna prueba.
}

#endif
