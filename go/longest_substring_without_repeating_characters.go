//lint:file-ignore U1000 Ignore all unused code
package main

func lengthOfLongestSubstring(s string) int {
	runeS := []rune(s)
	maxLen := 0
	head := 0
	using := make(map[rune]struct{})
	for tail := 0; tail < len(runeS); tail++ {
		if _, ok := using[runeS[tail]]; ok {
			for runeS[head] != runeS[tail] {
				delete(using, runeS[head])
				head++
			}
			head++
		} else {
			using[runeS[tail]] = struct{}{}
		}
		substring := runeS[head : tail+1]
		maxLen = max(maxLen, len(substring))
	}
	return maxLen
}
