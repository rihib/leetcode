package blind75

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	nm := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := nm[n]; ok {
			return true
		}
		nm[n] = struct{}{}
	}
	return false
}
