package operation

import (
	"math"
)

type Addition struct{}

func (Addition) OperandQuantity() int {
	return 2
}

func (Addition) Operate(operands []int64) (int64, error) {
	if len(operands) < 2 {
		return 0, InsufficientOperandsError{Operation: "addition"}
	}
	return operands[0] + operands[1], nil
}

type Subtraction struct{}

func (Subtraction) OperandQuantity() int {
	return 2
}

func (Subtraction) Operate(operands []int64) (int64, error) {
	if len(operands) < 2 {
		return 0, InsufficientOperandsError{Operation: "subtraction"}
	}
	return operands[0] - operands[1], nil
}

type Multiplication struct{}

func (Multiplication) OperandQuantity() int {
	return 2
}

func (Multiplication) Operate(operands []int64) (int64, error) {
	if len(operands) < 2 {
		return 0, InsufficientOperandsError{Operation: "multiplication"}
	}
	return operands[0] * operands[1], nil
}

type Division struct{}

func (Division) OperandQuantity() int {
	return 2
}

func (Division) Operate(operands []int64) (int64, error) {
	if len(operands) < 2 {
		return 0, InsufficientOperandsError{Operation: "division"}
	}
	if operands[1] == 0 {
		return 0, DivisionByZeroError{}
	}
	return operands[0] / operands[1], nil
}

type Power struct{}

func (Power) OperandQuantity() int {
	return 2
}

func (Power) Operate(operands []int64) (int64, error) {
	if len(operands) < 2 {
		return 0, InsufficientOperandsError{Operation: "power"}
	}
	if operands[1] < 0 {
		return 0, NegativeExponentError{}
	}
	return int64(math.Pow(float64(operands[0]), float64(operands[1]))), nil
}

type Logarithm struct{}

func (Logarithm) OperandQuantity() int {
	return 2
}

func (Logarithm) Operate(operands []int64) (int64, error) {
	if len(operands) < 2 {
		return 0, InsufficientOperandsError{Operation: "logarithm"}
	}
	if operands[1] < 2 {
		return 0, BaseLogarithmError{}
	}
	return int64(math.Log(float64(operands[0])) / math.Log(float64(operands[1]))), nil
}

type SquareRoot struct{}

func (SquareRoot) OperandQuantity() int {
	return 1
}

func (SquareRoot) Operate(operands []int64) (int64, error) {
	if len(operands) < 1 {
		return 0, InsufficientOperandsError{Operation: "square root"}
	}
	if operands[0] < 0 {
		return 0, SquareRootError{}
	}
	return int64(math.Sqrt(float64(operands[0]))), nil
}

type Ternary struct{}

func (Ternary) OperandQuantity() int {
	return 3
}

func (Ternary) Operate(operands []int64) (int64, error) {
	if len(operands) < 3 {
		return 0, InsufficientOperandsError{Operation: "ternary"}
	}
	if operands[0] != 0 {
		return operands[1], nil
	}
	return operands[2], nil
}
