//lint:file-ignore U1000 Ignore all unused code
package main

func majorityElement(nums []int) int {
	candidate, count := nums[0], 0
	for _, n := range nums {
		if n == candidate {
			count++
		} else {
			count--
		}
		if count < 0 {
			candidate = n
			count = 0
		}
	}
	return candidate
}
