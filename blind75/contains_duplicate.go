package blind75

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	m := make(map[int]struct{})
	for _, n := range nums {
		if _, found := m[n]; found {
			return true
		}
		m[n] = struct{}{}
	}
	return false
}
