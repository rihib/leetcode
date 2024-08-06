//lint:file-ignore U1000 Ignore all unused code
package main

func lengthOfLongestSubstring(s string) int {
	maxLen, head := 0, 0
	inUse := make(map[byte]bool)
	for tail := range s {
		for inUse[s[tail]] {
			delete(inUse, s[head])
			head++
		}
		inUse[s[tail]] = true
		maxLen = max(maxLen, tail-head+1)
	}
	return maxLen
}
