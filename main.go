package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/tour/reader"

	errors "github.com/moledoc/letsgo/Errors"
	helloworld "github.com/moledoc/letsgo/HelloWorld"
	ifelse "github.com/moledoc/letsgo/IfElse"
	loops "github.com/moledoc/letsgo/Loops"
	moretypes "github.com/moledoc/letsgo/MoreTypes"
	rdr "github.com/moledoc/letsgo/Readers"
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

	//// Map
	fmt.Println(moretypes.WordCount("The quick brown fox jumped over the lazy dog."))

	//// Function as value
	// moretypes.FuncAsVal uses x=7 and y=3
	Multiply := func(x float32, y float32) float32 {
		return x * y
	}
	// naive power function
	Power := func(x float32, y float32) float32 {
		var power float32 = 1
		for y > 0 {
			power *= x
			y--
		}
		return power
	}
	fmt.Println(moretypes.FuncAsVal(Multiply))
	fmt.Println(moretypes.FuncAsVal(Power))

	//// Function closure
	subtractor := moretypes.Subtractr()
	fmt.Println(subtractor(5)) // expected -5
	fmt.Println(subtractor(2)) // expected -7

	fib := moretypes.Fibonacci()
	for i := 0; i < 10; i++ {
		if i == 9 {
			fmt.Printf("%v\n", fib())
			break
		}
		fmt.Printf("%v,", fib())
	}

	// Errors
	fmt.Println(errors.Sqrt(2))
	fmt.Println(errors.Sqrt(-2))

	// Readers
	fmt.Println("Go tour Reader exercise validation")
	reader.Validate(rdr.MyReader{})
}
