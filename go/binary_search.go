//lint:file-ignore U1000 Ignore all unused code
package main

func binarySearchHalfClosed(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}

func binarySearchClosed(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
