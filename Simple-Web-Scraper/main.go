package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Variables for the response and error
	var r *http.Response
	var err error

	// Request index.html from example.com
	r, err = http.Get("http://www.example.com/index.html")

	// Kill the program and print the error to the console if there is a problem accessing the server
	if err != nil {
		panic(err)
	}

	// Check the status code returned by the server
	if r.StatusCode == 200 {
		// The request was successful!
		var webPageContent []byte

		// We already know the size of the response is 1270
		var bodyLength int = 1270

		// Initialize the byte array to the size of the data
		webPageContent = make([]byte, bodyLength)

		// Read the data from the server
		r.Body.Read(webPageContent)

		// Open a writable file on your computer (create if it does not exist)
		var out *os.File
		out, err = os.OpenFile("index.html", os.O_CREATE|os.O_WRONLY, 0664)

		if err != nil {
			panic(err)
		}

		// Write the contents to a file
		out.Write(webPageContent)
		out.Close()
	} else {
		log.Fatal("Failed to retrieve the webpage.  Received status code: ", r.Status)
	}
}
