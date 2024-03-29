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

func (p pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	} else {
		return p.datos[p.cantidad-1]
	}
}

func (p *pilaDinamica[T]) Apilar(valor T) {
	if p.cantidad == len(p.datos) {
		nuevosDatos := make([]T,(len(p.datos) * 2))
		copy(nuevosDatos, p.datos)
		p.datos = nuevosDatos
	}
	p.datos[p.cantidad] = valor
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	} else {
		valor := p.datos[p.cantidad-1]
		p.cantidad--
		if (p.cantidad * 4) <= len(p.datos) {
			nuevosDatos := make([]T,(len(p.datos) / 2))
			copy(nuevosDatos, p.datos)
			p.datos = nuevosDatos
		}
		return valor
	}
}
