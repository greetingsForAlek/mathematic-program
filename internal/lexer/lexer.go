package lexer

import (
	"fmt"
	"strings"
)

func Tokenize(data []byte) {
	source := string(data)

	lines := strings.Split(source, "\n")

	for _, line := range lines {
		segments := strings.Split(string(line), " ")
		
		for _, segment := range segments {
			fmt.Println(segment)
		}
	}
}