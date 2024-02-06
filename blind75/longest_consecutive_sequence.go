package blind75

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, n := range nums {
		m[n] = struct{}{}
	}

	max_len := 0
	for n := range m {
		if _, found := m[n-1]; found {
			continue
		}

		l := 1
		for i := 1; ; i++ {
			if _, found := m[n+i]; found {
				l++
			} else {
				break
			}
		}

		if l > max_len {
			max_len = l
		}
	}

	return max_len
}
