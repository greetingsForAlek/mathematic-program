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
		os.Exit(1)
	}

	tokens := lexer.Tokenize(data)
	operations, err := parser.Parse(tokens)
	if (err != nil) {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(operations)
}