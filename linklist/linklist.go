package main

import "fmt"

type node struct {
	data int   //data inside the node
	next *node //address of the next node
}

type linkedlist struct {
	head   *node
	length int //how long the linklist is...
}

func (l *linkedlist) prepend(n *node) {
	second := l.head     //temporary place,put current node on second
	l.head = n           //set the new head which is n
	l.head.next = second //let the new head point toward the old head
	l.length++
}

func (l linkedlist) printListData() {
	toprint := l.head
	for l.length != 0 {
		fmt.Println(toprint.data)
		toprint = toprint.next
		l.length--
	}
	fmt.Println("")
}

func (l *linkedlist) deletevalue(val int) {
	if l.length == 0 {
		return
	}
	if l.head.data == val {
		l.head = l.head.next
		l.length--
		return
	}

	previousHeadToDelete := l.head
	for previousHeadToDelete.next.data != val {
		if previousHeadToDelete.next.next == nil {
			return
		}
		previousHeadToDelete = previousHeadToDelete.next
	}
	previousHeadToDelete = previousHeadToDelete.next.next
	l.length--
}

func main() {
	mylist := linkedlist{}
	node1 := &node{data: 48}
	node2 := &node{data: 82}
	node3 := &node{data: 41}
	node4 := &node{data: 44}
	node5 := &node{data: 18}
	node6 := &node{data: 4}
	node7 := &node{data: 8}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.prepend(node7)
	fmt.Println("")
	fmt.Println(mylist)
	fmt.Println("")
	mylist.printListData()
}
