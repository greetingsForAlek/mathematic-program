package main

import (
	"os"
	"fmt"
)

func main() {
	data, err := os.ReadFile("./main.mp")
	if (err != nil) {
		fmt.Println("Error reading file: ", err)
	}

	fmt.Println(data)
}