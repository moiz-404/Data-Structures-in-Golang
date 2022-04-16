package main

import (
	"fmt"
)

//ArraySize is the size of the hash table array
const ArraySize = 7

//HashTable wiil hold an array
type HashTable struct {
	array [ArraySize]*Bucket
}

//Bucket is a linklist in each slot of the hash table array
type Bucket struct {
	head *BucketNode
}

//BucketNode is a linked list node that holds the key
type BucketNode struct {
	key  string
	next *BucketNode
}

// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search will take in a key and return ture if that key is stored in hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete will take in a key and delete it from rhe hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// Bucket Insert will take a key, create a node with the key and
//insert the node in the bucket
func (b *Bucket) insert(k string) {
	if !b.search(k) {
		newNode := &BucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exist")
	}
}

// Bucket Search will take in a key and return true if the bucket has that key
func (b *Bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// Bucket Delete will take in a key and delete the node from the bucket
func (b *Bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			//delete
			previousNode = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

//hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

//Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &Bucket{}
	}
	return result
}

func main() {
	HashTable := Init()
	list := []string{
		"Eric",
		"Kenny",
		"Kyle",
		"Stan",
		"Randy",
		"Butters",
		"Token",
	}
	for _, v := range list {
		HashTable.Insert(v)
	}
	HashTable.Delete("Randy")
	fmt.Println("Randy", HashTable.Search("Randy"))

	// fmt.Println(testHashtable)
	// fmt.Println(hash("Randy"))

	// testBucket := &Bucket{}
	// testBucket.Insert("Randy")
	// testBucket.Delete("Randy")
	// fmt.Println(testBucket.Search("Randy"))
}
