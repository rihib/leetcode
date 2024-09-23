//lint:file-ignore U1000 Ignore all unused code
package main

import "sort"

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

func binarySearchHalfClosed2(nums []int, target int) int {
	left, right := 0, len(nums)
	for {
		if left == right {
			break
		}
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

func binarySearchHalfClosed3(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func binarySearchHalfClosed4(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if right >= len(nums) || nums[right] != target {
		return -1
	}
	return right
}

func binarySearchClosed1(nums []int, target int) int {
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

func binarySearchClosed2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func search(nums []int, target int) int {
	index := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if index < len(nums) && nums[index] == target {
		return index
	}
	return -1
}
