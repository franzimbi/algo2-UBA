#include <stdbool.h>
#include <stdio.h>

bool es_magico_recursivo(int* arreglo, size_t inicio, size_t fin)
{
    if(inicio>fin)
        return false;

    size_t medio=(inicio+fin)/2;
    if(arreglo[medio]==medio)
        return true;
    else{
        return es_magico_recursivo(arreglo, inicio, medio-1) || es_magico_recursivo(arreglo, medio+1, fin);
    }
}

int main(void)
{
    int vec[]= {4,4,6,1,8,5,3,2,9};

    printf("%s\n", es_magico_recursivo(vec, 0, 8) ? "es magico" : "no es magico");

    return 0;
}
