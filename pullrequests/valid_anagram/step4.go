//lint:file-ignore U1000 Ignore all unused code
package validanagram

func isAnagramStep4(s string, t string) bool {
	return frequencies(s) == frequencies(t)
}

func frequencies(s string) [26]int {
	var f [26]int
	for _, r := range s {
		f[r-'a']++
	}
	return f
}

// ちゃんとUnicodeに対応させるなら結合文字などを考慮する必要がある
// https://github.com/rihib/leetcode/pull/5#discussion_r1706198268
func isAnagramUnicodeStep4(s string, t string) bool {
	frequencies := make(map[rune]int)
	for _, r := range s {
		frequencies[r]++
	}
	for _, r := range t {
		frequencies[r]--
	}
	for _, n := range frequencies {
		if n != 0 {
			return false
		}
	}
	return true
}
