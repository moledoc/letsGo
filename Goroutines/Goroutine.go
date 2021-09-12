package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.

// Reason why function Walk is not recursive, is because then we can use the `close` function to close the channel
func Walk(t *tree.Tree, ch chan int) {
	walkRecursive(t, ch)
	close(ch)
}

func walkRecursive(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkRecursive((*t).Left, ch)
	ch <- t.Value
	walkRecursive((*t).Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1Chan := make(chan int, 10)
	t2Chan := make(chan int, 10)
	go Walk(t1, t1Chan)
	go Walk(t2, t2Chan)
	for {
		t1Elem, ok1 := <-t1Chan
		t2Elem, ok2 := <-t2Chan
		// if elements differ, trees are not the same
		// if one tree ends before other, trees are not the same
		if t1Elem != t2Elem || ok1 != ok2 {
			return false
		}
		// if channels are closed, then break the loop
		// we can check only one channel, because we checked in prev if they are equal (ok1==ok2).
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int, 10)
	treeWalk := tree.New(1)
	fmt.Println("Test Walk function:")
	fmt.Println(treeWalk)
	go Walk(treeWalk, ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Test if trees are same:")
	fmt.Printf("\tIs the tree1 and tree1 same: %v\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("\tIs the tree1 and tree2 same: %v\n", Same(tree.New(1), tree.New(2)))
}
