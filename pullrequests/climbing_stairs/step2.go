//lint:file-ignore U1000 Ignore all unused code
package template

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
Step1の方法はどちらかというと特殊で、メモ化再帰を使った動的計画法で解くのが本命のような気がしたので、こちらでも解いた。
*/
func climbStairsDP(n int) int {
	m := make(map[int]int, n)
	return climbStairsHelper(n, m)
}

func climbStairsHelper(n int, m map[int]int) int {
	if n == 1 || n == 2 {
		return n
	}
	if ways, ok := m[n]; ok {
		return ways
	}
	m[n] = climbStairsHelper(n-1, m) + climbStairsHelper(n-2, m)
	return m[n]
}
