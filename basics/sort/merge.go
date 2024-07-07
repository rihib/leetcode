package sort

func Merge(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	left := Merge(nums[:len(nums)/2])
	right := Merge(nums[len(nums)/2:])

	i, l, r := 0, 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			nums[i] = left[l]
			l++
		} else {
			nums[i] = right[r]
			r++
		}
		i++
	}
	for ; l < len(left); l++ {
		nums[i] = left[l]
		i++
	}
	for ; r < len(right); r++ {
		nums[i] = right[r]
		i++
	}

	return nums
}
