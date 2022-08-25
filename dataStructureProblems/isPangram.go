package main
import (
	"fmt"
	"strings"
)
func isPangram (s string) bool {
	m := make(map[string]int) 
	s = strings.ToLower(s)

	for idx, char := range s {
		if char >= 'a' && char <= 'z' {
			if _, found := m[string(char)] ; !found {
				m[string(char)] = idx
			}
		}
	}
	return len(m) == 26
}
func main() {
	str := "The quick brown fox jumps over the lazy dog"
	fmt.Println(isPangram(str))
}

