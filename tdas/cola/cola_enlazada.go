package cola

type nodo[T any] struct {
	valor       T
	proximoNodo *nodo[T]
}

type colaEnlazada[T any] struct {
	primerNodo *nodo[T]
	ultimoNodo *nodo[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])

	return cola
}

func (cola colaEnlazada[T]) EstaVacia() bool {
	return cola.primerNodo == nil && cola.ultimoNodo == nil
}

func (cola colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	return cola.primerNodo.valor
}

func crearNodo[T any](valor T) *nodo[T] {
	NodoCola := new(nodo[T])
	NodoCola.valor = valor

	return NodoCola
}

func (cola *colaEnlazada[T]) Encolar(valor T) {
	nuevoNodo := crearNodo(valor)

	if cola.EstaVacia() {
		cola.primerNodo = nuevoNodo
	} else {
		cola.ultimoNodo.proximoNodo = nuevoNodo
	}
	cola.ultimoNodo = nuevoNodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	valor := cola.primerNodo.valor

	cola.primerNodo = cola.primerNodo.proximoNodo
	if cola.primerNodo == nil {
		cola.ultimoNodo = nil
	}

	return valor
}
