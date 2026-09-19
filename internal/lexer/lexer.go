package lexer

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Tokenize(data []byte) [][]Token {
	var programTokens [][]Token

	source := string(data)

	lines := strings.Split(source, "\n")

	for lineNum, line := range lines {
		var lineTokens []Token

		segments := strings.Split(string(line), " ")
		
		for _, segment := range segments {
			token := isNumber(segment, lineNum)

			lineTokens = append(lineTokens, token)
		}

		programTokens = append(programTokens, lineTokens)
	}

	return programTokens
}

func isNumber(char string, lineNum int) Token {
	_, err := strconv.ParseFloat(char, 64)

	if err == nil {
		return Token {
			Type: NUMBER,
			Value: char,
		}
	}

	return isOperator(char, lineNum)
}

func isOperator(char string, lineNum int) Token {
	var token Token
	switch char {
	case "+":
		token = Token {
			Type: PLUS,
			Value: char,
		}
	case "-":
		token = Token {
			Type: MINUS,
			Value: char,
		}
	case "*":
		token = Token {
			Type: MULTIPLY,
			Value: char,
		}
	case "/":
		token = Token {
			Type: DIVIDE,
			Value: char,
		}
	case "%":
		token = Token {
			Type: MODULUS,
			Value: char,
		}
	default:
		token = isSpecial(char, lineNum)
	}

	return token
}

func isSpecial(char string, lineNum int) Token {
	var token Token
	switch char {
	case "|":
		token = Token {
			Type: PIPE_RESULT,
			Value: char,
		}
	case "!":
		token = Token {
			Type: OUTPUT_RESULT,
			Value: char,
		}
	case "?":
		token = Token {
			Type: PIPED,
			Value: char,
		}
	case "&":
		token = Token {
			Type: SLEEP,
			Value: char,
		}

	case "@":
		token = Token {
			Type: INPUT,
			Value: char,
		}
	default:
		fmt.Println("Unrecognised character:", char, "on line", lineNum)
		os.Exit(1)
	}

	return token
}