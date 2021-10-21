package main

import "fmt"

func main() {
	xValue := 0
	fmt.Println("xValue Address before: ", &xValue)
	fmt.Println("xValue Variable Before: ", xValue)

	foo(&xValue)
}

func foo(yValue *int) {
	fmt.Println("yValue Variable before: ", yValue)
	fmt.Println("yValue pointer before: ", *yValue)
}
