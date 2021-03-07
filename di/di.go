package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello %s", name)
}

func MyGreethandler(w http.ResponseWriter, r *http.Request) {
		Greet(w, "Ndie")
}

func main() {
	// Greet(os.Stdout, "Ndie")
	http.ListenAndServe(":8000", http.HandlerFunc(MyGreethandler))
}
