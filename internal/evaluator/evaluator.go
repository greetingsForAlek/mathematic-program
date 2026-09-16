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
	return 10.0, nil // i lied i need to do smt first
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