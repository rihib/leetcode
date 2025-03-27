//lint:file-ignore U1000 Ignore all unused code
package template

/*
半閉区画でも解いてみました。
*/
func search_half_closed(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right-1] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	return -1
}
