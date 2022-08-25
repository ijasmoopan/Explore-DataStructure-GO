package main
import (
	"fmt"
)

// ----------Definition for singly Linked List ------------
type Node struct {
	Value int
	Next *Node
}
func newNode (data int, next *Node) *Node {
	var n Node
	n.Value = data
	n.Next = next
	return &n
}
func addNodeAtEnd (head *Node, data int) *Node {
	temp := head
	if temp == nil {
		temp = newNode(data, nil)
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = newNode(data, nil)
	}
	return head
}
// func (head *Node) display () {
// 	temp := head
// 	for temp != nil {
// 		fmt.Println("method:", temp.Value)
// 		temp = temp.Next
// 	}
// }
func traversalLL (head *Node) {
	temp := head
	for temp != nil {
		fmt.Println(temp)
		temp = temp.Next
	}
}
func addNodeAtFront (head *Node, data int) *Node {
	if head == nil {
		head = newNode(data, nil)
	} else {
		temp := head
		head = newNode(data, temp)
	}
	return head
}
func deleteANode (head *Node, data int) *Node {
	if head == nil {
		return head
	}
	if head.Value == data {
		head = head.Next
		return head
	}
	temp := head
	for temp.Next.Value != data {
		if temp.Next.Next == nil {
			return head
		}
		temp = temp.Next
	}
	temp.Next = temp.Next.Next
	return head
}
func reverseLL (head *Node) {
	if head == nil {
		return 
	}
	reverseLL(head.Next)
	fmt.Println(head)
}
func removeDuplicates (head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	temp := head
	for temp.Next != nil {
		if temp.Value == temp.Next.Value {
			temp.Next = temp.Next.Next
		}
		temp = temp.Next
	}
	return head
}
func middle (head *Node) int {
	mid := head
	end := head
	for end != nil && end.Next != nil {
		mid = mid.Next
		end = end.Next.Next
	}
	return mid.Value
}
func pairSwapping(head *Node) *Node{
	temp := head
	for temp != nil && temp.Next != nil {
		tmp := temp.Value
		temp.Value = temp.Next.Value
		temp.Next.Value = tmp
		if temp.Next.Next == nil {
			return head
		}
		temp = temp.Next.Next
	}
	return head
}
func traverseValues (head *Node) {
	temp := head
	for temp != nil {
		fmt.Print(temp.Value)
		temp = temp.Next
	}
}

func main () {
	head := newNode(10, newNode(10, newNode(20, newNode(30, newNode(30, nil)))))
	
	head = addNodeAtEnd(head, 40)
	head = addNodeAtEnd(head, 50)
	// traversalLL(head)
	// head = addNodeAtFront(head, 100)
	// list := Node{}
	// list.display()
	traversalLL(head)
	fmt.Println()
	// head = deleteANode(head, 30)
	// traversalLL(head)
	// head = removeDuplicates(head)
	// traversalLL(head)
	// reverseLL(head)
	// fmt.Println(middle(head))
	// fmt.Println(pairSwapping(head))
	head = pairSwapping(head)
	traversalLL(head)
	traverseValues(head)
}