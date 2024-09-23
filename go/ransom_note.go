//lint:file-ignore U1000 Ignore all unused code
package main

func canConstruct(ransomNote string, magazine string) bool {
	var frequencies [26]int
	for _, r := range magazine {
		frequencies[r-'a']++
	}
	for _, r := range ransomNote {
		frequencies[r-'a']--
		if frequencies[r-'a'] < 0 {
			return false
		}
	}
	return true
}
