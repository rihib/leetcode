//lint:file-ignore U1000 Ignore all unused code
package template

import "unicode"

/*
Two Pointersなのでi, jではなく、left, rightを使うようにした。

質問：sをrune型配列にキャストした変数名を`sRunes`や`runeS`などにしてみたのですが、あまりしっくりきてないです。良い名前や`sRunes`と`runeS`ならどちらの方が良さげなどありますか
*/
func isPalindromeStep2(s string) bool {
	runeS := []rune(s)
	left, right := 0, len(s)-1
	for left < right {
		if !(unicode.IsDigit(runeS[left]) || unicode.IsLetter(runeS[left])) {
			left++
			continue
		}
		if !(unicode.IsDigit(runeS[right]) || unicode.IsLetter(runeS[right])) {
			right--
			continue
		}
		if unicode.ToLower(runeS[left]) != unicode.ToLower(runeS[right]) {
			return false
		}
		left++
		right--
	}
	return true
}
