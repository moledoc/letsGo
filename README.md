# LetsGo

## Synopsis

This repository contains simple examples/functions of GoLang,
that are the result of me learning the programming language.
It is meant to be testing and showing Go and it's concepts/ideas/possibilities.

## Contents

* HelloWorld - variable related examples (passing args, func type, returns, variable assignments)
  * HelloWorld - returns greeting to given string value
  * HelloWorldMultiReturn - returns two different greetings
* Loops - different cycles in GoLang
  * ForLoop - typical `for` loop
  * BasicallyWhileLoop - `go` version of a `while` loop
* IfElse - `if-else` statement in GoLang
  * BasicIf - typical `if` statement
  * AssignInIf - `if` statement with declaring a variable for that corresponding `if` statement
  * AssignInIfElse - `if-else` statement with declaring a variable for that corresponding `if-else` statement
  * Switch - `switch` statemnt in GoLang/shorter way of writing `if-else` block
  * SwitchV2 - `switch` statemnt can have functions in its statements.
* Sqrt - just cool/interesting exercise in 'A tour of Go' about calculating the square root of a number
  * Sqrt - calc square root of a number with error < 0.001 with Newton's method.
* MoreTypes - multiline program level variable, struct, pointer, array, slice, map, function as variable, function closures
  * Date (struct) - shows how to make a `struct`
  * Pointer - how to derefence and inderect in Go
  * StructAndPointer - how to use `pointer` and `struct` together
  * Array - how to make a simple `array`
  * Slice - how to make a `slice` from an `array`
  * SliceAsRef - how `slice` is a reference to an `array`
  * SliceLiteral - how to make a `slice literal`
  * SliceMatrix - another/simplified example of `slice literal`
  * WordCount - example of a `map`
  * FuncAsVal - how `func` can be a value
  * Subtractr - example of function closure
  * Fibonacci - Fibonacci nr solution with function closure

## Further reading

The Go official site (https://golang.org) is a great source.
For example 

* getting started: https://golang.org/doc/tutorial/getting-started
* the Go tour: https://tour.golang.org
* effective go: https://golang.org/doc/effective_go

Part of learning Go (for me) was also writing a summary of the materials I went through.
This summary can be found at https://github.com/moledoc/moledoc/tree/main/GoLang.
It gives overview of the language basics with examples in a single document.

## Author

Written by
Meelis Utt