package main

import "fmt"

type datos struct {
	numero *int
	nombre string
}

func main() {
	edadManuekl := 2

	p1 := datos{nombre: "manolo", numero: &edadManuekl}

	p2 := datos{nombre: "flavio", numero: p1.numero}
	*p2.numero += 1

	fmt.Println(*p1.numero)
	fmt.Println(*p2.numero)

}
