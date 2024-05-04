//lint:file-ignore U1000 Ignore all unused code
package main

func majorityElement(nums []int) int {
	candidate := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			cnt++
		} else {
			cnt--
			if cnt == 0 {
				candidate = nums[i]
				cnt = 1
			}
		}
	}
	return candidate
}
