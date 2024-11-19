//lint:file-ignore U1000 Ignore all unused code
package template

import "unicode"

/*
時間：12分

解法自体はすぐに思いついたが、すんなり実装できないところがまだまだ課題。
continueすることを忘れていたりなどケアレスミスが目立った。
*/
func isPalindromeStep1(s string) bool {
	sRunes := []rune(s)
	for i, j := 0, len(s)-1; i < j; {
		if !(unicode.IsDigit(sRunes[i]) || unicode.IsLetter(sRunes[i])) {
			i++
			continue
		}
		if !(unicode.IsDigit(sRunes[j]) || unicode.IsLetter(sRunes[j])) {
			j--
			continue
		}
		if unicode.ToLower(sRunes[i]) != unicode.ToLower(sRunes[j]) {
			return false
		}
		i++
		j--
	}
	return true
}
