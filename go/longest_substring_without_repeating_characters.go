//lint:file-ignore U1000 Ignore all unused code
package main

// 重複する文字が含まれていない文字列の左端をleft、右端をrightとしたとき、
// rightが取りうる位置はlen(s)個であり、rightの位置が決まればleftの位置も決まるため、
// right-left+1を全て比較して最大の長さを求めれば良い。
func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	left := 0
	seen := make(map[rune]int)
	for right, r := range s {
		if i, ok := seen[r]; ok && left <= i {
			left = i + 1
		}
		maxLength = max(maxLength, right-left+1)
		seen[r] = right
	}
	return maxLength
}
