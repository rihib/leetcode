//lint:file-ignore U1000 Ignore all unused code
package backspacestringcompare

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
スタックを使って前から処理していく方法と、後ろから見ていく方法の２つを考えた。
スタックを使う場合はシンプルに実装できるが、O(n)の追加の空間が必要になる。
後ろから見ていく方法は追加の空間が必要なくなるが、実装が複雑になる。
個人的には必要がなければできるだけスタックを使ってシンプルに実装する方が良いと思っている。
*/
func backspaceCompareStep1(s string, t string) bool {
	return typedText(s) == typedText(t)
}

func typedText(s string) string {
	var stack []rune
	for _, r := range s {
		if r == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, r)
	}
	return string(stack)
}
