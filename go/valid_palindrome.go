package blind75

import "unicode"

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		si := rune(s[i])
		sj := rune(s[j])

		if !(unicode.IsLetter(si) || unicode.IsDigit(si)) {
			i++
			continue
		}
		if !(unicode.IsLetter(sj) || unicode.IsDigit(sj)) {
			j--
			continue
		}

		if unicode.ToLower(si) != unicode.ToLower(sj) {
			return false
		}

		i++
		j--
	}
	return true
}
