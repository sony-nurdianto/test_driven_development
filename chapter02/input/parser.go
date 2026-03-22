package input

import "github.com/sonynurdianto/test_driven_development/chapter02/calculator"

type Parser struct {
	engine    *calculator.Engine
	validator *Validator
}

func (p *Parser) ProcessExpression(expr string) (*string, error) {
	return nil, nil
}
