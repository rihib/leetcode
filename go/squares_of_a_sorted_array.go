//lint:file-ignore U1000 Ignore all unused code
package main

import "math"

func sortedSquares(nums []int) []int {
	sq := make([]int, len(nums))
	i := len(nums) - 1
	for l, r := 0, len(nums)-1; l <= r; {
		if math.Abs(float64(nums[l])) > math.Abs(float64(nums[r])) {
			sq[i] = nums[l] * nums[l]
			l++
		} else {
			sq[i] = nums[r] * nums[r]
			r--
		}
		i--
	}
	return sq
}
