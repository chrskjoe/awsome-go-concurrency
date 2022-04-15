package main

import (
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkRecursively(t, ch)
	close(ch)
}

func walkRecursively(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkRecursively(t.Left, ch)
	ch <- t.Value
	walkRecursively(t.Right, ch)
}

// Same determines whether the trees
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	select {
	default:
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}
