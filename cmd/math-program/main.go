package main

import (
	"os"
	"fmt"

	"github.com/greetingsForAlek/MathematicProgram/internal/lexer"
	"github.com/greetingsForAlek/MathematicProgram/internal/parser"
)

func main() {
	data, err := os.ReadFile("./main.mp")
	if (err != nil) {
		fmt.Println("Error reading file: ", err)
	}

	tokens := lexer.Tokenize(data)
	parser.Parse(tokens)
}