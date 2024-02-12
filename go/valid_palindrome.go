package blind75

import "unicode"

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		ri := rune(s[i])
		rj := rune(s[j])

		if !(unicode.IsLetter(ri) || unicode.IsDigit(ri)) {
			i++
			continue
		}
		if !(unicode.IsLetter(rj) || unicode.IsDigit(rj)) {
			j--
			continue
		}

		if unicode.ToLower(ri) != unicode.ToLower(rj) {
			return false
		}

		i++
		j--
	}
	return true
}
