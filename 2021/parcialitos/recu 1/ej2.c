#include <stdio.h>

int cantidad_unos(int* arr, int inicio, int fin){
    if(inicio==fin)
        return arr[fin]==1 ? 1 : 0;
    int medio = (inicio+fin)/2;
    return  cantidad_unos(arr, inicio, medio) + cantidad_unos(arr, medio+1, fin);

}
int main(void){
    int arr[5] = {0, 0, 0, 0, 1};

    printf("%i\n", cantidad_unos(arr, 0, 5));
    return 0;
}


//aplicando teorema maestro: A=2 (cantidad de llamados recursivos), B=2 (proporcion con la que llamamos recursivamente, osea dividimos en 2), C=0 ( O(1) es lo q cuesta sumar)
// log2 (2) = 1 > C = 0 -> O(N^(log b (a))) -> O(n)