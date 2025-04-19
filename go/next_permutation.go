//lint:file-ignore U1000 Ignore all unused code
package main

import "sort"

// [1, 3, 5, 4, 2]という入力があるとき
// 3 < 5なので、3を変える必要がある
// 3より大きい最小の値は4なので、3と4を入れ替えると
// [1, 4, 5, 3, 2]となる
// 3以降の要素を逆順にすると、
// next permutationの[1, 4, 2, 3, 5]となる。
// nums[sortedUntil] >= nums[sortedUntil+1]は>ではなく、>=でないと
// [1, 1]などではOut of indexになる
func nextPermutation(nums []int) {
	sortedUntil := len(nums) - 2 // len(nums) <= 1だとsortedUntil=-1になる点に注意
	for sortedUntil >= 0 && nums[sortedUntil] >= nums[sortedUntil+1] {
		sortedUntil--
	}
	if sortedUntil >= 0 {
		swapTarget := len(nums) - 1
		for nums[sortedUntil] >= nums[swapTarget] {
			swapTarget--
		}
		nums[sortedUntil], nums[swapTarget] = nums[swapTarget], nums[sortedUntil]
	}
	sort.Ints(nums[sortedUntil+1:]) // sortを使うとO(n log n)になるので、別途reverse関数を作るのもあり
}
