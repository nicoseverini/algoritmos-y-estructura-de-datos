package cola

type nodo[T any] struct {
	valor       T
	proximoNodo *nodo[T]
}

type colaEnlazada[T any] struct {
	primerNodo *nodo[T]
	ultimoNodo *nodo[T]
	cantidad   int
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := colaEnlazada[T]{
		primerNodo: nil,
		ultimoNodo: nil,
		cantidad:   0,
	}

	return &cola
}

func (cola colaEnlazada[T]) EstaVacia() bool {
	return cola.cantidad == 0
}

func (cola colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	return cola.primerNodo.valor
}

func crearNodo[T any](valor T) *nodo[T] {
	return &nodo[T]{
		valor:       valor,
		proximoNodo: nil,
	}
}

func (cola *colaEnlazada[T]) Encolar(valor T) {
	nuevoNodo := crearNodo(valor)

	if cola.EstaVacia() {
		cola.primerNodo = nuevoNodo
	} else {
		cola.ultimoNodo.proximoNodo = nuevoNodo
	}
	cola.ultimoNodo = nuevoNodo

	cola.cantidad++
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	valor := cola.primerNodo.valor
	if cola.cantidad == 1 {
		cola.primerNodo = nil
		cola.ultimoNodo = nil
	} else {
		cola.primerNodo = cola.primerNodo.proximoNodo
	}

	cola.cantidad--
	return valor
}
