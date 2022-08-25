package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// ----------INSERT--------------
// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// maxHeapifyUp will heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// -----------DELETEROOT-------------
// deleteRoot returns the largest key, and removes it from heap
func (h *MaxHeap) deleteRoot() int {
	toDelete := h.array[0]
	l := len(h.array) - 1
	h.array[0] = l

	if len(h.array) == 0 {
		fmt.Println("Cannot extract because array length is 0.")
		return -1
	}
	h.array = h.array[:l]
	h.maxHeapifyDown(0)
	return toDelete
}

// maxHeapifyDown will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	// loop while index has atleast one child
	for l <= lastIndex {

		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}
		// compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// parent for getting parent index
func parent(i int) int {
	return (i - 1) / 2
}

// left for getting left node index
func left(i int) int {
	return 2*i + 1
}

// right for getting right node index
func right(i int) int {
	return 2*i + 2
}

// swap for swapping 2 indices
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, elem := range buildHeap {
		m.Insert(elem)
		fmt.Println(m)
	}
	for i := 0; i < 5; i++ {
		m.deleteRoot()
		fmt.Println(m)
	}
	// fmt.Println(m)
}
