//lint:file-ignore U1000 Ignore all unused code
package palindromenumber

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
時間：8分50秒
最初に思いついたのは文字列に変換してチェックする方法だが、LeetCode上に文字列に変換しないで求めてみよと書かれていたのでまずは下記のようにして解いた。
*/
func isPalindromeStep1(x int) bool {
	var reversed int
	tmp := x
	for tmp > 0 {
		n := tmp % 10
		reversed = reversed*10 + n
		tmp /= 10
	}
	return x == reversed
}
