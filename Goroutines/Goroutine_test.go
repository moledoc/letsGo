package goroutine_test

import (
	"testing"

	gr "github.com/moledoc/letsGo/Goroutine"
	"golang.org/x/tour/tree"
)

func TestSame(t *testing.T) {
	tree1 := tree.New(1)
	tree2 := tree.New(2)
	if !gr.Same(tree1, tree1) {
		t.Fatalf("Trees %+v and %+v are different", tree1, tree1)
	}
	if gr.Same(tree1, tree2) {
		t.Fatalf("Trees %+v and %+v are different", tree1, tree2)
	}
}
