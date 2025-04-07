//lint:file-ignore U1000 Ignore all unused code
package main

func firstUniqChar(s string) int {
	frequencies := make(map[rune]int)
	for _, r := range s {
		frequencies[r]++
	}
	for i, r := range s {
		if frequencies[r] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar2(s string) int {
	var frequencies [26]int
	for _, r := range s {
		frequencies[r-'a']++
	}
	for i, r := range s {
		if frequencies[r-'a'] == 1 {
			return i
		}
	}
	return -1
}
