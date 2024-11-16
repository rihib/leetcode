//lint:file-ignore U1000 Ignore all unused code
package main

import "unicode"

func isPalindrome(s string) bool {
	runeS := []rune(s)
	i, j := 0, len(s)-1
	for i < j {
		if !unicode.IsDigit(runeS[i]) && !unicode.IsLetter(runeS[i]) {
			i++
			continue
		}
		if !unicode.IsDigit(runeS[j]) && !unicode.IsLetter(runeS[j]) {
			j--
			continue
		}
		if unicode.ToLower(runeS[i]) != unicode.ToLower(runeS[j]) {
			return false
		}
		i++
		j--
	}
	return true
}
