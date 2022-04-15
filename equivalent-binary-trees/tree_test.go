package main

import (
	"golang.org/x/tour/tree"
	"testing"
)

func TestSame(t *testing.T) {
	t.Run("Tree(1)", func(t *testing.T) {
		if !Same(tree.New(1), tree.New(1)) {
			t.Error("Tree(1) failed")
		}
	})
	t.Run("Tree(2)", func(t *testing.T) {
		if !Same(tree.New(2), tree.New(2)) {
			t.Error("Tree(2) failed")
		}
	})
	t.Run("Different", func(t *testing.T) {
		if Same(tree.New(1), tree.New(2)) {
			t.Error("Tree(1) Tree(2) failed")
		}
	})
}
