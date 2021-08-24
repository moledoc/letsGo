package moretypes

import (
	"fmt"
	"strings"
)

// this is a simple struct
type Date struct {
	Year  int
	Month int
	Day   int
}

// can define multiple program level variables
var (
	testDate1 = Date{}
	testDate2 = Date{1900, 1, 1}
	testDate3 = &Date{1800, 1, 1}
)

// dereferencing and inderecting
func Pointer() int {
	i, j := 42, 2701
	ptr1 := &i // point to i
	*ptr1 = 21 // set i through the pointer

	ptr2 := &j         // point to j
	*ptr2 = *ptr2 / 37 // divide j through the pointer

	return *ptr1 + *ptr2
}

// struct pointers
func StructAndPointer() string {
	date := Date{1900, 1, 1}
	// struct field
	day := date.Day
	// pointers to structs
	datePtr := &date
	(*datePtr).Month = 2 // change var date.Month to be 2 by dereferencing a pointer
	datePtr.Year = 1901  // change var date.Year to be 1901 by dereferencing a pointer (syntactic sugar)

	month := datePtr.Month
	year := datePtr.Year

	return fmt.Sprintf("%v-%v-%v", year, month, day)
}

// array sample
func Arrays() [3]Date {
	var dates [3]Date
	dates[0] = testDate1
	dates[1] = testDate2
	dates[2] = *testDate3
	return dates
}

// slice made from an array
func Slice() []int {
	var numbers = [6]int{0, 1, 2, 3, 4, 5}
	return numbers[3:5]
}

// slices are like references to arrays
func SliceAsRef() ([4]string, []string, []string) {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	a := names[0:2]
	b := names[1:3]

	b[0] = "XXX"
	return names, a, b
}

// how to define slice literals
func SliceLiteral() ([]int, []bool, []struct {
	i int
	b bool
}) {
	var nrs = []int{2, 3, 5, 7, 11, 13}

	bools := []bool{true, false, true, true, false, true}

	structSlice := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	return nrs, bools, structSlice
}

func SliceMatrix() [][]string {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	return board
}

// exampe of a map
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCount := make(map[string]int)

	for _, key := range words {
		_, ok := wordCount[key]
		if !ok {
			wordCount[key] = 0
		}
		wordCount[key] += 1
	}
	return wordCount
}

// example of a function as a value (arg and return)
func FuncAsVal(fn func(float32, float32) float32) float32 {
	compute := func(fn func(float32, float32) float32) float32 {
		return fn(7, 3)
	}
	return compute(fn)
}

// function closure
func Subtractr() func(int) int {
	result := 0
	return func(x int) int {
		result -= x
		return result
	}
}

// fibonacci with closure
func Fibonacci() func() int {
	fibCur := 1
	fibPrev := -1 // so we can start the seq with 0, 1, 1, ...
	return func() int {
		fibNext := fibCur + fibPrev
		fibPrev = fibCur
		fibCur = fibNext
		return fibNext
	}
}
