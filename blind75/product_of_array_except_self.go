package blind75

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	rp := 1
	for i := 0; i < len(nums); i++ {
		res[i] = rp
		rp *= nums[i]
	}

	lp := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= lp
		lp *= nums[i]
	}

	return res
}
