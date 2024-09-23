//lint:file-ignore U1000 Ignore all unused code
package main

func wordBreak(s string, wordDict []string) bool {
	canBreak := make([]bool, len(s)+1)
	canBreak[0] = true
	for i := 1; i <= len(s); i++ {
		for _, w := range wordDict {
			if i >= len(w) && canBreak[i-len(w)] && s[i-len(w):i] == w {
				canBreak[i] = true
				break
			}
		}
	}
	return canBreak[len(s)]
}
