//lint:file-ignore U1000 Ignore all unused code
package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring_step4(s string) int {
	maxLength, left, seen := 0, 0, make(map[rune]int)
	for right, r := range s {
		if lastIndex, ok := seen[r]; ok && lastIndex >= left {
			left = lastIndex + 1
		}
		seen[r] = right
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}
