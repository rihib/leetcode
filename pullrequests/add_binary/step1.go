//lint:file-ignore U1000 Ignore all unused code
package addbinary

import (
	"slices"
	"strings"
)

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
解法自体はすぐに思いついたが、文字列操作のメソッドの使い方の記憶が曖昧だったので、調べながら書いた。
*/
func addBinary(a string, b string) string {
	var reversed strings.Builder
	maxLength := max(len(a), len(b))
	carry := 0
	for i := 1; i <= maxLength; i++ {
		bitA, bitB := 0, 0
		if i <= len(a) {
			bitA = int(a[len(a)-i] - '0')
		}
		if i <= len(b) {
			bitB = int(b[len(b)-i] - '0')
		}
		sum := bitA + bitB + carry
		carry = sum / 2
		reversed.WriteByte(byte(sum%2 + '0'))
	}
	if carry == 1 {
		reversed.WriteByte(byte(carry + '0'))
	}
	result := []rune(reversed.String())
	slices.Reverse(result)
	return string(result)
}
