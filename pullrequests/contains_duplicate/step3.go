//lint:file-ignore U1000 Ignore all unused code
package containsduplicate

func containsDuplicateStep3(nums []int) bool {
	seen := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := seen[n]; ok {
			return true
		}
		seen[n] = struct{}{}
	}
	return false
}
