package blind75

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	lp := 1
	for i := 0; i < len(nums); i++ {
		res[i] = lp
		lp *= nums[i]
	}

	rp := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= rp
		rp *= nums[i]
	}

	return res
}
