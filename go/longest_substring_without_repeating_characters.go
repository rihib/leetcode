//lint:file-ignore U1000 Ignore all unused code
package main

func lengthOfLongestSubstring(s string) int {
	maxLen, head := 0, 0
	using := make(map[byte]struct{})
	for tail := range s {
		if _, ok := using[s[tail]]; ok {
			for s[head] != s[tail] {
				delete(using, s[head])
				head++
			}
			head++
		} else {
			using[s[tail]] = struct{}{}
		}
		maxLen = max(maxLen, tail-head+1)
	}
	return maxLen
}
