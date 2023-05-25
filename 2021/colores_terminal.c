#include <stdio.h>

#define COLOR "\033[0;31;43m" 
#define NORMAL "\033[0m"

int main(void)
{
    printf(COLOR "hola" NORMAL"\n");

    return 0;
}
