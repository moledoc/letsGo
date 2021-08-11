package helloworld_test

import (
	"testing"

	helloworld "github.com/moledoc/letsgo/HelloWorld"
)

func TestHelloWorld(t *testing.T) {
	if helloworld.HelloWorld("moledoc") != "Hello moledoc" {
		t.Fatal("HelloWorld function does not greet given string value")
	}
}
