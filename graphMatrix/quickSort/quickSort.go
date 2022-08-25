package main

import "fmt"

func main() {
	array := []int{50,40,30,20,10}
	// array := []int{32,41,16,2,21,62,4,25,11}
	result := QuickSort(array)
	fmt.Println(result)
}
func QuickSort (array []int) []int {
	if len(array) < 2 {
		return array
	}
	pivot := 0
	left := 1
	right := len(array)-1
	for left < right {
		if array[pivot] < array[left] && array[pivot] > array[right] {
			array[left] , array[right] = array[right], array[left]
			left ++
			right --
		}
		if array[pivot] > array[left] {
			left ++
		}
		if array[pivot] < array[right] {
			right --
		}
	}
	if array[pivot] > array[right] {
		array[pivot], array[right] = array[right], array[pivot]
	} else {
		array[pivot], array[right-1] = array[right-1], array[pivot]
	}
	// fmt.Println("first: ",array)
	QuickSort(array[:right])
	// fmt.Println("second: ",array)
	QuickSort(array[right:])
	// fmt.Println("third: ",array)
	return array
}

