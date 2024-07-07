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
	i := 1
	for j := 1; j < len(nums); j++ {
		if nums[j] < nums[0] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i-1], nums[0] = nums[0], nums[i-1]
	return i - 1
}
