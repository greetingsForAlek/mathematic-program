package parser

import (
	"fmt"

	"github.com/greetingsForAlek/MathematicProgram/internal/lexer"
)

func Parse(lines [][]lexer.Token) ([]Operation, error) {
	operations := make([]Operation, 0, len(lines))

	for _, line := range lines {
		operation, err := parseLine(line)

		if err != nil {
			return nil, err
		}

		operations = append(operations, operation)
	}

	return operations, nil
}

func parseLine(line []lexer.Token) (Operation, error) {
	if len(line) < 3 {
		return Operation{}, fmt.Errorf("Invalid Operation.")
	}

	left := Operand {
		Value: line[0].Value,
	}

	if line[0].Type == lexer.PIPED {
		left.Previous = true
	}

	operator := line[1].Value

	right := Operand {
		Value: line[2].Value,
	}

	if line[2].Type == lexer.PIPED {
		right.Previous = true
	}

	pipe := false
	if len(line) > 3 {
		if line[3].Type == lexer.PIPE_RESULT {
			pipe = true
		}
	}

	return Operation {
		Left: left,
		Opertator: operator,
		Right: right,
		Pipe: pipe,
	}, nil
}