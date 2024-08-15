//lint:file-ignore U1000 Ignore all unused code
package main

func isSubsequence(s string, t string) bool {
	current := 0
	for i := 0; i < len(t); i++ {
		if s[current] == t[i] {
			current++
		}
		if current == len(s) {
			return true
		}
	}
	return false
}
