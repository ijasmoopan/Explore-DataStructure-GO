package main

import "fmt"

const AlphabetSize = 26

// Node represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie will create new Trie
func InitTrie() *Trie {
	return &Trie{
		root: &Node{},
	}
}

// Insert will take a word and add it into the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		// fmt.Println(charIndex)
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search will take in a word and RETURN true if the word included in that trie
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

// suffixTree will make a group of words from one word in suffix method
func (t *Trie) suffixTrie(str string) {
	for i := 0; i < len(str); i++ {
		t.Insert(str[i:])
		// fmt.Println(str[i:])
	}
}

func (t *Trie) SearchWord (w string) string {
	wordLength := len(w)
	currentNode := t.root
	result := ""
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return "Not in this Trie"
		}
		result = result + string(charIndex + 'a')
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return result
	}
	return "Not in this Trie"
}

func main() {
	myTrie := InitTrie()
	toAdd := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}
	for _, v := range toAdd{
		myTrie.Insert(v)
	}
	// str := "ijas"
	// myTrie.suffixTrie(str)
	fmt.Println(myTrie.Search("oreo"))
	fmt.Println(myTrie.SearchWord("argo"))
}
