package operation

import (
	"fmt"
	"math"
)

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
