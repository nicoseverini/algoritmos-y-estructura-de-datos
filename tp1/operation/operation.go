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

type Operation interface {
	Operate(operands []int64) (int64, error)
}

type Addition struct{}

func (Addition) Operate(operands []int64) (int64, error) {
	if len(operands) < 2 {
		return 0, fmt.Errorf("insufficient operands for addition")
	}
	return operands[0] + operands[1], nil
}

type Subtraction struct{}

func (Subtraction) Operate(operands []int64) (int64, error) {
	if len(operands) < 2 {
		return 0, fmt.Errorf("insufficient operands for subtraction")
	}
	return operands[0] - operands[1], nil
}

type Multiplication struct{}

func (Multiplication) Operate(operands []int64) (int64, error) {
	if len(operands) < 2 {
		return 0, fmt.Errorf("insufficient operands for multiplication")
	}
	return operands[0] * operands[1], nil
}

type Division struct{}

func (Division) Operate(operands []int64) (int64, error) {
	if len(operands) < 2 {
		return 0, fmt.Errorf("insufficient operands for division")
	}
	if operands[1] == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return operands[0] / operands[1], nil
}

type Power struct{}

func (Power) Operate(operands []int64) (int64, error) {
	if len(operands) < 2 {
		return 0, fmt.Errorf("insufficient operands for power operation")
	}
	if operands[1] < 0 {
		return 0, fmt.Errorf("power with negative exponent")
	}
	return int64(math.Pow(float64(operands[0]), float64(operands[1]))), nil
}

type Logarithm struct{}

func (Logarithm) Operate(operands []int64) (int64, error) {
	if len(operands) < 2 {
		return 0, fmt.Errorf("insufficient operands for logarithm operation")
	}
	if operands[1] < 2 {
		return 0, fmt.Errorf("base of logarithm less than 2")
	}
	return int64(math.Log(float64(operands[0])) / math.Log(float64(operands[1]))), nil
}

type SquareRoot struct{}

func (SquareRoot) Operate(operands []int64) (int64, error) {
	if len(operands) < 1 {
		return 0, fmt.Errorf("insufficient operands for square root operation")
	}
	if operands[0] < 0 {
		return 0, fmt.Errorf("operand out of range for square root")
	}
	return int64(math.Sqrt(float64(operands[0]))), nil
}

type Ternary struct{}

func (Ternary) Operate(operands []int64) (int64, error) {
	if len(operands) < 3 {
		return 0, fmt.Errorf("insufficient operands for ternary operation")
	}
	if operands[0] != 0 {
		return operands[1], nil
	}
	return operands[2], nil
}

func isOperator(instruction string) bool {
	operators := "+-*/sqrt^log?" //TODO hacela global
	return strings.Contains(operators, instruction)
}

// TODO: cambiar nombre func
func operate(instruction string, stack pila.Pila[int64]) (int64, error) {
	operands := make([]int64, 0)
	var op Operation
	if instruction == SQRT {
		for i := 0; i < 1; i++ {
			if stack.EstaVacia() {
				return 0, fmt.Errorf("insufficient operands for operator %s", instruction)
			}
			operand := stack.Desapilar()
			operands = append([]int64{operand}, operands...)
		}

		op = SquareRoot{}
	} else if instruction == TERNARY {
		for i := 0; i < 3; i++ {
			if stack.EstaVacia() {
				return 0, fmt.Errorf("insufficient operands for operator %s", instruction)
			}
			operand := stack.Desapilar()
			operands = append([]int64{operand}, operands...)
		}
		op = Ternary{}
	} else {
		for i := 0; i < 2; i++ {
			if stack.EstaVacia() {
				return 0, fmt.Errorf("insufficient operands for operator %s", instruction)
			}
			operand := stack.Desapilar()
			operands = append([]int64{operand}, operands...)
		}
		switch instruction {
		case ADDITION:
			op = Addition{}
		case SUBTRACTION:
			op = Subtraction{}
		case MULTIPLICATION:
			op = Multiplication{}
		case DIVISION:
			op = Division{}
		case RAISE:
			op = Power{}
		case LOGARITHM:
			op = Logarithm{}
		}
	}
	result, err := op.Operate(operands)

	return result, err
}

func CalculateExpression(expression string) (int64, error) {
	instructions := strings.Fields(expression)
	stack := pila.CrearPilaDinamica[int64]()

	for _, instruction := range instructions {
		if isOperator(instruction) {
			if stack.EstaVacia() {
				return 0, fmt.Errorf("the stack is empty")
			}

			result, err := operate(instruction, stack)
			if err != nil {
				return 0, err
			}

			stack.Apilar(result)
		} else {
			number, err := strconv.ParseInt(instruction, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("problems when converting number to int64")
			}

			stack.Apilar(number)
		}
	}
	result := stack.Desapilar()

	if !stack.EstaVacia() {
		return 0, fmt.Errorf("the content of my stack should have only the result")
	}

	return result, nil
}
