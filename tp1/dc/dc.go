package dc

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"tdas/pila"
)

// declarar const TODO
func operate(operator string, operand1 int64, operand2 int64) (int64, error) {
	switch operator {
	case "+":
		return operand1 + operand2, nil
	case "-":
		return operand1 - operand2, nil
	case "*":
		return operand1 * operand2, nil
	case "/":
		if operand2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}

		return operand1 / operand2, nil
	case "sqrt":
		if operand1 < 0 {
			return 0, fmt.Errorf("operand out of range")
		}

		return int64(math.Sqrt(float64(operand1))), nil
	case "^":
		return int64(math.Pow(float64(operand1), float64(operand2))), nil
	case "log":
		if operand1 < 2 {
			return 0, fmt.Errorf("base of log less than 2")
		}
		return int64(math.Log(float64(operand2)) / math.Log(float64(operand1))), nil
	case "?":
		if operand1 != 0 {
			return operand2, nil
		} else {
			//TODO
		}
	}
}

func isOperator(instruction string) bool {
	return instruction == "+" ||
		instruction == "-" ||
		instruction == "*" ||
		instruction == "/" ||
		instruction == "sqrt" ||
		instruction == "^" ||
		instruction == "log" ||
		instruction == "?"
}

func calculateExpression(expression string) (int64, error) {
	instructions := strings.Fields(expression)
	stack := pila.CrearPilaDinamica[int64]()
	var stackQuantity int

	for _, instruction := range instructions {
		if isOperator(instruction) {
			if stackQuantity < 1 {
				return 0, fmt.Errorf("requires at least one operand to operate")
			}
			//analizar casos bordes como el ternario y raiz TODO
			//aplicar la idea de waldo(mientas voy viendo las instrucciones voy desapilando lo que quiero) TODO

			operand2 := stack.Desapilar()
			operand1 := stack.Desapilar()

			var result int64
			var err error

			result, err = operate(instruction, operand1, operand2)
			if err != nil {
				return 0, fmt.Errorf("problems when operating")
			}
			stackQuantity--
			stack.Apilar(result)

		} else {
			number, err := strconv.ParseInt(instruction, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("problems when converting to int")
			}

			stackQuantity++
			stack.Apilar(number)
		}
	}

	if stackQuantity != 1 {
		return 0, fmt.Errorf("the content of my stack should have only the result")
	}

	return stack.Desapilar(), nil
}
