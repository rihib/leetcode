//lint:file-ignore U1000 Ignore all unused code
package blind75

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, s := range strs {
		var freq [26]int
		for _, r := range s {
			freq[r-'a']++
		}
		m[freq] = append(m[freq], s)
	}

	res := make([][]string, len(m))
	i := 0
	for _, v := range m {
		res[i] = v
		i++
	}
	return res
}
