//lint:file-ignore U1000 Ignore all unused code
package main

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	head := 0
	m := make(map[rune]struct{})
	for tail := 1; tail <= len(s); tail++ {
		if _, ok := m[rune(s[tail-1])]; ok {
			for s[head] != s[tail-1] {
				delete(m, rune(s[head]))
				head++
			}
			head++
		} else {
			m[rune(s[tail-1])] = struct{}{}
		}
		substring := s[head:tail]
		maxLen = max(maxLen, len(substring))
	}
	return maxLen
}
