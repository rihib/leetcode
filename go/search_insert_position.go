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

// 長さ１のnums={1}があるとき、
// left=0, right=1, mid=0になる。
// target=2のとき、ループを抜ける直前にleft=mid+1されるので、ループを抜けたらleftを返せばいい。
// target=0のとき、left=0のままなので、ループを抜けたらleftを返せばい（-1を返すわけではないことに注意。targetはleftよりも小さく、leftの左の値よりも大きいわけなので、leftを返せば良い）
