#include <stdio.h>
#include <string.h>

int main(void){

    char arr[50];
    
    while(!feof(stdin)){
        fgets(arr, 50, stdin);
        printf("%s\n", arr);
    }
    fprintf(stderr, "se cerro? parece q si\n");
    return 0;
}