//lint:file-ignore U1000 Ignore all unused code
package main

func searchClosed(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// 半閉区画ではleft~right-1が未探索で、rightは常に探索済み
// そのため、for文の条件はleft < rightとなる（left=1, right=2のときmid=1となり、未探索のleftのみが探索の対象となる）
func searchHalfClosed(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] < nums[mid] {
			if nums[left] <= target && target < nums[mid] { // leftがtargetの可能性もあるので<=とする
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right-1] { // rightは最初は範囲外なので-1をする。その場合right-1は未探索なので<=とする
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	return -1
}
