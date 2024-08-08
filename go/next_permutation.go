//lint:file-ignore U1000 Ignore all unused code
package main

import "sort"

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			minN := nums[i]
			idx := i
			for j := i; j < len(nums); j++ {
				if nums[j] < minN && nums[i-1] < nums[j] {
					idx = j
					minN = nums[j]
				}
			}
			nums[i-1], nums[idx] = nums[idx], nums[i-1]
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
