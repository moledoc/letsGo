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

func TestHelloWorldMultiReturn(t *testing.T) {
	fstGreet, sndGreet := helloworld.HelloWorldMultiReturn("moledoc", "molecurrent")
	if fstGreet != "Hello moledoc" || sndGreet != "Bye molecurrent" {
		t.Fatal("HelloWorldMultiReturn function does not greet given string values correctly")
	}
}
