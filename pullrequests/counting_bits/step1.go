//lint:file-ignore U1000 Ignore all unused code
package template

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
時間：11分
問題を理解するのに3分かかってしまった（問題文の英語が難しかった、、）
解法を思いつくのにも少し時間がかかってしまった。
この方法だと内側のループでO(log n)、外側のループでO(n)かかるので、全体でO(n log n)かかる。
*/
func countBitsStep1(n int) []int {
	bitCounts := make([]int, n+1)
	for i := 0; i <= n; i++ {
		num := i
		count := 0
		for num > 0 {
			count += num % 2
			num /= 2
		}
		bitCounts[i] = count
	}
	return bitCounts
}
