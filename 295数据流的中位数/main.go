package main

import (
	"container/heap"
	"sort"
)

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.

type MinHeap struct{ sort.IntSlice }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := h.IntSlice
	x := old[len(old)-1]
	h.IntSlice = old[0 : len(old)-1]
	return x
}

// //
type MedianFinder struct {
	minHeap *MinHeap
	maxHeap *MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: &MinHeap{},
		maxHeap: &MinHeap{},
	}
}

// 小堆放大数,大堆放小数,
// 小堆的最小大于等于大堆的最大
func (ct *MedianFinder) AddNum(num int) {
	minQ, maxQ := ct.minHeap, ct.maxHeap
	if minQ.Len() == 0 || num <= -minQ.IntSlice[0] {
		heap.Push(minQ, -num)
		if minQ.Len() > maxQ.Len()+1 {
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	} else {
		heap.Push(maxQ, num)
		if minQ.Len() < maxQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}

func (ct *MedianFinder) FindMedian() float64 {
	if ct.minHeap.Len() == ct.maxHeap.Len() {
		return float64(ct.maxHeap.IntSlice[0]-ct.minHeap.IntSlice[0]) / 2
	} else {
		return float64(-ct.minHeap.IntSlice[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
