package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Tree represents a tree node
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func tranverse(t *Tree) {
	if t == nil {
		return
	}
	tranverse(t.Left)
	fmt.Print(t.Value, " ")
	tranverse(t.Right)
}

func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n)
		t = insert(t, temp)
	}
	return t
}

func insert(t *Tree, v int) *Tree {

	if t == nil {
		return &Tree{
			Left:  nil,
			Value: v,
			Right: nil,
		}
	}

	if v == t.Value {
		return t
	}

	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}

	t.Right = insert(t.Right, v)
	return t
}

func main() {
	tree := create(30)
	tranverse(tree)
	fmt.Println()
	fmt.Println("The value of the root of the tree is", tree.Value)
}
