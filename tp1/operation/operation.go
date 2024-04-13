package operation

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"tdas/pila"
)

const (
	ADDITION       = "+"
	SUBTRACTION    = "-"
	MULTIPLICATION = "*"
	DIVISION       = "/"
	SQRT           = "sqrt"
	RAISE          = "^"
	LOGARITHM      = "log"
	TERNARY        = "?"
)

func operate(operator string, stack pila.Pila[int64], stackQuantity *int) (int64, error) {
	if operator == SQRT {
		if *stackQuantity < 1 {
			return 0, fmt.Errorf("stack quantity cannot be less than 1")
		}
		operand1 := stack.Desapilar()
		*stackQuantity--

		if operand1 < 0 {
			return 0, fmt.Errorf("operand out of range")
		}
		return int64(math.Sqrt(float64(operand1))), nil
	} else if operator == TERNARY {
		if *stackQuantity < 3 {
			return 0, fmt.Errorf("stack quantity cannot be less than 3")
		}
		operand3 := stack.Desapilar()
		*stackQuantity--
		operand2 := stack.Desapilar()
		*stackQuantity--
		operand1 := stack.Desapilar()
		*stackQuantity--

		if operand1 != 0 {
			return operand2, nil
		} else {
			return operand3, nil
		}
	} else {
		if *stackQuantity < 2 {
			return 0, fmt.Errorf("stack quantity cannot be less than 2")
		}

		operand2 := stack.Desapilar()
		*stackQuantity--
		operand1 := stack.Desapilar()
		*stackQuantity--

		switch operator {
		case ADDITION:
			return operand1 + operand2, nil
		case SUBTRACTION:
			return operand1 - operand2, nil
		case MULTIPLICATION:
			return operand1 * operand2, nil
		case DIVISION:
			if operand2 == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			return operand1 / operand2, nil
		case RAISE:
			if operand2 < 0 {
				return 0, fmt.Errorf("power with negative exponent")
			}
			return int64(math.Pow(float64(operand1), float64(operand2))), nil
		case LOGARITHM:
			if operand2 < 2 {
				return 0, fmt.Errorf("base of log less than 2")
			}
			return int64(math.Log(float64(operand1)) / math.Log(float64(operand2))), nil
		}
	}
	return 0, fmt.Errorf("there is no operator")
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

func CalculateExpression(expression string) (int64, error) {
	instructions := strings.Fields(expression)
	stack := pila.CrearPilaDinamica[int64]()
	var stackQuantity int

	for _, instruction := range instructions {
		if isOperator(instruction) {
			if stack.EstaVacia() {
				return 0, fmt.Errorf("the stack is empty")
			}

			var result int64
			var err error

			result, err = operate(instruction, stack, &stackQuantity)
			if err != nil {
				return 0, fmt.Errorf("problems when operating")
			}

			stack.Apilar(result)
			stackQuantity++

		} else {
			number, err := strconv.ParseInt(instruction, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("problems when converting number to int64")
			}

			stack.Apilar(number)
			stackQuantity++
		}
	}

	if stackQuantity != 1 {
		return 0, fmt.Errorf("the content of my stack should have only the result")
	}

	return stack.Desapilar(), nil
}
