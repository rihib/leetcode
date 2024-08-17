//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"container/heap"
	"math/rand/v2"
	"sort"
)

type Element struct {
	num   int
	count int
}

func topKFrequentBucketSort(nums []int, k int) []int {
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}
	countToNum := make([][]int, len(nums)+1)
	for num, count := range frequency {
		countToNum[count] = append(countToNum[count], num)
	}
	topK := make([]int, 0, k)
	for i := len(countToNum) - 1; i >= 0 && len(topK) < k; i-- {
		topK = append(topK, countToNum[i]...)
	}
	return topK
}

func topKFrequentQuickselect(nums []int, k int) []int {
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}
	elements := make([]Element, 0, len(frequency))
	for num, count := range frequency {
		elements = append(elements, Element{num: num, count: count})
	}
	quickselect(elements, 0, len(elements)-1, len(elements)-k, partitionRandom)
	topK := make([]int, k)
	for i := 0; i < k; i++ {
		topK[i] = elements[len(elements)-1-i].num
	}
	return topK
}

func quickselect(
	elements []Element, left, right, k int,
	partition func([]Element, int, int) int) {
	for left < right {
		pivotIndex := partition(elements, left, right)
		if pivotIndex == k {
			return
		}
		if pivotIndex < k {
			left = pivotIndex + 1
		} else {
			right = pivotIndex - 1
		}
	}
}

func partitionRandom(elements []Element, left, right int) int {
	pivotIndex := left + rand.IntN(right-left+1)
	elements[pivotIndex], elements[right] = elements[right], elements[pivotIndex]
	pivot := elements[right].count
	storeIndex := left
	for i := left; i < right; i++ {
		if elements[i].count < pivot {
			elements[i], elements[storeIndex] = elements[storeIndex], elements[i]
			storeIndex++
		}
	}
	elements[storeIndex], elements[right] = elements[right], elements[storeIndex]
	return storeIndex
}

func partitionMedianOf3(elements []Element, left, right int) int {
	mid := left + (right-left)/2
	if elements[right].count < elements[left].count {
		elements[right], elements[left] = elements[left], elements[right]
	}
	if elements[mid].count < elements[left].count {
		elements[mid], elements[left] = elements[left], elements[mid]
	}
	if elements[right].count < elements[mid].count {
		elements[right], elements[mid] = elements[mid], elements[right]
	}
	pivotIndex := mid
	elements[pivotIndex], elements[right] = elements[right], elements[pivotIndex]
	pivot := elements[right].count
	storeIndex := left
	for i := left; i < right; i++ {
		if elements[i].count < pivot {
			elements[i], elements[storeIndex] = elements[storeIndex], elements[i]
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
		heap.Push(h, Element{num: num, count: count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	topK := make([]int, 0, k)
	for h.Len() > 0 {
		topK = append(topK, heap.Pop(h).(Element).num)
	}
	return topK
}

type MinHeap []Element

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
