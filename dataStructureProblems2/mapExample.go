package main
import (
	"fmt"
)
func main () {


	m := make(map[string]int)
	m = map[string]int{
		"a" : 1,
		"b" : 2,
		"c" : 3,
	}
	fmt.Println(m["a"])
	m["a"] = 5
	fmt.Println(m["a"])

	m2 := make(map[string]int)
	str := "ijaas"
	for i, char := range str {
		if _, found := m2[string(char)] ; !found {
				m2[string(char)] = i
		}
	}
	fmt.Println(m2)


nestedMap := make(map[int]map[int]int)
nestedMap = map[int]map[int]int {
	10 : map[int]int {
		20 : 30,
	},
	20 : map[int]int {
		30 : 40,
	},
	30 : map[int]int {
		40 : 50,
	},
}

if v, found := nestedMap[20] ; found {
	fmt.Println(v, found)
}
fmt.Println(nestedMap)

}