package sort

func Selection(nums []int) []int {
	// Selection Sort is unstable as shown below:
	// 1. {2', 1, 2, 1}
	// 2. {1, 2', 2, 1}
	// 3. {1, 1, 2, 2'}
	for i := 0; i < len(nums)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[minIdx] > nums[j] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
	return nums
}
