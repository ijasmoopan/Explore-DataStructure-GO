package main

import "fmt"

// Node
type Node struct {
	Key int
	Right *Node
	Left *Node
}

// Insert
func (node *Node) Insert (k int) {
	if node == nil {
		return 
	}
	if node.Key == k {
		return 
	}
	if node.Key != k {
		if node.Left == nil {
			
		}
	}
	if node.Left == nil {
		node.Left = &Node{Key: k}
	} else if node.Right == nil {
		node.Right = &Node{Key: k}
	} else {
		node.Left.Insert(k)
	}
}

// Search
// func (node *Node) Search (k int) bool {
// 	if node == nil {
// 		return false
// 	}
// 	if node.Key == k {
// 		return true
// 	} else if node.Left != nil {
// 		if node.Left.Key == k {
// 			return true
// 		} else {
// 			return false
// 		}
// 	} else if node.Right != nil {
// 		if node.Right.Key == k {
// 			return true
// 		} else {
// 			return false
// 		}
// 	} else {
// 		return node.Left.Search(k)
// 	}	
// 	return node.Right.Search(k)
// }
func (node *Node) Search (k int) bool {
	if node == nil {
		return false
	}
	if node.Key == k {
		return true
	}
	node.Left.Search(k)
	node.Right.Search(k)
	return true
}

// InOrderTraverse
func (node *Node) InOrderTraverse () {
	if node == nil {
		return 
	}
	node.Left.InOrderTraverse()
	fmt.Println(node.Key)
	node.Right.InOrderTraverse()
}

// PreOrderTraverse
func (node *Node) PreOrderTraverse () {
	if node == nil {
		return 
	}
	fmt.Println(node.Key)
	node.Left.PreOrderTraverse()
	node.Right.PreOrderTraverse()
}

// PostOrderTraverse
func (node *Node) PostOrderTraverse () {
	if node == nil {
		return 
	}
	node.Left.PostOrderTraverse()
	node.Right.PostOrderTraverse()
	fmt.Println(node.Key)
}

func main() {
	tree := &Node{Key: 10}
	tree.Insert(15)
	tree.Insert(1)
	tree.Insert(17)
	tree.Insert(8)
	// fmt.Println(tree)
	// fmt.Println("Inorder:")
	// tree.InOrderTraverse()
	// fmt.Println("Preorder:")
	// tree.PreOrderTraverse()
	// fmt.Println("Postorder:")
	// tree.PostOrderTraverse()
	fmt.Println(tree.Search(20))

}