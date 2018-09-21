package hello_test

import (
	"testing"

	"github.com/neild/hello"
)

func TestHello(t *testing.T) {
	if got, want := hello.Hello(), "Hello, world!"; got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
