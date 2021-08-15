package main

import (
	"fmt"
	"log"

	helloworld "github.com/moledoc/letsgo/HelloWorld"
	ifelse "github.com/moledoc/letsgo/IfElse"
	loops "github.com/moledoc/letsgo/Loops"
	sqrt "github.com/moledoc/letsgo/Sqrt"
)

func main() {
	// HelloWorld
	var greeting = helloworld.HelloWorld("moledoc")
	fmt.Println(greeting)
	fmt.Println(helloworld.HelloWorldDifferentAssign("moledoc"))
	//Loops
	endpoint := 10
	fmt.Printf("Summing all values from 0 to %d results in %v\n", endpoint, loops.ForLoop(endpoint))

	sqrtVal, err := sqrt.Sqrt(16)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("My Sqrt(16)=%v\n", sqrtVal)

	fmt.Printf("When is Saturday? %v\n", ifelse.SwitchV2())
}
