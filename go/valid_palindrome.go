//lint:file-ignore U1000 Ignore all unused code
package main

import "unicode"

func isPalindrome(s string) bool {
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
