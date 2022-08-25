package main

import "fmt"

//Node represents the components of a binary tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//Insert will add a node to a tree
func (node *Node) Insert(k int) {
	if node == nil {
		return 
	}
	if node.Key == k {
		return 
	}
	if node.Key < k {
		if node.Right == nil {
			node.Right = &Node{Key: k}
		} else {
			node.Right.Insert(k)
		}
	} else if node.Key > k {
		if node.Left == nil {
			node.Left = &Node{Key: k}
		} else {
			node.Left.Insert(k)
		}
	}
}

// Search a key in a binary search tree
func (node *Node) Search(k int) bool {
	if node == nil {
		return false
	}
	if node.Key < k {
		return node.Right.Search(k)
	} else if node.Key > k {
		return node.Left.Search(k)
	}
	return true
}

// Traverse in InOrderTraverse
func (node *Node) InOrderTraverse() {
	if node == nil {
		return
	}
	node.Left.InOrderTraverse()
	fmt.Println(node)
	node.Right.InOrderTraverse()
}

// Traverse in Preorder
func (node *Node) PreOrderTraverse() {
	if node == nil {
		return
	}
	fmt.Println(node)
	node.Left.PreOrderTraverse()
	node.Right.PreOrderTraverse()
}

// Traverse in PostOrder
func (node *Node) PostOrderTraverse() {
	if node == nil {
		return
	}
	node.Left.PostOrderTraverse()
	node.Right.PostOrderTraverse()
	fmt.Println(node)
}

// Min
func (node *Node) GetMin() int {
	if node.Left == nil {
		return node.Key
	}
	return node.Left.GetMin()
}

// Max
func (node *Node) GetMax() int {
	if node.Right == nil {
		return node.Key
	}
	return node.Right.GetMax()
}

// Delete a node with given value
func (node *Node) Delete(value int) *Node {
	if node == nil {
		return node
	}
	if node.Key > value {
		node.Left = node.Left.Delete(value)
	}
	if node.Key < value {
		node.Right = node.Right.Delete(value)
	}
	if node.Key == value {
		if node.Left == nil && node.Right == nil {
			node = nil
			return node
		}
		if node.Left != nil && node.Right == nil {
			temp := node.Left
			node = nil
			node = temp
			return node
		}
		if node.Left == nil && node.Right != nil {
			temp := node.Right
			node = nil
			node = temp
			return node
		}
		// Both node.Left & node.Right are not nil
		tempNode := minValued(node.Right)
		node.Key = tempNode.Key
		node.Right = node.Right.Delete(tempNode.Key)
	}
	return node
}
func minValued(node *Node) *Node {
	temp := node
	for temp != nil && temp.Left != nil {
		temp = temp.Left
	}
	return temp
}

func main() {
	tree := &Node{Key: 15}
	tree.Insert(10)
	tree.Insert(25)
	tree.Insert(5)
	tree.Insert(13)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(11)
	tree.Insert(14)
	tree.Insert(18)
	tree.Insert(22)
	tree.Insert(28)
	tree.Insert(35)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(8)
	// fmt.Println(tree)
	// fmt.Println(tree.Search(6))
	// fmt.Println("inorder : ")
	// tree.InOrderTraverse()
	// fmt.Println(tree.GetMin())
	// fmt.Println(tree.GetMax())
	// fmt.Println("preorder : ")
	tree.PreOrderTraverse()
	fmt.Println("postorder : ")
	// tree.PostOrderTraverse()
	// tree.Delete(11)
	// tree.InOrderTraverse()
}
