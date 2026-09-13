package parser

import (
	"fmt"

	"github.com/greetingsForAlek/MathematicProgram/internal/lexer"
)

func Parse(tokens [][]lexer.Token) {
	for _, line := range tokens {
		for _, token := range line {
			fmt.Println(token)
		}
	}
}