//lint:file-ignore U1000 Ignore all unused code
package ransomnote

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
時間：2分
似たような問題を何回か解いたことがあるので特に困らなかった。
入力が小文字に限らない場合は、runeをkeyとするmapを使えば良い。その際、Unicodeの扱いには気をつける必要がある。
https://github.com/rihib/leetcode/pull/5#issue-2446890745
*/
func canConstructStep1(ransomNote string, magazine string) bool {
	var frequency [26]int
	for _, r := range magazine {
		frequency[r-'a']++
	}
	for _, r := range ransomNote {
		frequency[r-'a']--
		if frequency[r-'a'] < 0 {
			return false
		}
	}
	return true
}
