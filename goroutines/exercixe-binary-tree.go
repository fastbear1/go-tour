package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t2, ch2)
	go Walk(t1, ch1)
	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	YetTree := tree.New(1)
	fmt.Println(YetTree)
	go Walk(YetTree, ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	r := Same(tree.New(1), tree.New(2))
	z := Same(tree.New(5), tree.New(5))
	fmt.Println(r, z)
}
