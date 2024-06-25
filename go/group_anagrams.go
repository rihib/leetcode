//lint:file-ignore U1000 Ignore all unused code
package main

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, word := range strs {
		var freq [26]int
		for _, r := range word {
			freq[r-'a']++
		}
		m[freq] = append(m[freq], word)
	}

	res := make([][]string, len(m))
	i := 0
	for _, words := range m {
		res[i] = words
		i++
	}
	return res
}
