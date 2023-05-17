package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	
	for i := 0; i < 10; i++ {
        n1 := <-c1
		n2 := <-c2
		if n1 != n2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	
	for i := 0; i < 10; i++ {
        n := <-ch
		fmt.Printf("%v,", n)
	}
	fmt.Printf("%v,%v",Same(tree.New(1), tree.New(1)),Same(tree.New(1), tree.New(2)))
}
