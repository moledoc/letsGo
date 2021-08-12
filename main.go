package main

import (
	"fmt"

	helloworld "github.com/moledoc/letsgo/HelloWorld"
	loops "github.com/moledoc/letsgo/Loops"
)

func main() {
	// HelloWorld
	var greeting = helloworld.HelloWorld("moledoc")
	fmt.Println(greeting)
	fmt.Println(helloworld.HelloWorldDifferentAssign("moledoc"))
	//Loops
	endpoint := 10
	fmt.Printf("Summing all values from 0 to %d results in %v\n", endpoint, loops.ForLoop(endpoint))
}
