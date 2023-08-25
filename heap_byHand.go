package main

type MaxHeap struct {
	arr []int
}

// NewMaxHeap 返回一个新的最大堆
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// BuildHeap 使用给定的数组构建一个新的最大堆
func (h *MaxHeap) BuildHeap(nums []int) {
	h.arr = nums
	for i := len(nums)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

// Insert 向堆中插入一个新的元素
func (h *MaxHeap) Insert(key int) {
	h.arr = append(h.arr, key)
	h.heapifyUp(len(h.arr) - 1)
}

// ExtractMax 移除并返回堆中的最大元素
func (h *MaxHeap) ExtractMax() int {
	if len(h.arr) == 0 {
		return 0
	}

	extract := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]

	h.heapifyDown(0)

	return extract
}

// heapifyUp 从下到上调整堆
func (h *MaxHeap) heapifyUp(i int) {
	for h.arr[parent(i)] < h.arr[i] {
		h.swap(parent(i), i)
		i = parent(i)
	}
}

// heapifyDown 从上到下调整堆
func (h *MaxHeap) heapifyDown(i int) {
	largest := i
	left := leftChild(i)
	right := rightChild(i)

	if left < len(h.arr) && h.arr[left] > h.arr[largest] {
		largest = left
	}
	if right < len(h.arr) && h.arr[right] > h.arr[largest] {
		largest = right
	}
	if largest != i {
		h.swap(i, largest)
		h.heapifyDown(largest)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(parent int) int {
	return 2*parent + 1
}

func rightChild(parent int) int {
	return 2*parent + 2
}

func (h *MaxHeap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}
func findKthLargest(nums []int, k int) int {
	maxHeap := NewMaxHeap()
	maxHeap.BuildHeap(nums)
	for i := 0; i < k-1; i++ {
		maxHeap.ExtractMax()
	}
	return maxHeap.ExtractMax()
}
