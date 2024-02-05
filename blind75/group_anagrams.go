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

	result := make([][]string, len(m))
	i := 0
	for _, v := range m {
		result[i] = v
		i++
	}
	return result
}
