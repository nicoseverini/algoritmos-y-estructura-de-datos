package operation

type Operation interface {
	Operate(operands []int64) (int64, error)
	OperandQuantity() int
}
