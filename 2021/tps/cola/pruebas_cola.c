#include "cola.h"
#include "testing.h"

#include <stdio.h>
#include <stdlib.h>

static void prueba_cola_vacia(){

    printf("\nINICIO DE PRUEBAS CON COLA VACIA\n");

    cola_t* cola=cola_crear();
    print_test("creo la cola", cola!=NULL);

    print_test("veo si la cola esta vacia", cola_esta_vacia(cola)==true);
    print_test("veo si tiene algo para desencolar", cola_desencolar(cola)==NULL);

    print_test("destruyo la cola", cola!=NULL);
    cola_destruir(cola, NULL);
}

static void prueba_cola_encolar_desencolar(){

    printf("\nINICIO DE PRUEBAS ENCOLANDO Y DESENCOLANDO\n");

    cola_t* cola=cola_crear();
    print_test("creo la cola", cola!=NULL);

    //creo elementos para encolar y los uso
    int* dato1=malloc(sizeof(int));
    if(dato1==NULL)
    {
        printf("\nFALLO DE MEMORIA\n");
        return;
    }
    *dato1 =  45;

    float* dato2=malloc(sizeof(float));
    if(dato2==NULL)
    {
        free(dato1);
        printf("\nFALLO DE MEMORIA\n");
        return;
    }
    *dato2 = 234.576f;

    char* dato3=malloc(sizeof(char));
    if(dato2==NULL)
    {
        free(dato1);
        free(dato2);
        printf("\nFALLO DE MEMORIA\n");
        return;
    }
    *dato3='A';

    print_test("encolo un elemento", cola_encolar (cola, dato1)==true);
    print_test("desencolo elemento", cola_desencolar(cola)==dato1);
    print_test("ver si cola vuelve a estar vacia", cola_esta_vacia(cola));
    print_test("encolo nuevamente el elemento", cola_encolar (cola, dato1)==true);
    print_test("encolo otro elemento", cola_encolar (cola, dato2)==true);
    print_test("encolo un 3er elemento", cola_encolar (cola, dato3)==true);
    print_test("desencolo elemento", cola_desencolar(cola)==dato1);
    print_test("me fijo si cola esta vacia", cola_esta_vacia(cola)==false);
    print_test("desencolo elemento", cola_desencolar(cola)==dato2);
    print_test("me fijo si cola esta vacia", cola_esta_vacia(cola)==false);
    print_test("desencolo elemento", cola_desencolar(cola)==dato3);
    print_test("me fijo si cola esta vacia", cola_esta_vacia(cola)==true);
    print_test("encolo un elemento", cola_encolar (cola, dato1)==true);
    print_test("veo tope", cola_ver_primero(cola)==dato1);
    print_test("me fijo si cola esta vacia", cola_esta_vacia(cola)==false);
    print_test("desencolo elemento", cola_desencolar(cola)==dato1);

    free(dato1);
    free(dato2);
    free(dato3);

    print_test("destruyo la cola", cola!=NULL);
    cola_destruir(cola, NULL);
}
#define CANTIDAD_DATOS 100

static void prueba_destruir_cola(){

    printf("\nINICIO PRUEBAS DE DESTRUCCION DE DATOS EN LA COLA\n");

    cola_t* arreglo_datos[CANTIDAD_DATOS];
    for(int i=0; i<CANTIDAD_DATOS; i++){
    
        arreglo_datos[i]=cola_crear();
        print_test("creo un dato", arreglo_datos[i]!=NULL);
        if(arreglo_datos[i]==NULL){

            for(; i>=0; --i){
                cola_destruir(arreglo_datos[i], (void (*)(void *)) cola_destruir);
            }
                return;
        }
    }
    cola_t* cola=cola_crear();
    print_test("creo la cola", cola!=NULL);
    for(size_t i=0; i<CANTIDAD_DATOS; i++){
    
        print_test("encolo un dato", cola_encolar(cola, (cola_t*) arreglo_datos[i])==true);
        print_test("me fijo si cola esta vacia", cola_esta_vacia(cola)==false);
    }
    cola_destruir(cola, (void (*)(void *)) cola_destruir);
}

void pruebas_cola_estudiante(){
    prueba_cola_vacia();
    prueba_cola_encolar_desencolar();
    prueba_destruir_cola();
}

#ifndef CORRECTOR  // Para que no dé conflicto con el main() del corrector.

int main(void) {
    pruebas_cola_estudiante();
    return failure_count() > 0;  // Indica si falló alguna prueba.
}

#endif
