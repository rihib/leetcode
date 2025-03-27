//lint:file-ignore U1000 Ignore all unused code
package template

/*
O(log n)で解くことを期待されていたので、何らかの形で二分探索を使うのだろうとは思っていたのですが、初見では解けなかったので、他の人の回答を参考にしました。
どちらか片方は必ず必ず昇順にならんでいるということを利用して二分探索を行っています。
下記は閉区画で解きました。
*/
func search_closed(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
