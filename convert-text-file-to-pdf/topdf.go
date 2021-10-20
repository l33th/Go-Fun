package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	file := "test.txt"

	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("%s file not found", file)
	}
	fmt.Println(content)
}
