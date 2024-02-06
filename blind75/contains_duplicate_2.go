package blind75

func containsDuplicate2(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	m := make(map[int]struct{})
	for _, n := range nums {
		m[n] = struct{}{}
	}

	if len(nums) != len(m) {
		return true
	}
	return false
}
