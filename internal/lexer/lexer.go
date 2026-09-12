package lexer

import (
	"fmt"
	"strings"
)

func Tokenize(data []byte) {
	source := string(data)

	lines := strings.Split(source, "\n")

	for lineNum, line := range lines {
		segments := strings.Split(string(line), " ")
		
		for _, segment := range segments {
			isNumber(segment, lineNum)
		}
	}
}

func isNumber(char string, lineNum int) {
	if char == "0" || 
	char == "1" || 
	char == "2" || 
	char == "3" || 
	char == "4" || 
	char == "5" || 
	char == "6" || 
	char == "7" || 
	char == "8" || 
	char == "9" {
		fmt.Println("char is a number.")
	} else {
		isOperator(char, lineNum)
	}
}

func isOperator(char string, lineNum int) {
	switch char {
	case "+":
		fmt.Println("Addition Operator")
	case "-":
		fmt.Println("Subtraction Operator")
	case "*":
		fmt.Println("Multiplication Operator")
	case "/":
		fmt.Println("Division Operator")
	case "%":
		fmt.Println("Modulus Operator")
	default:
		isSpecial(char, lineNum)
	}
}

func isSpecial(char string, lineNum int) {
	switch char {
	case "|":
		fmt.Println("Pipe Result")
	case "!":
		fmt.Println("Output Result")
	case "?":
		fmt.Println("Piped Result")
	case "&":
		fmt.Println("Sleep")
	default:
		fmt.Println("Unrecognised character:", char, "on line", lineNum)
	}
}