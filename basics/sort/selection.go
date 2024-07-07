package sort

func Selection(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		min, minIdx := nums[i], i
		for j := i + 1; j < len(nums); j++ {
			if min > nums[j] {
				min, minIdx = nums[j], j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
	return nums
}
