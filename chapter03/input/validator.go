package input

type ValidationHelper interface {
	CheckInput(operator string, operands []float64) error
}

type Validator struct{}
