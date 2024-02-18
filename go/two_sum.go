//lint:file-ignore U1000 Ignore all unused code
package blind75

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{i, j}
		}
		m[n] = i
	}
	return nil
}
