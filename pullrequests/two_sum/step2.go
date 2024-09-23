//lint:file-ignore U1000 Ignore all unused code
package twosum

/*
mではなく、よりわかりやすいようにnumsMapとしました。
GoogleのGoスタイルガイドには変数名に型名を使うのは良くないと書かれていますが、同時に下記のようにも書かれています。今回はnumsという配列をマップに変換したもの（配列もインデックスと値を情報としてもつ）と捉えることができるため、対応していることを示すためにnumsMapとしました。

`It is acceptable to include a type-like qualifier if there are two versions of a value in scope, for example you might have an input stored in ageString and use age for the parsed value.`

`numToIdx`というのもありそう（https://github.com/aoshi2025s/leetcode-review/pull/1#discussion_r1666780953）。

対応する組み合わせが見つからなかった際にどうするのかは難しいところ。
  - https://github.com/seal-azarashi/leetcode/pull/11#discussion_r1672537855
  - https://github.com/sendahuang14/leetcode/pull/11#discussion_r1702393602
*/
func twoSumStep2(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, n := range nums {
		if j, ok := numsMap[target-n]; ok {
			return []int{i, j}
		}
		numsMap[n] = i
	}
	return nil
}
