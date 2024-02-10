package blind75

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	cs := make([][]int, len(nums))
	for n, c := range freq {
		cs[c-1] = append(cs[c-1], n)
	}

	var res []int
	for i := len(cs) - 1; i >= 0; i-- {
		res = append(res, cs[i]...)
		if len(res) >= k {
			break
		}
	}
	return res
}
