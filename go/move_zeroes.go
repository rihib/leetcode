//lint:file-ignore U1000 Ignore all unused code
package main

func moveZeroes(nums []int) {
	zeroIndex := 0
	for i, n := range nums {
		if n == 0 {
			continue
		}
		nums[zeroIndex], nums[i] = nums[i], nums[zeroIndex]
		zeroIndex++
	}
}
