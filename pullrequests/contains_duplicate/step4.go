//lint:file-ignore U1000 Ignore all unused code
package containsduplicate

func containsDuplicateStep4(nums []int) bool {
	seen := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := seen[n]; ok {
			return true
		}
		seen[n] = struct{}{}
	}
	return false
}

// 若干トリッキーで意図がわかりづらいので上の方がベター
func containsDuplicate2Step4(nums []int) bool {
	numsMap := make(map[int]struct{})
	for _, n := range nums {
		numsMap[n] = struct{}{}
	}
	return len(nums) > len(numsMap)
}
