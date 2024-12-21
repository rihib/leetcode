//lint:file-ignore U1000 Ignore all unused code
package template

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
0~len(nums)の和を求めるには算数的に求める方法も取れる。
*/
func missingNumberStep2(nums []int) int {
	var total int
	for _, n := range nums {
		total += n
	}
	return (1+len(nums))*len(nums)/2 - total
}
