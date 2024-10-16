//lint:file-ignore U1000 Ignore all unused code
package template

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
DPのメモカを利用すればO(n)で解ける。
2の冪乗のときは必ず１になるし、それ以外の場合は前に計算した結果+1すれば良い。
*/
func countBitsStep2(n int) []int {
	bitCounts := make([]int, n+1)
	powerOfTwo := 1
	for i := 1; i <= n; i++ {
		if i == powerOfTwo*2 {
			powerOfTwo = i
		}
		bitCounts[i] = 1 + bitCounts[i-powerOfTwo]
	}
	return bitCounts
}
