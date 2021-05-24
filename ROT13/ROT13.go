package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check command line arguments
	if len(os.Args) != 2 {
		fmt.Println(`ROT13 ("rotate by 13 places", sometimes hyphenated ROT-13)`)
		fmt.Println("Usage: " + os.Args[0] + " <string>")
		os.Exit(1)
	}

	input := os.Args[1]

	mapped := strings.Map(rot13, input)
	fmt.Println(input)
	fmt.Println(mapped)
}

func rot13(rne rune) rune {
	if rne >= 'a' && rne <= 'z' {
		// Rotate lowercase letters 13 places.
		if rne >= 'm' {
			return rne - 13
		} else {
			return rne + 13
		}
	} else if rne >= 'A' && rne <= 'Z' {
		// Rotate uppercase letters 13 places.
		if rne >= 'M' {
			return rne - 13
		} else {
			return rne + 13
		}
	}
	// Do nothing.
	return rne
}
