package evaluator

import (
	"fmt"
	"strconv"

	"github.com/greetingsForAlek/MathematicProgram/internal/parser"
)

func Evaluate(operations []parser.Operation) ([]float64, error) {	
	results := make([]float64, 0, len(operations))

	var pipedResult float64
	hasPipedResult := false

	for _, operation := range operations {
		result, err := evaluateOperation(
			operation,
			pipedResult,
			hasPipedResult,
		)

		if err != nil {
			return nil, err
		}

		results = append(results, result)

		if operation.Pipe {
			pipedResult = result
			hasPipedResult = true
		} else {
			hasPipedResult = false
		}
	}

	return results, nil
}

func evaluateOperation(operation parser.Operation, previousResult float64, hasPipedResult bool) (float64, error) {
	left, err := resolveOperand(
		operation.Left,
		previousResult,
		hasPipedResult,
	)
	if err != nil {
		return 0, err
	}

	right, err := resolveOperand(
		operation.Right,
		previousResult,
		hasPipedResult,
	)

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
		return left / right, nil

	default:
		return 0, fmt.Errorf("unknown operator: %s", operation.Opertator)
	}
}

func resolveOperand(operand parser.Operand, previousResult float64, hasPipedResult bool) (float64, error) {
	if operand.Previous {
		if !hasPipedResult {
			return 0, fmt.Errorf(
				"cannot use '?' because the previous line did not pipe a result",
			)
		}

		return previousResult, nil
	}

	value, err := strconv.ParseFloat(operand.Value, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number %s", operand.Value)
	}

	return value, nil
}