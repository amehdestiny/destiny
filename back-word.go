package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "the sky is blue"

	words := strings.Fields(s)

	for i := len(words) - 1; i >= 0; i-- {
		fmt.Print(words[i])

		if i > 0 {
			fmt.Print(" ")
		}
	}

	fmt.Println()
}
