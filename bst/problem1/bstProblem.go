package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 
//  func sortedArrayToBST(nums []int) *TreeNode {
//     var Node *TreeNode
//     if nums == nil {
//         return Node
//     }
//     if len(nums) > 0 {
//         Node.Val = nums[0]
//     }
//     for i := 1 ; i < len(nums) ; i++ {
//         Node.insert(nums[i])
//     }
//     return Node
// }
// func (node *TreeNode) insert(value int) *TreeNode {
//     if node == nil {
//         return nil
//     }
//     if node.Val == value {
//         return nil
//     }
//     if node.Val > value {
// 		if node.Left == nil {
// 			node.Left = &TreeNode{Val : value}
// 		}
// 		node.Left.insert(value)
//     }
//     if node.Val < value {
// 		if node.Right == nil {
// 			node.Right = &TreeNode{Val: value}
// 		}
//         node.Right.insert(value)
//     }
// 	return nil
// }
func binarySearch(nums []int, target int) int {
	middle, left, right := 0, 0, len(nums)-1
	for left <= right {
		middle = left + (right - left)/2 
		if nums[middle] == target {
			return middle
		}
		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}

func main() {
	nums := []int{-8,-6,-3,-1,0,3,5,7,9}
	tree := &TreeNode{}
	tree = sortedArrayToBST2(nums)
	// fmt.Println(tree)
	tree.PrintPreOrder()

	// ------Binary Search In An Array------
	fmt.Println(binarySearch(nums, 10))
}

func sortedArrayToBST2(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    return &TreeNode{
        Val: nums[len(nums)/2],
        Left: sortedArrayToBST2(nums[:len(nums)/2]),
        Right: sortedArrayToBST2(nums[len(nums)/2 + 1:]),
    }
}

func (node *TreeNode) PrintPreOrder() {
	if node == nil {
		return 
	}
	fmt.Println(node)
	node.Left.PrintPreOrder()
	node.Right.PrintPreOrder()
}