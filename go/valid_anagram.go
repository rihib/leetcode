//lint:file-ignore U1000 Ignore all unused code
package main

func isAnagram(s string, t string) bool {
	var frequency [26]int
	for _, r := range s {
		frequency[r-'a']++
	}
	for _, r := range t {
		frequency[r-'a']--
	}
	for _, n := range frequency {
		if n != 0 {
			return false
		}
	}
	return true
}

func isAnagram_unicode(s string, t string) bool {
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
