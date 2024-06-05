package main

import (
	"bytes"
	"testing"
)

///The Buffer type from the bytes package implements the Writer interface, because it has the method Write(p []byte) (n int, err error).
//So we'll use it in our test to send in as our Writer and then we can check what was written to it after we invoke Greet

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
