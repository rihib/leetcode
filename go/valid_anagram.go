//lint:file-ignore U1000 Ignore all unused code
package main

func isAnagram(s string, t string) bool {
	return frequency(s) == frequency(t)
}

func frequency(s string) [26]int {
	var f [26]int
	for _, r := range s {
		f[r-'a']++
	}
	return f
}

// ちゃんとUnicodeに対応させるなら結合文字などを考慮する必要がある
// https://github.com/rihib/leetcode/pull/5#discussion_r1706198268
func isAnagramUnicode(s string, t string) bool {
	frequency := make(map[rune]int)
	for _, r := range s {
		frequency[r]++
	}
	for _, r := range t {
		frequency[r]--
	}
	for _, n := range frequency {
		if n != 0 {
			return false
		}
	}
	return true
}
