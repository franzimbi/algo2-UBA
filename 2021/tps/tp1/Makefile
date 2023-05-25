PROGRAM=fixcol
CC=gcc
CFLAGS=-g -O2 -std=c99 -Wall -Werror -pedantic -Wformat=2 -Wshadow -Wpointer-arith -Wunreachable-code -Wconversion -Wno-sign-conversion -Wbad-function-cast 

fixcol.o: fixcol.h fixcol.c
	$(CC) $(CFLAGS) -c fixcol.c

main.o: main.c fixcol.h fixcol.c
	$(CC) $(CFLAGS) -c main.c

build: main.o fixcol.o
	$(CC) $(CFLAGS) main.o fixcol.o -o $(PROGRAM)