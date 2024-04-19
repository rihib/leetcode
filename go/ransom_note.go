//lint:file-ignore U1000 Ignore all unused code
package main

func canConstruct(ransomNote string, magazine string) bool {
	var l [26]int

	for _, r := range magazine {
		l[r-'a']++
	}

	for _, r := range ransomNote {
		if l[r-'a'] <= 0 {
			return false
		}
		l[r-'a']--
	}
	return true
}
