package lexer

import (
	"fmt"
	"strings"
)

func Tokenize(data []byte) {
	source := string(data)

	lines := strings.Split(source, "\n")

	for line := range lines {
		fmt.Println(line)
	}
}