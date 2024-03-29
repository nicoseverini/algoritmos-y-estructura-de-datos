package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := pilaDinamica[T]{
		datos: make([]T, 5),
		cantidad: 0,
	}

	return &pila
}

func (pila pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	} else {
		return pila.datos[pila.cantidad-1]
	}
}

func (pila *pilaDinamica[T]) Apilar(valor T) {
	if pila.cantidad == len(pila.datos) {
		nuevosDatos := make([]T,(len(pila.datos) * 2))
		copy(nuevosDatos, pila.datos)
		pila.datos = nuevosDatos
	}
	pila.datos[pila.cantidad] = valor
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	} else {
		valor := pila.datos[pila.cantidad-1]
		pila.cantidad--
		if (pila.cantidad * 4) <= len(pila.datos) {
			nuevosDatos := make([]T,(len(pila.datos) / 2))
			copy(nuevosDatos, pila.datos)
			pila.datos = nuevosDatos
		}
		return valor
	}
}
