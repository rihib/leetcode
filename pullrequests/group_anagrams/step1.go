//lint:file-ignore U1000 Ignore all unused code
package groupanagrams

/*
	すぐに思いついたのはソートした文字列をkeyとしてハッシュテーブルを使ってグルーピングすること。
	ただ、クイックソートの平均計算量はO(n log n)なので、O(n)でできればやりたいと思い、前に似た
	ような問題を解いたことがあったので、与えられた条件に小文字しか使われないと書いてあったので、
	配列のインデックスを用いて頻度を取得し、keyにすることにした。
*/

func groupAnagramsStep1(strs []string) [][]string {
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
