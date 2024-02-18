//lint:file-ignore U1000 Ignore all unused code
package blind75

func containsDuplicate2(nums []int) bool {
	m := make(map[int]struct{})
	for _, n := range nums {
		m[n] = struct{}{}
	}
	if len(nums) > len(m) {
		return true
	}
	return false
}
