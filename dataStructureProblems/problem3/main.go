package main
import (
	"fmt"
)
type node struct {
	Val int
	Next *node
}
type ListNode struct {
	Head *node
	Tail *node
	length int
}
func (l *ListNode) append (node *node) {
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = l.Tail.Next
	}
	l.length++
}
func (l *ListNode) display () {
	list := l.Head
	for list != nil {
		fmt.Println(list)
		list = list.Next
	}
}
func (l *ListNode) deleteDuplicates () {
    if l.Head == nil {
        return 
    }        
    second := l.Head
	fmt.Println(second)
    for second != nil && second.Next != nil {
        if second.Val == second.Next.Val {
            second.Next = second.Next.Next
        } else {
            second = second.Next
        }
    }
	fmt.Println()
}
func (l *ListNode) deleteAllDuplicates () {
	if l.Head == nil {
		return 
	}
	// l.Head := l.Head
	if l.Head.Next == nil {
		return
	} else {
		if l.Head.Val == l.Head.Next.Val {
			if l.Head.Next.Next == nil {
				l.Head = nil
				// l.Head.Next = nil
			}
			l.Head = l.Head.Next.Next
			fmt.Println(l.Head)
		}
	}
	second := l.Head
	third := second.Next
	for third != nil && third.Next != nil {
	
		if third.Val == third.Next.Val {
			if third.Next.Next == nil {
				second = nil
			}
			second = third.Next.Next
		}
		second = second.Next
	}
	fmt.Println()
	return
}
// func deleteAllDuplicates2 (Head *ListNode) *ListNode {  
// 	// ----------------- this is not working --------------
// 	var newList *ListNode
// 	cur := Head
// 	if cur == nil || cur.Next == nil {
// 		return newList
// 	}
// 	for cur != nil && cur.Next != nil {
// 		if cur.Val != cur.Next.Val {
// 			newList = cur

// 		}
// 	}

// }
func middleElement (l *ListNode) *node {
	length := 0
	temp := l.Head
	for temp != nil {
		length++
		temp = temp.Next
	}
	temp = l.Head
	var result *node
	for i := 0 ; i <= length/2 ; i++ {
		result = temp
		temp = temp.Next
	}
	return result
}
func main() {
	list := ListNode{}
	node1 := node {
		Val: 10,
	}
	node2 := node {
		Val: 10,
	}
	node3 := node {
		Val: 20,
	}
	node4 := node {
		Val: 50,
	}
	node5 := node {
		Val: 50,
	}
	node6 := node {
		Val: 60,
	}
	list.append(&node1)
	list.append(&node2)
	list.append(&node3)
	list.append(&node4)
	list.append(&node5)
	list.append(&node6)
	list.display()
	fmt.Println()
	list.deleteDuplicates()
	// list.deleteAllDuplicates()
	list.display()
	fmt.Println()
	fmt.Println(middleElement(&list))
}