//lint:file-ignore U1000 Ignore all unused code
package main

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// forループを抜けるのは探索範囲の長さが１以下になったとき
// 長さ２のnums={1, 2}があるとき、
// left=0, right=2, mid=1になる。
// target=3のとき、ループを抜ける直前にleft=mid+1されるので、ループを抜けたらleftを返せばいい。
// target=0のとき、leftはそのままなので、ループを抜けたらleftを返せばいい
