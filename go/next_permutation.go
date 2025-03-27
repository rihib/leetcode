//lint:file-ignore U1000 Ignore all unused code
package main

// [1, 3, 5, 4, 2]という入力があるとき
// 3 < 5なので、3を変える必要がある
// 3より大きい最小の値は4なので、3と4を入れ替えると
// [1, 4, 5, 3, 2]となる
// 3以降の要素を逆順にすると、
// next permutationの[1, 4, 2, 3, 5]となる。
// nums[isSortedUntil] >= nums[isSortedUntil+1]は>ではなく、>=でないと
// [1, 1]などではOut of indexになる
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	isSortedUntil := len(nums) - 2
	for isSortedUntil >= 0 && nums[isSortedUntil] >= nums[isSortedUntil+1] {
		isSortedUntil--
	}
	if isSortedUntil >= 0 {
		swapTarget := len(nums) - 1
		for nums[isSortedUntil] >= nums[swapTarget] {
			swapTarget--
		}
		nums[isSortedUntil], nums[swapTarget] = nums[swapTarget], nums[isSortedUntil]
	}
	reverse(nums[isSortedUntil+1:])
}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
