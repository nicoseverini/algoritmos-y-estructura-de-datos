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
	var arry []int

	for s.Scan() {
		line := s.Text()
		num, err := strconv.Atoi(line)

		if err != nil {
			fmt.Println(err)
			return nil
		}

		arry = append(arry, num)
	}
	err = s.Err()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return arry
}

func calcularMayoryOrdenar(arry1 []int, arry2 []int) []int {
	i := ej.Comparar(arry1, arry2)
	var arryMayor []int
	if i == 1 {
		arryMayor = arry1
	} else if i == -1 {
		arryMayor = arry2
	} else if i == 0 {
		arryMayor = arry1
	}

	ej.Seleccion(arryMayor)

	return arryMayor
}

func imprimirArreglo(arry []int) {
	for i := range arry {
		fmt.Println(arry[i])
	}
}

func main() {
	var arryArchivo1 []int
	var arryArchivo2 []int

	arryArchivo1 = guardarArchivoEnArreglo(archivo1)
	arryArchivo2 = guardarArchivoEnArreglo(archivo2)

	arryMayorOrdenado := calcularMayoryOrdenar(arryArchivo1, arryArchivo2)

	imprimirArreglo(arryMayorOrdenado)

}
