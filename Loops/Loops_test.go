package loops_test

import (
	"testing"

	loops "github.com/moledoc/letsgo/Loops"
)

func TestForLoop(t *testing.T) {
	const control = 45
	var cycle10 = loops.ForLoop(10)
	if cycle10 != control {
		t.Fatalf("ForLoop returns incorrect result: expected %v, got %v", control, cycle10)
	}
}

func TestForLoopOnlyCond(t *testing.T) {
	var cycle10 = loops.BasicallyWhileLoop(10)
	if cycle10 != 10 {
		t.Fatalf("ForLoopOnlyCond returns incorrect result: expected %v, got %v", 10, cycle10)
	}
}
