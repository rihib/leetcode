//lint:file-ignore U1000 Ignore all unused code
package main

import "sort"

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			candidate, candidateIndex := nums[i], i
			for j := i; j < len(nums); j++ {
				if nums[j] < candidate && nums[i-1] < nums[j] {
					candidate, candidateIndex = nums[j], j
				}
			}
			nums[i-1], nums[candidateIndex] = nums[candidateIndex], nums[i-1]
			sort.Slice(nums[i:], func(a, b int) bool {
				return nums[i+a] < nums[i+b]
			})
			return
		}
	}
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})
}
