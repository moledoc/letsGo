package main

import (
	"fmt"

	helloworld "github.com/moledoc/letsgo/HelloWorld"
)

func main() {
	var greeting = helloworld.HelloWorld("moledoc")
	fmt.Println(greeting)
}
