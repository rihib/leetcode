//lint:file-ignore U1000 Ignore all unused code
package main

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	var mid int
	for left < right {
		mid = (left + right) / 2
		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if target < nums[mid] {
		return mid
	}
	return mid + 1
}
