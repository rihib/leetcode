package blind75

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if j, found := m[target-n]; found {
			return []int{i, j}
		}
		m[n] = i
	}
	return nil
}
