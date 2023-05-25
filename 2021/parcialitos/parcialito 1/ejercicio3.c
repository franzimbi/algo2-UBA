#include <stdio.h>

/* A un array ordenado se lo rotó k posiciones. Implementar una función por división y conquista
que permita encontrar la cantidad k de rotaciones que se le aplicó al array. Justificar el orden
del algoritmo. */

int rotaciones_arreglo_rec(int arr[], size_t ini, size_t fin){
    if(ini==fin)
        return 0;

    size_t medio=(ini+fin)/2; //o(1)
    int mitad_izq = rotaciones_arreglo_rec(arr, ini, medio);
    int mitad_der = rotaciones_arreglo_rec(arr, medio+1, fin);
    int mitad = arr[medio] > arr[medio+1] ? medio+1 : 0;

    return mitad_izq + mitad_der + mitad; // o(1)
}

int rotaciones_arreglo(int arr[], size_t n){
    
    if(n==0)
        return 0;

    return rotaciones_arreglo_rec(arr, 0, n-1);
}

/* utilizo el teorema maestro para calcular el orden de este algoritmo, siendo  este T(n)= AT (n/B) + O(n^C):

    -> A es la cantidad de llamados recursivos que hago de la funcion: en mi caso 2, usados en mitad_izq y mitad_der.
    -> B es la proporcion con la que llamamos recursivamente: en mi caso 2, ya que divido en dos el arreglo.
    -> C se calcula teniendo en cuenta lo que cuesta partir y volver a juntar: en mi caso 0, ya que lo no recursivo  es O(1).

    T(n) = 2T(n/2) + O(1) ->  log b (a) = log 2 (2) = 1 > C

    entonces: T(n) = N^(log b (a))          ->    T(n) = O(n)
