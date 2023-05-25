#include "pila.h"
#include <stdbool.h>
#include <string.h>
#include <stdio.h>

char apertura[] = {'(', '[', '{'};
char cerrado[] = { ')', ']', '}'};

bool son_pareja(char c, char d){
    if(c=='(' && d==')')
        return true;
    if(c=='[' && d==']')
        return true;
    if(c=='{' && d=='}')
        return true;
    return false;
}

bool balanceo(char * cadena){
    pila_t* pila = pila_crear();
    if(pila == NULL)
        return false;
    int pos_str =0;

    int tam = strlen(cadena); //O(n)
    //pila_apilar(pila, cadena[pos_str++]);
    while(pos_str < tam){ //O(n)
        char c = cadena[pos_str];
        if(c=='(' || c=='[' || c=='{'){
            pila_apilar(pila, c);
            pos_str++;
        }
        else if(pila_esta_vacia(pila)){
            pila_destruir(pila);
            return false;
        }
        else{
                char d = pila_desapilar(pila);
                printf("%c %c\n", d, c);
                if(!son_pareja(d, c)){
                    pila_desapilar(pila);
                    return false;
                }
            }
    }
    pila_destruir(pila);
    return true;
}

int main (void){

    char str[] = "[()]";
    if(balanceo(str)==true)
        printf("true\n");
    else
        printf("false\n");
    return 0;
}

// como recorre toda la cadena para calcular el tamano y para el while: O(n+n) -> O(n) 

/* NO LO PUDE HACER ANDAR, PERO LA IDEA ESTA BIEN. ENCOLO LOS  ( [ { Y CUANDO ME ENCUENTRO CON ALGUNO Q CIERRE LO COMPARO, SI ES OPUESTO SIGO, SINO RETURN FALSE.
SI LA PILA NO QUEDA VACIA TAMBIEN ES FALSE.
SINO RETURN TRUE 
*/