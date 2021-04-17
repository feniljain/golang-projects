package main

import "fmt"

//Node represens a node in the linked list
type Node struct {
	data int
	Next *Node
}

func addNode(t *Node, v int) int {
	if root == nil {
		t = &Node{
			data: v,
			Next: nil,
		}
		root = t
		fmt.Println("Creating root node with data:", v)
		return 0
	}

	if t.data == v {
		fmt.Println("Node already exists:", v)
		return -1
	}

	if t.Next == nil {
		t.Next = &Node{
			data: v,
			Next: nil,
		}
		fmt.Println("Creating new node with data:", v)
		return -2
	}

	return addNode(t.Next, v)
}

func transverse(t *Node) {
	if t == nil {
		fmt.Println("=> Empty List")
		return
	}

	for t != nil {
		fmt.Printf("%d => ", t.data)
		t = t.Next
	}

	fmt.Println()
}

var root = new(Node)

func main() {
	fmt.Println(root)
	root = nil
	fmt.Println(root)
	transverse(root)
	addNode(root, 1)
	addNode(root, 1)
	transverse(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 0)
	addNode(root, 0)
	transverse(root)
	addNode(root, 100)
	transverse(root)
}
