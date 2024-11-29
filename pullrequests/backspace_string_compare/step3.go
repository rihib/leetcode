//lint:file-ignore U1000 Ignore all unused code
package backspacestringcompare

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
重複をなくした実装。
*/
func backspaceCompareStep3(s string, t string) bool {
	runeS, runeT := []rune(s), []rune(t)
	i, j := len(runeS)-1, len(runeT)-1
	for i >= 0 || j >= 0 {
		i, j = nextIndex(runeS, i), nextIndex(runeT, j)
		if i >= 0 && j >= 0 && runeS[i] == runeT[j] {
			i--
			j--
			continue
		}
		if i >= 0 || j >= 0 {
			return false
		}
	}
	return true
}

func nextIndex(runeS []rune, index int) int {
	count := 0
	for {
		if index >= 0 && runeS[index] == '#' {
			count++
		} else if count > 0 {
			count--
		} else {
			break
		}
		index--
	}
	return index
}
