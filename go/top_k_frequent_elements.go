//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"container/heap"
	"math/rand/v2"
	"sort"
)

func topKFrequentBucketSort(nums []int, k int) []int {
	frequencies := make(map[int]int, len(nums))
	for _, n := range nums {
		frequencies[n]++
	}
	countToNum := make([][]int, len(nums)+1)
	for n, count := range frequencies {
		countToNum[count] = append(countToNum[count], n)
	}
	topK := make([]int, 0, k)
	for i := len(countToNum) - 1; i >= 0 && len(topK) < k; i-- {
		topK = append(topK, countToNum[i]...)
	}
	return topK
}

type element struct {
	num   int
	count int
}

func topKFrequentQuickselect(nums []int, k int) []int {
	frequencies := make(map[int]int, len(nums))
	for _, n := range nums {
		frequencies[n]++
	}
	elements := make([]element, 0, len(frequencies))
	for n, count := range frequencies {
		elements = append(elements, element{n, count})
	}
	quickselect(elements, 0, len(elements)-1, len(elements)-k)
	topK := make([]int, 0, k)
	for i := 0; i < k; i++ {
		topK = append(topK, elements[len(elements)-1-i].num)
	}
	return topK
}

func quickselect(elements []element, left, right, k int) {
	for left < right {
		pivotIndex := partition(elements, left, right)
		if k == pivotIndex {
			return
		}
		if k < pivotIndex {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}
	}
}

func partition(elements []element, left, right int) int {
	pivotIndex := left + rand.IntN(right-left+1)
	pivotValue := elements[pivotIndex].count
	elements[pivotIndex], elements[right] = elements[right], elements[pivotIndex]
	storeIndex := left
	for i := left; i < right; i++ {
		if elements[i].count < pivotValue {
			elements[storeIndex], elements[i] = elements[i], elements[storeIndex]
			storeIndex++
		}
	}
	elements[storeIndex], elements[right] = elements[right], elements[storeIndex]
	return storeIndex
}

func topKFrequentPDQSort(nums []int, k int) []int {
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}
	numsSet := make([]int, 0, len(frequency))
	for num := range frequency {
		numsSet = append(numsSet, num)
	}
	sort.Slice(numsSet, func(i, j int) bool {
		return frequency[numsSet[i]] > frequency[numsSet[j]]
	})
	return numsSet[:k]
}

func topKFrequentMinHeap(nums []int, k int) []int {
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}
	h := &MinHeap{}
	heap.Init(h)
	for num, count := range frequency {
		heap.Push(h, element{num: num, count: count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	topK := make([]int, 0, k)
	for h.Len() > 0 {
		topK = append(topK, heap.Pop(h).(element).num)
	}
	return topK
}

type MinHeap []element

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(element))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func topKFrequentMaxHeap(nums []int, k int) []int {
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}
	h := &MaxHeap{}
	heap.Init(h)
	for num, count := range frequency {
		heap.Push(h, element{num: num, count: count})
	}
	topK := make([]int, 0, k)
	for i := 0; i < k; i++ {
		topK = append(topK, heap.Pop(h).(element).num)
	}
	return topK
}

type MaxHeap []element

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(element))
}

func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
