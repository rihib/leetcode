package sort

func Insertion(nums []int) []int {
	for sorted := 0; sorted < len(nums)-1; sorted++ {
		if nums[sorted] > nums[sorted+1] {
			for j := sorted; j >= 0; j-- {
				if nums[j] <= nums[j+1] {
					break
				}
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
