//lint:file-ignore U1000 Ignore all unused code
package main

func containsDuplicate2(nums []int) bool {
	m := make(map[int]struct{})
	for _, n := range nums {
		m[n] = struct{}{}
	}
	return len(nums) > len(m)
}
