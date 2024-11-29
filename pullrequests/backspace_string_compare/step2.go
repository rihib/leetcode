//lint:file-ignore U1000 Ignore all unused code
package backspacestringcompare

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
後ろから見ていく方法。素直に実装している。
*/
func backspaceCompareStep2(s string, t string) bool {
	runeS, runeT := []rune(s), []rune(t)
	i, j := len(runeS)-1, len(runeT)-1
	sCount, tCount := 0, 0
	for i >= 0 || j >= 0 {
		if i >= 0 && runeS[i] == '#' {
			sCount++
			i--
			continue
		}
		if j >= 0 && runeT[j] == '#' {
			tCount++
			j--
			continue
		}
		if sCount > 0 {
			sCount--
			i--
			continue
		}
		if tCount > 0 {
			tCount--
			j--
			continue
		}
		if i >= 0 && j >= 0 && runeS[i] == runeT[j] {
			i--
			j--
			continue
		}
		return false
	}
	return true
}
