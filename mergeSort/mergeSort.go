package main

import "fmt"

func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	left := MergeSort(array[:len(array)/2])
	right := MergeSort(array[len(array)/2:])
	return sorting(left, right)
}
func sorting(left, right []int) []int {
	sortedArray := []int{}
	l, r := 0, 0
	ln := (len(left)+len(right))
	for len(sortedArray) < ln {
		if l == len(left) && r < len(right) {
			sortedArray = append(sortedArray, right[r])
			r++
		} else if r == len(right) && l < len(left) {
			sortedArray = append(sortedArray, left[l])
			l++
		} else if left[l] > right[r] {
			sortedArray = append(sortedArray, right[r])
			r++
		} else {
			sortedArray = append(sortedArray, left[l])
			l++
		}
	}
	return sortedArray
}
func main() {
	array := []int{4,6,72,7,4,85,32,75,2,45,1}
	fmt.Println(array)
	fmt.Println(MergeSort(array))
}