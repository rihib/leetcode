//lint:file-ignore U1000 Ignore all unused code
package searchinsertposition

/*
最後の部分はleftを返せ馬良いと気づき変更した。
その他もより読みやすいように直した。
*/
func searchInsert_step2(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
