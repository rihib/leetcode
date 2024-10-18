//lint:file-ignore U1000 Ignore all unused code
package topkfrequentelements

/*
時間：13分
とりあえずO(n)で解けそうだったので解いてみました。
（あとで他の人のPRを読んでこれがバケットソートと呼ばれているものだと知りました、、）
*/
func topKFrequentStep1(nums []int, k int) []int {
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
