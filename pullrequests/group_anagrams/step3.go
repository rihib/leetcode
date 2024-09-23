//lint:file-ignore U1000 Ignore all unused code
package groupanagrams

/*
	時間：3分
	変数名をリファクタした。

	質問：
		- resという変数名をどう思いますか？
		- mという変数名をどう思いますか？（mapのmです）
		- wordsという変数名をどう思いますか？
*/

func groupAnagramsStep3(strs []string) [][]string {
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
