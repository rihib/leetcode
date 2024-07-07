package sort

func Insertion(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		j := i - 1
		for ; j >= 0 && nums[j] > curr; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = curr
	}
	return nums
}
