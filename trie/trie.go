package main

import "fmt"

const Alphabetsize = 26

//node represents each node in the trie
type Node struct {
	children [Alphabetsize]*Node
	isEnd    bool
}

//trie represents the trie and pointer to the root node
type Trie struct {
	root *Node
}

//init will create new trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

//insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	wordlength := len(w)
	currentNode := t.root
	for i := 0; i < wordlength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

//search will take in a word and return true is that word is include in the	trie
func (t *Trie) Search(w string) bool {
	wordlength := len(w)
	currentNode := t.root
	for i := 0; i < wordlength; i++ {
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

func main() {
	myTrie := InitTrie()
	//fmt.Println(myTrie.root)
	// myTrie.Insert("aragoon")
	// fmt.Println(myTrie.Search("aragoon"))
	toAdd := []string{
		"aragorn",
		"aragon",
		"argoon",
		"eragon",
		"oragon",
		"orrgon",
		"oregona",
		"oreo",
	}
	for _, v := range toAdd {
		myTrie.Insert(v)
	}
	fmt.Println(myTrie.Search("aragorn"))
}
