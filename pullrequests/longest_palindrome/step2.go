//lint:file-ignore U1000 Ignore all unused code
package longestpalindrome

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
個人的にはStep1の方がわかりやすくて好きだが、こういう書き方もあると思うので書いてみた。
*/
func longestPalindromeStep2(s string) int {
	length := 0
	frequencies := make(map[rune]struct{})
	for _, r := range s {
		if _, ok := frequencies[r]; !ok {
			frequencies[r] = struct{}{}
		} else {
			delete(frequencies, r)
			length += 2
		}
	}
	if len(frequencies) > 0 {
		length++
	}
	return length
}
