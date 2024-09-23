//lint:file-ignore U1000 Ignore all unused code
package main

import "unicode"

func isPalindrome(s string) bool {
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
