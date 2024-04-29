package operation

import (
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
	OPERATORS      = "+-*/sqrt^log?"
)

func isOperator(instruction string) bool {
	return strings.Contains(OPERATORS, instruction)
}

func getOperands(quantity int, stack pila.Pila[int64]) ([]int64, error) {
	operands := make([]int64, 0)
	for i := 0; i < quantity; i++ {
		if stack.EstaVacia() {
			return nil, InsufficientOperandsError{}
		}
		operand := stack.Desapilar()
		operands = append([]int64{operand}, operands...)
	}
	return operands, nil
}

func operate(instruction string, stack pila.Pila[int64]) (int64, error) {
	var op Operation
	var operands []int64
	var result int64
	var err error

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
	case SQRT:
		op = SquareRoot{}
	case TERNARY:
		op = Ternary{}
	default:
		return 0, UnknownOperatorError{Operator: instruction}
	}

	operands, err = getOperands(op.OperandQuantity(), stack)
	if err != nil {
		return 0, err
	}

	result, err = op.Operate(operands)

	return result, err
}

func CalculateExpression(expression string) (int64, error) {
	instructions := strings.Fields(expression)
	stack := pila.CrearPilaDinamica[int64]()

	for _, instruction := range instructions {
		if isOperator(instruction) {
			if stack.EstaVacia() {
				return 0, StackEmptyError{}
			}

			result, err := operate(instruction, stack)
			if err != nil {
				return 0, err
			}

			stack.Apilar(result)
		} else {
			number, err := strconv.ParseInt(instruction, 10, 64)
			if err != nil {
				return 0, ConvertStringError{}
			}

			stack.Apilar(number)
		}
	}
	result := stack.Desapilar()

	if !stack.EstaVacia() {
		return 0, StackResultError{}
	}

	return result, nil
}
