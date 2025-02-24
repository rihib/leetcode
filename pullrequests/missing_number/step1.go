//lint:file-ignore U1000 Ignore all unused code
package template

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
方針としては0~len(nums)の和と実際の和の差分を取ればいい。
*/
func missingNumberStep1(nums []int) int {
	var difference int
	for i, n := range nums {
		difference += i + 1 - n
	}
	return difference
}
