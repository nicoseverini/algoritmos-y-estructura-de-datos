package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	ej "tp0/ejercicios"
)

const archivo1 = "archivo1.in"
const archivo2 = "archivo2.in"

func guardarArchivoEnArreglo(nombreArchivo string) []int {
	archivo, err := os.Open(nombreArchivo)

	if err != nil {
		fmt.Printf("Error %v al abrur el archivo %s", nombreArchivo, err)
		return nil
	}
	defer archivo.Close()

	s := bufio.NewScanner(archivo)
	var arreglo []int

	for s.Scan() {
		line := s.Text()
		num, err := strconv.Atoi(line)

		if err != nil {
			fmt.Println(err)
			return nil
		}

		arreglo = append(arreglo, num)
	}
	err = s.Err()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return arreglo
}

func calcularMayorYOrdenar(arreglo1 []int, arreglo2 []int) []int {
	i := ej.Comparar(arreglo1, arreglo2)

	arregloMayor := arreglo1
	if i == -1 {
		arregloMayor = arreglo2
	}

	ej.Seleccion(arregloMayor)

	return arregloMayor
}

func imprimirArreglo(arreglo []int) {
	for i := range arreglo {
		fmt.Println(arreglo[i])
	}
}

func main() {
	var arregloArchivo1 []int
	var arregloArchivo2 []int

	arregloArchivo1 = guardarArchivoEnArreglo(archivo1)
	arregloArchivo2 = guardarArchivoEnArreglo(archivo2)

	arregloMayorOrdenado := calcularMayorYOrdenar(arregloArchivo1, arregloArchivo2)

	imprimirArreglo(arregloMayorOrdenado)

}
