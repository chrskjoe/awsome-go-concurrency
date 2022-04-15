package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for n := range ch {
		fmt.Println(n)
	}
}
