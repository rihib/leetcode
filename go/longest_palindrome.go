//lint:file-ignore U1000 Ignore all unused code
package main

func longestPalindrome(s string) int {
	m := make(map[rune]struct{})
	l := 0

	for _, r := range s {
		if _, ok := m[r]; ok {
			delete(m, r)
			l += 2
		} else {
			m[r] = struct{}{}
		}
	}

	if len(m) > 0 {
		l++
	}

	return l
}
