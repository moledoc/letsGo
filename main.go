package main

import (
	"fmt"
	"log"
	"strings"

	helloworld "github.com/moledoc/letsgo/HelloWorld"
	ifelse "github.com/moledoc/letsgo/IfElse"
	loops "github.com/moledoc/letsgo/Loops"
	moretypes "github.com/moledoc/letsgo/MoreTypes"
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

	// sqrt implementation
	sqrtVal, err := sqrt.Sqrt(16)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("My Sqrt(16)=%v\n", sqrtVal)

	// ifelse
	fmt.Printf("When is Saturday? %v\n", ifelse.SwitchV2())

	// moretypes
	//// pointers
	fmt.Println(moretypes.Pointer())
	//// structs
	fmt.Println(moretypes.StructAndPointer())
	//// Arrays
	fmt.Println(moretypes.Arrays())
	//// Slices
	fmt.Println(moretypes.Slice())
	fmt.Println(moretypes.SliceAsRef())
	var nrs, bools, sliceLiteral = moretypes.SliceLiteral()
	fmt.Printf("This is int slice %v, this is a boolean slice %v, this a slice literal %v with types %T\n", nrs, bools, sliceLiteral, sliceLiteral)

	board := moretypes.SliceMatrix()
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}
