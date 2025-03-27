//lint:file-ignore U1000 Ignore all unused code
package nextpermutation

import "sort"

/*
時間：36分
解法を思いつくのに時間がかかってしまいました（15分ぐらい）。
後ろから見ていって昇順になっていない箇所を見つけたら、nums[i-1]よりも大きい中で最小の値を見つけてスワップし、後ろの要素をソートするというやり方を取りました。
*/
func nextPermutation_step1(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			minN := nums[i]
			idx := i
			for j := i; j < len(nums); j++ {
				if nums[j] < minN && nums[i-1] < nums[j] {
					idx = j
					minN = nums[j]
				}
			}
			nums[i-1], nums[idx] = nums[idx], nums[i-1]
			sort.Slice(nums[i:], func(a, b int) bool {
				return nums[i+a] < nums[i+b]
			})
			return
		}
	}
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})
}
