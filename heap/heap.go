package main

import (
	"fmt"
)

type Maxheap struct {
	items []int
}

//Insert adds an element to the heap
func (h *Maxheap) Insert(key int) {
	h.items = append(h.items, key)
	h.MaxHeapifyUP(len(h.items) - 1)
}

//Extract returns the largest, removes it from the heap
func (h *Maxheap) Extract() int {
	Extracted := h.items[0]
	l := len(h.items) - 1
	//when the array is empty
	if len(h.items) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}
	//take out the last index and put it in the root
	h.items[0] = h.items[l]
	h.items = h.items[:l]
	h.MaxHeapifydown(0)

	return Extracted
}

//MaxHeapifyUP will heaoify from bottom top
func (h *Maxheap) MaxHeapifyUP(index int) {
	for h.items[parent(index)] < h.items[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

//MaxHeapifydown will heaoify from bottom top
func (h *Maxheap) MaxHeapifydown(index int) {
	lastIndex := len(h.items) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	//loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { //when left child is the only child
			childToCompare = l
		} else if h.items[l] < h.items[r] { //when left child is larger
			childToCompare = l
		} else { //when right child is larger
			childToCompare = r
		}
		//compare array value of current index to larger child and swap if small
		if h.items[index] < h.items[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

//get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

//get the left index
func left(i int) int {
	return 2*i + 1
}

//get the right index
func right(i int) int {
	return 2*i + 2
}

//swap keys
func (h *Maxheap) swap(i1, i2 int) {
	h.items[i1], h.items[i2] = h.items[i2], h.items[i1]
}

func main() {
	m := &Maxheap{}
	fmt.Println(m)
	buildHeap := []int{10, 40, 30, 50, 14, 21, 45, 20, 58, 42, 15, 43, 31, 26, 27, 28, 8, 2, 8, 61, 68}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}

}
