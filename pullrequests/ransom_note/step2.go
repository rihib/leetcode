//lint:file-ignore U1000 Ignore all unused code
package ransomnote

/*
レビュワーの方へ：
  - このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/
func canConstructStep2(ransomNote string, magazine string) bool {
	var frequencies [26]int
	for _, r := range magazine {
		frequencies[r-'a']++
	}
	for _, r := range ransomNote {
		frequencies[r-'a']--
		if frequencies[r-'a'] < 0 {
			return false
		}
	}
	return true
}
