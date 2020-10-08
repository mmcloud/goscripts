package hello

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
}

// Dependency injection
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Matt") //send through buff

	got := buffer.String()
	want := "Hello, Matt"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
