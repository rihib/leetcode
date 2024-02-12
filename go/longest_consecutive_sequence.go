package blind75

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, n := range nums {
		m[n] = struct{}{}
	}

	max_len := 0
	for n := range m {
		if _, ok := m[n-1]; ok {
			continue
		}

		for l := 1; ; l++ {
			if _, ok := m[n+l]; !ok {
				if l > max_len {
					max_len = l
				}
				break
			}
		}
	}

	return max_len
}