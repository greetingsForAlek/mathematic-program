package evaluator

import (
	"fmt"
	"strconv"

	"github.com/greetingsForAlek/MathematicProgram/internal/parser"
)

func Evaluate(operations []parser.Operation) ([]float64, error) {
	results := make([]float64, 0, len(operations))

	var previousResult float64

	for _, operation := range operations {
		result, err := evaluateOperation(operation, previousResult)

		if err != nil {
			return nil, err
		}

		results = append(results, result)
		previousResult = result
	}

	return results, nil
}

func evaluateOperation(operation parser.Operation, previousResult float64) (float64, error) {
	left, err := resolveOperand(operation.Left, previousResult)
	if err != nil {
		return 0, err
	}

	right, err := resolveOperand(operation.Right, previousResult)
	if err != nil {
		return 0, err
	}

	switch operation.Opertator {
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		if right == 0 {
			return 0, fmt.Errorf("Division by Zero")
		}

		return left / right, nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", operation.Opertator)
	}
}

func resolveOperand(operand parser.Operand, previousResult float64) (float64, error) {
	if operand.Previous {
		return previousResult, nil
	}

	value, err := strconv.ParseFloat(operand.Value, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %s", operand.Value)
	}

	return value, nil
}