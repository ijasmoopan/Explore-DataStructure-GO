package main

import "fmt"

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		n--
	}
	return arr
}
func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j >= 1 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}
func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		smallest := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[smallest] {
				smallest = j
			}
		}
		arr[i], arr[smallest] = arr[smallest], arr[i]
	}
	return arr
}

func main() {
	array := []int{10, 5, 8, 4, 7, 2}
	// array = bubbleSort(array)
	array = insertionSort(array)
	// array = selectionSort(array)
	fmt.Println(array)
}
