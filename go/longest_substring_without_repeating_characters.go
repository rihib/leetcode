//lint:file-ignore U1000 Ignore all unused code
package main

// 各文字においてその段階でのSubstringの長さを求め、maxLengthを更新
// rightは現在の文字のインデックスを表す
// 常にleftとrightの間には重複が存在しないようにleftを管理する
// なのでrightを更新した際にleftとrightの間を見て。rightと重複する文字が存在するならその文字を含まないようにleftを更新する
// そうすることで常にleftとrightの間には重複が存在しないことを保証できるので、rightにおけるSubstringの長さはright-left+1で求めることができる
func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	left := 0
	seen := make(map[rune]int)
	for right, r := range s {
		if lastSeen, ok := seen[r]; ok && left <= lastSeen {
			left = lastSeen + 1
		}
		seen[r] = right
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}
