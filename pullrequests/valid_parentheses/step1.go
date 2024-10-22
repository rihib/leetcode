//lint:file-ignore U1000 Ignore all unused code
package validparentheses

/*
時間：11分
やけに時間がかかってしまった。まだまだコーディング力が弱い。
これぐらいのネストなら良いかなと思って書いたのですが、案外見辛いのと、変数名も良くない。

またこの部分`if len(stack) == 0`に気づかずエラーを出してしまった。
*/
func isValidStep1(s string) bool {
	var stack []rune
	brackets := map[rune]rune{')': '(', '}': '{', ']': '['}
	for _, bracket := range s {
		if b, ok := brackets[bracket]; !ok {
			stack = append(stack, bracket)
		} else {
			if len(stack) == 0 {
				return false
			}
			if b != stack[len(stack)-1] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}
