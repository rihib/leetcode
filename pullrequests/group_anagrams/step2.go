//lint:file-ignore U1000 Ignore all unused code
package groupanagrams

/*
	変数名を改善した。
*/

func groupAnagramsStep2(strs []string) [][]string {
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
