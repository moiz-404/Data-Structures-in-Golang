package main

import "fmt"

//Queues represents stack that hold a slice
type Queues struct {
	items []int
}

//Enqueue
func (q *Queues) Enqueue(i int) {
	q.items = append(q.items, i)
}

//Dequeue will remove a value at th front
//and RETURNs th remove vlue
func (q *Queues) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queues{}
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	myQueue.Enqueue(400)
	myQueue.Enqueue(500)
	myQueue.Enqueue(600)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
