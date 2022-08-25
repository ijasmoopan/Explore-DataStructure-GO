package main

import (
	"fmt"
	"io"
	"os"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}
type BinarySearchTree struct {
	root *Node
}

// Insert inserts the Item t in the tree
func (bst *BinarySearchTree) Insert(data int) {
	n := &Node{data, nil, nil}
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

// internal function to find the correct place for a node in a tree
func insertNode(node, newNode *Node) {
	if newNode.data < node.data {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

func print(w io.Writer, node *Node, ns int, ch rune) {
	if node == nil {
		return 
	}
	for i := 0 ; i < ns ; i++ {
		fmt.Fprintf(w, " ")
	}
	fmt.Fprintf(w, "%c : %v\n", ch, node)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func main() {
	bst := BinarySearchTree{}
	bst.Insert(50)
	bst.Insert(10)
	bst.Insert(60)
	bst.Insert(70)
	bst.Insert(55)
	bst.Insert(20)
	bst.Insert(8)

	print(os.Stdout, bst.root, 0, 'M')
}
