//lint:file-ignore U1000 Ignore all unused code
package main

func topKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}
	countToNum := make([][]int, len(nums)+1)
	for num, count := range frequency {
		countToNum[count] = append(countToNum[count], num)
	}
	var topK []int
	for i := len(countToNum) - 1; i >= 0; i-- {
		topK = append(topK, countToNum[i]...)
		if len(topK) >= k {
			break
		}
	}
	return topK
}
