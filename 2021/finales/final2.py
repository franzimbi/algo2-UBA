from platform import java_ver
from tkinter.messagebox import NO
import grafo
from collections import deque
import heapq
import math


# FINAL 19/07/22

#ejercicio 1
def dijkistra(grafo, origen):
    dist = {}
    padre = {}
    for v in grafo:
        dist[v]=float("inf")
    dist[origen] = 0
    padre[origen] = None
    q = heapq.heapify([(0, origen)])
    while len(q) != 0:
        v = heapq.heappush(q)
        for w in grafo.adyacentes(v):
            if dist[w] > dist[v] + grafo.peso(v, w):
                dist[w] = dist[v] + grafo.peso(v, w)
                padre[w] = v
                heapq.heappop(q,(dist[w], w))
    return dist, padre


def centralidad_armonica(grafo, origen):
    cent = 0
    for v in grafo:
        if v ==  origen:
            continue
        distancia, padre = dijkistra(grafo, origen)
        if padre[v] == None:
            continue
        cent += 1/distancia[v]
    return cent


# Escribir una función que reciba un arreglo de N números reales e informe los índices de dos números que estén entre sí a mínima distancia. Justificar el orden. 



def indices_a_minima_dist(arreglo, cantidad):
    nro1 = 0
    nro2 = 1
    dist = math.sqrt((arreglo[nro1]-arreglo[nro2])**2)
    for i in range(cantidad):
        for w in range(cantidad):
            if math.sqrt((arreglo[nro1]-arreglo[nro2])**2) < dist:
                dist = math.sqrt((arreglo[nro1]-arreglo[nro2])**2)
                nro1 = i
                nro2 = w
    return nro1, nro2

