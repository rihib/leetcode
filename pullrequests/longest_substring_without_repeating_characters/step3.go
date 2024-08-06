//lint:file-ignore U1000 Ignore all unused code
package longestsubstringwithoutrepeatingcharacters

/*
inUseの値にbool値を入れればよりシンプルに書けると思ったので、書き換えました。

  - inUseの値に最後に見たインデックスを入れることも考えたが、どちらにせよ`delete(inUse, s[head])`でそこまでの文字を消していく必要があるので、意味がないのでは？
  - head, tailか、left, rightのどちらの方がベターか。substringの先頭・末尾と捉えるのであれば前者、Two Pointersと捉えるのであれば後者？head, tailの方がinclusiveだとわかるのでより良いのでは？（https://github.com/sakzk/leetcode/pull/3#discussion_r1591191532）
  - inUseという変数名はどうか？
    ー 確かにsliding windowと捉えることもできる（https://github.com/sakzk/leetcode/pull/3#discussion_r1591194013）
*/
func lengthOfLongestSubstring_step3(s string) int {
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
