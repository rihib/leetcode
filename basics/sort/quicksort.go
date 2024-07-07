package sort

func Quicksort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pivot := partition(nums)
	Quicksort(nums[:pivot])
	Quicksort(nums[pivot+1:])
	return nums
}

func partition(nums []int) int {
	l, r := 1, len(nums)-1
	for l <= r {
		for l <= r && nums[l] <= nums[0] {
			l++
		}
		for l <= r && nums[r] > nums[0] {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[0], nums[r] = nums[r], nums[0]
	return r
}
