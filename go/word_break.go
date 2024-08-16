//lint:file-ignore U1000 Ignore all unused code
package main

func wordBreak(s string, wordDict []string) bool {
	canBreak := make([]bool, len(s)+1)
	canBreak[len(s)] = true
	for i := len(s) - 1; i >= 0; i-- {
		for _, w := range wordDict {
			if (i+len(w)) <= len(s) && s[i:i+len(w)] == w {
				canBreak[i] = canBreak[i+len(w)]
			}
			if canBreak[i] {
				break
			}
		}
	}
	return canBreak[0]
}
