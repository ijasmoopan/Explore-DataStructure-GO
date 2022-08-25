package main

import "fmt"

type maxHeap struct {
	array []int
}

func (h *maxHeap) Insert(v int) {
	h.array = append(h.array, v)
	h.HeapifyUp(len(h.array) - 1)
}
func (h *maxHeap) HeapifyUp(index int) {
	for h.array[index] > h.array[Parent(index)] {
		h.Swap(index, Parent(index))
		index = Parent(index)
	}
}
func (h *maxHeap) HeapSort() []int {
	p := len(h.array) - 1
	// sortedArray := []int{}
	sortedArray := make([]int, len(h.array))
	for p >= 0 {
		sortedArray[p] = h.DeleteRoot()
		// fmt.Println("sortedArray[p] = ",sortedArray[p])
		p--
	}
	return sortedArray
}

func (h *maxHeap) DeleteRoot() int {
	if len(h.array) == 0 {
		fmt.Println("Array is empty")
		return -1
	}
	toDelete := h.array[0]
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	h.HeapifyDown(0)
	// fmt.Println("toDelete = ", toDelete)
	return toDelete
}
func (h *maxHeap) HeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := Left(index), Right(index)
	// fmt.Println("Left : ", l, "Right: ", r)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.array[index] < h.array[childToCompare] {
			h.Swap(index, childToCompare)
			index = childToCompare
			l, r = Left(index), Right(index)
		} else {
			return
		}
	}
}

func Parent(i int) int {
	return (i - 1) / 2
}
func Left(i int) int {
	return 2*i + 1
}
func Right(i int) int {
	return 2*i + 2
}
func (h *maxHeap) Swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	mh := &maxHeap{}
	array := []int{50, 40, 30, 20, 10}
	for _, v := range array {
		mh.Insert(v)
	}
	fmt.Println("Unsorted Array : ", mh.array)
	// mh.DeleteRoot()
	// fmt.Println("After deleting root : ",mh.array)

	sortedArray := mh.HeapSort()
	fmt.Println("Sorted Array : ",sortedArray)
	str := "ijas"
	fmt.Println(str[1:])
}
