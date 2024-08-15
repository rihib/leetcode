//lint:file-ignore U1000 Ignore all unused code
package main

func isSubsequence(s string, t string) bool {
	i := 0
	for _, r := range s {
		for {
			if i >= len(t) {
				return false
			}
			if r == rune(t[i]) {
				i++
				break
			}
			i++
		}
	}
	return true
}
