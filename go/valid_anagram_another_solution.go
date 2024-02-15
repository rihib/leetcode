package blind75

func isAnagram_unicode(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freq := make(map[rune]int)
	for _, r := range s {
		freq[r]++
	}
	for _, r := range t {
		freq[r]--
		if freq[r] < 0 {
			return false
		}
	}
	return true
}
