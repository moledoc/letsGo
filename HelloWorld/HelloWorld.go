package helloworld

import (
	"fmt"
	"math/cmplx"
)

// package level variable assigning
var (
	ToBe   bool       = true
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Greet given `name`
func HelloWorld(name string) string {
	return "Hello " + name
}

// Give two greetings: say hello to first argument and say bye to the second
func HelloWorldMultiReturn(fstName string, sndName string) (string, string) {
	var fstGreet = "Hello " + fstName
	var sndGreet = "Bye " + sndName
	return fstGreet, sndGreet
}

// Demostrate different ways to assign values
func HelloWorldDifferentAssign(name string) (string, uint64) {
	// immutable
	const nameComma = "Hello"
	// function level var keyword; can also be at package level
	var fstGreet = "Hello " + name
	// keyword := can only be used at func level
	sndGreet := fstGreet + ", nice day!"
	// declaring variable with its zero value, in this case ""
	var thrdGreet string
	thrdGreet = sndGreet + " Well, bye now"
	// multiple variable declarations on the same line; each can have one value
	var frthGreet, number = thrdGreet, 2 + MaxInt
	if ToBe {
		fmt.Printf("Printing out different formats: %v (type = %T); default value of MaxInt = %q\n", z, z, MaxInt)
	}
	return frthGreet, number
}
