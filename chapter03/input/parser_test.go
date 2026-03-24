package input_test

import (
	"testing"

	"github.com/sony-nurdianto/test_driven_development/chapter03/calculator"
	"github.com/sony-nurdianto/test_driven_development/chapter03/input"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessExpression(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		// Arrange
		expr := "2 + 3"
		operator := "+"
		operands := []float64{2.0, 3.0}
		expectedResult := "2 + 3 = 5.5"
		engine := input.NewMockOperationProcessor(t)
		validator := input.NewMockValidationHelper(t)
		parser := input.NewParser(engine, validator)

		validator.On("CheckInput", operator, operands).Return(nil).Once()
		engine.On("ProcessOperation", calculator.Operation{
			Expression: expr,
			Operator:   operator,
			Operands:   operands,
		}).Return(&expectedResult, nil).Once()

		// Act
		result, err := parser.ProcessExpression(expr)

		// Assert
		require.Nil(t, err)
		require.NotNil(t, result)
		assert.Contains(t, *result, expectedResult)
		assert.Contains(t, *result, expr)
		validator.AssertExpectations(t)
		engine.AssertExpectations(t)
	})
}
