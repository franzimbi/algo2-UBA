from tdaGrafo import Grafo
import csv

_NOMBRE_CIUDAD = 0
_LATITUD_CIUDAD = 1
_LONGITUD_CIUDAD = 2
_CIUDAD_ORIGEN = 0
_CIUDAD_DESTINO = 1
_TIEMPO_CONEXION = 2


def CrearGrafoConPajek(archivo):
    f = open(archivo, 'r')
    #grafoDirigido = Grafo(True, True)
    grafoNoDirigido = Grafo(False, True)
    cantCiudades = f.readline()
    for i in range(int(cantCiudades)):
        v = f.readline()
        ciudad = v.split(",")
        # grafoDirigido.insertarVertice(
        #    ciudad[_NOMBRE_CIUDAD], (ciudad[_LATITUD_CIUDAD], ciudad[_LONGITUD_CIUDAD].rstrip()))
        grafoNoDirigido.insertarVertice(
            ciudad[_NOMBRE_CIUDAD], (ciudad[_LATITUD_CIUDAD], ciudad[_LONGITUD_CIUDAD].rstrip()))
    cantConexiones = f.readline()
    t = 0
    for i in range(int(cantConexiones)):
        e = f.readline()
        conexion = e.split(",")
        # grafoDirigido.insertarArista(
        #    conexion[_CIUDAD_ORIGEN], conexion[_CIUDAD_DESTINO], int(conexion[_TIEMPO_CONEXION].rstrip()))
        grafoNoDirigido.insertarArista(
            conexion[_CIUDAD_ORIGEN], conexion[_CIUDAD_DESTINO], int(conexion[_TIEMPO_CONEXION].rstrip()))
    f.close()
    return grafoNoDirigido


def CrearKMLCamino(nombreDocumento, nombreArchivo, grafo, camino):
    f = open(nombreArchivo, 'w')
    f.write("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
    f.write("<kml xmlns=\"http://earth.google.com/kml/2.1\">\n")
    f.write("\t<Document>\n")
    f.write("\t\t<name>" + nombreDocumento + "</name>\n\n")
    # f.write("\t\t<description>"+descripcion+"</description>\n\n")
    agregados = set()
    for i in camino:
        if i in agregados:
            continue
        agregados.add(i)
        f.write("\t\t<Placemark>\n\t\t\t<name>"+i+"</name>\n\t\t\t<Point>\n")
        long, lat = grafo.datoVertice(i)
        f.write("\t\t\t\t<coordinates>"+long+", "+lat +
                "</coordinates>\n\t\t\t</Point>\n\t\t</Placemark>\n")
    f.write("\n")
    for i in range(1, len(camino)):
        f.write("\t\t<Placemark>\n\t\t\t<LineString>\n")
        long1, lat1 = grafo.datoVertice(camino[i-1])
        long2, lat2 = grafo.datoVertice(camino[i])
        f.write("\t\t\t\t<coordinates>"+long1+", "+lat1 +
                " "+long2+", "+lat2+"</coordinates>\n\t\t\t</LineString>\n\t\t</Placemark>\n")
    f.write("\t</Document>\n</kml>\n")
    f.close()


def CrearGrafoTopologicoConCSV(archivo):
    f = open(archivo, 'r')
    csvReader = csv.reader(f)
    grafoTopologico = Grafo(True, False)
    for linea in csvReader:
        grafoTopologico.insertarArista(
            linea[_CIUDAD_ORIGEN], linea[_CIUDAD_DESTINO])
    f.close()
    return grafoTopologico


def CrearPajekArbolTendidoMinimo(grafo, aristas, nombreArchivo):
    f = open(nombreArchivo, 'w')
    f.write(str(len(grafo)) + "\n")
    for v in grafo:
        long, lat = grafo.datoVertice(v)
        f.write(v + "," + long + "," + lat + "\n")
    f.write(str(len(aristas)) + "\n")
    for e in aristas:
        origen, destino, peso = e
        f.write(origen + "," + destino + "," + str(peso) + "\n")
    f.close()


def FormatoDeSalidaDeCamino(camino):
    result = ""
    for c in range(len(camino)):
        result += camino[c]
        if c != len(camino)-1:
            result += " -> "
    return result
