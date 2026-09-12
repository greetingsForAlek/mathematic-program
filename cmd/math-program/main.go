package main

import (
	"os"
	"fmt"

	"github.com/greetingsForAlek/MathematicProgram/internal/lexer"
)

func main() {
	data, err := os.ReadFile("./main.mp")
	if (err != nil) {
		fmt.Println("Error reading file: ", err)
	}

	lexer.Tokenize(data)
}