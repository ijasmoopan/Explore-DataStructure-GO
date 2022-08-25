package main
import (
	"fmt"
)
type stack struct {
	items []string
}
func (st *stack) push (data string) {
	st.items = append(st.items, data)
}
func (st *stack) pop () {
	st.items = st.items[:len(st.items)-1]
}
func (st *stack) top () string {
	return st.items[len(st.items)-1]
}
func (st *stack) isEmpty () bool {
	return len(st.items) == 0
}
func (st stack) display() {
	fmt.Println(st.items)
}
func (st *stack) isPalindrome (str string) bool {
	l := len(str)
	count := 0
	// count, p := 0, 0
	// if l % 2 != 0 {
	// 	p = l
	// } 
	for _, char := range str {
		strchar := string(char)
		if count < l/2 {
			st.push(strchar)
		} else if count == l/2 {
			if l % 2 != 0 {
				st.pop()
				count++
				continue
			}
		} else if count >= l/2 {
			if st.top() == strchar {
				st.pop()
			} else {
				return false
			}
		}
		count++
	}
	return true
}
func main() {
	st := stack{}
	str := "malayyalam"
	// for _, char := range str {
	// 	st.push(string(char))
	// }
	fmt.Println(st.isPalindrome(str))
	st.display()
}