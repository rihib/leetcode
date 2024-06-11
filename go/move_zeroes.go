//lint:file-ignore U1000 Ignore all unused code
package main

func moveZeroes(nums []int) {
	for i, n := range nums {
		if n == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					tmp := nums[i]
					nums[i] = nums[j]
					nums[j] = tmp
					break
				}
			}
		}
	}
}
