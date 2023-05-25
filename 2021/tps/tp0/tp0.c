#include "tp0.h"

/* *****************************************************************
*                     FUNCIONES A COMPLETAR                       *
*         (ver en tp0.h la documentación de cada función)         *
* *****************************************************************/

void swap(int *x, int *y)
{
    int aux;

    aux= (*x);
    *x = (*y);
    *y = aux;
}


int maximo(int vector[], int n)
{
    if(n==0)
        return -1;

    int posicion_del_mayor=0;
    int mayor_aux = vector[0];

    for(int i=0; i<n; i++)
    {
        if(vector[i]>mayor_aux)
        {
            posicion_del_mayor=i;
            mayor_aux = vector[i];
        }
    }

    return posicion_del_mayor;
}


int comparar(int vector1[], int n1, int vector2[], int n2)
{
    int n=n2;

    if(n1<n2)
        n=n1;

    for(int i=0; i<n; i++)
    {
        if(vector1[i]<vector2[i])
            return -1;
        if(vector1[i]>vector2[i])
            return 1;
    }
    if(n1==n2)
        return 0;
    
    if(n1<n2)
        return -1;
    else{
        return 1;
    }
}

void seleccion(int v[], int n)
{
    for(int i=n-1; n>0; i--)
    {
        int p = maximo(v, n);
        swap(&v[i], &v[p]);
        n--;
    }
}
