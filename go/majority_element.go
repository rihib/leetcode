//lint:file-ignore U1000 Ignore all unused code
package main

func majorityElement(nums []int) int {
	frequencies := make(map[int]int, len(nums))
	for _, n := range nums {
		frequencies[n]++
		if frequencies[n] > len(nums)/2 {
			return n
		}
	}
	return -1
}

func majorityElementBoyerMooreMajorityVote(nums []int) int {
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
