#include "fixcol.h"
#include <stdlib.h>

#define MENSAJE_ERROR_PARAMETROS "Error: Cantidad erronea de parametros\n"
#define MENSAJE_ERROR_ARCHIVOS "Error: archivo fuente inaccesible\n"
#define MENSAJE_ERROR_MEMORIA "Error: falla de memoria"

long validar_cantidad_caracteres(char* argumento){
    char* invalid_ptr = NULL;
    long numero = strtol(argumento, &invalid_ptr, 10);
    if(*invalid_ptr!='\0' || numero<=0)
        return -1;
    return numero;
}

int main(int argc, char* argv[]){
    if((argc!=2 && argc!=3) || validar_cantidad_caracteres(argv[1])==-1){
        fprintf(stderr, MENSAJE_ERROR_PARAMETROS );
        return EXIT_FAILURE;
    }
    long argumento_nro = validar_cantidad_caracteres(argv[1]);
    FILE* archivo_entrada = argc == 3 ? fopen(argv[2], "r") : stdin;
    if(archivo_entrada==NULL){
        fprintf(stderr, MENSAJE_ERROR_ARCHIVOS);
        return EXIT_FAILURE;
    }
    fixcol(argumento_nro, archivo_entrada);
    if(archivo_entrada!=stdin)
        fclose(archivo_entrada);
    
    return EXIT_SUCCESS;
}
