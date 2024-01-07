package blind75

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	nm := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := nm[v]; ok {
			return true
		}
		nm[v] = struct{}{}
	}
	return false
}
