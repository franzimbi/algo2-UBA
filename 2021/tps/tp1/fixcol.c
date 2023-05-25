#include "fixcol.h"

void fixcol(long cant_caracteres, FILE* archivo){
    int c;
    long i=0;
    while((c=fgetc(archivo))!=EOF){
        if(i==0 && c=='\n')
            c=fgetc(archivo);
        if(c=='\n')
            i=-1;
        fputc(c, stdout);
        i++;
        if(i==cant_caracteres){
            i=0;
            fputc('\n', stdout);
        }

    }
}
//hola'\n'como