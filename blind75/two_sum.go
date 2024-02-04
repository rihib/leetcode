package blind75

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if v, found := m[target-n]; found {
			return []int{i, v}
		}
		m[n] = i
	}
	return nil
}
