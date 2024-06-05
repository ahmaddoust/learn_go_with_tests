package main

import (
	"fmt"
	"io"
	"os"
)

// /From this we can infer that os.Stdout implements io.Writer; Printf passes os.Stdout to Fprintf which expects an io.Writer.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
