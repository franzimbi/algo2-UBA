/*
Para una implementación completa de Lista (como se implementó en el trabajo de TDA Lista), implementar una función
void* borrar_ultimo_de_lista(lista_t* lista), que quite y devuelva el último elemento de la lista. En caso de que
la lista esté vacía devuelve NULL. Indicar y justificar la complejidad de la función implementada.
*/
#include "lista.h"

void* borrar_ultimo_de_lista(lista_t* lista){
    if(lista_esta_vacia(lista)) return NULL;
    lista_iter_t* iterador = lista_iter_crear(lista); //O(1)
    if(iterador==NULL) return NULL;

    for(size_t i=0; i<lista_largo(lista) - 1; i++){ //O(n)
        lista_iter_avanzar(iterador); //O(1)
    }
    void* dato = lista_iter_borrar(iterador); //O(1)
    lista_iter_destruir(iterador); //O(1)

    return dato;
}

// LA FUNCION DESARROLLADA ES DE COMPLEJIDAD O(n), YA QUE SE TRATA DE UNA ITERACION DEL TAMANO DE LA LISTA Y EL RESTO ES DE ORDEN O(1).



/*
int main(void){

    lista_t* lista = lista_crear();

    lista_insertar_ultimo(lista, "1");
    lista_insertar_ultimo(lista, "2");
    lista_insertar_ultimo(lista, "3");
    lista_insertar_ultimo(lista, "4");

    printf("borrado: %s\n", (char*) borrar_ultimo_de_lista(lista));
    printf("cant: %zu\n", lista_largo(lista));
    lista_destruir(lista, NULL);
    return 0;
}
*/