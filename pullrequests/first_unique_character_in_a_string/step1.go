//lint:file-ignore U1000 Ignore all unused code
package template

/*
時間：５分
これまでに似たような問題を何個か解いたことがあるので比較的早く解くことができた。
１回のループで解く（notRepeatedのリストを保持して重複するごとに削除していく）ことも考えたが、１回のループで倍の処理をするのであれば２回ループして処理するのと全体の仕事量としては変わらないと思い、よりシンプルでわかりやすい方法にした。

-1をreturnするのは行儀が良くないが、LeetCodeの制約がなければ返り値としてerrを返すのが一番良いと思う。該当する要素が存在しないことは正常な挙動の範疇なのでlog.Fatalはするべきではないし、ログに記録するほどのものでもないと思う。

また、「サロゲートペアや合成文字等、char の値を複数用いないと表現出来ない文字が来ないことを想定」している。（https://github.com/seal-azarashi/leetcode/pull/15#discussion_r1704423277）
*/
func firstUniqCharStep1(s string) int {
	frequency := make(map[rune]int)
	for _, r := range s {
		frequency[r]++
	}
	for i, r := range s {
		if frequency[r] == 1 {
			return i
		}
	}
	return -1
}
