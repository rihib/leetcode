//lint:file-ignore U1000 Ignore all unused code
package palindromenumber

import "strconv"

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
一応、文字列に変換する方法でも解いてみた。
*/
func isPalindromeStep2(x int) bool {
	strX := strconv.Itoa(x)
	i, j := 0, len(strX)-1
	for i < j {
		if strX[i] != strX[j] {
			return false
		}
		i++
		j--
	}
	return true
}
