package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Alex")

	got := buffer.String()
	want := "Hello, Alex"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
