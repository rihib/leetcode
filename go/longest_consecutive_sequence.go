//lint:file-ignore U1000 Ignore all unused code
package main

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, n := range nums {
		m[n] = struct{}{}
	}

	maxLen := 0
	for n := range m {
		if _, ok := m[n-1]; ok {
			continue
		}

		for l := 1; ; l++ {
			if _, ok := m[n+l]; !ok {
				if l > maxLen {
					maxLen = l
				}
				break
			}
		}
	}

	return maxLen
}
