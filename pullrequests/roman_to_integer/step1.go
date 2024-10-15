//lint:file-ignore U1000 Ignore all unused code
package template

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
時間：25分
方針自体はすぐに思いついたが、単純に実装に時間がかかってしまった。
*/
func romanToIntStep1(s string) int {
	symbolToInt := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := 0
	runeS := []rune(s)
	currIndex, nextIndex := 0, 1
	for currIndex < len(runeS) && nextIndex < len(runeS) {
		curr, next := runeS[currIndex], runeS[nextIndex]
		if (curr == 'I' && (next == 'V' || next == 'X')) ||
			(curr == 'X' && (next == 'L' || next == 'C')) ||
			(curr == 'C' && (next == 'D' || next == 'M')) {
			total += symbolToInt[next] - symbolToInt[curr]
			currIndex, nextIndex = currIndex+2, nextIndex+2
			continue
		}
		total += symbolToInt[curr]
		currIndex, nextIndex = currIndex+1, nextIndex+1
	}
	if currIndex < len(runeS) {
		total += symbolToInt[runeS[currIndex]]
	}
	return total
}
