package lexer

type TokenType string

const (
	// Number
	NUMBER TokenType = "NUMBER"

	// Operators
	PLUS TokenType = "PLUS"
	MINUS TokenType = "MINUS"
	MULTIPLY TokenType = "MULTIPLY"
	DIVIDE TokenType = "DIVIDE"
	MODULUS TokenType = "MODULUS"

	// Special
	PIPE_RESULT TokenType = "PIPE_RESULT"
	PIPED TokenType = "PIPED"
	OUTPUT_RESULT TokenType = "OUTPUT_RESULT"
	SLEEP TokenType = "SLEEP"
)