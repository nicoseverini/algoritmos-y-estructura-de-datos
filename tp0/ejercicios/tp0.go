package ejercicios

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	aux := *x
	*x = *y
	*y = aux
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) == 0 {
		return -1
	} else {
		valorMax := vector[0]
		indiceMax := 0

		for i := range vector {
			if vector[i] > valorMax {
				valorMax = vector[i]
				indiceMax = i
			}
		}

		return indiceMax
	}
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	largo1 := len(vector1)
	largo2 := len(vector2)

	minLargo := largo1
	if largo2 < minLargo {
		minLargo = largo2
	}

	for i := 0; i < minLargo; i++ {
		if vector1[i] < vector2[i] {
			return -1
		} else if vector1[i] > vector2[i] {
			return 1
		}
	}

	if largo1 < largo2 {
		return -1
	} else if largo1 > largo2 {
		return 1
	} else {
		return 0
	}

}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	var slice []int = vector
	for i := len(vector); i > 1; i-- {
		indexMax := Maximo(slice[:i])
		if indexMax != (i - 1) {
			Swap(&slice[indexMax], &slice[i-1])
		}
	}
}

func SumaAux(vector []int, indice int) int {
	if indice == len(vector) {
		return 0
	}

	return vector[indice] + SumaAux(vector, indice+1)
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	return SumaAux(vector, 0)
}

func CapicuaAux(cadena string, indice int) bool {
	if indice == len(cadena) {
		return true
	}

	if cadena[indice] != cadena[len(cadena)-indice-1] {
		return false
	}

	return CapicuaAux(cadena, indice+1)
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func EsCadenaCapicua(cadena string) bool {
	return CapicuaAux(cadena, 0)
}
