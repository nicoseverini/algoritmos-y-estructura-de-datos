package main

import (
	//"bufio"
	"dc/operation"
	"fmt"
	//"os"
)

func main() {
	//s := bufio.NewScanner(os.Stdin)

	//for s.Scan() {
	//	line := s.Text()
	line := "4 4 +" //TODO: esto es solo para dubuguear, borrar despues
	var result int64
	var err error

	result, err = operation.CalculateExpression(line)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(result)
	}
	//}
}
