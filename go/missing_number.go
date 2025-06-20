//lint:file-ignore U1000 Ignore all unused code
package main

func missingNumber(nums []int) int {
	expectedTotal := (1 + len(nums)) * len(nums) / 2
	var actualTotal int
	for _, n := range nums {
		actualTotal += n
	}
	return expectedTotal - actualTotal
}
