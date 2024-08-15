//lint:file-ignore U1000 Ignore all unused code
package main

func isSubsequence(s string, t string) bool {
	current := 0
	for i := 0; i < len(s); i++ {
		for {
			if current >= len(t) {
				return false
			}
			if s[i] == t[current] {
				current++
				break
			}
			current++
		}
	}
	return true
}
