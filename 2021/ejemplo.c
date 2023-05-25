#include <stdio.h>
#include <stdlib.h>

int main(){
    char* aux = malloc(sizeof(char) * 10);
    size_t tam = 10;
    size_t leidos = getline(&aux, &tam, stdin);
    aux[leidos-2] = '\0';
    printf("%s", aux);
    return 0;
}