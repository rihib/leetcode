//lint:file-ignore U1000 Ignore all unused code
package validanagram

/*
よりシンプルにわかりやすく書くようにしました。例えばfreqのように省略するのではなくfrequencyにしました。

上については、Goの場合、s[i]のようにアクセスするとバイト列にアクセスすることになるので、今回の場合はたまたまascii文字しか入力として与えられないので問題ないのですが、あまりよろしくないと思ったので文字としてアクセスするように変更しました。

下の解法についてはStep1のロジックがわかりづらかったので、より明確になるように変更しました。
*/
func isAnagram_step2(s string, t string) bool {
	var frequency [26]int
	for _, r := range s {
		frequency[r-'a']++
	}
	for _, r := range t {
		frequency[r-'a']--
	}
	for _, n := range frequency {
		if n != 0 {
			return false
		}
	}
	return true
}

func isAnagram_unicode_step2(s string, t string) bool {
	frequency := make(map[rune]int)
	for _, r := range s {
		frequency[r]++
	}
	for _, r := range t {
		frequency[r]--
	}
	for _, n := range frequency {
		if n != 0 {
			return false
		}
	}
	return true
}
