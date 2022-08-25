package main

import "fmt"

func changingArrayValues(array []int){
	array[0] = 5
	array[4] = 0
}
func main() {
	array := []int{1, 2, 3, 4, 5}
	fmt.Println(array)
	changingArrayValues(array)
	fmt.Println(array)
}