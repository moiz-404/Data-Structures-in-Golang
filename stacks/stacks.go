package main

import "fmt"

//stack represents stack that hold a slice
type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

//Pop will removea value at th end
//and RETURNs th remove vlue
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {
	myStack := Stack{}
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	myStack.Push(400)
	myStack.Push(500)
	myStack.Push(600)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
