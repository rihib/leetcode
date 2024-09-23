//lint:file-ignore U1000 Ignore all unused code
package longestsubstringwithoutrepeatingcharacters

/*
より変数名をわかりやすくし、無駄な部分を消しました。
*/
func lengthOfLongestSubstring(s string) int {
	maxLen, head := 0, 0
	inUse := make(map[byte]struct{})
	for tail := range s {
		if _, ok := inUse[s[tail]]; ok {
			for s[head] != s[tail] {
				delete(inUse, s[head])
				head++
			}
			head++
		} else {
			inUse[s[tail]] = struct{}{}
		}
		maxLen = max(maxLen, tail-head+1)
	}
	return maxLen
}
