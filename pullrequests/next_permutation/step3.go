//lint:file-ignore U1000 Ignore all unused code
package nextpermutation

/*
これまでの解法がソートを使っていたために時間計算量がO(n log n)になってしまったわけですが、他の人の回答を見てO(n)で解ける方法に変えました。
nums[i-1]よりも後ろは昇順になっているのを利用し、リバースすればソートできるということに気づきませんでした。
*/
func nextPermutation_step3(nums []int) {
	if len(nums) < 2 {
		return
	}
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
