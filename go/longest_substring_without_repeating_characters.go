//lint:file-ignore U1000 Ignore all unused code
package main

// 重複する文字が含まれていないことが保証されるleft~rightの範囲を作り、
// rightを１つずつ進めていき、各段階での長さを求め、最大の長さを更新していく。
func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	seen := make(map[byte]int, len(s))
	for left, right := 0, 0; right < len(s); right++ {
		c := s[right]
		if i, ok := seen[c]; ok && left <= i {
			left = i + 1
		}
		seen[c] = right
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}
