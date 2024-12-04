//lint:file-ignore U1000 Ignore all unused code
package longestcommonprefix

import "strings"

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
時間：５分
まず最初に思いついた方法は各文字をそれぞれ見ていくという方法。素直に実装した。
*/
func longestCommonPrefix(strs []string) string {
	var prefix strings.Builder
	for i := 0; ; i++ {
		var curr rune
		for _, word := range strs {
			if i >= len(word) {
				return prefix.String()
			}
			if curr == 0 {
				curr = rune(word[i])
				continue
			}
			if rune(word[i]) != curr {
				return prefix.String()
			}
		}
		prefix.WriteRune(curr)
	}
}
