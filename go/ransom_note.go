//lint:file-ignore U1000 Ignore all unused code
package main

func canConstruct(ransomNote string, magazine string) bool {
	var freq [26]int

	for _, r := range magazine {
		freq[r-'a']++
	}

	for _, r := range ransomNote {
		freq[r-'a']--
		if freq[r-'a'] < 0 {
			return false
		}
	}
	return true
}
