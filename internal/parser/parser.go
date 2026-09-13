package parser

import (
	"fmt"

	"github.com/greetingsForAlek/MathematicProgram/internal/lexer"
)

func Parse(tokens [][]lexer.Token) {
	for _, line := range tokens {
		parseLine(line)
	}
}

func parseLine(line []lexer.Token) error {
	if len(line) == 0 {
		return nil
	}

	for _, token := range line {
		fmt.Println(token)
	}

	return nil
}