package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func hello(done chan bool) {  
    fmt.Println("Hello world goroutine")
    done <- true
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

    fmt.Println(x, y, x + y)

    done := make(chan bool)
    go hello(done)
    <-done
    fmt.Println("main function")
}
