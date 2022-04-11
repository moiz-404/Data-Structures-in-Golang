package main

import "fmt"

var count int

//node represent the components of a binary search
type Node struct {
	key   int
	left  *Node
	right *Node
}

//Insert will add a node to the tree
//the key to add should not be already int tree
func (n *Node) Insert(k int) {
	if n.key > k {
		if n.left == nil {
			n.left = &Node{key: k}
		} else {
			n.left.Insert(k)
		}
	} else if n.key < k {
		if n.right == nil {
			n.right = &Node{key: k}
		} else {
			n.right.Insert(k)
		}
	}
}

//Search will take in a key value
//and RETURN true if is a node with that value
func (n *Node) Search(k int) bool {
	count++

	if n == nil {
		return false
	}

	if n.key < k {
		return n.right.Search(k)
	} else if n.key > k {
		return n.left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{key: 100}
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(150)
	tree.Insert(250)
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(5)
	tree.Insert(80)
	tree.Insert(60)
	tree.Insert(400)
	tree.Insert(350)
	fmt.Println(tree)
	x := tree.Search(400)
	fmt.Println(x)
	fmt.Println(count)
}
