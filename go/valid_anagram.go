//lint:file-ignore U1000 Ignore all unused code
package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freq [26]int
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}

	for _, v := range freq {
		if v != 0 {
			return false
		}
	}
	return true
}
