package calculator

import (
	"fmt"

	"github.com/sony-nurdianto/test_driven_development/chapter03/format"
)

type Operation struct {
	Expression string
	Operator   string
	Operands   []float64
}

type Engine struct {
	expectedLength  int
	validOperations map[string]func(x, y float64) (*float64, error)
}

func NewEngine() *Engine {
	e := Engine{
		expectedLength:  2,
		validOperations: make(map[string]func(x float64, y float64) (*float64, error)),
	}
	e.validOperations["+"] = e.Add
	e.validOperations["-"] = e.Sub
	e.validOperations["/"] = e.Div
	e.validOperations["*"] = e.Mult
	return &e
}

func (e *Engine) GetNumOperands() int {
	return e.expectedLength
}

func (e *Engine) GetValidOperators() []string {
	var ops []string
	for o := range e.validOperations {
		ops = append(ops, o)
	}

	return ops
}

func (e *Engine) ProcessOperation(operation Operation) (*string, error) {
	f, ok := e.validOperations[operation.Operator]
	if !ok {
		err := fmt.Errorf("no operation for operator %s found", operation.Operator)
		return nil, format.Error(operation.Expression, err)
	}
	res, err := f(operation.Operands[0], operation.Operands[1])
	if err != nil {
		return nil, format.Error(operation.Expression, err)
	}
	fres := format.Result(operation.Expression, *res)
	return &fres, nil
}

func (e *Engine) Add(x, y float64) (*float64, error) {
	result := x + y
	return &result, nil
}

func (e *Engine) Sub(x, y float64) (*float64, error) {
	result := x - y
	return &result, nil
}

func (e *Engine) Mult(x, y float64) (*float64, error) {
	result := x * y
	return &result, nil
}

func (e *Engine) Div(x, y float64) (*float64, error) {
	if y == 0 {
		return nil, fmt.Errorf("cannot divide by zero")
	}
	result := x / y
	return &result, nil
}
