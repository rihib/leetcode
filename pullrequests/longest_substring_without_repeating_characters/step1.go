//lint:file-ignore U1000 Ignore all unused code
package longestsubstringwithoutrepeatingcharacters

/*
時間：34分
最初に下記のやり方でやることを思いついたが、もう少し賢いやり方があるのではないかと10分ぐらい考えてしまい、思いつかなかったので、とりあえず書こうと思い、書きました。結果的には他の人のコードも見たりする限り、この方法で良かったのだとわかりました。
*/
func lengthOfLongestSubstring_step1(s string) int {
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
