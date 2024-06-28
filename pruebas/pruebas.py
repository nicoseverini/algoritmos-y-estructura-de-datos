from grafo import Grafo
from pila import Pila
from collections import deque
from random import choice
import sys


def inicializar_grafo(ruta):
    grafo = Grafo(es_dirigido=True)
    vertices_visitados = set()
    with open(f'./{ruta}') as f:
        for linea in f:
            mensaje = linea.split()
            id1, id2 = mensaje[0], mensaje[1]
            if id1 not in vertices_visitados:
                grafo.agregar_vertice(id1)
            if id2 not in vertices_visitados:
                grafo.agregar_vertice(id2)

            if not grafo.estan_unidos(id1, id2):
                peso_arista = grafo.peso_arista(id1, id2)
                grafo.agregar_arista(id1, id2, peso=1)
    return grafo


def comandos_eleccion(parametros, grafo_delincuentes):
    comando = parametros[0]

    if comando == "mas_imp":
        if len(parametros) < 2:
            raise "Error en comandos"
        cantidad = int(parametros[1])


def main(datos):
    if len(datos) != 1:
        raise ("Error en parametros de entrada")
    ruta = datos[0]
    contenido_archivo2 = sys.stdin.read().split()

    grafo_delincuentes = inicializar_grafo(ruta)
    comandos_eleccion(contenido_archivo2, grafo_delincuentes)



if __name__ == "__main__":
    main(sys.argv[1:])