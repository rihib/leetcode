//lint:file-ignore U1000 Ignore all unused code
package longestpalindrome

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
時間：6分

１周目はカウントだけして、２周目にlengthを求めるという方法も取れると思うが、for文を２つ書くのがめんどくさかった（計算量的にはどちらも大差ないと思う）ので１周で収めるようにした。
*/
func longestPalindrome(s string) int {
	length := 0
	frequencies := make(map[rune]int, len(s))
	for _, r := range s {
		frequencies[r]++
		if frequencies[r] == 2 {
			length += 2
			delete(frequencies, r)
		}
	}
	if len(frequencies) != 0 {
		length++
	}
	return length
}
