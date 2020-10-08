package main

import (
	"fmt"
	"io"
	"net/http"

	quoteV3 "rsc.io/quote/v3"
)

func Hello() string {
	return quoteV3.HelloV3()
}

func Proverb() string {
	return quoteV3.Concurrency()
}

func Greet(writer io.Writer, name string) {
	// fmt.Printf("Hello, %s", name)
	fmt.Fprintf(writer, "Hello, %s", name)
}

// send through http writer
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
