package sort

import "container/heap"

func Heapsort(nums []int) {
	h := &intHeap{}
	heap.Init(h)

	for _, n := range nums {
		heap.Push(h, n)
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = heap.Pop(h).(int)
	}
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	size := len(*h)
	val := (*h)[size - 1]
	*h = (*h)[0 : size -1]
	return val
}
