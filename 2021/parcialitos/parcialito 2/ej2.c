/* 
Escribir una función que reciba una palabra y un abecedario (una cadena de caracteres con cada letra del abecedario una
pegada a la otra) e indique si con el abecedario dado se puede formar dicha palabra. Indicar y justificar la complejidad de la
función implementada. Ejemplos de uso:
se_puede_formar("filo", "aeiflko") => True, ya que todas las letras de “filo” se encuentran en el abecedario
se_puede_formar("grafo", "aeiflko") => False, ya que ni “g” ni “r” forman parte del abecedario
*/

#include <stdbool.h>
#include <stdio.h>
#include "hash.h"

bool se_puede_formar(char* palabra, char* abc){
    hash_t* diccionario = hash_crear(NULL);
    char aux[2];
    aux[1] = '\0';

    for(size_t i=0; i<strlen(abc); i++){ //O(m)
        aux[0] = abc[i];
        hash_guardar(diccionario, aux, NULL); //O(1)
    }
    for(size_t i=0; i<strlen(palabra); i++){ //O(n)
        aux[0] = palabra[i];
        if(!hash_pertenece(diccionario, aux)){ //O(1)
            hash_destruir(diccionario); 
            return false;
        }
    }
    hash_destruir(diccionario);
    return true;
}

//LA FUNCION IMPLEMENTADA CON UN HASH ES DE ORDEN o(n+m), SIENDO n EL LARGO DE LA PALABRA Y m EL LARGO DEL ABECEDARIO.  YA QUE SE REALIZAN DOS CICLOS FOR,
// UNO PARA GUARDAR LAS LETRAS DEL ABECEDARIO, Y OTRO PARA VER SI TODAS LAS LETRAS DE LA PALABRA ESTAN EN EL DICCIONARIO. DENTRO DE LOS FOR TODO EN O(1).
