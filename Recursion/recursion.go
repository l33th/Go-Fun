package main

import (
	"fmt"
)

func main() {
	n := factorial(7)
	fmt.Println(n)
	m := loopFactorial(7)
	fmt.Println(m)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFactorial(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}