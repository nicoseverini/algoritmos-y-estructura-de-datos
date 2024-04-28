package operation

import "fmt"

type InsufficientOperandsError struct {
	Operation string
}

func (e InsufficientOperandsError) Error() string {
	return fmt.Sprintf("insufficient operands for %s operation", e.Operation)
}

type DivisionByZeroError struct{}

func (e DivisionByZeroError) Error() string {
	return "division by zero"
}

type NegativeExponentError struct{}

func (e NegativeExponentError) Error() string {
	return "power with negative exponent"
}

type BaseLogarithmError struct{}

func (e BaseLogarithmError) Error() string {
	return "base of logarithm less than 2"
}

type SquareRootError struct{}

func (e SquareRootError) Error() string {
	return "operand out of range for square root"
}

type UnknownOperatorError struct {
	Operator string
}

func (e UnknownOperatorError) Error() string {
	return fmt.Sprintf("unknown operator %s", e.Operator)
}

type StackEmptyError struct{}

func (e StackEmptyError) Error() string {
	return "the stack is empty"
}

type ConvertStringError struct{}

func (e ConvertStringError) Error() string {
	return "problems when converting string to int64"
}

type StackResultError struct{}

func (e StackResultError) Error() string {
	return "the content of my stack should have only the result"
}
