//lint:file-ignore U1000 Ignore all unused code
package validanagram

/*
かなり前に解いたものなので、詳細については忘れてしましました。
小文字のアルファベットのみが入力として与えられる場合と、Unicode文字も含まれる場合の２つの解法が含まれています。

基本的にはそれぞれの文字列に含まれる文字の種類と数が一致しているかを見ています。
下の解法について、40行目で0未満になった場合のみfalseを返すのは一見おかしいように見えますが、30行目で文字数が等しいことを確認しているため、一方より多い文字があれば何かの文字は少なくなるはずであるため、ロジック的には問題ないと思います。
*/
func isAnagramStep1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freq [26]int
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}
	for _, v := range freq {
		if v != 0 {
			return false
		}
	}
	return true
}

func isAnagramUnicodeStep1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freq := make(map[rune]int)
	for _, r := range s {
		freq[r]++
	}
	for _, r := range t {
		freq[r]--
		if freq[r] < 0 {
			return false
		}
	}
	return true
}
