package main

import "fmt"

type minHeap struct {
	array []int
}

func (h *minHeap) insert(key int) {
	h.array = append(h.array, key)
	h.minHeapifyUp(len(h.array) - 1)
}
func (h *minHeap) minHeapifyUp(index int) {
	for h.array[index] < h.array[parent(index)] {
		h.swap(index, parent(index))
		index = parent(index)
	}
}
func (h *minHeap) deleteRoot() int {
	toDelete := h.array[0]
	l := len(h.array) - 1
	h.array[0] = h.array[l]
	if len(h.array) == 0 {
		fmt.Println("Extract is not possible because length of the array is 0")
		return -1
	}
	h.array = h.array[:l]
	h.minHeapifyDown(0)
	return toDelete
}
func (h *minHeap) minHeapifyDown (index int) {
	lastIndex := len(h.array)-1
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] < h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.array[index] > h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return 
		}
	}
}
func parent(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return 2*i + 1
}
func right(i int) int {
	return 2*i + 2
}
func (h *minHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
func main() {
	m := &minHeap{}
	buildHeap := []int{49, 24, 52, 51, 5, 6, 32, 54, 25, 2, 9, 10}
	for _, elem := range buildHeap {
		m.insert(elem)
	}
	fmt.Println(m)
	m.deleteRoot()
	fmt.Println(m)
}