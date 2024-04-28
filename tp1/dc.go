package main

import (
	"bufio"
	"dc/operation"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		line := s.Text()

		var result int64
		var err error

		result, err = operation.CalculateExpression(line)
		if err != nil {
			fmt.Println("ERROR")
		} else {
			fmt.Println(result)
		}
	}
}
