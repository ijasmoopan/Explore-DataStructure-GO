package main

import (
	"encoding/base64"
	"fmt"
	"sort"
	"strconv"
)

func encodingAString(str string) string {
	var encoded string
	for _, char := range str {
		if char == 'z' {
			char = 'z' - 26
		}
		if char == 'y' {
			char = 'y' - 26
		}
		// fmt.Println(char)
		encoded = encoded + string(char+2)
	}
	return encoded
}

func encodingStr (str string) string {
	var encoded string
	for _, char := range str {
		char = char + 2
		if char > 122 {
			char = 96 + char % 122
		} 
		encoded = encoded + string(char)
	}
	return encoded
}

func countOfRepeated(str string) string {
	result := ""
	count := 0
	for i := 0; i < len(str); i = i + count {
		count = 0
		for _, char2 := range str[i:] {
			if str[i] == byte(char2) {
				count++
			}
		}
		result = result + strconv.Itoa(count) + string(str[i])
	}
	return result
}

// EncodingWithPackage
func EncodingWithPackage (str string) {
	encodedText := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(encodedText)
	decodedText, err := base64.StdEncoding.DecodeString(encodedText)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s",decodedText)
	fmt.Println(string(decodedText))
}

func main() {
	s := "abcdxyz"
	// s = encodingAString(s)
	s = encodingStr(s)
	fmt.Println(s)
	s2 := "aaabbccccddddde"
	// s2 = countOfRepeated(s2)
	s2 = countOfRepeatedCharacters(s2)
	fmt.Println(s2)
	// fmt.Println('A', 'Z', 'a')
	// fmt.Println(string(s2[3]))
	EncodingWithPackage(s2)
	fmt.Println('a','z','{','|')

	nums := []string {"hello","hella","xkh","bcj","abcd"}
	sort.Strings(nums)
	fmt.Println(nums)
}
func countOfRepeatedCharacters(str string) string {
	j, count := 0, 0
	result := ""
	flag := false
	for i := 0; i < len(str); {
		if str[j] == str[i] {
			count++
			i++
			flag= true
		} else {
			result = result + string(str[j]) + strconv.Itoa(count)
			j = j + count
			count = 0
			flag=false
		}
	}
	if flag {
		result = result + string(str[j]) + strconv.Itoa(count)
	}
	return result
}
