//lint:file-ignore U1000 Ignore all unused code
package majorityelement

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
時間：6分

Boyer-Moore Voteアルゴリズムでも解いてみた。このアルゴリズムもたまたま知っていたが、思い出しつつ書いたので少し時間がかかってしまった。
*/
func majorityElementBoyerMooreMajorityVote(nums []int) int {
	candidate, count := nums[0], 0
	for _, n := range nums {
		if n == candidate {
			count++
		} else {
			count--
		}
		if count < 0 {
			candidate = n
			count = 0
		}
	}
	return candidate
}
