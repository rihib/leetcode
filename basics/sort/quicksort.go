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

// Using the two-pointer technique from both ends for better performance
// The single-direction approach starting from the beginning is
// more likely to result in skewed partitions
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
