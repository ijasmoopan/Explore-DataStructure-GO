package main

import "fmt"

func reverseString(str string) string {
	if len(str) == 0 {
		return ""
	}
	fmt.Println(str)
	return reverseString(str[1:]) + string(str[0])
}
func isPalindrome (str string) bool {
	if len(str)==0 || len(str)==1 {
		return true
	}
	fmt.Println(str)
	if str[0] == str[len(str)-1] {
		return isPalindrome(str[1:len(str)-1])
	}
	return false
}
func decimalToBinary (dec int, result string) string {
	if dec == 0 {
		return result
	}
	result = fmt.Sprint(dec % 2 , result)
	// fmt.Println(result)
	// result = dec % 2 + result
	return decimalToBinary(dec / 2, result) 
}
func sumOfNaturalNumbers (num int, sum int) int {
	if num == 0 {
		return sum
	}
	sum = sum + num
	return sumOfNaturalNumbers(num-1, sum)
}
func sumOfNaturalNumbersCorrect (num int) int {
	if num <= 1 {
		return num
	}
	return num + sumOfNaturalNumbersCorrect(num-1)
}
func binarySearch (array []int, left, right, target int) bool {
	if left > right {
		return false
	}
	mid := left + (right - left) / 2
	if array[mid] == target {
		return true
	}
	if array[mid] > target {
		return binarySearch(array, left, mid-1, target)
	}
	return binarySearch(array, mid+1, right, target)
}
func fibonacciSeries (n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return fibonacciSeries(n-1) + fibonacciSeries(n-2)
	}
}
func main() {
	// str := "hello"
	// fmt.Println(str[:1])
	// fmt.Println(string(str[0]))
	// fmt.Println(reverseString(str))
	// fmt.Println(isPalindrome("kayayak"))
	// fmt.Println(decimalToBinary(10,""))
	// fmt.Println(sumOfNaturalNumbers(15,0))
	// fmt.Println(sumOfNaturalNumbersCorrect(22))
	// array := []int{1,2,3,4,5,6}
	// fmt.Println(binarySearch([]int{-2,4,5,6},0,5,-2))
	fmt.Println(fibonacciSeries(7))
}