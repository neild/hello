package hello_test

import (
	"testing"

	"github.com/neild/hello"
)

func TestHello(t *testing.T) {
	if got, want := hello.Hello(hello.English), "Hello, world!"; got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
