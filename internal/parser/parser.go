package parser

import (
	"fmt"
	"strconv"

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
	if line[0].Type == lexer.SLEEP {
		if len(line) != 2 {
			return Operation{}, fmt.Errorf("Invalid sleep operation.")
		}

		if line[1].Type != lexer.NUMBER {
			return Operation{}, fmt.Errorf("Sleep duration must be a number.")
		}

		duration, err := strconv.Atoi(line[1].Value)
		if err != nil {
			return Operation{}, fmt.Errorf("Invalid sleep duration: %s", line[1].Value)
		}

		return Operation {
			Sleep: true,
			SleepTime: duration,
		}, nil
	}

	if len(line) < 3 {
		return Operation{}, fmt.Errorf("Invalid Operation")
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
	output := false

	for _, token := range line[3:] {
		switch token.Type {
		case lexer.PIPE_RESULT:
			pipe = true
		
		case lexer.OUTPUT_RESULT:
			output = true

		default:
			return Operation{}, fmt.Errorf("Unexpected token: %s", token.Value)
		}
	}

	return Operation {
		Left: left,
		Opertator: operator,
		Right: right,
		Pipe: pipe,
		Output: output,
		Sleep: false,
		SleepTime: 0,
	}, nil
}