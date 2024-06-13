//lint:file-ignore U1000 Ignore all unused code
package main

func majorityElement(nums []int) int {
	candidate, cnt := nums[0], 0
	for _, n := range nums {
		if n == candidate {
			cnt++
		} else {
			cnt--
			if cnt == 0 {
				candidate = n
				cnt++
			}
		}
	}
	return candidate
}
