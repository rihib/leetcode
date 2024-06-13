//lint:file-ignore U1000 Ignore all unused code
package main

func moveZeroes(nums []int) {
	for nonZero, cur := 0, 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			nums[nonZero], nums[cur] = nums[cur], nums[nonZero]
			nonZero++
		}
	}
}
