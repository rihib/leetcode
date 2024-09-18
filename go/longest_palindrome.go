//lint:file-ignore U1000 Ignore all unused code
package main

func longestPalindrome(s string) int {
	length := 0
	frequencies := make(map[rune]struct{})
	for _, r := range s {
		if _, ok := frequencies[r]; !ok {
			frequencies[r] = struct{}{}
		} else {
			delete(frequencies, r)
			length += 2
		}
	}
	if len(frequencies) > 0 {
		length++
	}
	return length
}
