//lint:file-ignore U1000 Ignore all unused code
package main

func moveZeroes(nums []int) {
	zeroIdx := 0
	for i, n := range nums {
		if n != 0 {
			if i != zeroIdx {
				nums[i], nums[zeroIdx] = nums[zeroIdx], nums[i]
			}
			zeroIdx++
		}
	}
}
