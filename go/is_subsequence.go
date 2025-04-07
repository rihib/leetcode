//lint:file-ignore U1000 Ignore all unused code
package main

func isSubsequence(s string, t string) bool {
	sIndex, tIndex := 0, 0
	for sIndex < len(s) && tIndex < len(t) {
		if s[sIndex] == t[tIndex] {
			sIndex++
		}
		tIndex++
	}
	return sIndex == len(s)
}

func isSubsequence2(s string, t string) bool {
	sIndex := 0
	runeS := []rune(s)
	for _, r := range t {
		if sIndex < len(runeS) && r == runeS[sIndex] {
			sIndex++
		}
	}
	return sIndex == len(runeS)
}
