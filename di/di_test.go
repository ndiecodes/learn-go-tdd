package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Ndie")
	got := buffer.String()
	want := "Hello Ndie"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
