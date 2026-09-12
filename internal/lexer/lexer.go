package lexer

import "fmt"

func Tokenize(data []byte) {
	source := string(data)
	fmt.Println(source)
}