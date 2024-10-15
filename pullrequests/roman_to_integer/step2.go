//lint:file-ignore U1000 Ignore all unused code
package romantointeger

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
前のシンボルの方が後ろのシンボルよりも小さい場合として考えることができる。
ただ個人的にはStep1の方が要件に明確に沿っているので良いのではと思っている。
*/
func romanToIntStep2(s string) int {
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
	for i, r := range runeS {
		if i+1 < len(runeS) && symbolToInt[r] < symbolToInt[runeS[i+1]] {
			total -= symbolToInt[r]
		} else {
			total += symbolToInt[r]
		}
	}
	return total
}
