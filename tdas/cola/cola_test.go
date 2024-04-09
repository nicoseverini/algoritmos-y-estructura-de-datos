package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACola "tdas/cola"
	"testing"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestEncolarCola(t *testing.T) {
	colaInt := TDACola.CrearColaEnlazada[int]()
	require.True(t, colaInt.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaInt.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaInt.Desencolar() })

	colaInt.Encolar(3)
	require.False(t, colaInt.EstaVacia())
	require.Equal(t, 3, colaInt.VerPrimero(), "Si quiero ver el primer elemeto me devuelve el primer elemento de la cola")

	colaInt.Encolar(5)
	require.False(t, colaInt.EstaVacia())
	require.Equal(t, 3, colaInt.VerPrimero(), "Si quiero ver el primer elemeto me devuelve el primer elemento de la cola")

	colaStr := TDACola.CrearColaEnlazada[string]()
	require.True(t, colaStr.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaStr.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaStr.Desencolar() })

	colaStr.Encolar("test")
	require.False(t, colaStr.EstaVacia())
	require.Equal(t, "test", colaStr.VerPrimero(), "Si quiero ver el primer elemeto me devuelve el primer elemento de la cola")

	type ejemplo struct {
		texto string
	}
	ej := ejemplo{}
	colaStruct := TDACola.CrearColaEnlazada[ejemplo]()
	require.True(t, colaStruct.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaStruct.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaStruct.Desencolar() })

	colaStruct.Encolar(ej)
	require.Equal(t, ej, colaStruct.VerPrimero(), "Si quiero ver el primer elemeto me devuelve el primer elemento de la cola")
}

func TestDesencolarCola(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	cola.Encolar(1)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 1, cola.Desencolar(), "Si desencolo una cola me devuelve el primer elemento de la cola")
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	cola.Encolar(2)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 2, cola.VerPrimero(), "Si quiero ver el primer elemeto me devuelve el primer elemento de la cola")
	cola.Encolar(3)
	require.Equal(t, 2, cola.Desencolar(), "Si desencolo una cola me devuelve el primer elemento de la cola")
	require.Equal(t, 3, cola.VerPrimero(), "Si quiero ver el primer elemeto me devuelve el primer elemento de la cola")

	cola.Encolar(4)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 3, cola.Desencolar(), "Si desencolo una cola me devuelve el primer elemento de la cola")
	require.Equal(t, 4, cola.VerPrimero(), "Si quiero ver el primer elemeto me devuelve el primer elemento de la cola")
	require.False(t, cola.EstaVacia())
	require.Equal(t, 4, cola.Desencolar(), "Si desencolo una cola me devuelve el primer elemento de la cola")
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	colaStr := TDACola.CrearColaEnlazada[string]()
	require.True(t, colaStr.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaStr.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaStr.Desencolar() })

	colaStr.Encolar("a")
	require.False(t, colaStr.EstaVacia())
	require.Equal(t, "a", colaStr.VerPrimero(), "Si quiero ver el primer elemeto me devuelve el primer elemento de la cola")
	require.Equal(t, "a", colaStr.Desencolar(), "Si desencolo una cola me devuelve el primer elemento de la cola")
	require.True(t, colaStr.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaStr.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaStr.Desencolar() })

	colaStr.Encolar("b")
	require.False(t, colaStr.EstaVacia())
	require.Equal(t, "b", colaStr.VerPrimero(), "Si quiero ver el primer elemeto me devuelve el primer elemento de la cola")
	require.Equal(t, "b", colaStr.Desencolar(), "Si desencolo una cola me devuelve el primer elemento de la cola")
	require.True(t, colaStr.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaStr.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaStr.Desencolar() })
}

func TestVolumenCola(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	for i := range 100 {
		cola.Encolar(i)
		require.False(t, cola.EstaVacia())
		require.Equal(t, i, cola.VerPrimero(), "No importa el tamaño de la cola si quiero ver el primero me devuelve el primer elemento de la cola")

		require.Equal(t, i, cola.Desencolar(), "No importa el tamaño de la cola si desencolo me devuelve el primer elemento de la cola")

		require.True(t, cola.EstaVacia())
		require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
		require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	}

	volumenPila := 1000
	for i := range volumenPila {
		cola.Encolar(i + 1)
		require.Equal(t, 1, cola.VerPrimero(), "No importa el tamaño de la cola si quiero ver el primero me devuelve el primer elemento de la cola")
	}

	require.False(t, cola.EstaVacia())
	require.Equal(t, 1, cola.VerPrimero(), "No importa el tamaño de la cola si quiero ver el primero me devuelve el primer elemento de la cola")

	for i := range volumenPila {
		require.Equal(t, i+1, cola.Desencolar(), "No importa el tamaño de la cola si desencolo me devuelve el primer elemento de la cola")
	}

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

}
