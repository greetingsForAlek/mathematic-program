package evaluator

import (
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
	return 10.0, nil // 10.0 is not a permanent value imma fix that in the next commit
}