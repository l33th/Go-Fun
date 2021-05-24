package main

import (
	"fmt"
    "net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, _ *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}
