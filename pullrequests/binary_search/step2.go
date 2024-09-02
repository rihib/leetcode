//lint:file-ignore U1000 Ignore all unused code
package binarysearch

import "sort"

/*
PythonのbisectみたいなものがGoにもないかと調べたところ、sort.Search関数が使えそうだったので実装してみました。
*/
func search(nums []int, target int) int {
	index := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if index < len(nums) && nums[index] == target {
		return index
	}
	return -1
}
