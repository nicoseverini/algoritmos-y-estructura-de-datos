package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestApilarPila(t *testing.T) {
	pilaInt := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pilaInt.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaInt.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaInt.Desapilar() })
	pilaInt.Apilar(3)
	require.False(t, pilaInt.EstaVacia())
	require.Equal(t, 3, pilaInt.VerTope(), "Si apilo un elemento el tope me devuelve dicho elemento apilado")
	pilaInt.Apilar(33)
	require.False(t, pilaInt.EstaVacia())
	require.Equal(t, 33, pilaInt.VerTope(), "Si apilo un elemento el tope me devuelve dicho elemento apilado")


	pilaStr := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pilaStr.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaStr.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaStr.Desapilar() })

	pilaStr.Apilar("Test")
	require.False(t, pilaStr.EstaVacia())
	require.Equal(t, "Test", pilaStr.VerTope(), "Si apilo un elemento el tope me devuelve dicho elemento apilado")
	
	type ejemplo struct {
		texto string
	}
	ej := ejemplo{}
	pilaStruct := TDAPila.CrearPilaDinamica[ejemplo]()
	require.True(t, pilaStruct.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaStruct.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaStruct.Desapilar() })
	pilaStruct.Apilar(ej)
	require.Equal(t, ej, pilaStruct.VerTope(), "Si apilo un elemento el tope me devuelve dicho elemento apilado")
}

func TestDesapilarPila(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

	pila.Apilar(1)
	require.Equal(t, 1, pila.VerTope(), "Si apile el tope tiene que devolverme el ultimo elemento apilado")
	pila.Desapilar()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

	pila.Apilar(2)
	pila.Apilar(3)
	require.Equal(t, 3, pila.VerTope(), "Si apile el tope tiene que devolverme el ultimo elemento apilado")
	pila.Desapilar()
	require.Equal(t, 2, pila.VerTope(), "si desapile el tope tiene que devolverme el nuevo ultimo elemento de la pila")

	pila.Apilar(4)
	pila.Desapilar()
	pila.Desapilar()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

	pilaStr := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pilaStr.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaStr.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaStr.Desapilar() })

	pilaStr.Apilar("a")
	pilaStr.Desapilar()
	require.True(t, pilaStr.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaStr.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaStr.Desapilar() })
	
	pilaStr.Apilar("b")
	require.Equal(t, "b", pilaStr.VerTope(), "Si apile el tope tiene que devolverme el ultimo elemento apilado")

}

func TestVolumenPila(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

	volumenPila := 1000
	for i := range volumenPila {
		pila.Apilar(i+1)
		require.Equal(t, i+1, pila.VerTope(), "No importa el largo de mi pila si amplie el tope me tiene que devolver el ultimo elemento de esta")
	}

	require.False(t, pila.EstaVacia())
	require.Equal(t, volumenPila, pila.VerTope(), "No importa el largo de mi pila el tope me tiene que devolver el ultimo elemento de esta")
	
	for i := range volumenPila {
		require.Equal(t, volumenPila-i, pila.VerTope(), "No importa el largo de mi pila si desapile el tope me tiene que devolver el nuevo ultimo elemento de esta")
		pila.Desapilar()
	}

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}