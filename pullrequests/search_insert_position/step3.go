//lint:file-ignore U1000 Ignore all unused code
package searchinsertposition

/*
オーバーフローを起こさないようにmidの計算方法を修正した
https://github.com/hayashi-ay/leetcode/blob/064fca989bc4ecf9c7bce70237524a3e7ab1a21a/35.%20Search%20Insert%20Position.md?plain=1#L6
*/
func searchInsert_step3(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
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
