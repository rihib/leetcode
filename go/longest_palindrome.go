//lint:file-ignore U1000 Ignore all unused code
package main

func longestPalindrome(s string) int {
	length := 0
	frequencies := make(map[rune]int, len(s))
	for _, r := range s {
		frequencies[r]++
		if frequencies[r] == 2 {
			length += 2
			delete(frequencies, r)
		}
	}
	if len(frequencies) != 0 {
		length++
	}
	return length
}
