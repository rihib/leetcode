//lint:file-ignore U1000 Ignore all unused code
package main

func canConstruct(ransomNote string, magazine string) bool {
	var frequency [26]int
	for _, r := range magazine {
		frequency[r-'a']++
	}
	for _, r := range ransomNote {
		frequency[r-'a']--
		if frequency[r-'a'] < 0 {
			return false
		}
	}
	return true
}
